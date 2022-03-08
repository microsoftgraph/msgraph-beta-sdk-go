package getlicensesforappwithbundleid

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetLicensesForAppWithBundleIdRequestBuilder provides operations to call the getLicensesForApp method.
type GetLicensesForAppWithBundleIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetLicensesForAppWithBundleIdRequestBuilderGetOptions options for Get
type GetLicensesForAppWithBundleIdRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetLicensesForAppWithBundleIdRequestBuilderInternal instantiates a new GetLicensesForAppWithBundleIdRequestBuilder and sets the default values.
func NewGetLicensesForAppWithBundleIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, bundleId *string)(*GetLicensesForAppWithBundleIdRequestBuilder) {
    m := &GetLicensesForAppWithBundleIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/vppTokens/microsoft.graph.getLicensesForApp(bundleId='{bundleId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if bundleId != nil {
        urlTplParams["bundleId"] = *bundleId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetLicensesForAppWithBundleIdRequestBuilder instantiates a new GetLicensesForAppWithBundleIdRequestBuilder and sets the default values.
func NewGetLicensesForAppWithBundleIdRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetLicensesForAppWithBundleIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetLicensesForAppWithBundleIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getLicensesForApp
func (m *GetLicensesForAppWithBundleIdRequestBuilder) CreateGetRequestInformation(options *GetLicensesForAppWithBundleIdRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
