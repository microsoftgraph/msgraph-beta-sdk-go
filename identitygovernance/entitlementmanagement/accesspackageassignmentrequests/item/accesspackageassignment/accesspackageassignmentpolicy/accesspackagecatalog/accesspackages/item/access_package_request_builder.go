package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i190983adfb7376d33a6a5418558040dfde1bfc2ab00485633baac461cfca29b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/getapplicablepolicyrequirements"
    i1fa3f3c34f9287edb19ce39aecc3886ea62768f119a6a6aa9b907091b58eaab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatiblegroups"
    i269149454375ea30f4cfb872b8c74baf0c9977eaafde667d894a6548079027de "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageassignmentpolicies"
    i2b01482df2c7282d2aa6540846ae693b4ddcdfbea9e718555093cee2ec6b124b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackagecatalog"
    ia0a01e5b0369d088dc591253d0c2b38e34624b4455953416298d30c29ecfdab4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackagesincompatiblewith"
    iba8d237af5e644881fe81fafba276a5e3d8dc4a777718fddd11cc99e8099c5f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatibleaccesspackages"
    ifafad20ba2dcedcbad9a25b787dcf7c7f6c52a47092a7a82259c5a3ce2800e9a "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageresourcerolescopes"
    i2c4b635be08a34312e524a2a207aacee40b7a5ffae005a8673dbf43e78760487 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageresourcerolescopes/item"
    i6be66ce2b4dc2717ab24dfb516dd00cc87dfa8426a8f47c9065de4c0a992fb34 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageassignmentpolicies/item"
    if540078400aa415c313352d519746c9a446dcd93d273f58299a332d43f6c7b2d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatiblegroups/item"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}\accessPackageAssignment\accessPackageAssignmentPolicy\accessPackageCatalog\accessPackages\{accessPackage-id}
type AccessPackageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AccessPackageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AccessPackageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The access packages in this catalog. Read-only. Nullable.
type AccessPackageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AccessPackageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPolicies()(*i269149454375ea30f4cfb872b8c74baf0c9977eaafde667d894a6548079027de.AccessPackageAssignmentPoliciesRequestBuilder) {
    return i269149454375ea30f4cfb872b8c74baf0c9977eaafde667d894a6548079027de.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item.accessPackageAssignment.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackages.item.accessPackageAssignmentPolicies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i6be66ce2b4dc2717ab24dfb516dd00cc87dfa8426a8f47c9065de4c0a992fb34.AccessPackageAssignmentPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy_id"] = id
    }
    return i6be66ce2b4dc2717ab24dfb516dd00cc87dfa8426a8f47c9065de4c0a992fb34.NewAccessPackageAssignmentPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageCatalog()(*i2b01482df2c7282d2aa6540846ae693b4ddcdfbea9e718555093cee2ec6b124b.AccessPackageCatalogRequestBuilder) {
    return i2b01482df2c7282d2aa6540846ae693b4ddcdfbea9e718555093cee2ec6b124b.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopes()(*ifafad20ba2dcedcbad9a25b787dcf7c7f6c52a47092a7a82259c5a3ce2800e9a.AccessPackageResourceRoleScopesRequestBuilder) {
    return ifafad20ba2dcedcbad9a25b787dcf7c7f6c52a47092a7a82259c5a3ce2800e9a.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item.accessPackageAssignment.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackages.item.accessPackageResourceRoleScopes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i2c4b635be08a34312e524a2a207aacee40b7a5ffae005a8673dbf43e78760487.AccessPackageResourceRoleScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope_id"] = id
    }
    return i2c4b635be08a34312e524a2a207aacee40b7a5ffae005a8673dbf43e78760487.NewAccessPackageResourceRoleScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackagesIncompatibleWith()(*ia0a01e5b0369d088dc591253d0c2b38e34624b4455953416298d30c29ecfdab4.AccessPackagesIncompatibleWithRequestBuilder) {
    return ia0a01e5b0369d088dc591253d0c2b38e34624b4455953416298d30c29ecfdab4.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new AccessPackageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    m := &AccessPackageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/{accessPackageAssignmentRequest_id}/accessPackageAssignment/accessPackageAssignmentPolicy/accessPackageCatalog/accessPackages/{accessPackage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AccessPackageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// The access packages in this catalog. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The access packages in this catalog. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageRequestBuilder) CreateGetRequestInformation(options *AccessPackageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The access packages in this catalog. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageRequestBuilder) CreatePatchRequestInformation(options *AccessPackageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The access packages in this catalog. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageRequestBuilder) Delete(options *AccessPackageRequestBuilderDeleteOptions)(error) {
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
// The access packages in this catalog. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageRequestBuilder) Get(options *AccessPackageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackage() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage), nil
}
func (m *AccessPackageRequestBuilder) GetApplicablePolicyRequirements()(*i190983adfb7376d33a6a5418558040dfde1bfc2ab00485633baac461cfca29b0.GetApplicablePolicyRequirementsRequestBuilder) {
    return i190983adfb7376d33a6a5418558040dfde1bfc2ab00485633baac461cfca29b0.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleAccessPackages()(*iba8d237af5e644881fe81fafba276a5e3d8dc4a777718fddd11cc99e8099c5f4.IncompatibleAccessPackagesRequestBuilder) {
    return iba8d237af5e644881fe81fafba276a5e3d8dc4a777718fddd11cc99e8099c5f4.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleGroups()(*i1fa3f3c34f9287edb19ce39aecc3886ea62768f119a6a6aa9b907091b58eaab5.IncompatibleGroupsRequestBuilder) {
    return i1fa3f3c34f9287edb19ce39aecc3886ea62768f119a6a6aa9b907091b58eaab5.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item.accessPackageAssignment.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackages.item.incompatibleGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageRequestBuilder) IncompatibleGroupsById(id string)(*if540078400aa415c313352d519746c9a446dcd93d273f58299a332d43f6c7b2d.GroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group_id"] = id
    }
    return if540078400aa415c313352d519746c9a446dcd93d273f58299a332d43f6c7b2d.NewGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The access packages in this catalog. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageRequestBuilder) Patch(options *AccessPackageRequestBuilderPatchOptions)(error) {
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
