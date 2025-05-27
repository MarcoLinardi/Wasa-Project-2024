<script>
import axiosInstance from "@/services/axios";
import UserItem from "../components/UserItem.vue";
import NewChat from "../components/NewChat.vue";

export default {
  components: {
    UserItem,
    NewChat
  },
  data() {
    return {
      userName: "",
      users: [],
      selectedUser: null
    }
  },
  computed: {
    filteredUsers() {
      if (!this.userName) {
        return this.users;
      }
      const searchTerm = this.userName.toLowerCase();
      return this.users.filter(user => {
        return user.name && user.name.toLowerCase().includes(searchTerm);
      });
    }
  },
  methods: {
    async loadUsers() {
      try {
        const response = await axiosInstance.get("/users");
        this.users = response.data.users || response.data;
        this.users.sort((a, b) => (a.name || "").localeCompare(b.name || ""));
        console.log("Utenti caricati:", this.users);
      } catch (e) {
        console.error("[loadUsers] Errore nel caricamento degli utenti:", e);
        this.users = [];
      }
    },
    handleUserSelected(user) {
      this.selectedUser = user;
      console.log("Utente selezionato in UsersView:", this.selectedUser);
    },
    clearSelectedUser() {
      this.selectedUser = null;
    },
    async handleCreateNewChat(messageContent) {
      if (!this.selectedUser || !messageContent.trim()) {
        alert("Seleziona un utente e scrivi un messaggio.");
        return;
      }
      console.log(`Tentativo di creare una nuova chat con: ${this.selectedUser.name}`);
      console.log(`Messaggio: ${messageContent}`);
      
      // Qui implementerai la logica per creare la chat, ad esempio con una chiamata API:
      try {
        // Esempio di chiamata API (da adattare al tuo backend)
        // const response = await axiosInstance.post('/chats/create', {
        //   recipientUserId: this.selectedUser.userId,
        //   message: messageContent
        // });
        // console.log('Nuova chat creata:', response.data);
        
        alert(`Chat con ${this.selectedUser.name} iniziata (simulazione).\nMessaggio: "${messageContent}"\nImplementare la logica di backend qui!`);
        
        // Dopo aver creato la chat, potresti voler resettare lo stato
        this.selectedUser = null; 
      } catch (error) {
        console.error("Errore durante la creazione della chat:", error);
        alert("Errore durante la creazione della chat.");
      }
    }
  },
  created() {
    this.loadUsers();
  }
}
</script>

<template>
  <div class="users-list-view">
    <div class="users-list-header">
      <h2>Lista Utenti</h2>
      <div class="search-user">
        <label for="userName">Cerca utente</label>
        <input type="text" id="userName" v-model="userName" placeholder="Cerca utente">
      </div>
    </div>

    <div v-if="selectedUser" class="new-chat-view">
      <NewChat
        :user="selectedUser"
        @send-new-message="handleCreateNewChat"
        @close-chat-view="clearSelectedUser" 
      />
    </div>
    <div v-else class="users-list-container">
      <p v-if="!filteredUsers || filteredUsers.length === 0" class="no-users-message">
        Nessun utente disponibile o nessun risultato per la ricerca...
      </p>
      <ul class="user-list" v-else>
        <li v-for="user in filteredUsers" :key="user.userId" class="user-item-wrapper">
          <UserItem 
          :user="user"
          @select-user="handleUserSelected"
          />
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.users-list-view {
  width: 100%;
  max-width: 700px;
  height: 90%;
  margin: 2rem auto;
  padding: 30px;
  background-color: whitesmoke;
  border-radius: 10px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  border: 4px solid navy;
  overflow: hidden;
}

/* HEADER */
.users-list-header {
  text-align: center;
  margin-bottom: 25px;
}

.users-list-header h2 {
  color: #2c3e50;
  font-size: 2.5em;
  margin-bottom: 20px;
  position: relative;
  display: inline-block; /* Per allineare il pseudo-elemento */
}

.users-list-header h2::after {
  content: '';
  position: absolute;
  left: 50%;
  bottom: -10px;
  transform: translateX(-50%);
  width: 60px;
  height: 4px;
  background-color: navy; /* Colore d'accento di Vue.js */
  border-radius: 2px;
}

/* SEARCH SECTION */
.search-user {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  margin-top: 10px;
}

.search-user label {
  font-size: 1.1em;
  color: #34495e;
  margin-bottom: 5px;
  font-weight: 600;
}

.search-user input[type="text"] {
  width: 100%;
  max-width: 400px; /* Limita la larghezza dell'input */
  padding: 12px 18px;
  border: 2px solid navy;
  border-radius: 25px;
  font-size: 1.1em;
  outline: none;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
  background-color: #f8f8f8;
}

.search-user input[type="text"]:focus {
  border-color: navy;
  box-shadow: 0 0 0 4px rgba(66, 185, 131, 0.2);
}

.search-user input[type="text"]::placeholder {
  color: #95a5a6;
}

/* NO USERS MESSAGE */
.no-users-message {
  text-align: center;
  padding: 30px;
  background-color: #f0f4f7;
  border-radius: 8px;
  color: #7f8c8d;
  font-size: 1.2em;
  margin-top: 0px;
  box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.05);
}

.users-list-container {
  display: flex;
  flex-grow: 1;
  flex-direction: column;
  overflow-y: hidden;
  background-color: #f8f8f8;
  border: 2px solid navy;
  border-radius: 12px;
}

/* USER LIST */
.user-list {
  flex-grow: 1;
  overflow-y: auto;
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  gap: 2px;
  padding-top: 0px;
  flex-direction: column;
}

.user-item-wrapper {
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s ease-in-out;
}

.user-item-wrapper:last-child {
  border-bottom: none;
}

.user-item-wrapper:hover,
.user-item-wrapper:focus-within {
  background-color: #f0f0f0;
}

.user-item-wrapper.selected {
  background-color: navy;
  color: white;
}
.user-item-wrapper.selected .user-name {
  color: white;
}

.new-chat-view {
  display: flex;
  flex-grow: 1;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>
