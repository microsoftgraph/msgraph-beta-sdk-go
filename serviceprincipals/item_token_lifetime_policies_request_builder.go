package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTokenLifetimePoliciesRequestBuilder provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
type ItemTokenLifetimePoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTokenLifetimePoliciesRequestBuilderGetQueryParameters list the tokenLifetimePolicy objects that are assigned to a servicePrincipal. Only one object is returned in the collection because only one tokenLifetimePolicy can be assigned to a service principal.
type ItemTokenLifetimePoliciesRequestBuilderGetQueryParameters struct {
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
// ItemTokenLifetimePoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTokenLifetimePoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTokenLifetimePoliciesRequestBuilderGetQueryParameters
}
// ByTokenLifetimePolicyId provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder when successful
func (m *ItemTokenLifetimePoliciesRequestBuilder) ByTokenLifetimePolicyId(tokenLifetimePolicyId string)(*ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tokenLifetimePolicyId != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = tokenLifetimePolicyId
    }
    return NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTokenLifetimePoliciesRequestBuilderInternal instantiates a new ItemTokenLifetimePoliciesRequestBuilder and sets the default values.
func NewItemTokenLifetimePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenLifetimePoliciesRequestBuilder) {
    m := &ItemTokenLifetimePoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/tokenLifetimePolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTokenLifetimePoliciesRequestBuilder instantiates a new ItemTokenLifetimePoliciesRequestBuilder and sets the default values.
func NewItemTokenLifetimePoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenLifetimePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenLifetimePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTokenLifetimePoliciesCountRequestBuilder when successful
func (m *ItemTokenLifetimePoliciesRequestBuilder) Count()(*ItemTokenLifetimePoliciesCountRequestBuilder) {
    return NewItemTokenLifetimePoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the tokenLifetimePolicy objects that are assigned to a servicePrincipal. Only one object is returned in the collection because only one tokenLifetimePolicy can be assigned to a service principal.
// returns a TokenLifetimePolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-list-tokenlifetimepolicies?view=graph-rest-beta
func (m *ItemTokenLifetimePoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTokenLifetimePoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TokenLifetimePolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTokenLifetimePolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TokenLifetimePolicyCollectionResponseable), nil
}
// ToGetRequestInformation list the tokenLifetimePolicy objects that are assigned to a servicePrincipal. Only one object is returned in the collection because only one tokenLifetimePolicy can be assigned to a service principal.
// returns a *RequestInformation when successful
func (m *ItemTokenLifetimePoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTokenLifetimePoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTokenLifetimePoliciesRequestBuilder when successful
func (m *ItemTokenLifetimePoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemTokenLifetimePoliciesRequestBuilder) {
    return NewItemTokenLifetimePoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
