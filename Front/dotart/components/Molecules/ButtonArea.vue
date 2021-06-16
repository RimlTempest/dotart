<template>
    <v-row dense>
        <v-col cols="2">
            <div>
                <button @click="modeChange">{{ penMode }}</button>
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <undo-button :click-event="undo" />
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <redo-button :click-event="redo" />
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <button @click="drawGrid">grid</button>
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <button @click="Save">保存</button>
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <button @click="clockRotate">時計回り</button>
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <button @click="antiClockRotate">反時計回り</button>
            </div>
        </v-col>
    </v-row>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api';
import UndoButton from '@/components/Atomics/UndoButton.vue';
import RedoButton from '@/components/Atomics/RedoButton.vue';

export default defineComponent({
    name: 'ButtonArea',
    middleware: 'auth',
    components: {
        UndoButton,
        RedoButton,
    },
    props: {
        penMode: {
            type: String,
            required: true,
        },
        modeChange: {
            type: Function,
            required: true,
        },
        undo: {
            type: Function,
            required: true,
        },
        redo: {
            type: Function,
            required: true,
        },
        drawGrid: {
            type: Function,
            required: true,
        },
        save: {
            type: Function,
            required: true,
        },
        clockRotate: {
            type: Function,
            required: true,
        },
        antiClockRotate: {
            type: Function,
            required: true,
        },
    },
});
</script>
<style></style>
