package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates the list of owners for a website or domain.",
		Long:  "Updates the list of owners for a website or domain.",
		RunE:  update,
		Example: `gcloud-site-verify update -i dns://example.com -t INET_DOMAIN -m DNS_TXT -o 'foo@example.com,bar@example.com'
gcloud-site-verify update --identifier dns://example.com --owners 'foo@example.com,bar@example.com'`,
	}
)

func update(cmd *cobra.Command, args []string) error {

	var owners []string

	if siteOwners != "" {
		for _, v := range strings.Split(siteOwners, ",") {
			owners = append(owners, strings.TrimSpace(v))
		}
	} else {
		return errors.New("required flag(s) \"owners\" not set.")
	}

	resp, err := webResourceUpdate(owners)
	if err != nil {
		return err
	}

	cmd.SetOut(cmd.OutOrStdout())
	cmd.Println(string(resp))
	return err
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
