package tenantscustomizedinformation

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
    iaf3e02c828f39af7b769e35e7aa9da0803b0c8afb407564e2c5fb7a5dbbb26ca "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantscustomizedinformation/count"
)

// TenantsCustomizedInformationRequestBuilder provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantsCustomizedInformationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TenantsCustomizedInformationRequestBuilderGetOptions options for Get
type TenantsCustomizedInformationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TenantsCustomizedInformationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TenantsCustomizedInformationRequestBuilderGetQueryParameters the collection of tenant level customized information across managed tenants.
type TenantsCustomizedInformationRequestBuilderGetQueryParameters struct {
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
// TenantsCustomizedInformationRequestBuilderPostOptions options for Post
type TenantsCustomizedInformationRequestBuilderPostOptions struct {
    // 
    Body i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantCustomizedInformationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTenantsCustomizedInformationRequestBuilderInternal instantiates a new TenantsCustomizedInformationRequestBuilder and sets the default values.
func NewTenantsCustomizedInformationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TenantsCustomizedInformationRequestBuilder) {
    m := &TenantsCustomizedInformationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenantsCustomizedInformation{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantsCustomizedInformationRequestBuilder instantiates a new TenantsCustomizedInformationRequestBuilder and sets the default values.
func NewTenantsCustomizedInformationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TenantsCustomizedInformationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantsCustomizedInformationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TenantsCustomizedInformationRequestBuilder) Count()(*iaf3e02c828f39af7b769e35e7aa9da0803b0c8afb407564e2c5fb7a5dbbb26ca.CountRequestBuilder) {
    return iaf3e02c828f39af7b769e35e7aa9da0803b0c8afb407564e2c5fb7a5dbbb26ca.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of tenant level customized information across managed tenants.
func (m *TenantsCustomizedInformationRequestBuilder) CreateGetRequestInformation(options *TenantsCustomizedInformationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to tenantsCustomizedInformation for tenantRelationships
func (m *TenantsCustomizedInformationRequestBuilder) CreatePostRequestInformation(options *TenantsCustomizedInformationRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the collection of tenant level customized information across managed tenants.
func (m *TenantsCustomizedInformationRequestBuilder) Get(options *TenantsCustomizedInformationRequestBuilderGetOptions)(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantCustomizedInformationCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateTenantCustomizedInformationCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantCustomizedInformationCollectionResponseable), nil
}
// Post create new navigation property to tenantsCustomizedInformation for tenantRelationships
func (m *TenantsCustomizedInformationRequestBuilder) Post(options *TenantsCustomizedInformationRequestBuilderPostOptions)(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantCustomizedInformationable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateTenantCustomizedInformationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantCustomizedInformationable), nil
}
