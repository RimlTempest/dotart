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
                            <button @click="logout">ログアウト</button>
                        </v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </div>
        </v-list>
    </v-menu>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';

@Component({})
export default class LogoutMenu extends Vue {
    loginOfRegister: string =
        this.$store.state.auth.user.userID + 'でログイン中です。';

    logout(): void {
        (this as any).$auth.logout();
    }
}
</script>
