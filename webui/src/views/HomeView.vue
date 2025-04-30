<script>
import { sendMessage, deleteChat, createPrivateChat } from "@/services/api"; 
import axiosInstance from "@/services/axios";

export default {
  
  data() {
    return {
      mode: 'chats',      // chats oppure users
      chatList: [],
      userList: [],
      selectedChat: null,
      selectedUser: null,
      chatId: -1,
      token: -1,
    };
  },
  mounted() {
    this.token = localStorage.getItem('token')
  },
  methods: {

    async selectChat(input) {
      let chatFound = null 
      // controllo se c'è una chat con lo stesso nome con l'utente selezionato
      this.chatList.forEach(chat => {
        if (!chat.isGroup) {
          if (input.name == chat.name) {
            chatFound = chat; 
          }
        }
      });
      if (chatFound) {
        this.selectedChat = chatFound
        await this.loadMessages(chatFound.chatId)
      } else {
        // Se non esiste una chat con l'utente la crea
        await this.startChat(input)
      }
      
    },  

    startChat(user) {
      this.selectedChat = null;     // non c'è una chat creata
      this.selectedUser = user;     // mostriamo solo il nome utente selezionato
      this.messages = [];
      console.log("selectedUser =", this.selectedUser);
    },

    async handleSendMessage(chatId, content) {
      await this.sendMessage(chatId, content)
      await this.loadMessages(chatId)
    },

    async switchMode(newMode) {
      this.mode = newMode;
      if (newMode === 'chats') {
        await this.refreshChats();
      } else if (newMode === 'users') {
        await this.refreshUsers();
      }
    },
    async refreshChats() {
      try {
        const response = await axiosInstance.get("/chats");
        this.chatList = response.data.chats;
      } catch (e) {
        console.error(e);
      }
    },
    async refreshUsers() {
      try {
        const response = await axiosInstance.get("/users");
        this.userList = response.data.users;
        this.userList.sort((a, b) => {

          if (a.name < b.name) {
            return -1; // a viene prima di b
          }
          if (a.name > b.name) {
            return 1; // b viene prima di a
          }
          return 0; // a e b sono uguali
        })
      } catch (e) {
        console.error(e);
      }
    },

    async createGroup(group) {
      console.log("Create new group!");
      try {
        const response = await axiosInstance.post(`/chats`, {
          name: group.name, 
          users: group.participants, 
          isGroup: true}) 
        const newChat = {
          chatId: response.data.chatId,
          name: group.name,
          isGroup: true
      };
      this.chatList.push(newChat);
      console.log("Gruppo creato con successo");
      } catch (e) {
        console.error('Errore nella creazione gruppo: ',e);
      }
    },
    logout() {
      localStorage.removeItem('token');
      this.$router.push('/login');
      console.log('Logout effettuato');
    }
  }

}
</script>

<template>
  <div class="home-container">
    <div class="sidebar-wrapper">
      <Sidebar 
        :mode="mode"
        :chatList="chatList"
        :userList="userList"
        @selectChat="selectChat"
        @startChat="startChat"
        @createGroup="createGroup"
        @logout="logout"
        @switchMode="switchMode"
      />
    </div>

    <div class="chat-wrapper">
      <ChatArea 
        :selectedChat="selectedChat"
        :selectedUser="selectedUser"
        @sendFirstMessage="handleFirstMessage"
        @sendMessage="handleSendMessage"
      />
    </div>

  </div>
</template>

  
<style scoped>
 .home-container {
    display: flex;
    flex-direction: row;
    height: 100vh;
    width: 100%;
    overflow: hidden;
}

.sidebar-wrapper {
  width: 300px;
  flex-shrink: 0;
  height: 100%;
  background-color: #d59d63;
  z-index: 2;
}

.chat-wrapper {
  flex-grow: 1;
  height: 100%;
  overflow: hidden;
  background-color: rgb(210, 180, 140);
}

  </style>
  
