package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalItemRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a servicePrincipal object.
type ServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalItemRequestBuilderGetQueryParameters
}
// ServicePrincipalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddTokenSigningCertificate provides operations to call the addTokenSigningCertificate method.
func (m *ServicePrincipalItemRequestBuilder) AddTokenSigningCertificate()(*ItemAddTokenSigningCertificateRequestBuilder) {
    return NewItemAddTokenSigningCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManagementPolicies provides operations to manage the appManagementPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppManagementPolicies()(*ItemAppManagementPoliciesRequestBuilder) {
    return NewItemAppManagementPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppRoleAssignedTo provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedTo()(*ItemAppRoleAssignedToRequestBuilder) {
    return NewItemAppRoleAssignedToRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignments()(*ItemAppRoleAssignmentsRequestBuilder) {
    return NewItemAppRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ServicePrincipalItemRequestBuilder) CheckMemberGroups()(*ItemCheckMemberGroupsRequestBuilder) {
    return NewItemCheckMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ServicePrincipalItemRequestBuilder) CheckMemberObjects()(*ItemCheckMemberObjectsRequestBuilder) {
    return NewItemCheckMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPolicies()(*ItemClaimsMappingPoliciesRequestBuilder) {
    return NewItemClaimsMappingPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewServicePrincipalItemRequestBuilderInternal instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    m := &ServicePrincipalItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewServicePrincipalItemRequestBuilder instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) CreatedObjects()(*ItemCreatedObjectsRequestBuilder) {
    return NewItemCreatedObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatePasswordSingleSignOnCredentials provides operations to call the createPasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) CreatePasswordSingleSignOnCredentials()(*ItemCreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DelegatedPermissionClassifications provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassifications()(*ItemDelegatedPermissionClassificationsRequestBuilder) {
    return NewItemDelegatedPermissionClassificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a servicePrincipal object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-delete?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeletePasswordSingleSignOnCredentials provides operations to call the deletePasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) DeletePasswordSingleSignOnCredentials()(*ItemDeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Endpoints provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Endpoints()(*ItemEndpointsRequestBuilder) {
    return NewItemEndpointsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) FederatedIdentityCredentials()(*ItemFederatedIdentityCredentialsRequestBuilder) {
    return NewItemFederatedIdentityCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and relationships of a servicePrincipal object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-get?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *ServicePrincipalItemRequestBuilder) GetMemberGroups()(*ItemGetMemberGroupsRequestBuilder) {
    return NewItemGetMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *ServicePrincipalItemRequestBuilder) GetMemberObjects()(*ItemGetMemberObjectsRequestBuilder) {
    return NewItemGetMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetPasswordSingleSignOnCredentials provides operations to call the getPasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) GetPasswordSingleSignOnCredentials()(*ItemGetPasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewItemHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) LicenseDetails()(*ItemLicenseDetailsRequestBuilder) {
    return NewItemLicenseDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) MemberOf()(*ItemMemberOfRequestBuilder) {
    return NewItemMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrants()(*ItemOauth2PermissionGrantsRequestBuilder) {
    return NewItemOauth2PermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) OwnedObjects()(*ItemOwnedObjectsRequestBuilder) {
    return NewItemOwnedObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Owners()(*ItemOwnersRequestBuilder) {
    return NewItemOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of servicePrincipal object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-update?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable), nil
}
// Restore provides operations to call the restore method.
func (m *ServicePrincipalItemRequestBuilder) Restore()(*ItemRestoreRequestBuilder) {
    return NewItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Synchronization provides operations to manage the synchronization property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Synchronization()(*ItemSynchronizationRequestBuilder) {
    return NewItemSynchronizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePolicies()(*ItemTokenIssuancePoliciesRequestBuilder) {
    return NewItemTokenIssuancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePolicies()(*ItemTokenLifetimePoliciesRequestBuilder) {
    return NewItemTokenLifetimePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPatchRequestInformation update the properties of servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOf()(*ItemTransitiveMemberOfRequestBuilder) {
    return NewItemTransitiveMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdatePasswordSingleSignOnCredentials provides operations to call the updatePasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ServicePrincipalItemRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalItemRequestBuilder) {
    return NewServicePrincipalItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
