<script>
import axiosInstance from "@/services/axios";
import ChatItem from "../components/ChatItem.vue";
import ChatArea from "../components/ChatArea.vue";

export default {
  data() {
    return {
      chats: [],
      chatName: "",
      selectedChat: null,
    }
  },
  components: {
    ChatItem,
    ChatArea
  },
  mounted() {
    this.loadChats();
    this.loadLoggedUser();
    this.$watch(
      () => this.chats,
      (newChats) => {
        const selectedChatId = this.$route.query.selectedChatId;
        if (selectedChatId && newChats.length > 0) {
          const foundChat = newChats.find(chat => chat.chatId === selectedChatId);
          if (foundChat) {
            this.handleChatSelected(foundChat);
          }
        }
      },
      { immediate: true }
    );
  },
  computed: {
    filteredChats() {
      if (!this.chatName) {
        return this.chats;
      }
      const searchTerm = this.chatName.toLowerCase();
      return this.chats.filter(chat => {
        return chat.name && chat.name.toLowerCase().includes(searchTerm);
      });
    }
  },
  watch: {
    chats: {
      immediate: true,
      handler(newChats) {
        const selectedChatId = this.$route.query.selectedChatId;

        if (selectedChatId && newChats.length > 0) {
          const foundChat = newChats.find(
            chat => chat.chatId.toString() === selectedChatId
          );

          if (foundChat) {
            this.handleChatSelected(foundChat);
          }
        }
      }
    }
  },
  methods: {
    loadLoggedUser() {
      const userData = localStorage.getItem('user');
      if (userData) {
        this.loggedUser = JSON.parse(userData);
        console.log("Utente loggato caricato:", this.loggedUser);
      } else {
        console.warn("Nessun utente trovato in localStorage");
      }
    },
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
    handleChatSelected(chat) {
      this.selectedChat = chat;
    },
    handleMessageSent() {
      this.loadChats();
    },
    handleChatDeleted(chatId) {
      // 1. Rimuovi la chat dall'array reattivo
      this.chats = this.chats.filter(chat => chat.chatId !== chatId);

      // 2. Se era la chat selezionata, deseleziona
      if (this.selectedChat && this.selectedChat.chatId === chatId) {
        this.selectedChat = null;
      }
    }
  }
}
</script>


<template>
  <div class="home-container">
    <div class="sidebar-wrapper">
      <!-- Lista Chat -->
      <div class="chat-list">
        <h3>Le mie chat</h3>
        <div class="search-chat">
          <input type="text" id="chatName" v-model="chatName" placeholder="Cerca chat">
        </div>
        <div class="chats-list-container">
          <p v-if="!filteredChats || filteredChats.length === 0" class="no-users-message">
            Nessuna chat disponibile o nessun risultato per la ricerca...
          </p>
          <div class="chat-list" v-else>
            <li v-for="chat in filteredChats" :key="chat.chatId" class="chat-item">
              <ChatItem 
                :chat="chat"
                :loggedUser="loggedUser"
                @select-chat="handleChatSelected"
              />
            </li>
          </div>
        </div>
      </div>
    </div>
    <!-- Chat Area -->
    <div v-if="selectedChat" class="chat-wrapper">
      <ChatArea
        :selectedChat="selectedChat"
        :loggedUser="loggedUser"
        @select-chat="handleChatSelected"
        @message-sent="handleMessageSent"
        @chat-deleted="handleChatDeleted"
      />
    </div>
    <div v-else class="chat-placeholder"> <p>Seleziona una chat per visualizzare i messaggi.</p>
    </div>
  </div>
</template>

<style>
/* Layout principale */
.home-container {
  width: 90%;
  max-width: 80rem;
  height: 90vh;
  margin: 40px auto;
  background-color: whitesmoke;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: row;
  border: 4px solid navy;
  overflow: hidden;
} 

/* Sidebar */
.sidebar-wrapper {
  width: 30%;
  height: 100vh;
  background-color: #c7c6e4;
  color: white;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  border-right: 0.15rem solid navy;
}

.search-chat {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 1rem;
  margin-bottom: 1.5rem;
}

.search-chat input[type="text"] {
  width: 80%;
  padding: 0.45rem 1rem;
  border: 2px solid rgb(185, 181, 219);
  border-radius: 2rem;
  font-size: 1.1em;
  outline: none;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
  background-color: #f8f8f8;
}

.search-chat input[type="text"]:focus {
  border-color: navy;
}

.search-chat input[type="text"]::placeholder {
  color: #95a5a6;
}

.chat-list h3 {
  text-align: center;
  color: #444;
}

.chat-list{
  overflow-y: auto;
  max-height: calc(100vh - 200px);
  flex-grow: 1;
}

.chat-item {
  padding: 0.2rem;
  border-radius: 1rem;
  cursor: pointer;
}

/* Chat Area */
.chat-wrapper {
  flex-grow: 1;
  background-color: navy;
  display: flex;
  flex-direction: column;
}
</style>
