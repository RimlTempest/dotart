<template>
    <v-layout column justify-center align-center>
        <v-flex xs12 sm8 md6>
            <v-container fluid>
                <v-row dense>
                    <v-col
                        v-for="card in cardState.cards"
                        :key="card.title"
                        :cols="card.flex"
                    >
                        <v-card light hover max-width="1000" :to="card.link">
                            <!-- タイトル -->
                            <v-card-title
                                class="headline"
                                v-text="card.title"
                            ></v-card-title>
                            <!-- サブタイトル -->
                            <v-card-subtitle
                                v-text="card.subtitle"
                            ></v-card-subtitle>
                            <v-divider class="mx-4"></v-divider>
                            <v-card-text v-if="card.infomation != null">
                                <div>{{ card.infomation }}</div>
                            </v-card-text>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
        </v-flex>
    </v-layout>
</template>

<script lang="ts">
import { defineComponent, reactive } from '@nuxtjs/composition-api';

type Cards = {
    title: string;
    subtitle: string;
    link: string;
    infomation: string | null;
    flex: number;
};

export default defineComponent({
    name: 'index',
    setup() {
        const cardState = reactive<{ cards: Cards[] }>({
            cards: [
                // {
                //     title: 'ログイン',
                //     subtitle: 'ログイン、新規登録はこちらから出来ます。',
                //     link: '/login/',
                //     flex: 12,
                // },
                {
                    title: 'ドット絵クリエイター',
                    subtitle: 'ドット絵を作ってみよう！',
                    link: '/creator/',
                    infomation: null,
                    flex: 12,
                },
                // {
                //     title: '投稿ページ',
                //     subtitle: 'ドット絵の投稿ページ',
                //     link: '/sns/',
                //     flex: 6,
                // },
                {
                    title: 'お知らせ',
                    subtitle: 'version 0.0.5β',
                    link: '/',
                    infomation:
                        '現在テスト環境です。\nバグ等あるかもしれませんがご了承ください。',
                    flex: 12,
                },
            ],
        });
        return { cardState };
    },
});
</script>
