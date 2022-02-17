package runhuntingquery

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// RunHuntingQueryRequestBuilder builds and executes requests for operations under \security\microsoft.graph.security.runHuntingQuery
type RunHuntingQueryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RunHuntingQueryRequestBuilderPostOptions options for Post
type RunHuntingQueryRequestBuilderPostOptions struct {
    // 
    Body *RunHuntingQueryRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RunHuntingQueryResponse union type wrapper for classes huntingQueryResults
type RunHuntingQueryResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type huntingQueryResults
    huntingQueryResults *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HuntingQueryResults;
}
// NewRunHuntingQueryResponse instantiates a new runHuntingQueryResponse and sets the default values.
func NewRunHuntingQueryResponse()(*RunHuntingQueryResponse) {
    m := &RunHuntingQueryResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunHuntingQueryResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHuntingQueryResults gets the huntingQueryResults property value. Union type representation for type huntingQueryResults
func (m *RunHuntingQueryResponse) GetHuntingQueryResults()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HuntingQueryResults) {
    if m == nil {
        return nil
    } else {
        return m.huntingQueryResults
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RunHuntingQueryResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["huntingQueryResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewHuntingQueryResults() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHuntingQueryResults(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HuntingQueryResults))
        }
        return nil
    }
    return res
}
func (m *RunHuntingQueryResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RunHuntingQueryResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("huntingQueryResults", m.GetHuntingQueryResults())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunHuntingQueryResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHuntingQueryResults sets the huntingQueryResults property value. Union type representation for type huntingQueryResults
func (m *RunHuntingQueryResponse) SetHuntingQueryResults(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HuntingQueryResults)() {
    if m != nil {
        m.huntingQueryResults = value
    }
}
// NewRunHuntingQueryRequestBuilderInternal instantiates a new RunHuntingQueryRequestBuilder and sets the default values.
func NewRunHuntingQueryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RunHuntingQueryRequestBuilder) {
    m := &RunHuntingQueryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/microsoft.graph.security.runHuntingQuery";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRunHuntingQueryRequestBuilder instantiates a new RunHuntingQueryRequestBuilder and sets the default values.
func NewRunHuntingQueryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RunHuntingQueryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunHuntingQueryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action runHuntingQuery
func (m *RunHuntingQueryRequestBuilder) CreatePostRequestInformation(options *RunHuntingQueryRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action runHuntingQuery
func (m *RunHuntingQueryRequestBuilder) Post(options *RunHuntingQueryRequestBuilderPostOptions)(*RunHuntingQueryResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRunHuntingQueryResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RunHuntingQueryResponse), nil
}
