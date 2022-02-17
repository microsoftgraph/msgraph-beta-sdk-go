package updaterequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UpdateRequestRequestBuilder builds and executes requests for operations under \governanceRoleAssignmentRequests\{governanceRoleAssignmentRequest-id}\microsoft.graph.updateRequest
type UpdateRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdateRequestRequestBuilderPostOptions options for Post
type UpdateRequestRequestBuilderPostOptions struct {
    // 
    Body *UpdateRequestRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdateRequestResponse union type wrapper for classes governanceRoleAssignmentRequest
type UpdateRequestResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type governanceRoleAssignmentRequest
    governanceRoleAssignmentRequest *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest;
}
// NewUpdateRequestResponse instantiates a new updateRequestResponse and sets the default values.
func NewUpdateRequestResponse()(*UpdateRequestResponse) {
    m := &UpdateRequestResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateRequestResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetGovernanceRoleAssignmentRequest gets the governanceRoleAssignmentRequest property value. Union type representation for type governanceRoleAssignmentRequest
func (m *UpdateRequestResponse) GetGovernanceRoleAssignmentRequest()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.governanceRoleAssignmentRequest
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateRequestResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["governanceRoleAssignmentRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGovernanceRoleAssignmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGovernanceRoleAssignmentRequest(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest))
        }
        return nil
    }
    return res
}
func (m *UpdateRequestResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateRequestResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("governanceRoleAssignmentRequest", m.GetGovernanceRoleAssignmentRequest())
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
func (m *UpdateRequestResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGovernanceRoleAssignmentRequest sets the governanceRoleAssignmentRequest property value. Union type representation for type governanceRoleAssignmentRequest
func (m *UpdateRequestResponse) SetGovernanceRoleAssignmentRequest(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest)() {
    if m != nil {
        m.governanceRoleAssignmentRequest = value
    }
}
// NewUpdateRequestRequestBuilderInternal instantiates a new UpdateRequestRequestBuilder and sets the default values.
func NewUpdateRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateRequestRequestBuilder) {
    m := &UpdateRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceRoleAssignmentRequests/{governanceRoleAssignmentRequest_id}/microsoft.graph.updateRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateRequestRequestBuilder instantiates a new UpdateRequestRequestBuilder and sets the default values.
func NewUpdateRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateRequest
func (m *UpdateRequestRequestBuilder) CreatePostRequestInformation(options *UpdateRequestRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action updateRequest
func (m *UpdateRequestRequestBuilder) Post(options *UpdateRequestRequestBuilderPostOptions)(*UpdateRequestResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdateRequestResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UpdateRequestResponse), nil
}
