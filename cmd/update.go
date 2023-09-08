package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/siteverification/v1"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates the list of owners for a website or domain.",
		Long:  "Updates the list of owners for a website or domain.",
		RunE:  update,
		Example: `gcloud-site-verify update -i example.com -t INET_DOMAIN -m DNS_TXT -o 'foo@example.com,bar@example.com'
gcloud-site-verify update --identifier example.com --owners 'foo@example.com,bar@example.com'`,
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

	updateResp, err := client.WebResource.Update(addPrefix(siteIdentifier), &siteverification.SiteVerificationWebResourceResource{
		Owners: owners,
		Site: &siteverification.SiteVerificationWebResourceResourceSite{
			Identifier: trimPrefix(siteIdentifier),
			Type:       siteType,
		},
	}).Do()

	if err != nil {
		return err
	}

	cmd.SetOut(cmd.OutOrStdout())
	insertRespJSON, err := updateResp.MarshalJSON()
	if err != nil {
		return err
	}
	cmd.Println(string(insertRespJSON))

	return nil
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
