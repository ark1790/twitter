
import feedService from '../../services/FeedService';

export const mutations = {
    FEED_FETCH_SUCCESSFUL(state, data) {

    }
};


export const getters = {
    feeds: (state) => state.feeds,
};

export const actions = {

    list({ commit }, data) {
        return new Promise((resolve, reject) => {
            // Do something here... lets say, a http call using vue-resource
            feedService.list(data).then(response => {
                // http success, call the mutator and change something in state
                resolve(response);  // Let the calling function know that http is done. You may send some data back
                const data = response.data.data;
                commit('FEED_FETCH_SUCCESSFUL', data);
            }, error => {
                // http failed, let the calling function know that action did not work out
                commit('FEED_FETCH_FAILED', 'error');
                reject(error);
            })
        })
    },
};

export const state = {
    feeds: []
};



