import axios from 'axios';
import store from '../state/store'

var baseUrl = ``;

export const apiClient = axios.create({
    baseURL: 'http://localhost:8080/api/v1/',
    withCredentials: false, // This is the default
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
    },
    timeout: 10000
});

apiClient.interceptors.request.use((config) => {
    if (config.url.indexOf('/auth/') === -1) {
        let token = store.getters['auth/getToken'];
        config.headers['Authorization'] = 'Bearer ' + token;
        config.withCredentials = true;
    }
    return config;
}, (error) => {
    // Do something with request error
    return Promise.reject(error);
});





