package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGroupLifecyclePoliciesRenewGroupRequestBuilder provides operations to call the renewGroup method.
type ItemGroupLifecyclePoliciesRenewGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupLifecyclePoliciesRenewGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupLifecyclePoliciesRenewGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGroupLifecyclePoliciesRenewGroupRequestBuilderInternal instantiates a new RenewGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesRenewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesRenewGroupRequestBuilder) {
    m := &ItemGroupLifecyclePoliciesRenewGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/renewGroup", pathParameters),
    }
    return m
}
// NewItemGroupLifecyclePoliciesRenewGroupRequestBuilder instantiates a new RenewGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesRenewGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesRenewGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupLifecyclePoliciesRenewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy. This API is available in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsRenewGroupPostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/grouplifecyclepolicy-renewgroup?view=graph-rest-1.0
func (m *ItemGroupLifecyclePoliciesRenewGroupRequestBuilder) Post(ctx context.Context, body ItemGroupLifecyclePoliciesRenewGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesRenewGroupRequestBuilderPostRequestConfiguration)(ItemGroupLifecyclePoliciesRenewGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGroupLifecyclePoliciesRenewGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGroupLifecyclePoliciesRenewGroupResponseable), nil
}
// PostAsRenewGroupPostResponse renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/grouplifecyclepolicy-renewgroup?view=graph-rest-1.0
func (m *ItemGroupLifecyclePoliciesRenewGroupRequestBuilder) PostAsRenewGroupPostResponse(ctx context.Context, body ItemGroupLifecyclePoliciesRenewGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesRenewGroupRequestBuilderPostRequestConfiguration)(ItemGroupLifecyclePoliciesRenewGroupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGroupLifecyclePoliciesRenewGroupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGroupLifecyclePoliciesRenewGroupPostResponseable), nil
}
// ToPostRequestInformation renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy. This API is available in the following national cloud deployments.
func (m *ItemGroupLifecyclePoliciesRenewGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGroupLifecyclePoliciesRenewGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesRenewGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemGroupLifecyclePoliciesRenewGroupRequestBuilder) WithUrl(rawUrl string)(*ItemGroupLifecyclePoliciesRenewGroupRequestBuilder) {
    return NewItemGroupLifecyclePoliciesRenewGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
