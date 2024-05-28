package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemFollowedsitesRemoveRequestBuilder provides operations to call the remove method.
type ItemFollowedsitesRemoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFollowedsitesRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFollowedsitesRemoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemFollowedsitesRemoveRequestBuilderInternal instantiates a new ItemFollowedsitesRemoveRequestBuilder and sets the default values.
func NewItemFollowedsitesRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowedsitesRemoveRequestBuilder) {
    m := &ItemFollowedsitesRemoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/followedSites/remove", pathParameters),
    }
    return m
}
// NewItemFollowedsitesRemoveRequestBuilder instantiates a new ItemFollowedsitesRemoveRequestBuilder and sets the default values.
func NewItemFollowedsitesRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowedsitesRemoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFollowedsitesRemoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unfollow a user's site or multiple sites.
// Deprecated: This method is obsolete. Use PostAsRemovePostResponse instead.
// returns a ItemFollowedsitesRemoveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-unfollow?view=graph-rest-beta
func (m *ItemFollowedsitesRemoveRequestBuilder) Post(ctx context.Context, body ItemFollowedsitesRemovePostRequestBodyable, requestConfiguration *ItemFollowedsitesRemoveRequestBuilderPostRequestConfiguration)(ItemFollowedsitesRemoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFollowedsitesRemoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFollowedsitesRemoveResponseable), nil
}
// PostAsRemovePostResponse unfollow a user's site or multiple sites.
// returns a ItemFollowedsitesRemovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-unfollow?view=graph-rest-beta
func (m *ItemFollowedsitesRemoveRequestBuilder) PostAsRemovePostResponse(ctx context.Context, body ItemFollowedsitesRemovePostRequestBodyable, requestConfiguration *ItemFollowedsitesRemoveRequestBuilderPostRequestConfiguration)(ItemFollowedsitesRemovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFollowedsitesRemovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFollowedsitesRemovePostResponseable), nil
}
// ToPostRequestInformation unfollow a user's site or multiple sites.
// returns a *RequestInformation when successful
func (m *ItemFollowedsitesRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemFollowedsitesRemovePostRequestBodyable, requestConfiguration *ItemFollowedsitesRemoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFollowedsitesRemoveRequestBuilder when successful
func (m *ItemFollowedsitesRemoveRequestBuilder) WithUrl(rawUrl string)(*ItemFollowedsitesRemoveRequestBuilder) {
    return NewItemFollowedsitesRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
