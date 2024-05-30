package grouplifecyclepolicies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRemovegroupRemoveGroupRequestBuilder provides operations to call the removeGroup method.
type ItemRemovegroupRemoveGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRemovegroupRemoveGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemovegroupRemoveGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRemovegroupRemoveGroupRequestBuilderInternal instantiates a new ItemRemovegroupRemoveGroupRequestBuilder and sets the default values.
func NewItemRemovegroupRemoveGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovegroupRemoveGroupRequestBuilder) {
    m := &ItemRemovegroupRemoveGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groupLifecyclePolicies/{groupLifecyclePolicy%2Did}/removeGroup", pathParameters),
    }
    return m
}
// NewItemRemovegroupRemoveGroupRequestBuilder instantiates a new ItemRemovegroupRemoveGroupRequestBuilder and sets the default values.
func NewItemRemovegroupRemoveGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovegroupRemoveGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemovegroupRemoveGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action removeGroup
// Deprecated: This method is obsolete. Use PostAsRemoveGroupPostResponse instead.
// returns a ItemRemovegroupRemoveGroupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRemovegroupRemoveGroupRequestBuilder) Post(ctx context.Context, body ItemRemovegroupRemoveGroupPostRequestBodyable, requestConfiguration *ItemRemovegroupRemoveGroupRequestBuilderPostRequestConfiguration)(ItemRemovegroupRemoveGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRemovegroupRemoveGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRemovegroupRemoveGroupResponseable), nil
}
// PostAsRemoveGroupPostResponse invoke action removeGroup
// returns a ItemRemovegroupRemoveGroupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRemovegroupRemoveGroupRequestBuilder) PostAsRemoveGroupPostResponse(ctx context.Context, body ItemRemovegroupRemoveGroupPostRequestBodyable, requestConfiguration *ItemRemovegroupRemoveGroupRequestBuilderPostRequestConfiguration)(ItemRemovegroupRemoveGroupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRemovegroupRemoveGroupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRemovegroupRemoveGroupPostResponseable), nil
}
// ToPostRequestInformation invoke action removeGroup
// returns a *RequestInformation when successful
func (m *ItemRemovegroupRemoveGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemovegroupRemoveGroupPostRequestBodyable, requestConfiguration *ItemRemovegroupRemoveGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRemovegroupRemoveGroupRequestBuilder when successful
func (m *ItemRemovegroupRemoveGroupRequestBuilder) WithUrl(rawUrl string)(*ItemRemovegroupRemoveGroupRequestBuilder) {
    return NewItemRemovegroupRemoveGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
