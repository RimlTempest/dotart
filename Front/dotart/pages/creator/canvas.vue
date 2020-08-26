<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm12 md12>
      <v-container fluid>
        <v-row dense>
          <v-col cols="12">
            <div class="DrowCanvas">
              <canvas id="drowcanvas" width="384px" height="384px"></canvas>
              <div class="GridCanvas">
                <canvas
                  id="gridcanvas"
                  width="384px"
                  height="384px"
                  @mousedown="clickEvent"
                  @mouseup="dragEnd"
                  @mouseout="dragEnd"
                  @mousemove="dragEvent"
                  @touchstart="touchEvent"
                  @touchmove="swipeEvent"
                  @touchend="dragEnd"
                ></canvas>
              </div>
            </div>
          </v-col>
        </v-row>

        <v-row dense>
          <v-col cols="2">
            <div>
              <button v-on:click="modeChange">{{ penMode }}</button>
            </div>
          </v-col>
          <v-col cols="2">
            <div>
              <button v-on:click="undo">←</button>
            </div>
          </v-col>
          <v-col cols="2">
            <div>
              <button v-on:click="redo">→</button>
            </div>
          </v-col>
          <v-col cols="2">
            <div>
              <button v-on:click="drawGrid">grid</button>
            </div>
          </v-col>
          <v-col cols="2">
            <div>
              <button v-on:click="Save">保存</button>
            </div>
          </v-col>
        </v-row>

        <v-row dense class="palletarea">
          <v-layout justify-center>
            <pallet
              v-for="item in colorPallet"
              :key="item"
              :color="item"
              @getcolor="getPalletColor"
            ></pallet>
          </v-layout>
        </v-row>
      </v-container>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { canvasDataModule } from '../../store/modules/canvasData'
import pallet from '@/components/canvas/pallet.vue'

type Pointed = {
  X: number
  Y: number
}

const canvasConstants = {
  canvasStyleSize: 334,
  canvasSizeMagnification: 0.87
} as const

type canvasConstants = typeof canvasConstants[keyof typeof canvasConstants]

type Stack = {
  indexData: number[]
}

@Component({
  middleware: 'auth',
  components: {
    pallet
  }
})

// TypeScriptの処理
export default class CanvasPage extends Vue {
  get getRange(): number {
    return canvasDataModule.canvasRange
  }
  get getMagnification(): number {
    return canvasDataModule.canvasMagnification
  }
  get getColorPallet(): string[] {
    return canvasDataModule.palletColor
  }
  get getCanvasIndexData(): number[] {
    return canvasDataModule.canvasIndexData
  }

  // 初期値
  pointed: Pointed = { X: 0, Y: 0 } //現在のグリッド座標
  beforePointed: Pointed = { X: 0, Y: 0 } //一個前のグリッド座標
  colorPallet: string[] = this.getColorPallet //パレットの色
  selectingColor: string = this.colorPallet[1] //選択中の色
  palletIndex: number = 1 //選択中の色のパレットにおける順番
  canvasIndexData: number[] = this.getCanvasIndexData //キャンバスに塗られている色の保存領域
  stackMaxSize: number = 100 //巻き戻し可能な最大回数の設定
  undoDataStack: Stack[] = [] //巻き戻しに使う画面データ
  redoDataStack: Stack[] = [] //やり直しに使う画面データ
  canvas: HTMLCanvasElement | null = null //イラストを描くキャンバス
  canvasCtx: CanvasRenderingContext2D | null = null //↑のコンテキスト
  gridCanvas: HTMLCanvasElement | null = null //グリッド線が描かれたキャンバス
  gridCanvasCtx: CanvasRenderingContext2D | null = null //↑のコンテキスト
  rect: DOMRect | null = null //要素の寸法とそのビューポートに対する位置
  isDrag: boolean = false //ドラッグしているかのフラグ
  isGrid: boolean = false //グリッドの表示の有無のフラグ
  penMode: string = 'pen' //ペンのモード
  pageActive: boolean = false //画面が読み込まれたかどうかのフラグ
  canvasMagnification: number = this.getMagnification // 表示倍率
  canvasRange: number = this.getRange // キャンバス横幅.縦幅
  canvasStyleSize: number = 334
  canvasSizeMagnification: number = 0.87

