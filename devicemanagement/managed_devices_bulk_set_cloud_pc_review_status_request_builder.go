package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder provides operations to call the bulkSetCloudPcReviewStatus method.
type ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal instantiates a new ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    m := &ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/bulkSetCloudPcReviewStatus", pathParameters),
    }
    return m
}
// NewManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder instantiates a new ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action bulkSetCloudPcReviewStatus
// Deprecated: The BulkSetCloudPcReviewStatus action is deprecated and will stop supporting on September 30, 2024. Please use bulk action entity api. as of 2024-05/BulkSetCloudPcReviewStatus
// returns a CloudPcBulkRemoteActionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-bulksetcloudpcreviewstatus?view=graph-rest-beta
func (m *ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) Post(ctx context.Context, body ManagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
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
// ToPostRequestInformation invoke action bulkSetCloudPcReviewStatus
// Deprecated: The BulkSetCloudPcReviewStatus action is deprecated and will stop supporting on September 30, 2024. Please use bulk action entity api. as of 2024-05/BulkSetCloudPcReviewStatus
// returns a *RequestInformation when successful
func (m *ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The BulkSetCloudPcReviewStatus action is deprecated and will stop supporting on September 30, 2024. Please use bulk action entity api. as of 2024-05/BulkSetCloudPcReviewStatus
// returns a *ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder when successful
func (m *ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    return NewManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
