package assigntag

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// AssignTagRequestBuilder provides operations to call the assignTag method.
type AssignTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AssignTagRequestBuilderPostOptions options for Post
type AssignTagRequestBuilderPostOptions struct {
    // 
    Body AssignTagRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AssignTagResponse union type wrapper for classes tenantTag
type AssignTagResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type tenantTag
    tenantTag i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantTagable;
}
// NewAssignTagResponse instantiates a new assignTagResponse and sets the default values.
func NewAssignTagResponse()(*AssignTagResponse) {
    m := &AssignTagResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateAssignTagResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAssignTagResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignTagResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignTagResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tenantTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateTenantTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantTag(val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantTagable))
        }
        return nil
    }
    return res
}
// GetTenantTag gets the tenantTag property value. Union type representation for type tenantTag
func (m *AssignTagResponse) GetTenantTag()(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantTagable) {
    if m == nil {
        return nil
    } else {
        return m.tenantTag
    }
}
func (m *AssignTagResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignTagResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("tenantTag", m.GetTenantTag())
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
func (m *AssignTagResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTenantTag sets the tenantTag property value. Union type representation for type tenantTag
func (m *AssignTagResponse) SetTenantTag(value i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantTagable)() {
    if m != nil {
        m.tenantTag = value
    }
}
// AssignTagResponseable 
type AssignTagResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetTenantTag()(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantTagable)
    SetTenantTag(value i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantTagable)()
}
// NewAssignTagRequestBuilderInternal instantiates a new AssignTagRequestBuilder and sets the default values.
func NewAssignTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AssignTagRequestBuilder) {
    m := &AssignTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenantTags/{tenantTag_id}/microsoft.graph.managedTenants.assignTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAssignTagRequestBuilder instantiates a new AssignTagRequestBuilder and sets the default values.
func NewAssignTagRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AssignTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action assignTag
func (m *AssignTagRequestBuilder) CreatePostRequestInformation(options *AssignTagRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action assignTag
func (m *AssignTagRequestBuilder) Post(options *AssignTagRequestBuilderPostOptions)(AssignTagResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateAssignTagResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(AssignTagResponseable), nil
}
