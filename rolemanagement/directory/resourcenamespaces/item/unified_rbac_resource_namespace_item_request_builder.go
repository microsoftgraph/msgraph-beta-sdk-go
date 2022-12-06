package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i23eeb299d8c7655a5b8f9d91dcc2657946c4c182d94ebf89a3ccfe073344b511 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces/item/resourceactions"
    i57882143c214ec8547e9e29f83c78b6ba8a5a9e82bf6f8cf44bb035bec347ebb "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces/item/importresourceactions"
    i0ba042109e088f641da71389cf2ee1d6834428aef79b7f6483d9fafa2f43ab7a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces/item/resourceactions/item"
)

// UnifiedRbacResourceNamespaceItemRequestBuilder provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
type UnifiedRbacResourceNamespaceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters get resourceNamespaces from roleManagement
type UnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters
}
// UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRbacResourceNamespaceItemRequestBuilder) {
    m := &UnifiedRbacResourceNamespaceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRbacResourceNamespaceItemRequestBuilder instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewUnifiedRbacResourceNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resourceNamespaces for roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get resourceNamespaces from roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property resourceNamespaces in roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property resourceNamespaces for roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get resourceNamespaces from roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ImportResourceActions provides operations to call the importResourceActions method.
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) ImportResourceActions()(*i57882143c214ec8547e9e29f83c78b6ba8a5a9e82bf6f8cf44bb035bec347ebb.ImportResourceActionsRequestBuilder) {
    return i57882143c214ec8547e9e29f83c78b6ba8a5a9e82bf6f8cf44bb035bec347ebb.NewImportResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property resourceNamespaces in roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ResourceActions provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActions()(*i23eeb299d8c7655a5b8f9d91dcc2657946c4c182d94ebf89a3ccfe073344b511.ResourceActionsRequestBuilder) {
    return i23eeb299d8c7655a5b8f9d91dcc2657946c4c182d94ebf89a3ccfe073344b511.NewResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceActionsById provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActionsById(id string)(*i0ba042109e088f641da71389cf2ee1d6834428aef79b7f6483d9fafa2f43ab7a.UnifiedRbacResourceActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceAction%2Did"] = id
    }
    return i0ba042109e088f641da71389cf2ee1d6834428aef79b7f6483d9fafa2f43ab7a.NewUnifiedRbacResourceActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
