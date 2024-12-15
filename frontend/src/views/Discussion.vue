<template>
<div class="flex flex-col min-h-screen">
    <Nav />
    <SelectCrypto />
    <Comments />
    
</div>
</template>

<script setup lang="ts">
import Nav from '../components/Nav.vue';
import SelectCrypto from '../components/discussion/SelectCrypto.vue';
import Comments from '../components/discussion/Comments.vue';
</script>

<style>


</style>