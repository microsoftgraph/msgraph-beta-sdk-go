package setcloudpcreviewstatus

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/microsoft.graph.setCloudPcReviewStatus";
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
// CreatePostRequestInformation invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) CreatePostRequestInformation(body SetCloudPcReviewStatusPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *SetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) Post(body SetCloudPcReviewStatusPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action setCloudPcReviewStatus
func (m *SetCloudPcReviewStatusRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body SetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *SetCloudPcReviewStatusRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
