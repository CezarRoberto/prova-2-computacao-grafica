package main

import (
	"image/color"
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"

	"gocv.io/x/gocv"
)

func main() {
	imgFile := "./assets/image.jpg"
	img := loadImage(imgFile)
	defer img.Close()

	windowOriginal := gocv.NewWindow("Imagem Original")
	windowOriginal.IMShow(img)
	windowOriginal.WaitKey(0)

	grayscaleHist := calculateGrayscaleHistogram(img)
	plotHistogram(grayscaleHist, "Histograma em Escala de Cinza")

	rgbHist := calculateRGBHistogram(img)
	plotRGBHistogram(rgbHist)

	brightImg := applyBrightness(img, 50)

	windowAltered := gocv.NewWindow("Imagem Alterada")
	windowAltered.IMShow(brightImg)
	windowAltered.WaitKey(0)

	grayscaleHistAltered := calculateGrayscaleHistogram(brightImg)
	plotHistogram(grayscaleHistAltered, "Histograma em Escala de Cinza - Imagem Alterada")

	rgbHistAltered := calculateRGBHistogram(brightImg)
	plotRGBHistogram(rgbHistAltered)
}

func loadImage(file string) *gocv.Mat {
	img := gocv.IMRead(file, gocv.IMReadColor)
	if img.Empty() {
		log.Fatalf("Erro ao carregar a imagem: %s", file)
	}
	return &img
}

func calculateGrayscaleHistogram(img *gocv.Mat) plotter.Values {
	grayImg := gocv.NewMat()
	defer grayImg.Close()

	gocv.CvtColor(*img, &grayImg, gocv.ColorBGRToGray)

	hist := gocv.NewMat()
	defer hist.Close()

	gocv.CalcHist([]gocv.Mat{grayImg}, []int{0}, gocv.NewMat(), &hist, []int{256}, []float64{0, 256}, false)

	histogram := make(plotter.Values, hist.Rows())
	for i := 0; i < hist.Rows(); i++ {
		histogram[i] = hist.GetFloatAt(i, 0)
	}

	return histogram
}

func calculateRGBHistogram(img *gocv.Mat) [3]plotter.Values {
	channels := img.Channels()

	hist := [3]gocv.Mat{}
	for i := 0; i < channels; i++ {
		hist[i] = gocv.NewMat()
		defer hist[i].Close()
	}

	gocv.CalcHist([]gocv.Mat{*img}, []int{0}, gocv.NewMat(), &hist[0], []int{256}, []float64{0, 256}, false)
	gocv.CalcHist([]gocv.Mat{*img}, []int{1}, gocv.NewMat(), &hist[1], []int{256}, []float64{0, 256}, false)
	gocv.CalcHist([]gocv.Mat{*img}, []int{2}, gocv.NewMat(), &hist[2], []int{256}, []float64{0, 256}, false)

	histogram := [3]plotter.Values{}
	for i := 0; i < channels; i++ {
		histogram[i] = make(plotter.Values, hist[i].Rows())
		for j := 0; j < hist[i].Rows(); j++ {
			histogram[i][j] = hist[i].GetFloatAt(j, 0)
		}
	}

	return histogram
}

func plotHistogram(values plotter.Values, title string) {
	p, err := plot.New()
	if err != nil {
		log.Fatalf("Erro ao criar o gráfico: %v", err)
	}

	bar, err := plotter.NewBarChart(values, vg.Length(10))
	if err != nil {
		log.Fatalf("Erro ao criar o gráfico de barras: %v", err)
	}

	p.Add(bar)
	p.Title.Text = title

	err = p.Save(4*vg.Inch, 4*vg.Inch, "histograma.png")
	if err != nil {
		log.Fatalf("Erro ao salvar o gráfico: %v", err)
	}
}

func plotRGBHistogram(histogram [3]plotter.Values) {
	p, err := plot.New()
	if err != nil {
		log.Fatalf("Erro ao criar o gráfico: %v", err)
	}

	colors := []color.RGBA{
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		color.RGBA{R: 0, G: 255, B: 0, A: 255},
		color.RGBA{R: 0, G: 0, B: 255, A: 255},
	}

	for i, hist := range histogram {
		line, err := plotter.NewLine(hist)
		if err != nil {
			log.Fatalf("Erro ao criar a linha do gráfico: %v", err)
		}

		line.LineStyle.Color = colors[i]

		p.Add(line)
	}

	p.Title.Text = "Histograma RGB"
	p.Legend.Add("Red", histogram[0])
	p.Legend.Add("Green", histogram[1])
	p.Legend.Add("Blue", histogram[2])
	p.Legend.Top = true

	err = p.Save(4*vg.Inch, 4*vg.Inch, "histograma_rgb.png")
	if err != nil {
		log.Fatalf("Erro ao salvar o gráfico: %v", err)
	}
}

func applyBrightness(img *gocv.Mat, brightness int) *gocv.Mat {
	adjustedImg := gocv.NewMat()
	defer adjustedImg.Close()

	gocv.ConvertScaleAbs(*img, &adjustedImg, 1.0, float64(brightness))

	return &adjustedImg
}
