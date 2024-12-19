<template>
  <div class="container mx-auto p-4">
    <h1 class="text-4xl font-bold mb-6 text-center text-blue-600">News Management</h1>
    <button @click="showForm = true" class="mb-4 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
      Add News
    </button>
    <NewsList :newsList="newsList" />
    <NewsForm v-if="showForm" @close="showForm = false" @newsCreated="fetchNews" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';
import NewsList from '@/components/magazine/NewsList.vue';
import NewsForm from '@/components/magazine/NewsForm.vue';

export default defineComponent({
  name: 'Magazine',
  components: {
    NewsForm,
    NewsList
  },
  async setup() {
    const showForm = ref(false);
    const newsList = ref([]);
    
    const fetchNews = async () => {
      const response = await axios.get('/news');
      newsList.value = response.data;
    };

    await fetchNews();

    return { showForm, newsList, fetchNews };
  }
});
</script>

<style scoped>
body {
  background-color: #f9fafb;
}
</style>
