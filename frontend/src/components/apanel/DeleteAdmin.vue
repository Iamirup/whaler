// src/components/DeleteAdmin.vue
<template>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-6">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Delete Admin</h2>
    <div class="flex items-center mb-4">
      <input v-model="adminId" placeholder="Enter Admin ID" class="w-full p-3 border border-gray-300 rounded-l-md focus:outline-none" />
      <button @click="deleteAdmin" class="bg-red-600 text-white px-4 py-3 rounded-r-md hover:bg-red-700">Delete</button>
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
    deleteAdmin() {
      axios.delete(`/api/apanel/v1/admin`, { data: { id: this.adminId } })
        .then(() => {
          this.message = 'Admin deleted successfully';
          this.messageClass = 'text-green-600';
        })
        .catch(() => {
          this.message = 'Failed to delete admin';
          this.messageClass = 'text-red-600';
        });
    },
  },
});
</script>
