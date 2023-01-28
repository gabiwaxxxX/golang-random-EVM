package amm_calculation

import (
	"math/big"
	"fmt"
)

func GetAmountHypothetical(reserve1, reserve2, amountIn *big.Int ) *big.Int {
	amountWithFee:=new(big.Int).Mul(amountIn,big.NewInt(997)) 
	calc:=new(big.Int).Div(new(big.Int).Mul(reserve2, amountWithFee), new(big.Int).Add(reserve1, amountWithFee))
	return new(big.Int).Div(calc, big.NewInt(1000))
}


func CalculationOfTFloat(r1,r2,v,m *big.Int) *big.Float {
	t:=new(big.Float)
	c0:=new(big.Float).Mul(new(big.Float).Mul(big.NewFloat(9000000),new(big.Float).SetInt(m)),new(big.Float).Mul(new(big.Float).SetInt(r1),new(big.Float).SetInt(r1)))
	c1:=new(big.Float).Mul(new(big.Float).Mul(big.NewFloat(3976036000000),new(big.Float).SetInt(v)),new(big.Float).Mul(new(big.Float).SetInt(r1),new(big.Float).SetInt(r2)))
	c2:=new(big.Float).Mul(new(big.Float).Mul(big.NewFloat(5964054000),new(big.Float).SetInt(m)),new(big.Float).Mul(new(big.Float).SetInt(r1),new(big.Float).SetInt(v)))
	c3:=new(big.Float).Mul(new(big.Float).Mul(big.NewFloat(988053892081),new(big.Float).SetInt(m)),new(big.Float).Mul(new(big.Float).SetInt(v),new(big.Float).SetInt(v)))
	t.Add(c0,c1)
	t.Sub(t,c2)
	t.Add(t,c3)
	return t.Sqrt(t)
}

func CalculationOfT(r1,r2,v,m *big.Int) *big.Int {
	t:=new(big.Int)
	c0:=new(big.Int).Mul(new(big.Int).Mul(big.NewInt(9000000),m),new(big.Int).Mul(r1,r1))
	c1:=new(big.Int).Mul(new(big.Int).Mul(big.NewInt(3976036000000),v),new(big.Int).Mul(r1,r2))
	c2:=new(big.Int).Mul(new(big.Int).Mul(big.NewInt(5964054000),m),new(big.Int).Mul(r1,v))
	c3:=new(big.Int).Mul(new(big.Int).Mul(big.NewInt(988053892081),m),new(big.Int).Mul(v,v))
	t.Add(c0,c1)
	t.Sub(t,c2)
	t.Add(t,c3)
	return t.Sqrt(t)
}

func MaxInput(r1,r2,v,m *big.Int) *big.Int {

	max:=new(big.Int)
	t:=CalculationOfT(r1,r2,v,m)

	c0nmuMul:=new(big.Int).Mul(big.NewInt(501505),t)

	c0num:=new(big.Int).Div(c0nmuMul,big.NewInt(1000000000000))
	sqrtM:=new(big.Int).Sqrt(m)
	c0:=new(big.Int).Div(c0num,sqrtM)
	
	c1Mul:=new(big.Int).Mul(big.NewInt(10015),r1)
	c1:=new(big.Int).Div(c1Mul,big.NewInt(10000))

	c2Mul:=new(big.Int).Mul(big.NewInt(4985),v)
	c2:=new(big.Int).Div(c2Mul,big.NewInt(10000))

	fmt.Println("c0",c0)
	fmt.Println("c1",c1)
	fmt.Println("c2",c2)
	max.Sub(c0,c1)
	max.Sub(max,c2)
	return max
}

