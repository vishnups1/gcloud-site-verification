package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/siteverification/v1"
)

var (
	addownersCmd = &cobra.Command{
		Use:   "addowners",
		Short: "Adds owner(s) to a site or domain.",
		Long:  "Adds owner(s) to a site or domain. Before calling insert command, place the authenticated user's verification token on their website or domain.",
		RunE:  addowners,
		Example: `gcloud-site-verify addowners -i example.com -t INET_DOMAIN -m DNS_TXT -o "foo@example.com,bar@example.com"
gcloud-site-verify addowners --identifier example.com --owners "foo@example.com,bar@example.com"`,
	}
)

func addowners(cmd *cobra.Command, args []string) error {

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
	RootCmd.AddCommand(addownersCmd)
}
