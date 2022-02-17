package changedeploymentstatus

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ChangeDeploymentStatusRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\managementTemplateStepVersions\{managementTemplateStepVersion-id}\deployments\{managementTemplateStepDeployment-id}\microsoft.graph.managedTenants.changeDeploymentStatus
type ChangeDeploymentStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChangeDeploymentStatusRequestBuilderPostOptions options for Post
type ChangeDeploymentStatusRequestBuilderPostOptions struct {
    // 
    Body *ChangeDeploymentStatusRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChangeDeploymentStatusResponse union type wrapper for classes managementTemplateStepDeployment
type ChangeDeploymentStatusResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type managementTemplateStepDeployment
    managementTemplateStepDeployment *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplateStepDeployment;
}
// NewChangeDeploymentStatusResponse instantiates a new changeDeploymentStatusResponse and sets the default values.
func NewChangeDeploymentStatusResponse()(*ChangeDeploymentStatusResponse) {
    m := &ChangeDeploymentStatusResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeDeploymentStatusResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetManagementTemplateStepDeployment gets the managementTemplateStepDeployment property value. Union type representation for type managementTemplateStepDeployment
func (m *ChangeDeploymentStatusResponse) GetManagementTemplateStepDeployment()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplateStepDeployment) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateStepDeployment
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeDeploymentStatusResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementTemplateStepDeployment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagementTemplateStepDeployment() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepDeployment(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplateStepDeployment))
        }
        return nil
    }
    return res
}
func (m *ChangeDeploymentStatusResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChangeDeploymentStatusResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("managementTemplateStepDeployment", m.GetManagementTemplateStepDeployment())
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
func (m *ChangeDeploymentStatusResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagementTemplateStepDeployment sets the managementTemplateStepDeployment property value. Union type representation for type managementTemplateStepDeployment
func (m *ChangeDeploymentStatusResponse) SetManagementTemplateStepDeployment(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplateStepDeployment)() {
    if m != nil {
        m.managementTemplateStepDeployment = value
    }
}
// NewChangeDeploymentStatusRequestBuilderInternal instantiates a new ChangeDeploymentStatusRequestBuilder and sets the default values.
func NewChangeDeploymentStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChangeDeploymentStatusRequestBuilder) {
    m := &ChangeDeploymentStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion_id}/deployments/{managementTemplateStepDeployment_id}/microsoft.graph.managedTenants.changeDeploymentStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChangeDeploymentStatusRequestBuilder instantiates a new ChangeDeploymentStatusRequestBuilder and sets the default values.
func NewChangeDeploymentStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChangeDeploymentStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChangeDeploymentStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action changeDeploymentStatus
func (m *ChangeDeploymentStatusRequestBuilder) CreatePostRequestInformation(options *ChangeDeploymentStatusRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action changeDeploymentStatus
func (m *ChangeDeploymentStatusRequestBuilder) Post(options *ChangeDeploymentStatusRequestBuilderPostOptions)(*ChangeDeploymentStatusResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChangeDeploymentStatusResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ChangeDeploymentStatusResponse), nil
}
