package client

import (
	"context"
	"fmt"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"log"
	"strings"
)

type AggregatedClient struct {
	BaseUrl          string
	UploadUrl        string
	EnterpriseClient *github.Client
	Ctx              context.Context
}

func GetGitEntClient(baseUrl, uploadUrl, token, tfVersion string) (*AggregatedClient, error) {
	ctx := context.Background()
	if strings.EqualFold(baseUrl, "") {
		return nil, fmt.Errorf("the base url is required")
	}

	if strings.EqualFold(uploadUrl, "") {
		return nil, fmt.Errorf("the upload url is required")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	enterpriseClient, err := github.NewEnterpriseClient(baseUrl, uploadUrl, tc)

	if err != nil {
		log.Printf("GetGitEntClient(): github.NewEnterpriseClient failed.")
		return nil, err
	}

	aggregatedClient := &AggregatedClient{
		BaseUrl:          baseUrl,
		UploadUrl:        uploadUrl,
		EnterpriseClient: enterpriseClient,
		Ctx:              ctx,
	}
	log.Printf("GetGitEntClient(): Created Enterprise client successfully!")
	return aggregatedClient, nil
}
