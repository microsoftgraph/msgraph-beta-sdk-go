package softwareupdatestatussummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4066a058652009641ce5d44658b8d2e68d34c551acdb8b9fa9d7e83202402fb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/softwareupdatestatussummary/ref"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// SoftwareUpdateStatusSummaryRequestBuilder builds and executes requests for operations under \deviceManagement\softwareUpdateStatusSummary
type SoftwareUpdateStatusSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SoftwareUpdateStatusSummaryRequestBuilderGetOptions options for Get
type SoftwareUpdateStatusSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters the software update status summary.
type SoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewSoftwareUpdateStatusSummaryRequestBuilderInternal instantiates a new SoftwareUpdateStatusSummaryRequestBuilder and sets the default values.
func NewSoftwareUpdateStatusSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SoftwareUpdateStatusSummaryRequestBuilder) {
    m := &SoftwareUpdateStatusSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/softwareUpdateStatusSummary{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSoftwareUpdateStatusSummaryRequestBuilder instantiates a new SoftwareUpdateStatusSummaryRequestBuilder and sets the default values.
func NewSoftwareUpdateStatusSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SoftwareUpdateStatusSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSoftwareUpdateStatusSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the software update status summary.
func (m *SoftwareUpdateStatusSummaryRequestBuilder) CreateGetRequestInformation(options *SoftwareUpdateStatusSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the software update status summary.
func (m *SoftwareUpdateStatusSummaryRequestBuilder) Get(options *SoftwareUpdateStatusSummaryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SoftwareUpdateStatusSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSoftwareUpdateStatusSummary() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SoftwareUpdateStatusSummary), nil
}
func (m *SoftwareUpdateStatusSummaryRequestBuilder) Ref()(*i4066a058652009641ce5d44658b8d2e68d34c551acdb8b9fa9d7e83202402fb2.RefRequestBuilder) {
    return i4066a058652009641ce5d44658b8d2e68d34c551acdb8b9fa9d7e83202402fb2.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
