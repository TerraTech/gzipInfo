module github.com/gzipInfo

go 1.12

require (
	futurequest.net/FQgolibs v0.0.0-00010101000000-000000000000
	futurequest.net/FQgolibs/tools v0.0.0-00010101000000-000000000000
	github.com/TerraTech/gzipInfo v3.0.3+incompatible
	github.com/TerraTech/pflag v1.0.4-0.20190523114457-3d0973303b16
	github.com/spf13/pflag v1.0.3 // indirect
)

replace (
	futurequest.net/FQgolibs => ../../futurequest.net/FQgolibs
	futurequest.net/FQgolibs/tools => ../../futurequest.net/FQgolibs/tools
	github.com/TerraTech/gzipInfo/pkg/gzipInfo => ./pkg/gzipinfo
)
