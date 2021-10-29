package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b548a88d8027e48ea15291f1dfe39a8927fe93dcca07173b0472f3798f2d017 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackagesincompatiblewith"
    i8a8700fea501b7bf142c3430294dd2e707d3d2ce5820b7d1bef4ed9af4319f80 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageresourcerolescopes"
    i9833a1555432d936ca4a849db79f7f2f28bce6ec20b1f571bf7dd15e45592fd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatiblegroups"
    ib224a2189234906b261ee0e2da4596d3a9d8047d5957ca375469a34ba88dde3c "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/getapplicablepolicyrequirements"
    ic0419a87849e51ae8a63d5884bab9c32f3ddbf25318b1c51c2dd3e111752dcfd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackagecatalog"
    if2e34baa9d5f486470dc4a9e3087ec6cb55685341f5b383564a57c1cdd6a62ea "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageassignmentpolicies"
    if83c62939ddf8992405494eb0d52216cfd46c87f9a8c60c3763707c24531971f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatibleaccesspackages"
    i9616a87af7f3c7810303cf2a99a508cb66680536b1ac2e435928362fcc6514fc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageassignmentpolicies/item"
    ic15fa732fd94aca01d0c675909da26f44abd7b76c06c2524d9f8eb53be1aea72 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/incompatiblegroups/item"
    if10b7ce5e756ef96bfc4cf4bb17c7bd568a9095cd3f8b40831944167f04716b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item/accesspackageresourcerolescopes/item"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}\accessPackageAssignmentPolicy\accessPackageCatalog\accessPackages\{accessPackage-id}
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
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPolicies()(*if2e34baa9d5f486470dc4a9e3087ec6cb55685341f5b383564a57c1cdd6a62ea.AccessPackageAssignmentPoliciesRequestBuilder) {
    return if2e34baa9d5f486470dc4a9e3087ec6cb55685341f5b383564a57c1cdd6a62ea.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackages.item.accessPackageAssignmentPolicies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i9616a87af7f3c7810303cf2a99a508cb66680536b1ac2e435928362fcc6514fc.AccessPackageAssignmentPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy_id"] = id
    }
    return i9616a87af7f3c7810303cf2a99a508cb66680536b1ac2e435928362fcc6514fc.NewAccessPackageAssignmentPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageCatalog()(*ic0419a87849e51ae8a63d5884bab9c32f3ddbf25318b1c51c2dd3e111752dcfd.AccessPackageCatalogRequestBuilder) {
    return ic0419a87849e51ae8a63d5884bab9c32f3ddbf25318b1c51c2dd3e111752dcfd.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopes()(*i8a8700fea501b7bf142c3430294dd2e707d3d2ce5820b7d1bef4ed9af4319f80.AccessPackageResourceRoleScopesRequestBuilder) {
    return i8a8700fea501b7bf142c3430294dd2e707d3d2ce5820b7d1bef4ed9af4319f80.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackages.item.accessPackageResourceRoleScopes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*if10b7ce5e756ef96bfc4cf4bb17c7bd568a9095cd3f8b40831944167f04716b8.AccessPackageResourceRoleScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope_id"] = id
    }
    return if10b7ce5e756ef96bfc4cf4bb17c7bd568a9095cd3f8b40831944167f04716b8.NewAccessPackageResourceRoleScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackagesIncompatibleWith()(*i5b548a88d8027e48ea15291f1dfe39a8927fe93dcca07173b0472f3798f2d017.AccessPackagesIncompatibleWithRequestBuilder) {
    return i5b548a88d8027e48ea15291f1dfe39a8927fe93dcca07173b0472f3798f2d017.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new AccessPackageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    m := &AccessPackageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}/accessPackageAssignments/{accessPackageAssignment_id}/accessPackageAssignmentPolicy/accessPackageCatalog/accessPackages/{accessPackage_id}{?select,expand}";
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
func (m *AccessPackageRequestBuilder) GetApplicablePolicyRequirements()(*ib224a2189234906b261ee0e2da4596d3a9d8047d5957ca375469a34ba88dde3c.GetApplicablePolicyRequirementsRequestBuilder) {
    return ib224a2189234906b261ee0e2da4596d3a9d8047d5957ca375469a34ba88dde3c.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleAccessPackages()(*if83c62939ddf8992405494eb0d52216cfd46c87f9a8c60c3763707c24531971f.IncompatibleAccessPackagesRequestBuilder) {
    return if83c62939ddf8992405494eb0d52216cfd46c87f9a8c60c3763707c24531971f.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleGroups()(*i9833a1555432d936ca4a849db79f7f2f28bce6ec20b1f571bf7dd15e45592fd1.IncompatibleGroupsRequestBuilder) {
    return i9833a1555432d936ca4a849db79f7f2f28bce6ec20b1f571bf7dd15e45592fd1.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackages.item.incompatibleGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageRequestBuilder) IncompatibleGroupsById(id string)(*ic15fa732fd94aca01d0c675909da26f44abd7b76c06c2524d9f8eb53be1aea72.GroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group_id"] = id
    }
    return ic15fa732fd94aca01d0c675909da26f44abd7b76c06c2524d9f8eb53be1aea72.NewGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
