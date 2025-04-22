<script>

export default {

  data: function() {
    return{
      mode: "chats",
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
        this.userList.sort((a, b) => {

          if (a.name < b.name) {
            return -1;
          }
          if (a.name > b.name) {
            return 1;
          }
          return 0;
        })
      } catch (e) {
        console.error(e);
      }
    },
  },
  mounted() {
    this.refreshChats();
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
          v-if="mode === 'users'" 
          :users="userList" 
          @startChat="$emit('startChat', $event)"
          @createGroup="$emit('createGroup', $event)"
        />
      </div>
  
      <!-- BOTTOM: Logout -->
      <div class="sidebar-bottom">
        <button class="sidebar-button" @click="$emit('logout')">Logout</button>
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
    background-color: rgb(217, 128, 91);
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 10px;
    cursor: pointer;
    font-weight: bold;
    text-align: center;
    min-width: 100px;
  }
  
  .sidebar-button:hover {
    background-color: rgb(180, 90, 58);
  }
  
  .logout-button {
    background-color: rgb(217, 128, 91);
    color: black;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 10px;
    cursor: pointer;
    font-weight: bold;
  }
  </style>
  