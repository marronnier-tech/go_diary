import Vue from 'vue'
import Router from 'vue-router'
import Top from '@/components/Top'
import Service from '@/components/Service'
import Todo from '@/components/Todo'
import Profile from '@/components/Profile'
import Goal from '@/components/Goal'
import Mypage from '@/components/Mypage'
import Login from '@/components/Login'
import Register from '@/components/Register'
import Logout from '@/components/Logout'


Vue.use(Router)


export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Top',
      component: Top
    },
    {
      path: '',
      component: Service,
      children: [
        {
          path: '/mypage',
          name: 'Mypage',
          component: Mypage
        },
        {
          path: '/todo',
          name: 'Todo',
          component: Todo
        },
        {
          path: '/profile',
          name: 'Profile',
          component: Profile
        },
        {
          path: '/goal',
          name: 'Goal',
          component: Goal
        }
      ]
    },
    {

      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Regiter',
      component: Register
    },
    {
      path: '/logout',
      name: 'Logout',
      component: Logout
    }
  ]
})