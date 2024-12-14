<template>
  <ArticleList v-if="isLoggedIn"/>
  <div v-else>Oh no!</div>
</template>

<script lang="ts">
import ArticleList from '../components/blog/ArticleList.vue';
import { refreshService } from '../components/refreshJWT';
import { useRouter } from 'vue-router';
import { defineComponent, ref, onMounted } from 'vue';

export default defineComponent({
  setup() {
    const router = useRouter();
    const isLoggedIn = ref<boolean | null>(false);

    onMounted(async () => {
      isLoggedIn.value = await refreshService.refreshJWT(); 
      console.log("isLoggedIn.value: ", isLoggedIn.value);
      if (!isLoggedIn.value) {
        router.push('/login'); 
      }
    });

    return {
      isLoggedIn, // Expose the reactive state to the template
    };
  }
});


</script>

<style>

</style>