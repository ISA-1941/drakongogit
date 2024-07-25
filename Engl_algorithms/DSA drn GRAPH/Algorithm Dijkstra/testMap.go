package main
import (
	"fmt"
	"math"
)
func main () {
var mapLen = make(map[string] float64)
var mapPrice = make(map[string] float64)
var priceLen = make([] float64, 10)
var length float64 = 10000.00

mapLen["0-1"] = 10.45
mapLen["0-6"] = 13.57
mapLen["1-2"] = 18.83
mapLen["1-4"] = 14.23
mapLen["2-0"] = 43.67
mapLen["2-4"] = 27.93
mapLen["3-5"] = 47.33
mapLen["4-3"] = 29.27
mapLen["5-4"] = 57.69
mapLen["6-3"] = 33.31
//
mapPrice["0-1"] = 1.50
mapPrice["0-6"] = 1.80
mapPrice["1-2"] = 2.00
mapPrice["1-4"] = 3.00
mapPrice["2-0"] = 1.20
mapPrice["2-4"] = 1.30
mapPrice["3-5"] = 1.00
mapPrice["4-3"] = 1.30
mapPrice["5-4"] = 1.00
mapPrice["6-3"] = 1.60
//
	priceLen[0] = mapLen["0-1"] * mapPrice["0-1"] 
fmt.Println("Map --> ", length, mapLen["0-1"], mapPrice["0-1"], (math.Round(priceLen[0]*100)/100))
 // 12.35 (round to nearest)
}
