import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI, {Message} from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import './css/global.css'
import './router/axios'

Vue.prototype.$message = Message
Vue.use(ElementUI);

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
