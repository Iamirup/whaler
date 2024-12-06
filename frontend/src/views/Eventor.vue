<template>
<div>
    <svg class="hidden">
    <symbol id="chevron-left" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor">
      <path fill-rule="evenodd" d="M6.22 4.22a.75.75 0 0 1 1.06 0l3.25 3.25a.75.75 0 0 1 0 1.06l-3.25 3.25a.75.75 0 0 1-1.06-1.06L8.94 8 6.22 5.28a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd" />
    </symbol>    
    <symbol id="sun" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" d="M12 3v2.25m6.364.386-1.591 1.591M21 12h-2.25m-.386 6.364-1.591-1.591M12 18.75V21m-4.773-4.227-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z" />
    </symbol> 
    <symbol id="moon" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" d="M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z" />
    </symbol>           
    
    </svg>

    <Nav />

    <Description />

    <Table />

    <Report />

    <Info />

    <Numbers />

    <Footer />
</div>
</template>

<script lang="ts">
import Nav from '../components/Nav.vue';
import Description from '../components/eventor/Description.vue';
import Table from '../components/eventor/Table.vue';
import Report from '../components/eventor/Report.vue';
import Info from '../components/eventor/Info.vue';
import Numbers from '../components/eventor/Numbers.vue';
import Footer from '../components/eventor/Footer.vue';
import { defineComponent } from 'vue'


export default defineComponent({
  components: {
    Nav,
    Description,
    Table,
    Report,
    Info,
    Numbers,
    Footer
  },
})

</script>

<style scoped>

</style>