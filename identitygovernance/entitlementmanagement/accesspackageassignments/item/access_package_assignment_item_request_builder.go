package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2af0bfbcee7907a04a6c5b1b4818035a6651210ffff9b70fda593206c9d69a95 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/target"
    i3681a850a3cd793fe3c063807af12bdde2ee5a63a2cd6b2033458b57d5e3fa51 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy"
    i6b87ae9babacd24aa7786e6d45bebb8e19e6210a5805c3a06f65b372589cd0a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles"
    ia0acba2babd5b1475401d0b6eaac48aff95dc39f020731cafbfd07a3cd647d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage"
    iba6b2e1e9a7782cc4bbca9924b10deee1e26627f7845e0b11bbb91671f323380 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/reprocess"
    id7bb72014a00aae74285bd91b1a862ae93a5130eaa1f7e48ba20472c19da7c73 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests"
    i140a5f9ba5238539c561b0796e993a7f43398107e3e3cc9f8bf919276d650b44 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item"
    i5252ed1a2033545e7629ee16423aea516ec3f2de5b6f26fa3842c58927181b88 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests/item"
)

// AccessPackageAssignmentItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}
type AccessPackageAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentItemRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentItemRequestBuilderGetOptions options for Get
type AccessPackageAssignmentItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageAssignmentItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentItemRequestBuilderGetQueryParameters represents the grant of an access package to a subject (user or group).
type AccessPackageAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentItemRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackage()(*ia0acba2babd5b1475401d0b6eaac48aff95dc39f020731cafbfd07a3cd647d18.AccessPackageRequestBuilder) {
    return ia0acba2babd5b1475401d0b6eaac48aff95dc39f020731cafbfd07a3cd647d18.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentPolicy()(*i3681a850a3cd793fe3c063807af12bdde2ee5a63a2cd6b2033458b57d5e3fa51.AccessPackageAssignmentPolicyRequestBuilder) {
    return i3681a850a3cd793fe3c063807af12bdde2ee5a63a2cd6b2033458b57d5e3fa51.NewAccessPackageAssignmentPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentRequests()(*id7bb72014a00aae74285bd91b1a862ae93a5130eaa1f7e48ba20472c19da7c73.AccessPackageAssignmentRequestsRequestBuilder) {
    return id7bb72014a00aae74285bd91b1a862ae93a5130eaa1f7e48ba20472c19da7c73.NewAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackageAssignmentRequests.item collection
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*i5252ed1a2033545e7629ee16423aea516ec3f2de5b6f26fa3842c58927181b88.AccessPackageAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest_id"] = id
    }
    return i5252ed1a2033545e7629ee16423aea516ec3f2de5b6f26fa3842c58927181b88.NewAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentResourceRoles()(*i6b87ae9babacd24aa7786e6d45bebb8e19e6210a5805c3a06f65b372589cd0a8.AccessPackageAssignmentResourceRolesRequestBuilder) {
    return i6b87ae9babacd24aa7786e6d45bebb8e19e6210a5805c3a06f65b372589cd0a8.NewAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackageAssignmentResourceRoles.item collection
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*i140a5f9ba5238539c561b0796e993a7f43398107e3e3cc9f8bf919276d650b44.AccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole_id"] = id
    }
    return i140a5f9ba5238539c561b0796e993a7f43398107e3e3cc9f8bf919276d650b44.NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageAssignmentItemRequestBuilderInternal instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentItemRequestBuilder) {
    m := &AccessPackageAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentItemRequestBuilder instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) Delete(options *AccessPackageAssignmentItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) Get(options *AccessPackageAssignmentItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment), nil
}
// Patch represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) Patch(options *AccessPackageAssignmentItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AccessPackageAssignmentItemRequestBuilder) Reprocess()(*iba6b2e1e9a7782cc4bbca9924b10deee1e26627f7845e0b11bbb91671f323380.ReprocessRequestBuilder) {
    return iba6b2e1e9a7782cc4bbca9924b10deee1e26627f7845e0b11bbb91671f323380.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentItemRequestBuilder) Target()(*i2af0bfbcee7907a04a6c5b1b4818035a6651210ffff9b70fda593206c9d69a95.TargetRequestBuilder) {
    return i2af0bfbcee7907a04a6c5b1b4818035a6651210ffff9b70fda593206c9d69a95.NewTargetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
