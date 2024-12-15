<template>
  <div class="flex flex-col min-h-screen">
    <Nav class="bg-gray-800 text-white p-4" />
    <ArticleList v-if="isLoggedIn" class="flex-1 p-4"/>
  </div>

</template>

<script setup lang="ts">
import Nav from '../components/Nav.vue';
import ArticleList from '../components/blog/ArticleList.vue';
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

<style>

</style>