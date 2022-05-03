package createenrollmentnotificationconfiguration

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CreateEnrollmentNotificationConfigurationRequestBuilder provides operations to call the createEnrollmentNotificationConfiguration method.
type CreateEnrollmentNotificationConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCreateEnrollmentNotificationConfigurationRequestBuilderInternal instantiates a new CreateEnrollmentNotificationConfigurationRequestBuilder and sets the default values.
func NewCreateEnrollmentNotificationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateEnrollmentNotificationConfigurationRequestBuilder) {
    m := &CreateEnrollmentNotificationConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}/microsoft.graph.createEnrollmentNotificationConfiguration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateEnrollmentNotificationConfigurationRequestBuilder instantiates a new CreateEnrollmentNotificationConfigurationRequestBuilder and sets the default values.
func NewCreateEnrollmentNotificationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateEnrollmentNotificationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateEnrollmentNotificationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createEnrollmentNotificationConfiguration
func (m *CreateEnrollmentNotificationConfigurationRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action createEnrollmentNotificationConfiguration
func (m *CreateEnrollmentNotificationConfigurationRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *CreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action createEnrollmentNotificationConfiguration
func (m *CreateEnrollmentNotificationConfigurationRequestBuilder) Post()(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action createEnrollmentNotificationConfiguration
func (m *CreateEnrollmentNotificationConfigurationRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *CreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
