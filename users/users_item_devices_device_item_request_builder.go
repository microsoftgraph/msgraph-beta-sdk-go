package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemDevicesDeviceItemRequestBuilder provides operations to manage the devices property of the microsoft.graph.user entity.
type UsersItemDevicesDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemDevicesDeviceItemRequestBuilderGetQueryParameters get devices from users
type UsersItemDevicesDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemDevicesDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDevicesDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemDevicesDeviceItemRequestBuilderGetQueryParameters
}
// UsersItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *UsersItemDevicesDeviceItemRequestBuilder) CheckMemberGroups()(*UsersItemDevicesItemCheckMemberGroupsRequestBuilder) {
    return NewUsersItemDevicesItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *UsersItemDevicesDeviceItemRequestBuilder) CheckMemberObjects()(*UsersItemDevicesItemCheckMemberObjectsRequestBuilder) {
    return NewUsersItemDevicesItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commands provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) Commands()(*UsersItemDevicesItemCommandsRequestBuilder) {
    return NewUsersItemDevicesItemCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) CommandsById(id string)(*UsersItemDevicesItemCommandsCommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return NewUsersItemDevicesItemCommandsCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemDevicesDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewUsersItemDevicesDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDevicesDeviceItemRequestBuilder) {
    m := &UsersItemDevicesDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDevicesDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewUsersItemDevicesDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDevicesDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDevicesDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property devices for users
func (m *UsersItemDevicesDeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get devices from users
func (m *UsersItemDevicesDeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemDevicesDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property devices in users
func (m *UsersItemDevicesDeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *UsersItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property devices for users
func (m *UsersItemDevicesDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) Extensions()(*UsersItemDevicesItemExtensionsRequestBuilder) {
    return NewUsersItemDevicesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) ExtensionsById(id string)(*UsersItemDevicesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemDevicesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get devices from users
func (m *UsersItemDevicesDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemDevicesDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *UsersItemDevicesDeviceItemRequestBuilder) GetMemberGroups()(*UsersItemDevicesItemGetMemberGroupsRequestBuilder) {
    return NewUsersItemDevicesItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *UsersItemDevicesDeviceItemRequestBuilder) GetMemberObjects()(*UsersItemDevicesItemGetMemberObjectsRequestBuilder) {
    return NewUsersItemDevicesItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) MemberOf()(*UsersItemDevicesItemMemberOfRequestBuilder) {
    return NewUsersItemDevicesItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) MemberOfById(id string)(*UsersItemDevicesItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemDevicesItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property devices in users
func (m *UsersItemDevicesDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *UsersItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable), nil
}
// RegisteredOwners provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) RegisteredOwners()(*UsersItemDevicesItemRegisteredOwnersRequestBuilder) {
    return NewUsersItemDevicesItemRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.registeredOwners.item collection
func (m *UsersItemDevicesDeviceItemRequestBuilder) RegisteredOwnersById(id string)(*UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) RegisteredUsers()(*UsersItemDevicesItemRegisteredUsersRequestBuilder) {
    return NewUsersItemDevicesItemRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) RegisteredUsersById(id string)(*UsersItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *UsersItemDevicesDeviceItemRequestBuilder) Restore()(*UsersItemDevicesItemRestoreRequestBuilder) {
    return NewUsersItemDevicesItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) TransitiveMemberOf()(*UsersItemDevicesItemTransitiveMemberOfRequestBuilder) {
    return NewUsersItemDevicesItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) TransitiveMemberOfById(id string)(*UsersItemDevicesItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemDevicesItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) UsageRights()(*UsersItemDevicesItemUsageRightsRequestBuilder) {
    return NewUsersItemDevicesItemUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *UsersItemDevicesDeviceItemRequestBuilder) UsageRightsById(id string)(*UsersItemDevicesItemUsageRightsUsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return NewUsersItemDevicesItemUsageRightsUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
