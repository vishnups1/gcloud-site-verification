package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	addownersCmd = &cobra.Command{
		Use:   "addowners",
		Short: "Adds owner(s) to a site or domain.",
		Long:  "Adds owner(s) to a site or domain. Before calling insert command, place the authenticated user's verification token on their website or domain.",
		RunE:  addowners,
		Example: `gcloud-site-verify addowners -i dns://example.com -t INET_DOMAIN -m DNS_TXT -o "foo@example.com,bar@example.com"
gcloud-site-verify addowners --identifier dns://example.com --owners "foo@example.com,bar@example.com"`,
	}
)

func addowners(cmd *cobra.Command, args []string) error {

	var owners []string
	if siteOwners != "" {
		for _, v := range strings.Split(siteOwners, ",") {
			owners = append(owners, strings.TrimSpace(v))
		}
	}

	resp, err := weResourceInsert(owners)
	if err != nil {
		return err
	}

	cmd.SetOut(cmd.OutOrStdout())
	cmd.Println(string(resp))
	return err
}

func init() {
	RootCmd.AddCommand(addownersCmd)
}
