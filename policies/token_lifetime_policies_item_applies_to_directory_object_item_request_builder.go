package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.stsPolicy entity.
type TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderGetQueryParameters get appliesTo from policies
type TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewTokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal instantiates a new TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder and sets the default values.
func NewTokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder) {
    m := &TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/tokenLifetimePolicies/{tokenLifetimePolicy%2Did}/appliesTo/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder instantiates a new TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder and sets the default values.
func NewTokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get appliesTo from policies
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// ToGetRequestInformation get appliesTo from policies
// returns a *RequestInformation when successful
func (m *TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder when successful
func (m *TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*TokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder) {
    return NewTokenLifetimePoliciesItemAppliesToDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
