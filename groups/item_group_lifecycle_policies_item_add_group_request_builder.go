package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder provides operations to call the addGroup method.
type ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupLifecyclePoliciesItemAddGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupLifecyclePoliciesItemAddGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGroupLifecyclePoliciesItemAddGroupRequestBuilderInternal instantiates a new ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemAddGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder) {
    m := &ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/{groupLifecyclePolicy%2Did}/addGroup", pathParameters),
    }
    return m
}
// NewItemGroupLifecyclePoliciesItemAddGroupRequestBuilder instantiates a new ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemAddGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupLifecyclePoliciesItemAddGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addGroup
// Deprecated: This method is obsolete. Use PostAsAddGroupPostResponse instead.
// returns a ItemGroupLifecyclePoliciesItemAddGroupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder) Post(ctx context.Context, body ItemGroupLifecyclePoliciesItemAddGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilderPostRequestConfiguration)(ItemGroupLifecyclePoliciesItemAddGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGroupLifecyclePoliciesItemAddGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGroupLifecyclePoliciesItemAddGroupResponseable), nil
}
// PostAsAddGroupPostResponse invoke action addGroup
// returns a ItemGroupLifecyclePoliciesItemAddGroupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder) PostAsAddGroupPostResponse(ctx context.Context, body ItemGroupLifecyclePoliciesItemAddGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilderPostRequestConfiguration)(ItemGroupLifecyclePoliciesItemAddGroupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGroupLifecyclePoliciesItemAddGroupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGroupLifecyclePoliciesItemAddGroupPostResponseable), nil
}
// ToPostRequestInformation invoke action addGroup
// returns a *RequestInformation when successful
func (m *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGroupLifecyclePoliciesItemAddGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder when successful
func (m *ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder) WithUrl(rawUrl string)(*ItemGroupLifecyclePoliciesItemAddGroupRequestBuilder) {
    return NewItemGroupLifecyclePoliciesItemAddGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
