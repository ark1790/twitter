import Vue from 'vue';
import App from './App';
import store from '@state/store';
import router from '@router';
import Noty from 'vuejs-noty';

import 'normalize.css/normalize.css'


import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';


Vue.use(ElementUI);

import moment from 'moment'




Vue.use(Noty, {
  timeout: 1000
});

// Don't warn about using the dev version of Vue in development
Vue.config.productionTip = process.env.NODE_ENV === 'production';


Vue.prototype.moment = moment;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
