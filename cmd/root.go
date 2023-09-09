package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/siteverification/v1"
)

var (
	RootCmd                = newRootCmd()
	client                 *siteverification.Service
	siteType               string
	siteIdentifier         string
	siteVerificationMethod string
	siteOwners             string
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:               "gcloud-site-verify",
		Short:             "Google Cloud Site Verification CLI.",
		Long:              "A command-line utility for verifying and managing domain ownership using the Site Verification API.",
		Version:           "0.1.0",
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

func trimPrefix(s string) string {
	prefixes := []string{"http://", "https://", "dns://"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return strings.TrimPrefix(s, prefix)
		}
	}
	return s
}

func webResourceGet(id string) (response *siteverification.SiteVerificationWebResourceResource, err error) {
	getResp, err := client.WebResource.Get(id).Do()
	if err != nil {
		return nil, err
	}

	return getResp, err
}

func webResourceUpdate(owners []string) (responseBytes []byte, err error) {
	site := &siteverification.SiteVerificationWebResourceResourceSite{
		Identifier: trimPrefix(siteIdentifier),
		Type:       siteType,
	}

	updateResp, err := client.WebResource.Update(siteIdentifier, &siteverification.SiteVerificationWebResourceResource{
		Site:   site,
		Owners: owners,
	}).Do()
	if err != nil {
		return nil, err
	}

	responseBytes, err = updateResp.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return responseBytes, err
}

func weResourceInsert(owners []string) (responseBytes []byte, err error) {
	site := &siteverification.SiteVerificationWebResourceResourceSite{
		Identifier: trimPrefix(siteIdentifier),
		Type:       siteType,
	}

	insertResp, err := client.WebResource.Insert(siteVerificationMethod, &siteverification.SiteVerificationWebResourceResource{
		Owners: owners,
		Site:   site,
	}).Do()
	if err != nil {
		return nil, err
	}

	responseBytes, err = insertResp.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return responseBytes, err
}

func removeMembersFromList(members, membersToRemove []string) []string {
	result := []string{}

	// Create a map of items to remove for faster lookup
	membersToRemoveMap := make(map[string]bool)
	for _, item := range membersToRemove {
		membersToRemoveMap[item] = true
	}

	// Iterate through the original slice and add items not in the itemsToRemove map to the result slice
	for _, item := range members {
		if !membersToRemoveMap[item] {
			result = append(result, item)
		}
	}

	return result
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&siteType, "type", "t", "INET_DOMAIN", "The type of site to be verified. Valid values are ANDROID_APP, INET_DOMAIN or SITE.")
	RootCmd.PersistentFlags().StringVarP(&siteOwners, "owners", "o", "", "The list of owners to be added to the site. Separate multiple owners with a comma (e.g. foo@example.com,bar@example.com)")
	RootCmd.PersistentFlags().StringVarP(&siteIdentifier, "identifier", "i", "", "The identifier of the site to be verified (e.g. http://www.example.com/, https://www.example.com/, dns://example.com).")
	RootCmd.MarkPersistentFlagRequired("identifier")
	RootCmd.PersistentFlags().StringVarP(&siteVerificationMethod, "method", "m", "DNS_TXT", "The method to use for verifying a site. Check here for a list of valid methods: https://developers.google.com/site-verification/v1/getting_started#tokens")
}
