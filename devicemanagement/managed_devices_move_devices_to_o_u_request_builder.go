package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesMoveDevicesToOURequestBuilder provides operations to call the moveDevicesToOU method.
type ManagedDevicesMoveDevicesToOURequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesMoveDevicesToOURequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesMoveDevicesToOURequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesMoveDevicesToOURequestBuilderInternal instantiates a new ManagedDevicesMoveDevicesToOURequestBuilder and sets the default values.
func NewManagedDevicesMoveDevicesToOURequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesMoveDevicesToOURequestBuilder) {
    m := &ManagedDevicesMoveDevicesToOURequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/moveDevicesToOU", pathParameters),
    }
    return m
}
// NewManagedDevicesMoveDevicesToOURequestBuilder instantiates a new ManagedDevicesMoveDevicesToOURequestBuilder and sets the default values.
func NewManagedDevicesMoveDevicesToOURequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesMoveDevicesToOURequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesMoveDevicesToOURequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action moveDevicesToOU
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesMoveDevicesToOURequestBuilder) Post(ctx context.Context, body ManagedDevicesMoveDevicesToOUPostRequestBodyable, requestConfiguration *ManagedDevicesMoveDevicesToOURequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action moveDevicesToOU
// returns a *RequestInformation when successful
func (m *ManagedDevicesMoveDevicesToOURequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesMoveDevicesToOUPostRequestBodyable, requestConfiguration *ManagedDevicesMoveDevicesToOURequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedDevicesMoveDevicesToOURequestBuilder when successful
func (m *ManagedDevicesMoveDevicesToOURequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesMoveDevicesToOURequestBuilder) {
    return NewManagedDevicesMoveDevicesToOURequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
