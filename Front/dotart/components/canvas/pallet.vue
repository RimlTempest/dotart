<template>
  <div
    class="pallet"
    v-bind:style="{ background: color }"
    @mousedown="getcolor"
  ></div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { canvasDataModule } from '../../store/modules/canvasData'
@Component
export default class pallet extends Vue {
  @Prop({ type: String })
  color!: String

  selectedflg: boolean = false
  getcolor(e: any): void {
    this.$emit('getcolor', this.color)
  }
}
</script>
<style>
.pallet {
  display: flex;
  width: 23px;
  height: 40px;
  border: 2px solid rgb(235, 146, 227);
}
</style>
