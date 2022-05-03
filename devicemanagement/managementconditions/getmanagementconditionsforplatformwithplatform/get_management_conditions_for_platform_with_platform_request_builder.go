package getmanagementconditionsforplatformwithplatform

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetManagementConditionsForPlatformWithPlatformRequestBuilder provides operations to call the getManagementConditionsForPlatform method.
type GetManagementConditionsForPlatformWithPlatformRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetManagementConditionsForPlatformWithPlatformRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetManagementConditionsForPlatformWithPlatformRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal instantiates a new GetManagementConditionsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, platform *string)(*GetManagementConditionsForPlatformWithPlatformRequestBuilder) {
    m := &GetManagementConditionsForPlatformWithPlatformRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditions/microsoft.graph.getManagementConditionsForPlatform(platform='{platform}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if platform != nil {
        urlTplParams[""] = *platform
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagementConditionsForPlatformWithPlatformRequestBuilder instantiates a new GetManagementConditionsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionsForPlatformWithPlatformRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagementConditionsForPlatformWithPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) GetWithResponseHandler(requestConfiguration *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetRequestConfiguration)(GetManagementConditionsForPlatformWithPlatformResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) GetWithResponseHandler(requestConfiguration *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetManagementConditionsForPlatformWithPlatformResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagementConditionsForPlatformWithPlatformResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagementConditionsForPlatformWithPlatformResponseable), nil
}
