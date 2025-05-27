import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import UserItem from './components/UserItem.vue';
import UserProfile from './components/UserProfile.vue';
import ChatItem from './components/ChatItem.vue';
import ChatArea from './components/ChatArea.vue';
import NewChat from './components/NewChat.vue';

import './assets/dashboard.css'
import './assets/main.css'




const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("UserItem", UserItem)
app.component("UserProfile", UserProfile)
app.component("ChatItem", ChatItem)
app.component("ChatArea", ChatArea)
app.component("NewChat", NewChat)
app.use(router)
app.mount('#app')
