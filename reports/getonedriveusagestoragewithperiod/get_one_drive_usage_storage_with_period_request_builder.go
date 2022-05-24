package getonedriveusagestoragewithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetOneDriveUsageStorageWithPeriodRequestBuilder provides operations to call the getOneDriveUsageStorage method.
type GetOneDriveUsageStorageWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetOneDriveUsageStorageWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOneDriveUsageStorageWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal instantiates a new GetOneDriveUsageStorageWithPeriodRequestBuilder and sets the default values.
func NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetOneDriveUsageStorageWithPeriodRequestBuilder) {
    m := &GetOneDriveUsageStorageWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOneDriveUsageStorage(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams[""] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOneDriveUsageStorageWithPeriodRequestBuilder instantiates a new GetOneDriveUsageStorageWithPeriodRequestBuilder and sets the default values.
func NewGetOneDriveUsageStorageWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOneDriveUsageStorageWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getOneDriveUsageStorage
func (m *GetOneDriveUsageStorageWithPeriodRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getOneDriveUsageStorage
func (m *GetOneDriveUsageStorageWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetOneDriveUsageStorageWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getOneDriveUsageStorage
func (m *GetOneDriveUsageStorageWithPeriodRequestBuilder) Get()([]byte, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getOneDriveUsageStorage
func (m *GetOneDriveUsageStorageWithPeriodRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetOneDriveUsageStorageWithPeriodRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "byte", responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
