package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder provides operations to manage the device property of the microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethod entity.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderGetQueryParameters get device from me
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderGetQueryParameters
}
// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) CheckMemberGroups()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) CheckMemberObjects()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commands provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) Commands()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) CommandsById(id string)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) {
    m := &AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property device for me
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get device from me
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property device in me
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property device for me
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) Extensions()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceExtensionsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) ExtensionsById(id string)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get device from me
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) GetMemberGroups()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) GetMemberObjects()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) MemberOf()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) MemberOfById(id string)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in me
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) RegisteredOwners()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.registeredOwners.item collection
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) RegisteredOwnersById(id string)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) RegisteredUsers()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) RegisteredUsersById(id string)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) Restore()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRestoreRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) TransitiveMemberOf()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) TransitiveMemberOfById(id string)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) UsageRights()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceUsageRightsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) UsageRightsById(id string)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceUsageRightsUsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceUsageRightsUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
