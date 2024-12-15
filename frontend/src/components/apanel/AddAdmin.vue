// src/components/AddAdmin.vue
<template>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-6">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Add Admin</h2>
    <div class="flex items-center mb-4">
      <input v-model="adminId" placeholder="Enter Admin ID" class="w-full p-3 border border-gray-300 rounded-l-md focus:outline-none" />
      <button @click="addAdmin" class="bg-green-600 text-white px-4 py-3 rounded-r-md hover:bg-green-700">Add</button>
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
      adminId: '',
      message: '',
      messageClass: '',
    };
  },
  methods: {
    addAdmin() {
      axios.post(`/api/auth/v1/admin`, { id: this.adminId })
        .then(() => {
          this.message = 'Admin added successfully';
          this.messageClass = 'text-green-600';
        })
        .catch(() => {
          this.message = 'Failed to add admin';
          this.messageClass = 'text-red-600';
        });
    },
  },
});
</script>
