package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2cIdentityUserFlow entity.
type B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderGetQueryParameters get userFlowIdentityProviders from identity
type B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderGetQueryParameters
}
// NewB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    m := &B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userFlowIdentityProviders/{identityProviderBase%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get userFlowIdentityProviders from identity
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityProviderBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseable), nil
}
// ToGetRequestInformation get userFlowIdentityProviders from identity
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) WithUrl(rawUrl string)(*B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    return NewB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
