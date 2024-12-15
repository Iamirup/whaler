<template>
    <Nav class="bg-gray-100"/>
    <ArticleList v-if="isLoggedIn"/>
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