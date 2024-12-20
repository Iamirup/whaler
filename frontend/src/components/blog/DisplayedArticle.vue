<template>
  <div class="rounded-lg shadow-lg max-w-5xl mx-auto break-words p-6 bg-white" v-if="article">
    <h2 class="text-4xl font-bold text-gray-800 mb-4 border-b pb-2">{{ article.title }}</h2>
    <div class="flex items-center justify-between mb-4">
      <span class="text-gray-600">By {{ article.author_username }}</span>
      <span class="text-gray-600">{{ formatDate(article.date) }}</span>
    </div>
    <p class="text-gray-700 text-lg leading-relaxed mb-4">{{ article.content }}</p>
    <div class="flex items-center justify-between">
      <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-700">Like</button>
      <span class="text-gray-600">{{ article.likes }} Likes</span>
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

    onMounted(() => {
      const articleId = route.params.articleId;
      if (articleId) fetchArticle(articleId as string);
    });

    return { article, formatDate };
  },
});
</script>

<style scoped>
.bg-white {
  background-color: #ffffff;
}
</style>