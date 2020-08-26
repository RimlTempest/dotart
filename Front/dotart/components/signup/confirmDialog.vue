<template>
  <v-row justify="center">
    <v-dialog
      v-model="dialog"
      fullscreen
      hide-overlay
      transition="dialog-bottom-transition"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn color="primary" dark v-bind="attrs" v-on="on">新規作成</v-btn>
      </template>
      <!-- プログレスバー -->
      <v-overlay :absolute="absolute" :value="overlay" v-if="loadingflg">
        <v-progress-circular indeterminate color="primary" />
      </v-overlay>

      <v-card>
        <v-toolbar dark color="primary">
          <v-btn icon dark @click="close">
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>登録アカウント情報</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>

        <v-list three-line subheader>
          <v-list-item v-if="errorflg">
            <v-list-item-content>
              <v-alert dense outlined type="error"
                >ユーザID、パスワード、メールアドレスいずれかが条件を満たしていません。</v-alert
              >
            </v-list-item-content>
          </v-list-item>
          <!-- UserId -->
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title>UserId</v-list-item-title>
              <v-list-item-subtitle>{{ UserId }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <!-- パスワード -->
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title>Password</v-list-item-title>
              <v-list-item-subtitle>{{ Password }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <!-- メールアドレス -->
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title>MailAddress</v-list-item-title>
              <v-list-item-subtitle>{{ Mail }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <v-divider></v-divider>
          <!-- 登録ボタン -->

          <v-list-item>
            <v-list-item-content>
              <p class="text-right">
                以上の内容で登録してもよろしいでしょうか？
              </p>
              <v-btn tile large color="primary" icon @click="signup"
                >登録する</v-btn
              >
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import { registerModule } from '../../store/modules/register'
import axios from 'axios'

@Component({})
export default class ConfirmDialog extends Vue {
  dialog: boolean = false
  // プログレスバー
  loadingflg: boolean = false
  // エラーフラグ
  errorflg: boolean = false

  @Prop(String)
  UserId: string = ''
  @Prop(String)
  Password: string = ''
  @Prop(String)
  Mail: string = ''

  // storeを呼び出し値のセット
  setDatas(): void {
    registerModule.userIdAction(this.UserId)
    registerModule.passwordAction(this.Password)
    registerModule.mailAction(this.Mail)
    console.log(`USER_ID : ${registerModule.userId}`)
    console.log(`PASSWORD : ${registerModule.password}`)
    console.log(`MAIL : ${registerModule.mail}`)
  }

  // 新規作成処理
  signup(): void {
    this.setDatas()
    if (
      registerModule.userId === 'error' ||
      registerModule.password === 'error' ||
      registerModule.mail === 'error'
    ) {
      // 何かしらエラーあるのでエラーアラートを出す
      this.errorflg = true
      console.log('error')
    } else {
      // エラーフラグを閉じる
      this.errorflg = false

      this.asyncData()
      alert('成功しました。')
      this.$router.push('/login')
    }
  }

  // 登録処理
  async asyncData() {
    this.loadingflg = true
    const res = await axios
      .post(
        '/api/v1/user',
        {
          UserId: registerModule.userId,
          Password: registerModule.password,
          Mail: registerModule.mail
        },
        {
          headers: {
            'Content-Type': 'application/json'
          }
        }
      )
      .then((res: any) => {
        console.log('SUCCESS')
        this.loadingflg = false
        console.log(res)
      })
      .catch((e: any) => {
        this.loadingflg = false
        console.log('ERROR')
        console.log(e)
      })

    console.log(res)
  }

  close(): void {
    this.dialog = false
    this.errorflg = false
  }
}
</script>
