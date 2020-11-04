import { apiClient } from "./ApiService";

export default {

    list(data) {
        let url = "/feeds"
        if (data.username) {
            url += "?username=" + data.username + "&type=profile"
        }
        if (data.type == "home") {
            url += "?type=home"
        }
        return apiClient.get(url, data);
    },


};