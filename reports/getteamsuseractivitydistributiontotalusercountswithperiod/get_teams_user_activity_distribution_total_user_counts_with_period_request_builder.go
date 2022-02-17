package getteamsuseractivitydistributiontotalusercountswithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityDistributionTotalUserCounts(period='{period}')
type GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetOptions options for Get
type GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal instantiates a new GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    m := &GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsUserActivityDistributionTotalUserCounts(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder instantiates a new GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getTeamsUserActivityDistributionTotalUserCounts
func (m *GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getTeamsUserActivityDistributionTotalUserCounts
func (m *GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) Get(options *GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetOptions)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "byte", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
