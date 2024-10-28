<script setup lang="ts">

import { Button } from '@/components/ui/button'
import { computed, onMounted, ref } from 'vue';
import { RouterLink, useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/authStore';



const router = useRouter();
const authStore = useAuthStore();

const loginFlag = computed (() => (authStore.isLogin))

const onSubmit = async () => {
  router.push('/auth/login');
};


const handleLogout = async () => {
  try {
    await authStore.logoutUser();
    authStore.turnOffLogin()
    router.push('/');
  } catch (error) {
    console.error('Logout failed', error);
  }
};


</script>

<template>

  <div class="wrapper">
    <div class="content">
      <header>
        <nav class="navbar">
          <ul class="nav-links">
            <li class="nav-item">
              <a href="/" class="nav-link">Home</a>
            </li>
            <li class="nav-item">
              <a href="#" class="nav-link">About</a>
            </li>
            <li class="nav-item">
              <a href="/link_tree/" class="nav-link">Link Tree</a>
            </li>
            <div v-if="loginFlag">
              <li class="nav-item">
                <a href="/account/get_account" class="nav-link">Profile</a>
              </li>
              <Button variant="secondary" @click="handleLogout">Logout</Button>
            </div>
            <div v-else>
              <Button variant="secondary" @click="onSubmit">Login</Button>
            </div>

          </ul>
        </nav>
      </header>

      <div class="main">
        <p class="fancy-word">Your Links, Your Story</p>
        <p class="normal">Let's share it!</p>
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
  justify-content: center;
  align-items: center;
  flex-direction: column;
  color: #fff;
  height: 100vh;
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