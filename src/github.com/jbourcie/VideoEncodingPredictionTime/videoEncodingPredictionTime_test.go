package VideoEncodingPredictionTime

import (
	"fmt"
	"testing"
)

func TestInitializePredictor(t *testing.T) {
	t.Parallel()
	numberOfRegion := 10
	var defaultTime float64 = 250
	minChunkSize := 100000
	maxChunkSize := 500000

	p:= InitializePredictor("Test", numberOfRegion, defaultTime, minChunkSize, maxChunkSize)

	if p.Default != defaultTime{
		t.Errorf("Default time was supposed to be %v. Default time found %v",defaultTime,  p.Default)
	}

	if p.NumberOfRegion != numberOfRegion{
		t.Errorf("NumberOfRegion was supposed to be %v. NumberOfREgion found %v",numberOfRegion, p.NumberOfRegion)
	}

	if len(p.RegionDefinition) != numberOfRegion{
		t.Errorf("Lenght of RegionDefinition was supposed to be %v. Lenght of RegionDefinition found %v",numberOfRegion, len(p.RegionDefinition))
	}

	if len(p.GopsReg) != numberOfRegion{
		t.Errorf("Lenght of GopsReg was supposed to be %v. Lenght of GopsReg found %v",numberOfRegion, len(p.GopsReg))
	}

	if len(p.MRegSize) != numberOfRegion{
		t.Errorf("Lenght of MRegSize was supposed to be %v. Lenght of GopsReg found %v",numberOfRegion, len(p.MRegSize))
	}

	if len(p.MRegTime) != numberOfRegion{
		t.Errorf("Lenght of MRegTime was supposed to be %v. Lenght of MRegTime found %v",numberOfRegion, len(p.MRegTime))
	}

	if p.DistinctRegs !=0 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",0, p.DistinctRegs)
	}

	if p.Samples !=0 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",0, p.Samples)
	}

	//for i,r:= range p.RegionDefinition{
	//	fmt.Printf("Region %v : [%v, %v]\n", i, r.Start, r.End)
	//}
}

func TestAddMeasure(t *testing.T) {
	t.Parallel()
	numberOfRegion := 10
	var defaultTime float64 = 250
	minChunkSize := 100000
	maxChunkSize := 500000

	p:= InitializePredictor("Test", numberOfRegion, defaultTime, minChunkSize, maxChunkSize)
	var size float64 = 150000
	var time float64 = 300
	AddMeasure(size, time, p)

	if p.DistinctRegs !=1 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",1, p.DistinctRegs)
	}

	if p.Samples !=1 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",1, p.Samples)
	}

	if p.GopsReg[1]!=1{
		t.Errorf("GopsReg[1] was supposed to be %v. GopsReg[1] found %v",1, p.GopsReg[1])
	}

	if p.MRegSize[1]!=size{
		t.Errorf("MRegSize[1] was supposed to be %v. MRegSize[1] found %v",size, p.MRegSize[1])
	}

	if p.MRegTime[1]!=time{
		t.Errorf("MRegTime[1] was supposed to be %v. MRegTime[1] found %v",time, p.MRegTime[1])
	}

	if p.MSize!=size{
		t.Errorf("MSize was supposed to be %v. MSize found %v",size, p.MSize)
	}

	if p.MTime!=time{
		t.Errorf("MTime was supposed to be %v. MTime found %v",time, p.MTime)
	}

	AddMeasure(size,time, p)

	if p.DistinctRegs !=1 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",1, p.DistinctRegs)
	}

	if p.Samples !=2 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",2, p.Samples)
	}

	if p.GopsReg[1]!=2{
		t.Errorf("GopsReg[1] was supposed to be %v. GopsReg[1] found %v",2, p.GopsReg[1])
	}

	if p.MRegSize[1]!=size{
		t.Errorf("MRegSize[1] was supposed to be %v. MRegSize[1] found %v",size, p.MRegSize[1])
	}

	if p.MRegTime[1]!=time{
		t.Errorf("MRegTime[1] was supposed to be %v. MRegTime[1] found %v",time, p.MRegTime[1])
	}

	if p.MSize!=size{
		t.Errorf("MSize was supposed to be %v. MSize found %v",size, p.MSize)
	}

	if p.MTime!=time{
		t.Errorf("MTime was supposed to be %v. MTime found %v",time, p.MTime)
	}


	size = 120000
	time = 250
	AddMeasure(size,time, p)

	if p.DistinctRegs !=2 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",2, p.DistinctRegs)
	}

	if p.Samples !=3 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",3, p.Samples)
	}

	if p.GopsReg[0]!=1{
		t.Errorf("GopsReg[0] was supposed to be %v. GopsReg[0] found %v",1, p.GopsReg[0])
	}

	if p.MRegSize[0]!=size{
		t.Errorf("MRegSize[0] was supposed to be %v. MRegSize[0] found %v",size, p.MRegSize[0])
	}

	if p.MRegTime[0]!=time{
		t.Errorf("MRegTime[0] was supposed to be %v. MRegTime[0] found %v",time, p.MRegTime[0])
	}

	if p.MSize<=size{
		t.Errorf("MSize was supposed to be > %v. MSize found %v",size, p.MSize)
	}

	if p.MTime<=time{
		t.Errorf("MTime was supposed to be > %v. MTime found %v",time, p.MTime)
	}

	fmt.Printf("MSize %v, MTime %v\n", p.MSize, p.MTime)

	size=float64(maxChunkSize)
	time = 500
	AddMeasure(size,time, p)

	if p.DistinctRegs !=3 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",3, p.DistinctRegs)
	}

	if p.Samples !=4 {
		t.Errorf("DistinctRegs was supposed to be %v. DistinctRegs found %v",4, p.Samples)
	}

	if p.GopsReg[9]!=1{
		t.Errorf("GopsReg[9] was supposed to be %v. GopsReg[9] found %v",1, p.GopsReg[9])
	}

	if p.MRegSize[9]!=size{
		t.Errorf("MRegSize[9] was supposed to be %v. MRegSize[9] found %v",size, p.MRegSize[9])
	}

	if p.MRegTime[9]!=time{
		t.Errorf("MRegTime[9] was supposed to be %v. MRegTime[9] found %v",time, p.MRegTime[9])
	}

	if p.A!=0{
		t.Errorf("A was supposed to be %v. A found %v",0, p.A)
	}

	if p.B!=0{
		t.Errorf("B was supposed to be %v. B found %v",0, p.B)
	}


	fmt.Printf("MSize %v, MTime %v\n", p.MSize, p.MTime)


	size=250000
	time = 400
	AddMeasure(size,time, p)

	if p.A!=0{
		t.Errorf("A was supposed to be %v. A found %v",0, p.A)
	}

	if p.B!=0{
		t.Errorf("B was supposed to be %v. B found %v",0, p.B)
	}
	fmt.Printf("MSize %v, MTime %v\n", p.MSize, p.MTime)

	size=330000
	time = 430
	AddMeasure(size,time, p)

	if p.A==0{
		t.Errorf("A was supposed to be different from %v. A found %v",0, p.A)
	}

	if p.B==0{
		t.Errorf("B was supposed to be different from %v. B found %v",0, p.B)
	}
	fmt.Printf("MSize %v, MTime %v\n", p.MSize, p.MTime)

}

