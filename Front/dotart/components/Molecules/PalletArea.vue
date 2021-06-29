<template>
    <v-row dense class="palletArea">
        <v-layout justify-center class="palletitem">
            <pallet
                class="palletitem"
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

    @Prop({ type: Number })
    palletIndex!: number; // 最初に選択しているパレットの位置

    getPalletColor(newColor: string, newIndex: number): void {
        this.$emit('getPalletColor', newColor, newIndex);
    }
}
</script>
<style lang="scss" scoped>
.palletArea {
    display: inline-block;
    white-space: nowrap;
    justify-content: space-between;
}
</style>
