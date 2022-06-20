package getoffice365activationsusercounts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetOffice365ActivationsUserCountsRequestBuilder provides operations to call the getOffice365ActivationsUserCounts method.
type GetOffice365ActivationsUserCountsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetOffice365ActivationsUserCountsRequestBuilderInternal instantiates a new GetOffice365ActivationsUserCountsRequestBuilder and sets the default values.
func NewGetOffice365ActivationsUserCountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365ActivationsUserCountsRequestBuilder) {
    m := &GetOffice365ActivationsUserCountsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365ActivationsUserCounts()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOffice365ActivationsUserCountsRequestBuilder instantiates a new GetOffice365ActivationsUserCountsRequestBuilder and sets the default values.
func NewGetOffice365ActivationsUserCountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365ActivationsUserCountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365ActivationsUserCountsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getOffice365ActivationsUserCounts
func (m *GetOffice365ActivationsUserCountsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getOffice365ActivationsUserCounts
func (m *GetOffice365ActivationsUserCountsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getOffice365ActivationsUserCounts
func (m *GetOffice365ActivationsUserCountsRequestBuilder) Get()([]byte, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getOffice365ActivationsUserCounts
func (m *GetOffice365ActivationsUserCountsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetOffice365ActivationsUserCountsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "byte", responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
