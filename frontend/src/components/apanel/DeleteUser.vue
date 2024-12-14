// src/components/DeleteUser.vue
<template>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-6">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Delete User</h2>
    <div class="flex items-center mb-4">
      <input v-model="userId" placeholder="Enter User ID" class="w-full p-3 border border-gray-300 rounded-l-md focus:outline-none" />
      <button @click="deleteUser" class="bg-red-600 text-white px-4 py-3 rounded-r-md hover:bg-red-700">Delete</button>
    </div>
    <p v-if="message" :class="messageClass">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

export default defineComponent({
  data() {
    return {
      userId: '',
      message: '',
      messageClass: '',
    };
  },
  methods: {
    deleteUser() {
      axios.delete(`/api/auth/v1/user`, { data: { id: this.userId } })
        .then(() => {
          this.message = 'User deleted successfully';
          this.messageClass = 'text-green-600';
        })
        .catch(() => {
          this.message = 'Failed to delete user';
          this.messageClass = 'text-red-600';
        });
    },
  },
});
</script>
