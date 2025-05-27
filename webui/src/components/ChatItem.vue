<script>
export default {
  props: {
    chat: {
      type: Object,
      required: true,
    },
  },
  computed: {
    formattedTimestamp() {
      if (!this.chat.lastMessage || !this.chat.lastMessage.timestamp) {
        return '';
      }
      const messageDate = new Date(this.chat.lastMessage.timestamp);
      const now = new Date();

      // Formato "HH:MM" se è oggi
      if (
        messageDate.getDate() === now.getDate() &&
        messageDate.getMonth() === now.getMonth() &&
        messageDate.getFullYear() === now.getFullYear()
      ) {
        return messageDate.toLocaleTimeString('it-IT', { hour: '2-digit', minute: '2-digit' });
      }

      // Formato "ieri" se è ieri
      const yesterday = new Date(now);
      yesterday.setDate(now.getDate() - 1);
      if (
        messageDate.getDate() === yesterday.getDate() &&
        messageDate.getMonth() === yesterday.getMonth() &&
        messageDate.getFullYear() === yesterday.getFullYear()
      ) {
        return 'ieri';
      }

      // Formato "dd/mm/yyyy" per date più vecchie
      return messageDate.toLocaleDateString('it-IT');
    },
    profileImageUrl() {
      return this.chat.photo
    }
  },
  methods: {
    selectChat() {
      console.log('Chat selezionata:', this.chat.name);
      this.$emit('select-chat', this.chat);
    }
  }
};
</script>

<template>
  <div class="chat-item" @click="selectChat">
    <div class="chat-avatar">
      <img :src="profileImageUrl" :alt="chat.name" class="avatar-img" />
      </div>
    <div class="chat-content">
      <div class="chat-header">
        <span class="chat-name">{{ chat.name }}</span>
        <span class="chat-timestamp">{{ formattedTimestamp }}</span>
      </div>
      <div class="chat-last-message" v-if="chat.lastMessage">
        <span class="message-text">{{ chat.lastMessage.content }}</span>
      </div>
    </div>
    </div>
</template>

<style scoped>
.chat-item {
  display: flex;
  align-items: center;
  padding: 5px 8px;
  cursor: pointer;
  transition: background-color 0.2s ease;
  color: whitesmoke;
  font-family: Arial, sans-serif;
}

.chat-item:hover {
  background-color: navy;
}

.chat-avatar {
  position: relative;
  margin-right: 15px;
  flex-shrink: 0;
}

.avatar-img {
  width: 45px;
  height: 45px;
  border-radius: 50%;
  object-fit: cover;
}

.chat-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.chat-name {
  color: whitesmoke;
  font-size: 1em;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.chat-timestamp {
  font-size: 0.85em;
  color: #999; /* Timestamp grigio chiaro */
  white-space: nowrap;
  margin-left: 10px; /* Spazio tra nome e timestamp */
}

.chat-last-message {
  font-size: 0.9em;
  color: #b0b0b0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>