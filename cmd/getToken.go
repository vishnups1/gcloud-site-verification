package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/api/siteverification/v1"
)

var (
	getTokenCmd = &cobra.Command{
		Use:   "getToken",
		Short: "Get a verification token.",
		Long:  "Gets a verification token for the authenticated user to place on a website or domain.",
		RunE:  getToken,
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
		VerificationMethod: verificationMethod,
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
