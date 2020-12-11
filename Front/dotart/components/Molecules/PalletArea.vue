<template>
    <v-row dense class="palletarea">
        <v-layout justify-center>
            <pallet
                v-for="(item, index) in colorPallet"
                :key="item"
                :color="item"
                :index="index"
                :selected-index="palletIndex"
                @getColor="getPalletColor"
            ></pallet>
        </v-layout>
    </v-row>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'nuxt-property-decorator';
// import { canvasDataModule } from '../../store/modules/canvasData';
import Pallet from '@/components/Atomics/Pallet.vue';

@Component({
    components: {
        Pallet,
    },
})
export default class PalletArea extends Vue {
    @Prop({ type: Array })
    colorPallet!: string[]; // ページから渡されるパレットの色の配列

    @Prop({ type: Number })
    firstPalletIndex!: number; // 最初に選択しているパレットの位置

    palletIndex: number = this.firstPalletIndex;
    selectingColor: string = this.colorPallet[this.palletIndex];

    getPalletColor(newColor: string, newIndex: number): void {
        this.selectingColor = newColor;
        this.palletIndex = newIndex;
        this.$emit('getPalletColor', this.selectingColor, this.palletIndex);
    }
}
</script>
<style lang="scss" scoped>
.palletarea {
    display: flex;
    justify-content: space-between;
}
</style>
