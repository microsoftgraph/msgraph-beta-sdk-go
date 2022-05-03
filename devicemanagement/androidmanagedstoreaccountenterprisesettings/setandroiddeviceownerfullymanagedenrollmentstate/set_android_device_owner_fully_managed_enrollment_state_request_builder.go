package setandroiddeviceownerfullymanagedenrollmentstate

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
type SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal instantiates a new SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder and sets the default values.
func NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    m := &SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/microsoft.graph.setAndroidDeviceOwnerFullyManagedEnrollmentState";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder instantiates a new SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder and sets the default values.
func NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBodyable, requestConfiguration *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) PostWithResponseHandler(body SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBodyable, requestConfiguration *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) PostWithResponseHandler(body SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBodyable, requestConfiguration *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
