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
)

var (
	//original Wexler coefficients for Wexler's equation to compute vapor pressure over water
	wexlerCoeff = [8]float64{-2.8365744e3, -6.028076559e3, 1.954263612e1, -2.737830188e-2, 1.6261698e-5, 7.0229056e-10, -1.8680009e-13, 2.7150305}

	//ITS-90 coefficients using the Wexler's equation for computing vapor pressure over water
	its90coeff = [8]float64{-2.8365744e3, -6.028076559e3, 1.954263612e1, -2.737830188e-2, 1.6261698e-5, 7.0229056e-10, -1.8680009e-13, 2.7150305}
)

//Computes wexlersEquation using the passed coefficients
func wexlersEquation(tempInC float64, coeffs [8]float64) float64 {
	val := math.Exp(
		func() float64 {
			val := float64(0.0)
			for i := 0; i < 7; i++ {
				val += coeffs[i] * math.Pow(tempInC, float64(i-2))
			}
			return val
		}() + coeffs[7]*math.Log(tempInC))

	if math.IsInf(val, 1) || math.IsInf(val, -1) || math.IsNaN(val) {
		return math.NaN() //force +/-Inf
	}
	return val

}

/*VaporPressureOverWaterITS90 returns the water vapor pressure
(in kPa) over liquid water as per http://www.rhs.com/papers/its90form.pdf given
a sensor temperature (in C).  For more information, please see

	Hardy, B., 1998, ITS-90 Formulations for Vapor Pressure, Frostpoint Temperature, Dewpoint
	Temperature, and Enhancement Factors in the Range –100 to +100 °C, The Proceedings of the Third
	International Symposium on Humidity & Moisture, London, England, April 1998, Volume 1, 214-222.
	(http://www.decatur.de/javascript/dew/resources/its90formulas.pdf)

Any +/- INF values will be reset to NaN()
*/
func VaporPressureOverWaterITS90(tempInC float64) float64 {
	return wexlersEquation(tempInC, its90coeff)
}

/*VaporPressureOverWaterWexler76 returns the water vapor pressure
(in kPa) over liquid water as per the original Wexler-1976 paper:
	Wexler, A., Vapor Pressure Formulation for Water in Range 0 to 100°C. A Revision, Journal of Research of
	the National Bureau of Standards – A. Physics and Chemistry, September – December 1976, Vol. 80A, Nos.
	5 and 6, 775-785
Any +/- INF values will be reset to NaN()
*/
func VaporPressureOverWaterWexler76(tempInC float64) float64 {
	return wexlersEquation(tempInC, wexlerCoeff)
}

/*ITS90CorrectRH returns a corrected RH using the Hardy-Wexler correction given
in the Hardy-1998.  Functionally this is a convenience function for:
	sensorRH * VaporPressureOverWaterITS90(Tsensor) / VaporPressureOverWaterITS90(Tambient)

*/
func ITS90CorrectRH(sensorRH, tempSensorInC, tempAmbientInC float64) float64 {
	return sensorRH * VaporPressureOverWaterITS90(tempSensorInC) / VaporPressureOverWaterITS90(tempAmbientInC)
}
