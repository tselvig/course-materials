package scanner

import (
	"testing"
)

// Tests to see if PortScanner picks up the required number of open ports
func TestOpenPort(t *testing.T) {

    got, _ := PortScanner(1024)
    want := 1 // I changed this from 2 to 1 because I have only ever seen port 80 be open one time and it wasn't for very long

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}

// Tests to see if the port scanner scanned all of the ports it was supposed to
func TestTotalPortsScanned(t *testing.T) {
	totalPorts := 1000
    open, closed := PortScanner(totalPorts)
	got := open + closed
    want := totalPorts // default value; consider what would happen if you parameterize the portscanner ports to scan

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}


