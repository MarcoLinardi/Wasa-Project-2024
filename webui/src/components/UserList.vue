<script>
export default {
  props: {
    users: {
      type: Array,
      required: true
    }
  }
}
</script>

<template>
    <ul class="list">
      <li 
        class="list-item special"
        @click="$emit('createGroup')"
      >
        ➕ Crea Gruppo
      </li>
  
      <li 
        v-for="user in users" 
        :key="user.id" 
        class="list-item"
        @click="$emit('startChat', user)"
      >
        {{ user.name }}
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
  
  .special {
    font-weight: bold;
    color: #2e7d32;
  }
  </style>
  