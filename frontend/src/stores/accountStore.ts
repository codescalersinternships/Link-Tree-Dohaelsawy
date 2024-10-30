import { defineStore } from 'pinia'
import axios from '@/plugins/axios'
import type { APIResponse, User } from '../types/index'

axios.defaults.withCredentials = true;

export const useAccountStore = defineStore('AccountStore', {
    state: () => ({
        user: {} as User,
    }),

    actions: {

        async getAccount(): Promise<User> {
            return new Promise<User>(async (resolve, reject) => {
                try {
                    const { data } = await axios.get<APIResponse<{ user: User }>>(
                        '/account/get_account'
                    );  
                    this.user = data.data.user           
                    console.log('Success fetching account', data.data.user);
                    resolve(this.user);
                } catch (error) {
                    reject(error);
                }
            })
        },

        async updateAccount(form: Record<string, string>): Promise<User> {
            return new Promise<User>(async (resolve, reject) => {

                try {

                    const { data } = await axios.put<APIResponse<{ user: User }>>('account/edit_account/', {
                        ...form
                    });

                    console.log('Success updating user', data.data.user);

                    resolve(data.data.user)
                } catch (error) {
                    reject(error)
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
                    reject(error)
                }

            })

        },



        async updateImageAccount(image: FormData) {

            return new Promise<User>(async (resolve, reject) => {

                try {
                    const { data } = await axios.post( 
                        '/account/add_photo/',
                        image,
                        {headers: 
                            {'Content-Type': 'multipart/form-data'}
                        });
                    resolve(data.data.user)
                } catch (error) {
                    console.log(error)
                    reject(error)
                }

            })

        },

    },
})