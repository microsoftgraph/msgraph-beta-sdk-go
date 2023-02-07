package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicesItemMicrosoftGraphRestartRequestBuilder provides operations to call the restart method.
type DevicesItemMicrosoftGraphRestartRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DevicesItemMicrosoftGraphRestartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicesItemMicrosoftGraphRestartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicesItemMicrosoftGraphRestartRequestBuilderInternal instantiates a new MicrosoftGraphRestartRequestBuilder and sets the default values.
func NewDevicesItemMicrosoftGraphRestartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMicrosoftGraphRestartRequestBuilder) {
    m := &DevicesItemMicrosoftGraphRestartRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/devices/{teamworkDevice%2Did}/microsoft.graph.restart";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDevicesItemMicrosoftGraphRestartRequestBuilder instantiates a new MicrosoftGraphRestartRequestBuilder and sets the default values.
func NewDevicesItemMicrosoftGraphRestartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMicrosoftGraphRestartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemMicrosoftGraphRestartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restart the specified Microsoft Teams-enabled device asynchronously.  A device is restarted after the async operation completes successfully, which might occur subsequent to a response from this API.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/teamworkdevice-restart?view=graph-rest-1.0
func (m *DevicesItemMicrosoftGraphRestartRequestBuilder) Post(ctx context.Context, requestConfiguration *DevicesItemMicrosoftGraphRestartRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation restart the specified Microsoft Teams-enabled device asynchronously.  A device is restarted after the async operation completes successfully, which might occur subsequent to a response from this API.
func (m *DevicesItemMicrosoftGraphRestartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DevicesItemMicrosoftGraphRestartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
