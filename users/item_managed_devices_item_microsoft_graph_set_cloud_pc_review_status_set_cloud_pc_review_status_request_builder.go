package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder provides operations to call the setCloudPcReviewStatus method.
type ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderInternal instantiates a new SetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder) {
    m := &ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/microsoft.graph.setCloudPcReviewStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder instantiates a new SetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the review status of a specific Cloud PC device. Use this API to set the review status of a Cloud PC to in review if you consider a Cloud PC as suspicious. After the review is completed, use this API again to set the Cloud PC back to a normal state.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/manageddevice-setcloudpcreviewstatus?view=graph-rest-1.0
func (m *ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder) Post(ctx context.Context, body ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation set the review status of a specific Cloud PC device. Use this API to set the review status of a Cloud PC to in review if you consider a Cloud PC as suspicious. After the review is completed, use this API again to set the Cloud PC back to a normal state.
func (m *ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
