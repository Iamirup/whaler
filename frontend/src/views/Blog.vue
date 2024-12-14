<template>
  <ArticleList />
</template>

<script setup lang="ts">
import ArticleList from '../components/blog/ArticleList.vue';
import { refreshJWT } from '../components/refreshJWT';
refreshJWT();


</script>

<style>

</style>