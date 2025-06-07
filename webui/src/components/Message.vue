<script>
import axiosInstance from "@/services/axios";
export default {
  data() {
    return {
      showMenu: false,
      emojiPickerVisible: false,
      reaction: null,
      showReactionInfo: false,
      users: []
    };
  },
  props: {
    chat: {
      type: Object,
      required: true
    },
    message: {
      type: Object,
      required: true,
    },
    loggedUser: {
      type: Object,
      required: true
    }
  },
  mounted() {
    document.addEventListener("click", this.handleOutsideClick);
    this.loadUsers()
  },
  computed: {
    isSentByLoggedUser() {
      return this.message.senderId === this.loggedUser.userId;
    },
    formattedTimestamp() {
      if (!this.message.timestamp) return '';
      const date = new Date(this.message.timestamp);
      // Formatta l'ora come HH:MM
      return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', hour12: false });
    },
    visibleReactions() {
      return this.message.reactions.slice(0, 2);
    },
    hiddenReactions() {
      return this.message.reactions.slice(2).map(r => r.reaction).join(' ');
    }
    
  },
  methods: {
    toggleMenu(event) {
      this.showMenu = !this.showMenu;
      // evita chiusura immediata del menu cliccando sul bottone
      event.stopPropagation();
    },
    handleOutsideClick(event) {
      // Chiudi menu opzioni
      if (this.showMenu && !this.$el.contains(event.target)) {
        this.showMenu = false;
      }

      // Chiudi popup reaction info
      if (this.showReactionInfo && !this.$el.contains(event.target)) {
        this.showReactionInfo = false;
      }
    },
    onReply() {
      this.$emit("reply", this.message);
      this.showMenu = false;
    },
    onForward() {
      this.$emit("forward", this.message);
      this.showMenu = false;
    },
    onDelete() {
      this.$emit("delete", this.message);
      this.showMenu = false;
    },
    toggleEmojiPicker() {
      this.emojiPickerVisible = !this.emojiPickerVisible;
    },
    async react(reaction) {
      try {
        const response = await axiosInstance.post(`/chats/${this.chat.chatId}/messages/${this.message.messageId}/reactions`, { reaction });

        this.message.reaction = reaction;
        this.emojiPickerVisible = false;
        this.showMenu = false;
        this.$emit("reload-messages")
        } catch (err) {
          console.error('Errore di rete:', err);
        }
    },
    async removeReaction(reaction) {
      try {
        await axiosInstance.delete(`/chats/${this.chat.chatId}/messages/${this.message.messageId}/reactions`, {
          data: { reaction }
        });
        this.$emit("reload-messages");
      } catch (err) {
        console.error("Errore nella rimozione della reaction:", err);
      }
    },

    openReactionInfo(event) {
      this.showReactionInfo = !this.showReactionInfo;
      event.stopPropagation();
    },
    getUserNameById(userId) {
      if (userId === this.loggedUser.userId) {
        return "Tu";
      }
      const user = this.users.find(u => u.userId === userId);
      return user ? user.name : `Utente ${userId}`;
    },
    async loadUsers() {
      try {
        const response = await axiosInstance.get("/users");
        this.users = response.data.users || response.data;
        this.users.sort((a, b) => (a.name || "").localeCompare(b.name || ""));
      } catch (e) {
        console.error("[loadUsers] Errore nel caricamento degli utenti:", e);
        this.users = [];
      }
    },
  }
}
</script>

<template>
  <div class="message-wrapper" :class="{ 'sent-wrapper': isSentByLoggedUser, 'received-wrapper': !isSentByLoggedUser }">
   <div class="message-item" :class="{ 'sent': isSentByLoggedUser, 'received': !isSentByLoggedUser }">
      <!-- Wrapper: messaggio + menu -->
      <div class="message-body-wrapper" :class="{ 'sent-align': isSentByLoggedUser, 'received-align': !isSentByLoggedUser }">
        
        <!-- Bottone menu -->
        <button class="message-menu-btn" @click="toggleMenu">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#chevron-up"/></svg>
        </button>

        <!-- Messaggio -->
        <div class="message-content">
          <p class="message-text">{{ message.content }}</p>
          <span class="message-timestamp">{{ formattedTimestamp }}</span>
        </div>

        <!-- Reaction -->
        <div v-if="message.reactions && message.reactions.length" class="reaction-group">
          <span v-for="(r, index) in visibleReactions" :key="r.userId + '-' + index" class="reaction" @click="openReactionInfo">
            {{ r.reaction }}
          </span>
          <span v-if="message.reactions.length > 2" class="reaction-dots" @click="openReactionInfo">...</span>
          
          <!-- Popup con dettagli -->
          <div v-if="showReactionInfo" 
          class="reaction-info-popup" 
          :class="{ 'popup-left': isSentByLoggedUser, 'popup-right': !isSentByLoggedUser }" 
          ref="reactionPopup"
          >
            <div v-for="(r, index) in message.reactions" :key="'info-' + index" class="reaction-info-row">
              <span class="reaction-icon">{{ r.reaction }}</span>
              <span class="reaction-user">{{ getUserNameById(r.userId) }}</span>
              <!-- Mostra bottone di rimozione solo se la reaction è dell'utente loggato -->
              <button
                v-if="r.userId === loggedUser.userId"
                class="remove-reaction-btn"
                @click.stop="removeReaction(r.reaction)"
                title="Rimuovi la tua reaction">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#minus"/></svg>
              </button>
            </div>
          </div>
        </div>
        
        <!-- Menù a tendina -->
        <div v-if="showMenu" class="message-menu-dropdown" :class="isSentByLoggedUser ? 'dropdown-left' : 'dropdown-right'">
          <ul>
            <li @click="onReply">Rispondi</li>
            <li @click="onForward">Inoltra</li>
            <li @click="onDelete">Elimina</li>
            <li @click="toggleEmojiPicker">Reagisci</li>
          </ul>
          <div v-if="emojiPickerVisible" class="emoji-picker">
            <span @click="react('👍')">👍</span>
            <span @click="react('😂')">😂</span>
            <span @click="react('❤️')">❤️</span>
            <span @click="react('😮')">😮</span>
            <span @click="react('😢')">😢</span>
          </div>
        </div>
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
  padding: 0.7rem 1rem;
  border-radius: 1.3rem;
  position: relative;
  box-shadow: 0 0.2rem 0.3rem rgba(0, 0, 0, 0.1);
}

