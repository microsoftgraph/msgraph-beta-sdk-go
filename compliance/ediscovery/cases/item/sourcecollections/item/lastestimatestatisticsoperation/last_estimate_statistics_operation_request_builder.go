package lastestimatestatisticsoperation

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia17a50f392f233475f2726c08e20bb925fe8c75dc2a18a8c6c889953892e8e2c "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/lastestimatestatisticsoperation/ref"
)

// LastEstimateStatisticsOperationRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\sourceCollections\{sourceCollection-id}\lastEstimateStatisticsOperation
type LastEstimateStatisticsOperationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// LastEstimateStatisticsOperationRequestBuilderGetOptions options for Get
type LastEstimateStatisticsOperationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *LastEstimateStatisticsOperationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// LastEstimateStatisticsOperationRequestBuilderGetQueryParameters the last estimate operation associated with the sourceCollection.
type LastEstimateStatisticsOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewLastEstimateStatisticsOperationRequestBuilderInternal instantiates a new LastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewLastEstimateStatisticsOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LastEstimateStatisticsOperationRequestBuilder) {
    m := &LastEstimateStatisticsOperationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}/lastEstimateStatisticsOperation{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLastEstimateStatisticsOperationRequestBuilder instantiates a new LastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewLastEstimateStatisticsOperationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LastEstimateStatisticsOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLastEstimateStatisticsOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the last estimate operation associated with the sourceCollection.
func (m *LastEstimateStatisticsOperationRequestBuilder) CreateGetRequestInformation(options *LastEstimateStatisticsOperationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the last estimate operation associated with the sourceCollection.
func (m *LastEstimateStatisticsOperationRequestBuilder) Get(options *LastEstimateStatisticsOperationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EstimateStatisticsOperation, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEstimateStatisticsOperation() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EstimateStatisticsOperation), nil
}
func (m *LastEstimateStatisticsOperationRequestBuilder) Ref()(*ia17a50f392f233475f2726c08e20bb925fe8c75dc2a18a8c6c889953892e8e2c.RefRequestBuilder) {
    return ia17a50f392f233475f2726c08e20bb925fe8c75dc2a18a8c6c889953892e8e2c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
