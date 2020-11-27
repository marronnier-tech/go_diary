import Vue from 'vue'
import Router from 'vue-router'
import Top from '@/components/Top'
import Service from '@/components/Service'
import Todo from '@/components/Todo'
import Profile from '@/components/Profile'
import EditProfile from '@/components/EditProfile'
import MyGoal from '@/components/MyGoal'
import Mypage from '@/components/Mypage'
import Users from '@/components/UsersTodo'
import Profs from '@/components/UserProf'
import QandA from '@/components/QandA'
import Login from '@/components/Login'
import Register from '@/components/Register'
import Logout from '@/components/Logout'


Vue.use(Router)


export default new Router({
  mode: 'history',
  routes: [
    {
      path: '',
      component: Service,
      children: [
        {
          path: '/',
          name: 'Top',
          component: Top
        },
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
          path: '/edit-profile',
          name: 'EditProfile',
          component: EditProfile
        },
        {
          path: '/mygoal',
          name: 'MyGoal',
          component: MyGoal
        },
        {
          path: '/logout',
          name: 'Logout',
          component: Logout
        },
        {
          path: '/user/:id',
          name: 'Users',
          component: Users
        },
        {
          path: '/profile/:id',
          name: 'Profs',
          component: Profs
        },
        {
          path: '/question',
          name: 'QandA',
          component: QandA
        },
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

  ]
})