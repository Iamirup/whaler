<template>
  <div class="p-6 bg-white rounded-lg shadow-lg max-w-2xl mx-auto my-8" v-if="article">
    <h2 class="text-3xl font-bold text-gray-800 mb-4 border-b pb-2">{{ article.title }}</h2>
    <p class="text-gray-700 text-lg leading-relaxed">{{ article.content }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import axios from 'axios';
import { refreshJWT } from '../refreshJWT';

export default defineComponent({
  setup() {
    const route = useRoute();
    const article = ref<{ title: string; content: string } | null>(null);

    const fetchArticle = async (id: string) => {
      return await axios.get(`/api/blog/v1/article/${id}`)
        .then(response => {
          article.value = response.data;
        })
        .catch(error => {
          console.error('Failed to fetch article', error);
          refreshJWT();
        });
    };

    onMounted(() => {
      const { id } = route.params;
      if (id) fetchArticle(id as string);
    });

    return { article };
  },
});
</script>

<style scoped>
</style>