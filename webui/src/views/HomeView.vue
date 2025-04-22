<script>
export default {
  
  data() {
    return {
      mode: 'chats',      // chats oppure users
      chatList: [],
      userList: [],
      selectedChat: null,
      messages: [],
      chatId: -1,
      token: -1,
    };
  },
  mounted() {
    this.token = localStorage.getItem('token')
  },
  methods: {
    async selectChat(chat) {
      this.selectedChat = chat;
      await this.loadMessages(chat.chatId);
    },
    async loadMessages(chatId) {
      try {
        const response = await this.$axios.get(`/chats/${chatId}/messages`);
        this.messages = response.data;
      } catch (e) {
        console.error(e);
      }
    },
    async switchMode(newMode) {
      this.mode = newMode;
      if (newMode === 'chats') {
        await this.refreshChats();
      } else if (newMode === 'users') {
        await this.refreshUsers();
      }
    },
    
    async startChat(user) {
      try {
        const response = await this.$axios.post(`/chats`, {name: user.name, users: [user.userId], isGroup: false})
        this.chatId = response.data.chatId; 
        this.switchMode("chats");
        this.selectedChat = {
          chatId: response.data.chatId,
          name: user.name,
          isGroup: false
        };
        this.chatList.push(this.selectedChat);
        console.log("Chat privata creata con:", user.name);
      } catch (e) {
        console.error('Errore nella creazione chat: ',e);
      }
    },
    async createGroup(group) {
      console.log("Create new group!");
      try {
        const response = await this.$axios.post(`/chats`, {name: group.name, users: group.participants, isGroup: true}) 
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
  
