package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder provides operations to manage the federatedTokenValidationPolicy property of the microsoft.graph.policyRoot entity.
type FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderGetQueryParameters get a list of the federatedTokenValidationPolicy objects and their properties.
type FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderGetQueryParameters
}
// FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderInternal instantiates a new FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder and sets the default values.
func NewFederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) {
    m := &FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/federatedTokenValidationPolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder instantiates a new FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder and sets the default values.
func NewFederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property federatedTokenValidationPolicy for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) Delete(ctx context.Context, requestConfiguration *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get a list of the federatedTokenValidationPolicy objects and their properties.
// returns a FederatedTokenValidationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/policyroot-list-federatedtokenvalidationpolicy?view=graph-rest-beta
func (m *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedTokenValidationPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFederatedTokenValidationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedTokenValidationPolicyable), nil
}
// Patch update the properties of a federatedTokenValidationPolicy object.
// returns a FederatedTokenValidationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/federatedtokenvalidationpolicy-update?view=graph-rest-beta
func (m *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedTokenValidationPolicyable, requestConfiguration *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedTokenValidationPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFederatedTokenValidationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedTokenValidationPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property federatedTokenValidationPolicy for policies
// returns a *RequestInformation when successful
func (m *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get a list of the federatedTokenValidationPolicy objects and their properties.
// returns a *RequestInformation when successful
func (m *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a federatedTokenValidationPolicy object.
// returns a *RequestInformation when successful
func (m *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FederatedTokenValidationPolicyable, requestConfiguration *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder when successful
func (m *FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) WithUrl(rawUrl string)(*FederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder) {
    return NewFederatedtokenvalidationpolicyFederatedTokenValidationPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
