package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemReprovisionCloudPcRequestBuilder provides operations to call the reprovisionCloudPc method.
type ItemManagedDevicesItemReprovisionCloudPcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemReprovisionCloudPcRequestBuilderInternal instantiates a new ItemManagedDevicesItemReprovisionCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesItemReprovisionCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    m := &ItemManagedDevicesItemReprovisionCloudPcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/reprovisionCloudPc", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemReprovisionCloudPcRequestBuilder instantiates a new ItemManagedDevicesItemReprovisionCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesItemReprovisionCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemReprovisionCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action reprovisionCloudPc
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemReprovisionCloudPcRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action reprovisionCloudPc
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemReprovisionCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc
// returns a *ItemManagedDevicesItemReprovisionCloudPcRequestBuilder when successful
func (m *ItemManagedDevicesItemReprovisionCloudPcRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewItemManagedDevicesItemReprovisionCloudPcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
