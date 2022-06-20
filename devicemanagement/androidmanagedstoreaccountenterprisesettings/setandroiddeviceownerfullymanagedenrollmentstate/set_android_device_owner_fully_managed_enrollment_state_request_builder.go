package setandroiddeviceownerfullymanagedenrollmentstate

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
// CreatePostRequestInformation sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) CreatePostRequestInformation(body SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable, requestConfiguration *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) Post(body SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable, requestConfiguration *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
