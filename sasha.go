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

func testLine() {
  var x float64 = 1.0
  var y float64 = 1.0

  // A
  x = -2
  y = 1
  log.Println("Sector A")
  isPositionUnderTheLine(x, y)

  // B
  x = -3
  y = 6
  log.Println("Sector B")
  isPositionUnderTheLine(x, y)

  // C
  x = -1
  y = 2
  log.Println("Sector C")
  isPositionUnderTheLine(x, y)

  // E
  x = 0
  y = 4
  log.Println("Sector E")
  isPositionUnderTheLine(x, y)

  // On the line
  x = -1
  y = 3
  log.Println("On the line")
  isPositionUnderTheLine(x, y)
}

func testParabola() {
  var x float64 = 1.0
  var y float64 = 1.0

  // A
  x = -2
  y = 1
  log.Println("Sector A")
  isPositionInsideTheParabola(x, y)

  // B
  x = -3
  y = 6
  log.Println("Sector B")
  isPositionInsideTheParabola(x, y)

  // C
  x = -1
  y = 2
  log.Println("Sector C")
  isPositionInsideTheParabola(x, y)

  // E
  x = 0
  y = 4
  log.Println("Sector E")
  isPositionInsideTheParabola(x, y)

  // On the parabola
  x = 1
  y = 1
  log.Println("On the parabola")
  isPositionInsideTheParabola(x, y)
}

func testTamGdeNado() {
  var x float64 = 1.0
  var y float64 = 1.0

  // A
  x = -2
  y = 1
  log.Println("Sector A")
  isInSectorC(x, y)

  // B
  x = -3
  y = 6
  log.Println("Sector B")
  isInSectorC(x, y)

  // C
  x = -1
  y = 2
  log.Println("Sector C")
  isInSectorC(x, y)

  // E
  x = 0
  y = 4
  log.Println("Sector E")
  isInSectorC(x, y)

  // Tam gde nado
  x = 1
  y = 1
  log.Println("Tam gde nado ")
  isInSectorC(x, y)
}

func testSectorD() {
  var x float64 = 1.0
  var y float64 = 1.0

  // A
  x = -2
  y = 1
  log.Println("Sector A")
  isInSectorD(x, y)

  // B
  x = -3
  y = 6
  log.Println("Sector B")
  isInSectorD(x, y)

  // C
  x = -1
  y = 2
  log.Println("Sector C")
  isInSectorD(x, y)

  // E
  x = 0
  y = 4
  log.Println("Sector E")
  isInSectorD(x, y)

  // D
  x = 1
  y = 0.5
  log.Println("Sector D")
  isInSectorD(x, y)

  // D
  x = 1
  y = 1
  log.Println("Sector D Border")
  isInSectorD(x, y)

  // D
  x = 2
  y = 0
  log.Println("Sector D Angle")
  isInSectorD(x, y)
}

func isInSectorD(x, y float64) bool {
  if x >= 0 && y >= 0 && isPositionUnderTheLine(x, y) && !isPositionInsideTheParabola(x, y) {
    log.Println("IS IN SECTOR D")
    return true
  } else {
    log.Println("-------------")
    return false
  }

}

func isPositionUnderTheLine(x, y float64) bool {
  if line(x) >= y {
    log.Println("IS BELOW")
    return true
  } else {
    log.Println("IS ABOVE")
    return false
  }
}

func isPositionInsideTheParabola(x, y float64) bool {
  if parabola(x) > y {
    log.Println("IS OUTSIDE")
    return false
  } else {
    log.Println("IS INSIDE")
    return true
  }
}

func isInSectorC(x, y float64) bool {
    if isPositionUnderTheLine(x, y) && isPositionInsideTheParabola(x, y) {
      log.Println("IS IN SECTOR C")
      return true
    } else {
      log.Println("-------------")
      return false
    }
}


func isInTarget(x, y float64) bool {
  return isInSectorC(x, y) || isInSectorD(x, y)
}
