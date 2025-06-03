import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import UserItem from './components/UserItem.vue';
import ChatItem from './components/ChatItem.vue';
import ChatArea from './components/ChatArea.vue';
import NewChat from './components/NewChat.vue';
import Message from './components/Message.vue';
import ChatInfo from './components/ChatInfo.vue';
import ManageMember from './components/ManageMember.vue';

import './assets/dashboard.css'
import './assets/main.css'


const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("UserItem", UserItem)
app.component("ChatItem", ChatItem)
app.component("ChatArea", ChatArea)
app.component("ChatInfo", ChatInfo)
app.component("ManageMember", ManageMember)
app.component("NewChat", NewChat)
app.component("Message", Message)
app.use(router)
app.mount('#app')
