package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemPlayLostModeSoundRequestBuilder provides operations to call the playLostModeSound method.
type ManagedDevicesItemPlayLostModeSoundRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemPlayLostModeSoundRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemPlayLostModeSoundRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemPlayLostModeSoundRequestBuilderInternal instantiates a new PlayLostModeSoundRequestBuilder and sets the default values.
func NewManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    m := &ManagedDevicesItemPlayLostModeSoundRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/playLostModeSound", pathParameters),
    }
    return m
}
// NewManagedDevicesItemPlayLostModeSoundRequestBuilder instantiates a new PlayLostModeSoundRequestBuilder and sets the default values.
func NewManagedDevicesItemPlayLostModeSoundRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(urlParams, requestAdapter)
}
// Post play lost mode sound
func (m *ManagedDevicesItemPlayLostModeSoundRequestBuilder) Post(ctx context.Context, body ManagedDevicesItemPlayLostModeSoundPostRequestBodyable, requestConfiguration *ManagedDevicesItemPlayLostModeSoundRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation play lost mode sound
func (m *ManagedDevicesItemPlayLostModeSoundRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesItemPlayLostModeSoundPostRequestBodyable, requestConfiguration *ManagedDevicesItemPlayLostModeSoundRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedDevicesItemPlayLostModeSoundRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewManagedDevicesItemPlayLostModeSoundRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
