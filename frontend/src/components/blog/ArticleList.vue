<template>
  <Nav />
  <div class="pt-32 p-6 min-h-screen">
    <div v-for="article in articles" :key="article.article_id" @click="selectArticle(article)" class="p-4 bg-gradient-to-br from-orange-200 to-slate-50 rounded-lg shadow-lg mb-4 cursor-pointer hover:bg-gray-50 transition break-words duration-150">
      <h3 class="text-xl font-bold text-gray-800">{{ article.title }}</h3>
      <p class="text-gray-600 text-sm mt-1">{{ article.content }}</p>
    </div>
    <div class="article-list text-center p-5">
      <h1 @click="myArticles" class="text-4xl font-bold text-gray-800 mb-5 relative inline-block pb-2">
      My Articles
      <span class="block w-1/2 h-1 bg-green-500 absolute bottom-0 left-1/4 rounded"></span>
      </h1>
      <button @click="loadMore" v-if="cursor" class="bg-green-500 text-white py-3 px-8 rounded-full shadow-lg transition-transform transform hover:bg-green-600 hover:-translate-y-1 active:bg-green-700 active:translate-y-0">
        Load More
      </button>
    </div>
  </div>
</template>

<script lang="ts">

import { defineComponent, ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';
import { refreshService } from '../refreshJWT';
import { alertService } from '../alertor';

export default defineComponent({
  setup() {
    const articles = ref<Array<{ 
      article_id: string; 
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
      const url = cursorParam ? `/api/blog/v1/all-articles?limit=20&cursor=${cursorParam}` : '/api/blog/v1/all-articles?limit=20';
      await axios.get(url)
      .then(response => {
        articles.value.push(...response.data.articles);
        cursor.value = response.data.new_cursor;
      })
      .catch(async error => {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          await fetchArticles();
        } else {
          alertService.showAlert(error.response.data.errors[0].message, "error");
          console.error(error);
        }
      });
    };

    onMounted(async () => fetchArticles());

    const selectArticle = (article: { 
      article_id: string; 
      title: string; 
      content: string; 
      author_id: string; 
      author_username: string;
      likes: number;
      date: string;
    }) => {
      router.push({ path: `/article/${article.article_id}` });
    };

    const loadMore = () => {
      if (cursor.value) {
        fetchArticles(cursor.value);
      }
    };

    const myArticles = () => {
      router.push({ path: `/blog/manage` });
    }

    return {
      articles,
      cursor,
      selectArticle,
      loadMore,
      myArticles,
    };
  },
});
</script>

<style scoped>
</style>

