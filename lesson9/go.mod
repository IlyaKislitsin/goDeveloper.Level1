module github.com/IlyaKislitsin/goDeveloper.Level1/lesson9

go 1.16

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config => ./config

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config/console => ./config/console

replace github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config/file => ./config/file

require (
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config/console v0.0.0-00010101000000-000000000000
	github.com/IlyaKislitsin/goDeveloper.Level1/lesson9/config/file v0.0.0-00010101000000-000000000000
	github.com/pelletier/go-toml v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
