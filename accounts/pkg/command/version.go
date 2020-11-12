package command

import (
	"fmt"
	"os"

	"github.com/owncloud/ocis/ocis-pkg/registry"

	"github.com/micro/cli/v2"
	tw "github.com/olekukonko/tablewriter"
	"github.com/owncloud/ocis/accounts/pkg/config"
	"github.com/owncloud/ocis/accounts/pkg/flagset"
)

// PrintVersion prints the service versions of all running instances.
func PrintVersion(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Print the versions of the running instances",
		Flags: flagset.ListAccountsWithConfig(cfg),
		Action: func(c *cli.Context) error {
			reg := *registry.GetRegistry()
			services, err := reg.GetService(cfg.GRPC.Namespace + "." + cfg.Server.Name)
			if err != nil {
				fmt.Println(fmt.Errorf("could not get accounts services from the registry: %v", err))
				return err
			}

			if len(services) == 0 {
				fmt.Println("No running accounts service found.")
				return nil
			}

			table := tw.NewWriter(os.Stdout)
			table.SetHeader([]string{"Version", "Address", "Id"})
			table.SetAutoFormatHeaders(false)
			for _, s := range services {
				for _, n := range s.Nodes {
					table.Append([]string{s.Version, n.Address, n.Id})
				}
			}
			table.Render()
			return nil
		},
	}
}
