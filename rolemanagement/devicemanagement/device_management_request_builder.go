package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i16ba9d37411ef84c105365cabee4cb8ed289e66e837ce2bc6d3458b7cdcfb9e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments"
    i74d924764342fa6022f3fd5e5f9706420912c7cdd1a2bb0b50c8bf67c140c163 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces"
    ic7f2d26945adadc926e1a8a84b9142418a67273d0f9ba3ad2fda23bc4c71665e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roledefinitions"
    i2f6986ad54e9019a27955e607e7a5e6a3900d7d2a186631d40a82c7420f91a15 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item"
    ibf3ec838b277e1140e5e57824ced75b7c6eed602c42fa7b33e6783984c251265 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item"
    id1e540118321507d6c775d35c43257337ecb59533ab19347820eefca0118bb2c "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roledefinitions/item"
)

// DeviceManagementRequestBuilder provides operations to manage the deviceManagement property of the microsoft.graph.roleManagement entity.
type DeviceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementRequestBuilderGetQueryParameters the RbacApplication for Device Management
type DeviceManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementRequestBuilderGetQueryParameters
}
// DeviceManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementRequestBuilderInternal instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    m := &DeviceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceManagement for roleManagement
func (m *DeviceManagementRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceManagement in roleManagement
func (m *DeviceManagementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationMultipleable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceManagement for roleManagement
func (m *DeviceManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationMultipleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationMultipleable), nil
}
// Patch update the navigation property deviceManagement in roleManagement
func (m *DeviceManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationMultipleable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationMultipleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationMultipleable), nil
}
// ResourceNamespaces provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplicationMultiple entity.
func (m *DeviceManagementRequestBuilder) ResourceNamespaces()(*i74d924764342fa6022f3fd5e5f9706420912c7cdd1a2bb0b50c8bf67c140c163.ResourceNamespacesRequestBuilder) {
    return i74d924764342fa6022f3fd5e5f9706420912c7cdd1a2bb0b50c8bf67c140c163.NewResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplicationMultiple entity.
func (m *DeviceManagementRequestBuilder) ResourceNamespacesById(id string)(*i2f6986ad54e9019a27955e607e7a5e6a3900d7d2a186631d40a82c7420f91a15.UnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace%2Did"] = id
    }
    return i2f6986ad54e9019a27955e607e7a5e6a3900d7d2a186631d40a82c7420f91a15.NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplicationMultiple entity.
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*i16ba9d37411ef84c105365cabee4cb8ed289e66e837ce2bc6d3458b7cdcfb9e9.RoleAssignmentsRequestBuilder) {
    return i16ba9d37411ef84c105365cabee4cb8ed289e66e837ce2bc6d3458b7cdcfb9e9.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplicationMultiple entity.
func (m *DeviceManagementRequestBuilder) RoleAssignmentsById(id string)(*ibf3ec838b277e1140e5e57824ced75b7c6eed602c42fa7b33e6783984c251265.UnifiedRoleAssignmentMultipleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentMultiple%2Did"] = id
    }
    return ibf3ec838b277e1140e5e57824ced75b7c6eed602c42fa7b33e6783984c251265.NewUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplicationMultiple entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*ic7f2d26945adadc926e1a8a84b9142418a67273d0f9ba3ad2fda23bc4c71665e.RoleDefinitionsRequestBuilder) {
    return ic7f2d26945adadc926e1a8a84b9142418a67273d0f9ba3ad2fda23bc4c71665e.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplicationMultiple entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitionsById(id string)(*id1e540118321507d6c775d35c43257337ecb59533ab19347820eefca0118bb2c.UnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = id
    }
    return id1e540118321507d6c775d35c43257337ecb59533ab19347820eefca0118bb2c.NewUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
