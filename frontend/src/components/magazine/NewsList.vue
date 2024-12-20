<!-- NewsList.vue -->
<template>
  <div class="mt-10 container mx-auto p-6 bg-gradient-to-r from-neutral-300 to-stone-400 text-white rounded-md shadow-lg">
    <h1 class="text-4xl font-bold mb-6">News</h1>
    <div class="flex flex-col-reverse">
      <div v-for="news in newsList" :key="news.id" class="p-4 shadow-md bg-gray-900 rounded mb-4 break-words">
        <h2 class="text-2xl font-semibold">{{ news.title }}</h2>
        <p class="text-lg">{{ news.content }}</p>
        <p class="text-sm text-gray-400">{{ formatDate(news.date) }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';
import { alertService } from '../alertor';
import { useRouter } from 'vue-router';
import { refreshService } from '../refreshJWT';

export interface News {
  id: string;
  title: string;
  content: string;
  date: string;
}

export default defineComponent({
  setup() {
    const router = useRouter();
    const newsList = ref<News[]>([]);

    const fetchNews = async () => {
      try {
        const response = await axios.get('/api/magazine/v1/news');
        newsList.value = response.data.news;
      } catch (error: any) {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          await fetchNews();
        } else {
          alertService.showAlert(error.response.data.errors[0].message, "error");
          console.error(error);
        }
      }
    };

    const formatDate = (date?: string): string => {
      return date ? new Date(date).toLocaleString() : '';
    };

    onMounted(fetchNews);

    return {
      newsList,
      formatDate,
    };
  },
});
</script>

<style scoped>
/* Add any additional styles here */
</style>
