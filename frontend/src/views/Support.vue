<template>
  <Nav />
  <div  v-if="isLoggedIn" class="mt-6">
    <AdminTicket v-if="isAdmin"/>
    <UserTicket v-else/>
  </div>
</template>

<script setup lang="ts">
import Nav from '../components/Nav.vue';
import AdminTicket from '@/components/support/AdminTicket.vue';
import UserTicket from '../components/support/UserTicket.vue';
import { refreshService } from '../components/refreshJWT';
import { adminService } from '../components/admin_check';
import { useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';

const router = useRouter();
const isLoggedIn = ref<boolean | null>(false);
const isAdmin = ref<boolean | null>(false);

onMounted(async () => {
  isLoggedIn.value = await refreshService.refreshJWT(); 
  if (!isLoggedIn.value) {
    router.push('/login'); 
  }
  isAdmin.value = await adminService.isAdmin(); 
});




</script>

<style>

</style>