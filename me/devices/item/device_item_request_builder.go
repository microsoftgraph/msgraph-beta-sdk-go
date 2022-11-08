package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06de72e85c136548b632475aeec5362fa29354e14a94fe0f6befb3d7d0c7ee08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/extensions"
    i2d7c976ed5ceb8455050d02a0603e4daf6148156215fc4056825711cb8d13da3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof"
    i37a7b24866d0eac1ecc6144ca60c9e826f0db376946285c74089934926209b12 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/checkmembergroups"
    i41a37091391970a2699d29bd22aaa56b4501fb9220844645984c749d80067404 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/getmembergroups"
    i6b4891e8e524165f041b742d05729e1561eab385718b042b623c5f072b36c6b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/restore"
    i7767512c114fae7428f00ecc379bdb6c40770b44d591d609b62d288e8fb6e1a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/usagerights"
    ia09559070bc6c3c9adcb81ec9ee248e1edc481381bdc642dcbc165846ecfb443 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof"
    iafe6422085b6e90cc2b738a384d5bace00ab3e5402091eaf62c14c7e5aa0a0b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/checkmemberobjects"
    ib0f5f2d9bcd8e673f2e034767fd3f29d1635342a9f9433713ff733937a83bc7f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/getmemberobjects"
    ib367c374a26076087ddba4f9b886fc919d9fa57297ddc426513b74650121bfbb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/commands"
    ib911ca91b9101c7c48053251c8d96b7673ac91153da66f0307d0d4e1cfffd8b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers"
    ibe08c9723c639b1b431cfeb009d22d7bc584bc4ec32861b964ec222077f1feb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners"
    i2af1416da365c0340a39d8ecc556e26557cb9f58ce6560792d2198c137dc55b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/commands/item"
    i2e2e998caa42ebe0146db25ae496c4ffcc4120a2786ba8d28386223ee7843679 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item"
    i4a299936ca517941a98d45670fb79ba0a2da5a3ca11d12e1c5bbd1579d31d553 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item"
    i534ed1cb37dd623a3ea412e67b0dca570cb44b8681f511614639ffd95b6791c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/usagerights/item"
    i9fd4cc16d3894463aa8887400816cd691765b2f88389be9d15e959a82eb1a7ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/extensions/item"
    ib92c409c5832dab5202e7c30ad3be7d04087f9f45c98fb4009798c7f34318347 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item"
    iec54264a6c486f68435119cb745dbefdc6e215ba637f94299ef3561dd30389aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item"
)

