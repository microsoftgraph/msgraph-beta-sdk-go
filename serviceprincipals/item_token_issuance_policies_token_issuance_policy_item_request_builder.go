package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
type ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderGetQueryParameters the tokenIssuancePolicies assigned to this service principal. Supports $expand.
type ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderGetQueryParameters
}
// NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal instantiates a new ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder and sets the default values.
func NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    m := &ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/tokenIssuancePolicies/{tokenIssuancePolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder instantiates a new ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder and sets the default values.
func NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the tokenIssuancePolicies assigned to this service principal. Supports $expand.
// returns a TokenIssuancePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TokenIssuancePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTokenIssuancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TokenIssuancePolicyable), nil
}
// ToGetRequestInformation the tokenIssuancePolicies assigned to this service principal. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder when successful
func (m *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) WithUrl(rawUrl string)(*ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    return NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
