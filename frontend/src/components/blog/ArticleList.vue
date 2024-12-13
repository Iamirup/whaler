<template>
  <div class="p-6 bg-gray-100 min-h-screen">
    <div v-for="article in articles" :key="article.article_id" @click="selectArticle(article)" class="p-4 bg-white rounded-lg shadow-lg mb-4 cursor-pointer hover:bg-gray-50 transition duration-150">
      <h3 class="text-xl font-bold text-gray-800">{{ article.title }}</h3>
      <p class="text-gray-600 text-sm mt-1">{{ article.content }}</p>
    </div>
    <button @click="loadMore" v-if="cursor" class="px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700">Load More</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

export default defineComponent({
  setup() {
    const articles = ref<Array<{ 
      article_id: string; 
      url_path: string; 
      title: string; 
      content: string; 
      author_id: string; 
      author_username: string;
      likes: number;
      date: string;
    }>>([]);
    const cursor = ref<string | null>(null);
    const router = useRouter();

    const fetchArticles = async (cursorParam?: string) => {
      const url = cursorParam ? `/api/blog/v1/all-articles?limit=20&cursor=${cursorParam}` : '/articles?limit=20';
      await axios.get(url)
      .then(response => {
        articles.value.push(...response.data.articles);
        cursor.value = response.data.new_cursor;
      })
      .catch(error => {
        console.log(error.response);
        console.error('Failed to fetch articles', error);
      });
    };

    onMounted(() => fetchArticles());

    const selectArticle = (article: { 
      article_id: string; 
      url_path: string; 
      title: string; 
      content: string; 
      author_id: string; 
      author_username: string;
      likes: number;
      date: string;
    }) => {
      router.push({ path: `/api/blog/v1/article/${article.url_path}` });
    };

    const loadMore = () => {
      if (cursor.value) {
        fetchArticles(cursor.value);
      }
    };

    return {
      articles,
      cursor,
      selectArticle,
      loadMore,
    };
  },
});
</script>

<style scoped>
</style>

