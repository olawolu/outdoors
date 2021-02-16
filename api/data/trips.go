package data

import "strings"

type trip struct {
	Name         string
	Destinations []string
}

// Recommendations possible for each trip type
var Recommendations = []interface{}{
	trip{
		Name:         "Romantic",
		Destinations: []string{"restaurant", "bar", "park", "movie_theater"},
	},
	trip{
		Name:         "Shopping",
		Destinations: []string{"clothing_store", "jewelery_store", "supermarket", "mall"},
	},
	trip{
		Name:         "Night Out",
		Destinations: []string{"bar", "restaurant", "night_club", "movie_theater"},
	},
	trip{
		Name:         "Outdoors & Games",
		Destinations: []string{"paintball_arena", "pool"},
	},
	trip{
		Name:         "Pamper",
		Destinations: []string{"hair_care", "beauty_salon", "spa"},
	},
}

// Format controls how a trip object appears publicly
func (t trip) Format() interface{} {
	return map[string]interface{}{
		"name":         t.Name,
		"destinations": strings.Join(t.Destinations, "|"),
	}
}
