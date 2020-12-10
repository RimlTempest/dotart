<template>
    <v-layout column justify-center align-center>
        <v-card>
            <!-- v-model="valid" -->
            <v-form ref="form" lazy-validation>
                <v-card-title class="headline">新規作成</v-card-title>

                <v-card-actions>
                    <v-text-field v-model="user.userId" label="ユーザ名" />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-text-field
                        v-model="user.password"
                        label="パスワード"
                        :append-icon="passShowFlg ? 'mdi-eye' : 'mdi-eye-off'"
                        :type="passShowFlg ? 'text' : 'password'"
                        hint="5文字以上入力してください。"
                        class="input-group--focused"
                        @click:append="passShowFlg = !passShowFlg"
                    />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-text-field
                        v-model="user.mail"
                        label="メールアドレス"
                        :rules="[rules.required, rules.email]"
                    />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-spacer />
                    <!-- @click="loginWithAuthModule" -->
                    <ConfirmDialog
                        :user-id="user.userId"
                        :password="user.password"
                        :mail="user.mail"
                    />
                </v-card-actions>
            </v-form>
        </v-card>
        <!--<NuxtLink to="/">Home page</NuxtLink>-->
    </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator';
import { User } from '@/types/User/User';
import ConfirmDialog from '@/components/Signup/ConfirmDialog.vue';

type Rules = {
    required: any;
    counter: any;
    email: any;
};

// Componentの読み込み
@Component({
    components: {
        ConfirmDialog,
    },
})

// TypeScriptの処理
export default class SignupPage extends Vue {
    // 入力情報
    user: User = {
        userId: '',
        password: '',
        mail: '',
    };

    // パスワード表示非表示フラグ
    passShowFlg: boolean = false;
    // メールアドレスのバリデート
    rules: Rules = {
        required: (value: any) => !!value || 'Required.',
        counter: (value: any) =>
            value.length <= 20 || '20文字以上は入力できません。',
        email: (value: any) => {
            const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            return (
                pattern.test(value) || 'メールアドレスの形式になっていません。'
            );
        },
    };
}
</script>
