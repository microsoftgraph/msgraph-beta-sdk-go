package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
type DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters the available group policy categories for this account.
type DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.groupPolicyCategory entity.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Children()(*DeviceManagementGroupPolicyCategoriesItemChildrenRequestBuilder) {
    return NewDeviceManagementGroupPolicyCategoriesItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.groupPolicyCategory entity.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) ChildrenById(id string)(*DeviceManagementGroupPolicyCategoriesItemChildrenGroupPolicyCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyCategory%2Did1"] = id
    }
    return NewDeviceManagementGroupPolicyCategoriesItemChildrenGroupPolicyCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyCategories for deviceManagement
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the available group policy categories for this account.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyCategories in deviceManagement
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyCategory entity.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) DefinitionFile()(*DeviceManagementGroupPolicyCategoriesItemDefinitionFileRequestBuilder) {
    return NewDeviceManagementGroupPolicyCategoriesItemDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Definitions provides operations to manage the definitions property of the microsoft.graph.groupPolicyCategory entity.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Definitions()(*DeviceManagementGroupPolicyCategoriesItemDefinitionsRequestBuilder) {
    return NewDeviceManagementGroupPolicyCategoriesItemDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefinitionsById provides operations to manage the definitions property of the microsoft.graph.groupPolicyCategory entity.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) DefinitionsById(id string)(*DeviceManagementGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinition%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property groupPolicyCategories for deviceManagement
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available group policy categories for this account.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable), nil
}
// Parent provides operations to manage the parent property of the microsoft.graph.groupPolicyCategory entity.
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Parent()(*DeviceManagementGroupPolicyCategoriesItemParentRequestBuilder) {
    return NewDeviceManagementGroupPolicyCategoriesItemParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property groupPolicyCategories in deviceManagement
func (m *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable), nil
}
