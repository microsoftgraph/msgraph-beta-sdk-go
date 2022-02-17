package makeeligible

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// MakeEligibleRequestBuilder builds and executes requests for operations under \privilegedRoleAssignments\{privilegedRoleAssignment-id}\microsoft.graph.makeEligible
type MakeEligibleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MakeEligibleRequestBuilderPostOptions options for Post
type MakeEligibleRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MakeEligibleResponse union type wrapper for classes privilegedRoleAssignment
type MakeEligibleResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type privilegedRoleAssignment
    privilegedRoleAssignment *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment;
}
// NewMakeEligibleResponse instantiates a new makeEligibleResponse and sets the default values.
func NewMakeEligibleResponse()(*MakeEligibleResponse) {
    m := &MakeEligibleResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MakeEligibleResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPrivilegedRoleAssignment gets the privilegedRoleAssignment property value. Union type representation for type privilegedRoleAssignment
func (m *MakeEligibleResponse) GetPrivilegedRoleAssignment()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.privilegedRoleAssignment
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MakeEligibleResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *MakeEligibleResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MakeEligibleResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *MakeEligibleResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPrivilegedRoleAssignment sets the privilegedRoleAssignment property value. Union type representation for type privilegedRoleAssignment
func (m *MakeEligibleResponse) SetPrivilegedRoleAssignment(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment)() {
    if m != nil {
        m.privilegedRoleAssignment = value
    }
}
// NewMakeEligibleRequestBuilderInternal instantiates a new MakeEligibleRequestBuilder and sets the default values.
func NewMakeEligibleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MakeEligibleRequestBuilder) {
    m := &MakeEligibleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignments/{privilegedRoleAssignment_id}/microsoft.graph.makeEligible";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMakeEligibleRequestBuilder instantiates a new MakeEligibleRequestBuilder and sets the default values.
func NewMakeEligibleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MakeEligibleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMakeEligibleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action makeEligible
func (m *MakeEligibleRequestBuilder) CreatePostRequestInformation(options *MakeEligibleRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action makeEligible
func (m *MakeEligibleRequestBuilder) Post(options *MakeEligibleRequestBuilderPostOptions)(*MakeEligibleResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMakeEligibleResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*MakeEligibleResponse), nil
}
