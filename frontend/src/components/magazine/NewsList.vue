<!-- NewsList.vue -->
<template>
  <div class="container mx-auto p-6 bg-gradient-to-r from-green-400 to-blue-500 text-white rounded-md shadow-lg">
    <h1 class="text-4xl font-bold mb-6">News</h1>
    <div v-for="news in newsList" :key="news.id" class="p-4 bg-gray-900 rounded mb-4 break-words">
      <h2 class="text-2xl font-semibold">{{ news.title }}</h2>
      <p class="text-lg">{{ news.content }}</p>
      <p class="text-sm text-gray-400">{{ news.date }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';

export interface News {
  id: string;
  title: string;
  content: string;
  date: string;
}

export default defineComponent({
  setup() {
    const newsList = ref<News[]>([]);

    const fetchNews = async () => {
      try {
        const response = await axios.get('/news');
        newsList.value = response.data.news;
      } catch (error) {
        console.error('Error fetching news:', error);
      }
    };

    onMounted(fetchNews);

    return {
      newsList,
    };
  },
});
</script>

<style scoped>
/* Add any additional styles here */
</style>
