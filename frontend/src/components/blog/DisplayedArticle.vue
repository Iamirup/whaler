<template>
  <div class="bg-gradient-to-br from-slate-400 to-slate-100 min-h-screen flex items-center justify-center">
    <div class="bg-white rounded-lg shadow-lg max-w-3xl w-full p-8 transition-transform transform hover:scale-105" v-if="article">
      <div class="border-b-2 border-gray-200 pb-4 mb-6">
        <h2 class="text-4xl font-bold text-gray-800 mb-2 text-center">{{ article.title }}</h2>
        <div class="flex justify-between text-gray-600 text-sm">
          <span>By {{ article.author_username }}</span>
          <span>{{ formatDate(article.date) }}</span>
        </div>
      </div>
      <div class="text-gray-700 text-lg leading-relaxed mb-6">
        <p>{{ article.content }}</p>
      </div>
      <div class="flex justify-between items-center">
        <button class="bg-red-500 text-white px-4 py-2 rounded-full flex items-center hover:bg-red-600 transform transition-transform hover:scale-110" @click="likeArticle">
          <i class="fas fa-heart mr-2"></i> Like
        </button>
        <span class="text-gray-600 text-lg">{{ article.likes }} Likes</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import axios from 'axios';
import { refreshService } from '../refreshJWT';
import { alertService } from '../alertor';

export default defineComponent({
  setup() {
    const route = useRoute();
    const router = useRouter();
    const article = ref<{
      title: string;
      content: string;
      author_username: string;
      likes: number;
      date: string;
    } | null>(null);

    const fetchArticle = async (id: string) => {
      return await axios.get(`/api/blog/v1/article/${id}`)
        .then(response => {
            article.value = response.data.article;
        })
        .catch(async error => {
          if (error.response.data.need_refresh){
              const isRefreshed = await refreshService.refreshJWT(); 
              if (!isRefreshed) { router.push('/login'); return; }
              await fetchArticle(id);
          } else {
              alertService.showAlert(error.response.data.errors[0].message, "error");
              console.error('Failed to fetch article', error);
          }
        });
    };

    const formatDate = (dateString: string) => {
      const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric' };
      return new Date(dateString).toLocaleDateString(undefined, options);
    };

    const likeArticle = () => {
      // Implement like functionality here
      article.value!.likes += 1;
    };

    onMounted(() => {
      const articleId = route.params.articleId;
      if (articleId) fetchArticle(articleId as string);
    });

    return { article, formatDate, likeArticle };
  },
});
</script>

<style scoped>
/* No additional styles needed */
</style>