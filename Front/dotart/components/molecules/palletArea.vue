<template>
  <v-row dense class="palletarea">
    <v-layout justify-center>
      <pallet
        v-for="(item, index) in colorPallet"
        :key="item"
        :color="item"
        :index="index"
        :selectedIndex="palletIndex"
        @getcolor="getpalletcolor"
      ></pallet>
    </v-layout>
  </v-row>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { canvasDataModule } from '../../store/modules/canvasData'
import pallet from '@/components/atomics/pallet.vue'

@Component({
  middleware: 'auth',
  components: {
    pallet
  }
})
export default class palletArea extends Vue {
  @Prop({ type: Array })
  colorPallet!: string[] //ページから渡されるパレットの色の配列
  @Prop({ type: Number })
  firstPalletIndex!: number //最初に選択しているパレットの位置

  palletIndex: number = this.firstPalletIndex | 1
  selectingColor: string = this.colorPallet[this.palletIndex]

  getpalletcolor(newColor: string, newIndex: number): void {
    this.selectingColor = newColor
    this.palletIndex = newIndex
    this.$emit('getpalletcolor', this.selectingColor, this.palletIndex)
  }
}
</script>
<style>
.palletarea {
  display: flex;
  justify-content: space-between;
}
</style>