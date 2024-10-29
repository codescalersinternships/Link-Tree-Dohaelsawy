<script setup lang="ts">

import { Button } from '@/components/ui/button'
import Header from '@/components/header.vue'

import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/authStore';



const router = useRouter();
const authStore = useAuthStore();

const onSubmit = async () => {
  router.push('/auth/login');
};





</script>

<template>

  <div class="wrapper">
    <div class="content">
      <Header></Header>
      <div class="main">
        <p class="fancy-word">Your Links, Your Story</p>
        <p class="normal">share it!</p>
        <Button class="start-button" @click="onSubmit">Get Started!</Button>
      </div>
    </div>
  </div>



</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Lobster&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Lobster&family=Poiret+One&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Lobster&family=Noto+Sans+Mono:wght@100..900&family=Poiret+One&display=swap');


.main {
  display: flex;
  text-align: center;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  color: #fff;
  min-height: 100vh;
}


.fancy-word {
  font-family: "Lobster", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 120px;
}


.normal {
  font-size: 50px;
  font-family: "Poiret One", sans-serif;
  font-weight: 100;
}

.start-button {
  width: 200px;
  height: 50px;
  font-size: large;
}
</style>