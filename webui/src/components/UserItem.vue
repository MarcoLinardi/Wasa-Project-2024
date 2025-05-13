<script>
  export default{
    props: {
        user: {
          type: Object,
          required: true
        }
    },
    mounted() {

    },
    methods: {
      handleClick() {
        this.$emit('select-user', this.user);
    }
    }
  }
  
</script>

<template>
  <div class="user-item p-4 border-b cursor-pointer" @click="handleClick" >
    <img :src="user.photo" class="user-avatar" />
    <div class="font-semibold">{{ user.name }}</div>
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
  