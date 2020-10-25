package outdoors

var (
	// APIKey for Google's places API
	APIKey  string
	baseURL = "https://maps.googleapis.com/maps/api/place/"
	nearbySearchURL = baseURL + "nearbysearch/json?"
	photosURL = baseURL + "photo?"
)

// Geometry defines the coordinates of a place
type Geometry struct {
	Lat float64 `json:"lat"`
	Lng float32 `json:"lng"`
}

type photos struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

// Place defines the fields that describes a destination
type Place struct {
	*Geometry `json:"geometry"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Photos    []*photos `json:"photos"`
	Vicinity  string    `json:"vicinity"`
}

type response struct {
	Results []Place `json:"results"`
}

// Format controls how a Place object appears publicly
func (p *Place) Format() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}
