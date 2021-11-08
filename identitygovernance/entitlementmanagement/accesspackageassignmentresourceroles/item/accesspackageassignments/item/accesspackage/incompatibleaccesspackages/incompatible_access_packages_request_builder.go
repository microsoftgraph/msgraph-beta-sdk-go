package incompatibleaccesspackages

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i05723bdb202f576597238d803ea8e7b07e98df5cc7ea4400c6774866dbe96337 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/incompatibleaccesspackages/search"
    i66b5aa2cc0599554b04d530f7a9e7f5c6d9c87f817ec406351048191b5d91de7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/incompatibleaccesspackages/filterbycurrentuserwithon"
    i6f6400ba72ba46ab3ccf588707fb62ff944d4bdd7e1db51b36f7fe34d77d6ff7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/incompatibleaccesspackages/ref"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\incompatibleAccessPackages
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
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage/incompatibleAccessPackages{?top,skip,search,filter,count,orderby,select,expand}";
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
// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\incompatibleAccessPackages\microsoft.graph.filterByCurrentUser(on={on})
// Parameters:
//  - on : Usage: on={on}
func (m *IncompatibleAccessPackagesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*i66b5aa2cc0599554b04d530f7a9e7f5c6d9c87f817ec406351048191b5d91de7.FilterByCurrentUserWithOnRequestBuilder) {
    return i66b5aa2cc0599554b04d530f7a9e7f5c6d9c87f817ec406351048191b5d91de7.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
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
func (m *IncompatibleAccessPackagesRequestBuilder) Ref()(*i6f6400ba72ba46ab3ccf588707fb62ff944d4bdd7e1db51b36f7fe34d77d6ff7.RefRequestBuilder) {
    return i6f6400ba72ba46ab3ccf588707fb62ff944d4bdd7e1db51b36f7fe34d77d6ff7.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\incompatibleAccessPackages\microsoft.graph.Search()
func (m *IncompatibleAccessPackagesRequestBuilder) Search()(*i05723bdb202f576597238d803ea8e7b07e98df5cc7ea4400c6774866dbe96337.SearchRequestBuilder) {
    return i05723bdb202f576597238d803ea8e7b07e98df5cc7ea4400c6774866dbe96337.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
