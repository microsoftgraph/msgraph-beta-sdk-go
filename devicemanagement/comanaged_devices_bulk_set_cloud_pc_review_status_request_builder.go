package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder provides operations to call the bulkSetCloudPcReviewStatus method.
type ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal instantiates a new BulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    m := &ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/bulkSetCloudPcReviewStatus", pathParameters),
    }
    return m
}
// NewComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder instantiates a new BulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the review status of multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-bulksetcloudpcreviewstatus?view=graph-rest-1.0
func (m *ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) Post(ctx context.Context, body ComanagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
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
// ToPostRequestInformation set the review status of multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices.
func (m *ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
