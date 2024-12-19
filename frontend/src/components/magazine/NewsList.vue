<template>
    <div>
      <h2 class="text-2xl font-bold mb-4">News List</h2>
      <ul>
        <li v-for="news in newsList" :key="news.id" class="mb-4 p-4 bg-white shadow-md rounded-lg">
          <h3 class="text-xl font-semibold">{{ news.title }}</h3>
          <p class="text-gray-600">{{ news.content }}</p>
          <span class="text-gray-500 text-sm">{{ new Date(news.date).toLocaleDateString() }}</span>
        </li>
      </ul>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, type PropType } from 'vue';
  
  interface News {
    id: string;
    title: string;
    content: string;
    date: string;
  }
  
  export default defineComponent({
    name: 'NewsList',
    props: {
      newsList: {
        type: Array as PropType<News[]>,
        required: true
      }
    }
  });
  </script>
  