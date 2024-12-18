<template>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-6">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Top Authors</h2>
    <button @click="fetchTopAuthors" class="bg-blue-600 text-white px-4 py-3 rounded mb-4 hover:bg-blue-700">Refresh</button>
    <ul>
      <li v-for="author in topAuthors" :key="author.author_id" class="border-b py-2">{{ author.author_username }} with {{ author.likes }} likes❤️</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';

interface TopAuthor {
  author_id: string;
  author_username: string;
  likes: number;
}

export default defineComponent({
  setup() {

    const topAuthors = ref<TopAuthor[]>([]);

    const fetchTopAuthors = async () => {
      await axios.get(`/api/blog/v1/top-authors`)
        .then(response => {
          topAuthors.value = response.data.authors;
        })
        .catch(error => {
          console.error(error);
        });
    }

    onMounted(async () => {
      await fetchTopAuthors();
    });

    return {
      fetchTopAuthors,
      topAuthors
    }
  }
});
</script>
