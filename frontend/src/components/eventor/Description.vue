<template>
 <main class="pt-32 mx-3 xs:mx-5 sm:mx-14">
  <section class="text-end">
    <div class="">
      <div class="grid-cols-2 xl:flex xl:justify-between">
        <div class="p-5 mb-5 sm:pt-12 xl:mb-0 bg-blue-200 text-black rounded-2xl xl:w-[800px] w-full h-[320px] sm:h-[300px] xl:h-[335px] ">
          <span class="text-3xl sm:text-4xl">Lorem ipsum dolor kon waed<br>, sit amet <span class="font-bold">whaler </span><br> adipisicing elit.Saepe.</span>
          <span class="flex pb-4 sm:pb-0 pt-4">Lorem ipsum dolor sit amet consectetur adipisicing elit. Sed cumque earum molestiae facilis, autem pariatur odio aperiam itaque dolorum cum tenetur illo neque explicabo, quasi quo exercitationem architecto harum. Veniam?</span>
          <span @click="nextNumber" class="flex justify-center float-left mt-3 pt-0.5 w-16 h-8 rounded-xl bg-sky-200 hover:bg-sky-300 transition-all">
            <p class="mx-1.5 mt-0.5 no-select">
              {{ number }}
            </p>

            <svg class="mx-0.5 w-6 h-6" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="Capa_1" x="0px" y="0px" viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve">
        <polygon style="fill:#6C3BEB;" points="107.5,196.5 77.5,196.5 47.5,350.5 77.5,504.5 107.5,504.5 167.5,350.5 "/>
        <rect x="7.5" y="196.5" style="fill:#576CEB;"/>
        <polygon style="fill:#FFC46C;" points="290.596,196.5 107.5,466 107.5,504.5 386,504.5 386,196.5 "/>
        <path style="fill:#FF903D;" d="M309,7.5V46L107.596,436h-0.085l-0.011,30h0.096c101.068,0,183-81.932,183-183v-86.5  C327.385,154.912,347.5,102.006,347.5,46C347.5,24.737,330.263,7.5,309,7.5z"/>
        <path style="fill:#FFB03B;" d="M260.596,283v-97.865l7.53-8.512C299.965,140.63,317.5,94.24,317.5,46c0-21.263-3.806-38.5-8.5-38.5  c-21.263,0-38.5,17.237-38.5,38.5c0,83.119-67.381,150.5-150.5,150.5h-12.404L107.511,436h0.086  C191.96,436,260.596,367.364,260.596,283z"/>
        <path style="fill:#FF903D;" d="M406,427.5v-20l-20,20c-1.696,0-3.362,0.121-5,0.334v76.332c1.638,0.213,3.304,0.334,5,0.334h20  c21.263,0,38.5-17.237,38.5-38.5S427.263,427.5,406,427.5z"/>
        <path style="fill:#FFB03B;" d="M414.5,466c0-19.567-14.601-35.713-33.5-38.166c-18.899,2.453-33.5,18.599-33.5,38.166  s14.601,35.713,33.5,38.166C399.899,501.713,414.5,485.567,414.5,466z"/>
        <path style="fill:#FF903D;" d="M426,350.5v-20l-30,20v77h30c21.263,0,38.5-17.237,38.5-38.5S447.263,350.5,426,350.5z"/>
        <path style="fill:#FFB03B;" d="M434.5,389c0-21.263-17.237-38.5-38.5-38.5v-20l-10,20c-21.263,0-38.5,17.237-38.5,38.5  s17.237,38.5,38.5,38.5h10C417.263,427.5,434.5,410.263,434.5,389z"/>
        <g>
          <path style="fill:#FF903D;" d="M466,196.5h-30v77h30c21.263,0,38.5-17.237,38.5-38.5S487.263,196.5,466,196.5z"/>
          <path style="fill:#FF903D;" d="M446,273.5v-20l-30,20v77h30c21.263,0,38.5-17.237,38.5-38.5S467.263,273.5,446,273.5z"/>
        </g>
        <g>
          <path style="fill:#FFB03B;" d="M474.5,235c0-21.263-17.237-38.5-38.5-38.5h-50c-21.263,0-38.5,17.237-38.5,38.5   s17.237,38.5,38.5,38.5h50C457.263,273.5,474.5,256.263,474.5,235z"/>
          <path style="fill:#FFB03B;" d="M454.5,312c0-21.263-17.237-38.5-38.5-38.5v-20l-30,20c-21.263,0-38.5,17.237-38.5,38.5   s17.237,38.5,38.5,38.5h30C437.263,350.5,454.5,333.263,454.5,312z"/>
        </g>
        <circle style="fill:#FFFFFF;" cx="57.5" cy="454.5" r="20"/>
        <path d="M57.5,427C42.336,427,30,439.337,30,454.5S42.336,482,57.5,482S85,469.663,85,454.5S72.664,427,57.5,427z M57.5,467  c-6.893,0-12.5-5.607-12.5-12.5S50.607,442,57.5,442S70,447.607,70,454.5S64.393,467,57.5,467z"/>
        <path d="M512,235c0-25.364-20.636-46-46-46H306.481C337.868,148.161,355,97.983,355,46c0-25.364-20.636-46-46-46s-46,20.636-46,46  c0,78.851-64.149,143-143,143H7.5c-4.142,0-7.5,3.357-7.5,7.5v308c0,4.143,3.358,7.5,7.5,7.5H406c25.364,0,46-20.636,46-46  c0-12.775-5.239-24.347-13.676-32.69C457.721,427.907,472,410.095,472,389c0-12.775-5.239-24.347-13.676-32.69  C477.721,350.907,492,333.095,492,312c0-12.775-5.239-24.347-13.676-32.69C497.721,273.907,512,256.095,512,235z M278,46  c0-17.094,13.907-31,31-31s31,13.906,31,31c0,53.737-19.54,105.421-55.021,145.53c-0.031,0.035-0.057,0.073-0.087,0.108  c-0.07,0.082-0.137,0.167-0.203,0.252c-0.076,0.097-0.15,0.195-0.221,0.295c-0.067,0.095-0.13,0.192-0.193,0.29  c-0.061,0.096-0.12,0.193-0.177,0.291c-0.062,0.107-0.121,0.216-0.177,0.327c-0.048,0.095-0.094,0.19-0.138,0.287  c-0.054,0.117-0.106,0.235-0.154,0.356c-0.038,0.097-0.073,0.194-0.107,0.292c-0.043,0.122-0.085,0.245-0.121,0.37  c-0.031,0.106-0.056,0.214-0.083,0.322c-0.029,0.119-0.06,0.238-0.083,0.359c-0.025,0.131-0.042,0.265-0.061,0.398  c-0.014,0.101-0.032,0.2-0.042,0.302c-0.023,0.239-0.036,0.479-0.036,0.721v47c0,4.143,3.358,7.5,7.5,7.5s7.5-3.357,7.5-7.5V204  h53.947C344.567,212.183,340,223.068,340,235c0,16.095,8.316,30.278,20.871,38.5C348.316,281.722,340,295.905,340,312  s8.316,30.278,20.871,38.5C348.316,358.722,340,372.905,340,389s8.316,30.278,20.871,38.5C348.316,435.722,340,449.905,340,466  c0,11.933,4.569,22.817,12.047,31H115v-23.657C216.623,469.438,298.096,385.562,298.096,283v-9.5c0-4.143-3.358-7.5-7.5-7.5  s-7.5,3.357-7.5,7.5v9.5c0,94.29-74.747,171.438-168.096,175.333V204h5C207.122,204,278,133.121,278,46z M15,204h85v293H15V204z   M406,497h-20c-17.093,0-31-13.906-31-31s13.907-31,31-31h20c17.093,0,31,13.906,31,31S423.093,497,406,497z M426,420h-40  c-17.093,0-31-13.906-31-31s13.907-31,31-31h40c17.093,0,31,13.906,31,31S443.093,420,426,420z M446,343h-60  c-17.093,0-31-13.906-31-31s13.907-31,31-31h60c17.093,0,31,13.906,31,31S463.093,343,446,343z M466,266h-80  c-17.093,0-31-13.906-31-31s13.907-31,31-31h80c17.093,0,31,13.906,31,31S483.093,266,466,266z"/>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
        </g>
        <g>
            </g><g></g></svg>
          </span>
        </div>
       <div class="bg-red-600 w-[350px] sm:w-[420px] float-end 2xl:w-full xl:mr-5 2xl:mx-6 rounded-2xl">
          <div class="mt-6 bg-white inline-block opacity-50 rounded-r-2xl pr-2 pt-2.5 pb-2.5">
            Lorem ipsum dolor sit amet.
          </div>
          <div></div>
          <div class="w-70 mt-4 bg-white inline-block opacity-70 rounded-r-2xl pr-2 pt-2.5 pb-2.5">
            Lorem ipsum dolor sit amet dg.
          </div>
          <div></div>
          <span class="inline-flex float-start pt-6 pr-1 sm:p-4 xl:p-4">
            <svg class="text-white w-6 h-6 sm:w-7 sm:h-7 xl:w-10 xl:h-10 rotate-180">
              <use href="#chevron-left"></use>
            </svg>
          </span>
          <div class="w-80 mt-4 bg-white inline-block opacity-90 rounded-r-2xl pr-2 pt-2.5 pb-2.5">
            Lorem ipsum dolor sit, amet dol.
          </div>
          <div></div>
          <div class="w-70 mt-4 bg-white inline-block opacity-70 rounded-r-2xl pr-2 pt-2.5 pb-2.5">
            Lorem ipsum dolor sit amet dg.
          </div>
          <div></div>
          <div class="mt-4 mb-6 bg-white inline-block opacity-50 rounded-r-2xl pr-2 pt-2.5 pb-2.5">
            Lorem ipsum dolor sit amet.
          </div>
      </div>
      </div>
    </div>
  </section> 
 </main>
</template>

<script setup lang="ts">
import { ref } from 'vue';
const number = ref(0);
const reacted = ref(false)
function nextNumber () {
  reacted.value ? number.value-- : number.value++
  reacted.value=!reacted.value
}
</script>

<style>
.no-select {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>