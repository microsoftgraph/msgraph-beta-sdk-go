package setcloudpcreviewstatus

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetCloudPcReviewStatusRequestBuilder provides operations to call the setCloudPcReviewStatus method.
type SetCloudPcReviewStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetCloudPcReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetCloudPcReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSetCloudPcReviewStatusRequestBuilderInternal instantiates a new SetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewSetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetCloudPcReviewStatusRequestBuilder) {
    m := &SetCloudPcReviewStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/microsoft.graph.setCloudPcReviewStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetCloudPcReviewStatusRequestBuilder instantiates a new SetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewSetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetCloudPcReviewStatusRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetCloudPcReviewStatusRequestBodyable, requestConfiguration *SetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) PostWithResponseHandler(body SetCloudPcReviewStatusRequestBodyable, requestConfiguration *SetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) PostWithResponseHandler(body SetCloudPcReviewStatusRequestBodyable, requestConfiguration *SetCloudPcReviewStatusRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
