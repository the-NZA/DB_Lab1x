import { RouteRecordRaw, createRouter, createWebHistory} from 'vue-router'

import Homepage from "../views/Homepage.vue"
import Bookspage from "../views/Bookspage.vue"
import Genrespage from "../views/Genrespage.vue"
import Authorspage from "../views/Authorspage.vue"

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
		path: '/genres',
		name: 'Genres',
		component: Genrespage,
	},
	{
		path: '/authors',
		name: 'Authors',
		component: Authorspage,
	},
]

const router = createRouter({
	history: createWebHistory(),
	routes,
})

export default router