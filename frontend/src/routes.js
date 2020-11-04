// import Main from '@components/layouts/Main'
// import PreAuth from "@/components/layouts/PreAuth";

export default [
  {
    path: '/404',
    name: '404',
    component: require('@components/util/404').default
  },
  {
    path: '/',
    name: 'home',
    component: require('@components/home/home').default
  },
  {
    path: '/profile/:username',
    name: 'profile',
    component: require('@components/user/profile').default
  },
  {
    path: '/post',
    name: 'post',
    component: require('@components/tweet/post').default
  },
  {
    path: '/auth',
    name: 'auth',
    component: require('@components/user/login').default
  },
  {
    name: 'register',
    path: '/auth/signup',
    component: require('@components/user/register').default
  },
  {
    path: '*',
    name: 'not-found',
    component: require('@components/util/404').default
  },

];
