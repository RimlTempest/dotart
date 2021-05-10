<template>
    <v-row dense>
        <v-col cols="2">
            <div>
                <button @click="modeChange">{{ penMode }}</button>
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <undoButton :click-event="undo" />
            </div>
        </v-col>
        <v-col cols="2">
            <div>
                <redoButton :click-event="redo" />
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
                <button @click="anticlockRotate">反時計回り</button>
            </div>
        </v-col>
    </v-row>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';
import undoButton from '@/components/Atomics/undoButton.vue';
import redoButton from '@/components/Atomics/redoButton.vue';

@Component({
    middleware: 'auth',
    components: {
        undoButton,
        redoButton,
    },
})
export default class ButtonArea extends Vue {
    @Prop({ type: String })
    penMode!: string; // 最初に選択しているパレットの位置

    @Prop({ type: Function })
    modeChange!: Function;

    @Prop({ type: Function })
    undo!: Function;

    @Prop({ type: Function })
    redo!: Function;

    @Prop({ type: Function })
    drawGrid!: Function;

    @Prop({ type: Function })
    Save!: Function;

    @Prop({ type: Function })
    clockRotate!: Function;

    @Prop({ type: Function })
    anticlockRotate!: Function;
}
</script>
<style></style>
