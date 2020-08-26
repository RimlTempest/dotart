<template>
<v-carousel height="650" hide-delimiter-background show-arrows-on-hover>
    <v-carousel-item v-for="(slide, i) in slides" :key="i">
        <v-sheet :color="colors[i]" height="100%">
            <v-row class="fill-height" align="center" justify="center">
                <v-card class="cardColor" height="600" width="1000" outlined>
                    <v-list-item three-line>
                        <v-list-item-content>
                            <v-list-item-title class="headline">
                                <p>{{ slide }}</p>
                            </v-list-item-title>
                            <div class="sl1" v-if="i == 0">
                                <p>フロントエンド</p>
                                <p>・Riml</p>
                                <p>・ZIN</p>
                                <br />
                                <p>バックエンド</p>
                                <p>・Hash</p>
                                <p>・kimu</p>
                            </div>
                            <div v-if="i == 1">
                                <br />
                                <p>自分たちは将来的に役に立つだろう環境での開発を心がけて環境を構築しました。</p>
                                <p>等アプリケーションはWeb、モバイルに対応したドット絵をかんたんに作成できるクリエイターアプリです。</p>
                                <p>ドット絵が何かしらない人は調べてください。検索エンジンはGoogleがおすすめです。</p>
                                <p>完成率はまだ4割ほどですがフロント部分はある程度目処がたってきました。</p>
                                <p>バックエンドに関しては正直一切進んでません。</p>
                            </div>
                            <div v-if="i == 2">
                                <br />
                                <p>Frontend</p>
                                <p>lang: TypeScript</p>
                                <p>・Nuxtjs / Vuejs</p>
                                <br />
                                <p>Backend</p>
                                <p>lang: Golang</p>
                                <p>・gin / Gorm</p>
                                <p></p>
                                <br />
                                <p>Infra</p>
                                <p>・Docker</p>
                                <br />
                                <p>App</p>
                                <p>・MySQL / phpmyadmin</p>
                            </div>
                            <div class="sl3" v-if="i == 3">
                                <p>実機説明（もう動かしている）</p>
                            </div>
                        </v-list-item-content>
                    </v-list-item>
                </v-card>
            </v-row>
        </v-sheet>
    </v-carousel-item>
</v-carousel>
</template>

<script lang="ts">
import {
    Component,
    Vue
} from 'vue-property-decorator'
import {
    counterModule
} from '../../store/modules/counter'

@Component({
    //middleware: 'auth'
})

// vuexのサンプルのコード
export default class Counter extends Vue {
    colors = ['indigo', 'warning', 'pink darken-2', 'red lighten-1']
    slides = ['ドット絵アプリ　DotArt', 'アプリ概要', '技術', '']
}
</script>

<style>
.cardColor {
    background-color: rgba(255, 255, 255, 0) !important;
    border-color: rgba(255, 255, 255, 0) !important;
}

.sl1 {
    position: absolute;
    top: 300px;
    right: 0%;
}

.sl3 {
    position: absolute;
    text-align: center;
    top: 290px;
    right: 0;
    left: 0;
}
</style>
