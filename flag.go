package main
import "flag"
var nFlag = flag.Int("port", 8080, "the port listen on ")
//flag.Int(name,default,usage)
var user = flag.String("user","guest","username for login")
