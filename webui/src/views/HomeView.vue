<script>
import axiosInstance from "@/services/axios";
import UserItem from "../components/UserItem.vue";
export default {
  data() {
    return {
      users: [],
      chats: [],
      selectedSection: "chat",
    }
  },
  mounted() {
    this.loadUsers();  // Carica gli utenti quando il componente è montato
  },

  methods: {
    logout() {
      localStorage.removeItem('token');
      this.$router.push('/login');
      console.log('Logout effettuato');
    },
    async loadUsers() {
      try {
        const response = await axiosInstance.get("/users");
        this.users = response.data.users;
        console.log("lista utenti: " + this.users)
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
  }
}
</script>


<template>
  <div class="home-container">
    <!-- Sidebar -->
    <div class="sidebar-wrapper">
      <!-- Bottoni per scegliere tra Chat e Utenti -->
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
        <!-- Bottone Crea Gruppo Fisso -->
        <button class="btn-create-group">Crea Gruppo</button>

        <!-- Mappa degli utenti per visualizzarli come componenti UserItem -->
        <UserItem 
          v-for="user in users" 
          :key="user.id" 
          :user="user" 
        />
      </div>

      <!-- Altri controlli -->
      <div class="control-options">
        <button class="control-btn" @click="logout">Logout</button>
      </div>
    </div>

    <!-- Chat Area -->
    <div class="chat-wrapper">
      <div class="chat-header">
        <h2>Seleziona una chat</h2>
      </div>
      <div class="chat-messages">
        <div class="message">
          <strong>Utente:</strong> Messaggio di esempio
        </div>
      </div>
      <div class="chat-input">
        <input type="text" placeholder="Scrivi un messaggio..." />
      </div>
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
  background-color: rgb(202, 115, 79); /* Colore di hover leggermente più scuro */
}

.btn.active {
  background-color: rgb(180, 90, 58); /* Colore attivo (scuro) */
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
}

.control-btn {
  padding: 12px;
  background-color: rgb(217, 128, 91);
  color: white;
  border-radius: 15px;
  cursor: pointer;
  margin-top: 10px;
  border: none;
  width: 100px;
  text-align: center;        
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
  background-color: #f5f5f5;
  display: flex;
  flex-direction: column;
}

.chat-header {
  background-color: rgb(210, 180, 140);
  padding: 15px;
  border-bottom: 1px solid #ccc;
}

.chat-messages {
  flex-grow: 1;
  padding: 20px;
  overflow-y: auto;
}

.message {
  margin-bottom: 10px;
}

.chat-input {
  background-color: rgb(210, 180, 140);
  padding: 15px;
  border-top: 1px solid #ccc;
}

.chat-input input {
  width: 100%;
  padding: 12px;
  border-radius: 10px;
  border: 1px solid rgb(70, 70, 70);
  background-color: transparent;
}
</style>
