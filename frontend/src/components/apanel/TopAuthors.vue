// src/components/TopAuthors.vue
<template>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-6">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Top Authors</h2>
    <button @click="fetchTopAuthors" class="bg-blue-600 text-white px-4 py-3 rounded mb-4 hover:bg-blue-700">Refresh</button>
    <ul>
      <li v-for="author in topAuthors" :key="author.id" class="border-b py-2">{{ author.name }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

export interface User {
  id: string;
  name: string;
}
export interface Author {
  id: string;
  name: string;
}

export default defineComponent({
  data() {
    return {
      topAuthors: [] as Author[],
    };
  },
  methods: {
    fetchTopAuthors() {
      axios.get<Author[]>(`/api/blog/v1/top-authors`)
        .then(response => {
          this.topAuthors = response.data;
        })
        .catch(() => {
          alert('Failed to fetch top authors');
        });
    },
  },
});
</script>
