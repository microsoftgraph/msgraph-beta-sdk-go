package getemailappusageuserdetailwithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetEmailAppUsageUserDetailWithPeriodRequestBuilder provides operations to call the getEmailAppUsageUserDetail method.
type GetEmailAppUsageUserDetailWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal instantiates a new GetEmailAppUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    m := &GetEmailAppUsageUserDetailWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/microsoft.graph.getEmailAppUsageUserDetail(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetEmailAppUsageUserDetailWithPeriodRequestBuilder instantiates a new GetEmailAppUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetEmailAppUsageUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getEmailAppUsageUserDetail
func (m *GetEmailAppUsageUserDetailWithPeriodRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getEmailAppUsageUserDetail
func (m *GetEmailAppUsageUserDetailWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getEmailAppUsageUserDetail
func (m *GetEmailAppUsageUserDetailWithPeriodRequestBuilder) Get()([]byte, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getEmailAppUsageUserDetail
func (m *GetEmailAppUsageUserDetailWithPeriodRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
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
