<script>
import axiosInstance from "@/services/axios";
import Message from "./Message.vue";
import ChatInfo from "./ChatInfo.vue";

export default {
  props: {
    selectedChat: {
      type: Object,
      required: true
    },
    loggedUser: {
      type: Object,
      required: true
    }
  },
  component: {
    Message,
    ChatInfo
  },
  data() {
    return {
      messageText: "",
      messages: [],
      showInfo: false
    };
  },
  watch: {
    selectedChat: {
      handler(newChat, oldChat) {
        if (newChat && newChat.chatId && (!oldChat || newChat.chatId !== oldChat.chatId)) {
          this.loadMessages();
        } else if (!newChat) {
          this.messages = [];
        }
      },
      immediate: true
    }
  },
  computed: {
    otherParticipant() {
      if (this.selectedChat && !this.selectedChat.isGroup) {
        return this.selectedChat.participants.find(
          participant => participant.userId !== this.loggedUser.userId
        );
      }
      return null;
    },
    chatPhoto() {
      if (this.selectedChat?.isGroup) {
        if (this.selectedChat.photo.startsWith('/')) {
          return this.selectedChat.photo
        }
        return 'data:image/jpeg;base64,' + this.selectedChat.photo;
      }
      else if (this.selectedChat.participants && this.selectedChat.participants.length > 0) {
        const otherParticipant = this.selectedChat.participants.find(
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
          // Caso 3: è una stringa base64 senza prefisso
          return `data:image/png;base64,${photo}`;
        }
        return '/images/default-user-avatar.png';
      }
      return;
    },
    chatName() {
      if (this.selectedChat.isGroup) {
        return this.selectedChat.name || 'Gruppo'; // Nome del gruppo
      }

      // È una selectedChat privata, trova l'altro partecipante
      else if (this.selectedChat.participants && this.selectedChat.participants.length > 0) {
        const otherParticipant = this.selectedChat.participants.find(
          participant => participant.userId !== this.loggedUser.userId
        );

        if (otherParticipant && otherParticipant.name) {
          return otherParticipant.name;
        }
      }
      // Fallback se il nome della selectedChat fornito dal backend è già quello dell'altro utente
      return this.selectedChat.name || 'Chat';
    },
    participantNames() {
      const chat = this.selectedChat;
      if (chat && chat.isGroup && Array.isArray(chat.participants)) {
        return chat.participants.map(p => p.name).join(', ');
      }
      return '';
    }
  },

  methods: {
    async loadMessages() {
      try {
        const response = await axiosInstance.get(`/chats/${this.selectedChat.chatId}/messages`);
        this.messages = response.data || [];
        console.log("Messaggi caricati: " + JSON.stringify(this.messages, null, 2));

      } catch (e) {
        console.error(e);
      }
    },
    async handleSend() {
      if (!this.messageText.trim() || !this.selectedChat || !this.selectedChat.chatId) {
        console.warn("Nessun messaggio da inviare o nessuna chat selezionata.");
        return;
      }
      const textContentForMessage = this.messageText.trim();
      const messagePayload = {
        senderId: this.loggedUser.userId,
        content: textContentForMessage,
      };
      const chatId = this.selectedChat.chatId;

      try {
        const response = await axiosInstance.post(`/chats/${chatId}/messages`, messagePayload);

        if (response.data) {
          const newMessage = {
            id: response.data.id || Date.now(),
            content: textContentForMessage,
            senderId: messagePayload.senderId,
            timestamp: response.data.timestamp || new Date().toISOString(),
          };

          if (typeof newMessage.content !== 'string' || newMessage.content.trim() === '') {
            this.loadMessages();
          } else {
            this.messages.push(newMessage);
            this.$nextTick(() => {
              const chatBody = this.$refs.chatBody;
              if (chatBody) {
                chatBody.scrollTop = chatBody.scrollHeight;
              }
            });
          }
          this.$emit('message-sent', { chatId, newMessage });
        } else {
          this.loadMessages();
        }
        this.messageText = "";
        this.loadMessages();
        
      } catch (e) {
        console.error('[handleSend] Errore durante invio del messaggio:', e);
      }
    },
    handleChatDeleted(chatId) {
      this.showInfo = false;
      this.$emit('chat-deleted', chatId);
    },
    handleCloseMemberManager({ added, removed }) {
      if (!this.selectedChat || !Array.isArray(this.selectedChat.participants)) {
        console.warn("chat o chat.participants non disponibile");
        return;
      }

      // Aggiungi nuovi utenti
      for (const user of added) {
        if (!this.selectedChat.participants.some(p => p.userId === user.userId)) {
          this.selectedChat.participants.push(user);
        }
      }

      // Rimuovi utenti (se implementato)
      const removedIds = removed.map(u => u.userId);
      this.selectedChat.participants = this.selectedChat.participants.filter(
        p => !removedIds.includes(p.userId)
      );
    },
    handleLeaveGroup(chatId) {
      this.$emit("left-group", chatId);
    },
    async deleteMessage(message) {
      const chatId = this.selectedChat.chatId;
      const messageId = message.messageId;
      try {
        await axiosInstance.delete(`/chats/${chatId}/messages/${messageId}`);
        this.loadMessages();
      } catch (e) {
        console.error("Errore eliminazione:", e);
      }
    }

    }
  }
