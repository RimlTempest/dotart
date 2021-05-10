<template>
    <v-layout column justify-center align-center>
        <v-card>
            <!-- v-model="valid" -->
            <v-form ref="form" lazy-validation>
                <v-card-title class="headline">新規作成</v-card-title>

                <v-card-actions>
                    <v-text-field v-model="userState.userId" label="ユーザ名" />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-text-field
                        v-model="userState.password"
                        label="パスワード"
                        :append-icon="passShowFlg ? 'mdi-eye' : 'mdi-eye-off'"
                        :type="passState.passShowFlg ? 'text' : 'password'"
                        hint="5文字以上入力してください。"
                        class="input-group--focused"
                        @click:append="
                            passState.passShowFlg = !passState.passShowFlg
                        "
                    />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-text-field
                        v-model="userState.mail"
                        label="メールアドレス"
                        :rules="[mailState.required, mailState.email]"
                    />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-spacer />
                    <!-- @click="loginWithAuthModule" -->
                    <ConfirmDialog
                        :user-id="userState.userId"
                        :password="userState.password"
                        :mail="userState.mail"
                    />
                </v-card-actions>
            </v-form>
        </v-card>
        <!--<NuxtLink to="/">Home page</NuxtLink>-->
    </v-layout>
</template>

<script lang="ts">
import { defineComponent, reactive } from '@nuxtjs/composition-api';
import { User } from '@/types/User/UserType';
import ConfirmDialog from '@/components/signup/confirmDialog.vue';

type Rules = {
    required: any;
    counter: any;
    email: any;
};

export default defineComponent({
    name: 'signUpPage',
    components: {
        ConfirmDialog,
    },
    setup() {
        const userState = reactive<User>({
            userId: '',
            password: '',
            mail: '',
        });
        const passState = reactive<{ passShowFlg: boolean }>({
            // パスワード表示非表示フラグ
            passShowFlg: false,
        });
        const mailState = reactive<Rules>({
            required: (value: any) => !!value || 'Required.',
            counter: (value: any) =>
                value.length <= 20 || '20文字以上は入力できません。',
            email: (value: any) => {
                const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                return (
                    pattern.test(value) ||
                    'メールアドレスの形式になっていません。'
                );
            },
        });
        return { userState, passState, mailState };
    },
});
</script>
