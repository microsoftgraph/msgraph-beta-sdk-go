package enableandroiddeviceadministratorenrollment

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnableAndroidDeviceAdministratorEnrollmentRequestBuilder provides operations to call the enableAndroidDeviceAdministratorEnrollment method.
type EnableAndroidDeviceAdministratorEnrollmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EnableAndroidDeviceAdministratorEnrollmentRequestBuilderPostOptions options for Post
type EnableAndroidDeviceAdministratorEnrollmentRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal instantiates a new EnableAndroidDeviceAdministratorEnrollmentRequestBuilder and sets the default values.
func NewEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnableAndroidDeviceAdministratorEnrollmentRequestBuilder) {
    m := &EnableAndroidDeviceAdministratorEnrollmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.enableAndroidDeviceAdministratorEnrollment";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEnableAndroidDeviceAdministratorEnrollmentRequestBuilder instantiates a new EnableAndroidDeviceAdministratorEnrollmentRequestBuilder and sets the default values.
func NewEnableAndroidDeviceAdministratorEnrollmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnableAndroidDeviceAdministratorEnrollmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action enableAndroidDeviceAdministratorEnrollment
func (m *EnableAndroidDeviceAdministratorEnrollmentRequestBuilder) CreatePostRequestInformation(options *EnableAndroidDeviceAdministratorEnrollmentRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action enableAndroidDeviceAdministratorEnrollment
func (m *EnableAndroidDeviceAdministratorEnrollmentRequestBuilder) Post(options *EnableAndroidDeviceAdministratorEnrollmentRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
