<template>
    <div class="mainMenu">
        <div v-if="!drawerFlg" class="drawerMenuArea">
            <div class="drawerdefault">
                <div
                    class="scroll"
                    id="beforeScrollArea"
                    v-on:scroll="onScroll"
                >
                    <pallet-area
                        class="palletArea"
                        :colorPallet="colorPallet"
                        :firstPalletIndex="firstPalletIndex"
                        @getPalletColor="getPalletColor"
                    ></pallet-area>
                </div>
                <button class="switch" v-on:click="translate">▲</button>
            </div>
        </div>

        <transition name="popupMenu">
            <div v-if="drawerFlg" class="drawerMenuArea__Wrapper">
                <div class="drawerMenu">
                    <!-- ここにメニューの内容を書いていく -->
                    <div class="drawerdefault">
                        <div class="scroll" id="afterScrollArea">
                            <pallet-area
                                class="palletArea"
                                :colorPallet="colorPallet"
                                :firstPalletIndex="firstPalletIndex"
                                @getPalletColor="getPalletColor"
                            ></pallet-area>
                        </div>
                        <button class="switch" v-on:click="translate">▼</button>
                    </div>
                </div>
            </div>
        </transition>
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'nuxt-property-decorator';
import PalletArea from '@/components/Molecules/PalletArea.vue';

@Component({
    //middleware: 'auth',
    components: {
        PalletArea,
    },
})
export default class MainMenu extends Vue {
    @Prop({ type: Array })
    colorPallet!: string[]; //ページから渡されるパレットの色の配列
    @Prop({ type: Number })
    firstPalletIndex!: number; //最初に選択しているパレットの位置
    @Prop({ type: Function })
    getPalletColor!: Function;

    scrollArea: HTMLDivElement | null = null;
    drawerFlg: boolean = false;

    public mounted(): void {
        this.scrollArea = document.querySelector('#beforeScrollArea');
        console.log('this.scrollArea');
        this.scrollArea!.addEventListener('scroll', this.onScroll);
        return;
    }

    public translate(): void {
        this.drawerFlg = !this.drawerFlg;
        return;
    }

    public onScroll(e: any): void {
        console.log(e);
        return;
    }
}
</script>
<style lang="scss" scoped>
$defaultHeight: 45px; //格納状態でのメニューのheight
$movedHeight: 230px; //展開状態でのメニューのheight
$movePercentage: 100% * (1 - $defaultHeight/$movedHeight); //transformの割合
//@debug $movePercentage;
.mainMenu {
    display: grid;
}
.drawerdefault {
    display: ruby;
    vertical-align: top;
    justify-content: space-between;
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
    margin-top: 6px;
}
.popupMenu-enter-active,
.popupMenu-leave-active {
    transform: translate(0px, 0px);
    transition: transform 225ms cubic-bezier(0, 0, 0.2, 1) 0ms;
}
.popupMenu-enter,
.popupMenu-leave-to {
    transform: translateY($movePercentage);
}

.drawerMenuArea {
    position: absolute;
    text-align: center;
    margin-left: auto;
    margin-right: auto;
    z-index: 1;
    right: 0;
    left: 0;
    bottom: 0;
    width: 100%;
    height: $defaultHeight;
    background-color: rgba(233, 95, 192);
}
.drawerMenuArea__Wrapper {
    position: absolute;
    text-align: center;
    margin-left: auto;
    margin-right: auto;
    z-index: 2;
    right: 0;
    left: 0;
    bottom: 0;
    width: 100%;
    height: $movedHeight;
    background-color: rgba(233, 95, 192);
}
</style>