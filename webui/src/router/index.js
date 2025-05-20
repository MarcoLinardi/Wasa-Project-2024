import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import NewGroupView from '../views/NewGroupView.vue'
import UsersView from '../views/UsersView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login' },
		{path: '/login', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/profile', component: ProfileView},
		{path: '/newGroup', component: NewGroupView},
		{path: '/usersList', component: UsersView},
		{path: '/pathmatch(.*)*', redirect: '/login'},
	]
})

export default router
