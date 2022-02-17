package getoffice365activeuserdetailwithdate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOffice365ActiveUserDetailWithDateRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActiveUserDetail(date={date})
type GetOffice365ActiveUserDetailWithDateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetOffice365ActiveUserDetailWithDateRequestBuilderGetOptions options for Get
type GetOffice365ActiveUserDetailWithDateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal instantiates a new GetOffice365ActiveUserDetailWithDateRequestBuilder and sets the default values.
func NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*GetOffice365ActiveUserDetailWithDateRequestBuilder) {
    m := &GetOffice365ActiveUserDetailWithDateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365ActiveUserDetail(date={date})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if date != nil {
        urlTplParams["date"] = (*date).String()
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOffice365ActiveUserDetailWithDateRequestBuilder instantiates a new GetOffice365ActiveUserDetailWithDateRequestBuilder and sets the default values.
func NewGetOffice365ActiveUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetOffice365ActiveUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getOffice365ActiveUserDetail
func (m *GetOffice365ActiveUserDetailWithDateRequestBuilder) CreateGetRequestInformation(options *GetOffice365ActiveUserDetailWithDateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getOffice365ActiveUserDetail
func (m *GetOffice365ActiveUserDetailWithDateRequestBuilder) Get(options *GetOffice365ActiveUserDetailWithDateRequestBuilderGetOptions)([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365ActiveUserDetail, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOffice365ActiveUserDetail() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365ActiveUserDetail, len(res))
    for i, v := range res {
        val[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Office365ActiveUserDetail))
    }
    return val, nil
}
