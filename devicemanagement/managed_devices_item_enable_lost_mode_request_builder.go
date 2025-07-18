// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemEnableLostModeRequestBuilder provides operations to call the enableLostMode method.
type ManagedDevicesItemEnableLostModeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemEnableLostModeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemEnableLostModeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemEnableLostModeRequestBuilderInternal instantiates a new ManagedDevicesItemEnableLostModeRequestBuilder and sets the default values.
func NewManagedDevicesItemEnableLostModeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemEnableLostModeRequestBuilder) {
    m := &ManagedDevicesItemEnableLostModeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/enableLostMode", pathParameters),
    }
    return m
}
// NewManagedDevicesItemEnableLostModeRequestBuilder instantiates a new ManagedDevicesItemEnableLostModeRequestBuilder and sets the default values.
func NewManagedDevicesItemEnableLostModeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemEnableLostModeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemEnableLostModeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enable lost mode
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemEnableLostModeRequestBuilder) Post(ctx context.Context, body ManagedDevicesItemEnableLostModePostRequestBodyable, requestConfiguration *ManagedDevicesItemEnableLostModeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation enable lost mode
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemEnableLostModeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesItemEnableLostModePostRequestBodyable, requestConfiguration *ManagedDevicesItemEnableLostModeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedDevicesItemEnableLostModeRequestBuilder when successful
func (m *ManagedDevicesItemEnableLostModeRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemEnableLostModeRequestBuilder) {
    return NewManagedDevicesItemEnableLostModeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
