package managementactiontenantdeploymentstatuses

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id50624c28ac1a98a4cd1cc6f61b4dc20c4a7524a54c2bbd7c8cbd10f86ded55f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses/changedeploymentstatus"
)

// ManagementActionTenantDeploymentStatusesRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\managementActionTenantDeploymentStatuses
type ManagementActionTenantDeploymentStatusesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions options for Get
type ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters the tenant level status of management actions across managed tenants.
type ManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters struct {
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
// ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions options for Post
type ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementActionTenantDeploymentStatus;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) ChangeDeploymentStatus()(*id50624c28ac1a98a4cd1cc6f61b4dc20c4a7524a54c2bbd7c8cbd10f86ded55f.ChangeDeploymentStatusRequestBuilder) {
    return id50624c28ac1a98a4cd1cc6f61b4dc20c4a7524a54c2bbd7c8cbd10f86ded55f.NewChangeDeploymentStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagementActionTenantDeploymentStatusesRequestBuilderInternal instantiates a new ManagementActionTenantDeploymentStatusesRequestBuilder and sets the default values.
func NewManagementActionTenantDeploymentStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementActionTenantDeploymentStatusesRequestBuilder) {
    m := &ManagementActionTenantDeploymentStatusesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementActionTenantDeploymentStatuses{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementActionTenantDeploymentStatusesRequestBuilder instantiates a new ManagementActionTenantDeploymentStatusesRequestBuilder and sets the default values.
func NewManagementActionTenantDeploymentStatusesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementActionTenantDeploymentStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementActionTenantDeploymentStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the tenant level status of management actions across managed tenants.
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) CreateGetRequestInformation(options *ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation the tenant level status of management actions across managed tenants.
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) CreatePostRequestInformation(options *ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the tenant level status of management actions across managed tenants.
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) Get(options *ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions)(*ManagementActionTenantDeploymentStatusesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementActionTenantDeploymentStatusesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagementActionTenantDeploymentStatusesResponse), nil
}
// Post the tenant level status of management actions across managed tenants.
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) Post(options *ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementActionTenantDeploymentStatus, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagementActionTenantDeploymentStatus() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementActionTenantDeploymentStatus), nil
}
