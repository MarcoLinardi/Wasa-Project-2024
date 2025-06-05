<script>
import axiosInstance from "@/services/axios";
import ManageMember from "./ManageMember.vue";

export default {
  data() {
    return {
      isDeletingChat: false,
      isEditingName: false,
      editedName: '',
      editedPhoto: '',
      isManagingMembers: false
    };
  },
  props: {
    chat: {
      type: Object,
      required: true
    },
    loggedUserId: {
      type: Number,
      required: true
    }
  },
  component: {
    ManageMember
  },
  computed: {
    chatPhoto() {
      if (this.chat?.isGroup) {
        if (this.chat.photo.startsWith('/')) {
          return this.chat.photo
        }
        return 'data:image/jpeg;base64,' + this.chat.photo || 'images/default-group-avatar.png';
      }
      else if (this.chat.participants && this.chat.participants.length > 0) {
        const otherParticipant = this.chat.participants.find(
          participant => participant.userId !== this.loggedUserId
        );

        if (otherParticipant && otherParticipant.photo) {
          const photo = otherParticipant.photo.trim();
          const DEFAULT_USER_PHOTO = '/images/default-user-avatar.png';

          // Caso 1: è una base64 valida
          if (photo.startsWith('data:image')) {
            return photo;
          }
          // Caso 2: è un percorso (non base64)
          if (photo === DEFAULT_USER_PHOTO) {
            return photo;
          }
          // Caso 3: è una stringa base64 senza prefisso
          return `data:image/png;base64,${photo}`;
        }
        return '/images/default-user-avatar.png';
      }
    },
    chatName() {
      if (this.chat.isGroup) {
        return this.chat.name || 'Gruppo'; // Nome del gruppo
      }

      // È una chat privata, trova l'altro partecipante
      else if (this.chat.participants && this.chat.participants.length > 0) {
        const otherParticipant = this.chat.participants.find(
          participant => participant.userId !== this.loggedUserId
        );

        if (otherParticipant && otherParticipant.name) {
          return otherParticipant.name;
        }
      }
      // Fallback se il nome della chat fornito dal backend è già quello dell'altro utente
      return this.chat.name || 'Chat';
    }
  },
  methods: {
    async deleteChat(chatId) {
      try {
        console.log("Funzione deleteChat chiamata con chatId:", chatId);
        await axiosInstance.delete(`/chats/${chatId}`);
        this.$emit('chat-deleted', chatId);  // Notifica al componente genitore
      } catch (error) {
        console.error("Errore durante l'eliminazione della chat:", error);
      } 
    },
    getParticipantNames(chat) {
      if (chat && chat.isGroup && Array.isArray(chat.participants)) {
        return chat.participants.map(p => p.name).join(', ');
      }
      return '';
    },
    changeGroupName() {
      this.isEditingName = !this.isEditingName;
      this.editedName = this.chat.name;
    },

    changeGroupPhoto() {
      const fileInput = this.$refs.fileInput;
      if (fileInput) {
        fileInput.click(); // apre la finestra file
      }
      this.editedPhoto = this.chat.photo
    },

    cancelEdit() {
      this.isEditingName = false;
      this.editedName = this.chat.name;
    },

    async saveName() {
      const chatId = this.chat.chatId;
      try {
        const response = await axiosInstance.put(`/chats/${chatId}/name`, {
          newName: this.editedName
        });

        if (response.status === 200) {
          this.chat.name = this.editedName;
          this.isEditingName = false;
          console.log('Nome aggiornato con successo:', response.data);
        }
      } catch (error) {
        console.error("Errore durante l'aggiornamento del nome:", error);
        alert("Si è verificato un errore durante l'aggiornamento del nome");
      }
    },

    handlePhotoUpload(event) {
      const file = event.target.files[0];
      if (!file) return;

      // Controlla che sia PNG
      if (file.type !== "image/png") {
        alert("Sono accettate solo immagini in formato PNG.");
        return;
      }

      const reader = new FileReader();
      reader.onload = () => {
        const base64String = reader.result.split(',')[1]; // Rimuove "data:image/png;base64,"
        this.chat.photo = base64String;
        this.uploadPhoto(base64String);
      };
      reader.readAsDataURL(file);
    },

    async uploadPhoto(base64Image) {
      const chatId = this.chat.chatId;
      try {
        const response = await axiosInstance.put(`/chats/${chatId}/photo`, {
          newPhoto: base64Image
        });

        if (response.status === 200) {
          this.chat.photo = base64Image;
        }
        this.isEditingPhoto = !this.isEditingPhoto;
      } catch (error) {
        console.error("Errore durante l'upload della foto:", error);
        alert("Errore durante l'upload della foto");
      }
    },
    openMemberManager() {
      this.isManagingMembers = true;
    },
    closeMemberManager() {
      this.isManagingMembers = false;
    },
    handleCloseFromMember(payload) {
      this.isManagingMembers = false;
      this.$emit('close', payload); // Propaga a ChatArea
    },
    async leaveGroup() {
      const chatId = this.chat.chatId;
      const userId = this.loggedUserId;

      if (!chatId || !userId) {
        console.warn("Chat ID o utente loggato mancante");
        return;
      }

      try {
        const payload = { userIds: [userId] };
        console.log("Leaving group:", payload);

        const response = await axiosInstance.delete(`/chats/${chatId}/members`, {
          data: payload
        });

        console.log("Sei uscito dal gruppo:", response.data);

        // Emetti evento o naviga altrove
        this.$emit("left-group", chatId);
      } catch (error) {
        console.error("Errore durante l'uscita dal gruppo:", error);
      }
    }

  }
};
</script>

