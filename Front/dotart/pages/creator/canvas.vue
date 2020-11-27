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
                  @mousedown="click"
                  @mouseup="dragEnd"
                  @mouseout="dragEnd"
                  @mousemove="drag"
                  @touchstart="touch"
                  @touchmove="swipe"
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
              <button v-on:click="undo">
                <icon-Base width="36" height="36" icon-name="undoIcon"
                  ><undoIcon
                /></icon-Base>
              </button>
            </div>
          </v-col>
          <v-col cols="2">
            <div>
              <button v-on:click="redo">
                <icon-Base width="36" height="36" icon-name="redoIcon"
                  ><redoIcon
                /></icon-Base>
              </button>
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

        <palletArea
          :colorPallet="colorPallet"
          :pfirstPalletIndex="palletIndex"
          @getpalletcolor="getpalletcolor"
        ></palletArea>
      </v-container>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { canvasDataModule } from '../../store/modules/canvasData'
import palletArea from '@/components/molecules/palletArea.vue'
import redoIcon from '@/components/atomics/icons/redoIcon.vue'
import undoIcon from '@/components/atomics/icons/undoIcon.vue'
import iconBase from '@/components/atomics/icons/iconBase.vue'
import { Pointed } from '../../types/canvasPointed'
import { Stack } from '../../types/canvasStack'

@Component({
  middleware: 'auth',
  components: {
    palletArea,
    redoIcon,
    undoIcon,
    iconBase
  }
})

// TypeScriptの処理
export default class CanvasPage extends Vue {
  get Getrange(): number {
    return canvasDataModule.canvasRange
  }
  get Getmagnification(): number {
    return canvasDataModule.canvasMagnification
  }
  get Getcolorpallet(): string[] {
    return canvasDataModule.palletColor
  }
  get Getcanvasindexdata(): number[] {
    return canvasDataModule.canvasIndexData
  }

  pointed: Pointed = { X: 0, Y: 0 } //現在のグリッド座標
  beforePointed: Pointed = { X: 0, Y: 0 } //一個前のグリッド座標
  colorPallet: string[] = this.Getcolorpallet //パレットの色
  palletIndex: number = 1 //選択中の色のパレットにおける順番　初期値は2番目
  selectingColor: string = this.colorPallet[this.palletIndex] //選択中の色
  canvasIndexData: number[] = this.Getcanvasindexdata //キャンバスに塗られている色の保存領域
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
  canvasMagnification: number = this.Getmagnification // 表示倍率
  canvasRange: number = this.Getrange // キャンバス横幅.縦幅
  canvasStyreSize: number = 334 //キャンバスの外見上のサイズ
  canvasSizeMagnification: number = 0.87 //キャンパスの表示倍率　外見上のサイズと整合性つけるため必要

  //引数 coorX, coorY = キャンバス内のマウスのXY座標
  //引数 cellX, cellY = ↑から出したグリッドの座標

  handleTouchMove(event: UIEvent) {
    event.preventDefault()
  }
  public mounted(): void {
    //ページが立ち上がる時の処理
    this.canvas = document.querySelector('#drowcanvas')
    this.canvasCtx = this.canvas!.getContext('2d')
    //サイズ変更、枠線の追加
    this.canvas!.style.width = this.canvasStyreSize + 'px'
    this.canvas!.style.height = this.canvasStyreSize + 'px'
    this.canvas!.style.border = '1px solid rgb(0,0,0)'
    //初期色での塗りつぶし、グリッドの描画
    this.redraw(this.canvasIndexData)
    this.drawGrid()
    // スマホでのタッチ操作でのスクロール禁止
    document.addEventListener('touchmove', this.handleTouchMove, {
      passive: false
    })
    this.pageActive = true
  }
  public beforeDestroy(): void {
    // コンポーネントが破棄される直前の処理　スマホでのタッチ操作でのスクロール解禁
    document.removeEventListener('touchmove', this.handleTouchMove)
  }

  //メソッド
  //クリックしたパレットの色を取得
  getpalletcolor(newColor: string, newIndex: number): void {
    this.selectingColor = newColor
    this.palletIndex = newIndex
  }
  //ペンのモードチェンジ
  modeChange(): void {
    if (this.penMode == 'pen') {
      this.penMode = 'bucket'
    } else {
      this.penMode = 'pen'
    }
  }
  // マウス移動時の座標取得(mousemove)
  drag(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.rect = this.canvas!.getBoundingClientRect()
    //キャンバス内におけるXY座標を取得
    var x = (e.clientX - this.rect.left) / this.canvasSizeMagnification
    var y = (e.clientY - this.rect.top) / this.canvasSizeMagnification
    this.drowing(x, y)
  }

  //スワイプ時の座標取得(touchmove)
  swipe(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.rect = this.canvas!.getBoundingClientRect()
    //キャンバス内におけるXY座標を取得
    var x = (e.touches[0].pageX - this.rect.left) / this.canvasSizeMagnification
    var y = (e.touches[0].pageY - this.rect.top) / this.canvasSizeMagnification
    this.drowing(x, y)
  }

