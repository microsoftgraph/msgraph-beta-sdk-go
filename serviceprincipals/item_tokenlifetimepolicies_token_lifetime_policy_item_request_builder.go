package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
type ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetQueryParameters the tokenLifetimePolicies assigned to this service principal. Supports $expand.
type ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetQueryParameters
}
// NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal instantiates a new ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    m := &ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/tokenLifetimePolicies/{tokenLifetimePolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder instantiates a new ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the tokenLifetimePolicies assigned to this service principal. Supports $expand.
// returns a TokenLifetimePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TokenLifetimePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTokenLifetimePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TokenLifetimePolicyable), nil
}
// ToGetRequestInformation the tokenLifetimePolicies assigned to this service principal. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder when successful
func (m *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) WithUrl(rawUrl string)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    return NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
