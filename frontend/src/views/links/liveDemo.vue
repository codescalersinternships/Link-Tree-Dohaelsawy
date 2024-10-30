<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import { useLinkStore } from '@/stores/linkStore'
import type { Link } from '@/types/index';
import Header from '@/components/header.vue';

props: ['username'];

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
    }
}

onMounted(fetchLinks);

</script>

<template>
    <div class="wrapper">
        <div class="content">

            <Header></Header>

            <div class="container">
                <div class="sub-container">
                    <div v-if="isEmpty" class="empty">
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
