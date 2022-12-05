package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.privilegedAccess entity.
type PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetQueryParameters a collection of role assignments for the provider.
type PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    m := &PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/roleAssignments/{governanceRoleAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for privilegedAccess
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of role assignments for the provider.
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignments in privilegedAccess
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleAssignments for privilegedAccess
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of role assignments for the provider.
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable), nil
}
// LinkedEligibleRoleAssignment provides operations to manage the linkedEligibleRoleAssignment property of the microsoft.graph.governanceRoleAssignment entity.
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) LinkedEligibleRoleAssignment()(*PrivilegedAccessItemRoleAssignmentsItemLinkedEligibleRoleAssignmentRequestBuilder) {
    return NewPrivilegedAccessItemRoleAssignmentsItemLinkedEligibleRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property roleAssignments in privilegedAccess
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable), nil
}
// Resource provides operations to manage the resource property of the microsoft.graph.governanceRoleAssignment entity.
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Resource()(*PrivilegedAccessItemRoleAssignmentsItemResourceRequestBuilder) {
    return NewPrivilegedAccessItemRoleAssignmentsItemResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.governanceRoleAssignment entity.
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) RoleDefinition()(*PrivilegedAccessItemRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    return NewPrivilegedAccessItemRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Subject provides operations to manage the subject property of the microsoft.graph.governanceRoleAssignment entity.
func (m *PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Subject()(*PrivilegedAccessItemRoleAssignmentsItemSubjectRequestBuilder) {
    return NewPrivilegedAccessItemRoleAssignmentsItemSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
