package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesBulkRestoreCloudPcRequestBuilder provides operations to call the bulkRestoreCloudPc method.
type ItemManagedDevicesBulkRestoreCloudPcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesBulkRestoreCloudPcRequestBuilderInternal instantiates a new ItemManagedDevicesBulkRestoreCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesBulkRestoreCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesBulkRestoreCloudPcRequestBuilder) {
    m := &ItemManagedDevicesBulkRestoreCloudPcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/bulkRestoreCloudPc", pathParameters),
    }
    return m
}
// NewItemManagedDevicesBulkRestoreCloudPcRequestBuilder instantiates a new ItemManagedDevicesBulkRestoreCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesBulkRestoreCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesBulkRestoreCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesBulkRestoreCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices and a restore point date and time.
// Deprecated: The bulkRestoreCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkRestoreCloudPc
// returns a CloudPcBulkRemoteActionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-bulkrestorecloudpc?view=graph-rest-1.0
func (m *ItemManagedDevicesBulkRestoreCloudPcRequestBuilder) Post(ctx context.Context, body ItemManagedDevicesBulkRestoreCloudPcPostRequestBodyable, requestConfiguration *ItemManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcBulkRemoteActionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable), nil
}
// ToPostRequestInformation restore multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices and a restore point date and time.
// Deprecated: The bulkRestoreCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkRestoreCloudPc
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesBulkRestoreCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManagedDevicesBulkRestoreCloudPcPostRequestBodyable, requestConfiguration *ItemManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The bulkRestoreCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkRestoreCloudPc
// returns a *ItemManagedDevicesBulkRestoreCloudPcRequestBuilder when successful
func (m *ItemManagedDevicesBulkRestoreCloudPcRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesBulkRestoreCloudPcRequestBuilder) {
    return NewItemManagedDevicesBulkRestoreCloudPcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
