<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/authStore';
import { ref } from 'vue';



const router = useRouter();
const authStore = useAuthStore();
const isLogin = authStore.returnIsLogin

const onSubmitLogout = async () => {
    try {
        await authStore.logoutUser();
        window.location.reload();
    } catch (error) {
        console.error('Logout failed', error);
    }
};


const onSubmitLogin = async () => {
    router.push('/auth/login');
};


const getIsLogin = (isLogin: boolean) => {
    return isLogin
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
                    <a href="#" class="nav-link">About</a>
                </li>
                <li class="nav-item">
                    <a href="/link_tree/" class="nav-link">Link Tree</a>
                </li>
                <li v-if="!isLogin" class="nav-item">
                    <Button variant="secondary" @click="onSubmitLogin">Login</Button>
                </li>
                <li v-else class="nav-item">
                    <a href="/link_tree/" class="nav-link">Profile</a>
                    <Button variant="secondary" @click="onSubmitLogout">Logout</Button>
                </li>
            </ul>
        </nav>
    </header>
</template>


<style>
.navbar {
    padding-top: 20px;
    display: flex;
    text-align: center;
    align-items: center;
    justify-content: center
}


.nav-links {
    flex-direction: row;
    list-style: none;
    display: flex;
    text-align: center;
    align-items: center;
    justify-content: center;
}

.nav-item a{
    color: #fff;
    text-decoration: none;
    font-weight: 500;
    margin: 0px 20px;
    padding: 10px;
}
</style>