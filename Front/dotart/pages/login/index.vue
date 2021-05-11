<template>
    <v-layout column justify-center align-center>
        <!-- プログレスバー -->
        <v-overlay
            v-if="flaggerState.loadingFlg"
            :absolute="flaggerState.loadingFlg"
            :value="flaggerState.loadingFlg"
        >
            <v-progress-circular indeterminate color="primary" />
        </v-overlay>

        <div v-if="flaggerState.errorFlg">
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
                    <v-text-field v-model="userState.userId" label="ユーザ名" />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-text-field
                        v-model="userState.password"
                        type="password"
                        label="パスワード"
                    />
                    <v-spacer />
                </v-card-actions>

                <v-card-actions>
                    <v-spacer />
                    <!-- @click="loginWithAuthModule" -->
                    <v-btn color="primary" @click="loginWithAuthModule"
                        >ログイン</v-btn
                    >
                </v-card-actions>
            </v-form>
        </v-card>
        <!--<NuxtLink to="/">Home page</NuxtLink>-->
    </v-layout>
</template>

<script lang="ts">
import {
    defineComponent,
    reactive,
    useAsync,
    useContext,
} from '@nuxtjs/composition-api';

type User = {
    userId: string;
    password: string;
};
type Flagger = {
    loadingFlg: boolean;
    errorFlg: boolean;
};

export default defineComponent({
    name: 'signInPage',
    setup() {
        const userState = reactive<User>({
            userId: '',
            password: '',
        });
        const flaggerState = reactive<Flagger>({
            loadingFlg: false,
            errorFlg: false,
        });

        /* TODO: ログイン実装に用いているnuxt/authのcontextが取れているはずなのに $auth でエラーが吐かれる
         * Property '$auth' does not exist on type 'Store<any>'.
         */
        const loginWithAuthModule = useAsync(() => {
            const { store } = useContext();
            flaggerState.loadingFlg = true;
            // console.log(store.$auth);
            store.$auth
                .loginWith('local', {
                    data: {
                        userId: userState.userId,
                        password: userState.password,
                    },
                })
                .then(
                    (response: any) => {
                        console.log(`LOGIN_SUCCESS : `);
                        // console.log(response)
                        // store.state.nuxt.setLayout('default');
                        flaggerState.errorFlg = false;
                        flaggerState.loadingFlg = false;
                        return response;
                    },
                    (error: any) => {
                        console.log(userState.password);
                        console.log(`LOGIN_ERROR : ${error}`);
                        flaggerState.errorFlg = true;
                        flaggerState.loadingFlg = false;
                        return error;
                    }
                );
        });
        return { userState, flaggerState, loginWithAuthModule };
    },
});
</script>
