package identity

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\identityProviders\{identityProvider-id}
type IdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilderInternal instantiates a new IdentityProviderItemRequestBuilder and sets the default values.
func NewIdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder) {
    m := &IdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder{
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
// NewIdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder instantiates a new IdentityProviderItemRequestBuilder and sets the default values.
func NewIdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityContainer entities.
func (m *IdentityB2cUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder) Ref()(*IdentityB2cUserFlowsItemIdentityProvidersItemRefRequestBuilder) {
    return NewIdentityB2cUserFlowsItemIdentityProvidersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
