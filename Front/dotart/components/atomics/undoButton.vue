<template>
  <button v-on:click="clickEvent">
    <icon-Base width="36" height="36" icon-name="undoIcon"
      ><undoIcon
    /></icon-Base>
  </button>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { canvasDataModule } from '../../store/modules/canvasData'
import iconBase from '@/components/atomics/icons/iconBase.vue'
import undoIcon from '@/components/atomics/icons/undoIcon.vue'

@Component({
  middleware: 'auth',
  components: {
    iconBase,
    undoIcon
  }
})
export default class undoButton extends Vue {
  @Prop({ type: Function })
  clickEvent!: Function
}
</script>
<style>
</style>