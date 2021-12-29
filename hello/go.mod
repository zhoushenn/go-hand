module go-hand/hello

go 1.16

require (
	greeting v0.0.0-00010101000000-000000000000
	moul.io/banner v1.0.1
	rsc.io/quote v1.5.2
)

replace greeting => ./../greeting
