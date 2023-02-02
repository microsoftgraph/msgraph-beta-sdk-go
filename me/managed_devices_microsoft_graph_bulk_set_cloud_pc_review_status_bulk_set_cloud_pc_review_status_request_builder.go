package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder provides operations to call the bulkSetCloudPcReviewStatus method.
type ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilderInternal instantiates a new BulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder) {
    m := &ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/microsoft.graph.bulkSetCloudPcReviewStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder instantiates a new BulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the review status of multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/manageddevice-bulksetcloudpcreviewstatus?view=graph-rest-1.0
func (m *ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder) Post(ctx context.Context, body ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
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
// ToPostRequestInformation set the review status of multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices.
func (m *ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ManagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
