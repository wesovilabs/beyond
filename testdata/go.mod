module github.com/wesovilabs/goa/testdata

require github.com/wesovilabs/goa v0.0.0

replace (
	github.com/wesovilabs/goa v0.0.0 => ../
	github.com/wesovilabs/goa/examples => ./.goa/
)

go 1.13
