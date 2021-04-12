module github.com/IlyaKislitsin/goDeveloper.Level1/lesson3

go 1.16

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson3/calculator => ./calculator

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson3/simple => ./simple

require (
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson3/calculator v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson3/simple v0.0.0-00010101000000-000000000000
)
