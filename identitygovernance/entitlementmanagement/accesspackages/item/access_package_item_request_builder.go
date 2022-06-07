package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2a44e775fb7f737b94d47f06032cc0ce0b4bfd2fbfa2463f91ebde6ab7f3db40 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages"
    i7554110a94219c17393947ff9abe170d93bdfb30d153839818ba04ffcc76de4f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes"
    i76ccc4db9daa784ae822d619ad983c2223a197dcc8d1fde6c1f4e9a639bf2bf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies"
    i9f6aa13376612aa214454151aaed75311be92072af5b00d7314fae36adcdfa19 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatiblegroups"
    icf4fae5627b4c00615f25f4291b050076f3e43516464e10c1555e560fa24a900 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackagecatalog"
    idde732b29dba3e8a39bef8255c30333987cd56f4f2de35a68cc15f74a585f58f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackagesincompatiblewith"
    ifd34e90f17620729556c1246884066ebb12a7b35296a94b49e16baf9acf9aef9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/getapplicablepolicyrequirements"
    i1b0b6dccee25c45cfa4e69c60a7929eb9be56ec91385cab0b45c904775a1c060 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackagesincompatiblewith/item"
    i32cd1387ebc6076b2999c52277e7c8235a84ad82ab68b7e7f5c3787e7e50ad58 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item"
    i62c3918d05bd24a32f20be566f2beb5a984d56ff052983f3f2d0b2b4b7a6659e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages/item"
    i6f598a4c0565fbb06735b2ed6cb90e45996b75462797a784a018e5156dc02b37 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes/item"
    if5122358f2125de5d7aac67a870d8a285d01b2414de16d1936714c8a51c64611 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatiblegroups/item"
)

// AccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
type AccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessPackageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageItemRequestBuilderGetQueryParameters represents access package objects.
type AccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessPackageItemRequestBuilderGetQueryParameters
}
// AccessPackageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentPolicies the accessPackageAssignmentPolicies property
func (m *AccessPackageItemRequestBuilder) AccessPackageAssignmentPolicies()(*i76ccc4db9daa784ae822d619ad983c2223a197dcc8d1fde6c1f4e9a639bf2bf2.AccessPackageAssignmentPoliciesRequestBuilder) {
    return i76ccc4db9daa784ae822d619ad983c2223a197dcc8d1fde6c1f4e9a639bf2bf2.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackageAssignmentPolicies.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i32cd1387ebc6076b2999c52277e7c8235a84ad82ab68b7e7f5c3787e7e50ad58.AccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return i32cd1387ebc6076b2999c52277e7c8235a84ad82ab68b7e7f5c3787e7e50ad58.NewAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageCatalog the accessPackageCatalog property
func (m *AccessPackageItemRequestBuilder) AccessPackageCatalog()(*icf4fae5627b4c00615f25f4291b050076f3e43516464e10c1555e560fa24a900.AccessPackageCatalogRequestBuilder) {
    return icf4fae5627b4c00615f25f4291b050076f3e43516464e10c1555e560fa24a900.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopes the accessPackageResourceRoleScopes property
func (m *AccessPackageItemRequestBuilder) AccessPackageResourceRoleScopes()(*i7554110a94219c17393947ff9abe170d93bdfb30d153839818ba04ffcc76de4f.AccessPackageResourceRoleScopesRequestBuilder) {
    return i7554110a94219c17393947ff9abe170d93bdfb30d153839818ba04ffcc76de4f.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackageResourceRoleScopes.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i6f598a4c0565fbb06735b2ed6cb90e45996b75462797a784a018e5156dc02b37.AccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope%2Did"] = id
    }
    return i6f598a4c0565fbb06735b2ed6cb90e45996b75462797a784a018e5156dc02b37.NewAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackagesIncompatibleWith the accessPackagesIncompatibleWith property
func (m *AccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*idde732b29dba3e8a39bef8255c30333987cd56f4f2de35a68cc15f74a585f58f.AccessPackagesIncompatibleWithRequestBuilder) {
    return idde732b29dba3e8a39bef8255c30333987cd56f4f2de35a68cc15f74a585f58f.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesIncompatibleWithById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackagesIncompatibleWith.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*i1b0b6dccee25c45cfa4e69c60a7929eb9be56ec91385cab0b45c904775a1c060.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return i1b0b6dccee25c45cfa4e69c60a7929eb9be56ec91385cab0b45c904775a1c060.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageItemRequestBuilder) {
    m := &AccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents access package objects.
func (m *AccessPackageItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents access package objects.
func (m *AccessPackageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *AccessPackageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *AccessPackageItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents access package objects.
func (m *AccessPackageItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetApplicablePolicyRequirements the getApplicablePolicyRequirements property
func (m *AccessPackageItemRequestBuilder) GetApplicablePolicyRequirements()(*ifd34e90f17620729556c1246884066ebb12a7b35296a94b49e16baf9acf9aef9.GetApplicablePolicyRequirementsRequestBuilder) {
    return ifd34e90f17620729556c1246884066ebb12a7b35296a94b49e16baf9acf9aef9.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler represents access package objects.
func (m *AccessPackageItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AccessPackageItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable), nil
}
// IncompatibleAccessPackages the incompatibleAccessPackages property
func (m *AccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*i2a44e775fb7f737b94d47f06032cc0ce0b4bfd2fbfa2463f91ebde6ab7f3db40.IncompatibleAccessPackagesRequestBuilder) {
    return i2a44e775fb7f737b94d47f06032cc0ce0b4bfd2fbfa2463f91ebde6ab7f3db40.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleAccessPackages.item collection
func (m *AccessPackageItemRequestBuilder) IncompatibleAccessPackagesById(id string)(*i62c3918d05bd24a32f20be566f2beb5a984d56ff052983f3f2d0b2b4b7a6659e.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return i62c3918d05bd24a32f20be566f2beb5a984d56ff052983f3f2d0b2b4b7a6659e.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IncompatibleGroups the incompatibleGroups property
func (m *AccessPackageItemRequestBuilder) IncompatibleGroups()(*i9f6aa13376612aa214454151aaed75311be92072af5b00d7314fae36adcdfa19.IncompatibleGroupsRequestBuilder) {
    return i9f6aa13376612aa214454151aaed75311be92072af5b00d7314fae36adcdfa19.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleGroups.item collection
func (m *AccessPackageItemRequestBuilder) IncompatibleGroupsById(id string)(*if5122358f2125de5d7aac67a870d8a285d01b2414de16d1936714c8a51c64611.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return if5122358f2125de5d7aac67a870d8a285d01b2414de16d1936714c8a51c64611.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *AccessPackageItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