<template>
  <div class="chat-details-overlay" @click.self="$emit('close')">
    <div class="chat-details-modal">
      <button class="close-button" title="Chiudi" @click="$emit('close-modal')">
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
      </button>

      <div class="photo-container">
        <img :src="chatPhoto" alt="Foto chat" class="chat-photo" />

        <button v-if="chat.isGroup" class="edit-photo"  title="Cambia Foto" @click="changeGroupPhoto">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
        </button>
        <input
            type="file"
            accept="image/png"
            ref="fileInput"
            @change="handlePhotoUpload"
            class="photo-upload-input"
        />

      </div>

      <div class="name-container">
        <div v-if="isEditingName" class="edit-name-form">
          <input
            v-model="editedName"
            placeholder="Inserisci nuovo nome..."
            class="name-input"
            @keyup.enter="saveName"
          />
          <button @click="saveName" class="confirm-edit-button">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#check" /></svg>
          </button>
          <button @click="cancelEdit" class="cancel-edit-button">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x" /></svg>
          </button>
        </div>
        <div v-else class="name-display">
          <h3 class="chat-name"><strong>{{ chatName }}</strong></h3>
          <button v-if="chat.isGroup" class="edit-name" title="Cambia nome" @click="isEditingName = true">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit-3" /></svg>
          </button>
        </div>
      </div>

      <div v-if="this.chat.isGroup" class="member-container">
        <h6 class="participants-label">Partecipanti:</h6>
        <p class="participants-list">{{ getParticipantNames(chat) }}</p>
        <div class="member-button-container">
          <button class="edit-member-button" title="Rimuovi membro" @click="openMemberManager">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-minus" /></svg>
            <span>Rimuovi membri</span>
          </button>
          <button class="edit-member-button" title="Aggiungi membro" @click="openMemberManager">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-plus" /></svg>
            <span>Aggiungi membri</span>
          </button>
        </div>
      </div>
      
      <div v-if="this.chat.isGroup" class="leave-group-container">
        <button class="leaveGroup-button" @click="leaveGroup">
          Abbandona Gruppo
        </button>
      </div>

      <div v-else class="delete-container">
        <button class="delete-button" title="Elimina Chat" @click="isDeletingChat = !isDeletingChat">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
          <span>Elimina Chat</span>
        </button>
        <div v-if="isDeletingChat" class="edit-action">
          <button class="confirm-button" @click="deleteChat(chat.chatId)">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
          </button>
          <button class="cancel-button" @click="isDeletingChat = false">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#delete"/></svg>
          </button>
        </div>
      </div>
      <ManageMember
        v-if="isManagingMembers"
        :group-members="chat.participants"
        :chat="chat"
        @close-modal="closeMemberManager"
        @close="handleCloseFromMember"
      />
    </div>
  </div>
</template>