  handleTouchMove(event: UIEvent) {
    event.preventDefault()
  }
  mounted(): void {
    this.canvas = document.querySelector('#drowcanvas')
    this.canvasCtx = this.canvas!.getContext('2d')
    this.canvas!.style.width = 334 + 'px'
    this.canvas!.style.height = 334 + 'px'
    this.canvas!.style.border = '1px solid rgb(0,0,0)'
    this.redraw(this.canvasIndexData)
    this.drawGrid()
    // スマホでのタッチ操作でのスクロール禁止
    document.addEventListener('touchmove', this.handleTouchMove, {
      passive: false
    })
    this.pageActive = true
  }
  beforeDestroy(): void {
    // コンポーネントが破棄される直前の処理
    document.removeEventListener('touchmove', this.handleTouchMove)
  }

  //メソッド
  //クリックしたパレットの色を取得
  getPalletColor(newValue: string): void {
    this.selectingColor = newValue
    this.palletIndex = this.colorPallet.indexOf(this.selectingColor)
  }
  //ペンのモードチェンジ
  modeChange(): void {
    if (this.penMode == 'pen') {
      this.penMode = 'bucket'
    } else {
      this.penMode = 'pen'
    }
  }
  // 描画(mousemove)
  dragEvent(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.rect = this.canvas!.getBoundingClientRect()
    //キャンバス内におけるXY座標を取得
    var x = (e.clientX - this.rect.left) / 0.87
    var y = (e.clientY - this.rect.top) / 0.87
    this.drowingEvent(x, y)
  }

  swipeEvent(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.rect = this.canvas!.getBoundingClientRect()
    //キャンバス内におけるXY座標を取得
    var x = (e.touches[0].pageX - this.rect.left) / 0.87
    var y = (e.touches[0].pageY - this.rect.top) / 0.87
    this.drowingEvent(x, y)
  }

  drowingEvent(x: number, y: number): void {
    if (!this.pageActive) {
      return
    }
    //キャンバス内におけるXY座標を取得
    //ドットのグリッド座標を取得
    this.beforePointed['X'] = this.pointed['X']
    this.beforePointed['Y'] = this.pointed['Y']
    this.pointed['X'] = Math.floor(x / this.canvasMagnification)
    this.pointed['Y'] = Math.floor(y / this.canvasMagnification)
    switch (this.penMode) {
      case 'pen':
        //なめらかな線を描画するためドラッグ時は直線で描く
        this.drawLine(
          this.beforePointed['X'],
          this.pointed['X'],
          this.beforePointed['Y'],
          this.pointed['Y']
        )
        break
      case 'bucket':
        //バケツ中にドラッグしても何も起きない
        break
    }
  }

  // 描画開始（mousedown）
  clickEvent(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.isDrag = true
    this.beforeDraw() //undoの準備
    switch (this.penMode) {
      case 'pen':
        this.drawdot(this.pointed['X'], this.pointed['Y'])
        break

      case 'bucket':
        this.drawFill(this.pointed['X'], this.pointed['Y'])
    }
  }
  touchEvent(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.isDrag = true
    this.beforeDraw() //undoの準備

    this.rect = this.canvas!.getBoundingClientRect()
    var x = (e.touches[0].pageX - this.rect.left) / 0.87
    var y = (e.touches[0].pageY - this.rect.top) / 0.87

    this.beforePointed['X'] = this.pointed['X']
    this.beforePointed['Y'] = this.pointed['Y']
    this.pointed['X'] = Math.floor(x / this.canvasMagnification)
    this.pointed['Y'] = Math.floor(y / this.canvasMagnification)

    switch (this.penMode) {
      case 'pen':
        this.drawdot(this.pointed['X'], this.pointed['Y'])
        break

      case 'bucket':
        this.drawFill(this.pointed['X'], this.pointed['Y'])
    }
  }

