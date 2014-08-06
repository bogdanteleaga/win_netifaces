win_netifaces
=============

Windows ipconfig equivalent written in Go

Lists network interfaces on a system along with other attributes.
Uses go-ole to query WMI.

=============

Example usage:

package main

import (
        "fmt"

        "github.com/bogdan/win_netifaces"
)

func main() {
        interfaces, err := win_netifaces.GetAdapters(win_netifaces.Physical)
        if err != nil {
                fmt.Println(err)
        }
        for _, iface := range interfaces {
                fmt.Println("\n")
                win_netifaces.PrettyPrintInterface(iface)
                fmt.Println("\n")
        }
}
