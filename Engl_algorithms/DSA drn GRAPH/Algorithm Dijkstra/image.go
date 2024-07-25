package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// Load graph data from JSON or CSV file
	graphData := loadGraphData("graph.json")

	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, 1024, 768))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.RGBA{255, 255, 255, 255}), image.ZP, draw.Src)

	// Draw nodes
	for _, node := range graphData.Nodes {
		drawNode(img, node)
	}

	// Draw edges
	for _, edge := range graphData.Edges {
		drawEdge(img, edge)
	}

	// Save the image
	f, err := os.Create("graph.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}

func drawNode(img *image.RGBA, node Node) {
	// Draw a circle or rectangle for the node
	draw.Draw(img, image.Rect(node.X-5, node.Y-5, node.X+5, node.Y+5), image.NewUniform(color.RGBA{0, 0, 0, 255}), image.ZP, draw.Src)
}

func drawEdge(img *image.RGBA, edge Edge) {
	// Draw a line or arrow for the edge
	draw.Draw(img, image.Rect(edge.Source.X, edge.Source.Y, edge.Target.X, edge.Target.Y), image.NewUniform(color.RGBA{128, 128, 128, 255}), image.ZP, draw.Src)

	// Add edge weight as a text label
	font := font.Font{font: "Helvetica", size: 12}
	draw.Draw(img, image.Rect(edge.Source.X+10, edge.Source.Y+10, edge.Target.X-10, edge.Target.Y-10), image.NewUniform(color.RGBA{0, 0, 0, 255}), image.ZP, draw.Src)
	draw.DrawString(img, font, edge.Weight, image.Point{edge.Source.X + 10, edge.Source.Y + 10})
}