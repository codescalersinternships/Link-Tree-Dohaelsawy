<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/authStore';
import { ref } from 'vue';



const router = useRouter();
const authStore = useAuthStore();
const isLogin = ref(authStore.returnIsLogin);

const onSubmitLogout = async () => {
    try {
        await authStore.logoutUser();
        router.push('/').then(() => {
            window.location.reload();
        });
    } catch (error) {
        console.error('Logout failed', error);
        alert(error);
    }
};

const onSubmitLinkTree = () => {
    if (isLogin.value) {
        router.push('/link_tree');
    } else {
        router.push('/auth/login');
    }
}

const profileUrl = ref("/account/get_account/")

const onSubmitLogin = async () => {
    router.push('/auth/login');
};

</script>

<template>
    <header>
        <nav class="navbar">
            <ul class="nav-links">
                <li class="nav-item">
                    <a href="/" class="nav-link">Home</a>
                </li>
                <li class="nav-item">
                    <a href="/about" class="nav-link" cy="about-me">About Me</a>
                </li>
                <li class="nav-item">
                    <a @click="onSubmitLinkTree" class="nav-link" href="">Link Tree</a>
                </li>
                <li v-if="!isLogin" class="nav-item">
                    <Button variant="secondary" @click="onSubmitLogin" cy="login-btn">Login</Button>
                </li>
                <li v-if="isLogin" class="nav-item">
                    <a class="nav-link" :href="profileUrl" cy="profile">Profile</a>
                </li>
                <li v-if="isLogin" class="nav-item">
                    <Button variant="secondary" @click="onSubmitLogout" cy="logout-btn">Logout</Button>
                </li>
            </ul>
        </nav>
    </header>
</template>