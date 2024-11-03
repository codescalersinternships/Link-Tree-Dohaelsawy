import { defineStore } from 'pinia'
import axios from '@/plugins/axios'
import type { Analytics, APIResponse, ErrorRes, User } from '../types/index'
import type { AxiosError } from 'axios';

axios.defaults.withCredentials = true;

export const useAccountStore = defineStore('AccountStore', {
    state: () => ({
        user: {} as User,
    }),

    actions: {


        async getAnalytics(user_id: string): Promise<Analytics[]> {
            return new Promise<Analytics[]>(async (resolve, reject) => {
                try {
                    const { data } = await axios.get<APIResponse<{ analytics: Analytics[]}>>(
                        `/analytics/get_analytics/${user_id}`
                    );
                    console.log('Success getting analytics', data.data.analytics);
                    resolve(data.data.analytics);
                } catch (error) {
                    const err = error as AxiosError
                    console.log(err.response?.data)
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                    reject(error);
                }
            })
        },

        async getAccount(): Promise<User> {
            return new Promise<User>(async (resolve, reject) => {
                try {
                    const { data } = await axios.get<APIResponse<{ user: User }>>(
                        '/account/get_account'
                    );
                    this.user = data.data.user
                    localStorage.setItem("currentUser", JSON.stringify(data.data.user));
                    console.log('Success fetching account', data.data.user);
                    resolve(this.user);
                } catch (error) {
                    const err = error as AxiosError
                    console.log(err.response?.data)
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                    reject(error);
                }
            })
        },

        async updateAccount(form: Record<string, string>): Promise<User> {
            return new Promise<User>(async (resolve, reject) => {

                try {

                    const { data } = await axios.put<APIResponse<{ user: User }>>('account/edit_account', {
                        ...form
                    });

                    console.log('Success updating user', data.data.user);
                    resolve(data.data.user)
                } catch (error) {
                    const err = error as AxiosError
                    console.log(err.response?.data)
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }
            })

        },

        async deleteAccount() {
            return new Promise(async (resolve, reject) => {
                try {
                    const { data } = await axios.delete('/account/delete_account');
                    console.log('user', data.data);
                    resolve(data.data)
                } catch (error) {
                    const err = error as AxiosError
                    console.log(err.response?.data)
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }

            })

        },

        async updateImageAccount(image: FormData) {

            return new Promise<User>(async (resolve, reject) => {

                try {
                    const { data } = await axios.post(
                        '/account/add_photo',
                        image,
                        {
                            headers:
                                { 'Content-Type': 'multipart/form-data' }
                        });
                    localStorage.setItem("currentUserImage", data.data.user.image);
                    resolve(data.data.user)
                } catch (error) {
                    const err = error as AxiosError
                    console.log(err.response?.data)
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }

            })

        },

    },
})