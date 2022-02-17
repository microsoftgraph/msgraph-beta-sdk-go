package dailyprintusagesummariesbyprinter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// DailyPrintUsageSummariesByPrinterRequestBuilder builds and executes requests for operations under \reports\dailyPrintUsageSummariesByPrinter
type DailyPrintUsageSummariesByPrinterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DailyPrintUsageSummariesByPrinterRequestBuilderGetOptions options for Get
type DailyPrintUsageSummariesByPrinterRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DailyPrintUsageSummariesByPrinterRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DailyPrintUsageSummariesByPrinterRequestBuilderGetQueryParameters get dailyPrintUsageSummariesByPrinter from reports
type DailyPrintUsageSummariesByPrinterRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// DailyPrintUsageSummariesByPrinterRequestBuilderPostOptions options for Post
type DailyPrintUsageSummariesByPrinterRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrintUsageByPrinter;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal instantiates a new DailyPrintUsageSummariesByPrinterRequestBuilder and sets the default values.
func NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DailyPrintUsageSummariesByPrinterRequestBuilder) {
    m := &DailyPrintUsageSummariesByPrinterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/dailyPrintUsageSummariesByPrinter{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDailyPrintUsageSummariesByPrinterRequestBuilder instantiates a new DailyPrintUsageSummariesByPrinterRequestBuilder and sets the default values.
func NewDailyPrintUsageSummariesByPrinterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DailyPrintUsageSummariesByPrinterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get dailyPrintUsageSummariesByPrinter from reports
func (m *DailyPrintUsageSummariesByPrinterRequestBuilder) CreateGetRequestInformation(options *DailyPrintUsageSummariesByPrinterRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// CreatePostRequestInformation create new navigation property to dailyPrintUsageSummariesByPrinter for reports
func (m *DailyPrintUsageSummariesByPrinterRequestBuilder) CreatePostRequestInformation(options *DailyPrintUsageSummariesByPrinterRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Get get dailyPrintUsageSummariesByPrinter from reports
func (m *DailyPrintUsageSummariesByPrinterRequestBuilder) Get(options *DailyPrintUsageSummariesByPrinterRequestBuilderGetOptions)(*DailyPrintUsageSummariesByPrinterResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDailyPrintUsageSummariesByPrinterResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DailyPrintUsageSummariesByPrinterResponse), nil
}
// Post create new navigation property to dailyPrintUsageSummariesByPrinter for reports
func (m *DailyPrintUsageSummariesByPrinterRequestBuilder) Post(options *DailyPrintUsageSummariesByPrinterRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrintUsageByPrinter, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrintUsageByPrinter() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrintUsageByPrinter), nil
}
