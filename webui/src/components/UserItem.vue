<script>
  export default {
    props: {
        user: {
          type: Object,
          required: true
        },
        isSelected: {
          type: Boolean,
          default: false
        },
        isMember: {
          type: Boolean,
          default: false,
        }
    },
    methods: {
      handleClick() {
        this.$emit('select-user', this.user);
      },
      getUserPhoto(photo) {
        // Se è una path (foto di default)
        if (photo.startsWith('/')) {
          return photo;
        }
        // Altrimenti è una base64 pura
        return 'data:image/jpeg;base64,' + photo;
      }
    }
  }
  
</script>

<template>
  <div
    class="user-item"
    :class="{ 'is-member': isMember, 'selected': isSelected }"
    @click="handleClick"
  >
    <img
      :src="getUserPhoto(user.photo)"
      alt="User avatar"
      class="user-avatar"
    />

    <div class="user-details">
      <span class="user-name">{{ user.name }}</span>
      <span v-if="isMember" class="member-badge">Membro</span>
      <span v-if="isSelected" class="checkmark">✔</span>
    </div>
  </div>
</template>
  
<style scoped>
.user-item {
  background: white;
  border-radius: 1rem;
  padding: 0.6rem 0.8rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  margin-bottom: 0.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.user-item:hover {
  background-color: #eef0ff;
}

.user-item.is-member {
  background-color: #e3e6ff;
  border-left: 4px solid navy;
}

.user-item.selected {
  background-color: #cfd5ff;
}

.user-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  object-fit: cover;
}

.user-details {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  width: 100%;
  justify-content: space-between;
}

.user-name {
  color: black;
  font-size: 14px;
  font-weight: 500;
}

.member-badge {
  background-color: navy;
  color: white;
  padding: 0.2rem 0.6rem;
  font-size: 0.75rem;
  border-radius: 999px;
}

.checkmark {
  color: green;
  font-size: 1rem;
  font-weight: bold;
}
</style>