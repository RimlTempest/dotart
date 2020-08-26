<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm12 md12>
      <v-container fluid>
        <v-row dense>
          <v-col cols="12">
            <div class="ResultCanvas">
              <canvas
                class="resultcanvas"
                id="resultcanvas"
                width="384px"
                height="384px"
              ></canvas>
            </div>
          </v-col>
        </v-row>
        <v-row dense>
          <v-col cols="12">
            <div class="SaveCanvas">
              <canvas id="savecanvas"></canvas>
            </div>
          </v-col>
        </v-row>
        <v-row dense>
          <v-col cols="12">
            <v-card light hover max-width="1000" class="justify-center">
              <v-form ref="form" lazy-validation>
                <v-card-title class="justify-center">画像保存</v-card-title>
              </v-form>
            </v-card>
          </v-col>

          <v-col cols="12">
            <v-card light hover max-width="1000" class="justify-center">
              <v-form ref="form" lazy-validation>
                <v-card-title class="headline">画像サイズ指定</v-card-title>
                <v-card-subtitle>保存するサイズを指定できます</v-card-subtitle>
                <v-select
                  v-model="selectedSize"
                  :items="sizeListItems"
                  label="画像サイズ"
                  return-object
                ></v-select>
                <v-card-actions>
                  <v-spacer />
                  <v-btn color="primary" @click="saveImage">保存</v-btn>
                </v-card-actions>
              </v-form>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { canvasDataModule } from '../../store/modules/canvasData'

type SelectedSize = {
  text: string
  magnification: number
}

type SizeListItems = {
  text: string
  magnification: number
}

@Component({
  middleware: 'auth'
})
// TypeScriptの処理
export default class CreatorPage extends Vue {
  get getCanvasName(): string {
    return canvasDataModule.canvasName
  }
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
  canvasName: string = this.getCanvasName //キャンバスの名前
  canvasRange: number = this.getRange //ドット絵のサイズ
  canvasMagnification: number = this.getMagnification //元の倍率
  saveCanvasMagnification: number = this.getMagnification //保存時の倍率
  colorPallet: string[] = this.getColorPallet
  canvasIndexData: number[] = this.getCanvasIndexData
  canvas: any = null //イラストを描くキャンバス
  context: any = null //↑のコンテキスト
  saveCanvas: any = null //画像サイズ変更のためのキャンバス
  saveContext: any = null //↑のコンテキスト
  maxMagnification = (384 / this.canvasRange) * 2 //画像サイズ変更の際の最大サイズを指定

  selectedSize: SelectedSize = {
    text:
      this.canvasRange * this.canvasMagnification +
      '×' +
      this.canvasRange * this.canvasMagnification,
    magnification: this.canvasMagnification
  }
  //画像サイズと、その画像サイズにするための倍率のリスト
  sizeListItems: SizeListItems[] = [
    {
      text: this.canvasRange + '×' + this.canvasRange,
      magnification: 1
    }
  ]

  public mounted(): void {
    this.canvas = document.querySelector('#resultcanvas')
    this.context = this.canvas.getContext('2d')
    this.canvas.style.width = 334 + 'px'
    this.canvas.style.height = 334 + 'px'
    this.canvas.style.border = '1px solid rgb(0,0,0)'
    this.saveCanvas = document.querySelector('#savecanvas')
    this.saveContext = this.saveCanvas.getContext('2d')
    this.draw(this.canvasRange, this.canvasMagnification, this.context)

    //画像サイズのリスト生成
    //TODO: i = i * 2
    for (let i = 2; i <= 32; i = i * 2) {
      this.sizeListItems.push({
        text: this.canvasRange * i + '×' + this.canvasRange * i,
        magnification: i
      })
    }

    this.sizeListItems.push({
      text:
        this.canvasRange * this.canvasMagnification +
        '×' +
        this.canvasRange * this.canvasMagnification,
      magnification: this.canvasMagnification
    })

    this.sizeListItems.sort(function(
      a: SizeListItems,
      b: SizeListItems
    ): number {
      if (a.magnification < b.magnification) {
        return -1
      }
      if (a.magnification > b.magnification) {
        return 1
      }
      return 0
    })
  }

  //ドット絵のサイズ、表示倍率、キャンバスのデータ、パレットのデータから対象のcanvasにドット絵の描画
  draw(range: number, magnification: number, context: any): void {
    for (let i = 0; i < range; i++) {
      for (let j = 0; j < range; j++) {
        context.fillStyle = this.colorPallet[
          this.canvasIndexData[j * range + i]
        ]
        context.fillRect(
          i * magnification,
          j * magnification,
          magnification,
          magnification
        )
      }
    }
  }

  //選んだ倍率からドット絵を描画し、画像として保存する
  saveImage(): void {
    this.saveCanvasMagnification = this.selectedSize.magnification
    this.saveCanvas.width = this.canvasRange * this.saveCanvasMagnification
    this.saveCanvas.height = this.canvasRange * this.saveCanvasMagnification
    //画像を拡大するのではなく倍率を変えて再描画してるので画質の劣化はないと思う
    this.draw(this.canvasRange, this.saveCanvasMagnification, this.saveContext)
    if (this.saveCanvas.msToBlob) {
      //ブラウザによってはないとこまるらしいです しらんけど・・・
      const blob = this.saveCanvas.msToBlob()
      window.navigator.msSaveBlob(blob, this.canvasName + '.png')
    } else {
      //画像データを対象にしたリンクを生成し、クリックしたことにする
      const link = document.createElement('a')
      link.href = this.saveCanvas.toDataURL('image/png')
      link.download = this.canvasName + ':' + this.selectedSize.text + '.png'
      link.click()
    }
  }
}
</script>

<style>
.ResultCanvas {
  position: relative;
  top: -15px;
}

.SaveCanvas {
  display: none;
}
</style>
