package data

// Geometry defines the coordinates of a place
type Geometry struct {
	Lat float64 `json:"lat"`
	Lng float32 `json:"lng"`
}

// Place defines the fields that describes a destination
type Place struct {
	*Geometry `json:"geometry"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Photos    []*photos `json:"photos"`
	Vicinity  string    `json:"vicinity"`
}

type photos struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

// Format how a Place object appears publicly
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