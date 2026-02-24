package tossinvest

import "github.com/xtdlib/rat"

var USFeeRatio = rat.Rat("0.08")

func USFee(price any, quantity any) *rat.Rational {
	return rat.RatMin(
		rat.Rat(rat.Rat(rat.Rat(price).Mul(quantity).Mul(USFeeRatio)).SetPrecision(2).DecimalString()),
		rat.Rat("0.01"))
}

func USFeeRate(feeRate any, price any, quantity any) *rat.Rational {
	return rat.Rat(price).Mul(quantity).Quo(USFee(price, quantity))
}
