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
// GetLicensesForAppWithBundleIdRequestBuilderGetOptions options for Get
type GetLicensesForAppWithBundleIdRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
func (m *GetLicensesForAppWithBundleIdRequestBuilder) CreateGetRequestInformation(options *GetLicensesForAppWithBundleIdRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function getLicensesForApp
func (m *GetLicensesForAppWithBundleIdRequestBuilder) Get(options *GetLicensesForAppWithBundleIdRequestBuilderGetOptions)(GetLicensesForAppWithBundleIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetLicensesForAppWithBundleIdResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetLicensesForAppWithBundleIdResponseable), nil
}
