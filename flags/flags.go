package flags

import "flag"

var sesUsername = flag.String("username", "", "Username credential")
var sesPassword = flag.String("password", "", "Password credential")
// Set up flags
flag.Parse()

