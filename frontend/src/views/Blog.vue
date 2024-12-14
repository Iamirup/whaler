<template>
  <div v-if="isLoggedIn">
    <ArticleList />
  </div>
</template>

<script lang="ts">
import ArticleList from '../components/blog/ArticleList.vue';
import { refreshJWT } from '../components/refreshJWT';
import { useRouter } from 'vue-router';
import { defineComponent, ref, onMounted } from 'vue';

export default defineComponent({
  setup() {
    const router = useRouter();
    const isLoggedIn = ref(false); // Reactive state

    onMounted(async () => {
      const isThereUser = await refreshJWT(); 
      console.log("isThereUser: ", isThereUser);
      if (isThereUser == false) {
        router.push('/login'); 
      } else {
        isLoggedIn.value = true;
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