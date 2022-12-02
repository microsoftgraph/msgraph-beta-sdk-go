package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
type DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters the available group policy definitions for this account.
type DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Category provides operations to manage the category property of the microsoft.graph.groupPolicyDefinition entity.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) Category()(*DeviceManagementGroupPolicyDefinitionsItemCategoryRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionsItemCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyDefinitions for deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the available group policy definitions for this account.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyDefinitions in deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyDefinition entity.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) DefinitionFile()(*DeviceManagementGroupPolicyDefinitionsItemDefinitionFileRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionsItemDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property groupPolicyDefinitions for deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available group policy definitions for this account.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable), nil
}
// NextVersionDefinition provides operations to manage the nextVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) NextVersionDefinition()(*DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property groupPolicyDefinitions in deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable), nil
}
// Presentations provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) Presentations()(*DeviceManagementGroupPolicyDefinitionsItemPresentationsRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionsItemPresentationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationsById provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) PresentationsById(id string)(*DeviceManagementGroupPolicyDefinitionsItemPresentationsGroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentation%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyDefinitionsItemPresentationsGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PreviousVersionDefinition provides operations to manage the previousVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
func (m *DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) PreviousVersionDefinition()(*DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
