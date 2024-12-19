<!-- NewsPage.vue -->
<template>
  <Nav />
  <div class="container mx-auto">
    <AdminNews />
    <NewsList />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import AdminNews from '../components/magazine/AdminNews.vue';
import NewsList from '../components/magazine/NewsList.vue';

export default defineComponent({
  components: {
    AdminNews,
    NewsList,
  },
});
</script>

<style scoped>
/* Add any additional styles here */
</style>
