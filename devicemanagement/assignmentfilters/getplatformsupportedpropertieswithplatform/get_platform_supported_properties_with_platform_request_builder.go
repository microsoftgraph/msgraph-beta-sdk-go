package getplatformsupportedpropertieswithplatform

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetPlatformSupportedPropertiesWithPlatformRequestBuilder provides operations to call the getPlatformSupportedProperties method.
type GetPlatformSupportedPropertiesWithPlatformRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal instantiates a new GetPlatformSupportedPropertiesWithPlatformRequestBuilder and sets the default values.
func NewGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, platform *string)(*GetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    m := &GetPlatformSupportedPropertiesWithPlatformRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/assignmentFilters/microsoft.graph.getPlatformSupportedProperties(platform='{platform}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if platform != nil {
        urlTplParams["platform"] = *platform
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetPlatformSupportedPropertiesWithPlatformRequestBuilder instantiates a new GetPlatformSupportedPropertiesWithPlatformRequestBuilder and sets the default values.
func NewGetPlatformSupportedPropertiesWithPlatformRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getPlatformSupportedProperties
func (m *GetPlatformSupportedPropertiesWithPlatformRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getPlatformSupportedProperties
func (m *GetPlatformSupportedPropertiesWithPlatformRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getPlatformSupportedProperties
func (m *GetPlatformSupportedPropertiesWithPlatformRequestBuilder) Get()(GetPlatformSupportedPropertiesWithPlatformResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getPlatformSupportedProperties
func (m *GetPlatformSupportedPropertiesWithPlatformRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetPlatformSupportedPropertiesWithPlatformResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetPlatformSupportedPropertiesWithPlatformResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetPlatformSupportedPropertiesWithPlatformResponseable), nil
}