  //取得した座標から描画を行う
  drowing(coorX: number, coorY: number): void {
    if (!this.pageActive) {
      return
    }
    //キャンバス内におけるXY座標を取得
    //ドットのグリッド座標を取得
    this.beforePointed['X'] = this.pointed['X']
    this.beforePointed['Y'] = this.pointed['Y']
    this.pointed['X'] = Math.floor(coorX / this.canvasMagnification)
    this.pointed['Y'] = Math.floor(coorY / this.canvasMagnification)
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

  // クリックしたとき（mousedown）
  click(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.isDrag = true
    this.beforeDraw() //undoの準備

    //ペンモードによって処理の変更
    switch (this.penMode) {
      case 'pen':
        this.drawdot(this.pointed['X'], this.pointed['Y'])
        break
      case 'bucket':
        this.drawFill(this.pointed['X'], this.pointed['Y'])
    }
  }
  //タッチしたとき（touchstart）
  touch(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.isDrag = true
    this.beforeDraw() //undoの準備
    //タッチした座標の取得
    this.rect = this.canvas!.getBoundingClientRect()
    var x = (e.touches[0].pageX - this.rect.left) / 0.87
    var y = (e.touches[0].pageY - this.rect.top) / 0.87
    //マウスムーブによる座標の獲得がないので、ここでpointed等を更新する
    this.beforePointed['X'] = this.pointed['X']
    this.beforePointed['Y'] = this.pointed['Y']
    this.pointed['X'] = Math.floor(x / this.canvasMagnification)
    this.pointed['Y'] = Math.floor(y / this.canvasMagnification)
    //ペンモードによって処理の変更
    switch (this.penMode) {
      case 'pen':
        this.drawdot(this.pointed['X'], this.pointed['Y'])
        break
      case 'bucket':
        this.drawFill(this.pointed['X'], this.pointed['Y'])
    }
  }

  // 描画終了（mouseup, mouseout, touchend）
  dragEnd(): void {
    if (!this.pageActive) {
      return
    }
    this.canvasCtx!.closePath()
    this.isDrag = false
  }

  //直線
  //coorX1,coorY1 = 始点　coorX2,coorY2 = 終点
  drawLine(
    coorX1: number,
    coorX2: number,
    coorY1: number,
    coorY2: number
  ): void {
    let Xdif = Math.abs(coorX1 - coorX2), //X方向の移動距離の絶対値
      Ydif = Math.abs(coorY1 - coorY2), //Y方向の移動距離の絶対値
      Xdif2 = Xdif * 2,
      Ydif2 = Ydif * 2,
      Xvek = coorX2 > coorX1 ? 1 : -1, //X方向のベクトル
      Yvek = coorY2 > coorY1 ? 1 : -1, //Y方向のベクトル
      x = coorX1,
      y = coorY1
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

  //ドット描画
  drawdot(cellX: number, cellY: number): void {
    this.canvasCtx!.beginPath()
    if (!this.isDrag) {
      return
    }
    //該当の座標に色を塗るだけ
    this.canvasCtx!.fillStyle = this.selectingColor
    this.canvasCtx!.fillRect(
      cellX * this.canvasMagnification,
      cellY * this.canvasMagnification,
      this.canvasMagnification,
      this.canvasMagnification
    )
    this.canvasIndexData[cellY * this.canvasRange + cellX] = this.palletIndex
  }

  //塗りつぶし
  //引数は座標から出したドットのマス目の位置
  drawFill(cellX: number, cellY: number): void {
    //クリックした位置の色のindexを取得
    var color = this.canvasIndexData[cellY * this.canvasRange + cellX]
    //今の選択中の色と同じならキャンセル
    if (color == this.palletIndex) {
      return
    }
    //再帰処理を読んで走査
    this.f(cellX, cellY, color)
  }
  //塗りつぶしの再帰処理
  f(cellX: number, cellY: number, color: number): void {
    //今の選択中の色と同じならキャンセル
    if (color == this.palletIndex) {
      return
    }
    //クリックした場所と違う色に当たるか画面端に到達するまで上下左右に走査し続ける
    if (cellX >= this.canvasRange || cellX < 0) return
    if (cellY >= this.canvasRange || cellY < 0) return
    if (this.canvasIndexData[cellY * this.canvasRange + cellX] === color) {
      this.drawdot(cellX, cellY)
      this.f(cellX - 1, cellY, color)
      this.f(cellX + 1, cellY, color)
      this.f(cellX, cellY - 1, color)
      this.f(cellX, cellY + 1, color)
    }
  }

  //クリック時に最初に行う処理　やり直しのためのデータを処理する
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
  redraw(indexData: number[]): void {
    for (var i = 0; i < this.canvasRange; i++) {
      for (var j = 0; j < this.canvasRange; j++) {
        this.canvasCtx!.fillStyle = this.colorPallet[
          indexData[j * this.canvasRange + i]
        ]
        this.canvasCtx!.fillRect(
          i * this.canvasMagnification,
          j * this.canvasMagnification,
          this.canvasMagnification,
          this.canvasMagnification
        )
        this.canvasIndexData[j * this.canvasRange + i] =
          indexData[j * this.canvasRange + i]
      }
    }
  }

  //グリッドのON、OFF
  drawGrid(): void {
    this.gridCanvas = document.querySelector('#gridcanvas')
    this.gridCanvasCtx = this.gridCanvas!.getContext('2d')
    //サイズの変更、枠線の追加
    this.gridCanvas!.style.width = this.canvasStyreSize + 'px'
    this.gridCanvas!.style.height = this.canvasStyreSize + 'px'
    this.gridCanvas!.style.border = '1px solid rgb(0, 0, 0)'
    this.gridCanvasCtx!.beginPath()
    this.gridCanvasCtx!.globalCompositeOperation = 'source-over'
    // 線の色・幅
    this.gridCanvasCtx!.strokeStyle = 'rgb(0, 0, 0)'
    this.gridCanvasCtx!.lineWidth = 1

    if (this.isGrid == false) {
      //グリッドの描画
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
      //削除する
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

  //画像保存ページへの遷移
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
</style>
