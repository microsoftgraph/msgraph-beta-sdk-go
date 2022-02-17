package accesspackagesincompatiblewith

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i45ebc34581437a79433cf4aa3876992373c137af7a7637663271606abbcf3a5e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackagesincompatiblewith/search"
    i92feb2d0991460a5350aa50e3dd293b750ef10e0ebb969a34f3bc81581c92021 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackagesincompatiblewith/filterbycurrentuserwithon"
    ib61b8206722061a54c5b0d950b74dbf691189ad6adca7951e55f274324f6ed77 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackagesincompatiblewith/ref"
)

// AccessPackagesIncompatibleWithRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\accessPackagesIncompatibleWith
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
// AccessPackagesIncompatibleWithRequestBuilderGetQueryParameters the access packages that are incompatible with this package. Read-only.
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
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewAccessPackagesIncompatibleWithRequestBuilderInternal instantiates a new AccessPackagesIncompatibleWithRequestBuilder and sets the default values.
func NewAccessPackagesIncompatibleWithRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackagesIncompatibleWithRequestBuilder) {
    m := &AccessPackagesIncompatibleWithRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage/accessPackagesIncompatibleWith{?top,skip,search,filter,count,orderby,select,expand}";
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\accessPackagesIncompatibleWith\microsoft.graph.filterByCurrentUser(on={on})
func (m *AccessPackagesIncompatibleWithRequestBuilder) FilterByCurrentUserWithOn(on *string)(*i92feb2d0991460a5350aa50e3dd293b750ef10e0ebb969a34f3bc81581c92021.FilterByCurrentUserWithOnRequestBuilder) {
    return i92feb2d0991460a5350aa50e3dd293b750ef10e0ebb969a34f3bc81581c92021.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get the access packages that are incompatible with this package. Read-only.
func (m *AccessPackagesIncompatibleWithRequestBuilder) Get(options *AccessPackagesIncompatibleWithRequestBuilderGetOptions)(*AccessPackagesIncompatibleWithResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackagesIncompatibleWithResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AccessPackagesIncompatibleWithResponse), nil
}
func (m *AccessPackagesIncompatibleWithRequestBuilder) Ref()(*ib61b8206722061a54c5b0d950b74dbf691189ad6adca7951e55f274324f6ed77.RefRequestBuilder) {
    return ib61b8206722061a54c5b0d950b74dbf691189ad6adca7951e55f274324f6ed77.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Search builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\accessPackagesIncompatibleWith\microsoft.graph.Search()
func (m *AccessPackagesIncompatibleWithRequestBuilder) Search()(*i45ebc34581437a79433cf4aa3876992373c137af7a7637663271606abbcf3a5e.SearchRequestBuilder) {
    return i45ebc34581437a79433cf4aa3876992373c137af7a7637663271606abbcf3a5e.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
