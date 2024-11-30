import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/kon',
      name: 'kon',
      component: () => import('@/views/Koni.vue'),
    },
    {
      path: '/eventor',
      name: 'eventor',
      component: () => import('@/views/Eventor.vue'),
    },
  ],
})

export default router

// import { createRouter, createWebHistory } from 'vue-router'
// import Koni from '@/views/Koni.vue'
// import Eventor from '@/views/Eventor.vue'

// const router = createRouter({
//   history: createWebHistory(import.meta.env.BASE_URL),
//   routes: [
//     {
//       path: '/kon',
//       name: 'kon',
//       component: Koni
//       // component: () => import('@/components/Koni.vue'),
//     },
//     {
//       path: '/eventor',
//       name: 'eventor',
//       component: Eventor
//       // component: () => import('@/components/Koni.vue'),
//     },
//   ],
// })

// export default router
