<script setup lang="ts">
import type { User } from '@/types';
import { onMounted, ref, type Ref } from 'vue';
import { useAccountStore } from '@/stores/accountStore'
import { Button } from '@/components/ui/button';


const accountStore = useAccountStore();
const defaultImage = '/src/assets/user.png';

const user = ref({
    first_name: "",
    last_name: "",
    phone: "",
    image: "",
    bio: "",
});

const form = ref({
    media: {},
});

const imageSrc = ref([]);


const updateProfile = async () => {
    try {
        user.value = await accountStore.updateAccount(user.value);
    } catch (error) {
        console.error("Error fetching user:", error);
    }
}

const handelImageUpload = async (e: Event) => {
    // var files = e.target.files || e.dataTransfer.files;

};

const getAccountData = async () => {
    try {
        user.value = await accountStore.getAccount();
        if (user.value.image === "") {
            user.value.image = defaultImage;
        }
    } catch (error) {
        console.error("Error fetching user:", error);
    }
}
onMounted(getAccountData);
</script>

<template>
    <h2>Edit Profile</h2>
    <div class="edit-profile-container">
        <div v-if="user">
            <form @submit.prevent="updateProfile">
                <!-- Profile Image -->
                <div class="profile-image">
                    <input type="file" id="media" accept="image/*" multiple
                        @change="(event) => handelImageUpload(event)" />
                    <img :src="user.image" alt="Profile Image" />
                </div>

                <!-- First Name -->
                <div class="form-group">
                    <label for="firstName">First Name</label>
                    <input type="text" id="firstName" v-model="accountStore.user.first_name" />
                </div>

                <!-- Last Name -->
                <div class="form-group">
                    <label for="lastName">Last Name</label>
                    <input type="text" id="lastName" v-model="accountStore.user.last_name" />
                </div>

                <!-- Phone -->
                <div class="form-group">
                    <label for="phone">Phone</label>
                    <input type="tel" id="phone" v-model="accountStore.user.phone" />
                </div>

                <!-- Bio -->
                <div class="form-group">
                    <label for="bio">Bio</label>
                    <textarea id="bio" v-model="accountStore.user.bio"></textarea>
                </div>

                <!-- Submit Button -->
                <Button variant="default" @click="updateProfile">Submit</Button>
            </form>
        </div>

    </div>
</template>