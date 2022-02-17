package incompatibleaccesspackages

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i23dc487e3641978a29545d0342f937e1cbc55e1587cdf8fb15c6d30280ecfc44 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages/filterbycurrentuserwithon"
    i32d4eb085a1ed7f6c3cdae2d6e021dbec724d19797fd5d2f47e32eac72e92f34 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages/search"
    i606e76decb617ef392f0a2ced4701e81499efef20fabff5ac750a6e0dabbe071 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages/ref"
)

// IncompatibleAccessPackagesRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}\incompatibleAccessPackages
type IncompatibleAccessPackagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IncompatibleAccessPackagesRequestBuilderGetOptions options for Get
type IncompatibleAccessPackagesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IncompatibleAccessPackagesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IncompatibleAccessPackagesRequestBuilderGetQueryParameters the  access packages whose assigned users are ineligible to be assigned this access package.
type IncompatibleAccessPackagesRequestBuilderGetQueryParameters struct {
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
// NewIncompatibleAccessPackagesRequestBuilderInternal instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewIncompatibleAccessPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncompatibleAccessPackagesRequestBuilder) {
    m := &IncompatibleAccessPackagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage_id}/incompatibleAccessPackages{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIncompatibleAccessPackagesRequestBuilder instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewIncompatibleAccessPackagesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncompatibleAccessPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIncompatibleAccessPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the  access packages whose assigned users are ineligible to be assigned this access package.
func (m *IncompatibleAccessPackagesRequestBuilder) CreateGetRequestInformation(options *IncompatibleAccessPackagesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}\incompatibleAccessPackages\microsoft.graph.filterByCurrentUser(on={on})
func (m *IncompatibleAccessPackagesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*i23dc487e3641978a29545d0342f937e1cbc55e1587cdf8fb15c6d30280ecfc44.FilterByCurrentUserWithOnRequestBuilder) {
    return i23dc487e3641978a29545d0342f937e1cbc55e1587cdf8fb15c6d30280ecfc44.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get the  access packages whose assigned users are ineligible to be assigned this access package.
func (m *IncompatibleAccessPackagesRequestBuilder) Get(options *IncompatibleAccessPackagesRequestBuilderGetOptions)(*IncompatibleAccessPackagesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncompatibleAccessPackagesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*IncompatibleAccessPackagesResponse), nil
}
func (m *IncompatibleAccessPackagesRequestBuilder) Ref()(*i606e76decb617ef392f0a2ced4701e81499efef20fabff5ac750a6e0dabbe071.RefRequestBuilder) {
    return i606e76decb617ef392f0a2ced4701e81499efef20fabff5ac750a6e0dabbe071.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Search builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}\incompatibleAccessPackages\microsoft.graph.Search()
func (m *IncompatibleAccessPackagesRequestBuilder) Search()(*i32d4eb085a1ed7f6c3cdae2d6e021dbec724d19797fd5d2f47e32eac72e92f34.SearchRequestBuilder) {
    return i32d4eb085a1ed7f6c3cdae2d6e021dbec724d19797fd5d2f47e32eac72e92f34.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
