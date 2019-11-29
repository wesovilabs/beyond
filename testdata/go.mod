module github.com/wesovilabs/beyond/testdata

require github.com/wesovilabs/beyond v0.0.0

replace (
	github.com/wesovilabs/beyond v0.0.0 => ../
	github.com/wesovilabs/beyond/examples => ./.beyond/
)

go 1.13
