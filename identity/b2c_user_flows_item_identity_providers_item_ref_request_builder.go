package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder provides operations to manage the collection of identityContainer entities.
type B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cUserFlowsItemIdentityProvidersItemRefRequestBuilderDeleteQueryParameters delete an identity provider from a b2cIdentityUserFlow object. For more information about identity providers available for user flows, see the identityProviders API reference.
type B2cUserFlowsItemIdentityProvidersItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// B2cUserFlowsItemIdentityProvidersItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemIdentityProvidersItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cUserFlowsItemIdentityProvidersItemRefRequestBuilderDeleteQueryParameters
}
// NewB2cUserFlowsItemIdentityProvidersItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewB2cUserFlowsItemIdentityProvidersItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder) {
    m := &B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/identityProviders/{identityProvider%2Did}/$ref{?%40id*}", pathParameters),
    }
    return m
}
// NewB2cUserFlowsItemIdentityProvidersItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewB2cUserFlowsItemIdentityProvidersItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemIdentityProvidersItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an identity provider from a b2cIdentityUserFlow object. For more information about identity providers available for user flows, see the identityProviders API reference.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider on 2021-08-24 and will be removed 2023-03-15
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/b2cidentityuserflow-delete-identityproviders?view=graph-rest-1.0
func (m *B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2cUserFlowsItemIdentityProvidersItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation delete an identity provider from a b2cIdentityUserFlow object. For more information about identity providers available for user flows, see the identityProviders API reference.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider on 2021-08-24 and will be removed 2023-03-15
func (m *B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemIdentityProvidersItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider on 2021-08-24 and will be removed 2023-03-15
func (m *B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder) WithUrl(rawUrl string)(*B2cUserFlowsItemIdentityProvidersItemRefRequestBuilder) {
    return NewB2cUserFlowsItemIdentityProvidersItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
