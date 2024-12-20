<!-- AdminNews.vue -->
<template>
  <div v-if="isAdmin" class="p-6 bg-gray-800 rounded-md shadow-lg">
    <h2 class="text-3xl font-bold mb-6">Add News</h2>
    <form @submit.prevent="addNews" class="mb-6">
      <div class="mb-4">
        <label for="title" class="block text-lg font-semibold mb-2">Title</label>
        <input v-model="newTitle" id="title" type="text" class="w-full p-2 rounded bg-gray-700 text-white">
      </div>
      <div class="mb-4">
        <label for="content" class="block text-lg font-semibold mb-2">Content</label>
        <textarea v-model="newContent" id="content" class="w-full p-2 rounded bg-gray-700 text-white"></textarea>
      </div>
      <button type="submit" class="px-4 py-2 bg-blue-600 rounded hover:bg-blue-700">Add News</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';
import { alertService } from '../alertor';
import { adminService } from '../admin_check';

export default defineComponent({
  setup() {
    const isAdmin = ref<boolean | null>(false);
    const newTitle = ref('');
    const newContent = ref('');

    const checkAdmin = async () => {
      isAdmin.value = await adminService.isAdmin();
    };

    const addNews = async () => {
      try {
        await axios.post('/api/magazine/v1/news', {
          title: newTitle.value,
          content: newContent.value,
        });
        newTitle.value = '';
        newContent.value = '';
      } catch (error: any) {
        alertService.showAlert(error.response.data.errors[0].message, "error");
        console.error('Error adding news:', error);
      }
    };

    onMounted(checkAdmin);

    return {
      isAdmin,
      newTitle,
      newContent,
      addNews,
    };
  },
});
</script>

<style scoped>

</style>
