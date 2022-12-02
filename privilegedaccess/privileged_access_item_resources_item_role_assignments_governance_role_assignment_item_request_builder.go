package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
type PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetQueryParameters the collection of role assignments for the resource.
type PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    m := &PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleAssignments/{governanceRoleAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for privilegedAccess
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of role assignments for the resource.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of role assignments for the resource.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
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
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) LinkedEligibleRoleAssignment()(*PrivilegedAccessItemResourcesItemRoleAssignmentsItemLinkedEligibleRoleAssignmentRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentsItemLinkedEligibleRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property roleAssignments in privilegedAccess
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
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
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Resource()(*PrivilegedAccessItemResourcesItemRoleAssignmentsItemResourceRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentsItemResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.governanceRoleAssignment entity.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) RoleDefinition()(*PrivilegedAccessItemResourcesItemRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Subject provides operations to manage the subject property of the microsoft.graph.governanceRoleAssignment entity.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) Subject()(*PrivilegedAccessItemResourcesItemRoleAssignmentsItemSubjectRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentsItemSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
