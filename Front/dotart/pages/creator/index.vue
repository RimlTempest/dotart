<template>
    <v-layout column justify-center align-center>
        <v-flex xs12 sm8 md6>
            <v-container fluid>
                <v-row dense>
                    <v-col cols="12">
                        <v-card
                            light
                            hover
                            max-width="1000"
                            class="justify-center"
                        >
                            <v-form ref="form" lazy-validation>
                                <v-card-title class="justify-center"
                                    >ドット絵作成</v-card-title
                                >
                            </v-form>
                        </v-card>
                    </v-col>

                    <v-col cols="12">
                        <v-card
                            light
                            hover
                            max-width="1000"
                            class="justify-center"
                        >
                            <v-form ref="form" lazy-validation>
                                <v-card-title class="headline"
                                    >キャンバス名</v-card-title
                                >
                                <v-card-subtitle
                                    >保存するときの名前です</v-card-subtitle
                                >
                                <v-text-field
                                    v-model="canvasName"
                                    label="キャンバス名"
                                />
                            </v-form>
                        </v-card>
                    </v-col>

                    <v-col cols="12">
                        <v-card
                            light
                            hover
                            max-width="1000"
                            class="justify-center"
                        >
                            <v-form ref="form" lazy-validation>
                                <v-card-title class="headline"
                                    >パレット選択</v-card-title
                                >
                                <v-card-subtitle
                                    >使うパレットを選びましょう</v-card-subtitle
                                >
                                <v-select
                                    v-model="selectedPallet"
                                    :items="palletListItems"
                                    label="パレット"
                                    return-object
                                ></v-select>
                                <small>{{ selectedPallet['summary'] }}</small>
                                <div class="palletPreview">
                                    <div
                                        v-for="item in selectedPallet['pallet']"
                                        :key="item"
                                        class="palletColor"
                                        :style="{ background: item }"
                                    ></div>
                                </div>
                            </v-form>
                        </v-card>
                    </v-col>

                    <v-col cols="12">
                        <v-card
                            light
                            hover
                            max-width="1000"
                            class="justify-center"
                        >
                            <v-form ref="form" lazy-validation>
                                <v-card-title class="headline"
                                    >キャンバスサイズ指定</v-card-title
                                >
                                <v-card-subtitle
                                    >キャンバスサイズを決めましょう</v-card-subtitle
                                >
                                <v-select
                                    v-model="selectedSize"
                                    :items="sizeListItems"
                                    label="キャンバスサイズ"
                                    return-object
                                ></v-select>
                                <v-card-actions>
                                    <v-spacer />
                                    <v-btn color="primary" @click="startDraw"
                                        >start</v-btn
                                    >
                                </v-card-actions>
                            </v-form>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
        </v-flex>
    </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator';
import { CanvasDataModule } from '@/store/modules/canvasData';

type SelectedSize = {
    text: string;
    range: number;
    magnification: number;
};

type SizeListItems = {
    text: string;
    range: number;
    magnification: number;
};

type SelectedPallet = {
    text: string;
    summary: string;
    pallet: string[];
};

@Component({
    // middleware: 'auth'
})
// TypeScriptの処理
export default class CreatorPage extends Vue {
    canvasName: string = 'キャンバス';
    // キャンパスサイズ
    selectedSize: SelectedSize = {
        text: '16×16',
        range: 16,
        magnification: 24,
    };

    sizeListItems: SizeListItems[] = [
        { text: '16×16', range: 16, magnification: 24 },
        { text: '32×32', range: 32, magnification: 12 },
        { text: '48×48', range: 48, magnification: 8 },
        { text: '64×64', range: 64, magnification: 6 },
        { text: '96×96', range: 96, magnification: 4 },
    ];

    palletListItems: SelectedPallet[] = [
        {
            text: 'スタンダード',
            summary: '使いやすそうな色をまとめてみました。',
            pallet: [
                'rgb(255, 255, 255)',
                'rgb(125, 125, 125)',
                'rgb(0, 0, 0)',
                'rgb(108, 57, 0)',
                'rgb(243, 55, 55)',
                'rgb(212, 110, 229)',
                'rgb(180, 27, 235)',
                'rgb(189, 137, 207)',
                'rgb(150, 150, 215)',
                'rgb(90, 90, 180)',
                'rgb(82, 226, 226)',
                'rgb(137, 255, 146)',
                'rgb(199, 243, 118)',
                'rgb(255, 245, 70)',
                'rgb(255, 195, 100)',
                'rgb(255, 228, 175)',
            ],
        },
        {
            text: 'モノクロ',
            summary: '白黒のグラデーションです。',
            pallet: [
                'rgb(255, 255, 255)',
                'rgb(235, 235, 235)',
                'rgb(220, 220, 220)',
                'rgb(205, 205, 205)',
                'rgb(190, 190, 190)',
                'rgb(175, 175, 175)',
                'rgb(160, 160, 160)',
                'rgb(145, 145, 145)',
                'rgb(130, 130, 130)',
                'rgb(115, 115, 115)',
                'rgb(100, 100, 100)',
                'rgb(85, 85, 85)',
                'rgb(70, 70, 70)',
                'rgb(55, 55, 55)',
                'rgb(40, 40, 40)',
                'rgb(0, 0, 0)',
            ],
        },
        {
            text: 'レトロゲーム4色',
            summary: '当時の4色ディスプレイ風味。',
            pallet: [
                'rgb(238, 238, 162)',
                'rgb(238, 238, 162)',
                'rgb(238, 238, 162)',
                'rgb(238, 238, 162)',
                'rgb(186, 200, 112)',
                'rgb(186, 200, 112)',
                'rgb(186, 200, 112)',
                'rgb(186, 200, 112)',
                'rgb(107, 131, 56)',
                'rgb(107, 131, 56)',
                'rgb(107, 131, 56)',
                'rgb(107, 131, 56)',
                'rgb(38, 55, 0)',
                'rgb(38, 55, 0)',
                'rgb(38, 55, 0)',
                'rgb(38, 55, 0)',
            ],
        },
    ];

    // 初期値
    selectedPallet: SelectedPallet = this.palletListItems[0];

    startDraw(): void {
        // お絵描きページへ遷移
        CanvasDataModule.setCanvasName(this.canvasName);
        CanvasDataModule.setCanvasRange(this.selectedSize.range);
        CanvasDataModule.setCanvasMagnification(
            this.selectedSize.magnification
        );
        CanvasDataModule.setPalletName(this.selectedPallet.text);
        CanvasDataModule.setPalletColor(this.selectedPallet.pallet);
        this.$router.push('/creator/canvas');
    }
}
</script>

<style>
.palletPreview {
    display: flex;
    pointer-events: none;
}
.palletColor {
    display: flex;
    width: 20px;
    height: 20px;
    border: 2px solid rgb(50, 50, 50);
    pointer-events: none;
    border-radius: 30%;
}
</style>
