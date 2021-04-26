module github.com/IlyaKislitsin/goDeveloper.Level1/lesson8

go 1.16

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson8/config => ./config

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator => ../lesson7/calculator

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/float => ../lesson7/calculator/float

require (
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson7/calculator/float v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson8/config v0.0.0-00010101000000-000000000000
)
