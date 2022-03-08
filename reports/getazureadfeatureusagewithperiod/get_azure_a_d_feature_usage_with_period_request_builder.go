package getazureadfeatureusagewithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetAzureADFeatureUsageWithPeriodRequestBuilder provides operations to call the getAzureADFeatureUsage method.
type GetAzureADFeatureUsageWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetAzureADFeatureUsageWithPeriodRequestBuilderGetOptions options for Get
type GetAzureADFeatureUsageWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetAzureADFeatureUsageWithPeriodRequestBuilderInternal instantiates a new GetAzureADFeatureUsageWithPeriodRequestBuilder and sets the default values.
func NewGetAzureADFeatureUsageWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetAzureADFeatureUsageWithPeriodRequestBuilder) {
    m := &GetAzureADFeatureUsageWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getAzureADFeatureUsage(period='{period}')";
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
// NewGetAzureADFeatureUsageWithPeriodRequestBuilder instantiates a new GetAzureADFeatureUsageWithPeriodRequestBuilder and sets the default values.
func NewGetAzureADFeatureUsageWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAzureADFeatureUsageWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAzureADFeatureUsageWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getAzureADFeatureUsage
func (m *GetAzureADFeatureUsageWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetAzureADFeatureUsageWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getAzureADFeatureUsage
func (m *GetAzureADFeatureUsageWithPeriodRequestBuilder) Get(options *GetAzureADFeatureUsageWithPeriodRequestBuilderGetOptions)(GetAzureADFeatureUsageWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetAzureADFeatureUsageWithPeriodResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetAzureADFeatureUsageWithPeriodResponseable), nil
}
