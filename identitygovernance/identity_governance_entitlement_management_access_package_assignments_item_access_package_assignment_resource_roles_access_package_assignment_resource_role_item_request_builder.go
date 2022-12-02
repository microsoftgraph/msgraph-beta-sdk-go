package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.accessPackageAssignment entity.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters the resource roles delivered to the target user for this assignment. Read-only. Nullable.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignments provides operations to manage the accessPackageAssignments property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignments()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentsById provides operations to manage the accessPackageAssignments property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignmentsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment%2Did1"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceRole provides operations to manage the accessPackageResourceRole property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceRole()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScope provides operations to manage the accessPackageResourceScope property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceScope()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageSubject provides operations to manage the accessPackageSubject property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageSubject()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentResourceRoles for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, error) {
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
