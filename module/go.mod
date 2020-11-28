module module_test

go 1.14

require (
	github.com/JameyWoo/modtest v0.0.0
	github.com/JameyWoo/goto/ds v0.0.0
	"github.com/JameyWoo/goto/module" v0.0.0
)

replace (
	github.com/JameyWoo/modtest v0.0.0 => ../../modtest
	github.com/JameyWoo/goto/ds v0.0.0 => ../ds
	"github.com/JameyWoo/goto/module" v0.0.0 => ./
)