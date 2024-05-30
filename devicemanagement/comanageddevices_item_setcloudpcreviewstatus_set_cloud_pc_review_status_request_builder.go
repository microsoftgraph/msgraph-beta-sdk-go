package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder provides operations to call the setCloudPcReviewStatus method.
type ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderInternal instantiates a new ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) {
    m := &ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/setCloudPcReviewStatus", pathParameters),
    }
    return m
}
// NewComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder instantiates a new ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the review status of a specific Cloud PC device. Use this API to set the review status of a Cloud PC to in review if you consider a Cloud PC as suspicious. After the review is completed, use this API again to set the Cloud PC back to a normal state.
// Deprecated: The setCloudPcReviewStatus API is deprecated and will stop returning data on Apr 30, 2024. Please use the new setReviewStatus API as of 2024-01/setCloudPcReviewStatus
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-setcloudpcreviewstatus?view=graph-rest-beta
func (m *ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) Post(ctx context.Context, body ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set the review status of a specific Cloud PC device. Use this API to set the review status of a Cloud PC to in review if you consider a Cloud PC as suspicious. After the review is completed, use this API again to set the Cloud PC back to a normal state.
// Deprecated: The setCloudPcReviewStatus API is deprecated and will stop returning data on Apr 30, 2024. Please use the new setReviewStatus API as of 2024-01/setCloudPcReviewStatus
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The setCloudPcReviewStatus API is deprecated and will stop returning data on Apr 30, 2024. Please use the new setReviewStatus API as of 2024-01/setCloudPcReviewStatus
// returns a *ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder when successful
func (m *ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) {
    return NewComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
