<script>
export default {
  props: {
    selectedChat: Object,
    messages: Array
  }
}
</script>

<template>
    <div class="chat-area">
      
      <!-- HEADER: Nome della chat -->
      <div class="chat-header">
        <h2 v-if="selectedChat">
          {{ selectedChat.name }}
        </h2>
        <h2 v-else>
          Seleziona una chat
        </h2>
      </div>
  
      <!-- BODY: Messaggi -->
      <div class="chat-body">
        <div v-if="messages.length === 0">
          Nessun messaggio
        </div>
        <div v-else>
          <div 
            v-for="message in messages" 
            :key="message.id" 
            class="message"
          >
            {{ message.content }}
          </div>
        </div>
      </div>
  
    </div>
  </template>
  
  <style scoped>
  .chat-area {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    background-color: rgb(210, 180, 140);
  }
  
  .chat-header {
    height: 80px;
    background-color: #f0f2f5;
    border-bottom: 1px solid #ddd;
    display: flex;
    align-items: center;
    padding: 0 1rem;
    box-sizing: border-box;
  }
  
  .chat-header h2 {
    font-size: 1.2rem;
    margin: 0;
    color: #333;
  }
  
  .chat-body {
    flex-grow: 1;
    background-color: rgb(210, 180, 140);
    padding: 1rem;
    overflow-y: auto;
  }
  
  .message {
    background-color: white;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    margin-bottom: 0.5rem;
    max-width: 60%;
    word-wrap: break-word;
  }
  </style>
  