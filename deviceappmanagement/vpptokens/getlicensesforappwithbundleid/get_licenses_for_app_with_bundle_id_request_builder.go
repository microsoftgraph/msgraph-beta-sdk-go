package getlicensesforappwithbundleid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetLicensesForAppWithBundleIdRequestBuilder provides operations to call the getLicensesForApp method.
type GetLicensesForAppWithBundleIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetLicensesForAppWithBundleIdRequestBuilderInternal instantiates a new GetLicensesForAppWithBundleIdRequestBuilder and sets the default values.
func NewGetLicensesForAppWithBundleIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, bundleId *string)(*GetLicensesForAppWithBundleIdRequestBuilder) {
    m := &GetLicensesForAppWithBundleIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/vppTokens/microsoft.graph.getLicensesForApp(bundleId='{bundleId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if bundleId != nil {
        urlTplParams[""] = *bundleId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetLicensesForAppWithBundleIdRequestBuilder instantiates a new GetLicensesForAppWithBundleIdRequestBuilder and sets the default values.
func NewGetLicensesForAppWithBundleIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetLicensesForAppWithBundleIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetLicensesForAppWithBundleIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getLicensesForApp
func (m *GetLicensesForAppWithBundleIdRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getLicensesForApp
func (m *GetLicensesForAppWithBundleIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getLicensesForApp
func (m *GetLicensesForAppWithBundleIdRequestBuilder) Get()(GetLicensesForAppWithBundleIdResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getLicensesForApp
func (m *GetLicensesForAppWithBundleIdRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetLicensesForAppWithBundleIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetLicensesForAppWithBundleIdResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetLicensesForAppWithBundleIdResponseable), nil
}