  // 描画終了（mouseup, mouseout）
  dragEnd(): void {
    if (!this.pageActive) {
      return
    }
    this.canvasCtx!.closePath()
    this.isDrag = false
  }

  //直線
  //x1,y1 = 始点　x2,y2 = 終点
  drawLine(x1: number, x2: number, y1: number, y2: number): void {
    let Xdif = Math.abs(x1 - x2), //X方向の移動距離の絶対値
      Ydif = Math.abs(y1 - y2), //Y方向の移動距離の絶対値
      Xdif2 = Xdif * 2,
      Ydif2 = Ydif * 2,
      Xvek = x2 > x1 ? 1 : -1, //X方向のベクトル
      Yvek = y2 > y1 ? 1 : -1, //Y方向のベクトル
      x = x1,
      y = y1
    if (Xdif >= Ydif) {
      //Xに何マス進んだらY方向に1進むか計算し、そこから直線を描画する
      let e = -Xdif
      for (let i = 0; i <= Xdif; i++) {
        if (x < 0 || x >= this.canvasRange || y < 0 || y >= this.canvasRange)
          break

        this.drawdot(x, y)
        x += Xvek //Xが1進む
        e += Ydif2
        if (e >= 0) {
          y += Yvek //Yが1進む
          e -= Xdif2 //割り切れない場合を考え端数は切り捨てない
        }
      }
    } else {
      //上と同様
      let e = -Ydif
      for (let i = 0; i <= Ydif; i++) {
        if (x < 0 || x >= this.canvasRange || y < 0 || y >= this.canvasRange)
          break

        this.drawdot(x, y)
        y += Yvek
        e += Xdif2
        if (e >= 0) {
          x += Xvek
          e -= Ydif2
        }
      }
    }
  }

  //塗りつぶし
  drawFill(x: number, y: number): void {
    var color = this.canvasIndexData[y * this.canvasRange + x]
    if (color == this.palletIndex) {
      return
    }
    this.f(x, y, color)
  }
  //塗りつぶしの再帰処理
  f(x: number, y: number, color: number): void {
    if (color == this.palletIndex) {
      return
    }
    //クリックした場所と違う色に当たるまで上下左右に走査し続ける
    if (x >= this.canvasRange || x < 0) return
    if (y >= this.canvasRange || y < 0) return
    if (this.canvasIndexData[y * this.canvasRange + x] === color) {
      this.drawdot(x, y)
      this.f(x - 1, y, color)
      this.f(x + 1, y, color)
      this.f(x, y - 1, color)
      this.f(x, y + 1, color)
    }
  }

  //ドット描画
  drawdot(x: number, y: number): void {
    this.canvasCtx!.beginPath()
    if (!this.isDrag) {
      return
    }
    this.canvasCtx!.fillStyle = this.selectingColor
    this.canvasCtx!.fillRect(
      x * this.canvasMagnification,
      y * this.canvasMagnification,
      this.canvasMagnification,
      this.canvasMagnification
    )
    this.canvasIndexData[y * this.canvasRange + x] = this.palletIndex
  }

  //クリック時、やり直しのためのデータを処理する
  beforeDraw(): void {
    // やり直し用スタックの中身を削除
    this.redoDataStack = []
    // 元に戻す用の配列が最大保持数より大きくなっているかどうか
    if (this.undoDataStack.length >= this.stackMaxSize) {
      // 条件に該当する場合末尾の要素を削除
      this.undoDataStack.pop()
    }
    // 元に戻す配列の先頭にデータを保持する
    this.undoDataStack.unshift({
      indexData: this.canvasIndexData.slice()
    })
  }

