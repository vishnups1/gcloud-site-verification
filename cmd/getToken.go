package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/api/siteverification/v1"
)

var (
	getTokenCmd = &cobra.Command{
		Use:   "gettoken",
		Short: "Gets a verification token for the authenticated user to place on a website or domain.",
		Long:  "Gets a verification token for the authenticated user to place on a website or domain.",
		RunE:  getToken,
		Example: `gcloud-site-verify gettoken -i example.com -t INET_DOMAIN -m DNS_TXT
gcloud-site-verify gettoken --identifier example.com`,
	}
)

func getToken(cmd *cobra.Command, args []string) error {
	if siteIdentifier == "" {
		return errors.New("required flag(s) \"identifier\" not set.")
	}

	getTokenRequest := &siteverification.SiteVerificationWebResourceGettokenRequest{
		Site: &siteverification.SiteVerificationWebResourceGettokenRequestSite{
			Identifier: siteIdentifier,
			Type:       siteType,
		},
		VerificationMethod: siteVerificationMethod,
	}

	getTokenResp, err := client.WebResource.GetToken(getTokenRequest).Do()
	if err != nil {
		return err
	}

	tokenBytes, err := getTokenResp.MarshalJSON()
	if err != nil {
		return err
	}

	cmd.SetOut(os.Stdout)
	cmd.Println(string(tokenBytes))

	return nil
}

func init() {
	RootCmd.AddCommand(getTokenCmd)
}
