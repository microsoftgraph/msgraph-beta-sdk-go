package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
type DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters the group policy presentations associated with the definition.
type DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal instantiates a new GroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/previousVersionDefinition/presentations/{groupPolicyPresentation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder instantiates a new GroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property presentations for deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the group policy presentations associated with the definition.
func (m *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property presentations in deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Definition provides operations to manage the definition property of the microsoft.graph.groupPolicyPresentation entity.
func (m *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Definition()(*DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property presentations for deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the group policy presentations associated with the definition.
func (m *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable), nil
}
// Patch update the navigation property presentations in deviceManagement
func (m *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable), nil
}
