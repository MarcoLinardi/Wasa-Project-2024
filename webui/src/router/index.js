import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', redirect: '/login' },
		{path: '/home', component: HomeView},
		{path: '/login', component: LoginView},
		{path: '/pathmatch(.*)*', redirect: '/login'},
	]
})

export default router
