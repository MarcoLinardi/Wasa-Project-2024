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
      users: [],
      showInfo: false,
      replyMessage: null,
    };
  },
  mounted() {
    this.loadUsers()
  },
  watch: {
    selectedChat: {
      async handler(newChat, oldChat) {
        if (newChat && newChat.chatId && (!oldChat || newChat.chatId !== oldChat.chatId)) {
          // Aspetta finché gli utenti non sono stati caricati
          while (this.users.length === 0) {
            await new Promise(resolve => setTimeout(resolve, 100)); // aspetta 100ms
          }
          await this.loadMessages();
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
      return '/images/default-user-avatar.png';
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
        const rawMessages = response.data || [];
        // Risolvi i replyTo
        this.messages = rawMessages.map(msg => {
          if (msg.replyToId) {
            const repliedTo = rawMessages.find(m => m.messageId === msg.replyToId);
            if (repliedTo) {
              console.log("Replied message trovato:", repliedTo);
              console.log("Sender ID da cercare:", repliedTo.senderId);
              console.log("Lista utenti:", this.users);
              console.log("Nome trovato:", this.getUserNameById?.(repliedTo.senderId));
              msg.replyTo = {
                content: repliedTo.content,
                senderName: this.getUserNameById?.(repliedTo.senderId) || "Utente"
              };
            }
          }
          return msg;
        });
      } catch (e) {
        console.error("[loadMessages] Errore nel caricamento dei messaggi:", e);
      }
    },
    getUserNameById(userId) {
      const user = this.users?.find(u => u.userId === userId);
      return user?.name || "Utente";
    },
    async loadUsers() {
      try {
        const response = await axiosInstance.get("/users");
        this.users = response.data.users || response.data;
        this.users.push(this.loggedUser);
        this.users.sort((a, b) => (a.name || "").localeCompare(b.name || ""));
      } catch (e) {
        console.error("[loadUsers] Errore nel caricamento degli utenti:", e);
        this.users = [];
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
        ...(this.replyMessage && { replyToId: this.replyMessage.messageId })
      };
      const chatId = this.selectedChat.chatId;

      try {
        const response = await axiosInstance.post(`/chats/${chatId}/messages`, messagePayload);

        if (response.data) {
          const newMessage = {
            messageId: response.data.id || Date.now(),
            content: textContentForMessage,
            senderId: messagePayload.senderId,
            timestamp: response.data.timestamp || new Date().toISOString(),
            ...(this.replyMessage && { replyToId: this.replyMessage })
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
        this.replyMessage = null;
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
      this.$emit("update-members", { added, removed });
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
    },
    handleMessageReactionUpdate({ messageId, reaction }) {
      const updatedMessages = this.messages.map(msg => {
        if (msg.messageId === messageId) {
          return { ...msg, reaction };
        }
        return msg;
      });
      this.messages = updatedMessages;
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
      @update-chat-name="$emit('update-chat-name', $event)"
      @update-chat-photo="$emit('update-chat-photo', $event)"
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
        @reply="replyMessage = $event"
        @update-message-reaction="handleMessageReactionUpdate"
      />
      
    </div>

    <div v-if="replyMessage" class="reply-preview">
      <div class="reply-header">
        Rispondi a:
        <button class="close-reply" @click="replyMessage = null">✕</button>
      </div>
      <div class="reply-content">{{ replyMessage.content }}</div>
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
  height: 100vh;
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

.reply-preview {
  background-color: rgba(0, 0, 128, 0.32);
  padding: 0.5rem;
  border-left: 3px solid navy;
  margin: 0 auto;
  width: 95%;
  display: flex;
  flex-direction: column;
  border-radius: 0.4rem;
}
.reply-header {
  display: flex;
  justify-content: space-between;
}
.close-reply {
  background: none;
  border: none;
  font-size: 1rem;
  cursor: pointer;
}
.reply-content {
  font-style: italic;
  color: black;
}

</style>
