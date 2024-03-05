package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemSitesRemoveRequestBuilder provides operations to call the remove method.
type ItemSitesItemSitesRemoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemSitesRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemSitesRemoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemSitesRemoveRequestBuilderInternal instantiates a new ItemSitesItemSitesRemoveRequestBuilder and sets the default values.
func NewItemSitesItemSitesRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemSitesRemoveRequestBuilder) {
    m := &ItemSitesItemSitesRemoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/sites/remove", pathParameters),
    }
    return m
}
// NewItemSitesItemSitesRemoveRequestBuilder instantiates a new ItemSitesItemSitesRemoveRequestBuilder and sets the default values.
func NewItemSitesItemSitesRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemSitesRemoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemSitesRemoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unfollow a user's site or multiple sites.
// Deprecated: This method is obsolete. Use PostAsRemovePostResponse instead.
// returns a ItemSitesItemSitesRemoveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-unfollow?view=graph-rest-1.0
func (m *ItemSitesItemSitesRemoveRequestBuilder) Post(ctx context.Context, body ItemSitesItemSitesRemovePostRequestBodyable, requestConfiguration *ItemSitesItemSitesRemoveRequestBuilderPostRequestConfiguration)(ItemSitesItemSitesRemoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemSitesRemoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemSitesRemoveResponseable), nil
}
// PostAsRemovePostResponse unfollow a user's site or multiple sites.
// returns a ItemSitesItemSitesRemovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-unfollow?view=graph-rest-1.0
func (m *ItemSitesItemSitesRemoveRequestBuilder) PostAsRemovePostResponse(ctx context.Context, body ItemSitesItemSitesRemovePostRequestBodyable, requestConfiguration *ItemSitesItemSitesRemoveRequestBuilderPostRequestConfiguration)(ItemSitesItemSitesRemovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemSitesRemovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemSitesRemovePostResponseable), nil
}
// ToPostRequestInformation unfollow a user's site or multiple sites.
// returns a *RequestInformation when successful
func (m *ItemSitesItemSitesRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemSitesRemovePostRequestBodyable, requestConfiguration *ItemSitesItemSitesRemoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemSitesRemoveRequestBuilder when successful
func (m *ItemSitesItemSitesRemoveRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemSitesRemoveRequestBuilder) {
    return NewItemSitesItemSitesRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
