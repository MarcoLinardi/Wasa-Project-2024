<script>
import axiosInstance from "@/services/axios";

export default {
  props: {
    user: {
      type: Object,
      required: true
    },
    isUserLogged: {
      type: Boolean,
      required: true
    }
  },

  data() {
    return {
      isEditingName: false,
      editedName: ""
    };
  },

  methods: {
    async saveName() {
      try {
        const response = await axiosInstance.put('/user/name', {
          newName: this.editedName
        });

        if (response.status === 200) {
          this.user.name = this.editedName;
          this.isEditingName = false;
          this.$emit('update-name', this.editedName);
          console.log('Nome aggiornato con successo:', response.data);
        }
      } catch (error) {
        console.error('Errore durante l\'aggiornamento del nome:', error);
        alert('Si è verificato un errore durante l\'aggiornamento del nome');
      }
    },

    cancelEdit() {
      this.isEditingName = false;
      this.editedName = this.user.name; // ripristina il nome precedente
    },
    editPhoto() {
      alert("Modifica foto non ancora implementata");
    }
  }
}
</script>

<template>
  <div class="user-profile">
    <button class="close-btn" title="Chiudi" @click="$emit('close')">
      <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
    </button>

    <div v-if="!isUserLogged" class="profile-content">
      <img :src="user.photo" class="user-photo" />
      <div class="profile-name">{{ user.name }}</div>
      <button class="new-chat" title="Inizia una conversazione" @click="$emit('start-new-chat', user)">
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit"/></svg>
      </button>
    </div>

    <div v-else class="profile-content">
      <div class="photo-section">
        <img :src="user.photo" class="user-photo" />
        <button class="edit-photo-btn" title="Cambia foto" @click="editPhoto">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit-2"/></svg>
        </button>
      </div>

      <div class="name-section">
        <template v-if="isEditingName">
          <input v-model="editedName" class="name-input" />
          <button @click="saveName" class="save-btn">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#check"/></svg>
          </button>
          <button @click="cancelEdit" class="cancel-btn">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
          </button>
        </template>
        <template v-else>
          <div class="profile-name">{{ user.name }}</div>
          <button class="edit-name-btn" title="Cambia nome" @click="isEditingName = true">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit-3"/></svg>
          </button>
        </template>
      </div>
    </div>
  </div>
</template>


<style>
.user-profile {
  position: relative;
  background-color: #f4f4f4;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  padding: 24px;
  width: 28rem;
  max-width: 90%;
  margin: auto;
}

.profile-content {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 16px;
}

.photo-section {
  position: relative;
}

.user-photo {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 50%;
  border: 2px solid #ccc;
}

.edit-photo-btn {
  position: absolute;
  bottom: 4px;
  right: 4px;
  background: #dad9d9;
  border: none;
  border-radius: 50%;
  cursor: pointer;
}

.name-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.profile-name {
  font-size: 1.2rem;
  font-weight: 600;
  color: #333;
}

.name-input {
  font-size: 1rem;
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

.new-chat,
.edit-name-btn,
.save-btn,
.cancel-btn {
  border-radius: 50%;
  border: none;
  cursor: pointer;
  font-size: 1rem;
}

.close-btn {
  position: absolute;
  top: 12px;
  right: 12px;
  background-color: transparent;
  color: #444;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-btn svg {
  transition: transform 0.2s ease;
}

.close-btn:hover svg {
  transform: scale(1.25);
}
</style>
