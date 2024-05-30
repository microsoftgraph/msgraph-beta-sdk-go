package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ForwardingpoliciesForwardingPolicyItemRequestBuilder provides operations to manage the forwardingPolicies property of the microsoft.graph.networkaccess.networkAccessRoot entity.
type ForwardingpoliciesForwardingPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ForwardingpoliciesForwardingPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ForwardingpoliciesForwardingPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ForwardingpoliciesForwardingPolicyItemRequestBuilderGetQueryParameters retrieve information about a specific forwarding policy.
type ForwardingpoliciesForwardingPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ForwardingpoliciesForwardingPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ForwardingpoliciesForwardingPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ForwardingpoliciesForwardingPolicyItemRequestBuilderGetQueryParameters
}
// ForwardingpoliciesForwardingPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ForwardingpoliciesForwardingPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewForwardingpoliciesForwardingPolicyItemRequestBuilderInternal instantiates a new ForwardingpoliciesForwardingPolicyItemRequestBuilder and sets the default values.
func NewForwardingpoliciesForwardingPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForwardingpoliciesForwardingPolicyItemRequestBuilder) {
    m := &ForwardingpoliciesForwardingPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/forwardingPolicies/{forwardingPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewForwardingpoliciesForwardingPolicyItemRequestBuilder instantiates a new ForwardingpoliciesForwardingPolicyItemRequestBuilder and sets the default values.
func NewForwardingpoliciesForwardingPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForwardingpoliciesForwardingPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewForwardingpoliciesForwardingPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property forwardingPolicies for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ForwardingpoliciesForwardingPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve information about a specific forwarding policy.
// returns a ForwardingPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/networkaccess-forwardingpolicy-get?view=graph-rest-beta
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ForwardingpoliciesForwardingPolicyItemRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateForwardingPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingPolicyable), nil
}
// MicrosoftGraphNetworkaccessUpdatePolicyRules provides operations to call the updatePolicyRules method.
// returns a *ForwardingpoliciesItemMicrosoftgraphnetworkaccessupdatepolicyrulesMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder when successful
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) MicrosoftGraphNetworkaccessUpdatePolicyRules()(*ForwardingpoliciesItemMicrosoftgraphnetworkaccessupdatepolicyrulesMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder) {
    return NewForwardingpoliciesItemMicrosoftgraphnetworkaccessupdatepolicyrulesMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property forwardingPolicies in networkAccess
// returns a ForwardingPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingPolicyable, requestConfiguration *ForwardingpoliciesForwardingPolicyItemRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateForwardingPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingPolicyable), nil
}
// PolicyRules provides operations to manage the policyRules property of the microsoft.graph.networkaccess.policy entity.
// returns a *ForwardingpoliciesItemPolicyrulesPolicyRulesRequestBuilder when successful
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) PolicyRules()(*ForwardingpoliciesItemPolicyrulesPolicyRulesRequestBuilder) {
    return NewForwardingpoliciesItemPolicyrulesPolicyRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property forwardingPolicies for networkAccess
// returns a *RequestInformation when successful
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ForwardingpoliciesForwardingPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve information about a specific forwarding policy.
// returns a *RequestInformation when successful
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ForwardingpoliciesForwardingPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property forwardingPolicies in networkAccess
// returns a *RequestInformation when successful
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingPolicyable, requestConfiguration *ForwardingpoliciesForwardingPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ForwardingpoliciesForwardingPolicyItemRequestBuilder when successful
func (m *ForwardingpoliciesForwardingPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ForwardingpoliciesForwardingPolicyItemRequestBuilder) {
    return NewForwardingpoliciesForwardingPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
