<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      <v-container fluid>
        <v-row dense>
          <v-col v-for="card in cards" :key="card.title" :cols="card.flex">
            <v-card light hover max-width="1000" :to="card.link">
              <!-- タイトル -->
              <v-card-title class="headline" v-text="card.title"></v-card-title>
              <!-- サブタイトル -->
              <v-card-subtitle v-text="card.subtitle"></v-card-subtitle>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

type Cards = {
  title: string
  subtitle: string
  link: string
  flex: number
}

@Component({})
export default class IndexPage extends Vue {
  cards: Cards[] = [
    {
      title: 'ログイン',
      subtitle: 'ログイン、新規登録はこちらから出来ます。',
      link: '/login/',
      flex: 12
    },
    {
      title: 'ドット絵クリエイター',
      subtitle: 'ドット絵を作ってみよう！',
      link: '/creator/',
      flex: 12
    },
    {
      title: '投稿ページ',
      subtitle: 'ドット絵の投稿ページ',
      link: '/sns/',
      flex: 6
    },
    {
      title: 'お知らせ',
      subtitle: 'version 0.0.2β',
      link: '/infomation',
      flex: 6
    }
  ]
}
</script>
