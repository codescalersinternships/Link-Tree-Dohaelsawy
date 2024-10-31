import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/Auth/login.vue'
import Register from '../views/Auth/register.vue'
import Links from '../views/links/links.vue'
import LiveDemo from '@/views/links/liveDemo.vue'
import Profile from "@/views/account/profile.vue"
import analytics from "@/views/account/analytics.vue"




const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/auth/login',
      name: 'auth-login',
      component: Login
    },
    {
      path: '/auth/register',
      name: 'auth-register',
      component: Register
    },
    {
      path: '/link_tree/',
      name: 'links',
      component: Links
    },
    {
      path: '/link_tree/:username',
      name: 'links-demo',
      component: LiveDemo
    },
    {
      path: '/account/get_account/',
      name: 'profile',
      component: Profile
    },
    {
      path: '/get_analytics/:user_id',
      name: 'analytics',
      component: analytics
    },
  ]
})

export default router
