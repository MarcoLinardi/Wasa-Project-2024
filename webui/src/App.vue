<script setup>
import { RouterView, useRoute, useRouter } from 'vue-router';

const route = useRoute(); // Istanza per accedere alle informazioni della route corrente
const router = useRouter(); // Istanza per la navigazione programmatica

function  logout() {
  localStorage.removeItem('token');
  localStorage.removeItem('user')
  console.log('Logout effettuato');
  router.push('/login');
}
</script>

<template>
  <div id="app-layout" :class="{ 'no-sidebar': route.path === '/login' }">
    <aside v-if="route.path !== '/login'" class="app-sidebar">
      <div class="sidebar-header">
        <h3>WasaText</h3>
      </div>
      <nav class="sidebar-navigation">
        <button class="sidebar-button" @click="router.push('/profile')"> 
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
          <span>Profilo</span>
        </button>
        <button class="sidebar-button" @click="router.push('/home')"> 
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-square"/></svg>
          <span>Chat</span>
        </button>
        <button class="sidebar-button" @click="router.push('/usersList')"> 
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#users"/></svg>
          <span>Utenti</span>
        </button>
        <button class="sidebar-button" @click="router.push('/newGroup')">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="feather"
          >
            <path stroke-linecap="round" 
              stroke-linejoin="round" 
              d="M18 18.72a9.094 9.094 0 0 0 3.741-.479 3 3 0 0 0-4.682-2.72m.94 3.198.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0 1 12 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 0 1 6 18.719m12 0a5.971 5.971 0 0 0-.941-3.197m0 0A5.995 5.995 0 0 0 12 12.75a5.995 5.995 0 0 0-5.058 2.772m0 0a3 3 0 0 0-4.681 2.72 8.986 8.986 0 0 0 3.74.477m.94-3.197a5.971 5.971 0 0 0-.94 3.197M15 6.75a3 3 0 1 1-6 0 3 3 0 0 1 6 0Zm6 3a2.25 2.25 0 1 1-4.5 0 2.25 2.25 0 0 1 4.5 0Zm-13.5 0a2.25 2.25 0 1 1-4.5 0 2.25 2.25 0 0 1 4.5 0Z" />
          </svg>
          <span>Crea Gruppo</span>
        </button>
        <button class="sidebar-button" title="Refresh" @click="() => router.go(0)"> 
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#refresh-cw"/></svg>
          <span>Refresh</span>
        </button>
        <button class="sidebar-button" title="Logout" @click="logout">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
          <span>Logout</span>
        </button>
      </nav>
    </aside>

    <main class="content-area">
      <RouterView />
    </main>
  </div>
</template>

<style>
#app-layout {
  display: flex;
  height: 100vh;
}

.app-sidebar {
  width: 18%;
  border-right: 1px solid antiquewhite;
  background-color: rgb(34, 34, 133);
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.sidebar-header h3 {
  margin-top: 0;
  color: antiquewhite;
}

.sidebar-navigation {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  margin-top: 2rem;
}

.sidebar-button {
  background: navy;
  border: none;
  border-radius: 10px;
  padding: 10px;
  text-align: left;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px; /* Spazio tra icona e testo */
  color: antiquewhite;
  box-shadow: 6px 6px 10px rgba(0,0,0,1),
  1px 1px 10px rgba(255, 255, 255, 0.6);
  transition: background-color 0.3s ease;
}

.sidebar-button:hover {
  background-color: #6159d9;
}

.sidebar-button:active {
  background-color: #6159d9;
}

.feather {
  width: 20px;
  height: 20px;
}

.content-area {
  flex-grow: 1;
  background-image: url('@/assets/AppBackground.jpg');
  background-size: cover;         /* riempie tutto */
  background-position: center;    /* centrato */
  background-repeat: no-repeat;
}

#app-layout.no-sidebar .content-area {
  width: 100%;
}
</style>