  //やり直し(undo)
  undo(): void {
    // データがなければ処理を終了する
    if (this.undoDataStack.length <= 0) return
    // やり直し用の配列に元に戻す操作をする前のデータをスタックしておく
    this.redoDataStack.unshift({
      indexData: this.canvasIndexData.slice()
    })
    // 元に戻す配列の先頭からデータを取得
    var imageData: number[] = this.undoDataStack.shift()!.indexData
    this.redraw(imageData)
  }

  //やり直しの取り消し(redo)
  redo(): void {
    // データがなければ処理を終了する
    if (this.redoDataStack.length <= 0) return
    // 元に戻す用の配列にやり直し操作をする前のデータをスタックしておく
    this.undoDataStack.unshift({
      indexData: this.canvasIndexData!.slice()
    })
    // やり直す配列の先頭からデータを取得
    var imageData: number[] = this.redoDataStack.shift()!.indexData
    this.redraw(imageData)
  }

  //渡されたcanvasのindexdataからドット絵を再描画する
  redraw(data: number[]): void {
    for (var i = 0; i < this.canvasRange; i++) {
      for (var j = 0; j < this.canvasRange; j++) {
        this.canvasCtx!.fillStyle = this.colorPallet[
          data[j * this.canvasRange + i]
        ]
        this.canvasCtx!.fillRect(
          i * this.canvasMagnification,
          j * this.canvasMagnification,
          this.canvasMagnification,
          this.canvasMagnification
        )
        this.canvasIndexData[j * this.canvasRange + i] =
          data[j * this.canvasRange + i]
      }
    }
  }

  //グリッドの描画
  drawGrid(): void {
    this.gridCanvas = document.querySelector('#gridcanvas')
    this.gridCanvasCtx = this.gridCanvas!.getContext('2d')
    this.gridCanvas!.style.width = 334 + 'px'
    this.gridCanvas!.style.height = 334 + 'px'
    this.gridCanvas!.style.border = '1px solid rgb(0, 0, 0)'
    this.gridCanvasCtx!.beginPath()
    this.gridCanvasCtx!.globalCompositeOperation = 'source-over'
    // 線の色・幅
    this.gridCanvasCtx!.strokeStyle = 'rgb(0, 0, 0)'
    this.gridCanvasCtx!.lineWidth = 1

    if (this.isGrid == false) {
      // 縦線
      for (var i = 1; i < this.canvasRange + 1; i++) {
        this.gridCanvasCtx!.moveTo(i * this.canvasMagnification + 0.5, 0.5)
        this.gridCanvasCtx!.lineTo(
          i * this.canvasMagnification + 0.5,
          this.canvasRange * this.canvasMagnification + 0.5
        )
      }
      // 横線
      for (var i = 1; i < this.canvasRange + 1; i++) {
        this.gridCanvasCtx!.moveTo(0.5, i * this.canvasMagnification + 0.5)
        this.gridCanvasCtx!.lineTo(
          this.canvasRange * this.canvasMagnification + 0.5,
          i * this.canvasMagnification + 0.5
        )
      }
      //描画
      this.gridCanvasCtx!.stroke()
      this.isGrid = true
    } else {
      //外枠を残して削除する
      this.gridCanvasCtx!.beginPath()
      this.gridCanvasCtx!.clearRect(
        0,
        0,
        this.canvasRange * this.canvasMagnification,
        this.canvasRange * this.canvasMagnification
      )
      this.isGrid = false
    }
  }

  Save(): void {
    //canvasのインデックスデータとパレットデータをストアへ
    canvasDataModule.setCanvasIndexData(this.canvasIndexData)
    canvasDataModule.setPalletColor(this.colorPallet)
    this.$router.push('/creator/save')
    return
  }
}
</script>

<style>
.DrowCanvas {
  position: relative;
  top: -15px;
}
.GridCanvas {
  position: absolute;
  opacity: 0.5;
  top: 0px;
  left: 0px;
}
.palletarea {
  display: flex;
  justify-content: space-between;
}
</style>
