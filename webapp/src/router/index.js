import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/pages/Login.vue'
import Create from '@/pages/Create.vue'
import Test from '@/pages/Test.vue'
import Test_commit from '@/pages/Test_commit.vue'
import Practice from '@/pages/Practice.vue'
import Practice_commit from '@/pages/Practice_commit.vue'
import User from '@/pages/User.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },

    {
      path: '/login/create',
      name: 'create',
      component: Create
    },

    {
      path: '/login/user',
      name: 'user',
      component: User
    },

    {
      path: '/login/create/test',
      name: 'test',
      component: Test
    },

    {
      path: '/login/create/practice',
      name: 'practice',
      component: Practice
    },
    {
      path: '/login/create/test/testcommit',
      name: 'testcommit',
      component: Test_commit
    },

    {
      path: '/login/create/practice/practicecommit',
      name: 'practicecommit',
      component: Practice_commit
    }
  ]
})
