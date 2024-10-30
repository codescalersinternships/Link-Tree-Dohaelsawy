<script setup lang="ts">
import type { User } from '@/types';
import { onMounted, ref, type Ref } from 'vue';
import { useAccountStore } from '@/stores/accountStore'
import { Button } from '@/components/ui/button';


const accountStore = useAccountStore();
const defaultImage = '/src/assets/user.png';
// const user = ref<User>();
const username = ref(localStorage.getItem("currentUsername"));
const userReq = ref({
    first_name: accountStore.user.first_name,
    last_name: accountStore.user.last_name,
    phone: accountStore.user.phone,
    bio: accountStore.user.bio,
});
console.log(userReq.value)

const formData = new FormData();
const handleImageUpload = async (event: Event) => {
    const target = event.target as HTMLInputElement;

    if (target && target.files && username.value) {
        formData.append("image", target.files[0],target.files[0].name);
    }

};

const getAccountData = async () => {
    try {
        userReq.value = await accountStore.getAccount();
    } catch (error) {
        console.error("Error fetching user:", error);
    }
}


const updateProfile = async () => {
    try {
        userReq.value = await accountStore.updateAccount(userReq.value);
        userReq.value = await accountStore.updateImageAccount(formData);
        console.log(userReq.value)

    } catch (error) {
        console.error("Error fetching user:", error);
    }
}
onMounted(getAccountData);
</script>

<template>
    <h2>Edit Profile</h2>
    <div class="edit-profile-container">
        <div v-if="userReq">
            <form @submit.prevent="updateProfile" enctype="multipart/form-data">
                <!-- Profile Image -->
                <div class="profile-image">
                    <input type="file" @change="handleImageUpload($event)" capture accept="image/*" class="input-file">
                    <img :src="accountStore.user.image || defaultImage" alt="Profile Image" />
                </div>

                <!-- First Name -->
                <div class="form-group">
                    <label for="firstName">First Name</label>
                    <input type="text" id="firstName" v-model="userReq.first_name" />
                </div>

                <!-- Last Name -->
                <div class="form-group">
                    <label for="lastName">Last Name</label>
                    <input type="text" id="lastName" v-model="userReq.last_name" />
                </div>

                <!-- Phone -->
                <div class="form-group">
                    <label for="phone">Phone</label>
                    <input type="tel" id="phone" v-model="userReq.phone" />
                </div>

                <!-- Bio -->
                <div class="form-group">
                    <label for="bio">Bio</label>
                    <textarea id="bio" v-model="userReq.bio"></textarea>
                </div>

                <!-- Submit Button -->
                <Button variant="default" @click="updateProfile">Submit</Button>
            </form>
        </div>

    </div>
</template>