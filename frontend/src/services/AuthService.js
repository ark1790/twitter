import { apiClient } from "./ApiService";

export default {

    login(userData) {
        return apiClient.post('/users/login', userData);
    },
    signUp(userData) {
        return apiClient.post('/users', userData);
    }

};