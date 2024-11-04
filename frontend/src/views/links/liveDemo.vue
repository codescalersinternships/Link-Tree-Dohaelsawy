<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import { useLinkStore } from '@/stores/linkStore'
import type { Link, User } from '@/types/index';
import Header from '@/components/header.vue';

props: ['username'];

const defaultImage = '/src/assets/profile.png';
const userImage = ref(localStorage.getItem("currentUserImage"))
const userData = ref<User | null>(JSON.parse(localStorage.getItem("currentUser") || "null") as User | null);
const isEmpty = ref(true);
const username = ref(localStorage.getItem("currentUsername"));
const linkStore = useLinkStore();
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
        alert(error);
    }
}

onMounted(fetchLinks);

</script>

<template>
    <div class="wrapper">
        <Header></Header>
        <div class="content">
            <div class="container">
                <div class="flex flex-col justify-center items-center">
                    <div v-if="isEmpty" class="empty">
                        <img src="/src/assets/website-design.png" alt="empty">
                    </div>
                    <div v-else>
                        <div class="profile-image">
                            <img :src="userImage || defaultImage" />
                        </div>
                        <div class="data">
                            <p class="username">{{ username }}</p>
                        </div>
                        <div class="data">
                            <p class="username">{{ userData?.phone }}</p>
                        </div>
                        <div class="data">
                            <p class="username">{{ userData?.bio }}</p>
                        </div>
                        <div v-for="link in links" class="link">
                            <div class="link-content">
                                <p class="link-name">{{ link.name }}</p>
                                <a :href="link.url" class="link-url">{{ link.url }}</a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
