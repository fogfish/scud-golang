package core_test

import (
	"testing"

	core "github.com/fogfish/blueprint-serverless-golang"
	"github.com/fogfish/it/v2"
)

func TestPet(t *testing.T) {
	var p core.Pet

	it.Then(t).Should(
		it.Equal(p.ID, ""),
		it.Equal(p.Category, ""),
		it.Equal(p.Price, 0.0),
	)
}
