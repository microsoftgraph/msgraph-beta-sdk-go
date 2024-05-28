package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder provides operations to call the resizeCloudPc method.
type ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderInternal instantiates a new ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder and sets the default values.
func NewItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) {
    m := &ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/resizeCloudPc", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder instantiates a new ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder and sets the default values.
func NewItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upgrade or downgrade an existing Cloud PC to another configuration with a new virtual CPU (vCPU) and storage size.
// Deprecated: The resizeCloudPc API is deprecated and will stop returning on Oct 30, 2023. Please use resize instead as of 2023-05/resizeCloudPc
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-resizecloudpc?view=graph-rest-beta
func (m *ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) Post(ctx context.Context, body ItemManageddevicesItemResizecloudpcResizeCloudPcPostRequestBodyable, requestConfiguration *ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation upgrade or downgrade an existing Cloud PC to another configuration with a new virtual CPU (vCPU) and storage size.
// Deprecated: The resizeCloudPc API is deprecated and will stop returning on Oct 30, 2023. Please use resize instead as of 2023-05/resizeCloudPc
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManageddevicesItemResizecloudpcResizeCloudPcPostRequestBodyable, requestConfiguration *ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The resizeCloudPc API is deprecated and will stop returning on Oct 30, 2023. Please use resize instead as of 2023-05/resizeCloudPc
// returns a *ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder when successful
func (m *ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) {
    return NewItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
