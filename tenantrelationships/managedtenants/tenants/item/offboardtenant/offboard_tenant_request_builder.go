package offboardtenant

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// OffboardTenantRequestBuilder provides operations to call the offboardTenant method.
type OffboardTenantRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OffboardTenantRequestBuilderPostOptions options for Post
type OffboardTenantRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOffboardTenantRequestBuilderInternal instantiates a new OffboardTenantRequestBuilder and sets the default values.
func NewOffboardTenantRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OffboardTenantRequestBuilder) {
    m := &OffboardTenantRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenants/{tenant_id}/microsoft.graph.managedTenants.offboardTenant";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOffboardTenantRequestBuilder instantiates a new OffboardTenantRequestBuilder and sets the default values.
func NewOffboardTenantRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OffboardTenantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOffboardTenantRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action offboardTenant
func (m *OffboardTenantRequestBuilder) CreatePostRequestInformation(options *OffboardTenantRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action offboardTenant
func (m *OffboardTenantRequestBuilder) Post(options *OffboardTenantRequestBuilderPostOptions)(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateTenantFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable), nil
}
