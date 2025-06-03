<script>
export default {
  props: {
    messageData: {
      type: Object,
      required: true,
    },
    loggedUser: {
      type: Object,
      required: true
    }
  },
  computed: {
    isSentByLoggedUser() {
      // Confronta il senderId del messaggio con l'ID dell'utente corrente
      return this.messageData.senderId === this.loggedUser.userId;
    },
    formattedTimestamp() {
      if (!this.messageData.timestamp) return '';
      const date = new Date(this.messageData.timestamp);
      // Formatta l'ora come HH:MM
      return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', hour12: false });
    }
  }
}
</script>

<template>
  <div class="message-wrapper" :class="{ 'sent-wrapper': isSentByLoggedUser, 'received-wrapper': !isSentByLoggedUser }">
    <div class="message-item" :class="{ 'sent': isSentByLoggedUser, 'received': !isSentByLoggedUser }">
      <div class="message-content">
        <p class="message-text">{{ messageData.content }}</p>
        <span class="message-timestamp">{{ formattedTimestamp }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.message-wrapper {
  display: flex;
  margin-bottom: 0.6rem;
  width: 100%;
}

.message-wrapper.sent-wrapper {
  justify-content: flex-end;
}

.message-wrapper.received-wrapper {
  justify-content: flex-start;
}

.message-item {
  max-width: 60%;
  word-wrap: break-word;
}

.message-content {
  padding: 10px 15px;
  border-radius: 1.3rem;
  position: relative;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}


.message-item.sent .message-content {
  background-color: #6565deca;
  color: black;
  border-bottom-right-radius: 0.2rem;
}

.message-item.received .message-content {
  background-color: rgba(78, 139, 183, 0.675);
  color: black;
  border: 1px solid #f0f0f0; 
  border-bottom-left-radius: 0.2rem;
  
}

.sender-name {
  font-size: 0.8em;
  font-weight: bold;
  margin-bottom: 4px;
  color: #555;
}

.message-text {
  margin: 0 0 5px 0;
  white-space: pre-wrap;
  font-size: 0.95em;
  line-height: 1.4;
}

.message-timestamp {
  font-size: 0.75em;
  color: rgba(0, 0, 0, 0.512);
  display: block;
  text-align: right;
  margin-top: 4px;
}

</style>