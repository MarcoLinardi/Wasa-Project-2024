<script>
  export default{
    props: {
        user: {
          type: Object,
          required: true
        },
        isSelected: {
          type: Boolean,
          required: true,
          default: false
        }
    },
    mounted() {

    },
    methods: {
      handleClick() {
        console.log("UserItem cliccato:", this.user.name);
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
    class="user-item p-4 border-b cursor-pointer flex items-center gap-4 rounded-md"
    :class="{ 'bg-orange-300 text-white font-semibold': isSelected }"
    @click="handleClick"
  >
    <img
      :src="getUserPhoto(user.photo)"
      class="user-avatar w-10 h-10 rounded-full object-cover"
      alt="User avatar"
    />
    <div class="flex justify-between items-center w-full">
      <span>{{ user.name }}</span>
      <span v-if="isSelected" class="ml-2">✔️</span>
    </div>
  </div>
</template>
  
<style scoped>
  .user-item {
    display: flex;
    align-items: center;
    padding: 8px 12px;
    border-radius: 6px;
    cursor: pointer;
    gap: 12px;
    height: 4rem;
    transition: background-color 0.2s;
  }

  .user-item:hover {
    background-color: #007bff;
  }

  .user-avatar {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    object-fit: cover;
    border: none;
    box-shadow: none;
  }

  .user-name {
    color: white;
    font-size: 14px;
    font-weight: 500;
  }
</style>
  