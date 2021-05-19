import Vue from 'vue';
import Vuex from 'vuex';
import { State } from '~/types/Store/StateType';
import createPersistedState from 'vuex-persistedstate';

// Vuexの設定
Vue.use(Vuex);

// storeの定義
export default new Vuex.Store<State>({ 
  plugins: [createPersistedState({
      key: 'dotArtStore',
      paths: ['canvasData'],
      storage: window.sessionStorage})]
  });
