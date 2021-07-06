<template>
    <button @click="clickEvent">
        <icon-Base width="36" height="36" icon-name="redoIcon"
            ><redo-Icon
        /></icon-Base>
    </button>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api';
// import { canvasDataModule } from '../../store/modules/canvasData';
import IconBase from '@/components/Atomics/Icons/IconBase.vue';
import RedoIcon from '@/components/Atomics/Icons/RedoIcon.vue';

export default defineComponent({
    name: 'RedoButton',
    components: {
        IconBase,
        RedoIcon,
    },
    props: {
        clickEvent: {
            type: Function,
            required: true,
        },
    },
});
</script>
<style lang="scss" scoped>
/* */
</style>
