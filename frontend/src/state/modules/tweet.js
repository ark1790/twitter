
import TweetService from '../../services/TweetService';

export const mutations = {};


export const getters = {
    feeds: (state) => state.feeds,
};

export const actions = {

    postTweet({ commit }, data) {
        return new Promise((resolve, reject) => {
            // Do something here... lets say, a http call using vue-resource
            TweetService.postTweet(data).then(response => {
                // http success, call the mutator and change something in state
                resolve(response);  // Let the calling function know that http is done. You may send some data back
                const data = response.data.data;

                commit('POST_TWEET_SUCCESSFUL', data);
            }, error => {
                // http failed, let the calling function know that action did not work out
                commit('POST_TWEET_FAILED', 'error');
                reject(error);
            })
        })
    },
};

export const state = {
    feeds: []
};



