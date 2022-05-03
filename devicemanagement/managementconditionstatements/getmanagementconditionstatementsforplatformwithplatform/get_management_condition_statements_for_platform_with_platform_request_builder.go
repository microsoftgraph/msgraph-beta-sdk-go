package getmanagementconditionstatementsforplatformwithplatform

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder provides operations to call the getManagementConditionStatementsForPlatform method.
type GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetManagementConditionStatementsForPlatformWithPlatformRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetManagementConditionStatementsForPlatformWithPlatformRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetManagementConditionStatementsForPlatformWithPlatformRequestBuilderInternal instantiates a new GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionStatementsForPlatformWithPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, platform *string)(*GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) {
    m := &GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditionStatements/microsoft.graph.getManagementConditionStatementsForPlatform(platform='{platform}')";
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
// NewGetManagementConditionStatementsForPlatformWithPlatformRequestBuilder instantiates a new GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionStatementsForPlatformWithPlatformRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagementConditionStatementsForPlatformWithPlatformRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getManagementConditionStatementsForPlatform
func (m *GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getManagementConditionStatementsForPlatform
func (m *GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetManagementConditionStatementsForPlatformWithPlatformRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler invoke function getManagementConditionStatementsForPlatform
func (m *GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) GetWithResponseHandler(requestConfiguration *GetManagementConditionStatementsForPlatformWithPlatformRequestBuilderGetRequestConfiguration)(GetManagementConditionStatementsForPlatformWithPlatformResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getManagementConditionStatementsForPlatform
func (m *GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) GetWithResponseHandler(requestConfiguration *GetManagementConditionStatementsForPlatformWithPlatformRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetManagementConditionStatementsForPlatformWithPlatformResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagementConditionStatementsForPlatformWithPlatformResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagementConditionStatementsForPlatformWithPlatformResponseable), nil
}
