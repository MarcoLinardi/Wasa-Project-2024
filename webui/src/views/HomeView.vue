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
      chatId: -1,
    };
  },
  methods: {
    async loadMessages(chatId) {
      try {
        const response = await this.$axios.get(`/chats/${chatId}/messages`);
        this.messages = response.data;
      } catch (e) {
        console.error(e);
      }
    },
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
  
