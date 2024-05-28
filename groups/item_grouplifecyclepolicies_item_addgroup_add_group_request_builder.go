package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder provides operations to call the addGroup method.
type ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderInternal instantiates a new ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) {
    m := &ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/{groupLifecyclePolicy%2Did}/addGroup", pathParameters),
    }
    return m
}
// NewItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder instantiates a new ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addGroup
// Deprecated: This method is obsolete. Use PostAsAddGroupPostResponse instead.
// returns a ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) Post(ctx context.Context, body ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostRequestBodyable, requestConfiguration *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderPostRequestConfiguration)(ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGrouplifecyclepoliciesItemAddgroupAddGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponseable), nil
}
// PostAsAddGroupPostResponse invoke action addGroup
// returns a ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) PostAsAddGroupPostResponse(ctx context.Context, body ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostRequestBodyable, requestConfiguration *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderPostRequestConfiguration)(ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponseable), nil
}
// ToPostRequestInformation invoke action addGroup
// returns a *RequestInformation when successful
func (m *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostRequestBodyable, requestConfiguration *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) WithUrl(rawUrl string)(*ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) {
    return NewItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
