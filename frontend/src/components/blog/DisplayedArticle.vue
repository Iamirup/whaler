<template>
  <div class="article-container" v-if="article">
    <div class="article-header">
      <h2 class="article-title">{{ article.title }}</h2>
      <div class="article-meta">
        <span class="article-author">By {{ article.author_username }}</span>
        <span class="article-date">{{ formatDate(article.date) }}</span>
      </div>
    </div>
    <div class="article-content">
      <p>{{ article.content }}</p>
    </div>
    <div class="article-footer">
      <button class="like-button" @click="likeArticle">
        <i class="fas fa-heart"></i> Like
      </button>
      <span class="article-likes">{{ article.likes }} Likes</span>
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
/* .article-container {
  background-color: #ffffff;
  border-radius: 15px;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
  max-width: 800px;
  margin: 40px auto;
  padding: 30px;
  transition: transform 0.3s, box-shadow 0.3s;
}

.article-container:hover {
  transform: translateY(-10px);
  box-shadow: 0 15px 30px rgba(0, 0, 0, 0.2);
}

.article-header {
  border-bottom: 2px solid #eaeaea;
  padding-bottom: 15px;
  margin-bottom: 25px;
}

.article-title {
  font-size: 2.5rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 10px;
  text-align: center;
}

.article-meta {
  display: flex;
  justify-content: space-between;
  font-size: 0.9rem;
  color: #777;
  margin-bottom: 20px;
}

.article-content {
  font-size: 1.2rem;
  line-height: 1.8;
  color: #555;
  margin-bottom: 30px;
  text-align: justify;
}

.article-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.like-button {
  background-color: #e0245e;
  color: #fff;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.3s;
  display: flex;
  align-items: center;
}

.like-button i {
  margin-right: 8px;
}

.like-button:hover {
  background-color: #c81e4d;
  transform: scale(1.05);
}

.article-likes {
  font-size: 1rem;
  color: #777;
} */
</style>