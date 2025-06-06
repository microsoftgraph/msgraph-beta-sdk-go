// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder provides operations to call the bulkReprovisionCloudPc method.
type ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesBulkReprovisionCloudPcRequestBuilderInternal instantiates a new ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesBulkReprovisionCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder) {
    m := &ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/bulkReprovisionCloudPc", pathParameters),
    }
    return m
}
// NewItemManagedDevicesBulkReprovisionCloudPcRequestBuilder instantiates a new ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesBulkReprovisionCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesBulkReprovisionCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post bulk reprovision a set of Cloud PC devices with Intune managed device IDs.
// Deprecated: The bulkReprovisionCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkReprovisionCloudPc on 2023-05-24 and will be removed 2023-09-24
// returns a CloudPcBulkRemoteActionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-bulkreprovisioncloudpc?view=graph-rest-beta
func (m *ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder) Post(ctx context.Context, body ItemManagedDevicesBulkReprovisionCloudPcPostRequestBodyable, requestConfiguration *ItemManagedDevicesBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
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
// ToPostRequestInformation bulk reprovision a set of Cloud PC devices with Intune managed device IDs.
// Deprecated: The bulkReprovisionCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkReprovisionCloudPc on 2023-05-24 and will be removed 2023-09-24
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManagedDevicesBulkReprovisionCloudPcPostRequestBodyable, requestConfiguration *ItemManagedDevicesBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The bulkReprovisionCloudPc action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkReprovisionCloudPc on 2023-05-24 and will be removed 2023-09-24
// returns a *ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder when successful
func (m *ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesBulkReprovisionCloudPcRequestBuilder) {
    return NewItemManagedDevicesBulkReprovisionCloudPcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
