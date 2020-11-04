import { apiClient } from "./ApiService";

export default {
    signUp(userData) {
        return apiClient.post('/users', userData);
    },
    getProfile(data) {
        return apiClient.get('/users/' + data.username, data);
    },
    toggleFollow(data) {
        return apiClient.post('/follows', data);
    }
};