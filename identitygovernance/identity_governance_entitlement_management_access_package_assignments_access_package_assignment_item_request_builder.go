package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder provides operations to manage the accessPackageAssignments property of the microsoft.graph.entitlementManagement entity.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderGetQueryParameters the assignment of an access package to a subject for a period of time.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) AccessPackage()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPolicy provides operations to manage the accessPackageAssignmentPolicy property of the microsoft.graph.accessPackageAssignment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentPolicy()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentPolicyRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequests provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.accessPackageAssignment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentRequests()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.accessPackageAssignment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentResourceRoles provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.accessPackageAssignment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentResourceRoles()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.accessPackageAssignment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderInternal instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignments for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the assignment of an access package to a subject for a period of time.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageAssignments in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageAssignments for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the assignment of an access package to a subject for a period of time.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable), nil
}
// Patch update the navigation property accessPackageAssignments in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable), nil
}
// Reprocess provides operations to call the reprocess method.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) Reprocess()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemReprocessRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Target provides operations to manage the target property of the microsoft.graph.accessPackageAssignment entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) Target()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemTargetRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemTargetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
