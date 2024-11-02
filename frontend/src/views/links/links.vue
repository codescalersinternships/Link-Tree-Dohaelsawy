<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { useLinkStore } from '@/stores/linkStore'
import { useAuthStore } from '@/stores/authStore'
import Header from '@/components/header.vue';
import type { Link } from '@/types/index';

const authStore = useAuthStore();
const isEmpty = ref(true);
const username = ref(localStorage.getItem("currentUsername"));
const linkStore = useLinkStore();
const defaultImage = 'src/assets/profile.png';
const userImage = ref(localStorage.getItem("currentUserImage"))


const links: Ref<Link[]> = ref([]);

type UpdateLink = {
    link_id: number;
    wantUpdate: boolean;
}
const wantUpdateLink = ref<UpdateLink>({
    wantUpdate: false,
    link_id: 0
});
const wantAddLink = ref(false);

type PAYLOAD = {
    name: string;
    url: string;
};

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
        alert(error);
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
        alert(error);
    }
}

const onSubmitUpdateLink = async (link_id: number) => {
    try {
        await linkStore.updateLink(link_id, form.value);
        window.location.reload();
    } catch (error) {
        console.error('Logout failed', error);
        alert(error);
    }
}

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
        <div class="content">

            <Header></Header>

            <div class="container">
                <div class="sub-container">
                    <div v-if="isEmpty" class="empty">
                        <h2>Nothing!..</h2>
                        <img src="/src/assets/website-design.png" alt="empty">
                    </div>
                    <div v-else>
                        <div class="profile-image">
                            <img :src="userImage || defaultImage" alt="Profile Image" />
                        </div>
                        <div class="username">
                            <p class="username">{{ username }}</p>
                        </div>

                        <div v-for="link in links" class="link">
                            <div class="link-content">
                                <div v-if="wantUpdateLink.wantUpdate === true && wantUpdateLink.link_id === link.id">
                                    <form @submit.prevent="onSubmitUpdateLink(link.id)">
                                        <div class="grid gap-2 flex flex-col justify-center items-center">
                                            <div class="link-content">
                                                <Input id="name" type="text" v-model="form.name" />
                                                <Input id="url" type="text" v-model="form.url" />
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
