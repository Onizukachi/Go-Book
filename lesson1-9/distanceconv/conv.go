package distanceconv

func MetrToFt(m Metr) Ft { return Ft(m * 3.28084) }
func FtToMetr(f Ft) Metr { return Metr(f * 0.3048) }
