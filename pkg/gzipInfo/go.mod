module github.com/gzipinfo/pkg/gzipinfo

go 1.12

replace (
	futurequest.net/FQgolibs => ../../../../futurequest.net/FQgolibs
	futurequest.net/FQgolibs/tools => ../../../../futurequest.net/FQgolibs/tools
)

require (
	futurequest.net/FQgolibs v0.0.0-00010101000000-000000000000
	futurequest.net/FQgolibs/tools v0.0.0-00010101000000-000000000000
	github.com/TerraTech/gzipInfo v3.0.3+incompatible
	github.com/stretchr/testify v1.3.0
)
