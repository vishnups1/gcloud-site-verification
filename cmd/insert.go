package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/siteverification/v1"
)

var (
	insertCmd = &cobra.Command{
		Use:   "insert",
		Short: "Verifies ownership of a website or domain.",
		Long:  "Verifies ownership of a website or domain. Before calling insert command, place the authenticated user's verification token on their website or domain.",
		RunE:  insert,
		Example: `gcloud-site-verify insert -i example.com -t INET_DOMAIN -m DNS_TXT -o "foo@example.com,bar@example.com"
gcloud-site-verify insert --identifier example.com --owners "foo@example.com,bar@example.com"`,
	}
)

func insert(cmd *cobra.Command, args []string) error {

	var owners []string
	if siteOwners != "" {
		for _, v := range strings.Split(siteOwners, ",") {
			owners = append(owners, strings.TrimSpace(v))
		}
	}

	insertResp, err := client.WebResource.Insert(siteVerificationMethod, &siteverification.SiteVerificationWebResourceResource{
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
	insertRespJSON, err := insertResp.MarshalJSON()
	if err != nil {
		return err
	}
	cmd.Println(string(insertRespJSON))

	return nil
}

func init() {
	RootCmd.AddCommand(insertCmd)
}
