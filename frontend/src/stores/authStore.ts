import { defineStore } from 'pinia'
import axios from '@/plugins/axios'
import type { APIResponse, User, ErrorRes } from '../types/index'
import { ref } from 'vue';
import router from '@/router';
import type { AxiosError } from 'axios';

axios.defaults.withCredentials = true;

export const useAuthStore = defineStore('AuthStore', {
    state: () => ({
        user: {} as User,
        isLogin: false,
    }),

    getters: {
        returnIsLogin(): boolean {
            return localStorage.getItem("currentUsername") !== ""
        },
    },


    actions: {



        async registerUser(form: Record<string, string>) {

            return new Promise<any>(async (resolve, reject) => {

                try {

                    const { data } = await axios.post<APIResponse<{ user: User }>>('/auth/register', {
                        ...form
                    });
                    console.log('Success Registration', data.data);
                    resolve(data.data)
                } catch (error) {
                    const err = error as AxiosError
                    console.log(err.response?.data)
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }

            })

        },


        async loginUser(form: Record<string, string>) {

            return new Promise<string>(async (resolve, reject) => {

                try {

                    const { data } = await axios.post<APIResponse<{ access_token: string, user: User }>>('/auth/login', {
                        ...form
                    });
                    console.log('Success Login ana henaaaaa', data.data.access_token);


                    function setCookie(name: string, value: string, days: number) {
                        const expires = new Date(Date.now() + days * 864e5).toUTCString();
                        document.cookie = `${name}=${encodeURIComponent(value)}; expires=${expires}; path=/`;
                    }

                    setCookie("Authorization", data.data.access_token, 3);
                    localStorage.setItem("currentUser", JSON.stringify(data.data.user));
                    localStorage.setItem("currentUsername", data.data.user.username);
                    this.isLogin = true;

                    resolve(data.data.access_token)
                } catch (error) {
                    const err = error as AxiosError
                    console.log(err.response?.data)
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }

            })

        },


        async logoutUser() {

            return new Promise<string>(async (resolve, reject) => {

                try {
                    const { data } = await axios.get('/auth/logout', {});
                    localStorage.setItem("currentUser", "");
                    localStorage.setItem("currentUsername", "");
                    this.isLogin = true;
                    router.push("/");
                    resolve(data)
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