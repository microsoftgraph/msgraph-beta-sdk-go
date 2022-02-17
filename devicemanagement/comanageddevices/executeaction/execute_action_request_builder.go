package executeaction

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExecuteActionRequestBuilder builds and executes requests for operations under \deviceManagement\comanagedDevices\microsoft.graph.executeAction
type ExecuteActionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExecuteActionRequestBuilderPostOptions options for Post
type ExecuteActionRequestBuilderPostOptions struct {
    // 
    Body *ExecuteActionRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExecuteActionResponse union type wrapper for classes bulkManagedDeviceActionResult
type ExecuteActionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type bulkManagedDeviceActionResult
    bulkManagedDeviceActionResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BulkManagedDeviceActionResult;
}
// NewExecuteActionResponse instantiates a new executeActionResponse and sets the default values.
func NewExecuteActionResponse()(*ExecuteActionResponse) {
    m := &ExecuteActionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExecuteActionResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBulkManagedDeviceActionResult gets the bulkManagedDeviceActionResult property value. Union type representation for type bulkManagedDeviceActionResult
func (m *ExecuteActionResponse) GetBulkManagedDeviceActionResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BulkManagedDeviceActionResult) {
    if m == nil {
        return nil
    } else {
        return m.bulkManagedDeviceActionResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExecuteActionResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bulkManagedDeviceActionResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBulkManagedDeviceActionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBulkManagedDeviceActionResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BulkManagedDeviceActionResult))
        }
        return nil
    }
    return res
}
func (m *ExecuteActionResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExecuteActionResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bulkManagedDeviceActionResult", m.GetBulkManagedDeviceActionResult())
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
func (m *ExecuteActionResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBulkManagedDeviceActionResult sets the bulkManagedDeviceActionResult property value. Union type representation for type bulkManagedDeviceActionResult
func (m *ExecuteActionResponse) SetBulkManagedDeviceActionResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BulkManagedDeviceActionResult)() {
    if m != nil {
        m.bulkManagedDeviceActionResult = value
    }
}
// NewExecuteActionRequestBuilderInternal instantiates a new ExecuteActionRequestBuilder and sets the default values.
func NewExecuteActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExecuteActionRequestBuilder) {
    m := &ExecuteActionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/microsoft.graph.executeAction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExecuteActionRequestBuilder instantiates a new ExecuteActionRequestBuilder and sets the default values.
func NewExecuteActionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExecuteActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecuteActionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action executeAction
func (m *ExecuteActionRequestBuilder) CreatePostRequestInformation(options *ExecuteActionRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action executeAction
func (m *ExecuteActionRequestBuilder) Post(options *ExecuteActionRequestBuilderPostOptions)(*ExecuteActionResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExecuteActionResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ExecuteActionResponse), nil
}