</script>

<template>
  <div class="chat-area">

    <div class="chat-header" v-if="selectedChat" @click="showInfo = true">
      <img :src="chatPhoto" class="chat-photo" :alt="selectedChat.name" />
      <div class="chat-info">
        <h2>{{ chatName }}</h2>
        <h4 v-if="selectedChat.isGroup" class="participants">
          {{ participantNames }}
        </h4>
      </div>
    </div>
    <ChatInfo
      v-if="showInfo"
      :chat="selectedChat"
      :loggedUserId="loggedUser.userId"
      @close-modal="showInfo = false"
      @chat-deleted="handleChatDeleted"
      @close="handleCloseMemberManager"
      @left-group="handleLeaveGroup"
    />

    <div class="chat-body" ref="chatBody">
      
      <div v-if="selectedChat && messages.length === 0" class="no-messages">
        Nessun messaggio in questa chat. Inizia tu la conversazione!
      </div>
      <Message
        v-for="message in messages || []"
        :key="message.id"
        :message="message"
        :loggedUser="loggedUser"
        :chat="selectedChat"
        @reload-messages="loadMessages"
        @delete="deleteMessage"
      />
      
    </div>

    <div class="chat-footer" v-if="selectedChat">
      <div class="input-container">
        <input
          v-model="messageText"
          placeholder="Scrivi un messaggio..."
          @keyup.enter="handleSend"
        />
        <button class="send-button" title="Scegli foto">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
        </button>
        <button class="send-button" title="Invia" @click="handleSend" :disabled="!messageText.trim()">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#send"/></svg>
        </button>
      </div>
    </div>

  </div>
</template>

<style scoped>
.chat-area {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  height: 100%;
  background-color: rgb(220, 212, 248);
  overflow: hidden;
}

.chat-header {
  height: 5rem;
  background-color: #c7c6e4;
  border-bottom: 0.1rem solid navy;
  display: flex;
  align-items: center;
  padding: 0 2rem;
  box-sizing: border-box;
  flex-shrink: 0;
}

.chat-photo {
  width: 4rem;
  height: 4rem;
  object-fit: cover;
  border-radius: 50%;
  margin-right: 1rem;
}

.chat-info {
  display: flex;
  flex-direction: column;
  justify-content: center;
  overflow: hidden;
}

.chat-info h2 {
  font-size: 1.3rem;
  margin: 0;
  color: #333;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 1.4;
}

.chat-info .participants {
  font-size: 0.9rem;
  margin: 0.3rem 0 0 0;
  color: #222;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  font-weight: normal;
}

.chat-body {
  flex-grow: 1; 
  flex-direction: column;
  overflow-y: auto;
  background-color: whitesmoke;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  padding: 1rem;
}

.chat-footer {
  background-color: #c7c6e4;
  padding: 0.5rem;
  width: 100%;
  border-top: 2px solid navy;
  display: flex;
  align-items: center;
  justify-content: center;
  position: sticky;
  bottom: 0;
  z-index: 2;
}

.input-container {
  background-color: transparent;
  display: flex;
  align-items: center;
  width: 100%;
  padding: 0.3rem 0.8rem;
  border-radius: 30px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.input-container input {
  border: none;
  outline: none;
  flex-grow: 1;
  padding: 0.6rem;
  font-size: 1rem;
  background: transparent;
  color: #333;
}

.send-button {
  border: none;
  border-radius: 50%;
  height: 2.5rem;
  width: 2.5rem;
  background-color: transparent;
}

.send-button:hover {
  background-color: navy;
  color: rgb(220, 212, 248);
}

.send-button svg {
  transition: transform 0.2s ease;
}

.send-button:hover svg {
  transform: scale(1.25);
}
</style>
