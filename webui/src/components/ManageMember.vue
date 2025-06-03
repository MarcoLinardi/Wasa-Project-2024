<script>
import axiosInstance from "@/services/axios";
import UserItem from "./UserItem.vue";

export default {
  data() {
    return {
      users: [],
      userName: "",
      selectedToAdd: [],
      selectedToRemove: [],
    };
  },
  props: {
    chat: {
      type: Object,
      required: true
    },
    groupMembers: {
      type: Array,
      required: true,
      default: () => []
    }
  },
  component: {
    UserItem
  },
  created() {
    this.loadUsers()
  },
  computed: {
    filteredUsers() {
      if (!this.userName) {
        return this.users;
      }
      const userName = this.userName.toLowerCase();
      return this.users.filter(user => {
        return user.name && user.name.toLowerCase().includes(userName);
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
    isAlreadyMember(user) {
      return this.groupMembers.some(member => member.userId === user.userId);
    },
    isSelected(user) {
      return this.selectedToAdd.includes(user) || this.selectedToRemove.includes(user);
    },

    toggleUser(user) {
      if (this.isAlreadyMember(user)) {
        this.toggleSelection(user, this.selectedToRemove);
      } else {
        this.toggleSelection(user, this.selectedToAdd);
      }
    },

    toggleSelection(user, list) {
      const index = list.findIndex(u => u.userId === user.userId);
      if (index === -1) list.push(user);
      else list.splice(index, 1);
    },

    async addSelected() {
      // TODO: chiamata API per aggiungere selectedToAdd
      const chatId = this.chat.chatId;
      const response = await axiosInstance.post(`/chats/${chatId}/members`)
    },

    async removeSelected() {
      // TODO: chiamata API per rimuovere selectedToRemove
      console.log("Rimuovi utenti:", this.selectedToRemove);
      this.selectedToRemove = [];
    }
  }
}
</script>

<template>
  <div class="modal-backdrop">
    <div class="modal-content">

      <!-- Header -->
      <div class="modal-header">
        <h2>Gestisci Membri</h2>
        <button @click="$emit('close')" class="close-button">&times;</button>
      </div>

      <!-- Ricerca -->
      <div class="search-container">
        <input
          v-model="userName"
          type="text"
          placeholder="Cerca utente..."
          class="search-input"
        />
      </div>

      <!-- Lista utenti -->
      <div class="user-list-container">
        <p v-if="!filteredUsers || filteredUsers.length === 0" class="no-users-message">
          Nessun utente disponibile o nessun risultato per la ricerca...
        </p>
        <ul v-else class="user-list">
          <li
            v-for="user in filteredUsers"
            :key="user.userId"
          >
            <UserItem
              :user="user"
              :is-member="isAlreadyMember(user)"
              :is-selected="isSelected(user)"
              @select-user="toggleUser(user)"
            />
          </li>
        </ul>
      </div>

      <!-- Footer -->
      <div class="modal-footer">
        <button class="manage-button remove" :disabled="selectedToRemove.length === 0" @click="removeSelected">
          Rimuovi
        </button>
        <button class="manage-button add" :disabled="selectedToAdd.length === 0" @click="addSelected">
          Aggiungi
        </button>
      </div>

    </div>
  </div>
</template>


<style scoped>
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
}

.modal-content {
  background: #c7c6e4;
  padding: 2rem;
  border-radius: 1.3rem;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
  border: 2px solid navy;
  width: 33rem;
  max-height: 80vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-left: 2rem;
  padding-right: 1.3rem;
}

.modal-header h2 {
  font-size: 1.5rem;
  margin: 0;
  color: navy;
}

.close-button {
  background: none;
  border: none;
  font-size: 1.8rem;
  color: navy;
  cursor: pointer;
}

.user-list {
  list-style: none;
  padding: 1rem;
  margin: 0;
  border-radius: 0.6rem;
}

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

.modal-footer {
  display: flex;
  justify-content: flex-end;
  margin-top: 1rem;
}

.search-container {
  padding-bottom: 1rem;
  padding-left: 1.5rem;
  padding-right: 1.5rem;
}

.search-input {
  width: 100%;
  padding: 0.6rem 1rem;
  border: 2px solid navy;
  border-radius: 0.8rem;
  font-size: 1rem;
}

.user-list-container {
  max-height: 50vh;
  overflow-y: auto;
  margin-bottom: 1rem;
  padding-right: 0.5rem;
}

.modal-footer {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
}

.manage-button {
  flex: 1;
  padding: 0.6rem;
  border-radius: 0.8rem;
  border: none;
  font-weight: bold;
  cursor: pointer;
  font-size: 1rem;
  box-shadow: 2px 2px 6px rgba(0, 0, 0, 0.2);
  transition: background-color 0.3s;
}

.manage-button.add {
  background-color: navy;
  color: white;
}

.manage-button.remove {
  background-color: crimson;
  color: white;
}

.manage-button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

</style>