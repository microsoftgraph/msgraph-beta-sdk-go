package importresourceactions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ImportResourceActionsRequestBuilder provides operations to call the importResourceActions method.
type ImportResourceActionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ImportResourceActionsRequestBuilderPostOptions options for Post
type ImportResourceActionsRequestBuilderPostOptions struct {
    // 
    Body ImportResourceActionsRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ImportResourceActionsResponse union type wrapper for classes unifiedRbacResourceNamespace
type ImportResourceActionsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type unifiedRbacResourceNamespace
    unifiedRbacResourceNamespace i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespaceable;
}
// NewImportResourceActionsResponse instantiates a new importResourceActionsResponse and sets the default values.
func NewImportResourceActionsResponse()(*ImportResourceActionsResponse) {
    m := &ImportResourceActionsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateImportResourceActionsResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewImportResourceActionsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportResourceActionsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportResourceActionsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["unifiedRbacResourceNamespace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnifiedRbacResourceNamespace(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespaceable))
        }
        return nil
    }
    return res
}
// GetUnifiedRbacResourceNamespace gets the unifiedRbacResourceNamespace property value. Union type representation for type unifiedRbacResourceNamespace
func (m *ImportResourceActionsResponse) GetUnifiedRbacResourceNamespace()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespaceable) {
    if m == nil {
        return nil
    } else {
        return m.unifiedRbacResourceNamespace
    }
}
func (m *ImportResourceActionsResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ImportResourceActionsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("unifiedRbacResourceNamespace", m.GetUnifiedRbacResourceNamespace())
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
func (m *ImportResourceActionsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUnifiedRbacResourceNamespace sets the unifiedRbacResourceNamespace property value. Union type representation for type unifiedRbacResourceNamespace
func (m *ImportResourceActionsResponse) SetUnifiedRbacResourceNamespace(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespaceable)() {
    if m != nil {
        m.unifiedRbacResourceNamespace = value
    }
}
// ImportResourceActionsResponseable 
type ImportResourceActionsResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetUnifiedRbacResourceNamespace()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespaceable)
    SetUnifiedRbacResourceNamespace(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespaceable)()
}
// NewImportResourceActionsRequestBuilderInternal instantiates a new ImportResourceActionsRequestBuilder and sets the default values.
func NewImportResourceActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportResourceActionsRequestBuilder) {
    m := &ImportResourceActionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/resourceNamespaces/{unifiedRbacResourceNamespace_id}/microsoft.graph.importResourceActions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportResourceActionsRequestBuilder instantiates a new ImportResourceActionsRequestBuilder and sets the default values.
func NewImportResourceActionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportResourceActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportResourceActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action importResourceActions
func (m *ImportResourceActionsRequestBuilder) CreatePostRequestInformation(options *ImportResourceActionsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action importResourceActions
func (m *ImportResourceActionsRequestBuilder) Post(options *ImportResourceActionsRequestBuilderPostOptions)(ImportResourceActionsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateImportResourceActionsResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(ImportResourceActionsResponseable), nil
}
