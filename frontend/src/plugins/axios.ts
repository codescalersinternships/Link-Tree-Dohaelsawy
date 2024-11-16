import axios from 'axios'

const axiosInstance = axios.create({
    baseURL: import.meta.env.VITE_AXIOS_API
})


const envVars = import.meta.env;
console.log(envVars);

axiosInstance.interceptors.request.use((config) => {

    const token =localStorage.getItem('Authorization');
    if (token) {

        config.headers = config.headers || {};
        const unprotectedURls = ['/auth/login/', '/auth/register/','/account/search'];

        if (config.url && !unprotectedURls.includes(config.url)) {

            config.headers.Authorization = token ? `${token}` : '';

        }

    }

    return config

})


export default axiosInstance;