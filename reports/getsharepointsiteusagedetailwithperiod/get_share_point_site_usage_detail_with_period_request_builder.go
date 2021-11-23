package getsharepointsiteusagedetailwithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// getSharePointSiteUsageDetailWithPeriodRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageDetail(period='{period}')
type GetSharePointSiteUsageDetailWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetSharePointSiteUsageDetailWithPeriodRequestBuilderGetOptions options for Get
type GetSharePointSiteUsageDetailWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal instantiates a new GetSharePointSiteUsageDetailWithPeriodRequestBuilder and sets the default values.
func NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    m := &GetSharePointSiteUsageDetailWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getSharePointSiteUsageDetail(period='{period}')";
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
// NewGetSharePointSiteUsageDetailWithPeriodRequestBuilder instantiates a new GetSharePointSiteUsageDetailWithPeriodRequestBuilder and sets the default values.
func NewGetSharePointSiteUsageDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getSharePointSiteUsageDetail
func (m *GetSharePointSiteUsageDetailWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetSharePointSiteUsageDetailWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getSharePointSiteUsageDetail
func (m *GetSharePointSiteUsageDetailWithPeriodRequestBuilder) Get(options *GetSharePointSiteUsageDetailWithPeriodRequestBuilderGetOptions)([]GetSharePointSiteUsageDetailWithPeriod, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetSharePointSiteUsageDetailWithPeriod() }, nil)
    if err != nil {
        return nil, err
    }
    val := make([]GetSharePointSiteUsageDetailWithPeriod, len(res))
    for i, v := range res {
        val[i] = *(v.(*GetSharePointSiteUsageDetailWithPeriod))
    }
    return val, nil
}
