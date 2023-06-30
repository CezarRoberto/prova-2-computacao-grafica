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
	// Abre o arquivo de vídeo
	videoFile := "./question2/assets/video"
	video, err := gocv.VideoCaptureFile(videoFile)
	if err != nil {
		log.Fatalf("Erro ao abrir o vídeo: %v", err)
	}
	defer video.Close()

	// Lê o primeiro quadro do vídeo
	frame := gocv.NewMat()
	defer frame.Close()

	if ok := video.Read(&frame); !ok {
		log.Fatal("Não foi possível ler o primeiro quadro do vídeo")
	}

	// Cria uma janela para exibir o vídeo
	window := gocv.NewWindow("Selecione o objeto")
	defer window.Close()

	// Exibe o primeiro quadro na janela
	window.IMShow(frame)

	// Aguarda o usuário selecionar o objeto
	rect := gocv.SelectROI("Selecione o objeto", frame)
	if rect.Empty() {
		log.Fatal("Objeto não selecionado")
	}

	// Inicializa as variáveis para armazenar as posições do objeto
	var positionsX []float64
	var positionsY []float64

	// Loop para processar cada quadro do vídeo
	for {
		// Lê o próximo quadro
		if ok := video.Read(&frame); !ok {
			break
		}

		// Extrai a região de interesse (objeto) do quadro
		object := frame.Region(rect)

		// Obtém as coordenadas do centro do objeto
		centerX := float64(rect.Min.X + rect.Dx()/2)
		centerY := float64(rect.Min.Y + rect.Dy()/2)

		// Armazena as posições do objeto
		positionsX = append(positionsX, centerX)
		positionsY = append(positionsY, centerY)

		// Exibe as posições do objeto
		fmt.Printf("Posição: X=%.2f, Y=%.2f\n", centerX, centerY)

		// Desenha um retângulo em volta do objeto
		gocv.Rectangle(&frame, rect, color.RGBA{0, 255, 0, 0}, 2)

		// Exibe o quadro com o objeto destacado
		window.IMShow(frame)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

	// Plota o gráfico com as posições do objeto
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
