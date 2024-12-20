<script lang="ts">

import { defineComponent } from 'vue'


export default defineComponent({

  
})

</script>


<template class="bg-gradient-to-br from-slate-300 to-slate-100">
  <div class="font">
    <router-view></router-view>
  </div>
</template>

<style scoped>
@font-face {
  font-family: "CustomFont";
  src: url("../src/assets/CustomFont.ttf");
  src: url("../src/assets/CustomFont.ttf") format("ttf"),
  url("../src/assets/CustomFont.ttf") format("opentype"),
}

.font {
  font-family: 'CustomFont', Arial, sans-serif;
  font-weight:normal;
}
</style>
