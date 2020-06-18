package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/lytics/anomalyzer"
)

func main() {
	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  80000,
		LowerBound:  100,
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "highrank", "lowrank", "magnitude"},
		// Methods: []string{"diff", "fence", "highrank"},
	}
	// _ = conf

	fp, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []float64{}

	rd := bufio.NewScanner(fp)
	for rd.Scan() {
		cur, _ := strconv.ParseFloat(rd.Text(), 0)
		data = append(data, cur)
	}

	anom, _ := anomalyzer.NewAnomalyzer(conf, data)
	prob := anom.Push(7000)
	fmt.Println("Anomalous Probability:", prob)
}
