<script>
export default {
  props: {
    chat: {
      type: Object,
      required: true,
    },
    loggedUser: {
      type: Object,
      required: true
    }
  },
  computed: {
    formattedTimestamp() {
      if (!this.chat.lastMessage || !this.chat.lastMessage.timestamp) {
        return '';
      }
      const messageDate = new Date(this.chat.lastMessage.timestamp);
      const now = new Date();

      // Formato "HH:MM" se è oggi
      if (
        messageDate.getDate() === now.getDate() &&
        messageDate.getMonth() === now.getMonth() &&
        messageDate.getFullYear() === now.getFullYear()
      ) {
        return messageDate.toLocaleTimeString('it-IT', { hour: '2-digit', minute: '2-digit' });
      }

      // Formato "ieri" se è ieri
      const yesterday = new Date(now);
      yesterday.setDate(now.getDate() - 1);
      if (
        messageDate.getDate() === yesterday.getDate() &&
        messageDate.getMonth() === yesterday.getMonth() &&
        messageDate.getFullYear() === yesterday.getFullYear()
      ) {
        return 'ieri';
      }

      // Formato "dd/mm/yyyy" per date più vecchie
      return messageDate.toLocaleDateString('it-IT');
    },
    displayPhoto() {
      if (this.chat.isGroup) {
        if (this.chat.photo.startsWith('/')) {
          return this.chat.photo
        }
        return 'data:image/jpeg;base64,' + this.chat.photo || 'images/default-group-avatar.png'; // Fallback per foto di gruppo
      }
      // È una chat privata, trova l'altro partecipante
      if (this.chat.participants && this.chat.participants.length > 0) {
        const otherParticipant = this.chat.participants.find(
          participant => participant.userId !== this.loggedUser.userId
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
          // Caso 3: è una stringa base64 *senza* prefisso
          return `data:image/png;base64,${photo}`;
        }
        return '/images/default-user-avatar.png'; // Fallback definitivo
      }
      return '/images/default-user-avatar.png';
    },

    displayName() {
      if (this.chat.isGroup) {
        return this.chat.name || 'Gruppo'; // Nome del gruppo
      }

      // È una chat privata, trova l'altro partecipante
      if (this.chat.participants && this.chat.participants.length > 0) {
        const otherParticipant = this.chat.participants.find(
          participant => participant.userId !== this.loggedUser.userId
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
    selectChat() {
      console.log('Chat selezionata:'+ this.chat.name);
      this.$emit('select-chat', this.chat);
    },
  }
};
</script>

<template>
  <div class="chat-item" @click="selectChat">
    <div class="chat-avatar">
      <img :src="displayPhoto" :alt="displayName" class="avatar-img" />
    </div>
    <div class="chat-content">
      <div class="chat-header">
        <span class="chat-name">{{ displayName }}</span>
        <span class="chat-timestamp">{{ formattedTimestamp }}</span>
      </div>
      <div class="chat-last-message" v-if="chat.lastMessage">
        <span class="message-text">{{ chat.lastMessage.content }}</span>
      </div>
    </div>
    </div>
</template>

<style scoped>
.chat-item {
  display: flex;
  align-items: center;
  padding: 0.7rem 1rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
  color: #444;
  font-family: Arial, sans-serif;
  background-color: rgba(0, 0, 128, 0.100);
}

.chat-item:hover {
  background-color: rgba(0, 0, 128, 0.300);
  color: white;
}

.chat-avatar {
  position: relative;
  margin-right: 15px;
  flex-shrink: 0;
}

.avatar-img {
  width: 3.2rem;
  height: 3.2rem;
  border-radius: 50%;
  object-fit: cover;
}

.chat-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.chat-name {
  color: black;
  font-size: 1em;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.chat-timestamp {
  font-size: 0.85em;
  color: rgba(0, 0, 0, 0.700);
  white-space: nowrap;
}

.chat-last-message {
  font-size: 0.9em;
  color: rgba(0, 0, 0, 0.700);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>