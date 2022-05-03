package unshareforschooldatasyncservice

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UnshareForSchoolDataSyncServiceRequestBuilder provides operations to call the unshareForSchoolDataSyncService method.
type UnshareForSchoolDataSyncServiceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUnshareForSchoolDataSyncServiceRequestBuilderInternal instantiates a new UnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewUnshareForSchoolDataSyncServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnshareForSchoolDataSyncServiceRequestBuilder) {
    m := &UnshareForSchoolDataSyncServiceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/microsoft.graph.unshareForSchoolDataSyncService";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnshareForSchoolDataSyncServiceRequestBuilder instantiates a new UnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewUnshareForSchoolDataSyncServiceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnshareForSchoolDataSyncServiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnshareForSchoolDataSyncServiceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action unshareForSchoolDataSyncService
func (m *UnshareForSchoolDataSyncServiceRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action unshareForSchoolDataSyncService
func (m *UnshareForSchoolDataSyncServiceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *UnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action unshareForSchoolDataSyncService
func (m *UnshareForSchoolDataSyncServiceRequestBuilder) Post()(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action unshareForSchoolDataSyncService
func (m *UnshareForSchoolDataSyncServiceRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *UnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
