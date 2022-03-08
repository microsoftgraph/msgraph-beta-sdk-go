package getteamsteamactivitycountswithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetTeamsTeamActivityCountsWithPeriodRequestBuilder provides operations to call the getTeamsTeamActivityCounts method.
type GetTeamsTeamActivityCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetTeamsTeamActivityCountsWithPeriodRequestBuilderGetOptions options for Get
type GetTeamsTeamActivityCountsWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal instantiates a new GetTeamsTeamActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    m := &GetTeamsTeamActivityCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsTeamActivityCounts(period='{period}')";
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
// NewGetTeamsTeamActivityCountsWithPeriodRequestBuilder instantiates a new GetTeamsTeamActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getTeamsTeamActivityCounts
func (m *GetTeamsTeamActivityCountsWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetTeamsTeamActivityCountsWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getTeamsTeamActivityCounts
func (m *GetTeamsTeamActivityCountsWithPeriodRequestBuilder) Get(options *GetTeamsTeamActivityCountsWithPeriodRequestBuilderGetOptions)(GetTeamsTeamActivityCountsWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTeamsTeamActivityCountsWithPeriodResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTeamsTeamActivityCountsWithPeriodResponseable), nil
}
