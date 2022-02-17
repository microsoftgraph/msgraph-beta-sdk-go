package managementtemplates

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3cdb6022c8caa8f7c6f791f0249e3ee549d0d83c9bd530870a7cbb0c2f85fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatecollections/item/managementtemplates/ref"
)

// ManagementTemplatesRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\managementTemplateCollections\{managementTemplateCollection-id}\managementTemplates
type ManagementTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagementTemplatesRequestBuilderGetOptions options for Get
type ManagementTemplatesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementTemplatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementTemplatesRequestBuilderGetQueryParameters get managementTemplates from tenantRelationships
type ManagementTemplatesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewManagementTemplatesRequestBuilderInternal instantiates a new ManagementTemplatesRequestBuilder and sets the default values.
func NewManagementTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementTemplatesRequestBuilder) {
    m := &ManagementTemplatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateCollections/{managementTemplateCollection_id}/managementTemplates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementTemplatesRequestBuilder instantiates a new ManagementTemplatesRequestBuilder and sets the default values.
func NewManagementTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get managementTemplates from tenantRelationships
func (m *ManagementTemplatesRequestBuilder) CreateGetRequestInformation(options *ManagementTemplatesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get managementTemplates from tenantRelationships
func (m *ManagementTemplatesRequestBuilder) Get(options *ManagementTemplatesRequestBuilderGetOptions)(*ManagementTemplatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplatesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagementTemplatesResponse), nil
}
func (m *ManagementTemplatesRequestBuilder) Ref()(*i3cdb6022c8caa8f7c6f791f0249e3ee549d0d83c9bd530870a7cbb0c2f85fe01.RefRequestBuilder) {
    return i3cdb6022c8caa8f7c6f791f0249e3ee549d0d83c9bd530870a7cbb0c2f85fe01.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
