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
import { defineComponent, PropType } from '@nuxtjs/composition-api';
// import { canvasDataModule } from '../../store/modules/canvasData';
import Pallet from '@/components/Atomics/Pallet.vue';

export default defineComponent({
    name: 'PalletArea',
    components: {
        Pallet,
    },
    props: {
        // ページから渡されるパレットの色の配列
        colorPallet: {
            type: Array as PropType<string[]>,
            required: true,
        },
        // 最初に選択しているパレットの位置
        firstPalletIndex: {
            type: Number,
            required: true,
        },
        palletIndex: {
            type: Number,
            requied: true,
        },
    },
    setup(_props, context) {
        const getPalletColor = (newColor: string, newIndex: number) => {
            context.emit('getPalletColor', newColor, newIndex);
        };
        return {
            getPalletColor,
        };
    },
});
</script>
<style lang="scss" scoped>
.palletarea {
    display: inline-block;
    white-space: nowrap;
    justify-content: space-between;
}
</style>
