<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import type { Link } from '@/types/index';
import { RouterLink, useRouter } from 'vue-router';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useAuthStore } from '@/stores/authStore'
import { useLinkStore } from '@/stores/linkStore'

import { BoxIcon } from '@radix-icons/vue';

type PAYLOAD = {
    name: string;
    url: string;
};
const isEmpty = ref(true);
const username = ref(localStorage.getItem("currentUsername"));
const form = ref<PAYLOAD>({
    name: '',
    url: '',
});
const router = useRouter();
const authStore = useAuthStore();
const linkStore = useLinkStore();
const handleLogout = async () => {
    try {
        await authStore.logoutUser();
        authStore.turnOffLogin()
        router.push('/');
    } catch (error) {
        console.error('Logout failed', error);
    }
};

const links: Ref<Link[]> = ref([]);
const fetchLinks = async () => {
    try {
        if (username.value !== null) {
            links.value = await linkStore.getLinks(username.value);

            if (links.value.length !== 0) {
                isEmpty.value = false;
            }
        }
    } catch (error) {
        console.error("Error fetching links:", error);
    }
}
onMounted(fetchLinks);
</script>

<template>
    <div class="wrapper">
        <div class="content">
            <nav class="navbar">
                <ul class="nav-links">
                    <li class="nav-item">
                        <a href="/" class="nav-link">Home</a>
                    </li>
                    <li class="nav-item">
                        <a href="#" class="nav-link">About</a>
                    </li>
                    <li class="nav-item">
                        <a href="#" class="nav-link">Future Work</a>
                    </li>
                    <li class="nav-item">
                        <a href="/account/get_account" class="nav-link">Profile</a>
                    </li>
                    <Button variant="secondary" @click="handleLogout">Logout</Button>
                </ul>
            </nav>

            <div class="container">
                <div class="sub-container">
                    <div v-if="isEmpty">
                        <img src="../../assets/website-design.png" alt="empty">
                    </div>
                    <div v-else>
                        <div class="circular--landscape">
                            <img src="../../assets/profile.png" />
                        </div>
                        <div class="username">
                            <p class="username">{{ username }}</p>
                        </div>
                        <div v-for="link in links" class="link">
                            <p class="link-name">
                                {{ link.name }}
                            </p>
                            <a :href="link.url" class="link-url">
                                {{ link.url }}
                            </a>
                        </div>

                    </div>
                    <div class="buttons">
                        <Button>Add Link</Button>
                        <Button>Live Demo</Button>
                    </div>

                </div>

            </div>
        </div>
    </div>


</template>




<style>
* {
    padding: 0;
    margin: 0;
}

.wrapper {
    position: relative;
    width: 100%;
    background-image: url("../src/assets/pexels-codioful-6985048.jpg");
    background-size: cover;
    background-position: center;
    background-repeat: repeat-y;
    font-family: "Noto Sans Mono", monospace;
    height: auto;
}

.content {
    background-color: #00000085;
}


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

.nav-item {
    color: #fff;
    text-decoration: none;
    font-weight: 500;
    margin: 0px 20px;

}

.container {
    top: 0;
    bottom: 0;
    margin-top: 60px;
    /* padding-bottom: 50px; */
    padding: 10px;
    display: flex;
    border: 10px;
    border-radius: 20px;
    background-color: rgba(255, 245, 232, 0.341);
    text-align: center;
    justify-content: center;
    flex-direction: column;
    overflow-y: scroll;

}


.sub-container {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    height: max-content;
}


Button {
    width: 150px;
    height: 100px;
}

.buttons {
    display: flex;
    flex-direction: row;
    gap: 30px;
    margin-top: 20px;
}



.link {
    display: flex;
    text-align: center;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    background-color: #001d18e3;
    border-radius: 20px;
    margin-top: 20px;
    width: 500px;
    height: 10vh;
    color: #ffffff;
}

.link-name {
    width: 350px;
    color: #e8fef5;
    border-radius: 10px;
    padding: 5px;
    font-weight: 800;
    font-size: 20px;
}

.username {
    font-weight: 800;
    font-size: 20px;
    color: #ffffff;
}



.circular--landscape {
    display: inline-block;
    position: relative;
    width: 200px;
    height: 120px;
    overflow: hidden;
    border-radius: 50%;
}

.circular--landscape img {
    width: 100px;
    height: 100px;
    margin-left: 50px;
}
</style>