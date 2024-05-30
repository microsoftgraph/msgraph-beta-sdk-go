package identity

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\identityProviders\{identityProvider-id}
type B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewB2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilderInternal instantiates a new B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder and sets the default values.
func NewB2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder) {
    m := &B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/identityProviders/{identityProvider%2Did}", pathParameters),
    }
    return m
}
// NewB2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder instantiates a new B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder and sets the default values.
func NewB2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityContainer entities.
// returns a *B2cuserflowsItemIdentityprovidersItemRefRequestBuilder when successful
func (m *B2cuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder) Ref()(*B2cuserflowsItemIdentityprovidersItemRefRequestBuilder) {
    return NewB2cuserflowsItemIdentityprovidersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
