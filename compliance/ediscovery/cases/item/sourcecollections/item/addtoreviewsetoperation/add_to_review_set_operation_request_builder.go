package addtoreviewsetoperation

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3e461ab8d8fb4c7edda4b883dc3fc5ec0a6c31a0ada6ed226bb5392ee265e579 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/addtoreviewsetoperation/ref"
)

// AddToReviewSetOperationRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\sourceCollections\{sourceCollection-id}\addToReviewSetOperation
type AddToReviewSetOperationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AddToReviewSetOperationRequestBuilderGetOptions options for Get
type AddToReviewSetOperationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AddToReviewSetOperationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AddToReviewSetOperationRequestBuilderGetQueryParameters adds the results of the sourceCollection to the specified reviewSet.
type AddToReviewSetOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewAddToReviewSetOperationRequestBuilderInternal instantiates a new AddToReviewSetOperationRequestBuilder and sets the default values.
func NewAddToReviewSetOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AddToReviewSetOperationRequestBuilder) {
    m := &AddToReviewSetOperationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}/addToReviewSetOperation{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAddToReviewSetOperationRequestBuilder instantiates a new AddToReviewSetOperationRequestBuilder and sets the default values.
func NewAddToReviewSetOperationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AddToReviewSetOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAddToReviewSetOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation adds the results of the sourceCollection to the specified reviewSet.
func (m *AddToReviewSetOperationRequestBuilder) CreateGetRequestInformation(options *AddToReviewSetOperationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get adds the results of the sourceCollection to the specified reviewSet.
func (m *AddToReviewSetOperationRequestBuilder) Get(options *AddToReviewSetOperationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AddToReviewSetOperation, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAddToReviewSetOperation() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AddToReviewSetOperation), nil
}
func (m *AddToReviewSetOperationRequestBuilder) Ref()(*i3e461ab8d8fb4c7edda4b883dc3fc5ec0a6c31a0ada6ed226bb5392ee265e579.RefRequestBuilder) {
    return i3e461ab8d8fb4c7edda4b883dc3fc5ec0a6c31a0ada6ed226bb5392ee265e579.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
