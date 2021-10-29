import {
	createRouter,
	createWebHistory
} from 'vue-router'
import Login from '../views/Login.vue'
import Home from "../views/Home.vue"

const routes = [{
		path: '/',
		name: 'Login',
		component: Login
	},
	{
		path: '/home',
		name: 'Home',
		component: Home,
		children: [{
				path: "/home/all",
				name: "AllFolder",
				component: () => import("../views/AllFolder.vue"),
			},
			{
				path: "/home/all/folder/files",
				name: "FolderFile",
				component: () => import("../views/FolderFiles.vue"),
			},
			{
				path: "/home/delete",
				name: "Delete",
				component: () => import("../views/Delete.vue")
			},
			{
				path: "/home/message",
				name: "Message",
				component: () => import("../views/Message.vue")
			},
			{
				path: "/home/pass",
				name: "Password",
				component: () => import("../views/Password.vue")
			},
			{
				path: "/home/upload",
				name: "Upload",
				component: () => import("../components/Upload.vue")
			}
		],
	},
]

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes
})

export default router