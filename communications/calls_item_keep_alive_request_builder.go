package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemKeepAliveRequestBuilder provides operations to call the keepAlive method.
type CallsItemKeepAliveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemKeepAliveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemKeepAliveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemKeepAliveRequestBuilderInternal instantiates a new CallsItemKeepAliveRequestBuilder and sets the default values.
func NewCallsItemKeepAliveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemKeepAliveRequestBuilder) {
    m := &CallsItemKeepAliveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/keepAlive", pathParameters),
    }
    return m
}
// NewCallsItemKeepAliveRequestBuilder instantiates a new CallsItemKeepAliveRequestBuilder and sets the default values.
func NewCallsItemKeepAliveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemKeepAliveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemKeepAliveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post make a request to this API every 15 to 45 minutes to ensure that an ongoing call remains active. A call that doesn't receive this request within 45 minutes is considered inactive and ends. At least one successful request must be made within 45 minutes of the previous request, or the start of the call. We recommend that you send a request in shorter time intervals (every 15 minutes). Make sure that these requests are successful to prevent the call from timing out and ending. Attempting to send a request to a call that ended results in a 404 Not Found error. The resources related to the call should be cleaned up on the application side.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-keepalive?view=graph-rest-beta
func (m *CallsItemKeepAliveRequestBuilder) Post(ctx context.Context, requestConfiguration *CallsItemKeepAliveRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation make a request to this API every 15 to 45 minutes to ensure that an ongoing call remains active. A call that doesn't receive this request within 45 minutes is considered inactive and ends. At least one successful request must be made within 45 minutes of the previous request, or the start of the call. We recommend that you send a request in shorter time intervals (every 15 minutes). Make sure that these requests are successful to prevent the call from timing out and ending. Attempting to send a request to a call that ended results in a 404 Not Found error. The resources related to the call should be cleaned up on the application side.
// returns a *RequestInformation when successful
func (m *CallsItemKeepAliveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CallsItemKeepAliveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallsItemKeepAliveRequestBuilder when successful
func (m *CallsItemKeepAliveRequestBuilder) WithUrl(rawUrl string)(*CallsItemKeepAliveRequestBuilder) {
    return NewCallsItemKeepAliveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
