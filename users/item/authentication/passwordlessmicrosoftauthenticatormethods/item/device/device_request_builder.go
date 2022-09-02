package device

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06f637b5817ea7dcc7e77afe66dd17410c8989ef5f95b6792b641e407ed01d4f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/restore"
    i2bb29e3be96b80c4d1ea202bf9795454e1ac40085fbd8c28dcd549c121f4a2c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/usagerights"
    i2e51ac4f4142590854e942ff543005c994156800001ab732ac27492f0464c529 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/extensions"
    i3ca0cc6485aa121740580898158dd7023feb4cad3c8de1ae9cd39a41f33c7aa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/getmemberobjects"
    i6c88bea172a45d33e766ba9c9c0c5a9cc98205bcfb3c0848fcee43c54a4d07cb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers"
    i8b40ca361810f5142e8fbad8621b4c7a2abdf1b7211f1b19a3a75dbb588f3a2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof"
    i913d39905322d9e23c153dbf87462ae1a40b9fcd61f767115de7af3c3f9ab951 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/checkmembergroups"
    iab3b8054460e67a1a7eb284cb569867947437ebfb46d1002190c5f161255e01f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/getmembergroups"
    ib358a2ab61a067507dd7a6352c2c0ff271513a523654acb731f4967917cabec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/checkmemberobjects"
    ib3e43d1e1863c4cbfa1da65b8879db46d5dedc7401f1d876012888c7a3035fac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners"
    idbb7b382dd5de4788dc803c48840308863a56c97f923dcd5b1f96d9ecc1eb4ea "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/commands"
    if7d4ffa6f2a65f3e6632d676843b6f973374a880bdeffecc316991121b8076f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof"
    i0b3d3eff6f0c692fa600a2cdc0e24410417bd326cff0c598f184fe73be909004 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/commands/item"
    i146d79045ea17b9287dcf1e2b1127dd2a767666d763a9f93d1093c6122814b2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item"
    i191b3f6554c0035f7a55eef4d87d782ea73ac3729d12d74a885232394ea732a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/usagerights/item"
    i409d9eed47e75160313c229e4e11041cc4f3fe1e57a6d39f4e6d32d6386cc472 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item"
    i5b3935114692a197c635227b4624459fd562ad55fbd4eba43dfc168dcc1a08c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item"
    i69d90a742e916ade151e50cbab5a77ffd7bd6ab29ccf3035a60a517e0a0c3baf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item"
    i93aafe8bcd1a528ebb2b886da1bc4674d9896032247de7be28665296f2789ad0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/extensions/item"
)

