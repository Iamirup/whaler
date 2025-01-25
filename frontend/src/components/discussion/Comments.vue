<template>
  <div class="max-w-md mx-auto mt-10">
    <div class="bg-white shadow-md rounded-lg p-4">
      <h2 class="text-2xl font-semibold mb-4 text-blue-700">Comments: {{ currency }}</h2>
      <div class="max-h-64 overflow-y-auto bg-gray-100 rounded-lg p-2">
        <ul>
          <li v-for="comment in comments" :key="comment.comment_id" class="mb-3 p-2 bg-white rounded-lg shadow-sm">
            <div class="flex items-center mb-2">
              <div class="font-bold text-sm mr-2 text-green-600">{{ comment.username }}</div>
              <div class="text-gray-500 text-xs">{{ comment.date }}</div>
            </div>
            <p class="text-gray-800">{{ comment.text }}</p>
          </li>
        </ul>
      </div>
      <div class="mt-4">
        <textarea v-model="newCommentText" placeholder="Write a comment..." class="w-full p-2 border rounded"></textarea>
        <button @click="addComment" class="bg-blue-500 text-white px-4 py-2 rounded mt-2">Add Comment</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch, computed } from 'vue';
import { useCurrencyStore } from '../../stores/currencyStore';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { refreshService } from '../refreshJWT';
import { alertService } from '../alertor';

export default defineComponent({
  name: 'CommentsBox',
  setup() {
    const store = useCurrencyStore();
    const currency = computed(() => store.currency);
    const ownUsername = ref('');
    const commentId = ref(0);

    const router = useRouter();

    const comments = ref<Array<{ 
      comment_id: number; 
      currency: string; 
      username: string; 
      text: string; 
      date: string;
    }>>([]);
    
    const newCommentText = ref('');

    const formatToJalali = (datetime: Date): string => {
      const date = new Date(datetime);

      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0');
      const day = String(date.getDate()).padStart(2, '0');
      const hours = String(date.getHours()).padStart(2, '0');
      const minutes = String(date.getMinutes()).padStart(2, '0');
      const seconds = String(date.getSeconds()).padStart(2, '0');
      const formattedDateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
      return formattedDateTime;
    };

    const fetchComments = async () => {
      try {
        const response = await axios.get(`/api/discussion/v1/comments/${currency.value}`);
        comments.value = response.data.comments.map((comment: { date: string }) => {
          const date = new Date(comment.date);
          return {
            ...comment,
            date: formatToJalali(date),
          };
        });
        ownUsername.value = response.data.own_username;
      } catch (error: any) {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          await fetchComments();
        } else {
          alertService.showAlert(error.response.data.errors[0].message, "error");
          console.error('Error fetching comments:', error);
        }
      }
    };

    const addComment = async () => {
      if (newCommentText.value.trim() === '') return;
      try {
        const commentData = { currency: currency.value, text: newCommentText.value};
        const response = await axios.post(`/api/discussion/v1/comment`, commentData);
        commentId.value = response.data.comment_id;
      } catch (error: any) {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          await addComment();
        } else {
          alertService.showAlert(error.response.data.errors[0].message, "error");
          console.error('Error fetching comments:', error);
          return;
        }
      }
      comments.value.push({
        comment_id: commentId.value,
        currency: currency.value,
        username: ownUsername.value,
        date: formatToJalali(new Date()),
        text: newCommentText.value,
      });
      newCommentText.value = '';
    };

    watch(currency, async () => {
      await fetchComments();
    });

    return {
      currency,
      comments,
      newCommentText,
      addComment,
    };
  },
});
</script>

<style scoped>
textarea {
  resize: none;
}
.max-h-64 {
  max-height: 16rem;
}
.bg-gray-100 {
  background-color: #f7fafc;
}


</style>