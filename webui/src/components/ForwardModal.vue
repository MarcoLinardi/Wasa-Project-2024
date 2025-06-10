<script>
import axiosInstance from "@/services/axios";

export default {
  data() {
    return {
      chats: [],
    };
  },
  props: {
    message: {
      type: Object,
      required: true
    },
    originalChatId: {
      type: Number,
      required: true
    }
  },
  mounted() {
    this.loadChats();
  },
  methods: {
    async loadChats() {
      try {
        const response = await axiosInstance.get("/chats");
        this.chats = response.data.chats || [];
        this.chats.sort((a, b) => {
          const tA = a.lastMessage?.timestamp || 0;
          const tB = b.lastMessage?.timestamp || 0;
          return new Date(tB) - new Date(tA); // Ordine decrescente (più recenti prima)
        });

      } catch (e) {
        console.error(e);
      }
    },
    async forwardMessageToChat(destinationChat) {
      try {
        const originalChatId = this.originalChatId;
        const originalMessageId = this.message.messageId;
        const destinationChatId = destinationChat.chatId;

        const requestBody = {
          destinationChatId: destinationChatId,
        };
        
        await axiosInstance.post(`/chats/${originalChatId}/messages/${originalMessageId}/forward`, requestBody);

        alert(`Messaggio inoltrato con successo a "${destinationChat.name}"!`);
        this.$emit('close');

      } catch (error) {
        console.error("Errore durante l'inoltro del messaggio:", error);
        alert("Si è verificato un errore. Impossibile inoltrare il messaggio.");
      }
    },
  }
}
</script>

<template>
  <div class="forward-modal">
    <div class="modal-content">
      <h3>Inoltra messaggio a:</h3>
      <div class="chat-list">
        <ul class="chat-list">
          <li v-for="chat in chats" :key="chat.chatId" @click="forwardMessageToChat(chat)">
            {{ chat.name }}
          </li>
        </ul>
      </div>
      <button @click="$emit('close')">Annulla</button>
    </div>
  </div>
</template>

<style scoped>
.forward-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* Sfondo scuro trasparente */
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background-color: #c7c6e4;
  padding: 2rem;
  border-radius: 1rem;
  width: 90%;
  max-width: 400px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  text-align: center;
  border: 0.3rem solid navy;
}

.modal-content h3 {
  margin-bottom: 1rem;
  font-size: 1.25rem;
  color: #333333;
}

.chat-list {
  max-height: 300px;
  overflow-y: auto;
  margin-bottom: 1rem;
}

.chat-list ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.chat-list li {
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  border-bottom: 1px solid navy;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.chat-list li:hover {
  background-color: rgba(0, 0, 128, 0.208);
}

button {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 0.6rem 1.2rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.2s ease;
}

button:hover {
  background-color: #c0392b;
}

</style>
