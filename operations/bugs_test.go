package operations_test

import (
	"testing"

	"github.com/automateafrica/rcc/hamlet"
	"github.com/automateafrica/rcc/operations"
)

func TestHashMatchingIsNotCaseSensitive(t *testing.T) {
	must, wont := hamlet.Specifications(t)

	sut := operations.MetaTemplates{
		Hash: "\t\tCatsAndDogs\r\n",
	}

	must.True(sut.MatchingHash(" catsanddogs "))
	wont.True(sut.MatchingHash(" dogsandcats "))

	sut = operations.MetaTemplates{
		Hash: "catsanddogs",
	}

	must.True(sut.MatchingHash(" CatsAndDogs "))
	wont.True(sut.MatchingHash(" dogsandcats "))
}
