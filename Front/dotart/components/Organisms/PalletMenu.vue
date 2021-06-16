<template>
    <transition name="palletMenu">
        <div v-show="palletDrawerFlag" class="drawerMenuArea__Wrapper">
            <!-- ここにメニューの内容を書いていく -->
            <div class="menuContents">
                <div class="scroll">
                    <pallet-area
                        class="palletArea"
                        :color-pallet="colorPallet"
                        :first-pallet-index="firstPalletIndex"
                        :palletIndex="palletIndex"
                        @getPalletColor="getPalletColorMethod"
                    ></pallet-area>
                </div>
                <button class="switch" @click="translate">X</button>
            </div>
        </div>
    </transition>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'nuxt-property-decorator';
import PalletArea from '@/components/Molecules/PalletArea.vue';

@Component({
    // middleware: 'auth',
    components: {
        PalletArea,
    },
})
export default class PalletMenu extends Vue {
    @Prop({ type: Array })
    colorPallet!: string[]; // ページから渡されるパレットの色の配列

    @Prop({ type: Number })
    firstPalletIndex!: number; // 最初に選択しているパレットの位置

    @Prop({ type: Boolean })
    palletDrawerFlag!: Boolean;

    @Prop({ type: Function })
    getPalletColor!: Function;

    @Prop({ type: Function })
    translate!: Function;

    palletIndex: number = this.firstPalletIndex;
    selectingColor: string = this.colorPallet[this.palletIndex];

    public created(): void {
        return;
    }

    public mounted(): void {}

    public getPalletColorMethod(newColor: string, newIndex: number): void {
        this.selectingColor = newColor;
        this.palletIndex = newIndex;
        console.log(this.palletIndex);
        this.getPalletColor(this.selectingColor, this.palletIndex);
    }
}
</script>
<style lang="scss" scoped>
$defaultHeight: 0px; //格納状態でのメニューのheight
$movedHeight: 50px; //展開状態でのメニューのheight
$movePercentage: 100% * (1 - $defaultHeight/$movedHeight); //transformの割合
//@debug $movePercentage;
.drawerdefault {
    display: flex;
    vertical-align: top;
    justify-content: space-between;
}
.menuContents {
    display: -webkit-box; //これ
    display: flex;
    margin-left: 5px;
    margin-right: 5px;
    width: 100%;
    max-width: 450px;
    vertical-align: top;
}
.scroll {
    overflow-x: scroll;
    max-width: 90%;
}
.palletArea {
    margin: auto;
}
.switch {
    vertical-align: top;
    margin-left: auto;
    margin-right: auto;
}
.palletMenu-enter-active,
.palletMenu-leave-active {
    transform: translate(0px, 0px);
    transition: transform 225ms cubic-bezier(0, 0, 0.2, 1) 0ms;
}
.palletMenu-enter,
.palletMenu-leave-to {
    transform: translateY($movePercentage) translateY($defaultHeight);
}

.drawerMenuArea__Wrapper {
    position: absolute;
    text-align: center;
    display: ruby;
    margin-left: auto;
    margin-right: auto;
    z-index: 2;
    right: 0;
    left: 0;
    bottom: 0;
    width: 100%;
    max-width: 100%;
    height: $movedHeight;
    background-color: rgb(114, 11, 92);
}
</style>