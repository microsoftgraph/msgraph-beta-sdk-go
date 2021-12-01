package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i19bb6b2247a0ce2eb5ddf900b69ec6a56655ebbd9dade3103ea4619c4dd6f5dc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/reprocess"
    i3c2c3e953d1ab01b11fd6abf2b2655653badcee87a7615e205d4ebb812bdfacb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy"
    i3d5a23907f9b30748283783ab5dea1f4376f78845d0e79f6c02d90dc7a01afb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentresourceroles"
    i660156b0613442327faf16ac6012969a764222e7492b930b3263eeae3b36193d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/target"
    idc493656467acca9133b4eddf84e365cb5f571cc845ea45f382d2545ece9e5cc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentrequests"
    iec81b0373c4da350b26d8984817114f28880005dadeeefba0eaccb6aa243f8ea "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage"
    i27994df1d0f4c91ae8ca583362cd3504a5e0201c5d06854f136e8de282fbae2a "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentresourceroles/item"
    i48399c4b2202ca662aaa0213380e828fbb875b4ba9647a0e692b5ac58ed4deed "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentrequests/item"
)

// AccessPackageAssignmentRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}
type AccessPackageAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentRequestBuilderGetOptions options for Get
type AccessPackageAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentRequestBuilderGetQueryParameters the access package assignments resulting in this role assignment. Read-only. Nullable.
type AccessPackageAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackage()(*iec81b0373c4da350b26d8984817114f28880005dadeeefba0eaccb6aa243f8ea.AccessPackageRequestBuilder) {
    return iec81b0373c4da350b26d8984817114f28880005dadeeefba0eaccb6aa243f8ea.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentPolicy()(*i3c2c3e953d1ab01b11fd6abf2b2655653badcee87a7615e205d4ebb812bdfacb.AccessPackageAssignmentPolicyRequestBuilder) {
    return i3c2c3e953d1ab01b11fd6abf2b2655653badcee87a7615e205d4ebb812bdfacb.NewAccessPackageAssignmentPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentRequests()(*idc493656467acca9133b4eddf84e365cb5f571cc845ea45f382d2545ece9e5cc.AccessPackageAssignmentRequestsRequestBuilder) {
    return idc493656467acca9133b4eddf84e365cb5f571cc845ea45f382d2545ece9e5cc.NewAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentRequests.item collection
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*i48399c4b2202ca662aaa0213380e828fbb875b4ba9647a0e692b5ac58ed4deed.AccessPackageAssignmentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest_id"] = id
    }
    return i48399c4b2202ca662aaa0213380e828fbb875b4ba9647a0e692b5ac58ed4deed.NewAccessPackageAssignmentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentResourceRoles()(*i3d5a23907f9b30748283783ab5dea1f4376f78845d0e79f6c02d90dc7a01afb1.AccessPackageAssignmentResourceRolesRequestBuilder) {
    return i3d5a23907f9b30748283783ab5dea1f4376f78845d0e79f6c02d90dc7a01afb1.NewAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentResourceRoles.item collection
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*i27994df1d0f4c91ae8ca583362cd3504a5e0201c5d06854f136e8de282fbae2a.AccessPackageAssignmentResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole_id1"] = id
    }
    return i27994df1d0f4c91ae8ca583362cd3504a5e0201c5d06854f136e8de282fbae2a.NewAccessPackageAssignmentResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageAssignmentRequestBuilderInternal instantiates a new AccessPackageAssignmentRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentRequestBuilder) {
    m := &AccessPackageAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}/accessPackageAssignments/{accessPackageAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentRequestBuilder instantiates a new AccessPackageAssignmentRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentRequestBuilder) Delete(options *AccessPackageAssignmentRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentRequestBuilder) Get(options *AccessPackageAssignmentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageAssignment() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment), nil
}
// Patch the access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentRequestBuilder) Patch(options *AccessPackageAssignmentRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AccessPackageAssignmentRequestBuilder) Reprocess()(*i19bb6b2247a0ce2eb5ddf900b69ec6a56655ebbd9dade3103ea4619c4dd6f5dc.ReprocessRequestBuilder) {
    return i19bb6b2247a0ce2eb5ddf900b69ec6a56655ebbd9dade3103ea4619c4dd6f5dc.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) Target()(*i660156b0613442327faf16ac6012969a764222e7492b930b3263eeae3b36193d.TargetRequestBuilder) {
    return i660156b0613442327faf16ac6012969a764222e7492b930b3263eeae3b36193d.NewTargetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
