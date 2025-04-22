<script>
export default {
  props: {
    users: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      creatingGroup: false,
      group: {
        name: "",
        participants: []
      }
    };
  },
  methods: {
    startGroupCreation() {
      this.creatingGroup = true;
      this.group = {
        name: "",
        participants: []
      };
    },
    handleUserClick(user) {
    if (this.creatingGroup) {
      const index = this.group.participants.indexOf(user.userId);
      if (index === -1) {
        this.group.participants.push(user.userId);
      } else {
        this.group.participants.splice(index, 1);
      }
    } else {
      this.$emit("startChat", user);
    }
  },
    confirmGroup() {
      if (!this.group.name || this.group.participants.length < 1) {
        alert("Inserisci un nome e seleziona almeno un partecipante");
        return;
      }
      this.$emit("createGroup", this.group);
      this.group = {
        name: "",
        participants: []
      };
      this.creatingGroup = false;
    }
  }
};
</script>
<template>
  <div class="list">
    <!-- Bottone per attivare la modalità creazione gruppo -->
    <li class="list-item special" @click="startGroupCreation">
      Crea Gruppo
    </li>

    <!-- Campo per inserire il nome del gruppo -->
    <li v-if="creatingGroup" class="list-item">
      <input
        v-model="group.name"
        type="text"
        placeholder="Inserisci un nome per il gruppo"
        class="group-name-input"
      />
    </li>

    <!-- Lista utenti -->
    <li
      v-for="user in users"
      :key="user.userId"
      class="list-item"
      :class="{ selectable: creatingGroup, selected: group.participants.includes(user.userId) }"
      @click="handleUserClick(user)"
    >
      {{ user.name }}
    </li>

    <!-- Bottone per confermare il gruppo -->
    <li
      v-if="creatingGroup"
      class="list-item confirm"
      @click="confirmGroup"
    >
      ✅ Conferma Gruppo
    </li>
  </div>
</template>
  
  <style scoped>
  .list {
    list-style: none;
    padding: 0;
    margin: 0;
  }
  
  .list-item {
    padding: 0.5rem 1rem;
    cursor: pointer;
    border-radius: 12px;
  }
  
  .list-item:hover {
    background-color: rgb(224, 160, 100);
    transition: background-color 0.3s ease;
  }
  
  .special {
    font-weight: bold;
    color: black;
    border-radius: 12px;
  }

  .list-item.selectable {
  cursor: pointer;
  }
  .list-item.selected {
    background-color: rgb(217, 128, 91);
    border-radius: 12px;
  }
  .list-item.confirm {
    font-weight: bold;
    color: green;
    cursor: pointer;
    border-radius: 12px;
  }
  .group-name-input {
  width: 100%;
  padding: 8px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 12px;
  background-color: #f9f9f9;
  color: #333;
}

.group-name-input::placeholder {
  color: #aaa;
  font-style: italic;
}
  </style>
  