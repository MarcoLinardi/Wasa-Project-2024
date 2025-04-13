<script>
import Sidebar from '../components/Sidebar.vue';
import ChatArea from '../components/ChatArea.vue';
export default {
  components: {
    Sidebar,
    ChatArea
  },
  data() {
    return {
      mode: 'chats',      // chats oppure users
      chatList: [],
      userList: [],
      selectedChat: null,
      messages: [],
    };
  },
  methods: {
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
        const response = await this.$axios.get("/chats");
        this.chatList = response.data;
      } catch (e) {
        console.error(e);
      }
    },
    async refreshUsers() {
      try {
        const response = await this.$axios.get("/users");
        this.userList = response.data;
      } catch (e) {
        console.error(e);
      }
    },
    async selectChat(chat) {
      this.selectedChat = chat;
      await this.loadMessages(chat.id);
    },
    async loadMessages(chatId) {
      try {
        const response = await this.$axios.get(`/chats/${chatId}/messages`);
        this.messages = response.data;
      } catch (e) {
        console.error(e);
      }
    },
    async startChat(user) {
      console.log("Start chat with user:", user);
      // Qui puoi creare una nuova chat POST /chats
    },
    async createGroup() {
      console.log("Create new group!");
      // Qui puoi aprire una modale per creare un gruppo
    },
    logout() {
      localStorage.removeItem('token');
      this.$router.push('/login');
      console.log('Logout effettuato');
    }
  },
  mounted() {
    this.refreshChats(); // Quando monti carichi le chat
  }
}
</script>

<template>
	<div class="home-container">
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
  
	  <ChatArea 
		:selectedChat="selectedChat"
		:messages="messages"
	  />
	</div>
  </template>
  
  <style scoped>
  .home-container {
	display: flex;
	height: 100vh;
	width: 100vw;
	overflow: hidden;
  }
  </style>
  
