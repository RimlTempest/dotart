import Vue from 'vue'
import Vuex from 'vuex'
import { CounterState } from '@/store/modules/counter'
import { RegisterState } from '@/store/modules/register'
import { CanvasDataState } from '@/store/modules/canvasData'

// Vuexの設定
Vue.use(Vuex)

// stateの型を縛る
export interface State {
  counter: CounterState
  register: RegisterState
  canvasData: CanvasDataState
}

// storeの定義
export default new Vuex.Store<State>({})
