package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetQueryParameters retrieve an access package with a list of accessPackageResourceRoleScope objects. These objects represent the resource roles that an access package assigns to each subject. Each object links to an accessPackageResourceRole and an accessPackageResourceScope.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentPolicies provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesAccessPackageAssignmentPoliciesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) AccessPackageAssignmentPolicies()(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesAccessPackageAssignmentPoliciesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesAccessPackageAssignmentPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageCatalog provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) AccessPackageCatalog()(*EntitlementmanagementAccesspackagesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRoleScopes provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesAccessPackageResourceRoleScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) AccessPackageResourceRoleScopes()(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesAccessPackageResourceRoleScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesAccessPackageResourceRoleScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackagesIncompatibleWith provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackagesIncompatibleWithRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackagesIncompatibleWithRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackagesIncompatibleWithRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder instantiates a new EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an accessPackage object. You can't delete an access package if it has any accessPackageAssignment. To delete the access package, first query if there are any assignments with a filter to indicate the specific access package, such as: $filter=accessPackage/id eq 'a914b616-e04e-476b-aa37-91038f0b165b'. For more information on how to remove assignments that are still in the delivered state, see Remove an assignment.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-delete?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve an access package with a list of accessPackageResourceRoleScope objects. These objects represent the resource roles that an access package assigns to each subject. Each object links to an accessPackageResourceRole and an accessPackageResourceScope.
// returns a AccessPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-list-accesspackageresourcerolescopes?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable), nil
}
// GetApplicablePolicyRequirements provides operations to call the getApplicablePolicyRequirements method.
// returns a *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) GetApplicablePolicyRequirements()(*EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncompatibleAccessPackages provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncompatibleGroups provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) IncompatibleGroups()(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MoveToCatalog provides operations to call the moveToCatalog method.
// returns a *EntitlementmanagementAccesspackagesItemMovetocatalogMoveToCatalogRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) MoveToCatalog()(*EntitlementmanagementAccesspackagesItemMovetocatalogMoveToCatalogRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemMovetocatalogMoveToCatalogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update an existing accessPackage object to change one or more of its properties, such as the display name or description.
// returns a AccessPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-update?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable), nil
}
// ToDeleteRequestInformation delete an accessPackage object. You can't delete an access package if it has any accessPackageAssignment. To delete the access package, first query if there are any assignments with a filter to indicate the specific access package, such as: $filter=accessPackage/id eq 'a914b616-e04e-476b-aa37-91038f0b165b'. For more information on how to remove assignments that are still in the delivered state, see Remove an assignment.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve an access package with a list of accessPackageResourceRoleScope objects. These objects represent the resource roles that an access package assigns to each subject. Each object links to an accessPackageResourceRole and an accessPackageResourceScope.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update an existing accessPackage object to change one or more of its properties, such as the display name or description.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageable, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
