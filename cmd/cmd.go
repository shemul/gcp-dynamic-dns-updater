package cmd

import (
	"fmt"
	"github.com/shemul/gcp-dynamic-dns-updater/gcp"
	"github.com/shemul/gcp-dynamic-dns-updater/ip"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func Run(app *cli.App) {
	app.Name = "GCP Dynamic DNS updater"
	app.Description = `
		Make sure you pass GOOGLE_APPLICATION_CREDENTIALS and GOOGLE_PROJECT env
		if you consider run in Docker, make sure you pass --net=host and restart policy as always
	`
	app.HideHelp = true
	app.UsageText = ""

	app.Commands = []cli.Command{
		{
			Name:  "update-dns",
			Usage: "Update DNS record for once using the current IP",
			Action: func(c *cli.Context) {
				updateDNS()
			},
		},
		{
			Name:  "list-dns",
			Usage: "Print the list of DNS with their Currennt IP",
			Action: func(c *cli.Context) {
				listDns()
			},
		},
		{
			Name:  "list-ip",
			Usage: "Print the current IP",
			Action: func(c *cli.Context) {
				listIP()
			},
		},
	}
	app.Run(os.Args)
}

func updateDNS() {
	names := paramDnsNames()
	project := paramGoogleProject()
	client := gcp.New(project)

	currentIP := readCurrentIP()
	handleChangedIP(currentIP, names, client)
}

func handleChangedIP(currentIP string, names []string, client *gcp.Client) {
	fmt.Printf("Current IP is %v\n", currentIP)

	fmt.Printf("Updating DNS records %v\n", names)
	records := readDnsRecords(client, names)
	updated := updateDnsRecords(client, records, []string{currentIP})

	if len(updated) == 0 {
		log.Println("Everything upto date!")
	} else {
		fmt.Printf("Updated %d DNS records:\n", len(updated))
		for _, record := range updated {
			fmt.Printf("    %s -> %v\n", record.Name, record.Rrdatas)
		}
	}
}

func listIP() {
	currentIP := readCurrentIP()
	fmt.Printf("Current IP is %v\n", currentIP)
}

func listDns() {
	names := paramDnsNames()
	project := paramGoogleProject()

	client := gcp.New(project)
	records := readDnsRecords(client, names)

	fmt.Printf("DNS %v\nIP %v\nTTL %v\n", records.NamesAndTypes()[0], records[0].Rrdatas[0], records[0].Ttl)
}

func paramDnsNames() []string {
	names := os.Getenv("DNS_NAMES")
	if names == "" {
		fmt.Printf("Environment variable DNS_NAMES not set.")
	}
	return strings.Split(names, " ")
}

func paramGoogleProject() string {
	project := os.Getenv("GOOGLE_PROJECT")
	if project == "" {
		fmt.Printf("Environment variable GOOGLE_PROJECT not set.")
	}
	return project
}

func readCurrentIP() string {
	var currentIP string
	var err error

	currentIP, err = ip.EgressIP()
	if err != nil {
		fmt.Printf("Failed to read current IP: ", err)
	}

	return currentIP
}

func readDnsRecords(client *gcp.Client, names []string) gcp.DnsRecords {
	records, err := client.DnsRecordsByNameAndType(names, "A")
	if err != nil {
		fmt.Printf("Failed to read DNS records: ", err)
	}
	return records
}

func updateDnsRecords(client *gcp.Client, records gcp.DnsRecords, newValues []string) gcp.DnsRecords {
	updated, err := client.UpdateDnsRecords(records, newValues)
	if err != nil {
		fmt.Printf("Failed to updateDNS DNS records: ", err)
	}
	return updated
}
