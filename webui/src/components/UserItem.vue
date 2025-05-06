<template>
  <div class="user-item p-4 border-b cursor-pointer">
    <!-- Immagine utente -->
    <img :src="user.photo" alt="User" class="user-avatar" />
    <div class="font-semibold">{{ user.name }}</div>
  </div>
</template>
  
<script>
  export default{
    props: {
        user: {
          type: Object,
          required: true
        }
    },
    mounted() {
      console.log("UserItem mounted:", this.user);
    }
  }
  
</script>
  
  
<style scoped>
  .user-item {
    display: flex;
    align-items: center;
    padding: 8px 12px;
    border-radius: 6px;
    cursor: pointer;
    gap: 12px;
    transition: background-color 0.2s;
  }

  .user-item:hover {
    background-color: rgb(217, 128, 91);
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
  