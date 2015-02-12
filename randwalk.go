package main

import (
    "fmt"
    "math"
    "math/rand"
    "os"
    "strconv"
)

//This function randomly generate the angle that the point will move to the next time.
func randAngle() float64 {
    var random float64 = rand.Float64()
    var pi float64 = math.Pi
    var angle float64 = random * 2.0 * pi
    return angle
}

//This funtion chenks that whether the point is in the limited field or not.
func inField(coord, bound float64) bool {
    return coord >= 0 && coord < bound
}

//This function let the point move to the next position randomly
func randStep(x, y, w, h, d float64) (nx, ny float64) {
    nx = x
    ny = y
    for (nx == x && ny == y) || !inField(nx, w) || !inField(ny, h) {
        var moveAngle float64 = randAngle()
        nx = x + d * math.Cos(moveAngle)
        ny = y + d * math.Sin(moveAngle)
    }
    return
}

//This function enables the point to move for the given times
func randWalk(w float64, h float64, d float64, n int) float64 {
    var x float64 = w / 2
    var y float64 = h / 2
    fmt.Println(x, y)

    for i := 0; i < n; i++ {
        x, y = randStep(x, y, w, h, d)
        fmt.Println(x, y)
    }

    var dist float64 = math.Sqrt((x - w/2) * (x - w/2) + (y - h/2) * (y - h/2))
    return dist

}

func main() {
    if len(os.Args) != 6 {
        fmt.Println("Error:  Not enough input number")
        return
    }

    w, err1 := strconv.ParseFloat(os.Args[1], 64)
    h, err2 := strconv.ParseFloat(os.Args[2], 64)
    d, err3 := strconv.ParseFloat(os.Args[3], 64)
    n, err4 := strconv.Atoi(os.Args[4])
    s, err5 := strconv.ParseInt(os.Args[5], 0, 64)



    if err1 != nil || w <= 0 {
        fmt.Println("Error:  Width should be a positive number")
        return
    }
    if err2 != nil || h <= 0 {
        fmt.Println("Error:  Height should be a positive number")
        return
    }
    if err3 != nil || d <= 0 {
        fmt.Println("Error:  Stepsize should be a positive number")
        return
    }
    if err4 != nil || n % 1 != 0 || n <= 0 {
        fmt.Println("Error:  Number of steps should be a positive integer")
        return
    }
    if err5 != nil || s % 1 != 0 || s <= 0 {
        fmt.Println("Error:  Seed should be a positive integer")
        return
    }

    rand.Seed(s)
    var far float64 = randWalk(w, h, d, n)
    fmt.Println("Distance =", far)
}