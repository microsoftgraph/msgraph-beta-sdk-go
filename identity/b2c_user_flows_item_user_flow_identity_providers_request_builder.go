package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2cIdentityUserFlow entity.
type B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetQueryParameters get userFlowIdentityProviders from identity
type B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetQueryParameters
}
// ByIdentityProviderBaseId provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2cIdentityUserFlow entity.
// returns a *B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder when successful
func (m *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) ByIdentityProviderBaseId(identityProviderBaseId string)(*B2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityProviderBaseId != "" {
        urlTplParams["identityProviderBase%2Did"] = identityProviderBaseId
    }
    return NewB2cUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal instantiates a new B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) {
    m := &B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userFlowIdentityProviders{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder instantiates a new B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *B2cUserFlowsItemUserFlowIdentityProvidersCountRequestBuilder when successful
func (m *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) Count()(*B2cUserFlowsItemUserFlowIdentityProvidersCountRequestBuilder) {
    return NewB2cUserFlowsItemUserFlowIdentityProvidersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get userFlowIdentityProviders from identity
// returns a IdentityProviderBaseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityProviderBaseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProviderBaseCollectionResponseable), nil
}
// ToGetRequestInformation get userFlowIdentityProviders from identity
// returns a *RequestInformation when successful
func (m *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder when successful
func (m *B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) WithUrl(rawUrl string)(*B2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder) {
    return NewB2cUserFlowsItemUserFlowIdentityProvidersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
