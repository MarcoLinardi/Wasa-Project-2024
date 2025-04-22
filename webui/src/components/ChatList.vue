<script>
export default {
  props: {
    chats: {
      type: Array,
      default: () => [],
      required: true
    }
  }
}
</script>

<template>
    <ul class="list">
      <li 
        v-for="chat in chats" 
        :key="chat.id" 
        class="list-item"
        @click="$emit('selectChat', chat)"
      >
        {{ chat.name || "Chat senza nome" }}
      </li>
    </ul>
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
    border-bottom: 1px solid #ddd;
  }
  
  .list-item:hover {
    background-color: #ececec;
  }
  </style>
  