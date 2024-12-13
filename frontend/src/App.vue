<script lang="ts">
import Nav from './components/Nav.vue';
import Description from './components/eventor/Description.vue';
import Table from './components/eventor/Table.vue';
import Report from './components/eventor/Report.vue';
import Info from './components/eventor/Info.vue';
import Numbers from './components/eventor/Numbers.vue';
import Footer from './components/eventor/Footer.vue';
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
  }



  
})

</script>


<template >
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

url("../src/assets/CustomFont.svg") format("svg");

}


.font {

font-family: 'CustomFont', Arial, sans-serif;

font-weight:normal;
}
</style>
