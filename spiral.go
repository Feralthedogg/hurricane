// spiral.go

package Hurricane

import "math"

func DefaultSpiralFunc(radiusStep, thetaStep, heightStep float64) SpiralFunc {
	return func(index int) (r, theta, z float64) {
		r = float64(index) * radiusStep
		theta = float64(index) * thetaStep
		z = float64(index) * heightStep
		return
	}
}

func LogSpiralFunc(a, b, heightStep float64) SpiralFunc {
	return func(index int) (r, theta, z float64) {
		theta = float64(index) * b
		r = a * math.Exp(b*theta)
		z = float64(index) * heightStep
		return
	}
}

func ArchimedeanSpiralFunc(a, b, heightStep float64) SpiralFunc {
	return func(index int) (r, theta, z float64) {
		theta = float64(index) * b
		r = a + b*theta
		z = float64(index) * heightStep
		return
	}
}
