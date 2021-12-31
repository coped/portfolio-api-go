package flags

import "flag"

type Flags interface {
	sesUsername 
	sesPassword
}

var sesUsername = flag.String("username", "", "Username credential")
var sesPassword = flag.String("password", "", "Password credential")
var production = flag.Bool("production", false, "Production mode")

// Set up flags
flag.Parse()

func Flags() Flags {

}

