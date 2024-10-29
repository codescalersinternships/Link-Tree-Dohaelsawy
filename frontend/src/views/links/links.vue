<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { useLinkStore } from '@/stores/linkStore'

import Header from '@/components/header.vue';

import type { Link } from '@/types/index';


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
type PAYLOAD = {
    name: string;
    url: string;
};

type UpdateLink = {
    link_id: number;
    wantUpdate: boolean;
}
const wantUpdateLink = ref<UpdateLink>({
    wantUpdate: false,
    link_id: 0
});
const wantAddLink = ref(false);


const form = ref<PAYLOAD>({
    name: '',
    url: '',
});
const router = useRouter();

const opposite = (flag: boolean) => {
    return !flag
}

const onSubmitLink = async () => {
    try {
        await linkStore.CreateLink(form.value);
        window.location.reload();
    } catch (error) {
        console.error('Logout failed', error);
    }
}
const onSubmitLiveDemo = () => {
    router.push('/link_tree/' + localStorage.getItem("currentUsername"));
}

const onSubmitDeleteLink = async (link_id: number) => {
    try {
        await linkStore.deleteLink(link_id);
        window.location.reload();
    } catch (error) {
        console.error('Logout failed', error);
    }
}

const onSubmitUpdateLink = async (link_id: number) => {
    try {
        await linkStore.updateLink(link_id, form.value);
        window.location.reload();
    } catch (error) {
        console.error('Logout failed', error);
    }
}
onMounted(fetchLinks);

</script>

<template>
    <link rel="stylesheet"
        href="https://cdn.jsdelivr.net/npm/@fortawesome/fontawesome-free@6.2.1/css/fontawesome.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.6.0/css/all.min.css"
        integrity="sha512-Kc323vGBEqzTmouAECnVceyQqyqdsSiqLQISBL29aUW4U/M7pSPA/gEUZQqv1cwx4OnYxTxve5UMg5GT6L4JJg==" />
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
                                <div v-if="wantUpdateLink.wantUpdate === true && wantUpdateLink.link_id === link.id">
                                    <form @submit.prevent="onSubmitUpdateLink(link.id)">
                                        <div class="grid gap-2 flex flex-col justify-center items-center">
                                            <div class="link-content ">
                                                <Input id="name" type="text" placeholder="link name"
                                                    v-model="form.name" />
                                                <Input id="url" type="text" placeholder="link url" v-model="form.url" />
                                                <Button variant="outline">Submit</Button>
                                            </div>
                                        </div>
                                    </form>
                                </div>
                                <div v-else>
                                    <p class="link-name">{{ link.name }}</p>
                                    <a :href="link.url" class="link-url">{{ link.url }}</a>
                                </div>
                            </div>
                            <div class="link-icon">
                                <i class="fa-solid fa-trash" @click="onSubmitDeleteLink(link.id)"></i>
                                <i class="fa-regular fa-pen-to-square"
                                    @click="wantUpdateLink.wantUpdate = opposite(wantUpdateLink.wantUpdate); wantUpdateLink.link_id = link.id"></i>
                            </div>
                        </div>
                    </div>
                    <div v-if="wantAddLink === true" class="">
                        <form @submit.prevent="onSubmitLink">
                            <div class="grid gap-2  flex flex-col justify-center items-center link">
                                <div class="link-content ">
                                    <Input id="name" type="text" placeholder="link name" v-model="form.name" />
                                    <Input id="url" type="text" placeholder="link url" v-model="form.url" />
                                    <Button variant="outline">Submit</Button>
                                </div>
                            </div>
                        </form>
                    </div>
                    <div class="buttons">
                        <Button @click="wantAddLink = opposite(wantAddLink)">Add Link</Button>
                        <Button @click="onSubmitLiveDemo">Live Demo</Button>
                    </div>

                </div>

            </div>
        </div>
    </div>


</template>




<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;800&family=VT323&display=swap');

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
}

.sub-container {
    display: flex;
    justify-content: center;
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
    flex-direction: row;
    justify-content: space-around;
    padding: 10px 10px;
    text-align: center;
    background-color: #0807435b;
    border-radius: 20px;
    margin: 20px 20px;
    min-width: 360px;
    height: max-content;
    border-radius: 10px;
    color: #e8fef5;
}


.link-icon {
    display: flex;
    flex-direction: row;
    align-items: center;
}

i {
    padding: 5px 10px;
}

.link-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.link-name {
    color: #e8fef5;
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