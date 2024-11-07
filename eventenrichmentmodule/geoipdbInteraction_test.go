package eventenrichmentmodule_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/eventenrichmentmodule"
)

func groupIpInfoResult(infoEvent datamodels.InformationFromEventEnricher) struct{ city, country, countryCode string } {
	sources := [...]string{"GeoipNoc", "MAXMIND", "DBIP", "AriadnaDB"}
	customIpResult := struct{ city, country, countryCode string }{}

	for _, ip := range infoEvent.GetIpAddresses() {
		for _, source := range sources {
			if city, ok := infoEvent.SearchCity(ip, source); ok && city != "" {
				if customIpResult.city != "" {
					continue
				}

				customIpResult.city = city
			}

			if country, ok := infoEvent.SearchCountry(ip, source); ok && country != "" {
				if customIpResult.country != "" {
					continue
				}

				customIpResult.country = country
			}

			if countryCode, ok := infoEvent.SearchCountryCode(ip, source); ok && countryCode != "" {
				if customIpResult.countryCode != "" {
					continue
				}

				customIpResult.countryCode = countryCode
			}
		}
	}

	return customIpResult
}

var ipAddresses = [...]string{"104.16.167.228", "82.221.129.24", "192.168.152.1", "56.14.66.100:8954", "82.157.247.165"}

func TestGeoIpInteraction(t *testing.T) {
	settingsResponse := eventenrichmentmodule.SettingsChanOutputEEM{
		IpAddresses:     []string(nil),
		IpAddressesInfo: []eventenrichmentmodule.GeoIpInformation(nil),
	}

	client, err := eventenrichmentmodule.NewGeoIpClient(
		eventenrichmentmodule.WithHost("pg2.cloud.gcm"),
		eventenrichmentmodule.WithPort(88),
		eventenrichmentmodule.WithPath("ip"),
		eventenrichmentmodule.WithConnectionTimeout(10*time.Second))
	assert.NoError(t, err)

	ctxTimeout, ctxCancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer ctxCancel()

	for _, ip := range ipAddresses {
		fmt.Println("----==== ip", ip, " ====----")

		res, err := client.GetGeoInformation(ctxTimeout, ip)
		assert.NoError(t, err)

		//resByte, err := json.MarshalIndent(res, "", " ")
		//assert.NoError(t, err)
		//fmt.Println("----- GeoIP Info -----")
		//fmt.Println(string(resByte))

		settingsResponse.IpAddresses = append(settingsResponse.IpAddresses, ip)
		settingsResponse.IpAddressesInfo = append(settingsResponse.IpAddressesInfo, res)

		customIpInfo := groupIpInfoResult(settingsResponse)

		fmt.Println("city:", customIpInfo.city)
		fmt.Println("country:", customIpInfo.country)
		fmt.Println("country code:", customIpInfo.countryCode)
		fmt.Println("")
	}

	assert.True(t, true)
}
