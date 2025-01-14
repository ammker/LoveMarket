import { createRouter, createWebHashHistory } from 'vue-router'
import Squre from '../views/Square.vue'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Postposts from '../views/Postposts.vue'
import ProfilePage from '../views/ProfilePage.vue';
import SettingsPage from '../views/SettingsPage.vue';
import ChatLayout from '../views/ChatLayout.vue'
import ChatView from '../components/ChatView.vue'
import Myposts from '../views/view_my_posts.vue'


const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            redirect: '/login'
        },
        {
            path: '/login',
            component: Login
        },
        {
            path: '/register',
            component: Register
        },
        {   
            path: '/home',
            name: 'home',
            component: Home
        },
        {   
            path: '/square',
            name: 'square',
            component: Squre
        },
        {
            path: '/posts/:id',
            name: 'posts',
            component: () => import ('../views/Postdetail.vue')
        },
        {
            path: '/postposts',
            name: 'postposts',
            component: Postposts
        },
        {
            path: '/search/:searchdetail',
            name: 'search',
            component: () => import ('../views/Searchpost.vue')
        },
        {
            path: '/profile',
            name: 'ProfilePage',
            component: ProfilePage
        },
        {
            path: '/settings',
            name: 'SettingsPage',
            component: SettingsPage
        },
        {
            path: '/chat',
            name:'chat',
            component: ChatLayout,
            children: [
              {
                path: ':friendId',
                name: 'ChatView',
                component: ChatView
              }
            ]
        },
        {
            path: '/myposts',
            name:'myposts',
            component: Myposts,
        }
    ]
})

export default router
