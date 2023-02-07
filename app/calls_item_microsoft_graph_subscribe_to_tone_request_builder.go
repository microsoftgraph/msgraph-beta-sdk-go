package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemMicrosoftGraphSubscribeToToneRequestBuilder provides operations to call the subscribeToTone method.
type CallsItemMicrosoftGraphSubscribeToToneRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsItemMicrosoftGraphSubscribeToToneRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemMicrosoftGraphSubscribeToToneRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemMicrosoftGraphSubscribeToToneRequestBuilderInternal instantiates a new MicrosoftGraphSubscribeToToneRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphSubscribeToToneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphSubscribeToToneRequestBuilder) {
    m := &CallsItemMicrosoftGraphSubscribeToToneRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/calls/{call%2Did}/microsoft.graph.subscribeToTone";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsItemMicrosoftGraphSubscribeToToneRequestBuilder instantiates a new MicrosoftGraphSubscribeToToneRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphSubscribeToToneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphSubscribeToToneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemMicrosoftGraphSubscribeToToneRequestBuilderInternal(urlParams, requestAdapter)
}
// Post subscribe to DTMF (dual-tone multi-frequency signaling). This allows you to be notified when the user presses keys on a 'Dialpad'.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/call-subscribetotone?view=graph-rest-1.0
func (m *CallsItemMicrosoftGraphSubscribeToToneRequestBuilder) Post(ctx context.Context, body CallsItemMicrosoftGraphSubscribeToToneSubscribeToTonePostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphSubscribeToToneRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubscribeToToneOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSubscribeToToneOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubscribeToToneOperationable), nil
}
// ToPostRequestInformation subscribe to DTMF (dual-tone multi-frequency signaling). This allows you to be notified when the user presses keys on a 'Dialpad'.
func (m *CallsItemMicrosoftGraphSubscribeToToneRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemMicrosoftGraphSubscribeToToneSubscribeToTonePostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphSubscribeToToneRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
