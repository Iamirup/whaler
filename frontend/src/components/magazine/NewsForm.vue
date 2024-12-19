<template>
    <div class="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-75">
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-xl font-bold">Add News</h2>
          <button @click="$emit('close')" class="text-gray-600 hover:text-gray-800">&times;</button>
        </div>
        <form @submit.prevent="submitNews" class="space-y-4">
          <div>
            <label class="block text-gray-700 text-sm font-bold mb-2" for="title">Title</label>
            <input v-model="title" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="title" type="text" placeholder="Title">
          </div>
          <div>
            <label class="block text-gray-700 text-sm font-bold mb-2" for="content">Content</label>
            <textarea v-model="content" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="content" placeholder="Content"></textarea>
          </div>
          <div class="flex justify-end space-x-2">
            <button type="button" @click="$emit('close')" class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Cancel</button>
            <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
              Add News
            </button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';

export default defineComponent({
  name: 'NewsForm',
  emits: ['close', 'newsCreated'],
  setup(props, { emit }) {
    const title = ref('');
    const content = ref('');

    const submitNews = async () => {
      await axios.post('/news', {
        title: title.value,
        content: content.value
      });
      emit('newsCreated');
      emit('close');
      title.value = '';
      content.value = '';
    };

    return { title, content, submitNews };
  }
});

  </script>
  