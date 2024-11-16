<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue';
import { useAccountStore } from '@/stores/accountStore'
import type { Analytics } from '@/types/index';
import Header from '@/components/header.vue';
import * as moment from 'moment';

const accountStore = useAccountStore();
const analytics = ref({} as Analytics[])
const user_id = ref(localStorage.getItem("currentUserID"))

const getAnalyticsData = async () => {
    try {
        if (user_id.value !== null) {
            analytics.value = await accountStore.getAnalytics(user_id.value);
        }
    } catch (error) {
        console.error("Error fetching user:", error);
        alert(error)
    }
}

const analyticsUrl = ref(`/get_analytics/${localStorage.getItem("currentUserID")}`);


onMounted(getAnalyticsData);

</script>





<template>
    <div class="wrapper">
        <Header></Header>
        <div class="content">
            <div class="container">
                <div class="flex flex-col justify-center items-center">
                    <div class="profile-bar py-3">
                        <a href="/account/get_account/" cy="profile-edit">Edit Profile</a>
                        <a :href="analyticsUrl" cy="profile-analysis"> Show Analytics</a>
                    </div>
                    <table>
                        <thead>
                            <tr>
                                <th> Click Count </th>
                                <th> Guest Username</th>
                                <th> Last Time </th>
                            </tr>
                        </thead>
                        <tbody v-for="analytic in analytics">
                            <tr>
                                <td> {{ analytic.click_count }} </td>
                                <td> {{ analytic.guest_username }}</td>
                                <td> {{ moment.default(analytic.updated_at) }} </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>

    </div>
</template>
