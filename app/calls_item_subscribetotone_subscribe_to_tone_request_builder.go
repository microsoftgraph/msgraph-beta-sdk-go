package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemSubscribetotoneSubscribeToToneRequestBuilder provides operations to call the subscribeToTone method.
type CallsItemSubscribetotoneSubscribeToToneRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemSubscribetotoneSubscribeToToneRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemSubscribetotoneSubscribeToToneRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemSubscribetotoneSubscribeToToneRequestBuilderInternal instantiates a new CallsItemSubscribetotoneSubscribeToToneRequestBuilder and sets the default values.
func NewCallsItemSubscribetotoneSubscribeToToneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemSubscribetotoneSubscribeToToneRequestBuilder) {
    m := &CallsItemSubscribetotoneSubscribeToToneRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/calls/{call%2Did}/subscribeToTone", pathParameters),
    }
    return m
}
// NewCallsItemSubscribetotoneSubscribeToToneRequestBuilder instantiates a new CallsItemSubscribetotoneSubscribeToToneRequestBuilder and sets the default values.
func NewCallsItemSubscribetotoneSubscribeToToneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemSubscribetotoneSubscribeToToneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemSubscribetotoneSubscribeToToneRequestBuilderInternal(urlParams, requestAdapter)
}
// Post subscribe to DTMF (dual-tone multi-frequency signaling) to allow you to be notified when the user presses keys on a dialpad. This action is supported only for calls that are initiated with serviceHostedMediaConfig.
// returns a SubscribeToToneOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-subscribetotone?view=graph-rest-beta
func (m *CallsItemSubscribetotoneSubscribeToToneRequestBuilder) Post(ctx context.Context, body CallsItemSubscribetotoneSubscribeToTonePostRequestBodyable, requestConfiguration *CallsItemSubscribetotoneSubscribeToToneRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubscribeToToneOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSubscribeToToneOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubscribeToToneOperationable), nil
}
// ToPostRequestInformation subscribe to DTMF (dual-tone multi-frequency signaling) to allow you to be notified when the user presses keys on a dialpad. This action is supported only for calls that are initiated with serviceHostedMediaConfig.
// returns a *RequestInformation when successful
func (m *CallsItemSubscribetotoneSubscribeToToneRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemSubscribetotoneSubscribeToTonePostRequestBodyable, requestConfiguration *CallsItemSubscribetotoneSubscribeToToneRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemSubscribetotoneSubscribeToToneRequestBuilder when successful
func (m *CallsItemSubscribetotoneSubscribeToToneRequestBuilder) WithUrl(rawUrl string)(*CallsItemSubscribetotoneSubscribeToToneRequestBuilder) {
    return NewCallsItemSubscribetotoneSubscribeToToneRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
