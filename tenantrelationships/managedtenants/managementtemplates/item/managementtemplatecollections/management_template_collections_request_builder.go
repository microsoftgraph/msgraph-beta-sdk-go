package managementtemplatecollections

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    if2cb735b8e1f5a4dc100e2d27310069fcdb495776e2b40116f0238c1a3805e9c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplates/item/managementtemplatecollections/ref"
)

// ManagementTemplateCollectionsRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\managementTemplates\{managementTemplate-id}\managementTemplateCollections
type ManagementTemplateCollectionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagementTemplateCollectionsRequestBuilderGetOptions options for Get
type ManagementTemplateCollectionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementTemplateCollectionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementTemplateCollectionsRequestBuilderGetQueryParameters get managementTemplateCollections from tenantRelationships
type ManagementTemplateCollectionsRequestBuilderGetQueryParameters struct {
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
// NewManagementTemplateCollectionsRequestBuilderInternal instantiates a new ManagementTemplateCollectionsRequestBuilder and sets the default values.
func NewManagementTemplateCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementTemplateCollectionsRequestBuilder) {
    m := &ManagementTemplateCollectionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplates/{managementTemplate_id}/managementTemplateCollections{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementTemplateCollectionsRequestBuilder instantiates a new ManagementTemplateCollectionsRequestBuilder and sets the default values.
func NewManagementTemplateCollectionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementTemplateCollectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementTemplateCollectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get managementTemplateCollections from tenantRelationships
func (m *ManagementTemplateCollectionsRequestBuilder) CreateGetRequestInformation(options *ManagementTemplateCollectionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get managementTemplateCollections from tenantRelationships
func (m *ManagementTemplateCollectionsRequestBuilder) Get(options *ManagementTemplateCollectionsRequestBuilderGetOptions)(*ManagementTemplateCollectionsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateCollectionsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagementTemplateCollectionsResponse), nil
}
func (m *ManagementTemplateCollectionsRequestBuilder) Ref()(*if2cb735b8e1f5a4dc100e2d27310069fcdb495776e2b40116f0238c1a3805e9c.RefRequestBuilder) {
    return if2cb735b8e1f5a4dc100e2d27310069fcdb495776e2b40116f0238c1a3805e9c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
