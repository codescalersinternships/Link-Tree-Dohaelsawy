<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import { useLinkStore } from '@/stores/linkStore'
import type { Link } from '@/types/index';
import Header from '@/components/header.vue';
import Profile from '@/components/profile.vue';

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

const analyticsUrl = ref(`/get_analytics/${localStorage.getItem("currentUserID")}`);
console.log(analyticsUrl.value);

onMounted(fetchLinks);

</script>

<template>
    <div class="wrapper">
        <div class="content">
            <Header></Header>
            <div class="container">
                <div class="sub-container">
                    <div class="profile-bar">
                        <a href="/account/get_account/">Edit Profile</a>
                        <a :href="analyticsUrl"> Show Analytics</a>
                    </div>
                    <Profile></Profile>
                </div>
            </div>
        </div>

    </div>
</template>
