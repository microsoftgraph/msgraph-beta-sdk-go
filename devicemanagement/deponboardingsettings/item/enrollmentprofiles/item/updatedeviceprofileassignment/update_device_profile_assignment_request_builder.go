package updatedeviceprofileassignment

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
// CreatePostRequestInformationWithRequestConfiguration invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body UpdateDeviceProfileAssignmentRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body UpdateDeviceProfileAssignmentRequestBodyable, requestConfiguration *UpdateDeviceProfileAssignmentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) PostWithResponseHandler(body UpdateDeviceProfileAssignmentRequestBodyable, requestConfiguration *UpdateDeviceProfileAssignmentRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) PostWithResponseHandler(body UpdateDeviceProfileAssignmentRequestBodyable, requestConfiguration *UpdateDeviceProfileAssignmentRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
