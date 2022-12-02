package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
type RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderGetQueryParameters specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
type RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderGetQueryParameters
}
// NewRoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderInternal instantiates a new RoleDefinitionRequestBuilder and sets the default values.
func NewRoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    m := &RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/roleDefinition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder instantiates a new RoleDefinitionRequestBuilder and sets the default values.
func NewRoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
func (m *RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
func (m *RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
