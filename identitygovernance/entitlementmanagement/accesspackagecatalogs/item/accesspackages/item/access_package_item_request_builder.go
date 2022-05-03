package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackagecatalog"
    i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackagesincompatiblewith"
    i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageresourcerolescopes"
    iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageassignmentpolicies"
    ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/getapplicablepolicyrequirements"
    idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/incompatibleaccesspackages"
    ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/incompatiblegroups"
    i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageassignmentpolicies/item"
    i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageresourcerolescopes/item"
    i6f48378de44ff087fa9221ec421efe295539bc5f8d3b597309f610323971f5b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackagesincompatiblewith/item"
    icf321ca159da7fc158842233fd8950ddac3820adfad012fa13287c5768d9b6ae "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/incompatibleaccesspackages/item"
    if8d5f68ff23b3dce782f1ab3c253f9a2850bce73a70d781dd74153c1477ce5be "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/incompatiblegroups/item"
)

// AccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
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
// AccessPackageItemRequestBuilderGetQueryParameters the access packages in this catalog. Read-only. Nullable. Supports $expand.
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
func (m *AccessPackageItemRequestBuilder) AccessPackageAssignmentPolicies()(*iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e.AccessPackageAssignmentPoliciesRequestBuilder) {
    return iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.accessPackageAssignmentPolicies.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9.AccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9.NewAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageCatalog the accessPackageCatalog property
func (m *AccessPackageItemRequestBuilder) AccessPackageCatalog()(*i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22.AccessPackageCatalogRequestBuilder) {
    return i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopes the accessPackageResourceRoleScopes property
func (m *AccessPackageItemRequestBuilder) AccessPackageResourceRoleScopes()(*i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d.AccessPackageResourceRoleScopesRequestBuilder) {
    return i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.accessPackageResourceRoleScopes.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd.AccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope%2Did"] = id
    }
    return i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd.NewAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackagesIncompatibleWith the accessPackagesIncompatibleWith property
func (m *AccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980.AccessPackagesIncompatibleWithRequestBuilder) {
    return i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesIncompatibleWithById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.accessPackagesIncompatibleWith.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*i6f48378de44ff087fa9221ec421efe295539bc5f8d3b597309f610323971f5b7.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return i6f48378de44ff087fa9221ec421efe295539bc5f8d3b597309f610323971f5b7.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageItemRequestBuilder) {
    m := &AccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AccessPackageItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AccessPackageItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetApplicablePolicyRequirements the getApplicablePolicyRequirements property
func (m *AccessPackageItemRequestBuilder) GetApplicablePolicyRequirements()(*ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb.GetApplicablePolicyRequirementsRequestBuilder) {
    return ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler the access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AccessPackageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the access packages in this catalog. Read-only. Nullable. Supports $expand.
func (m *AccessPackageItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AccessPackageItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
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
func (m *AccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe.IncompatibleAccessPackagesRequestBuilder) {
    return idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.incompatibleAccessPackages.item collection
func (m *AccessPackageItemRequestBuilder) IncompatibleAccessPackagesById(id string)(*icf321ca159da7fc158842233fd8950ddac3820adfad012fa13287c5768d9b6ae.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return icf321ca159da7fc158842233fd8950ddac3820adfad012fa13287c5768d9b6ae.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IncompatibleGroups the incompatibleGroups property
func (m *AccessPackageItemRequestBuilder) IncompatibleGroups()(*ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1.IncompatibleGroupsRequestBuilder) {
    return ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.incompatibleGroups.item collection
func (m *AccessPackageItemRequestBuilder) IncompatibleGroupsById(id string)(*if8d5f68ff23b3dce782f1ab3c253f9a2850bce73a70d781dd74153c1477ce5be.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return if8d5f68ff23b3dce782f1ab3c253f9a2850bce73a70d781dd74153c1477ce5be.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *AccessPackageItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *AccessPackageItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
