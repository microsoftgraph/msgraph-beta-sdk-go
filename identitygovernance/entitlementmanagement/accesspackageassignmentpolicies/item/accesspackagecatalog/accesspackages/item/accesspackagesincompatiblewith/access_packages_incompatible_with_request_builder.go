package accesspackagesincompatiblewith

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i04dcb68f459702c4277c836c5e28c543f056deb106fe0af31150a60003b27382 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackages/item/accesspackagesincompatiblewith/ref"
    i66c81427715cde95da33a3d507741a3e1b5ea5c0a92bbb958951b308d99dc3ba "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackages/item/accesspackagesincompatiblewith/search"
    i92347ef4a66bd2a1de15b377d6f9863ddb6270e05a020c1f482b92ad70bdd2d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackages/item/accesspackagesincompatiblewith/filterbycurrentuserwithon"
)

// accessPackagesIncompatibleWithRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentPolicies\{accessPackageAssignmentPolicy-id}\accessPackageCatalog\accessPackages\{accessPackage-id}\accessPackagesIncompatibleWith
type AccessPackagesIncompatibleWithRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackagesIncompatibleWithRequestBuilderGetOptions options for Get
type AccessPackagesIncompatibleWithRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackagesIncompatibleWithRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// accessPackagesIncompatibleWithRequestBuilderGetQueryParameters the access packages that are incompatible with this package. Read-only.
type AccessPackagesIncompatibleWithRequestBuilderGetQueryParameters struct {
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
// NewAccessPackagesIncompatibleWithRequestBuilderInternal instantiates a new AccessPackagesIncompatibleWithRequestBuilder and sets the default values.
func NewAccessPackagesIncompatibleWithRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackagesIncompatibleWithRequestBuilder) {
    m := &AccessPackagesIncompatibleWithRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy_id}/accessPackageCatalog/accessPackages/{accessPackage_id}/accessPackagesIncompatibleWith{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackagesIncompatibleWithRequestBuilder instantiates a new AccessPackagesIncompatibleWithRequestBuilder and sets the default values.
func NewAccessPackagesIncompatibleWithRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackagesIncompatibleWithRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackagesIncompatibleWithRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the access packages that are incompatible with this package. Read-only.
func (m *AccessPackagesIncompatibleWithRequestBuilder) CreateGetRequestInformation(options *AccessPackagesIncompatibleWithRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentPolicies\{accessPackageAssignmentPolicy-id}\accessPackageCatalog\accessPackages\{accessPackage-id}\accessPackagesIncompatibleWith\microsoft.graph.filterByCurrentUser(on={on})
func (m *AccessPackagesIncompatibleWithRequestBuilder) FilterByCurrentUserWithOn(on *string)(*i92347ef4a66bd2a1de15b377d6f9863ddb6270e05a020c1f482b92ad70bdd2d8.FilterByCurrentUserWithOnRequestBuilder) {
    return i92347ef4a66bd2a1de15b377d6f9863ddb6270e05a020c1f482b92ad70bdd2d8.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get the access packages that are incompatible with this package. Read-only.
func (m *AccessPackagesIncompatibleWithRequestBuilder) Get(options *AccessPackagesIncompatibleWithRequestBuilderGetOptions)(*AccessPackagesIncompatibleWithResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackagesIncompatibleWithResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AccessPackagesIncompatibleWithResponse), nil
}
func (m *AccessPackagesIncompatibleWithRequestBuilder) Ref()(*i04dcb68f459702c4277c836c5e28c543f056deb106fe0af31150a60003b27382.RefRequestBuilder) {
    return i04dcb68f459702c4277c836c5e28c543f056deb106fe0af31150a60003b27382.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Search builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentPolicies\{accessPackageAssignmentPolicy-id}\accessPackageCatalog\accessPackages\{accessPackage-id}\accessPackagesIncompatibleWith\microsoft.graph.Search()
func (m *AccessPackagesIncompatibleWithRequestBuilder) Search()(*i66c81427715cde95da33a3d507741a3e1b5ea5c0a92bbb958951b308d99dc3ba.SearchRequestBuilder) {
    return i66c81427715cde95da33a3d507741a3e1b5ea5c0a92bbb958951b308d99dc3ba.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
