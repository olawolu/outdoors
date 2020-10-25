package outdoors

import (
	"errors"
	"strings"
)

// PriceLevel that a destination costs
type PriceLevel int8

// price Pricelevels
const (
	_ PriceLevel = iota
	PriceLevel1
	PriceLevel2
	PriceLevel3
	PriceLevel4
	PriceLevel5
)

var priceStrings = map[string]PriceLevel{
	"$":     PriceLevel1,
	"$$":    PriceLevel2,
	"$$$":   PriceLevel3,
	"$$$$":  PriceLevel4,
	"$$$$$": PriceLevel5,
}

// PriceRange of the destinations
type PriceRange struct {
	From PriceLevel
	To   PriceLevel
}

func (p PriceRange) String() string {
	return p.From.String() + "..." + p.To.String()
}

func (l PriceLevel) String() string {
	for s, v := range priceStrings {
		if l == v {
			return s
		}
	}
	return "invalid!"
}

// ParsePriceRange returns the PriceLevel equivalent of a set of strings
func ParsePriceRange(s string) (PriceRange, error) {
	var r PriceRange
	seg := strings.Split(s, "...")
	if len(seg) != 2 {
		return r, errors.New("Invalid price range")
	}
	r.From = ParsePrice(seg[0])
	r.To = ParsePrice(seg[1])
	return r, nil
}

// ParsePrice returns the PriceLevel equivalent of a string
func ParsePrice(s string) PriceLevel {
	return priceStrings[s]
}