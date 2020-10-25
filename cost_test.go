package outdoors_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/olawolu/outdoors"
)

func TestLevels(t *testing.T) {
	is := is.New(t)
	is.Equal(int(outdoors.PriceLevel1), 1)
	is.Equal(int(outdoors.PriceLevel2), 2)
	is.Equal(int(outdoors.PriceLevel3), 3)
	is.Equal(int(outdoors.PriceLevel4), 4)
	is.Equal(int(outdoors.PriceLevel5), 5)
}

func TestPriceString(t *testing.T) {
	is := is.New(t)
	is.Equal(outdoors.PriceLevel1.String(), "$")
	is.Equal(outdoors.PriceLevel2.String(), "$$")
	is.Equal(outdoors.PriceLevel3.String(), "$$$")
	is.Equal(outdoors.PriceLevel4.String(), "$$$$")
	is.Equal(outdoors.PriceLevel5.String(), "$$$$$")
}

func TestParsePrice(t *testing.T) {
	is := is.New(t)
	is.Equal(outdoors.PriceLevel1, outdoors.ParsePrice("$"))
	is.Equal(outdoors.PriceLevel2, outdoors.ParsePrice("$$"))
	is.Equal(outdoors.PriceLevel3, outdoors.ParsePrice("$$$"))
	is.Equal(outdoors.PriceLevel4, outdoors.ParsePrice("$$$$"))
	is.Equal(outdoors.PriceLevel5, outdoors.ParsePrice("$$$$$"))
}

func TestParsePriceRange(t *testing.T) {
	is := is.New(t)
	var r outdoors.PriceRange
	var err error
	r, err = outdoors.ParsePriceRange("$$...$$$")
	is.NoErr(err)
	is.Equal(r.From, outdoors.PriceLevel2)
	is.Equal(r.To, outdoors.PriceLevel3)
	r, err = outdoors.ParsePriceRange("$...$$$$$")
	is.NoErr(err)
	is.Equal(r.From, outdoors.PriceLevel1)
	is.Equal(r.To, outdoors.PriceLevel5)
}

func TestPriceRangeString(t *testing.T) {
	is := is.New(t)
	r := outdoors.PriceRange{
		From: outdoors.PriceLevel2,
		To:   outdoors.PriceLevel4,
	}
	is.Equal("$$...$$$$", r.String())
}
