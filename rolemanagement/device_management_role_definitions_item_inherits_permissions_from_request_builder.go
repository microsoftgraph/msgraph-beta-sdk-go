package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
type DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderGetQueryParameters read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles support this attribute.
type DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderGetQueryParameters
}
// DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleDefinitionId1 provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
// returns a *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromUnifiedRoleDefinitionItemRequestBuilder when successful
func (m *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) ByUnifiedRoleDefinitionId1(unifiedRoleDefinitionId1 string)(*DeviceManagementRoleDefinitionsItemInheritsPermissionsFromUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleDefinitionId1 != "" {
        urlTplParams["unifiedRoleDefinition%2Did1"] = unifiedRoleDefinitionId1
    }
    return NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderInternal instantiates a new DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder and sets the default values.
func NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) {
    m := &DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/deviceManagement/roleDefinitions/{unifiedRoleDefinition%2Did}/inheritsPermissionsFrom{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder instantiates a new DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder and sets the default values.
func NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromCountRequestBuilder when successful
func (m *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) Count()(*DeviceManagementRoleDefinitionsItemInheritsPermissionsFromCountRequestBuilder) {
    return NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles support this attribute.
// returns a UnifiedRoleDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionCollectionResponseable), nil
}
// Post create new navigation property to inheritsPermissionsFrom for roleManagement
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
// ToGetRequestInformation read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles support this attribute.
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to inheritsPermissionsFrom for roleManagement
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder when successful
func (m *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) {
    return NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
