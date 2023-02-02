package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder provides operations to call the bulkReprovisionCloudPc method.
type ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilderInternal instantiates a new BulkReprovisionCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder) {
    m := &ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/microsoft.graph.bulkReprovisionCloudPc";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder instantiates a new BulkReprovisionCloudPcRequestBuilder and sets the default values.
func NewItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post bulk reprovision a set of Cloud PC devices with Intune managed device IDs.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/manageddevice-bulkreprovisioncloudpc?view=graph-rest-1.0
func (m *ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder) Post(ctx context.Context, body ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBodyable, requestConfiguration *ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcBulkRemoteActionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable), nil
}
// ToPostRequestInformation bulk reprovision a set of Cloud PC devices with Intune managed device IDs.
func (m *ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBodyable, requestConfiguration *ItemManagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
