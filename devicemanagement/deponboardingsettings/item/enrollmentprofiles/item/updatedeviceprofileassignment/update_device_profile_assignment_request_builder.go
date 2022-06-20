package updatedeviceprofileassignment

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UpdateDeviceProfileAssignmentRequestBuilder provides operations to call the updateDeviceProfileAssignment method.
type UpdateDeviceProfileAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UpdateDeviceProfileAssignmentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UpdateDeviceProfileAssignmentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUpdateDeviceProfileAssignmentRequestBuilderInternal instantiates a new UpdateDeviceProfileAssignmentRequestBuilder and sets the default values.
func NewUpdateDeviceProfileAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdateDeviceProfileAssignmentRequestBuilder) {
    m := &UpdateDeviceProfileAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/microsoft.graph.updateDeviceProfileAssignment";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateDeviceProfileAssignmentRequestBuilder instantiates a new UpdateDeviceProfileAssignmentRequestBuilder and sets the default values.
func NewUpdateDeviceProfileAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdateDeviceProfileAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateDeviceProfileAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) CreatePostRequestInformation(body UpdateDeviceProfileAssignmentPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body UpdateDeviceProfileAssignmentPostRequestBodyable, requestConfiguration *UpdateDeviceProfileAssignmentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) Post(body UpdateDeviceProfileAssignmentPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body UpdateDeviceProfileAssignmentPostRequestBodyable, requestConfiguration *UpdateDeviceProfileAssignmentRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
