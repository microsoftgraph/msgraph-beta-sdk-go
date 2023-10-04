package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesBulkRestoreCloudPcRequestBuilder provides operations to call the bulkRestoreCloudPc method.
type ManagedDevicesBulkRestoreCloudPcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesBulkRestoreCloudPcRequestBuilderInternal instantiates a new BulkRestoreCloudPcRequestBuilder and sets the default values.
func NewManagedDevicesBulkRestoreCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesBulkRestoreCloudPcRequestBuilder) {
    m := &ManagedDevicesBulkRestoreCloudPcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/bulkRestoreCloudPc", pathParameters),
    }
    return m
}
// NewManagedDevicesBulkRestoreCloudPcRequestBuilder instantiates a new BulkRestoreCloudPcRequestBuilder and sets the default values.
func NewManagedDevicesBulkRestoreCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesBulkRestoreCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesBulkRestoreCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices and a restore point date and time. This API is supported in the following national cloud deployments.
// Deprecated: The bulkRestoreCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkRestoreCloudPc on 2023-05-24 and will be removed 2023-09-24
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-bulkrestorecloudpc?view=graph-rest-1.0
func (m *ManagedDevicesBulkRestoreCloudPcRequestBuilder) Post(ctx context.Context, body ManagedDevicesBulkRestoreCloudPcPostRequestBodyable, requestConfiguration *ManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation restore multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices and a restore point date and time. This API is supported in the following national cloud deployments.
// Deprecated: The bulkRestoreCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkRestoreCloudPc on 2023-05-24 and will be removed 2023-09-24
func (m *ManagedDevicesBulkRestoreCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesBulkRestoreCloudPcPostRequestBodyable, requestConfiguration *ManagedDevicesBulkRestoreCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The bulkRestoreCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkRestoreCloudPc on 2023-05-24 and will be removed 2023-09-24
func (m *ManagedDevicesBulkRestoreCloudPcRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesBulkRestoreCloudPcRequestBuilder) {
    return NewManagedDevicesBulkRestoreCloudPcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
