import UserService from '@services/UserService';

export const mutations = {
    GET_PROFILE_SUCCESSFUL(state, user) { }
};

export const actions = {

    signUp({ commit }, user) {
        const data = {
            name: user.name,
            password: user.password,
            username: user.username,
            private: user.private
        };

        return new Promise((resolve, reject) => {
            // Do something here... lets say, a http call using vue-resource
            UserService.signUp(data).then(response => {
                // http success, call the mutator and change something in state
                resolve(response);  // Let the calling function know that http is done. You may send some data back
                const data = response.data.data;
                console.log(resp.data, data)
                commit('SIGNUP_SUCCESSFUL', data);
            }, error => {
                // http failed, let the calling function know that action did not work out
                commit('SIGNUP_FAILED', 'error');
                reject(error);
            })
        })
    },

    getProfile({ commit }, payload) {

        return new Promise((resolve, reject) => {
            // Do something here... lets say, a http call using vue-resource
            UserService.getProfile(payload).then(response => {
                // http success, call the mutator and change something in state
                resolve(response);  // Let the calling function know that http is done. You may send some data back
                const data = response.data.data;

                commit('GET_PROFILE_SUCCESSFUL', data);
            }, error => {
                // http failed, let the calling function know that action did not work out
                commit('GET_PROFILE_FAILED', 'error');
                reject(error);
            })
        })
    },

    toggleFollow({ commit }, payload) {

        return new Promise((resolve, reject) => {
            // Do something here... lets say, a http call using vue-resource
            UserService.toggleFollow(payload).then(response => {
                // http success, call the mutator and change something in state
                resolve(response);  // Let the calling function know that http is done. You may send some data back

                commit('TOGGLE_SUCCESSFUL');
            }, error => {
                // http failed, let the calling function know that action did not work out
                commit('TOGGLE_FAILED', 'error');
                reject(error);
            })
        })
    }
};

export const state = {
    user: null
};

export const getters = {};

