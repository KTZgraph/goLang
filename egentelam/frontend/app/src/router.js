import Vue from 'vue'
import Router from 'vue-router'
import Users from '@/views/Users'
import Game from '@/views/Game'
import Cards from '@/views/Cards'
import Login from '@/views/Login'
import Logout from '@/views/Logout'
import Register from '@/views/Register'


Vue.use(Router)


export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/game',
      name: 'game',
      component: Game,
      // beforeEnter: ifAuthenticated,
    },
    {
      path: '/dzentelmani',
      name: 'users',
      component: Users,
      // beforeEnter: ifAuthenticated,
    },
    {
      path: '/karty',
      name: 'cards',
      component: Cards,
      // beforeEnter: ifAuthenticated,
    },
    {
      path: '/register',
      name: 'register',
      component: Register
    },
    {
      path: '/', //doylsnie
      name: 'login',
      component: Login
    },
    // {
    //   path: '/login', //doylsnie
    //   name: 'login',
    //   component: Login
    // },
    {
      path: '/logout',
      name: 'logout',
      component: Logout,
      // beforeEnter: ifNotAuthenticated,
    },
  ]
})
