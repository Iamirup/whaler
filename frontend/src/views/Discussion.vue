<template>
<div class="flex flex-col min-h-screen">
    <Nav class="bg-gray-100"/>
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