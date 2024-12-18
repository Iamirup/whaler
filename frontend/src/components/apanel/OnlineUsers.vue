<template>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-6">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Online Users</h2>
    <button @click="fetchOnlineUsers" class="bg-blue-600 text-white px-4 py-3 rounded mb-4 hover:bg-blue-700">Refresh</button>
    <ul>
      <li v-for="user in onlineUsers" :key="user.id" class="border-b py-2">{{ user.username }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';
import { alertService } from '../alertor';
import { useRouter } from 'vue-router';
import { refreshService } from '../refreshJWT';

interface User {
  id: string;
  email: string;
  username: string;
  created_at: string;
}

export default defineComponent({
  setup() {

    const router = useRouter()
    const onlineUsers = ref<User[]>([]);

    const fetchOnlineUsers = async () => {
      await axios.get(`/api/auth/v1/onlines`)
        .then(response => {
          onlineUsers.value = response.data.online_users;
        })
        .catch(async error => {
          if (error.response.data.need_refresh){
            const isRefreshed = await refreshService.refreshJWT(); 
            if (!isRefreshed) { router.push('/login'); return; }
            await fetchOnlineUsers();
          } else {
            alertService.showAlert(error.response.data.errors[0].message, "error");
            console.error(error);
          }
        });
    }

    onMounted(async () => {
      await fetchOnlineUsers();
    });

    return {
      fetchOnlineUsers,
      onlineUsers
    }
  }
});
</script>
