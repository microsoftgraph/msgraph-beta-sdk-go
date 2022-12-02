package privilegedroleassignments

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder provides operations to manage the roleInfo property of the microsoft.graph.privilegedRoleAssignment entity.
type PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderGetQueryParameters read-only. Nullable. The associated role information.
type PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderGetQueryParameters
}
// PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.privilegedRole entity.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) Assignments()(*PrivilegedRoleAssignmentsItemRoleInfoAssignmentsRequestBuilder) {
    return NewPrivilegedRoleAssignmentsItemRoleInfoAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.privilegedRole entity.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) AssignmentsById(id string)(*PrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignment%2Did1"] = id
    }
    return NewPrivilegedRoleAssignmentsItemRoleInfoAssignmentsPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrivilegedRoleAssignmentsItemRoleInfoRequestBuilderInternal instantiates a new RoleInfoRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentsItemRoleInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) {
    m := &PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignments/{privilegedRoleAssignment%2Did}/roleInfo{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedRoleAssignmentsItemRoleInfoRequestBuilder instantiates a new RoleInfoRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentsItemRoleInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedRoleAssignmentsItemRoleInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleInfo for privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. The associated role information.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleInfo in privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleInfo for privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. The associated role information.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable), nil
}
// Patch update the navigation property roleInfo in privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, requestConfiguration *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable), nil
}
// SelfActivate provides operations to call the selfActivate method.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) SelfActivate()(*PrivilegedRoleAssignmentsItemRoleInfoSelfActivateRequestBuilder) {
    return NewPrivilegedRoleAssignmentsItemRoleInfoSelfActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SelfDeactivate provides operations to call the selfDeactivate method.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) SelfDeactivate()(*PrivilegedRoleAssignmentsItemRoleInfoSelfDeactivateRequestBuilder) {
    return NewPrivilegedRoleAssignmentsItemRoleInfoSelfDeactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.privilegedRole entity.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) Settings()(*PrivilegedRoleAssignmentsItemRoleInfoSettingsRequestBuilder) {
    return NewPrivilegedRoleAssignmentsItemRoleInfoSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Summary provides operations to manage the summary property of the microsoft.graph.privilegedRole entity.
func (m *PrivilegedRoleAssignmentsItemRoleInfoRequestBuilder) Summary()(*PrivilegedRoleAssignmentsItemRoleInfoSummaryRequestBuilder) {
    return NewPrivilegedRoleAssignmentsItemRoleInfoSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