<style scoped>
.chat-details-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(2px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.chat-details-modal {
  background: #c7c6e4;
  padding: 2rem;
  border-radius: 1.3rem;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
  text-align: center;
  width: 30rem;               /* puoi anche renderlo max-width */
  max-height: 90vh;           /* limite massimo per schermi piccoli */
  overflow-y: auto;           /* scroll se diventa troppo lungo */
  position: relative;
  border: 2px solid navy;
}


.photo-container {
  position: relative;
  width: fit-content;
  margin: 0 auto;
}

.chat-photo {
  width: 8rem;
  height: 8rem;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid navy;
}

.edit-photo {
  position: absolute;
  bottom: 0;
  right: 1.8rem;
  transform: translateX(100%);
  background: navy;
  border: none;
  border-radius: 1rem;
  padding: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  color: antiquewhite;
  box-shadow: 0.125rem 0.125rem 0.375rem rgba(0,0,0,1),
  1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.edit-photo:hover,
.edit-name:hover,
.confirm-edit-button:hover,
.cancel-edit-button:hover {
  background-color: #6159d9;
}

.edit-name {
  background: navy;
  border: none;
  color: white;
  padding: 0.45rem;
  border-radius: 1rem;
  display: flex;
  align-items: center;
  box-shadow: 0.125rem 0.125rem 0.375rem rgba(0,0,0,1),
  1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.photo-upload-input {
  display: none;
}

.name-container {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  padding-top: 0.7rem;
}

.name-display {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.edit-name-form {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  flex-wrap: wrap;
  margin-left: auto;  /* lo spinge tutto a destra nel suo contenitore */
  margin-right: 2rem; /* oppure regoli questo */
  padding-top: 0.7rem;
  padding-bottom: 0.7rem;
}

.name-input {
  padding: 0.4rem 0.6rem;
  border: 2px solid navy;
  border-radius: 0.5rem;
  font-size: 1rem;
  width: 12rem;
}

.confirm-edit-button,
.cancel-edit-button {
  background-color: navy;
  color: white;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 1rem;
  padding: 0.5rem;
  cursor: pointer;
  box-shadow: 0.125rem 0.125rem 0.375rem rgba(0, 0, 0, 0.3);
}


.close-button {
  position: absolute;
  top: 1rem;
  right: 1rem;
  background: none;
  border: none;
  color: #6b7280;
  cursor: pointer;
  width: 2.5rem;
  height: 2.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  z-index: 10;
}

.close-button .feather {
  width: 1.5rem;
  height: 1.5rem;
}

.close-button:hover {
  color: #374151;
  background-color: rgba(0, 0, 0, 0.05);
  border-radius: 50%;
}

.member-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 0.4rem;
}

.participants-label {
  font-size: 1rem;
  font-weight: bold;
  margin: 0.5rem 0 0.2rem;
  color: navy;
}

.participants-list {
  font-size: 0.95rem;
  margin-bottom: 0.8rem;
  text-align: center;
  color: #333;
  line-height: 1.4;
  max-width: 100%;
  word-break: break-word;
}

.member-button-container {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 0.5rem;
  margin-top: 0.4rem;
}

.edit-member-button {
  background: navy;
  border: none;
  border-radius: 0.6rem;
  padding: 0.6rem 1rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: antiquewhite;
  box-shadow:
    0.125rem 0.125rem 0.375rem rgba(0, 0, 0, 1),
    1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.edit-member-button:hover {
  background-color: #6159d9;
}


.delete-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
  margin-top: 1rem;
}

.edit-action {
  display: flex;
  flex-direction: row;
  width: 100%;
  align-items: center;
  margin-top: 0;
}

.leave-group-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
  margin-top: 1rem;
}

.leaveGroup-button {
  background: #7a1e1e;
  border: none;
  border-radius: 0.6rem;
  padding: 0.6rem 1rem;
  margin-bottom: -0.3rem;
  text-align: left;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  color: white;
  box-shadow:
    0.125rem 0.125rem 0.375rem rgba(0, 0, 0, 1),
    1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.leaveGroup-button:hover {
  background-color: #c0392b;
}

/* Bottone principale */
.delete-button {
  background: navy;
  border: none;
  border-radius: 0.6rem;
  padding: 0.6rem 1rem;
  margin-bottom: -0.3rem;
  text-align: left;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  color: antiquewhite;
  box-shadow:
    0.125rem 0.125rem 0.375rem rgba(0, 0, 0, 1),
    1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.delete-button:hover {
  background-color: #6159d9;
}


/* Contenitore dei bottoni di conferma */
.edit-action {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
}

/* Bottone conferma */
.confirm-button {
  background: darkred;
  border: none;
  border-radius: 10px;
  padding: 10px 16px;
  text-align: left;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  color: antiquewhite;
  box-shadow:
    0.125rem 0.125rem 0.375rem rgba(0, 0, 0, 1),
    1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.confirm-button:hover {
  background-color: #8b0000;
}

/* Bottone annulla */
.cancel-button {
  background: #444;
  border: none;
  border-radius: 10px;
  padding: 10px 16px;
  text-align: left;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  color: antiquewhite;
  box-shadow:
    0.125rem 0.125rem 0.375rem rgba(0, 0, 0, 1),
    1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.cancel-button:hover {
  background-color: #2b2b2b;
}

</style>