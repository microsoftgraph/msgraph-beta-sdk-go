package managementtemplate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i81d954f6bbde9039cb01411f190fff066a48fef31d6fa87ea4ad54005c369eac "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteps/item/managementtemplate/ref"
)

// ManagementTemplateRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\managementTemplateSteps\{managementTemplateStep-id}\managementTemplate
type ManagementTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagementTemplateRequestBuilderGetOptions options for Get
type ManagementTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementTemplateRequestBuilderGetQueryParameters get managementTemplate from tenantRelationships
type ManagementTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewManagementTemplateRequestBuilderInternal instantiates a new ManagementTemplateRequestBuilder and sets the default values.
func NewManagementTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementTemplateRequestBuilder) {
    m := &ManagementTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateSteps/{managementTemplateStep_id}/managementTemplate{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementTemplateRequestBuilder instantiates a new ManagementTemplateRequestBuilder and sets the default values.
func NewManagementTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get managementTemplate from tenantRelationships
func (m *ManagementTemplateRequestBuilder) CreateGetRequestInformation(options *ManagementTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get managementTemplate from tenantRelationships
func (m *ManagementTemplateRequestBuilder) Get(options *ManagementTemplateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagementTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementTemplate), nil
}
func (m *ManagementTemplateRequestBuilder) Ref()(*i81d954f6bbde9039cb01411f190fff066a48fef31d6fa87ea4ad54005c369eac.RefRequestBuilder) {
    return i81d954f6bbde9039cb01411f190fff066a48fef31d6fa87ea4ad54005c369eac.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
