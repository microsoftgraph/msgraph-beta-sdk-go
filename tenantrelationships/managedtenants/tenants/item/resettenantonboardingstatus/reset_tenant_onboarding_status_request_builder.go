package resettenantonboardingstatus

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ResetTenantOnboardingStatusRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\tenants\{tenant-id}\microsoft.graph.managedTenants.resetTenantOnboardingStatus
type ResetTenantOnboardingStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResetTenantOnboardingStatusRequestBuilderPostOptions options for Post
type ResetTenantOnboardingStatusRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResetTenantOnboardingStatusResponse union type wrapper for classes tenant
type ResetTenantOnboardingStatusResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type tenant
    tenant *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Tenant;
}
// NewResetTenantOnboardingStatusResponse instantiates a new resetTenantOnboardingStatusResponse and sets the default values.
func NewResetTenantOnboardingStatusResponse()(*ResetTenantOnboardingStatusResponse) {
    m := &ResetTenantOnboardingStatusResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetTenantOnboardingStatusResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTenant gets the tenant property value. Union type representation for type tenant
func (m *ResetTenantOnboardingStatusResponse) GetTenant()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Tenant) {
    if m == nil {
        return nil
    } else {
        return m.tenant
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResetTenantOnboardingStatusResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tenant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTenant() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenant(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Tenant))
        }
        return nil
    }
    return res
}
func (m *ResetTenantOnboardingStatusResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResetTenantOnboardingStatusResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("tenant", m.GetTenant())
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
func (m *ResetTenantOnboardingStatusResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTenant sets the tenant property value. Union type representation for type tenant
func (m *ResetTenantOnboardingStatusResponse) SetTenant(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Tenant)() {
    if m != nil {
        m.tenant = value
    }
}
// NewResetTenantOnboardingStatusRequestBuilderInternal instantiates a new ResetTenantOnboardingStatusRequestBuilder and sets the default values.
func NewResetTenantOnboardingStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResetTenantOnboardingStatusRequestBuilder) {
    m := &ResetTenantOnboardingStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenants/{tenant_id}/microsoft.graph.managedTenants.resetTenantOnboardingStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResetTenantOnboardingStatusRequestBuilder instantiates a new ResetTenantOnboardingStatusRequestBuilder and sets the default values.
func NewResetTenantOnboardingStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResetTenantOnboardingStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResetTenantOnboardingStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action resetTenantOnboardingStatus
func (m *ResetTenantOnboardingStatusRequestBuilder) CreatePostRequestInformation(options *ResetTenantOnboardingStatusRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action resetTenantOnboardingStatus
func (m *ResetTenantOnboardingStatusRequestBuilder) Post(options *ResetTenantOnboardingStatusRequestBuilderPostOptions)(*ResetTenantOnboardingStatusResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResetTenantOnboardingStatusResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ResetTenantOnboardingStatusResponse), nil
}
