<!-- NewsPage.vue -->
<template>
  <Nav />
  <div class="container mx-auto">
    <AdminNews />
    <NewsList />
  </div>
</template>

<script setup lang="ts">
import Nav from '../components/Nav.vue';
import AdminNews from '../components/magazine/AdminNews.vue';
import NewsList from '../components/magazine/NewsList.vue';

</script>

<style scoped>
/* Add any additional styles here */
</style>
