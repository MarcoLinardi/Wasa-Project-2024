<script>
import axiosInstance from "@/services/axios";

export default {
  props: {
    selectedChat: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      messageText: "",
      messages: []
    };
  },
  mounted() {
    this.loadMessages(this.selectedChat.chatId)
  },
  methods: {
    getParticipantNames(chat) {
      if (chat && chat.isGroup && Array.isArray(chat.participants)) {
        return chat.participants.map(p => p.name).join(', ');
      }
      return '';
    },
    async loadMessages(chatId) {
      try {
        const response = await axiosInstance.get(`/chats/${chatId}/messages/status`);
        this.messages = response.data;
      } catch (e) {
        console.error(e);
      }
    },
    }
  }
</script>

<template>
  <div class="chat-area">
    
    <!-- HEADER -->
    <div class="chat-header" v-if="selectedChat">
      <img :src="selectedChat.photo" class="chat-photo" />
      <div class="chat-info">
        <h2>{{ selectedChat.name }}</h2>
        <h4 v-if="selectedChat.isGroup" class="participants">
          {{ getParticipantNames(selectedChat) }}
        </h4>
      </div>
    </div>

    <!-- BODY -->
    <div class="chat-body">
      
    </div>

    <!-- FOOTER -->
    <div class="chat-footer" v-if="selectedChat">
      <div class="input-container">
        <input
          v-model="messageText"
          placeholder="Scrivi un messaggio..."
          @keyup.enter="handleSend"
        />
        <button class="send-button" title="Scegli foto" >
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
        </button>
        <button class="send-button" title="Invia" @click="handleSend">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#send"/></svg>
        </button>
      </div>
    </div>

  </div>
</template>

<style scoped>
.chat-area {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  height: 100%;
  background-color: rgb(220, 212, 248);
  overflow: hidden;
}

.chat-header {
  height: 80px;
  background-color: rgb(220, 212, 248);
  border-bottom: 2px solid navy;
  display: flex;
  align-items: center;
  padding: 0 1rem;
  box-sizing: border-box;
  flex-shrink: 0;
}

.chat-photo {
  width: 4rem;
  height: 4rem;
  object-fit: cover;
  border-radius: 50%;
  margin-right: 1rem;
}

.chat-info {
  display: flex;
  flex-direction: column;
  justify-content: center;
  overflow: hidden;
}

.chat-info h2 {
  font-size: 1.3rem;
  margin: 0;
  color: #333;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.chat-info .participants {
  font-size: 0.9rem;
  margin: 2px 0 0 0;
  color: #222;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  font-weight: normal;
}

.chat-body {
  flex-grow: 1; 
  overflow-y: auto;
  background-color: whitesmoke;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 1rem;
}

.chat-footer {
  background-color: rgb(220, 212, 248);
  padding: 0.5rem;
  width: 100%;
  border-top: 2px solid navy;
  display: flex;
  align-items: center;
  justify-content: center;
  position: sticky;
  bottom: 0;
  z-index: 2;
}

.input-container {
  background-color: transparent;
  display: flex;
  align-items: center;
  width: 100%;
  padding: 0.3rem 0.8rem;
  border-radius: 30px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.input-container input {
  border: none;
  outline: none;
  flex-grow: 1;
  padding: 0.6rem;
  font-size: 1rem;
  background: transparent;
  color: #333;
}

.send-button {
  border: none;
  border-radius: 50%;
  height: 2.5rem;
  width: 2.5rem;
  background-color: transparent;
}

.send-button:hover {
  background-color: navy;
  color: rgb(220, 212, 248);
}

.send-button svg {
  transition: transform 0.2s ease;
}

.send-button:hover svg {
  transform: scale(1.25);
}
</style>
