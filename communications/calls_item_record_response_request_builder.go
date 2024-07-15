package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemRecordResponseRequestBuilder provides operations to call the recordResponse method.
type CallsItemRecordResponseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemRecordResponseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemRecordResponseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemRecordResponseRequestBuilderInternal instantiates a new CallsItemRecordResponseRequestBuilder and sets the default values.
func NewCallsItemRecordResponseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemRecordResponseRequestBuilder) {
    m := &CallsItemRecordResponseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/recordResponse", pathParameters),
    }
    return m
}
// NewCallsItemRecordResponseRequestBuilder instantiates a new CallsItemRecordResponseRequestBuilder and sets the default values.
func NewCallsItemRecordResponseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemRecordResponseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemRecordResponseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post record a short audio response from the caller. A bot can use this API to capture a voice response from a caller after they're prompted for a response. For more information about how to handle operations, see commsOperation. This action isn't intended to record the entire call. The maximum length of recording is 2 minutes.The Cloud Communications Platform doesn't save the recording permanently and discards it shortly after the call ends. The bot must download the recording promptly after the recording operation finishes by using the recordingLocation value provided in the completed notification.
// returns a RecordOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-record?view=graph-rest-beta
func (m *CallsItemRecordResponseRequestBuilder) Post(ctx context.Context, body CallsItemRecordResponsePostRequestBodyable, requestConfiguration *CallsItemRecordResponseRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecordOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRecordOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecordOperationable), nil
}
// ToPostRequestInformation record a short audio response from the caller. A bot can use this API to capture a voice response from a caller after they're prompted for a response. For more information about how to handle operations, see commsOperation. This action isn't intended to record the entire call. The maximum length of recording is 2 minutes.The Cloud Communications Platform doesn't save the recording permanently and discards it shortly after the call ends. The bot must download the recording promptly after the recording operation finishes by using the recordingLocation value provided in the completed notification.
// returns a *RequestInformation when successful
func (m *CallsItemRecordResponseRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemRecordResponsePostRequestBodyable, requestConfiguration *CallsItemRecordResponseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemRecordResponseRequestBuilder when successful
func (m *CallsItemRecordResponseRequestBuilder) WithUrl(rawUrl string)(*CallsItemRecordResponseRequestBuilder) {
    return NewCallsItemRecordResponseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
