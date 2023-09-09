package cmd

import (
	"github.com/spf13/cobra"
)

var (
	getResourceCmd = &cobra.Command{
		Use:   "getresource",
		Short: "Retrieves the most current data for a website or domain.",
		Long:  "Retrieves the most current data for a website or domain.",
		RunE:  getresource,
		Example: `gcloud-site-verify getresource -i dns://example.com"
gcloud-site-verify getresource --identifier dns://example.com`,
	}
)

func getresource(cmd *cobra.Command, args []string) error {
	resp, err := webResourceGet(siteIdentifier)
	if err != nil {
		return err
	}

	getRespJSON, err := resp.MarshalJSON()
	if err != nil {
		return err
	}

	cmd.SetOut(cmd.OutOrStdout())
	cmd.Println(string(getRespJSON))
	return err
}

func init() {
	RootCmd.AddCommand(getResourceCmd)
}
