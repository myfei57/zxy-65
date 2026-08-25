package cable

type Rating string

const (
	RatingNormal Rating = "normal"
	RatingWarm   Rating = "warm"
	RatingHot    Rating = "hot"
)

func Classify(temp, rated float64) Rating {
	if temp >= rated {
		return RatingHot
	}
	if temp >= rated*0.8 {
		return RatingWarm
	}
	return RatingNormal
}
