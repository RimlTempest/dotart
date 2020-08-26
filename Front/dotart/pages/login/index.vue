<template>
  <v-layout column justify-center align-center>
    <!-- プログレスバー -->
    <v-overlay :absolute="absolute" :value="overlay" v-if="flag.loadingFlg">
      <v-progress-circular indeterminate color="primary" />
    </v-overlay>

    <div v-if="flag.errorFlg">
      <v-alert dense outlined type="error"
        >ユーザID、パスワード、メールアドレスいずれかが条件を満たしていません。</v-alert
      >
    </div>
    <v-card>
      <!-- v-model="valid" -->
      <v-form ref="form" lazy-validation>
        <v-card-title class="headline">ログイン</v-card-title>
        <v-card-text>
          <p>ユーザー名とパスワードを入力してください。</p>
        </v-card-text>

        <v-card-actions>
          <v-text-field label="ユーザ名" v-model="user.userId" />
          <v-spacer />
        </v-card-actions>

        <v-card-actions>
          <v-text-field
            type="password"
            label="パスワード"
            v-model="user.password"
          />
          <v-spacer />
        </v-card-actions>

        <v-card-actions>
          <v-spacer />
          <!-- @click="loginWithAuthModule" -->
          <v-btn color="primary" @click="loginWithAuthModule">ログイン</v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
    <!--<NuxtLink to="/">Home page</NuxtLink>-->
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

type User = {
  userId: string
  password: string
}
type Flagger = {
  loadingFlg: boolean
  errorFlg: boolean
}

// Componentの読み込み
@Component({})

// TypeScriptの処理
export default class SigninPage extends Vue {
  user: User = {
    userId: '',
    password: ''
  }
  flag: Flagger = {
    loadingFlg: false,
    errorFlg: false
  }
  //UserId: User = ''
  //Password: User = ''
  // プログレスバー
  //loadingflg: boolean = false
  //errorflg: boolean = false

  // methods
  async loginWithAuthModule() {
    this.flag.loadingFlg = true
    await this.$auth
      .loginWith('local', {
        data: {
          userId: this.user.userId,
          password: this.user.password
        }
      })
      .then(
        (response: any) => {
          console.log(`LOGIN_SUCCESS : `)
          //console.log(response)
          this.$nuxt.setLayout('default')
          this.flag.errorFlg = false
          this.flag.loadingFlg = false
          return response
        },
        (error: any) => {
          console.log(this.user.password)
          console.log(`LOGIN_ERROR : ${error}`)
          this.flag.errorFlg = true
          this.flag.loadingFlg = false
          return error
        }
      )
  }
}
</script>
