package validatefilter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ValidateFilterRequestBuilder builds and executes requests for operations under \deviceManagement\assignmentFilters\microsoft.graph.validateFilter
type ValidateFilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ValidateFilterRequestBuilderPostOptions options for Post
type ValidateFilterRequestBuilderPostOptions struct {
    // 
    Body *ValidateFilterRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ValidateFilterResponse union type wrapper for classes assignmentFilterValidationResult
type ValidateFilterResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type assignmentFilterValidationResult
    assignmentFilterValidationResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterValidationResult;
}
// NewValidateFilterResponse instantiates a new validateFilterResponse and sets the default values.
func NewValidateFilterResponse()(*ValidateFilterResponse) {
    m := &ValidateFilterResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateFilterResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignmentFilterValidationResult gets the assignmentFilterValidationResult property value. Union type representation for type assignmentFilterValidationResult
func (m *ValidateFilterResponse) GetAssignmentFilterValidationResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterValidationResult) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterValidationResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidateFilterResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterValidationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAssignmentFilterValidationResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterValidationResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterValidationResult))
        }
        return nil
    }
    return res
}
func (m *ValidateFilterResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ValidateFilterResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("assignmentFilterValidationResult", m.GetAssignmentFilterValidationResult())
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
func (m *ValidateFilterResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignmentFilterValidationResult sets the assignmentFilterValidationResult property value. Union type representation for type assignmentFilterValidationResult
func (m *ValidateFilterResponse) SetAssignmentFilterValidationResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterValidationResult)() {
    if m != nil {
        m.assignmentFilterValidationResult = value
    }
}
// NewValidateFilterRequestBuilderInternal instantiates a new ValidateFilterRequestBuilder and sets the default values.
func NewValidateFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValidateFilterRequestBuilder) {
    m := &ValidateFilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/assignmentFilters/microsoft.graph.validateFilter";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewValidateFilterRequestBuilder instantiates a new ValidateFilterRequestBuilder and sets the default values.
func NewValidateFilterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValidateFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewValidateFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action validateFilter
func (m *ValidateFilterRequestBuilder) CreatePostRequestInformation(options *ValidateFilterRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action validateFilter
func (m *ValidateFilterRequestBuilder) Post(options *ValidateFilterRequestBuilderPostOptions)(*ValidateFilterResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewValidateFilterResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ValidateFilterResponse), nil
}
