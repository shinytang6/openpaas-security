import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue'
import router from './router'
import axios from 'axios'

Vue.use(ElementUI);
Vue.config.productionTip = false
Vue.prototype.$axios=axios

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
