import Vue from 'vue';
import Vuex from 'vuex';
import { State } from '~/types/Store/StateType';

// Vuexの設定
Vue.use(Vuex);

// storeの定義
export default new Vuex.Store<State>({});
