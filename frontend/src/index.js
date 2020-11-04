import Vue from 'vue';
import VueRouter from 'vue-router';

import NProgress from 'nprogress/nprogress';

// import store from '@state/store';
import routes from './routes';

Vue.use(VueRouter);

const router = new VueRouter({
  routes,
  mode: 'history',
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    return { x: 0, y: 0 };
  }
});


router.beforeResolve((routeTo, routeFrom, next) => {
  if (routeFrom.name) {
    NProgress.start();
  }

  next();
});

router.afterEach((routeTo, routeFrom) => {
  NProgress.done();
});

export default router;
