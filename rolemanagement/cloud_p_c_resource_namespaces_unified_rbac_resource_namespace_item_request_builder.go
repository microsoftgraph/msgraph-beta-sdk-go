package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplicationMultiple entity.
type CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters get resourceNamespaces from roleManagement
type CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters
}
// CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewCloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    m := &CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/cloudPC/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewCloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resourceNamespaces for roleManagement
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get resourceNamespaces from roleManagement
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// MicrosoftGraphImportResourceActions provides operations to call the importResourceActions method.
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) MicrosoftGraphImportResourceActions()(*CloudPCResourceNamespacesItemMicrosoftGraphImportResourceActionsImportResourceActionsRequestBuilder) {
    return NewCloudPCResourceNamespacesItemMicrosoftGraphImportResourceActionsImportResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property resourceNamespaces in roleManagement
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ResourceActions provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActions()(*CloudPCResourceNamespacesItemResourceActionsRequestBuilder) {
    return NewCloudPCResourceNamespacesItemResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourceActionsById provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActionsById(id string)(*CloudPCResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceAction%2Did"] = id
    }
    return NewCloudPCResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property resourceNamespaces for roleManagement
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get resourceNamespaces from roleManagement
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property resourceNamespaces in roleManagement
func (m *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *CloudPCResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
