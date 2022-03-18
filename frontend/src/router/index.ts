import { RouteRecordRaw, createRouter, createWebHistory} from 'vue-router'

import Homepage from "../views/Homepage.vue"
import Bookspage from "../views/Bookspage.vue"
import Genrespage from "../views/Genrespage.vue"
import Authorspage from "../views/Authorspage.vue"
import Loginpage from "../views/Loginpage.vue"
import SingleBookpage from "../views/SingleBookpage.vue"

const routes: RouteRecordRaw[] = [
	{
		path: '/',
		name: 'Home',
		component: Homepage,
	},
	{
		path: '/books',
		name: 'Books',
		component: Bookspage,
	},
	{
		path: '/books/:bookId',
		name: 'SingleBook',
		component: SingleBookpage,
	},
	{
		path: '/genres',
		name: 'Genres',
		component: Genrespage,
	},
	{
		path: '/authors',
		name: 'Authors',
		component: Authorspage,
	},
	{
		path: '/login',
		name: 'Login',
		component: Loginpage,
	},
	{
		path: '/signup',
		name: 'Signup',
		component: Loginpage,
	},
]

const router = createRouter({
	history: createWebHistory(),
	routes,
})

export default router