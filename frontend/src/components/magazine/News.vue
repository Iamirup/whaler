<template>
  <div class="container mx-auto p-6 bg-gradient-to-r from-green-400 to-blue-500 text-white rounded-md shadow-lg">
    <h1 class="text-4xl font-bold mb-6">News</h1>
    <form @submit.prevent="addNews" class="mb-6">
      <div class="mb-4">
        <label for="title" class="block text-lg font-semibold mb-2">Title</label>
        <input v-model="newTitle" id="title" type="text" class="w-full p-2 rounded bg-gray-800 text-white">
      </div>
      <div class="mb-4">
        <label for="content" class="block text-lg font-semibold mb-2">Content</label>
        <textarea v-model="newContent" id="content" class="w-full p-2 rounded bg-gray-800 text-white"></textarea>
      </div>
      <button type="submit" class="px-4 py-2 bg-blue-700 rounded hover:bg-blue-800">Add News</button>
    </form>
    <div v-for="news in newsList" :key="news.id" class="p-4 bg-gray-900 break-words rounded mb-4">
      <h2 class="text-2xl font-semibold">{{ news.title }}</h2>
      <p class="text-lg">{{ news.content }}</p>
      <p class="text-sm text-gray-400">{{ news.date }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';

interface News {
  id: string;
  title: string;
  content: string;
  date: string;
}

export default defineComponent({
  setup() {
    const newsList = ref<News[]>([]);
    const newTitle = ref('');
    const newContent = ref('');

    const fetchNews = async () => {
      try {
        const response = await axios.get('/api/magazine/v1/news');
        newsList.value = response.data.news;
      } catch (error) {
        console.error('Error fetching news:', error);
      }
    };

    const addNews = async () => {
      try {
        const response = await axios.post('/api/magazine/v1/news', {
          title: newTitle.value,
          content: newContent.value,
        });
        newsList.value.push({
          id: response.data.news_id,
          title: newTitle.value,
          content: newContent.value,
          date: new Date().toISOString(), // Assuming the date is current date
        });
        newTitle.value = '';
        newContent.value = '';
      } catch (error) {
        console.error('Error adding news:', error);
      }
    };

    onMounted(fetchNews);

    return {
      newsList,
      newTitle,
      newContent,
      addNews,
    };
  },
});
</script>

<style scoped>
.container {
  max-width: 800px;
}
</style>
