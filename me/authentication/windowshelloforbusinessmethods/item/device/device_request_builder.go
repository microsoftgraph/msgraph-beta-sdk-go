package device

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07787b1628e2ea4526e2bee97e454b775916d3ef976868879eb64d2cc7f990a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers"
    i0b69de9349cd639549fc13841dae0b1d38e03c90f7c3cc140eaeb3265aad263f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners"
    i2344c54ddceaebae238610cb1c790ee9b7d3fb290f26e3b927cc79621a9ae7f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof"
    i2e6d67e0f80847cff52dba1e4d1598a67ecf098eafd2cad5755e90916d3b9367 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/checkmemberobjects"
    i6441e4b2aada4b4bd8a3e9791bac08da7e5aceecfef053530084824849eea00c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/commands"
    i6dd64153fd704ce1f022973ba51e19af334b8963fa1e1a604e1cba1b96b641d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/extensions"
    i7d305435fbf84c5562d1728a0911f436ed7537c5f5b801c04a02926650667d48 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/checkmembergroups"
    i841a8aff510e5a4a5e6ca9cc82a3fd85b32116a3df8402eb6639e6c034e4c22d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/restore"
    ia48b33dfe54fcfdf9df39b172d11e22f38b578000d0213acc0f939891157fe7b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/getmemberobjects"
    iba2750a5ce5fb00e221d191ba048d6089cb76e72d0f2b1c81d7fdd428c9bec2b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof"
    id71eee9323f05db3bb4943ca473a4bb89e65f7b50801607e1d1c66d276eb0596 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/usagerights"
    ieabfd3cf86d3b9b959f3b9fbb0db3e0089e63afebf0d0b8714b57dc1a75bfb7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/getmembergroups"
    i1fdde29b02c229c38836c664d8fb1511c356f305332b7025ad88f70c9d0b4700 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item"
    i28a37de6f15aa82d6bc7dc5073800894307c79d7c64c9408be0b83b92abdab93 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item"
    i5f4c5d57de967145a0d5c0c6ecd710addc462978b49f475115ba5217e626e7c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/extensions/item"
    i7dda5860a4a8309befbfd03a674813d893d3a67dff6a7b73805a3314cc10829a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item"
    ic42cf15b8530c0f65d196d0489bbea95d4b2ed3209eadabd02d3d0810c9d6c87 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/commands/item"
    if6d3c38bad9c36e3ff7262325c54f549ecdfb5c84752c38631852e15f6370511 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item"
    ifa8b3e6e745dac61d2f9c3a0f0d73a16ad981b9d7e0b0d7e93398091c2043b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/usagerights/item"
)

