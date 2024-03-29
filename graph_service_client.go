package msgraphbetasdkgo

import (
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/store"
	icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users"
	az "github.com/microsoftgraph/msgraph-sdk-go-core/authentication"
)

type GraphServiceClient struct {
	GraphBaseServiceClient
}

func NewGraphServiceClient(adapter abstractions.RequestAdapter) *GraphServiceClient {
	client := NewGraphBaseServiceClient(adapter, store.BackingStoreFactoryInstance)
	return &GraphServiceClient{
		*client,
	}
}

// NewGraphServiceClientWithCredentials instantiates a new GraphServiceClient with provided credentials and scopes
func NewGraphServiceClientWithCredentials(credential azcore.TokenCredential, scopes []string) (*GraphServiceClient, error) {
	return NewGraphServiceClientWithCredentialsAndHosts(credential, scopes, nil)
}

// NewGraphServiceClientWithCredentialsAndHosts instantiates a new GraphServiceClient with provided credentials , scopes and validhosts
func NewGraphServiceClientWithCredentialsAndHosts(credential azcore.TokenCredential, scopes []string, validhosts []string) (*GraphServiceClient, error) {
	if credential == nil {
		return nil, errors.New("credential cannot be nil")
	}

	if validhosts == nil || len(validhosts) == 0 {
		validhosts = []string{"graph.microsoft.com", "graph.microsoft.us", "dod-graph.microsoft.us", "graph.microsoft.de", "microsoftgraph.chinacloudapi.cn", "canary.graph.microsoft.com"}
	}

	if scopes == nil || len(scopes) == 0 {
		scopes = []string{"https://graph.microsoft.com/.default"}
	}

	auth, err := az.NewAzureIdentityAuthenticationProviderWithScopesAndValidHosts(credential, scopes, validhosts)
	if err != nil {
		return nil, err
	}

	adapter, err := NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, err
	}

	client := NewGraphServiceClient(adapter)
	return client, nil
}

// GetAdapter returns the client current adapter, Method should only be called when the user is certain an adapter has been provided
func (m *GraphBaseServiceClient) GetAdapter() abstractions.RequestAdapter {
	if m.BaseRequestBuilder.RequestAdapter == nil {
		panic(errors.New("request adapter has not been initialized"))
	}
	return m.BaseRequestBuilder.RequestAdapter
}

// Me provides operations to manage the user singleton.
func (m *GraphBaseServiceClient) Me() *icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UserItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["user%2Did"] = "me-token-to-replace"
	return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
