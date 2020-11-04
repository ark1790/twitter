import { apiClient } from "./ApiService";

export default {

    postTweet(data) {
        return apiClient.post('/tweets', data);
    },


};