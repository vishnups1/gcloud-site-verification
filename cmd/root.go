package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/api/siteverification/v1"
)

var (
	RootCmd            = newRootCmd()
	client             *siteverification.Service
	siteType           string
	siteIdentifier     string
	verificationMethod string
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:               "gcloud-site-verify",
		Short:             "Google Cloud Site Verification CLI.",
		Long:              "A command-line utility for verifying and managing domain ownership using the Site Verification API.",
		Version:           "0.0.1",
		PersistentPreRunE: setup,
	}
}

func setup(cmd *cobra.Command, args []string) error {
	var err error
	client, err = siteverification.NewService(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&siteType, "type", "t", "INET_DOMAIN", "The type of site to be verified. Valid values are ANDROID_APP, INET_DOMAIN or DOMAIN.")
	RootCmd.PersistentFlags().StringVarP(&siteIdentifier, "identifier", "i", "", "The identifier of the site to be verified (e.g. http://www.example.com/, https://www.example.com/, dns://example.com).")
	RootCmd.PersistentFlags().StringVarP(&verificationMethod, "method", "m", "DNS_TXT", "The method to use for verifying a site. Check here for a list of valid methods: https://developers.google.com/site-verification/v1/verifyingOwnership#methods")
}
