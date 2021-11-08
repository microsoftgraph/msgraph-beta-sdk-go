package incompatibleaccesspackages

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i28c20aacdc2a68a9ad67d99715fdf733971ee129989f81f4e7c808a3810aaed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatibleaccesspackages/search"
    ib8517794777a6fc9de103765ea80795cd790ceeaa1f110bdb38b1d6f5d4fbd0f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatibleaccesspackages/ref"
    ic7d63a266018683f0c48b1578c1867f0a92d568a82ef03091addab18c0da0785 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatibleaccesspackages/filterbycurrentuserwithon"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackageAssignmentPolicy\accessPackageCatalog\accessPackages\{accessPackage-id}\incompatibleAccessPackages
type IncompatibleAccessPackagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
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
// The  access packages whose assigned users are ineligible to be assigned this access package.
type IncompatibleAccessPackagesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewIncompatibleAccessPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncompatibleAccessPackagesRequestBuilder) {
    m := &IncompatibleAccessPackagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackageAssignmentPolicy/accessPackageCatalog/accessPackages/{accessPackage_id}/incompatibleAccessPackages{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewIncompatibleAccessPackagesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncompatibleAccessPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIncompatibleAccessPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// The  access packages whose assigned users are ineligible to be assigned this access package.
// Parameters:
//  - options : Options for the request
func (m *IncompatibleAccessPackagesRequestBuilder) CreateGetRequestInformation(options *IncompatibleAccessPackagesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackageAssignmentPolicy\accessPackageCatalog\accessPackages\{accessPackage-id}\incompatibleAccessPackages\microsoft.graph.filterByCurrentUser(on={on})
// Parameters:
//  - on : Usage: on={on}
func (m *IncompatibleAccessPackagesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*ic7d63a266018683f0c48b1578c1867f0a92d568a82ef03091addab18c0da0785.FilterByCurrentUserWithOnRequestBuilder) {
    return ic7d63a266018683f0c48b1578c1867f0a92d568a82ef03091addab18c0da0785.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// The  access packages whose assigned users are ineligible to be assigned this access package.
// Parameters:
//  - options : Options for the request
func (m *IncompatibleAccessPackagesRequestBuilder) Get(options *IncompatibleAccessPackagesRequestBuilderGetOptions)(*IncompatibleAccessPackagesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncompatibleAccessPackagesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*IncompatibleAccessPackagesResponse), nil
}
func (m *IncompatibleAccessPackagesRequestBuilder) Ref()(*ib8517794777a6fc9de103765ea80795cd790ceeaa1f110bdb38b1d6f5d4fbd0f.RefRequestBuilder) {
    return ib8517794777a6fc9de103765ea80795cd790ceeaa1f110bdb38b1d6f5d4fbd0f.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackageAssignmentPolicy\accessPackageCatalog\accessPackages\{accessPackage-id}\incompatibleAccessPackages\microsoft.graph.Search()
func (m *IncompatibleAccessPackagesRequestBuilder) Search()(*i28c20aacdc2a68a9ad67d99715fdf733971ee129989f81f4e7c808a3810aaed2.SearchRequestBuilder) {
    return i28c20aacdc2a68a9ad67d99715fdf733971ee129989f81f4e7c808a3810aaed2.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
