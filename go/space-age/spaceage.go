package space

//Planet is the string name of a planet
type Planet string

//Age calcs how many years old someone is based off how many earth seconds olds they are
func Age(age float64, planet Planet) float64 {

	switch planet {
	case "Earth":
		return secondsToPlanetRelativeYears(age, 1)
	case "Mercury":
		return secondsToPlanetRelativeYears(age, 0.2408467)
	case "Venus":
		return secondsToPlanetRelativeYears(age, 0.61519726)
	case "Mars":
		return secondsToPlanetRelativeYears(age, 1.8808158)
	case "Jupiter":
		return secondsToPlanetRelativeYears(age, 11.862615)
	case "Saturn":
		return secondsToPlanetRelativeYears(age, 29.447498)
	case "Uranus":
		return secondsToPlanetRelativeYears(age, 84.016846)
	case "Neptune":
		return secondsToPlanetRelativeYears(age, 164.79132)
	default:
		return age
	}

}

func secondsToPlanetRelativeYears(age, relativeModifier float64) float64 {
	return age * 3.16880878e-8 / relativeModifier
}
