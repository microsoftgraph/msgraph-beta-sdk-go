package tenantgroups

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
    i86e8e63e85de67171644c463e36cc57d3c920f484f6937360f52c82139946fb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups/tenantsearch"
    ia08441d1849e327095c234a187264cbf5a076c3016bbc7c3a3f4b741b9a7f50a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups/count"
)

// TenantGroupsRequestBuilder provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TenantGroupsRequestBuilderGetOptions options for Get
type TenantGroupsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TenantGroupsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TenantGroupsRequestBuilderGetQueryParameters the collection of a logical grouping of managed tenants used by the multi-tenant management platform.
type TenantGroupsRequestBuilderGetQueryParameters struct {
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
// TenantGroupsRequestBuilderPostOptions options for Post
type TenantGroupsRequestBuilderPostOptions struct {
    // 
    Body i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantGroupable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTenantGroupsRequestBuilderInternal instantiates a new TenantGroupsRequestBuilder and sets the default values.
func NewTenantGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TenantGroupsRequestBuilder) {
    m := &TenantGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenantGroups{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantGroupsRequestBuilder instantiates a new TenantGroupsRequestBuilder and sets the default values.
func NewTenantGroupsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TenantGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TenantGroupsRequestBuilder) Count()(*ia08441d1849e327095c234a187264cbf5a076c3016bbc7c3a3f4b741b9a7f50a.CountRequestBuilder) {
    return ia08441d1849e327095c234a187264cbf5a076c3016bbc7c3a3f4b741b9a7f50a.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *TenantGroupsRequestBuilder) CreateGetRequestInformation(options *TenantGroupsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to tenantGroups for tenantRelationships
func (m *TenantGroupsRequestBuilder) CreatePostRequestInformation(options *TenantGroupsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *TenantGroupsRequestBuilder) Get(options *TenantGroupsRequestBuilderGetOptions)(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantGroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateTenantGroupCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantGroupCollectionResponseable), nil
}
// Post create new navigation property to tenantGroups for tenantRelationships
func (m *TenantGroupsRequestBuilder) Post(options *TenantGroupsRequestBuilderPostOptions)(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantGroupable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateTenantGroupFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantGroupable), nil
}
func (m *TenantGroupsRequestBuilder) TenantSearch()(*i86e8e63e85de67171644c463e36cc57d3c920f484f6937360f52c82139946fb4.TenantSearchRequestBuilder) {
    return i86e8e63e85de67171644c463e36cc57d3c920f484f6937360f52c82139946fb4.NewTenantSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
