package getazureaduserfeatureusage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetAzureADUserFeatureUsageRequestBuilder provides operations to call the getAzureADUserFeatureUsage method.
type GetAzureADUserFeatureUsageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetAzureADUserFeatureUsageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAzureADUserFeatureUsageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetAzureADUserFeatureUsageRequestBuilderInternal instantiates a new GetAzureADUserFeatureUsageRequestBuilder and sets the default values.
func NewGetAzureADUserFeatureUsageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAzureADUserFeatureUsageRequestBuilder) {
    m := &GetAzureADUserFeatureUsageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getAzureADUserFeatureUsage()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAzureADUserFeatureUsageRequestBuilder instantiates a new GetAzureADUserFeatureUsageRequestBuilder and sets the default values.
func NewGetAzureADUserFeatureUsageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAzureADUserFeatureUsageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAzureADUserFeatureUsageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getAzureADUserFeatureUsage
func (m *GetAzureADUserFeatureUsageRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getAzureADUserFeatureUsage
func (m *GetAzureADUserFeatureUsageRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetAzureADUserFeatureUsageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getAzureADUserFeatureUsage
func (m *GetAzureADUserFeatureUsageRequestBuilder) Get()(GetAzureADUserFeatureUsageResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getAzureADUserFeatureUsage
func (m *GetAzureADUserFeatureUsageRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetAzureADUserFeatureUsageRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetAzureADUserFeatureUsageResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetAzureADUserFeatureUsageResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetAzureADUserFeatureUsageResponseable), nil
}
