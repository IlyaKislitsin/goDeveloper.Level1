module github.com/IlyaKislitsin/goDeveloper.Level1/lesson7

go 1.16

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator => ./calculator

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/float => ./calculator/float

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/int => ./calculator/int

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/uint => ./calculator/uint

require (
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/float v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/int v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/uint v0.0.0-00010101000000-000000000000
)
