package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder provides operations to call the runDiagnostics method.
type DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicesItemMicrosoftGraphRunDiagnosticsRequestBuilderInternal instantiates a new MicrosoftGraphRunDiagnosticsRequestBuilder and sets the default values.
func NewDevicesItemMicrosoftGraphRunDiagnosticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder) {
    m := &DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/devices/{teamworkDevice%2Did}/microsoft.graph.runDiagnostics";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder instantiates a new MicrosoftGraphRunDiagnosticsRequestBuilder and sets the default values.
func NewDevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemMicrosoftGraphRunDiagnosticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run and generate diagnostic logs for the specified Microsoft Teams-enabled device. This API triggers a long-running operation used to generate logs for a device.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/teamworkdevice-rundiagnostics?view=graph-rest-1.0
func (m *DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder) Post(ctx context.Context, requestConfiguration *DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation run and generate diagnostic logs for the specified Microsoft Teams-enabled device. This API triggers a long-running operation used to generate logs for a device.
func (m *DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DevicesItemMicrosoftGraphRunDiagnosticsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
