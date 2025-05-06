<script>
import { sendMessage, deleteChat, createPrivateChat } from "@/services/api"; 
export default {
  props: {
    selectedChat: Object,
    selectedUser: Object,
  },
  data() {
    return {
      messageText: "",
      messages: []
    };
  },
  methods: {

    async loadMessages(chatId) {
      try {
        const response = await axiosInstance.get(`/chats/${chatId}/messages/status`);
        this.messages = response.data;
      } catch (e) {
        console.error(e);
      }
    },

    async handleFirstMessage({ user, content }) {
      try {
        console.log("User della funzione handleFirstMessage: " + this.selectedUser.name)
        const newChatId = await createPrivateChat(this.selectedUser); // crea la chat
        try {
          await sendMessage(newChatId, content); // invia il primo messaggio

          // aggiorna frontend
          this.selectedChat = {
            chatId: newChatId,
            name: user.name,
            isGroup: false
          };
          this.selectedUser = null;
          await this.loadMessages(newChatId);
          this.chatList.push(this.selectedChat);

        } catch (errorMessage) {
          await deleteChat(newChatId); // cancella la chat se errore invio
          console.error("Errore invio primo messaggio, chat eliminata:", errorMessage);
        }

      } catch (e) {
        console.error("Errore durante creazione chat:", e);
      }
    },

    handleSend() {
      const text = this.messageText.trim();
      if (!text) return;

      console.log("Funzione handleSend: " + this.selectedUser.name)
      if (this.selectedChat) {
        console.log("Selected Chat: " + this.selectedChat.chatId)
        // Se la chat esiste già, invia il messaggio normalmente
        this.$emit("handleSendMessage", { chatId: this.selectedChat.chatId, content: text });

      } else if (this.selectedUser && !this.selectedChat) {
        // Se non c'è chat ma c'è un utente selezionato, creala
        this.handleFirstMessage(this.selectedUser, text );
      }
      this.messageText = ""; // svuota campo input
      },
    }
  }
</script>

<template>
  <div class="chat-area">
    
    <!-- HEADER -->
    <div class="chat-header">
      <h2 v-if="selectedChat">{{ selectedChat.name }}</h2>
      <h2 v-else-if="selectedUser">{{ selectedUser.name }}</h2>
      <h2 v-else>Nessuna chat selezionata</h2>
    </div>

    <!-- BODY -->
    <div class="chat-body">
      <template v-if="messages && messages.length > 0">
        <div v-for="message in messages" :key="message.id" class="message">
          {{ message.content }}
        </div>
      </template>
      <template v-else-if="selectedUser && !selectedChat">
        <div class="preview-message">
          Inizia una nuova chat con <strong>{{ selectedUser.name }}</strong>
        </div>
      </template>
      <template v-else>
        <div class="preview-message">
          Seleziona una chat per iniziare
        </div>
      </template>
    </div>

    <!-- FOOTER sempre visibile -->
    <div class="chat-footer" v-if="selectedChat || selectedUser">
      <div class="input-container">
        <input
          v-model="messageText"
          placeholder="Scrivi un messaggio..."
          @keyup.enter="handleSend"
        />
        <button class="send-button" @click="handleSend">
          <svg xmlns="http://www.w3.org/2000/svg" height="24" width="24" viewBox="0 0 24 24">
            <path fill="#ffffff" d="M2 21l21-9L2 3v7l15 2-15 2z"/>
          </svg>
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
  background-color: rgb(210, 180, 140);
  overflow: hidden;
}

.chat-header {
  height: 80px;
  background-color: #f0f2f5;
  border-bottom: 1px solid #ddd;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 1rem;
  box-sizing: border-box;
  flex-shrink: 0;
}

.chat-header h2 {
  font-size: 1.2rem;
  margin: 0;
  color: #333;
  white-space: nowrap;      /* evita spezzamenti strani */
  overflow: hidden;         /* evita sbordamenti */
  text-overflow: ellipsis;  /* aggiunge "..." se troppo lungo */
}

.chat-body {
  flex-grow: 1; 
  overflow-y: auto;
  background-color: rgb(210, 180, 140);
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 1rem;
}

.message {
  background-color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  margin-bottom: 0.5rem;
  max-width: 60%;
  word-wrap: break-word;
}

.chat-body.preview {
  display: flex;
  justify-content: center;
  align-items: center;
  color: #333;
  font-size: 1.2rem;
  font-weight: 500;
  background-color: rgb(210, 180, 140);
}

.preview-message {
  background-color: transparent;
  padding: 1.5rem 1.5rem;
  align-items: center;
}


.chat-footer {
  background-color: rgb(221, 172, 116);
  padding: 0.5rem;
  border-top: 1px solid #ddd;
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
  max-width: 800px;
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
  background-color: transparent;
  border: none;
  border-radius: 50%;
  padding: 0.6rem;
  margin-left: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.3s ease;
}

.send-button:hover {
  background-color: rgb(224, 160, 100);
  transition: background-color 0.3s ease;
}

.send-button svg {
  width: 24px;
  height: 24px;
}


</style>
