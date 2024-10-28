import { defineStore } from 'pinia'
import axios from '@/plugins/axios'
import type { APIResponse, Link } from '../types/index'
import type { promises } from 'dns'

export const useLinkStore = defineStore('LinkStore', {
    state: () => ({
        link: {} as Link
    }),

    actions: {

        async CreateLink(form: Record<string, string>) {

            return new Promise<Link>(async (resolve, reject) => {

                try {

                    const { data } = await axios.post<APIResponse<{ link: Link }>>('/link/create_link', {
                        ...form
                    });
                    console.log('Success creating link', data.data.link);


                    resolve(data.data.link)
                } catch (error) {
                    reject(error)
                }

            })

        },


        // async UpdateLink(form: Record<string, string>) {

        //     return new Promise<Link>(async (resolve, reject) => {

        //         try {

        //             const { data } = await axios.post<APIResponse<{ link: Link }>>('/link/update_link', {
        //                 ...form
        //             });
        //             console.log('Success creating link', data.data.link);


        //             resolve(data.data.link)
        //         } catch (error) {
        //             reject(error)
        //         }

        //     })

        // },

        async getLinks(username: string): Promise<Link[]> {
            return new Promise<Link[]>(async (resolve, reject) => {
                try {
                    const { data } = await axios.get<APIResponse<{links:Link[]}>>(`/link_tree/${username}`);
                    // console.log('links', data.data);
                    resolve(data.data.links)
                } catch (error) {
                    reject(error)
                }

            })

        },

        async deleteProduct(link_id: string) {
            return new Promise(async (resolve, reject) => {
                try {

                    const { data } = await axios.delete(`/link/delete_link/${link_id}`);
                    console.log('products', data.data);                
                    resolve(data.data)
                } catch (error) {
                    reject(error)
                }

            })

        },

    },
})