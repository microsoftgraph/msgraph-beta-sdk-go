package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder provides operations to call the renewGroup method.
type ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderInternal instantiates a new ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) {
    m := &ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/renewGroup", pathParameters),
    }
    return m
}
// NewItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder instantiates a new ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy.
// Deprecated: This method is obsolete. Use PostAsRenewGroupPostResponse instead.
// returns a ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/grouplifecyclepolicy-renewgroup?view=graph-rest-beta
func (m *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) Post(ctx context.Context, body ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostRequestBodyable, requestConfiguration *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderPostRequestConfiguration)(ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGrouplifecyclepoliciesRenewgroupRenewGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponseable), nil
}
// PostAsRenewGroupPostResponse renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy.
// returns a ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/grouplifecyclepolicy-renewgroup?view=graph-rest-beta
func (m *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) PostAsRenewGroupPostResponse(ctx context.Context, body ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostRequestBodyable, requestConfiguration *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderPostRequestConfiguration)(ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponseable), nil
}
// ToPostRequestInformation renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy.
// returns a *RequestInformation when successful
func (m *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostRequestBodyable, requestConfiguration *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) WithUrl(rawUrl string)(*ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) {
    return NewItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
