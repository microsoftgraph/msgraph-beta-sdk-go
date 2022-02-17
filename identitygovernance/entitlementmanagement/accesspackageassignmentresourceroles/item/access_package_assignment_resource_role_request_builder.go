package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7d890f65851793ed799c10549d6fd6c6c2494d30953973e4207fcc7acf22a83b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments"
    iaa99fd9136f8f73ab1315e849913867001cc91a1d284f0e03cee25e796bbae38 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcescope"
    ic9d50657d48a7afc267e2cecedcbf53309e95d9f9bb8c499a5ff668c121cdc13 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcerole"
    ife3670cfcf012fea44d391fd12342b004d1d1ec7881d957f768ace00e28c468f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackagesubject"
)

// AccessPackageAssignmentResourceRoleRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}
type AccessPackageAssignmentResourceRoleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentResourceRoleRequestBuilderGetOptions options for Get
type AccessPackageAssignmentResourceRoleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageAssignmentResourceRoleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentResourceRoleRequestBuilderGetQueryParameters represents the resource-specific role which a subject has been assigned through an access package assignment.
type AccessPackageAssignmentResourceRoleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRole;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageAssignments()(*i7d890f65851793ed799c10549d6fd6c6c2494d30953973e4207fcc7acf22a83b.AccessPackageAssignmentsRequestBuilder) {
    return i7d890f65851793ed799c10549d6fd6c6c2494d30953973e4207fcc7acf22a83b.NewAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageResourceRole()(*ic9d50657d48a7afc267e2cecedcbf53309e95d9f9bb8c499a5ff668c121cdc13.AccessPackageResourceRoleRequestBuilder) {
    return ic9d50657d48a7afc267e2cecedcbf53309e95d9f9bb8c499a5ff668c121cdc13.NewAccessPackageResourceRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageResourceScope()(*iaa99fd9136f8f73ab1315e849913867001cc91a1d284f0e03cee25e796bbae38.AccessPackageResourceScopeRequestBuilder) {
    return iaa99fd9136f8f73ab1315e849913867001cc91a1d284f0e03cee25e796bbae38.NewAccessPackageResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageSubject()(*ife3670cfcf012fea44d391fd12342b004d1d1ec7881d957f768ace00e28c468f.AccessPackageSubjectRequestBuilder) {
    return ife3670cfcf012fea44d391fd12342b004d1d1ec7881d957f768ace00e28c468f.NewAccessPackageSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentResourceRoleRequestBuilderInternal instantiates a new AccessPackageAssignmentResourceRoleRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentResourceRoleRequestBuilder) {
    m := &AccessPackageAssignmentResourceRoleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentResourceRoleRequestBuilder instantiates a new AccessPackageAssignmentResourceRoleRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentResourceRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentResourceRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentResourceRoleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) Delete(options *AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions)(error) {
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
// Get represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) Get(options *AccessPackageAssignmentResourceRoleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRole, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageAssignmentResourceRole() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRole), nil
}
// Patch represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) Patch(options *AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions)(error) {
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