// DeviceRequestBuilder provides operations to manage the device property of the microsoft.graph.windowsHelloForBusinessAuthenticationMethod entity.
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
// DeviceRequestBuilderGetQueryParameters the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
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
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *DeviceRequestBuilder) CheckMemberGroups()(*i7d305435fbf84c5562d1728a0911f436ed7537c5f5b801c04a02926650667d48.CheckMemberGroupsRequestBuilder) {
    return i7d305435fbf84c5562d1728a0911f436ed7537c5f5b801c04a02926650667d48.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *DeviceRequestBuilder) CheckMemberObjects()(*i2e6d67e0f80847cff52dba1e4d1598a67ecf098eafd2cad5755e90916d3b9367.CheckMemberObjectsRequestBuilder) {
    return i2e6d67e0f80847cff52dba1e4d1598a67ecf098eafd2cad5755e90916d3b9367.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commands provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) Commands()(*i6441e4b2aada4b4bd8a3e9791bac08da7e5aceecfef053530084824849eea00c.CommandsRequestBuilder) {
    return i6441e4b2aada4b4bd8a3e9791bac08da7e5aceecfef053530084824849eea00c.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) CommandsById(id string)(*ic42cf15b8530c0f65d196d0489bbea95d4b2ed3209eadabd02d3d0810c9d6c87.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return ic42cf15b8530c0f65d196d0489bbea95d4b2ed3209eadabd02d3d0810c9d6c87.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property device for me
func (m *DeviceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property device in me
func (m *DeviceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property device for me
func (m *DeviceRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceRequestBuilder) Extensions()(*i6dd64153fd704ce1f022973ba51e19af334b8963fa1e1a604e1cba1b96b641d0.ExtensionsRequestBuilder) {
    return i6dd64153fd704ce1f022973ba51e19af334b8963fa1e1a604e1cba1b96b641d0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*i5f4c5d57de967145a0d5c0c6ecd710addc462978b49f475115ba5217e626e7c2.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i5f4c5d57de967145a0d5c0c6ecd710addc462978b49f475115ba5217e626e7c2.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *DeviceRequestBuilder) GetMemberGroups()(*ieabfd3cf86d3b9b959f3b9fbb0db3e0089e63afebf0d0b8714b57dc1a75bfb7a.GetMemberGroupsRequestBuilder) {
    return ieabfd3cf86d3b9b959f3b9fbb0db3e0089e63afebf0d0b8714b57dc1a75bfb7a.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *DeviceRequestBuilder) GetMemberObjects()(*ia48b33dfe54fcfdf9df39b172d11e22f38b578000d0213acc0f939891157fe7b.GetMemberObjectsRequestBuilder) {
    return ia48b33dfe54fcfdf9df39b172d11e22f38b578000d0213acc0f939891157fe7b.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) MemberOf()(*i2344c54ddceaebae238610cb1c790ee9b7d3fb290f26e3b927cc79621a9ae7f3.MemberOfRequestBuilder) {
    return i2344c54ddceaebae238610cb1c790ee9b7d3fb290f26e3b927cc79621a9ae7f3.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) MemberOfById(id string)(*i1fdde29b02c229c38836c664d8fb1511c356f305332b7025ad88f70c9d0b4700.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i1fdde29b02c229c38836c664d8fb1511c356f305332b7025ad88f70c9d0b4700.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in me
func (m *DeviceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *DeviceRequestBuilder) RegisteredOwners()(*i0b69de9349cd639549fc13841dae0b1d38e03c90f7c3cc140eaeb3265aad263f.RegisteredOwnersRequestBuilder) {
    return i0b69de9349cd639549fc13841dae0b1d38e03c90f7c3cc140eaeb3265aad263f.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item.device.registeredOwners.item collection
func (m *DeviceRequestBuilder) RegisteredOwnersById(id string)(*if6d3c38bad9c36e3ff7262325c54f549ecdfb5c84752c38631852e15f6370511.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return if6d3c38bad9c36e3ff7262325c54f549ecdfb5c84752c38631852e15f6370511.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) RegisteredUsers()(*i07787b1628e2ea4526e2bee97e454b775916d3ef976868879eb64d2cc7f990a1.RegisteredUsersRequestBuilder) {
    return i07787b1628e2ea4526e2bee97e454b775916d3ef976868879eb64d2cc7f990a1.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) RegisteredUsersById(id string)(*i28a37de6f15aa82d6bc7dc5073800894307c79d7c64c9408be0b83b92abdab93.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i28a37de6f15aa82d6bc7dc5073800894307c79d7c64c9408be0b83b92abdab93.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *DeviceRequestBuilder) Restore()(*i841a8aff510e5a4a5e6ca9cc82a3fd85b32116a3df8402eb6639e6c034e4c22d.RestoreRequestBuilder) {
    return i841a8aff510e5a4a5e6ca9cc82a3fd85b32116a3df8402eb6639e6c034e4c22d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*iba2750a5ce5fb00e221d191ba048d6089cb76e72d0f2b1c81d7fdd428c9bec2b.TransitiveMemberOfRequestBuilder) {
    return iba2750a5ce5fb00e221d191ba048d6089cb76e72d0f2b1c81d7fdd428c9bec2b.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) TransitiveMemberOfById(id string)(*i7dda5860a4a8309befbfd03a674813d893d3a67dff6a7b73805a3314cc10829a.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i7dda5860a4a8309befbfd03a674813d893d3a67dff6a7b73805a3314cc10829a.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) UsageRights()(*id71eee9323f05db3bb4943ca473a4bb89e65f7b50801607e1d1c66d276eb0596.UsageRightsRequestBuilder) {
    return id71eee9323f05db3bb4943ca473a4bb89e65f7b50801607e1d1c66d276eb0596.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *DeviceRequestBuilder) UsageRightsById(id string)(*ifa8b3e6e745dac61d2f9c3a0f0d73a16ad981b9d7e0b0d7e93398091c2043b3c.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return ifa8b3e6e745dac61d2f9c3a0f0d73a16ad981b9d7e0b0d7e93398091c2043b3c.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
