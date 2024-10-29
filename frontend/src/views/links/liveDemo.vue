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
}

.content {
    background-color: #00000085;
    padding-bottom: 50px;
    min-height: 100vh;
    display: flex;
    text-align: center;
    align-items: center;
    justify-content: center;
    flex-direction: column;
}

.container {
    top: 0;
    bottom: 0;
    margin-top: 60px;
    padding: 10px;
    display: flex;
    border: 10px;
    border-radius: 20px;
    background-color: rgba(255, 245, 232, 0.341);
    text-align: center;
    justify-content: center;
    flex-direction: column;
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