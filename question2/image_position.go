package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"

	"gocv.io/x/gocv"
)

func main() {
	videoFile := "caminho_para_o_video.mp4"
	video, err := gocv.VideoCaptureFile(videoFile)
	if err != nil {
		log.Fatalf("Erro ao abrir o vídeo: %v", err)
	}
	defer video.Close()

	frame := gocv.NewMat()
	defer frame.Close()

	if ok := video.Read(&frame); !ok {
		log.Fatal("Não foi possível ler o primeiro quadro do vídeo")
	}

	window := gocv.NewWindow("Selecione o objeto")
	defer window.Close()

	window.IMShow(frame)

	rect := gocv.SelectROI("Selecione o objeto", frame)
	if rect.Empty() {
		log.Fatal("Objeto não selecionado")
	}

	var positionsX []float64
	var positionsY []float64

	for {
		if ok := video.Read(&frame); !ok {
			break
		}

		object := frame.Region(rect)

		centerX := float64(rect.Min.X + rect.Dx()/2)
		centerY := float64(rect.Min.Y + rect.Dy()/2)

		positionsX = append(positionsX, centerX)
		positionsY = append(positionsY, centerY)

		fmt.Printf("Posição: X=%.2f, Y=%.2f\n", centerX, centerY)

		gocv.Rectangle(&frame, rect, color.RGBA{0, 255, 0, 0}, 2)

		window.IMShow(frame)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

	p, err := plot.New()
	if err != nil {
		log.Fatalf("Erro ao criar o gráfico: %v", err)
	}

	points := make(plotter.XYs, len(positionsX))
	for i := range positionsX {
		points[i].X = positionsX[i]
		points[i].Y = positionsY[i]
	}

	s, err := plotter.NewScatter(points)
	if err != nil {
		log.Fatalf("Erro ao criar o gráfico de dispersão: %v", err)
	}

	p.Add(s)
	p.Title.Text = "Posições do Objeto"
	p.X.Label.Text = "Largura"
	p.Y.Label.Text = "Altura"

	err = p.Save(4*vg.Inch, 4*vg.Inch, "grafico.png")
	if err != nil {
		log.Fatalf("Erro ao salvar o gráfico: %v", err)
	}
}
