package main 

import (
    log "github.com/Sirupsen/logrus"
    "traintracks/lib1"
    "traintracks/lib2"
)

func main() {
    log.Println("I'm a cool app and I'm " +
                "gonna add together the results of two functions " +
                "in two imported libraries " +
                "and then use a library downloaded from the " +
                "internet to log the result ")
     log.Println("Result:", lib1.One() + lib2.Two())
}
