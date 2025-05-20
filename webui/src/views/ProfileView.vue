<script>
import axiosInstance from '@/services/axios.js';

export default {
  data() {
    return {
      loggedUser: {
        name: '',
        photo: ''
      },
      editedName: '',
      isEditingName: false,
      isEditingPhoto: false,
      editedPhoto: ''
    };
  },
  mounted() {
    this.loadLoggedUser();
  },
  methods: {
    loadLoggedUser() {
      const userData = localStorage.getItem('user');
      if (userData) {
        try {
          this.loggedUser = JSON.parse(userData);
          console.log("Utente loggato caricato:", this.loggedUser);
        } catch (e) {
          console.error("Errore nel parsing dei dati utente:", e);
        }
      } else {
        console.warn("Nessun utente trovato in localStorage");
      }
    },

    changeUserName() {
      this.isEditingName = !this.isEditingName;
      this.editedName = this.loggedUser.name;
    },

    changeUserPhoto() {
      this.isEditingPhoto = !this.isEditingPhoto;
      this.editedPhoto = this.loggedUser.photo
    },

    cancelEdit() {
      this.isEditingName = false;
      this.editedName = this.loggedUser.name; // ripristina il nome precedente
    },

    async saveName() {
      try {
        const response = await axiosInstance.put('/user/name', {
          newName: this.editedName
        });

        if (response.status === 200) {
          this.loggedUser.name = this.editedName;
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

      // Limite di 1MB 
      const maxSize = 1 * 1024 * 1024;
      if (file.size > maxSize) {
        alert("L'immagine non può superare 1MB.");
        return;
      }

      const reader = new FileReader();
      reader.onload = () => {
        const base64String = reader.result.split(',')[1]; // Rimuove "data:image/png;base64,"
        this.loggedUser.photo = base64String;
        this.uploadPhoto(base64String);
      };
      reader.readAsDataURL(file);
    },

    async uploadPhoto(base64Image) {
      try {
        const response = await axiosInstance.put('/user/photo', {
          photo: base64Image
        });

        if (response.status === 200) {
          this.loggedUser.photo = base64Image;
        }
        const updatedUser = { ...this.loggedUser, photo: this.loggedUser.photo };
        localStorage.setItem('user', JSON.stringify(updatedUser));
      } catch (error) {
        console.error("Errore durante l'upload della foto:", error);
        alert("Errore durante l'upload della foto");
      }
    },
    getUserPhoto(photo) {
      // Se è una path (foto di default)
      if (photo.startsWith('/')) {
        return photo;
      }
      // Altrimenti è una base64 pura
      return 'data:image/jpeg;base64,' + photo;
    }
  }
};
</script>

<template>
  <div class="profile-view-container">
    <div class="profile-card">
      <div class="profile-image-container" v-if="loggedUser">
        <img :src="getUserPhoto(loggedUser.photo)" class="profile-image" />
      </div>

      <h2 class="user-name">{{ loggedUser.name}}</h2>

      <div class="action-buttons">
        <button class="edit-button" @click="changeUserPhoto">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
          <span>Cambia Foto</span>
        </button>
        <div v-if="isEditingPhoto" class="edit-photo-form">
          <input
            type="file"
            accept="image/png"
            @change="handlePhotoUpload"
            class="photo-upload-input"
          />
        </div>
        <button class="edit-button" @click="changeUserName">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit-3"/></svg>
          <span>Cambia Username</span>
        </button>
      </div>
      <!-- BLOCCO PER INSERIRE NUOVO USERNAME -->
      <div v-if="isEditingName" class="edit-username-form">
        <input
          type="text"
          v-model="editedName"
          placeholder="Inserisci nuovo username"
          class="username-input"
        />
        <div class="edit-actions">
          <button class="confirm-button" @click="saveName">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#check"/></svg>
            <span>Conferma</span>
          </button>
          <button class="cancel-button" @click="cancelEdit">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
            <span>Annulla</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-view-container {
  background-image: url('@/assets/AppBackground.jpg');
  background-size: cover;         /* riempie tutto */
  background-position: center;    /* centrato */
  background-repeat: no-repeat;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  padding: 20px;
  box-sizing: border-box;
}

.profile-card {
  background-color: rgb(220, 212, 248);
  padding: 30px 40px;
  border-radius: 12px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 50%;
  border: 4px solid navy;

}

.profile-image-container {
  margin-bottom: 20px;
  width: 150px;
  height: 150px;
  border-radius: 50%;
  overflow: hidden;
  margin-left: auto;
  margin-right: auto;
  border: 4px solid navy;
  background-color: #e0e0e0;
}

.profile-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.profile-image-placeholder {
  width: 100%;
  height: 100%;
}

.user-name {
  font-size: 1.8em;
  font-weight: 600;
  color: black;
  margin-top: 0;
  margin-bottom: 25px;
}

.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  align-items: center;
}

.edit-button {
  background: navy;
  border: none;
  border-radius: 10px;
  padding: 10px;
  text-align: left;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px; /* Spazio tra icona e testo */
  color: antiquewhite;
  box-shadow: 6px 6px 10px rgba(0,0,0,1),
  1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.edit-button:hover {
  background-color: #6159d9;
  transform: translateY(-1px);
}

.edit-button:active {
  background-color: #6159d9;
  transform: translateY(0);
}

.edit-username-form {
  margin-top: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
}

/* Input per il nuovo nome */
.username-input {
  padding: 10px 15px;
  border-radius: 10px;
  border: 1px solid navy;
  outline: none;
  width: 60%;
  max-width: 300px;
  font-size: 16px;
  background-color: whitesmoke;
  box-shadow: inset 1px 1px 4px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.2s ease-in-out;
}

.username-input:focus {
  box-shadow: 0 0 0 2px #6159d9;
}

/* Contenitore bottoni */
.edit-actions {
  display: flex;
  gap: 10px;
}

/* Stile condiviso per entrambi i bottoni */
.confirm-button,
.cancel-button {
  padding: 8px 16px;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s;
  box-shadow: 4px 4px 10px rgba(0, 0, 0, 0.2);
  gap: 8px;
  padding: 10px;
  text-align: left;
  display: flex;
  align-items: center;
}

/* Bottone conferma */
.confirm-button {
  background-color: #6159d9;
  color: white;
}

.confirm-button:hover {
  background-color: #514ac7;
  transform: translateY(-1px);
}

/* Bottone annulla */
.cancel-button {
  background-color: #ccc;
  color: #333;
}

.cancel-button:hover {
  background-color: #bbb;
  transform: translateY(-1px);
}

.photo-upload-input {
  margin-top: 10px;
  padding: 5px;
}

</style>
