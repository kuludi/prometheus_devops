import Vue from 'vue'
import VueRouter from 'vue-router'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Login',
    component: () => import('../components/Login.vue')
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('../components/Home.vue'),
    redirect: "/welcome",
    children: [
      {
        path: "/welcome",
        component: () => import('../views/Default')
      },

      {
        path: "/users",
        component: () => import('../views/User')
      },
      // {
      //   path: "/roles",
      //   component: Role
      // },
      // {
      //   path: "/rights",
      //   component: Right
      // }
    ]
  },


  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = new VueRouter({
  routes
})

export default router
