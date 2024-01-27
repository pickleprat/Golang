module example.com/app

go 1.20

replace example.com/randomf => ../Randomizer

replace example.com/erroneous => ../Erroneous

require (
	example.com/erroneous v0.0.0-00010101000000-000000000000
	example.com/randomf v0.0.0-00010101000000-000000000000
)
