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
import { Vue, Component, Prop, Watch } from 'nuxt-property-decorator';
// import { canvasDataModule } from '../../store/modules/canvasData';
@Component
export default class Pallet extends Vue {
    @Prop({ type: String })
    color!: string;

    @Prop({ type: Number })
    index!: number;

    @Prop({ type: Number })
    selectedIndex!: number;

    margin: string = '2px';
    width: string = '23px';
    height: string = '23px';
    border: string = '2px solid rgb(87, 56, 84)';

    selectedflg: boolean = false;

    public mounted(): void {
        if (this.selectedIndex === this.index) {
            this.margin = '1px';
            this.width = '25px';
            this.height = '25px';
            this.border = '3px solid rgb(235, 146, 227)';
        }
    }

    @Watch('selectedIndex')
    public checkSelected(newIndex: number, _oldIndex: number) {
        if (newIndex === this.index) {
            this.margin = '1px';
            this.width = '25px';
            this.height = '25px';
            this.border = '3px solid rgb(235, 146, 227)';
        } else {
            this.margin = '2px';
            this.width = '23px';
            this.height = '23px';
            this.border = '2px solid rgb(87, 56, 84)';
        }
    }

    public getColor(_e: any): void {
        this.$emit('getColor', this.color, this.index);
    }
}
</script>
<style lang="scss" scoped>
.pallet {
    display: inline-block;
    border-radius: 35%;
}
</style>
