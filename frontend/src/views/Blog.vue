<template>
  <div v-if="isLoggedIn">
    <ArticleList />
  </div>
</template>

<script lang="ts">
import ArticleList from '../components/blog/ArticleList.vue';
import { refreshJWT } from '../components/refreshJWT';
import { useRouter } from 'vue-router';
import { alertService } from '../components/alertor';
import { defineComponent, ref, onMounted } from 'vue';

export default defineComponent({
  setup() {
    const router = useRouter();
    const isLoggedIn = ref(false); // Reactive state

    onMounted(async () => {
      const isThereUser = await refreshJWT(); 
      if (!isThereUser) { 
        router.push('/login'); 
      } else {
        isLoggedIn.value = true; // Use .value to update the reactive state
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