// DeviceItemRequestBuilder provides operations to manage the devices property of the microsoft.graph.user entity.
type DeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceItemRequestBuilderGetQueryParameters get devices from me
type DeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceItemRequestBuilderGetQueryParameters
}
// DeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *DeviceItemRequestBuilder) CheckMemberGroups()(*i37a7b24866d0eac1ecc6144ca60c9e826f0db376946285c74089934926209b12.CheckMemberGroupsRequestBuilder) {
    return i37a7b24866d0eac1ecc6144ca60c9e826f0db376946285c74089934926209b12.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *DeviceItemRequestBuilder) CheckMemberObjects()(*iafe6422085b6e90cc2b738a384d5bace00ab3e5402091eaf62c14c7e5aa0a0b1.CheckMemberObjectsRequestBuilder) {
    return iafe6422085b6e90cc2b738a384d5bace00ab3e5402091eaf62c14c7e5aa0a0b1.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commands provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) Commands()(*ib367c374a26076087ddba4f9b886fc919d9fa57297ddc426513b74650121bfbb.CommandsRequestBuilder) {
    return ib367c374a26076087ddba4f9b886fc919d9fa57297ddc426513b74650121bfbb.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) CommandsById(id string)(*i2af1416da365c0340a39d8ecc556e26557cb9f58ce6560792d2198c137dc55b5.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return i2af1416da365c0340a39d8ecc556e26557cb9f58ce6560792d2198c137dc55b5.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property devices for me
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get devices from me
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property devices in me
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property devices for me
func (m *DeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceItemRequestBuilder) Extensions()(*i06de72e85c136548b632475aeec5362fa29354e14a94fe0f6befb3d7d0c7ee08.ExtensionsRequestBuilder) {
    return i06de72e85c136548b632475aeec5362fa29354e14a94fe0f6befb3d7d0c7ee08.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) ExtensionsById(id string)(*i9fd4cc16d3894463aa8887400816cd691765b2f88389be9d15e959a82eb1a7ed.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9fd4cc16d3894463aa8887400816cd691765b2f88389be9d15e959a82eb1a7ed.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get devices from me
func (m *DeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *DeviceItemRequestBuilder) GetMemberGroups()(*i41a37091391970a2699d29bd22aaa56b4501fb9220844645984c749d80067404.GetMemberGroupsRequestBuilder) {
    return i41a37091391970a2699d29bd22aaa56b4501fb9220844645984c749d80067404.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *DeviceItemRequestBuilder) GetMemberObjects()(*ib0f5f2d9bcd8e673f2e034767fd3f29d1635342a9f9433713ff733937a83bc7f.GetMemberObjectsRequestBuilder) {
    return ib0f5f2d9bcd8e673f2e034767fd3f29d1635342a9f9433713ff733937a83bc7f.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) MemberOf()(*ia09559070bc6c3c9adcb81ec9ee248e1edc481381bdc642dcbc165846ecfb443.MemberOfRequestBuilder) {
    return ia09559070bc6c3c9adcb81ec9ee248e1edc481381bdc642dcbc165846ecfb443.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) MemberOfById(id string)(*i2e2e998caa42ebe0146db25ae496c4ffcc4120a2786ba8d28386223ee7843679.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i2e2e998caa42ebe0146db25ae496c4ffcc4120a2786ba8d28386223ee7843679.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property devices in me
func (m *DeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *DeviceItemRequestBuilder) RegisteredOwners()(*ibe08c9723c639b1b431cfeb009d22d7bc584bc4ec32861b964ec222077f1feb1.RegisteredOwnersRequestBuilder) {
    return ibe08c9723c639b1b431cfeb009d22d7bc584bc4ec32861b964ec222077f1feb1.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item.registeredOwners.item collection
func (m *DeviceItemRequestBuilder) RegisteredOwnersById(id string)(*i4a299936ca517941a98d45670fb79ba0a2da5a3ca11d12e1c5bbd1579d31d553.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i4a299936ca517941a98d45670fb79ba0a2da5a3ca11d12e1c5bbd1579d31d553.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) RegisteredUsers()(*ib911ca91b9101c7c48053251c8d96b7673ac91153da66f0307d0d4e1cfffd8b2.RegisteredUsersRequestBuilder) {
    return ib911ca91b9101c7c48053251c8d96b7673ac91153da66f0307d0d4e1cfffd8b2.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) RegisteredUsersById(id string)(*iec54264a6c486f68435119cb745dbefdc6e215ba637f94299ef3561dd30389aa.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return iec54264a6c486f68435119cb745dbefdc6e215ba637f94299ef3561dd30389aa.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *DeviceItemRequestBuilder) Restore()(*i6b4891e8e524165f041b742d05729e1561eab385718b042b623c5f072b36c6b1.RestoreRequestBuilder) {
    return i6b4891e8e524165f041b742d05729e1561eab385718b042b623c5f072b36c6b1.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) TransitiveMemberOf()(*i2d7c976ed5ceb8455050d02a0603e4daf6148156215fc4056825711cb8d13da3.TransitiveMemberOfRequestBuilder) {
    return i2d7c976ed5ceb8455050d02a0603e4daf6148156215fc4056825711cb8d13da3.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) TransitiveMemberOfById(id string)(*ib92c409c5832dab5202e7c30ad3be7d04087f9f45c98fb4009798c7f34318347.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ib92c409c5832dab5202e7c30ad3be7d04087f9f45c98fb4009798c7f34318347.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) UsageRights()(*i7767512c114fae7428f00ecc379bdb6c40770b44d591d609b62d288e8fb6e1a7.UsageRightsRequestBuilder) {
    return i7767512c114fae7428f00ecc379bdb6c40770b44d591d609b62d288e8fb6e1a7.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *DeviceItemRequestBuilder) UsageRightsById(id string)(*i534ed1cb37dd623a3ea412e67b0dca570cb44b8681f511614639ffd95b6791c1.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return i534ed1cb37dd623a3ea412e67b0dca570cb44b8681f511614639ffd95b6791c1.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
