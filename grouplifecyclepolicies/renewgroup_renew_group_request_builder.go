package grouplifecyclepolicies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RenewgroupRenewGroupRequestBuilder provides operations to call the renewGroup method.
type RenewgroupRenewGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RenewgroupRenewGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RenewgroupRenewGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRenewgroupRenewGroupRequestBuilderInternal instantiates a new RenewgroupRenewGroupRequestBuilder and sets the default values.
func NewRenewgroupRenewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RenewgroupRenewGroupRequestBuilder) {
    m := &RenewgroupRenewGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groupLifecyclePolicies/renewGroup", pathParameters),
    }
    return m
}
// NewRenewgroupRenewGroupRequestBuilder instantiates a new RenewgroupRenewGroupRequestBuilder and sets the default values.
func NewRenewgroupRenewGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RenewgroupRenewGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRenewgroupRenewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy.
// Deprecated: This method is obsolete. Use PostAsRenewGroupPostResponse instead.
// returns a RenewgroupRenewGroupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/grouplifecyclepolicy-renewgroup?view=graph-rest-beta
func (m *RenewgroupRenewGroupRequestBuilder) Post(ctx context.Context, body RenewgroupRenewGroupPostRequestBodyable, requestConfiguration *RenewgroupRenewGroupRequestBuilderPostRequestConfiguration)(RenewgroupRenewGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRenewgroupRenewGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RenewgroupRenewGroupResponseable), nil
}
// PostAsRenewGroupPostResponse renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy.
// returns a RenewgroupRenewGroupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/grouplifecyclepolicy-renewgroup?view=graph-rest-beta
func (m *RenewgroupRenewGroupRequestBuilder) PostAsRenewGroupPostResponse(ctx context.Context, body RenewgroupRenewGroupPostRequestBodyable, requestConfiguration *RenewgroupRenewGroupRequestBuilderPostRequestConfiguration)(RenewgroupRenewGroupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRenewgroupRenewGroupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RenewgroupRenewGroupPostResponseable), nil
}
// ToPostRequestInformation renew a group's expiration. When a group is renewed, the group expiration is extended by the number of days defined in the policy.
// returns a *RequestInformation when successful
func (m *RenewgroupRenewGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body RenewgroupRenewGroupPostRequestBodyable, requestConfiguration *RenewgroupRenewGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RenewgroupRenewGroupRequestBuilder when successful
func (m *RenewgroupRenewGroupRequestBuilder) WithUrl(rawUrl string)(*RenewgroupRenewGroupRequestBuilder) {
    return NewRenewgroupRenewGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
