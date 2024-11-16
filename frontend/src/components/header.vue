<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/authStore';
import { useAccountStore } from '@/stores/accountStore';

import { ref } from 'vue';



const router = useRouter();
const authStore = useAuthStore();
const accountStore = useAccountStore();
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

const searchTerm = ref("")
const searchResult = ref("")
const setSearchTerm = (event: Event) => {
    const input = event.target as HTMLInputElement
    searchTerm.value = input.value
    // console.log(searchTerm.value)
}

const sendSearchTerm = async () => {
    if (searchTerm.value !== "") {
        try {
            searchResult.value = await accountStore.getUsernameSearch(searchTerm.value);
            router.push(`/link_tree/${searchResult.value}`).then(() => {
            window.location.reload();
        });
        } catch (error) {
            router.push("/error")
        }
    }
}

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
                    <a @click="onSubmitLinkTree" class="nav-link" cy="link-tree">Link Tree</a>
                </li>
                <li v-show="isLogin" class="nav-item">
                    <a class="nav-link" :href="profileUrl" cy="profile">Profile</a>
                </li>
                <li class="nav-item">
                    <input type="search" name="search" id="search" placeholder="search username"
                        class="bg-transparent  text-white outline-white px-3 py-3 border-white placeholder:text-white rounded-lg border-2"
                        :onchange="setSearchTerm" cy="search-input">
                    <i class="fa-solid fa-magnifying-glass text-white px-3" @click="sendSearchTerm" cy="search-btn"></i>
                </li>
                <li v-show="!isLogin" class="nav-item">
                    <Button variant="secondary" @click="onSubmitLogin" cy="login-btn">Login</Button>
                </li>
                <li v-show="isLogin" class="nav-item">
                    <Button variant="secondary" @click="onSubmitLogout" cy="logout-btn">Logout</Button>
                </li>
            </ul>
        </nav>
    </header>
</template>