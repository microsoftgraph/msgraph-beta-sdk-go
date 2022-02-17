package getskypeforbusinessparticipantactivitycountswithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessParticipantActivityCounts(period='{period}')
type GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetOptions options for Get
type GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal instantiates a new GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    m := &GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getSkypeForBusinessParticipantActivityCounts(period='{period}')";
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
// NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder instantiates a new GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getSkypeForBusinessParticipantActivityCounts
func (m *GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getSkypeForBusinessParticipantActivityCounts
func (m *GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) Get(options *GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetOptions)([]byte, error) {
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
