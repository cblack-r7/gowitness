package utils

import (
	"github.com/tomsteele/go-nmap"
	"strconv"
	"strings"
)

// NmapUrls returns a slice of URL strings from a NmapRun object
func NmapUrls(nr nmap.NmapRun, svcs []string) ([]string, error) {
	var urls []string
	for _, host := range nr.Hosts {
		for _, port := range host.Ports {
			for _, service := range svcs {
				if strings.ToLower(service) == strings.ToLower(port.Service.Name) {
					for _, addr := range host.Addresses {
						urls = append(urls, addr.Addr+":"+strconv.Itoa(port.PortId))
					}
				}
			}
		}
	}
	return urls, nil
}