func TestPredict(t *testing.T) {

	t.Parallel()
	numberOfRegion := 10
	var defaultTime float64 = 250
	minChunkSize := 100000
	maxChunkSize := 500000

	p:= InitializePredictor("Test", numberOfRegion, defaultTime, minChunkSize, maxChunkSize)

	predictedTime, result := Predict(0, p)

	if predictedTime!= defaultTime {
		t.Errorf("Predicted Time was supposed to be the default time")
	}

	if result != false {
		t.Errorf("Result was supposed to be false indicating that the result has not been calculated through linear regression")
	}


	var size float64 = 150000
	var time float64 = 300
	AddMeasure(size, time, p)

	predictedTime, result = Predict(size, p)

	if predictedTime!= p.MTime {
		t.Errorf("Predicted Time was supposed to be MTime")
	}

	if result != false {
		t.Errorf("Result was supposed to be false indicating that the result has not been calculated through linear regression")
	}


	AddMeasure(size,time, p)
	predictedTime, result = Predict(size, p)

	if predictedTime!= p.MTime {
		t.Errorf("Predicted Time was supposed to be MTime")
	}

	if result != false {
		t.Errorf("Result was supposed to be false indicating that the result has not been calculated through linear regression")
	}

	size = 120000
	time = 250
	AddMeasure(size,time, p)

	predictedTime, result = Predict(size, p)


	fmt.Printf("Predicted time is : %v \n", predictedTime)
	if result != false {
		t.Errorf("Result was supposed to be false indicating that the result has not been calculated through linear regression")
	}


	size=float64(maxChunkSize)
	time = 500
	AddMeasure(size,time, p)

	predictedTime, result = Predict(size, p)


	fmt.Printf("Predicted time is : %v \n", predictedTime)
	if result != false {
		t.Errorf("Result was supposed to be false indicating that the result has not been calculated through linear regression")
	}


	size=250000
	time = 400
	AddMeasure(size,time, p)

	predictedTime, result = Predict(size, p)


	fmt.Printf("Predicted time is : %v \n", predictedTime)
	if result != false {
		t.Errorf("Result was supposed to be false indicating that the result has not been calculated through linear regression")
	}

	size=330000
	time = 430
	AddMeasure(size,time, p)

	predictedTime, result = Predict(size, p)


	fmt.Printf("Predicted time is : %v \n", predictedTime)
	if result != true {
		t.Errorf("Result was supposed to be true indicating that the result has been calculated through linear regression")
	}

}