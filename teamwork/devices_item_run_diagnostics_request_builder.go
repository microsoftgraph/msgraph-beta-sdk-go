package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicesItemRunDiagnosticsRequestBuilder provides operations to call the runDiagnostics method.
type DevicesItemRunDiagnosticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicesItemRunDiagnosticsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicesItemRunDiagnosticsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicesItemRunDiagnosticsRequestBuilderInternal instantiates a new RunDiagnosticsRequestBuilder and sets the default values.
func NewDevicesItemRunDiagnosticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemRunDiagnosticsRequestBuilder) {
    m := &DevicesItemRunDiagnosticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/devices/{teamworkDevice%2Did}/runDiagnostics", pathParameters),
    }
    return m
}
// NewDevicesItemRunDiagnosticsRequestBuilder instantiates a new RunDiagnosticsRequestBuilder and sets the default values.
func NewDevicesItemRunDiagnosticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemRunDiagnosticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemRunDiagnosticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run and generate diagnostic logs for the specified Microsoft Teams-enabled device. This API triggers a long-running operation used to generate logs for a device.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamworkdevice-rundiagnostics?view=graph-rest-1.0
func (m *DevicesItemRunDiagnosticsRequestBuilder) Post(ctx context.Context, requestConfiguration *DevicesItemRunDiagnosticsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation run and generate diagnostic logs for the specified Microsoft Teams-enabled device. This API triggers a long-running operation used to generate logs for a device.
func (m *DevicesItemRunDiagnosticsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DevicesItemRunDiagnosticsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
