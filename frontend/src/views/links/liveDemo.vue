<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import { useLinkStore } from '@/stores/linkStore';
import { useAccountStore } from '@/stores/accountStore';
import type { Link, User } from '@/types/index';
import Header from '@/components/header.vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const username = route.params
const defaultImage = '/src/assets/profile.png';
const isEmpty = ref(true);
const linkStore = useLinkStore();
const links: Ref<Link[]> = ref([]);

const accountStore = useAccountStore();
const user = ref({} as User);
const fetchLinks = async () => {
    try {
        if (username.username !== null) {
            links.value = await linkStore.getLinks(username.username as string);
            user.value = await accountStore.getAccountByUsername(username.username as string);
            console.log(user.value.image)
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
                    <div class="profile-image">
                        <img :src="user.image || 'https://link-tree.s3.eu-north-1.amazonaws.com/profile.png'" />
                    </div>
                    <div class="data">
                        <p class="username">{{ username.username }}</p>
                    </div>
                    <div class="data">
                        <p class="username">{{ user.phone }}</p>
                    </div>
                    <div class="data">
                        <p class="username">{{ user.bio }}</p>
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
</template>
