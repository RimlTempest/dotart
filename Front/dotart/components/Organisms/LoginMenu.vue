<template>
    <v-menu offset-y>
        <template v-slot:activator="{ on }">
            <v-btn text outlined v-on="on">{{ loginOfRegister }}</v-btn>
        </template>
        <v-list>
            <!-- ログイン -->
            <div>
                <v-list-item>
                    <v-list-item-content>
                        <v-list-item-title>
                            <!-- 遷移 -->
                            <nuxt-link to="/login/"
                                >ログインの方はこちら</nuxt-link
                            >
                        </v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <!-- 新規登録 -->
                <v-list-item>
                    <v-list-item-content>
                        <!-- 遷移 -->
                        <nuxt-link to="/login/Signup"
                            >新規登録の方はこちら</nuxt-link
                        >
                    </v-list-item-content>
                </v-list-item>
            </div>
        </v-list>
    </v-menu>
</template>

<script lang="ts">
// TODO: 本実装する際にcompotisionApiにする
import { Component, Vue } from 'nuxt-property-decorator';

@Component({})
export default class LoginMenu extends Vue {
    loginOfRegister: string = 'ログイン / 新規登録';
}
</script>
