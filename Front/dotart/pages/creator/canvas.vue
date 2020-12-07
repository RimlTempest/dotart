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
                  width="383px"
                  height="383px"
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
              <undoButton :clickEvent="undo" />
            </div>
          </v-col>
          <v-col cols="2">
            <div>
              <redoButton :clickEvent="redo" />
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
import undoButton from '@/components/atomics/undoButton.vue'
import redoButton from '@/components/atomics/redoButton.vue'
import { Point } from '../../types/canvasPoint'
import { Stack } from '../../types/canvasStack'

@Component({
  middleware: 'auth',
  components: {
    palletArea,
    undoButton,
    redoButton
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

  pointed: Point = { X: 0, Y: 0 } //現在のグリッド座標
  beforePointed: Point = { X: 0, Y: 0 } //一個前のグリッド座標
  colorPallet: string[] = this.Getcolorpallet //パレットの色
  palletIndex: number = 1 //選択中の色のパレットにおける順番　初期値は2番目
  selectingColor: string = this.colorPallet[this.palletIndex] //選択中の色
  canvasIndexData: number[] = this.Getcanvasindexdata //キャンバスに塗られている色の保存領域
  stackMaxSize: number = 100 //巻き戻し可能な最大回数の設定
  undoRedoDataStack: Stack[] = [] //undo,redoに使う画面データの配列
  undoRedoDataIndex: number = -1 //↑の、「現在表示している画面のデータ」が格納されている部分の添え字を示す
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

  handleTouchMove(event: UIEvent) {
    event.preventDefault()
  }
  public mounted(): void {
    //ページが立ち上がる時の処理
    //canvasのコンテキスト取得(絵を描く領域)
    this.canvas = document.querySelector('#drowcanvas')
    this.canvasCtx = this.canvas!.getContext('2d')
    //サイズ変更、枠線の追加
    this.canvas!.style.width = this.canvasStyreSize + 'px'
    this.canvas!.style.height = this.canvasStyreSize + 'px'
    this.canvas!.style.border = '1px solid rgb(0,0,0)'
    //canvasのコンテキスト取得(グリッドの領域)
    this.gridCanvas = document.querySelector('#gridcanvas')
    this.gridCanvasCtx = this.gridCanvas!.getContext('2d')
    //サイズの変更、枠線の追加
    this.gridCanvas!.style.width = this.canvasStyreSize + 'px'
    this.gridCanvas!.style.height = this.canvasStyreSize + 'px'
    this.gridCanvas!.style.border = '1px solid rgb(0, 0, 0)'

    //初期色での塗りつぶし、グリッドの描画、undo,redo用配列に追加
    this.redraw(this.canvasIndexData)
    this.drawGrid()
    this.afterDraw()

    // スマホでのタッチ操作でのスクロール禁止
    document.addEventListener('touchmove', this.handleTouchMove, {
      passive: false
    })
    this.pageActive = true
  }
  public beforeDestroy(): void {
    //コンポーネントが破棄される直前の処理
    //スマホでのタッチ操作でのスクロール禁止を解除
    document.removeEventListener('touchmove', this.handleTouchMove)
  }

  /*
  メソッド
  引数 coorX, coorY = キャンバス内のマウスのXY座標
  引数 cellX, cellY = ↑から出したグリッドの座標
  */

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
    let x = (e.clientX - this.rect.left) / this.canvasSizeMagnification
    let y = (e.clientY - this.rect.top) / this.canvasSizeMagnification
    //TODO:関数化
    if (this.isDrag) {
      this.drowing(x, y)
    }
  }

  //スワイプ時の座標取得(touchmove)
  swipe(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.rect = this.canvas!.getBoundingClientRect()
    //キャンバス内におけるXY座標を取得
    let x = (e.touches[0].pageX - this.rect.left) / this.canvasSizeMagnification
    let y = (e.touches[0].pageY - this.rect.top) / this.canvasSizeMagnification
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

    if (!this.isDrag) {
      return
    }
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
        //FIXME: バケツ中にドラッグしても何も起きない
        break
    }
  }

  // クリックしたとき（mousedown）
  click(e: any): void {
    if (!this.pageActive) {
      return
    }
    this.isDrag = true

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
    //タッチした座標の取得
    this.rect = this.canvas!.getBoundingClientRect()
    let x = (e.touches[0].pageX - this.rect.left) / this.canvasSizeMagnification
    let y = (e.touches[0].pageY - this.rect.top) / this.canvasSizeMagnification
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
    //描画を行っていたときのみ動かす
    if (!this.isDrag) return
    this.afterDraw() //undo,redo用配列を追加
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
    let xdiff = Math.abs(coorX1 - coorX2), //X方向の移動距離の絶対値
      ydiff = Math.abs(coorY1 - coorY2), //Y方向の移動距離の絶対値
      xdiff2 = xdiff * 2,
      ydiff2 = ydiff * 2,
      Xvek = coorX2 > coorX1 ? 1 : -1, //X方向のベクトル
      Yvek = coorY2 > coorY1 ? 1 : -1, //Y方向のベクトル
      x = coorX1,
      y = coorY1
    if (xdiff >= ydiff) {
      //Xに何マス進んだらY方向に1進むか計算し、そこから直線を描画する
      let e = -xdiff
      for (let i = 0; i <= xdiff; i++) {
        if (x < 0 || x >= this.canvasRange || y < 0 || y >= this.canvasRange)
          break

        this.drawdot(x, y)
        x += Xvek //Xが1進む
        e += ydiff2
        if (e >= 0) {
          y += Yvek //Yが1進む
          e -= xdiff2 //割り切れない場合を考え端数は切り捨てない
        }
      }
    } else {
      //上と同様
      let e = -ydiff
      for (let i = 0; i <= ydiff; i++) {
        if (x < 0 || x >= this.canvasRange || y < 0 || y >= this.canvasRange)
          break

        this.drawdot(x, y)
        y += Yvek
        e += xdiff2
        if (e >= 0) {
          x += Xvek
          e -= ydiff2
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
    let color = this.canvasIndexData[cellY * this.canvasRange + cellX]
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
  afterDraw(): void {
    //やり直しをした後だった場合、現在の表示内容以降のデータは削除
    if (this.undoRedoDataIndex < this.undoRedoDataStack.length - 1) {
      this.undoRedoDataStack.splice(this.undoRedoDataIndex + 1)
    }
    //巻き戻し最大回数より多かったら先頭を削除、そうでなければ追加
    if (this.undoRedoDataIndex >= this.stackMaxSize) {
      this.undoRedoDataStack.shift()
    } else {
      ++this.undoRedoDataIndex
    }
    //データのプッシュ
    this.undoRedoDataStack.push({
      indexData: this.canvasIndexData.slice()
    })
  }

  //やり直し(undo)
  undo(): void {
    //現在の表示内容が配列の先頭であれば処理を終了する
    if (this.undoRedoDataIndex <= 0) {
      return
    }
    //現在の表示内容の添え字をデクリメントし、それをもとにindexDataを取得して再描画
    this.redraw(this.undoRedoDataStack[--this.undoRedoDataIndex].indexData)
  }

  //やり直しの取り消し(redo)
  redo(): void {
    //現在の表示内容が配列の末尾であれば処理を終了する
    if (this.undoRedoDataIndex >= this.undoRedoDataStack.length - 1) {
      return
    }
    //現在の表示内容の添え字をインクリメントし、それをもとにindexDataを取得して再描画
    this.redraw(this.undoRedoDataStack[++this.undoRedoDataIndex].indexData)
  }

  //渡されたcanvasのindexdataからドット絵を再描画する
  redraw(indexData: number[]): void {
    for (let x = 0; x < this.canvasRange; x++) {
      for (let y = 0; y < this.canvasRange; y++) {
        this.canvasCtx!.fillStyle = this.colorPallet[
          indexData[y * this.canvasRange + x]
        ]
        this.canvasCtx!.fillRect(
          x * this.canvasMagnification,
          y * this.canvasMagnification,
          this.canvasMagnification,
          this.canvasMagnification
        )
        this.canvasIndexData[y * this.canvasRange + x] =
          indexData[y * this.canvasRange + x]
      }
    }
  }

  //グリッドのON、OFF
  drawGrid(): void {
    this.gridCanvasCtx!.beginPath()
    this.gridCanvasCtx!.globalCompositeOperation = 'source-over'
    // 線の色・幅
    this.gridCanvasCtx!.strokeStyle = 'rgb(0, 0, 0)'
    this.gridCanvasCtx!.lineWidth = 1

    if (this.isGrid == false) {
      //グリッドの描画
      // 縦線
      for (let i = 1; i < this.canvasRange + 1; i++) {
        this.gridCanvasCtx!.moveTo(i * this.canvasMagnification - 0.5, -0.5)
        this.gridCanvasCtx!.lineTo(
          i * this.canvasMagnification - 0.5,
          this.canvasRange * this.canvasMagnification - 0.5
        )
      }
      // 横線
      for (let i = 1; i < this.canvasRange + 1; i++) {
        this.gridCanvasCtx!.moveTo(-0.5, i * this.canvasMagnification - 0.5)
        this.gridCanvasCtx!.lineTo(
          this.canvasRange * this.canvasMagnification - 0.5,
          i * this.canvasMagnification - 0.5
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
