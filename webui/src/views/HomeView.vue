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
  created() {
    this.loadChats();
    this.loadLoggedUser();
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
        this.chats = response.data.chats;
        this.chats.sort((a, b) => (a.name || "").localeCompare(b.name || ""));
        console.log('chat caricate:', JSON.stringify(this.chats, null, 2));

      } catch (e) {
        console.error(e);
      }
    },
    handleChatSelected(chat) {
      this.selectedChat = chat;
      console.log("Chat selezionata in HomeView:", this.selectedChat);
    }
  }
}
</script>


<template>
  <div class="home-container">
    <div class="sidebar-wrapper">
      <!-- Lista Chat -->
      <div class="chat-list">
        <h2>Le mie chat</h2>
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
        @select-chat="handleChatSelected"
      />
    </div>
    <div v-else class="chat-placeholder"> <p>Seleziona una chat per visualizzare i messaggi.</p>
    </div>
  </div>
</template>

<style>
/* Layout principale */
.home-container {
  width: 100%;
  max-width: 1000px;
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
  width: 25%;
  height: 100%;
  background-color: rgb(210, 180, 140);
  color: white;
  padding: 20px;
  display: flex;
  flex-direction: column;
  border-right: 2px solid navy;
}

.search-chat {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  margin-top: 10px;
  margin-bottom: 5px;
}

.search-chat input[type="text"] {
  width: 90%;
  padding: 6px 18px;
  border: 2px solid rgb(185, 181, 219);
  border-radius: 25px;
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

.chat-list h2 {
  text-align: center;
}

.chat-list{
  overflow-y: auto;
  max-height: calc(100vh - 200px);
  flex-grow: 1;
}

.chat-item {
  padding: 1px;
  border-radius: 5px;
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
