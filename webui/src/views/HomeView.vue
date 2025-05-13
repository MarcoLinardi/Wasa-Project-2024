<script>
import axiosInstance from "@/services/axios";
import UserItem from "../components/UserItem.vue";
import UserProfile from "../components/UserProfile.vue";
import ChatArea from "../components/ChatArea.vue";

export default {
  data() {
    return {
      users: [],
      chats: [],
      selectedSection: "chat",
      loggedUser: null,
      selectedUser: null,
      showUserProfile: false,
      showChatArea: false,
      isUserLogged: false,
      selectedChat: null,
    }
  },
  components: {
    UserItem,
    UserProfile,
    ChatArea
  },
  mounted() {
    this.loadUsers();
    this.loadLoggedUser();
  },

  methods: {
    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('user')
      this.$router.push('/login');
      console.log('Logout effettuato');
    },
    loadLoggedUser() {
      const userData = localStorage.getItem('user');
      if (userData) {
        this.loggedUser = JSON.parse(userData);
        console.log("Utente loggato caricato:", this.loggedUser);
      } else {
        console.warn("Nessun utente trovato in localStorage");
      }
    },
    async loadUsers() {
      try {
        const response = await axiosInstance.get("/users");
        this.users = response.data.users;
        this.users.sort((a, b) => {
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
    openUserProfile() {
      this.selectedUser = this.loggedUser;
      this.showUserProfile = true;
      this.isUserLogged = true;
      this.showChatArea = false;
    },
    closeUserProfile() {
      this.showUserProfile = false;
      this.selectedUser = null;
      this.isUserLogged = false
    },
    openChat(user) {
      this.selectedUser = user;
      this.selectedChat = null;
      this.showUserProfile = false;
      this.showChatArea = true;
    },
    handleUserClick(user) {
      this.selectedUser = user;
      this.showUserProfile = true;
      this.showChatArea = false;
      this.isUserLogged = false;
    },
    updateUserName(newName) {
    console.log("Nuovo nome ricevuto:", newName);
    this.loggedUser.name = newName;
  },
  }
}
</script>


<template>
  <div class="home-container">
    <div class="sidebar-wrapper">
      <div class="buttons">
        <button 
          class="btn" 
          :class="{ 'active': selectedSection === 'chat' }"
          @click="selectedSection = 'chat'">
          Chat
        </button>
        <button 
          class="btn" 
          :class="{ 'active': selectedSection === 'user' }"
          @click="selectedSection = 'user'">
          Utenti
        </button>
      </div>

      <!-- Lista Chat -->
      <div v-if="selectedSection === 'chat'" class="chat-list">
        <div class="chat-item">
          <div class="font-semibold">Chat 1</div>
          <div class="text">Ultimo messaggio...</div>
        </div>
        <div class="chat-item">
          <div class="font-semibold">Chat 2</div>
          <div class="text">Ultimo messaggio...</div>
        </div>
      </div>

      <!-- Lista Utenti -->
      <div v-if="selectedSection === 'user'" class="user-list">
        <!-- Bottone Crea Gruppo -->
        <button class="btn-create-group">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#users"/></svg> Crea Gruppo
        </button>

        <!-- Lista degli utenti come componenti UserItem -->
        <UserItem 
          v-for="user in users" 
          :key="user.id" 
          :user="user" 
          @select-user="handleUserClick"
        />
      </div>

      <!-- Bottoni di Controllo -->
      <div class="control-options">
        <button class="control-btn" title="Profilo"  @click="openUserProfile">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/> </svg></button>
        <button class="control-btn" title="Logout" @click="logout">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/> </svg></button>
      </div>
    </div>

    <!-- Chat Area -->
    <div class="chat-wrapper">
      <UserProfile 
        v-if="showUserProfile" 
        :user="selectedUser"
        :isUserLogged="isUserLogged"
        @close="closeUserProfile"
        @start-new-chat="openChat"
      />
      <ChatArea
        v-if="showChatArea"
        :selectedUser="selectedUser"
        :selectedChat="selectedChat"
      />
    </div>
  </div>
</template>

<style>
/* Layout principale */
.home-container {
  display: flex;
  height: 100vh;
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
  border-right: 1px solid #444;
}

.buttons {
  display: flex;
  justify-content: space-around;
  margin-bottom: 20px;
}

.btn {
  background-color: rgb(217, 128, 91);
  color: white;
  padding: 10px 20px;
  border-radius: 15px;
  cursor: pointer;
  border: none;
  transition: background-color 0.3s ease;
}
.btn:hover {
  background-color: rgb(202, 115, 79);
}
.btn.active {
  background-color: rgb(180, 90, 58);
}

.chat-list, .user-list {
  overflow-y: auto;
  max-height: calc(100vh - 200px); /* Spazio per bottoni e controlli */
  flex-grow: 1;
}

.chat-item {
  padding: 12px;
  border-radius: 5px;
  cursor: pointer;
}
.chat-item:hover {
  background-color: rgb(217, 128, 91);
}

.control-options {
  margin-top: auto;
  display: flex;
  justify-content: center;
  width: 100%;
  justify-content: space-around;
}
.control-btn {
  display: flex;
  justify-content: center;
  padding: 12px;
  background-color: rgb(217, 128, 91);
  color: white;
  border-radius: 15px;
  cursor: pointer;
  margin-top: 10px;
  border: none;
  width: 3.5rem;
}
.control-btn:hover {
  background-color: rgb(202, 115, 79);
}

/* Bottone crea gruppo */
.btn-create-group {
  background-color: transparent;
  width: 100%;
  color: white;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  border: none;
  transition: background-color 0.3s ease;
}
.btn-create-group:hover {
  background-color: rgb(202, 115, 79);
}

/* Chat Area */
.chat-wrapper {
  flex-grow: 1;
  background-color: rgb(210, 180, 140);
  display: flex;
  flex-direction: column;
}
</style>
