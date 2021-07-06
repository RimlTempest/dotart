<template>
    <svg
        xmlns="http://www.w3.org/2000/svg"
        :width="width"
        :height="height"
        viewBox="0 0 18 18"
        :aria-labelledby="iconName"
        role="presentation"
    >
        <title :id="iconName" lang="en">{{ iconName }} icon</title>
        <g :fill="iconColor" :stroke="iconColor">
            <slot />
        </g>
    </svg>
</template>
<script lang="ts">
import { defineComponent, onMounted } from '@nuxtjs/composition-api';

export default defineComponent({
    name: 'IconBase',
    components: {},
    props: {
        iconName: {
            type: String,
            required: true,
            default: '',
        },
        iconColor: {
            type: String,
            required: true,
            default: 'currentColor',
        },
        width: {
            type: String,
            required: true,
            default: '32',
        },
        height: {
            type: String,
            required: true,
            default: '32',
        },
    },
    setup() {
        onMounted(() => {
            console.log('IconBase');
        });
    },
});
</script>
