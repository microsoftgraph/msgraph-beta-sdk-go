package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder provides operations to call the translateExchangeIds method.
type ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderInternal instantiates a new ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder and sets the default values.
func NewItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) {
    m := &ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/translateExchangeIds", pathParameters),
    }
    return m
}
// NewItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder instantiates a new ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder and sets the default values.
func NewItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post translate identifiers of Outlook-related resources between formats.
// Deprecated: This method is obsolete. Use PostAsTranslateExchangeIdsPostResponse instead.
// returns a ItemTranslateexchangeidsTranslateExchangeIdsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-translateexchangeids?view=graph-rest-beta
func (m *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) Post(ctx context.Context, body ItemTranslateexchangeidsTranslateExchangeIdsPostRequestBodyable, requestConfiguration *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderPostRequestConfiguration)(ItemTranslateexchangeidsTranslateExchangeIdsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTranslateexchangeidsTranslateExchangeIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTranslateexchangeidsTranslateExchangeIdsResponseable), nil
}
// PostAsTranslateExchangeIdsPostResponse translate identifiers of Outlook-related resources between formats.
// returns a ItemTranslateexchangeidsTranslateExchangeIdsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-translateexchangeids?view=graph-rest-beta
func (m *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) PostAsTranslateExchangeIdsPostResponse(ctx context.Context, body ItemTranslateexchangeidsTranslateExchangeIdsPostRequestBodyable, requestConfiguration *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderPostRequestConfiguration)(ItemTranslateexchangeidsTranslateExchangeIdsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTranslateexchangeidsTranslateExchangeIdsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTranslateexchangeidsTranslateExchangeIdsPostResponseable), nil
}
// ToPostRequestInformation translate identifiers of Outlook-related resources between formats.
// returns a *RequestInformation when successful
func (m *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTranslateexchangeidsTranslateExchangeIdsPostRequestBodyable, requestConfiguration *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder when successful
func (m *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) WithUrl(rawUrl string)(*ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) {
    return NewItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
