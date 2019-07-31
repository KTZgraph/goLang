import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import 'vuetify/src/stylus/app.styl'

Vue.use(Vuetify, {
  iconfont: 'md',
  theme: { //nadpisanie słów kluczowych
    primary: '#212121',
    success: '#EEEEEE',
    info: '#ffaa2c',
    error: '#D50000',
    myTheme: '#483D8B'
  }
})
