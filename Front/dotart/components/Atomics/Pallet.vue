<template>
    <div
        class="pallet"
        :style="{
            background: color,
            margin: margin,
            border: border,
            width: width,
            height: height,
        }"
        @mousedown="getColor"
    ></div>
</template>

<script lang="ts">
import {
    defineComponent,
    reactive,
    onMounted,
    watch,
} from '@nuxtjs/composition-api';

export default defineComponent({
    name: 'Pallet',
    props: {
        color: {
            type: String,
            required: true,
        },
        index: {
            type: Number,
            default: 0,
            required: true,
        },
        selectedIndex: {
            type: Number,
            default: 0,
            required: true,
        },
    },
    setup({ color, index, selectedIndex }, context) {
        const styleState = reactive<{
            margin: string;
            width: string;
            height: string;
            border: string;
        }>({
            margin: '2px',
            width: '23px',
            height: '23px',
            border: '2px solid rgb(87, 56, 84)',
        });

        onMounted(() => {
            if (selectedIndex === index) {
                styleState.margin = '1px';
                styleState.width = '25px';
                styleState.height = '25px';
                styleState.border = '3px solid rgb(235, 146, 227)';
            }
        });

        watch(
            () => selectedIndex,
            (newIndex: number, _oldIndex: number) => {
                if (newIndex === index) {
                    styleState.margin = '1px';
                    styleState.width = '25px';
                    styleState.height = '25px';
                    styleState.border = '3px solid rgb(235, 146, 227)';
                } else {
                    styleState.margin = '2px';
                    styleState.width = '23px';
                    styleState.height = '23px';
                    styleState.border = '2px solid rgb(87, 56, 84)';
                }
            }
        );

        const getColor = (_e: any): void => {
            context.emit('getColor', color, index);
        };

        return {
            getColor,
        };
    },
});
</script>
<style lang="scss" scoped>
.pallet {
    display: inline-block;
    border-radius: 35%;
}
</style>