// DeviceRequestBuilder provides operations to manage the device property of the microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethod entity.
type DeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceRequestBuilderGetQueryParameters get device from users
type DeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceRequestBuilderGetQueryParameters
}
// DeviceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups the checkMemberGroups property
func (m *DeviceRequestBuilder) CheckMemberGroups()(*i913d39905322d9e23c153dbf87462ae1a40b9fcd61f767115de7af3c3f9ab951.CheckMemberGroupsRequestBuilder) {
    return i913d39905322d9e23c153dbf87462ae1a40b9fcd61f767115de7af3c3f9ab951.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DeviceRequestBuilder) CheckMemberObjects()(*ib358a2ab61a067507dd7a6352c2c0ff271513a523654acb731f4967917cabec4.CheckMemberObjectsRequestBuilder) {
    return ib358a2ab61a067507dd7a6352c2c0ff271513a523654acb731f4967917cabec4.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commands the commands property
func (m *DeviceRequestBuilder) Commands()(*idbb7b382dd5de4788dc803c48840308863a56c97f923dcd5b1f96d9ecc1eb4ea.CommandsRequestBuilder) {
    return idbb7b382dd5de4788dc803c48840308863a56c97f923dcd5b1f96d9ecc1eb4ea.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.commands.item collection
func (m *DeviceRequestBuilder) CommandsById(id string)(*i0b3d3eff6f0c692fa600a2cdc0e24410417bd326cff0c598f184fe73be909004.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return i0b3d3eff6f0c692fa600a2cdc0e24410417bd326cff0c598f184fe73be909004.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property device for users
func (m *DeviceRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property device for users
func (m *DeviceRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get device from users
func (m *DeviceRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get device from users
func (m *DeviceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property device in users
func (m *DeviceRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property device in users
func (m *DeviceRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property device for users
func (m *DeviceRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Extensions the extensions property
func (m *DeviceRequestBuilder) Extensions()(*i2e51ac4f4142590854e942ff543005c994156800001ab732ac27492f0464c529.ExtensionsRequestBuilder) {
    return i2e51ac4f4142590854e942ff543005c994156800001ab732ac27492f0464c529.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.extensions.item collection
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*i93aafe8bcd1a528ebb2b886da1bc4674d9896032247de7be28665296f2789ad0.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i93aafe8bcd1a528ebb2b886da1bc4674d9896032247de7be28665296f2789ad0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get device from users
func (m *DeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// GetMemberGroups the getMemberGroups property
func (m *DeviceRequestBuilder) GetMemberGroups()(*iab3b8054460e67a1a7eb284cb569867947437ebfb46d1002190c5f161255e01f.GetMemberGroupsRequestBuilder) {
    return iab3b8054460e67a1a7eb284cb569867947437ebfb46d1002190c5f161255e01f.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DeviceRequestBuilder) GetMemberObjects()(*i3ca0cc6485aa121740580898158dd7023feb4cad3c8de1ae9cd39a41f33c7aa9.GetMemberObjectsRequestBuilder) {
    return i3ca0cc6485aa121740580898158dd7023feb4cad3c8de1ae9cd39a41f33c7aa9.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *DeviceRequestBuilder) MemberOf()(*i8b40ca361810f5142e8fbad8621b4c7a2abdf1b7211f1b19a3a75dbb588f3a2d.MemberOfRequestBuilder) {
    return i8b40ca361810f5142e8fbad8621b4c7a2abdf1b7211f1b19a3a75dbb588f3a2d.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.memberOf.item collection
func (m *DeviceRequestBuilder) MemberOfById(id string)(*i5b3935114692a197c635227b4624459fd562ad55fbd4eba43dfc168dcc1a08c8.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i5b3935114692a197c635227b4624459fd562ad55fbd4eba43dfc168dcc1a08c8.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in users
func (m *DeviceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// RegisteredOwners the registeredOwners property
func (m *DeviceRequestBuilder) RegisteredOwners()(*ib3e43d1e1863c4cbfa1da65b8879db46d5dedc7401f1d876012888c7a3035fac.RegisteredOwnersRequestBuilder) {
    return ib3e43d1e1863c4cbfa1da65b8879db46d5dedc7401f1d876012888c7a3035fac.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.registeredOwners.item collection
func (m *DeviceRequestBuilder) RegisteredOwnersById(id string)(*i69d90a742e916ade151e50cbab5a77ffd7bd6ab29ccf3035a60a517e0a0c3baf.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i69d90a742e916ade151e50cbab5a77ffd7bd6ab29ccf3035a60a517e0a0c3baf.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers the registeredUsers property
func (m *DeviceRequestBuilder) RegisteredUsers()(*i6c88bea172a45d33e766ba9c9c0c5a9cc98205bcfb3c0848fcee43c54a4d07cb.RegisteredUsersRequestBuilder) {
    return i6c88bea172a45d33e766ba9c9c0c5a9cc98205bcfb3c0848fcee43c54a4d07cb.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.registeredUsers.item collection
func (m *DeviceRequestBuilder) RegisteredUsersById(id string)(*i409d9eed47e75160313c229e4e11041cc4f3fe1e57a6d39f4e6d32d6386cc472.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i409d9eed47e75160313c229e4e11041cc4f3fe1e57a6d39f4e6d32d6386cc472.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore the restore property
func (m *DeviceRequestBuilder) Restore()(*i06f637b5817ea7dcc7e77afe66dd17410c8989ef5f95b6792b641e407ed01d4f.RestoreRequestBuilder) {
    return i06f637b5817ea7dcc7e77afe66dd17410c8989ef5f95b6792b641e407ed01d4f.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*if7d4ffa6f2a65f3e6632d676843b6f973374a880bdeffecc316991121b8076f2.TransitiveMemberOfRequestBuilder) {
    return if7d4ffa6f2a65f3e6632d676843b6f973374a880bdeffecc316991121b8076f2.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.transitiveMemberOf.item collection
func (m *DeviceRequestBuilder) TransitiveMemberOfById(id string)(*i146d79045ea17b9287dcf1e2b1127dd2a767666d763a9f93d1093c6122814b2c.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i146d79045ea17b9287dcf1e2b1127dd2a767666d763a9f93d1093c6122814b2c.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights the usageRights property
func (m *DeviceRequestBuilder) UsageRights()(*i2bb29e3be96b80c4d1ea202bf9795454e1ac40085fbd8c28dcd549c121f4a2c7.UsageRightsRequestBuilder) {
    return i2bb29e3be96b80c4d1ea202bf9795454e1ac40085fbd8c28dcd549c121f4a2c7.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item.device.usageRights.item collection
func (m *DeviceRequestBuilder) UsageRightsById(id string)(*i191b3f6554c0035f7a55eef4d87d782ea73ac3729d12d74a885232394ea732a2.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return i191b3f6554c0035f7a55eef4d87d782ea73ac3729d12d74a885232394ea732a2.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