/* Mostra il bottone solo quando si passa sopra l’intero messaggio */
.message-item:hover .message-menu-btn {
  opacity: 1;
}

.message-item.sent .message-content {
  background-color: #6565deca;
  color: black;
  border-bottom-right-radius: 0.2rem;
}

.message-item.received .message-content {
  background-color: rgba(78, 139, 183, 0.675);
  color: black;
  border-bottom-left-radius: 0.2rem;
  
}

.sender-name {
  font-size: 0.8em;
  font-weight: bold;
  margin-bottom: 0.25;
  color: #555;
}

.message-text {
  margin: 0 0 0.40rem 0;
  white-space: pre-wrap;
  font-size: 0.95em;
  line-height: 1.4;
}

.message-timestamp {
  font-size: 0.75em;
  color: rgba(0, 0, 0, 0.512);
  display: block;
  text-align: right;
  margin-top: 0.25rem;
}

.message-body-wrapper {
  display: flex;
  align-items: flex-start;
  position: relative;
  margin-bottom: 0.25rem;
}

.sent-align {
  flex-direction: row;
  justify-content: flex-end;
}

.received-align {
  flex-direction: row-reverse;
  justify-content: flex-start;
}

.message-menu-btn {
  background: transparent;
  border: none;
  font-size: 1rem;
  cursor: pointer;
  margin: 0 0.37rem;
  padding: 0;
  opacity: 0;
  transition: opacity 0.2s;
}

.message-content {
  position: relative;
  max-width: 18rem;
  word-break: break-word;
}

.message-menu-dropdown {
  position: absolute;
  top: 0;
  display: flex;
  flex-direction: row-reverse;
  transform: translateY(-6rem); /* allinea l'ultimo elemento al bottone */
  background: #c7c6e4;
  border: 1px solid #ccc;
  border-radius: 0.6rem;
  box-shadow: 0 0.15rem 0.37rem rgba(0, 0, 0, 0.2);
  z-index: 100;
  min-width: 100px;
}

.dropdown-left {
  right: 100%;
  margin-right: 0.37rem;
}

.dropdown-right {
  left: 100%;
  margin-left: 0.37rem;
}

.message-menu-dropdown ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.message-menu-dropdown li {
  padding: 0.37rem 0.75rem;
  cursor: pointer;
}

.message-menu-dropdown li:hover {
  background-color: #dcdbe3;
}

.emoji-picker {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
  background-color: white;
  border-radius: 8px;
  padding: 6px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.reaction-group {
  position: absolute;
  bottom: -0.6rem; /* metà dentro, metà fuori */
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 0.25rem;
  align-items: center;
  z-index: 1;
}

.reaction {
  font-size: 1.1rem;
  cursor: pointer;
  user-select: none;
}

.reaction-dots {
  font-size: 1.1rem;
  color: #444;
  margin-left: 0.25rem;
  user-select: none;
}

.reaction-info-popup {
  position: absolute;
  top: -100%;
  min-width: 10rem;
  background: #c7c6e4;
  border: 1px solid navy;
  border-radius: 6px;
  padding: 6px 8px;
  box-shadow: 0 2px 6px rgba(0,0,0,0.2);
  font-size: 0.9rem;
  z-index: 10;
}

/* Se il messaggio è a destra, il popup va a sinistra */
.popup-left {
  right: 100%;
  margin-right: 8px;
}

/* Se il messaggio è a sinistra, il popup va a destra */
.popup-right {
  left: 100%;
  margin-left: 8px;
}

.reaction-info-row {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 4px 0;
  border-bottom: 1px solid rgba(0, 0, 128, 0.345);
}

.reaction-info-row:last-child {
  border-bottom: none;
}

.reaction-icon {
  font-size: 1.2rem;
}

.reaction-user {
  flex-grow: 1;
  white-space: nowrap;
}

.remove-reaction-btn {
  margin-left: auto;
  background: transparent;
  border: none;
  color: #888;
  cursor: pointer;
  font-size: 0.9rem;
}

.remove-reaction-btn:hover {
  color: red;
}

</style>