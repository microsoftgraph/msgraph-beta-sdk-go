package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i38206d39ea5f422af2e9cd600d589c67ed3afc8b6dac7ef86b13c2c78850bddf "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/identityproviders/item/ref"
)

// IdentityProviderItemRequestBuilder builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\identityProviders\{identityProvider-id}
type IdentityProviderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdentityProviderItemRequestBuilderInternal instantiates a new IdentityProviderItemRequestBuilder and sets the default values.
func NewIdentityProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProviderItemRequestBuilder) {
    m := &IdentityProviderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/identityProviders/{identityProvider%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityProviderItemRequestBuilder instantiates a new IdentityProviderItemRequestBuilder and sets the default values.
func NewIdentityProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the ref property
func (m *IdentityProviderItemRequestBuilder) Ref()(*i38206d39ea5f422af2e9cd600d589c67ed3afc8b6dac7ef86b13c2c78850bddf.RefRequestBuilder) {
    return i38206d39ea5f422af2e9cd600d589c67ed3afc8b6dac7ef86b13c2c78850bddf.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
