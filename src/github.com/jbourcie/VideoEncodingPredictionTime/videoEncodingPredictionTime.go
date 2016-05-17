package VideoEncodingPredictionTime

import "fmt"

type floatCouple struct{
	Start float64
	End float64
}
type predictor struct {
	Name   string
	NumberOfRegion int
	RegionDefinition map[int]*floatCouple
	GopsReg []int
	MRegSize []float64
	MRegTime []float64
	MSize float64
	MTime float64
	A float64 //slope value
	B float64 // initial value at x = 0
	Samples int
	DistinctRegs int
	EnoughRegs int
	Default float64
	Alpha float64 // alpha indices for the exponential moving average
}


type Point struct {
	X float64
	Y float64
}

func linearRegressionLSE(series []*Point) (float64, float64) {

	q := len(series)

	if q == 0 {
		return 0,0
	}

	p := float64(q)

	sum_x, sum_y, sum_xx, sum_xy := 0.0, 0.0, 0.0, 0.0

	for _, p := range series {
		sum_x += p.X
		sum_y += p.Y
		sum_xx += p.X * p.X
		sum_xy += p.X * p.Y
	}

	a := (p*sum_xy - sum_x*sum_y) / (p*sum_xx - sum_x*sum_x)
	b := (sum_y / p) - (a * sum_x / p)

	return a,b
}


func AddMeasure(size float64, time float64, p *predictor){

	if p.Samples>0 {
		// compute the exponential moving average
		p.MSize = p.Alpha * size + (1 - p.Alpha) * p.MSize
		p.MTime = p.Alpha * time + (1 - p.Alpha) * p.MTime
	} else {
		p.MSize = size
		p.MTime = time
	}

	p.Samples ++
	var i int = 0
	for size>p.RegionDefinition[i].End && i<len(p.RegionDefinition){
		i++
	}

	//fmt.Printf("i = %v\n", i)
	if p.GopsReg[i]==0{
		p.DistinctRegs ++
	}

	if p.GopsReg[i]>0 {
		// compute the exponential moving average
		p.MRegSize[i] = p.Alpha * size + (1 - p.Alpha) * p.MRegSize[i]
		p.MRegTime[i] = p.Alpha * time + (1 - p.Alpha) * p.MRegTime[i]
	} else {
		p.MRegSize[i] = size
		p.MRegTime[i] = time
	}

	p.GopsReg[i] ++

	if p.DistinctRegs >= p.EnoughRegs {
		series := make([]*Point, p.NumberOfRegion)
		for r,x := range p.MRegSize {
			series[r] = new(Point)
			series[r].X = x
			series[r].Y = p.MRegTime[r]
		}
		p.A, p.B = linearRegressionLSE(series)
	}

}

func InitializePredictor (name string, numberOfRegions int, defaultTime float64, minChunkSize int, maxChunkSize int) *predictor {
	p:=new(predictor)
	p.Alpha = 0.3
	p.Name = name
	p.NumberOfRegion = numberOfRegions
	p.GopsReg = make([]int, numberOfRegions)
	p.MRegSize = make([]float64, numberOfRegions)
	p.MRegTime = make([]float64, numberOfRegions)
	p.Default = defaultTime
	p.Samples = 0
	p.DistinctRegs = 0
	p.EnoughRegs = numberOfRegions/2
	var regionSize float64 = float64(maxChunkSize - minChunkSize)/float64(numberOfRegions);
	p.RegionDefinition = make(map[int]*floatCouple, numberOfRegions)
	var start, end float64 = 0, float64(minChunkSize)+regionSize
	for i:=0; i<numberOfRegions; i++{
		couple := new(floatCouple)
		couple.Start = start
		couple.End = end
		p.RegionDefinition[i] = couple
		start = end
		end += regionSize
	}
	return p
}

func Predict (size float64, predict *predictor) (float64, bool) {
	if predict.Samples==0 {
		return predict.Default, false
	}
	if predict.DistinctRegs < predict.EnoughRegs {
		return predict.MTime, false
	}
	return predict.A*size + predict.B, true
}