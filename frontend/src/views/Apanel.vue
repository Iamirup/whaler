// src/views/AdminPanel.vue
<template>
  <div class="container mx-auto p-8" v-if="isLoggedIn">
    <h1 class="text-4xl font-bold text-gray-800 mb-8">Admin Panel</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <DeleteUser />
      <OnlineUsers />
      <AddAdmin />
      <DeleteAdmin />
      <TopAuthors />
    </div>
  </div>
</template>

<script setup lang="ts">
import DeleteUser from '../components/apanel/DeleteUser.vue';
import OnlineUsers from '../components/apanel/OnlineUsers.vue';
import AddAdmin from '../components/apanel/AddAdmin.vue';
import DeleteAdmin from '../components/apanel/DeleteAdmin.vue';
import TopAuthors from '../components/apanel/TopAuthors.vue';
import { refreshService } from '../components/refreshJWT';
import { useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';

const router = useRouter();
const isLoggedIn = ref<boolean | null>(false);

onMounted(async () => {
  isLoggedIn.value = await refreshService.refreshJWT(); 
  if (!isLoggedIn.value) {
    router.push('/login'); 
  }
});

</script>

<style scoped>
@import 'tailwindcss/tailwind.css';
</style>
