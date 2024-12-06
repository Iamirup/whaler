<template>
    <div class="max-w-md mx-auto mt-10">
      <div class="bg-white shadow-md rounded-lg p-4">
        <h2 class="text-2xl font-semibold mb-4 text-blue-700">Comments: {{ currency }}</h2>
        <div class="max-h-64 overflow-y-auto bg-gray-100 rounded-lg p-2">
          <ul>
            <li v-for="comment in comments" :key="comment.id" class="mb-3 p-2 bg-white rounded-lg shadow-sm">
              <div class="flex items-center mb-2">
                <div class="font-bold text-sm mr-2 text-green-600">{{ comment.userName }}</div>
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
  import { defineComponent, ref, onMounted, computed } from 'vue';
  import { useCurrencyStore } from '../../stores/currencyStore';
  import axios from 'axios';
  
  export default defineComponent({
    name: 'CommentsBox',
    setup() {
      const store = useCurrencyStore();
      const currency = computed(() => store.currency);

      const comments = ref([]);
      const newCommentText = ref('');
  
      const fetchComments = async () => {
        try {
          const response = await axios.get(`/api/comments/`);
          comments.value = response.data;
        } catch (error) {
          console.error('Error fetching comments:', error);
        }
      };
  
      const addComment = () => {
        if (newCommentText.value.trim() === '') return;
        comments.value.push({
          id: comments.value.length + 1,
          userName: 'Anonymous',
          date: new Date().toISOString().split('T')[0],
          text: newCommentText.value,
        });
        newCommentText.value = '';
      };
  
      onMounted(() => {
        fetchComments();
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