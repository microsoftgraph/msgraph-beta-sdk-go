package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder provides operations to call the updatePolicyRules method.
type ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderInternal instantiates a new ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder and sets the default values.
func NewForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder) {
    m := &ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/forwardingPolicies/{forwardingPolicy%2Did}/microsoft.graph.networkaccess.updatePolicyRules", pathParameters),
    }
    return m
}
// NewForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder instantiates a new ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder and sets the default values.
func NewForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the rules within a forwarding policy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/networkaccess-forwardingpolicy-updatepolicyrules?view=graph-rest-1.0
func (m *ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder) Post(ctx context.Context, body ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesUpdatePolicyRulesPostRequestBodyable, requestConfiguration *ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation update the rules within a forwarding policy.
// returns a *RequestInformation when successful
func (m *ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesUpdatePolicyRulesPostRequestBodyable, requestConfiguration *ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder when successful
func (m *ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder) WithUrl(rawUrl string)(*ForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder) {
    return NewForwardingPoliciesItemMicrosoftGraphNetworkaccessUpdatePolicyRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
