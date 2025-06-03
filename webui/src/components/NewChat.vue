<script>
export default {
  props: {
    user: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      newMessage: ''
    };
  },
  methods: {
    getUserPhoto(photo) {
        // Se è una path (foto di default)
        if (photo.startsWith('/')) {
          return photo;
        }
        // Altrimenti è una base64 pura
        return 'data:image/jpeg;base64,' + photo;
    },
    sendMessage() {
      if (this.newMessage.trim() === '') {
        alert("Il messaggio non può essere vuoto.");
        return;
      }
      this.$emit('send-new-message', this.newMessage);
      this.newMessage = '';
    },
    closeChat() {
      this.$emit('close-chat-view');
    }
  },
  computed: {
    userPhotoUrl() {
      return this.getUserPhoto(this.user.photo);
    }
  }
}
</script>

<template>
  <div class="new-chat-component">
    <div class="chat-header">
      <img 
        :src="userPhotoUrl" 
        class="user-avatar"
        :alt="'Avatar di ' + user.name" 
      />
      <div class="user-info">
        <h3>{{ user.name }}</h3>
      </div>
      <button class="close-button" title="Chiudi" @click="closeChat">
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
      </button>
    </div>

    <div class="chat-body">
      <p class="chat-body-placeholder">
        Invia un messaggio per iniziare una nuova conversazione con <strong>{{ user.name }}</strong>
      </p>
      </div>

    <div class="message-input-area">
      <form @submit.prevent="sendMessage" class="message-form">
        <input 
          type="text" 
          v-model="newMessage" 
          placeholder="Scrivi un messaggio..."
          class="message-input"
          aria-label="Nuovo messaggio"
        />
        <button 
          type="submit" 
          class="send-button"
          :disabled="!newMessage.trim()"
        >
          Invia
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped>
.new-chat-component {
  display: flex;
  flex-direction: column;
  height: 83%;
  width: 80%; 
  background-color: #ffffff;
  border: 2px solid navy;
  border-radius: 1rem;
  overflow: hidden;
  overflow-y: auto;
}

/* Header della Chat */
.chat-header {
  display: flex;
  align-items: center;
  padding: 0.6rem;
  border-bottom: 1px solid #e5e7eb;
}

.user-avatar {
  width: 3.2rem; 
  height: 3.2rem;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 0.75rem;
}

.user-info {
  flex-grow: 1;
}

.user-info h3 {
  font-size: 1.125rem;
  font-weight: 600;
  margin: 0;
}

.close-button {
  background: none;
  border: none;
  color: #6b7280;
  cursor: pointer;
  padding: 0.5rem;
  line-height: 1;
  width: 2.5rem;
  height: 2.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-button .feather {
  width: 1.5rem;
  height: 1.5rem;
}


.close-button:hover {
  color: #374151;
}

/* Corpo della Chat */
.chat-body {
  flex-grow: 1;
  
  background-color: #f9fafb;
}

.chat-body-placeholder {
  font-size: 0.875rem;
  color: #4b5563;
  font-style: italic;
  text-align: center;
  margin-top: 1rem;
}

/* Area Input Messaggio */
.message-input-area {
  padding: 1rem;
  border-top: 1px solid #e5e7eb;
  background-color: #ffffff;
}

.message-form {
  display: flex;
  align-items: center;
}

.message-input {
  flex-grow: 1;
  padding: 0.75rem;
  border: 1px solid #d1d5db;
  border-radius: 0.5rem;
  font-size: 1rem;
  margin-right: 0.5rem;
}

.message-input:focus {
  outline: none;
  border-color: navy;
}

.send-button {
  background-color: navy;
  color: white;
  font-weight: 600;
  padding: 0.75rem 1.25rem;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: background-color 0.15s ease-in-out;
}

.send-button:hover {
  background-color: navy;
}

.send-button:disabled {
  background-color: rgb(77, 77, 187);
  cursor: not-allowed;
}
</style>