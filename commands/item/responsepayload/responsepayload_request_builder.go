package responsepayload

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iba18db208a8013762074839c8fb77f8d0fc88e67d4c6d1d663f90ddd47f87754 "github.com/microsoftgraph/msgraph-beta-sdk-go/commands/item/responsepayload/ref"
)

// ResponsepayloadRequestBuilder builds and executes requests for operations under \commands\{command-id}\responsepayload
type ResponsepayloadRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResponsepayloadRequestBuilderGetOptions options for Get
type ResponsepayloadRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ResponsepayloadRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResponsepayloadRequestBuilderGetQueryParameters get responsepayload from commands
type ResponsepayloadRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewResponsepayloadRequestBuilderInternal instantiates a new ResponsepayloadRequestBuilder and sets the default values.
func NewResponsepayloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResponsepayloadRequestBuilder) {
    m := &ResponsepayloadRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/commands/{command_id}/responsepayload{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResponsepayloadRequestBuilder instantiates a new ResponsepayloadRequestBuilder and sets the default values.
func NewResponsepayloadRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResponsepayloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResponsepayloadRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get responsepayload from commands
func (m *ResponsepayloadRequestBuilder) CreateGetRequestInformation(options *ResponsepayloadRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get responsepayload from commands
func (m *ResponsepayloadRequestBuilder) Get(options *ResponsepayloadRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PayloadResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPayloadResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PayloadResponse), nil
}
func (m *ResponsepayloadRequestBuilder) Ref()(*iba18db208a8013762074839c8fb77f8d0fc88e67d4c6d1d663f90ddd47f87754.RefRequestBuilder) {
    return iba18db208a8013762074839c8fb77f8d0fc88e67d4c6d1d663f90ddd47f87754.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
