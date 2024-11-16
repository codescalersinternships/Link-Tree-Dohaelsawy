import { defineStore } from 'pinia'
import axios from '@/plugins/axios'
import type { APIResponse, ErrorRes, Link, User } from '../types/index'
import type { AxiosError } from 'axios'

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

                    resolve(data.data.link)
                } catch (error) {
                    const err = error as AxiosError
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }

            })

        },


        async updateLink(link_id:number, form: Record<string, string>) {

            return new Promise<Link>(async (resolve, reject) => {

                try {

                    const { data } = await axios.put<APIResponse<{ link: Link }>>(`/link/update_link/${link_id}`, {
                        ...form
                    });
                    resolve(data.data.link)
                } catch (error) {
                    const err = error as AxiosError
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }
            })
        },

        async getLinks(username: string): Promise<Link[]> {
            return new Promise<Link[]>(async (resolve, reject) => {
                try {
                    const { data } = await axios.get<APIResponse<{links:Link[]}>>(`/link_tree/${username}`);
                    resolve(data.data.links)
                } catch (error) {
                    const err = error as AxiosError
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }

            })
                    
        },

        async deleteLink(link_id: number) {
            return new Promise(async (resolve, reject) => {
                try {

                    const { data } = await axios.delete(`/link/delete_link/${link_id}`);
                    resolve(data.data)
                } catch (error) {
                    const err = error as AxiosError
                    const response = err.response?.data as ErrorRes
                    reject(response.error)
                }

            })

        },

    },
})