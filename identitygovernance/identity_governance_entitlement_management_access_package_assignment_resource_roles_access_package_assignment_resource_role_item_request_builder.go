package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.entitlementManagement entity.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters represents the resource-specific role which a subject has been assigned through an access package assignment.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignments provides operations to manage the accessPackageAssignments property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignments()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentsById provides operations to manage the accessPackageAssignments property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignmentsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceRole provides operations to manage the accessPackageResourceRole property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceRole()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScope provides operations to manage the accessPackageResourceScope property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceScope()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageSubject provides operations to manage the accessPackageSubject property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageSubject()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentResourceRoles for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageAssignmentResourceRoles in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property accessPackageAssignmentResourceRoles for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable), nil
}
// Patch update the navigation property accessPackageAssignmentResourceRoles in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable), nil
}
