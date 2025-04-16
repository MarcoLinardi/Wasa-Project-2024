<script>

export default {

  data: function() {
    return{
      mode: "chat",
      userList: [],
      chatList: [],
    }
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
        this.chatList = response.data.chats;
      } catch (e) {
        console.error(e);
      }
    },
    async refreshUsers() {
      try {
        const response = await this.$axios.get("/users");
        this.userList = response.data.users;
      } catch (e) {
        console.error(e);
      }
    },
    async selectChat(chat) {
      this.selectedChat = chat;
      await this.loadMessages(chat.id);
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
    this.refreshUsers();
  }
}
</script>

<template>
    <div class="sidebar">
      <!-- TOP: bottoni Chat / Nuova Chat -->
      <div class="sidebar-top">
        <button class="sidebar-button" @click="this.switchMode('chats')">Chat</button>
        <button class="sidebar-button" @click="this.switchMode('users')">Nuova Chat</button>
      </div>
  
      <!-- MIDDLE: ChatList o UserList -->
      <div class="sidebar-middle">
        <ChatList 
          v-if="mode === 'chats'" 
          :chats="chatList" 
          @selectChat="$emit('selectChat', $event)"
        />
        <UserList 
          v-else-if="mode === 'users'" 
          :users="userList" 
          @startChat="$emit('startChat', $event)"
          @createGroup="$emit('createGroup')"
        />
      </div>
  
      <!-- BOTTOM: Logout -->
      <div class="sidebar-bottom">
        <button class="logout-button" @click="this.logout">Logout</button>
      </div>
    </div>
  </template>  
  
  <style scoped>
  .sidebar {
    display: flex;
    flex-direction: column;
    height: 100vh;
    width: 30%;
    background-color: rgb(221, 172, 116);
    border-right: 1px solid #ddd;
    padding: 1rem;
    box-sizing: border-box;
  }
  
  .sidebar-top {
    display: flex;
    gap: 0.5rem;
    justify-content: center;
    flex-wrap: wrap;
  }
  
  .sidebar-middle {
    flex-grow: 1;
    overflow-y: auto;
    margin-top: 1rem;
  }
  
  .sidebar-bottom {
    margin-top: auto;
    display: flex;
    justify-content: center;
  }
  
  .sidebar-button {
    background-color: #d9d9d9;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 10px;
    cursor: pointer;
    font-weight: bold;
    text-align: center;
    min-width: 100px;
  }
  
  .sidebar-button:hover {
    background-color: #cccccc;
  }
  
  .logout-button {
    background-color: rgb(217, 128, 91);
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 10px;
    cursor: pointer;
    font-weight: bold;
  }

  .logout-button:hover {
    background-color: rgb(180, 90, 58);
  }
  </style>
  