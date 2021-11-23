package accesspackagesincompatiblewith

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0d58ccbeb54572e060c60cf5cd44e654a2bdfffc31e49e9bc54b8de9f326d829 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackage/accesspackagesincompatiblewith/filterbycurrentuserwithon"
    ic3bcc5e7ed3e4d985eee773d4b921f72ac36167e48d94d47ba0b7435104269a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackage/accesspackagesincompatiblewith/search"
    ie25857b1c79307bfbf5c4b5950130d8c2e3de3e015661d9764fb5e80c30ab216 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackage/accesspackagesincompatiblewith/ref"
)

// accessPackagesIncompatibleWithRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}\accessPackageAssignment\accessPackage\accessPackagesIncompatibleWith
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
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/{accessPackageAssignmentRequest_id}/accessPackageAssignment/accessPackage/accessPackagesIncompatibleWith{?top,skip,search,filter,count,orderby,select,expand}";
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}\accessPackageAssignment\accessPackage\accessPackagesIncompatibleWith\microsoft.graph.filterByCurrentUser(on={on})
func (m *AccessPackagesIncompatibleWithRequestBuilder) FilterByCurrentUserWithOn(on *string)(*i0d58ccbeb54572e060c60cf5cd44e654a2bdfffc31e49e9bc54b8de9f326d829.FilterByCurrentUserWithOnRequestBuilder) {
    return i0d58ccbeb54572e060c60cf5cd44e654a2bdfffc31e49e9bc54b8de9f326d829.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
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
func (m *AccessPackagesIncompatibleWithRequestBuilder) Ref()(*ie25857b1c79307bfbf5c4b5950130d8c2e3de3e015661d9764fb5e80c30ab216.RefRequestBuilder) {
    return ie25857b1c79307bfbf5c4b5950130d8c2e3de3e015661d9764fb5e80c30ab216.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Search builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}\accessPackageAssignment\accessPackage\accessPackagesIncompatibleWith\microsoft.graph.Search()
func (m *AccessPackagesIncompatibleWithRequestBuilder) Search()(*ic3bcc5e7ed3e4d985eee773d4b921f72ac36167e48d94d47ba0b7435104269a0.SearchRequestBuilder) {
    return ic3bcc5e7ed3e4d985eee773d4b921f72ac36167e48d94d47ba0b7435104269a0.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
