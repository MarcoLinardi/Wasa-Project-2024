import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Sidebar from './components/Sidebar.vue';
import ChatArea from './components/ChatArea.vue';
import ChatList from './components/ChatList.vue';
import UserList from './components/UserList.vue';

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Sidebar", Sidebar)
app.component("ChatArea", ChatArea)
app.component("ChatList", ChatList)
app.component("UserList", UserList)
app.use(router)
app.mount('#app')
