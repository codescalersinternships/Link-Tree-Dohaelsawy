<script setup lang="ts">

import { Button } from '@/components/ui/button'
import Header from '@/components/header.vue'

import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/authStore';


const router = useRouter();
const authStore = useAuthStore();

const isLogin = ref(authStore.returnIsLogin)

const onSubmit = async () => {
  if (isLogin.value) {
    router.push('/link_tree');
  } else {
    router.push('/auth/login');
  }
};


</script>

<template>
  <div class="wrapper">
    <Header></Header>
    <div class="content flex flex-col justify-center items-center text-center py-20">
      <h1 class="fancy-word" cy="welcome-word">Your Links, Your Story
        <div class="aurora">
          <div class="aurora__item"></div>
          <div class="aurora__item"></div>
          <div class="aurora__item"></div>
          <div class="aurora__item"></div>
        </div>
      </h1>
      <p class="normal text-white">share it!</p>
      <Button class="start-button" @click="onSubmit" cy="get-start-btn">Get Started!</Button>
    </div>
  </div>



</template>