import {
  Mutation,
  MutationAction,
  Action,
  VuexModule,
  getModule,
  Module
} from 'vuex-module-decorators'
import store from '@/store/store'

// 結果を返す型
export interface CounterState {
  incrementCounter: number
  decrementCounter: number
}

@Module({ dynamic: true, store, name: 'counter', namespaced: true })
class Counter extends VuexModule implements CounterState {
  // state
  incrementCounter: number = 0
  decrementCounter: number = 1000

  // カウンタに値をセットする mutation
  @Mutation
  public SetIncrementCounter(num: number) {
    this.incrementCounter = num
  }
  @Mutation
  public setDecrementCounter(num: number) {
    this.decrementCounter = num
  }

  // stateに向けての値の処理
  @Action({})
  public incrementHundred() {
    this.SetIncrementCounter(this.incrementCounter + 100)
  }
  @Action({})
  public decrementHundred() {
    this.setDecrementCounter(this.decrementCounter - 100)
  }

  // Mutation,Action両方を実行する 型推論がうまく働かなくなるためあんまり使わないほうがいいけど例なので一応残しとく
  @MutationAction({ mutate: ['incrementCounter', 'decrementCounter'] })
  async resetCounter() {
    return {
      incrementCounter: 0,
      decrementCounter: 1000
    }
  }
}

// モジュール化
export const counterModule = getModule(Counter)
