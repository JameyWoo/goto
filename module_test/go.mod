module module_test

go 1.14

require (
	count v0.0.0
	github.com/JameyWoo/goto/ds v0.0.0-20201128135356-3ca7a0369284
	github.com/JameyWoo/goto/module_test v0.0.0
	github.com/JameyWoo/modtest v0.0.0
	github.com/Jameywoo/modtest/helloworld v0.0.0
	helloworld v0.0.0
)

replace (
	count v0.0.0 => ./src/count
	github.com/JameyWoo/goto/module_test v0.0.0 => ./
	github.com/JameyWoo/modtest v0.0.0 => ../../modtest
	helloworld v0.0.0 => ./src/helloworld
	github.com/Jameywoo/modtest/helloworld v0.0.0 => ../../modtest/helloworld
)
