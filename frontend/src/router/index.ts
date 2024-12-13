import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/eventor',
    },
    {
      path: '/discussion',
      name: 'discussion',
      component: () => import('@/views/Discussion.vue'),
    },
    {
      path: '/blog',
      name: 'blog',
      component: () => import('@/views/Blog.vue'),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/Login.vue'),
    },
    {
      path: '/magazine',
      name: 'magazine',
      component: () => import('@/views/Magazine.vue'),
    },
    {
      path: '/support',
      name: 'support',
      component: () => import('@/views/Support.vue'),
    },
    {
      path: '/eventor',
      name: 'eventor',
      component: () => import('@/views/Eventor.vue'),
    },
    {
      path: '/blog/manage',
      name: 'blog-management',
      component: () => import('@/components/blog/ArticleManagement.vue'),
    },
    {
      path: '/article/:articleId',
      name: 'article',
      component: () => import('@/components/blog/DisplayedArticle.vue'),
      props: true,
    },
  ],
})

export default router
