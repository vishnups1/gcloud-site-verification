package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

var (
	removeOwnersCmd = &cobra.Command{
		Use:   "removeowners",
		Short: "Removes owner(s) to a site or domain.",
		Long:  "Removes owner(s) to a site or domain. Before calling insert command, place the authenticated user's verification token on their website or domain.",
		RunE:  removeowners,
		Example: `gcloud-site-verify removeowners -i dns://example.com -t INET_DOMAIN -m DNS_TXT -o "foo@example.com,bar@example.com"
gcloud-site-verify removeowners --identifier dns://example.com --owners "foo@example.com,bar@example.com"`,
	}
)

func removeowners(cmd *cobra.Command, args []string) error {
	getResp, err := webResourceGet(siteIdentifier)
	if err != nil {
		return err
	}

	currentOwners := getResp.Owners

	var ownersToRemove []string
	if siteOwners != "" {
		for _, v := range strings.Split(siteOwners, ",") {
			ownersToRemove = append(ownersToRemove, strings.TrimSpace(v))
		}
	} else {
		return errors.New("required flag(s) \"owners\" not set.")
	}

	newOwners := removeMembersFromList(currentOwners, ownersToRemove)

	resp, err := webResourceUpdate(newOwners)
	if err != nil {
		return err
	}

	cmd.SetOut(cmd.OutOrStdout())
	cmd.Println(string(resp))
	return err
}

func init() {
	RootCmd.AddCommand(removeOwnersCmd)
}
