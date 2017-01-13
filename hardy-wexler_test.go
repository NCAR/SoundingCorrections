package sndcors

/*
Copyright (c) 2015-2016 University Corporation for Atmospheric Research (UCAR).
All rights reserved. Developed by NCAR's Earth Observing Laboratory, UCAR, www.eol.ucar.edu.

This file is part of the AVAPS® software infrastructure. AVAPS® is a registered trademark
of the University Corporation for Atmospheric Research, in the United States and/or other
countries worldwide.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"math"
	"testing"
)

/*nanChecker returns true if:
1) want and got are both NANs
2) want and got are not NANs and want == got
*/
func nanChecker(want, got float64) bool {
	a, b, c := math.IsNaN(want), math.IsNaN(got), (want == got)
	return (a && b) || (!a && !b && c)
}

func Test_wexlersEquation(t *testing.T) {
	type x struct {
		in, out float64
		coeffs  [8]float64
	}

	tests := []x{
		x{in: math.NaN(), out: math.NaN()},
		x{in: 0, out: 0, coeffs: [8]float64{1, 1, 1, 1, 1, 1, 1, 1}},
		x{in: 1, out: 0, coeffs: [8]float64{1, 1, 1, 1, 1, 1, 1, 1}},
		x{in: 1, out: 1},
	}

	for _, x := range tests {
		v := wexlersEquation(x.in, x.coeffs)
		if !nanChecker(x.out, v) {
			t.Log("In: ", x.in)
			t.Log("Coeff", x.coeffs)
			t.Log("Got: ", v)
			t.Log("Wnt: ", x.out)
			t.Error("Didint get the correct value for the output")
		}
	}
}

func Test_VaporPressureOverWaterITS90(t *testing.T) {
	for in, out := range map[float64]float64{
		1:          657.0806266167413,
		math.NaN(): math.NaN(),
		0.0:        611.2129106975888,
		-50.0:      6.437948772985935,
	} {
		if got := VaporPressureOverWaterITS90(in); !nanChecker(out, got) {
			t.Log("In ", in)
			t.Log("Got", got)
			t.Log("Wnt", out)
			t.Errorf("Did not compute ITS90 value properly")
		}
	}
}

func Test_VaporPressureOverWaterWexler76(t *testing.T) {
	for in, out := range map[float64]float64{
		1:          657.0694165167837,
		math.NaN(): math.NaN(),
		0.0:        611.2129098607411,
		-50.0:      6.445011485594395,
	} {
		if got := VaporPressureOverWaterWexler76(in); !nanChecker(out, got) {
			t.Log("In ", in)
			t.Log("Got", got)
			t.Log("Wnt", out)
			t.Errorf("Did not compute Wexler76 value properly")
		}
	}
}

func Test_ITS90CorrectRH(t *testing.T) {
	for out, in := range map[float64]struct{ RH, T1, T2 float64 }{
		math.NaN(): {math.NaN(), math.NaN(), math.NaN()},
	} {
		if got := ITS90CorrectRH(in.RH, in.T1, in.T2); !nanChecker(out, got) {
			t.Log("In ", in)
			t.Log("Got", got)
			t.Log("Wnt", out)
			t.Errorf("Did not compute Wexler76 value properly")
		}
	}
}
