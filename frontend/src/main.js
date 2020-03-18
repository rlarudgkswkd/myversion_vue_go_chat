import Vue from 'vue'
import App from './App.vue'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.css'
import 'vue-material/dist/theme/black-green-light.css'
import Directives from '../plugin/directives'

//import io from 'socket.io-client';

// socket.on('서버에서 받는 이벤트 이름', function("데이터")){}
// socket.emit('서버로 보낼 이벤트 명', 데이터);

Vue.use(VueMaterial)
Vue.use(Directives)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
