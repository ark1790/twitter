import AuthService from '@services/AuthService'

export const state = {
    ...getSavedState('victorTwitter:auth'),
    loggingIn: false
};

export const mutations = {
    AUTH_REQUEST(state) {
        state.loggingIn = false;
    },
    AUTH_REQUEST_SUCCESSFUL(state, token) {

        state.loggingIn = true;
        state.token = token;
        saveState('victorTwitter:auth', state);
        console.log("AUTH_REQUEST_SUCCESSFUL")
    },
    AUTH_REQUEST_FAILED(state) {
        state.loggingIn = false;
    },
    LOG_OUT(state) {
        state.token = '';
        state.user = {};
        saveState('victorTwitter:auth', state);
    }
};

export const getters = {
    isLoggedIn: (state, getters) => {
        return true
    },
    getToken: (state, getters) => {
        return state.token;
    },
    loggingIn: state => false,
    loggedIn: state => !!state.token,
    token: state => state.token,
};

export const actions = {
    logIn({ commit }, user) {
        const data = {
            password: user.password,
            username: user.username
        };
        return new Promise((resolve, reject) => {
            // Do something here... lets say, a http call using vue-resource
            AuthService.login(data).then(response => {
                // http success, call the mutator and change something in state
                resolve(response);  // Let the calling function know that http is done. You may send some data back
                let tkn = response.data.data.token;
                commit('AUTH_REQUEST_SUCCESSFUL', tkn);
            }, error => {
                // http failed, let the calling function know that action did not work out
                commit('AUTH_REQUEST_FAILED', 'error');
                reject(error);
            })
        })
    },
    logout({ commit }) {
        console.log("logging out");
        commit('LOG_OUT');
    }
};

function getSavedState(key) {
    const state = window.localStorage.getItem(key);
    if (state) {
        return JSON.parse(state);
    }
    return {
        token: '',
        user: {}
    };
}

function saveState(key, state) {
    window.localStorage.setItem(key, JSON.stringify(state));
}