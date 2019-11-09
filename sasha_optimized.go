package main

import (
    "log"
    "fmt"
    "os"
)

func parabola(x float64) float64 {
  return x * x
}

func line(x float64) float64 {
  return 2 - x
}

func main() {
    var x float64 = 1.0
    var y float64 = 1.0

    log.Println("Enter X")
    fmt.Fscan(os.Stdin, &x)

    log.Println("Enter Y")
    fmt.Fscan(os.Stdin, &y)

    log.Println(isInTarget(x, y))
}

func isInSectorD(x, y float64) bool {
  return x >= 0 && y >= 0 && isPositionUnderTheLine(x, y) && !isPositionInsideTheParabola(x, y)
}

func isPositionUnderTheLine(x, y float64) bool {
  return line(x) >= y
}

func isPositionInsideTheParabola(x, y float64) bool {
  return !(parabola(x) > y)
}

func isInSectorC(x, y float64) bool {
  return isPositionUnderTheLine(x, y) && isPositionInsideTheParabola(x, y)
}

func isInTarget(x, y float64) bool {
  return isInSectorC(x, y) || isInSectorD(x, y)
}
