<script>
import axiosInstance from "@/services/axios";
import UserItem from "../components/UserItem.vue";

export default {
  components: {
    UserItem
  },
  data() {
    return {
      groupName: '',
      availableUsers: [],
      selectedUserIds: [],
    };
  },
  computed: {
    selectedUsers() {
      return this.availableUsers.filter(user => this.selectedUserIds.includes(user.userId));
    },
    selectedMemberNames() {
      return this.selectedUsers.map(user => user.name).join(', ');
    }
  },
  methods: {
    toggleUserSelection(userId) {
      const index = this.selectedUserIds.indexOf(userId);
      if (index > -1) {
        this.selectedUserIds.splice(index, 1);
      } else {
        if (!this.selectedUserIds.includes(userId)) {
            this.selectedUserIds.push(userId);
        }
      }
    },
    isUserSelected(userId) {
      return this.selectedUserIds.includes(userId);
    },
    async createGroup() {
      if (!this.groupName.trim()) {
        alert("Per favore, inserisci un nome per il gruppo.");
        return;
      }
      if (this.selectedUserIds.length === 0) {
        alert("Per favore, seleziona almeno un membro per il gruppo.");
        return;
      }
      const newGroupRequest = {
        name: this.groupName,
        users: this.selectedUserIds,
        isGroup: true
      };
      try {
        const response = await axiosInstance.post('/chats', newGroupRequest);
        const newChat = {
          chatId: response.data.chatId,
          name: newGroupRequest.name,
          isGroup: true
        };
        console.log('Gruppo creato con successo', newChat);
        this.groupName = '';
        this.selectedUserIds = [];
        this.$router.push("/home");
      } catch (e) {
        console.error('Errore durante la creazione del gruppo:', e.response || e.message || e);
        alert('Errore durante la creazione del gruppo.');
      }
    },
    async loadUsers() {
      try {
        const response = await axiosInstance.get("/users");
        this.availableUsers = response.data.users || response.data;
        this.availableUsers.sort((a, b) => (a.name || "").localeCompare(b.name || ""));
      } catch (e) {
        console.error("[loadUsers] Errore nel caricamento degli utenti:", e);
        this.availableUsers = [];
      }
    },
  },
  created() {
    this.loadUsers();
  }
}
</script>

<template>
  <div class="create-group-view">
    <div class="create-group-header">
      <h2>Crea Nuovo Gruppo</h2>
      <div class="form-group">
        <label for="groupName">Nome del Gruppo:</label>
        <input type="text" id="groupName" v-model="groupName" placeholder="Es: Progetto WasaText">
      </div>
    </div>

    <div class="user-selection-container">
      <h3>Seleziona Membri</h3>
      <p v-if="!availableUsers || availableUsers.length === 0" class="no-users-message">
        Nessun utente disponibile o caricamento in corso...
      </p>
      <ul class="user-list" v-else>
        <li
          v-for="user in availableUsers"
          :key="user.userId"
          class="user-item-wrapper"
          :class="{ 'selected': isUserSelected(user.userId) }"
          @click="toggleUserSelection(user.userId)"
          tabindex="0"
          @keydown.enter="toggleUserSelection(user.userId)"
          @keydown.space.prevent="toggleUserSelection(user.userId)"
        >
          <UserItem :user="user" :is-selected="isUserSelected(user.userId)" />
        </li>
      </ul>
    </div>

    <div class="create-group-footer">
      <div class="selected-members-preview" v-if="selectedUsers.length > 0">
        <h4>Membri Selezionati:</h4>
        <p class="selected-names-display">{{ selectedMemberNames }}</p>
      </div>
      <button @click="createGroup" class="create-button" :disabled="!groupName.trim() || selectedUserIds.length === 0">
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#check"/></svg>
        <span>Crea Gruppo</span>
      </button>
    </div>
  </div>
</template>

<style scoped>

.create-group-view {
  display: flex;
  flex-direction: column;
  height: 90%;
  width: 100%;
  max-width: 700px;
  margin: 2rem auto;
  background-color: whitesmoke;
  border-radius: 10px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  font-family: 'Arial', sans-serif;
  overflow: hidden;
  border: 4px solid navy;
}

.create-group-header {
  padding: 15px 25px 1px 25px;
  flex-shrink: 0;
}

.create-group-header h2 {
  text-align: center;
  color: #2c3e50;
  margin-top: 0;
  margin-bottom: 0px;
}

.form-group {
  margin-bottom: 0;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: bold;
  color: #555;
}

.form-group input[type="text"] {
  width: 100%;
  padding: 8px;
  border: 1px solid navy;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
  box-sizing: border-box;
}

.form-group input[type="text"]:focus {
  border-color: #007bff;
  outline: none;
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}

/* Contenitore per la selezione utenti che gestirà lo scroll della lista */
.user-selection-container {
  padding: 15px 25px;
  flex-grow: 1;
  overflow-y: hidden;
  display: flex;
  flex-direction: column;
  min-height: 150px;
}

.user-selection-container h3 {
  margin-top: 0;
  margin-bottom: 5px;
  color: #333;
}

.no-users-message {
  color: #777;
  text-align: center;
  padding: 20px;
}

.user-list {
  flex-grow: 1;
  overflow-y: auto;
  list-style-type: none;
  padding: 0;
  margin: 0;
  border: 1px solid navy;
  border-radius: 6px;
  background-color: #fff;
}

/* Footer fisso in basso */
.create-group-footer {
  padding: 1px 25px 20px 25px;
  flex-shrink: 0; /* Non si restringe */
  background-color: #f9f9f9;
  position: sticky; bottom: 0; z-index: 10; /* Attiva se vuoi che il footer si incolli in fondo alla finestra se il .create-group-view stesso scrolla */
}

.selected-members-preview {
  margin-bottom: 15px;
}

.selected-members-preview h4 {
  margin-top: 0;
  margin-bottom: 0.5rem;
  color: #333;
  font-size: 1rem;
}

.selected-names-display {
  font-size: 0.9rem;
  color: #495057;
  line-height: 1.5;
  background-color: #e9ecef;
  padding: 8px 12px;
  border-radius: 4px;
  min-height: 1.5em;
}

.feather {
  width: 20px;
  height: 20px;
}

.create-button {
  display: block;
  width: 100%;
  padding: 12px 15px;
  gap: 2rem;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1.1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.create-button:hover:not(:disabled) {
  background-color: #218838;
}

.create-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
  opacity: 0.7;
}

</style>