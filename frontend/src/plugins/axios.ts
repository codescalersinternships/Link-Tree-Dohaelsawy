import axios from 'axios'

const axiosInstance = axios.create({
    baseURL: "http://185.206.122.17:31010"
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