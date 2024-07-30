package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleAssignment entity.
type EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderGetQueryParameters the directory object that is the scope of the assignment. Read-only. Supports $expand for the directory provider.
type EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderGetQueryParameters
}
// NewEntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderInternal instantiates a new EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder and sets the default values.
func NewEntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder) {
    m := &EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/transitiveRoleAssignments/{unifiedRoleAssignment%2Did}/directoryScope{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder instantiates a new EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder and sets the default values.
func NewEntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the directory object that is the scope of the assignment. Read-only. Supports $expand for the directory provider.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// ToGetRequestInformation the directory object that is the scope of the assignment. Read-only. Supports $expand for the directory provider.
// returns a *RequestInformation when successful
func (m *EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder when successful
func (m *EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder) {
    return NewEntitlementManagementTransitiveRoleAssignmentsItemDirectoryScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
