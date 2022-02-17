package makepermanent

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// MakePermanentRequestBuilder builds and executes requests for operations under \privilegedRoleAssignments\{privilegedRoleAssignment-id}\microsoft.graph.makePermanent
type MakePermanentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MakePermanentRequestBuilderPostOptions options for Post
type MakePermanentRequestBuilderPostOptions struct {
    // 
    Body *MakePermanentRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MakePermanentResponse union type wrapper for classes privilegedRoleAssignment
type MakePermanentResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type privilegedRoleAssignment
    privilegedRoleAssignment *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment;
}
// NewMakePermanentResponse instantiates a new makePermanentResponse and sets the default values.
func NewMakePermanentResponse()(*MakePermanentResponse) {
    m := &MakePermanentResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MakePermanentResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPrivilegedRoleAssignment gets the privilegedRoleAssignment property value. Union type representation for type privilegedRoleAssignment
func (m *MakePermanentResponse) GetPrivilegedRoleAssignment()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.privilegedRoleAssignment
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MakePermanentResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["privilegedRoleAssignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivilegedRoleAssignment(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment))
        }
        return nil
    }
    return res
}
func (m *MakePermanentResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MakePermanentResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("privilegedRoleAssignment", m.GetPrivilegedRoleAssignment())
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
func (m *MakePermanentResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPrivilegedRoleAssignment sets the privilegedRoleAssignment property value. Union type representation for type privilegedRoleAssignment
func (m *MakePermanentResponse) SetPrivilegedRoleAssignment(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment)() {
    if m != nil {
        m.privilegedRoleAssignment = value
    }
}
// NewMakePermanentRequestBuilderInternal instantiates a new MakePermanentRequestBuilder and sets the default values.
func NewMakePermanentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MakePermanentRequestBuilder) {
    m := &MakePermanentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignments/{privilegedRoleAssignment_id}/microsoft.graph.makePermanent";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMakePermanentRequestBuilder instantiates a new MakePermanentRequestBuilder and sets the default values.
func NewMakePermanentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MakePermanentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMakePermanentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action makePermanent
func (m *MakePermanentRequestBuilder) CreatePostRequestInformation(options *MakePermanentRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action makePermanent
func (m *MakePermanentRequestBuilder) Post(options *MakePermanentRequestBuilderPostOptions)(*MakePermanentResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMakePermanentResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*MakePermanentResponse), nil
}
