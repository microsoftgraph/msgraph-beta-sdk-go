package getmanagementconditionstatementexpressionstring

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetManagementConditionStatementExpressionStringRequestBuilder builds and executes requests for operations under \deviceManagement\managementConditionStatements\{managementConditionStatement-id}\microsoft.graph.getManagementConditionStatementExpressionString()
type GetManagementConditionStatementExpressionStringRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetManagementConditionStatementExpressionStringRequestBuilderGetOptions options for Get
type GetManagementConditionStatementExpressionStringRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetManagementConditionStatementExpressionStringResponse union type wrapper for classes managementConditionExpressionString
type GetManagementConditionStatementExpressionStringResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type managementConditionExpressionString
    managementConditionExpressionString *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementConditionExpressionString;
}
// NewGetManagementConditionStatementExpressionStringResponse instantiates a new getManagementConditionStatementExpressionStringResponse and sets the default values.
func NewGetManagementConditionStatementExpressionStringResponse()(*GetManagementConditionStatementExpressionStringResponse) {
    m := &GetManagementConditionStatementExpressionStringResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetManagementConditionStatementExpressionStringResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetManagementConditionExpressionString gets the managementConditionExpressionString property value. Union type representation for type managementConditionExpressionString
func (m *GetManagementConditionStatementExpressionStringResponse) GetManagementConditionExpressionString()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementConditionExpressionString) {
    if m == nil {
        return nil
    } else {
        return m.managementConditionExpressionString
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetManagementConditionStatementExpressionStringResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementConditionExpressionString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagementConditionExpressionString() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementConditionExpressionString(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementConditionExpressionString))
        }
        return nil
    }
    return res
}
func (m *GetManagementConditionStatementExpressionStringResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetManagementConditionStatementExpressionStringResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("managementConditionExpressionString", m.GetManagementConditionExpressionString())
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
func (m *GetManagementConditionStatementExpressionStringResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagementConditionExpressionString sets the managementConditionExpressionString property value. Union type representation for type managementConditionExpressionString
func (m *GetManagementConditionStatementExpressionStringResponse) SetManagementConditionExpressionString(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementConditionExpressionString)() {
    if m != nil {
        m.managementConditionExpressionString = value
    }
}
// NewGetManagementConditionStatementExpressionStringRequestBuilderInternal instantiates a new GetManagementConditionStatementExpressionStringRequestBuilder and sets the default values.
func NewGetManagementConditionStatementExpressionStringRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetManagementConditionStatementExpressionStringRequestBuilder) {
    m := &GetManagementConditionStatementExpressionStringRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditionStatements/{managementConditionStatement_id}/microsoft.graph.getManagementConditionStatementExpressionString()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagementConditionStatementExpressionStringRequestBuilder instantiates a new GetManagementConditionStatementExpressionStringRequestBuilder and sets the default values.
func NewGetManagementConditionStatementExpressionStringRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetManagementConditionStatementExpressionStringRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagementConditionStatementExpressionStringRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getManagementConditionStatementExpressionString
func (m *GetManagementConditionStatementExpressionStringRequestBuilder) CreateGetRequestInformation(options *GetManagementConditionStatementExpressionStringRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getManagementConditionStatementExpressionString
func (m *GetManagementConditionStatementExpressionStringRequestBuilder) Get(options *GetManagementConditionStatementExpressionStringRequestBuilderGetOptions)(*GetManagementConditionStatementExpressionStringResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetManagementConditionStatementExpressionStringResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetManagementConditionStatementExpressionStringResponse), nil
}
