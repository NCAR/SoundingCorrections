# Sounding Corrections

These are various utility functions with original sources that can be applied to various sounding data sets.

# Vapor Pressure over Liquid Water

## Wexler-76 Vapor Pressure over liquid water
The function ```VaporPressureOverWaterWexler``` is an implementation of the Wexler-1976 paper to find the Vapor Pressure over liquid water given a temperature. This has not been greatly optimized for speed, but purposefully follows the expanded series form given in the paper.

## ITS-90 / Wexler-Hardy 1998 Humidity Correction
The function ```VaporPressureOverWaterITS90``` is an implementation of [Equation 2 of ITS-90 FORMULATIONS FOR VAPOR PRESSURE, FROSTPOINT TEMPERATURE, DEWPOINT TEMPERATURE, AND ENHANCEMENT FACTORS IN THE RANGE –100 TO +100 C ](http://www.rhs.com/papers/its90form.pdf) with the ITS-90 scale. This has not been greatly optimized for speed, but purposefully follows the expanded series form given in the paper.

## Correction of RH via Wexler-Hardy-1998
This corrects a raw RH value by scaling by the ratios ```VaporPressureOverWaterITS90```.  This is a simply a convenience function equivilant to ```sensorRH * VaporPressureOverWaterITS90(Tsensor) / VaporPressureOverWaterITS90(Tambient)```.