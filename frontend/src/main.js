import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue'
import router from './router'
import store from './store'
import axios from "axios";


Vue.config.productionTip = false;

axios.defaults.baseURL='https://veric-mvp.netlify.app/api';
Vue.prototype.$axios = axios;

Vue.use(ElementUI);

router.beforeEach((to, from, next) => {
  if (to.name !== 'home' && store.state.apiToken === "") next({ name: 'home' })
  else next()
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
