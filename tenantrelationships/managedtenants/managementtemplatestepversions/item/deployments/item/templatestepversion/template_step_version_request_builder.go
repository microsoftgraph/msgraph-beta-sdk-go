package templatestepversion

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4261fb17e0c76843a152f58b34dc951cd1cd3b513cd9ff4370402cbe2d7ccd76 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item/deployments/item/templatestepversion/ref"
    ic3ac3b84723781bdf65b34095cd8952b449b29f28b5193fc410efb700a4d0b4c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item/deployments/item/templatestepversion/deploy"
)

// TemplateStepVersionRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\managementTemplateStepVersions\{managementTemplateStepVersion-id}\deployments\{managementTemplateStepDeployment-id}\templateStepVersion
type TemplateStepVersionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TemplateStepVersionRequestBuilderGetOptions options for Get
type TemplateStepVersionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TemplateStepVersionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TemplateStepVersionRequestBuilderGetQueryParameters get templateStepVersion from tenantRelationships
type TemplateStepVersionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewTemplateStepVersionRequestBuilderInternal instantiates a new TemplateStepVersionRequestBuilder and sets the default values.
func NewTemplateStepVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TemplateStepVersionRequestBuilder) {
    m := &TemplateStepVersionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion_id}/deployments/{managementTemplateStepDeployment_id}/templateStepVersion{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTemplateStepVersionRequestBuilder instantiates a new TemplateStepVersionRequestBuilder and sets the default values.
func NewTemplateStepVersionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TemplateStepVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplateStepVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get templateStepVersion from tenantRelationships
func (m *TemplateStepVersionRequestBuilder) CreateGetRequestInformation(options *TemplateStepVersionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
func (m *TemplateStepVersionRequestBuilder) Deploy()(*ic3ac3b84723781bdf65b34095cd8952b449b29f28b5193fc410efb700a4d0b4c.DeployRequestBuilder) {
    return ic3ac3b84723781bdf65b34095cd8952b449b29f28b5193fc410efb700a4d0b4c.NewDeployRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get templateStepVersion from tenantRelationships
func (m *TemplateStepVersionRequestBuilder) Get(options *TemplateStepVersionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplateStepVersion, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagementTemplateStepVersion() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplateStepVersion), nil
}
func (m *TemplateStepVersionRequestBuilder) Ref()(*i4261fb17e0c76843a152f58b34dc951cd1cd3b513cd9ff4370402cbe2d7ccd76.RefRequestBuilder) {
    return i4261fb17e0c76843a152f58b34dc951cd1cd3b513cd9ff4370402cbe2d7ccd76.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
