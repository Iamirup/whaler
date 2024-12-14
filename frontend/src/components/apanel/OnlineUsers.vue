// src/components/OnlineUsers.vue
<template>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-6">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Online Users</h2>
    <button @click="fetchOnlineUsers" class="bg-blue-600 text-white px-4 py-3 rounded mb-4 hover:bg-blue-700">Refresh</button>
    <ul>
      <li v-for="user in onlineUsers" :key="user.id" class="border-b py-2">{{ user.name }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

export interface User {
  id: string;
  name: string;
}
export interface Author {
  id: string;
  name: string;
}

export default defineComponent({
  data() {
    return {
      onlineUsers: [] as User[],
    };
  },
  methods: {
    fetchOnlineUsers() {
      axios.get<User[]>(`/api/auth/v1/onlines`)
        .then(response => {
          this.onlineUsers = response.data;
        })
        .catch(() => {
          alert('Failed to fetch online users');
        });
    },
  },
});
</script>
