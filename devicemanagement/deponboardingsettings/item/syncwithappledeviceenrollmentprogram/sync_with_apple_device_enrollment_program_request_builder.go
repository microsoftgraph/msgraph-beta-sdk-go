package syncwithappledeviceenrollmentprogram

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SyncWithAppleDeviceEnrollmentProgramRequestBuilder provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
type SyncWithAppleDeviceEnrollmentProgramRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal instantiates a new SyncWithAppleDeviceEnrollmentProgramRequestBuilder and sets the default values.
func NewSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    m := &SyncWithAppleDeviceEnrollmentProgramRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/microsoft.graph.syncWithAppleDeviceEnrollmentProgram";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSyncWithAppleDeviceEnrollmentProgramRequestBuilder instantiates a new SyncWithAppleDeviceEnrollmentProgramRequestBuilder and sets the default values.
func NewSyncWithAppleDeviceEnrollmentProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration synchronizes between Apple Device Enrollment Program and Intune
func (m *SyncWithAppleDeviceEnrollmentProgramRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration synchronizes between Apple Device Enrollment Program and Intune
func (m *SyncWithAppleDeviceEnrollmentProgramRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *SyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler synchronizes between Apple Device Enrollment Program and Intune
func (m *SyncWithAppleDeviceEnrollmentProgramRequestBuilder) PostWithResponseHandler(requestConfiguration *SyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler synchronizes between Apple Device Enrollment Program and Intune
func (m *SyncWithAppleDeviceEnrollmentProgramRequestBuilder) PostWithResponseHandler(requestConfiguration *SyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
