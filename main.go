package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/:precision", getEuler)
	r.Run()
}

func getEuler(c *gin.Context) {
	precision, _ := strconv.ParseInt(c.Param("precision"), 10, 64)
	sum := 0.0
	for i := int64(0); i < precision; i++ {
		e := euler(250)
		sum += e
	}
	c.String(http.StatusOK, fmt.Sprintf("%f", sum))
}

func euler(precision float64) float64 {
	fact := uint64(1)
	e := 2.0
	n := uint64(2)
	epsilon := 1.0 / math.Pow(10.0, precision)
	//fmt.Println(epsilon)
	for {
		e0 := e
		fact *= n
		n++
		//fmt.Println(n)
		e += 1.0 / float64(fact)
		//fmt.Println(math.Abs(e - e0))
		if math.Abs(e - e0) <= epsilon {
			break
		}
	}
	return e
}