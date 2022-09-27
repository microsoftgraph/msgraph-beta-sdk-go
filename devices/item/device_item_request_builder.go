package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/getmemberobjects"
    i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof"
    i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/checkmembergroups"
    i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/commands"
    i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/getmembergroups"
    i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof"
    ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/checkmemberobjects"
    ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/restore"
    id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredusers"
    ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/extensions"
    iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners"
    if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/usagerights"
    i17b2e41daecab12778ef8849a52a0050afea7ab215b7a063c8d145b656ff350c "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof/item"
    i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/usagerights/item"
    ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/commands/item"
    ic5b79b8426b0398a1d37364aa2476ccd846f17a8e2130477d3235a55b6728c2f "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/item"
    id7b0a8873e3dc9a1070cd7d3a137f0ed2968cdf066a85a0810adb726fef41b0e "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredusers/item"
    if0991c114c37c9c6d2fde813890c596fa6c31a4b084d49dd6eb0e8c33c7f8af6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof/item"
    if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/extensions/item"
)

// DeviceItemRequestBuilder provides operations to manage the collection of device entities.
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
// DeviceItemRequestBuilderGetQueryParameters get the properties and relationships of a device object. Since the **device** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **device** instance.
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
// CheckMemberGroups the checkMemberGroups property
func (m *DeviceItemRequestBuilder) CheckMemberGroups()(*i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862.CheckMemberGroupsRequestBuilder) {
    return i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DeviceItemRequestBuilder) CheckMemberObjects()(*ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06.CheckMemberObjectsRequestBuilder) {
    return ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commands the commands property
func (m *DeviceItemRequestBuilder) Commands()(*i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4.CommandsRequestBuilder) {
    return i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.commands.item collection
func (m *DeviceItemRequestBuilder) CommandsById(id string)(*ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete a registered device.
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete a registered device.
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get the properties and relationships of a device object. Since the **device** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **device** instance.
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the properties and relationships of a device object. Since the **device** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **device** instance.
func (m *DeviceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a device. Only certain properties of a device can be updated through approved Mobile Device Management (MDM) apps.
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of a device. Only certain properties of a device can be updated through approved Mobile Device Management (MDM) apps.
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete a registered device.
func (m *DeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceItemRequestBuilder) Extensions()(*ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc.ExtensionsRequestBuilder) {
    return ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.extensions.item collection
func (m *DeviceItemRequestBuilder) ExtensionsById(id string)(*if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of a device object. Since the **device** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **device** instance.
func (m *DeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *DeviceItemRequestBuilder) GetMemberGroups()(*i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4.GetMemberGroupsRequestBuilder) {
    return i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DeviceItemRequestBuilder) GetMemberObjects()(*i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3.GetMemberObjectsRequestBuilder) {
    return i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *DeviceItemRequestBuilder) MemberOf()(*i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1.MemberOfRequestBuilder) {
    return i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.memberOf.item collection
func (m *DeviceItemRequestBuilder) MemberOfById(id string)(*if0991c114c37c9c6d2fde813890c596fa6c31a4b084d49dd6eb0e8c33c7f8af6.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return if0991c114c37c9c6d2fde813890c596fa6c31a4b084d49dd6eb0e8c33c7f8af6.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a device. Only certain properties of a device can be updated through approved Mobile Device Management (MDM) apps.
func (m *DeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// RegisteredOwners the registeredOwners property
func (m *DeviceItemRequestBuilder) RegisteredOwners()(*iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043.RegisteredOwnersRequestBuilder) {
    return iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.registeredOwners.item collection
func (m *DeviceItemRequestBuilder) RegisteredOwnersById(id string)(*ic5b79b8426b0398a1d37364aa2476ccd846f17a8e2130477d3235a55b6728c2f.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ic5b79b8426b0398a1d37364aa2476ccd846f17a8e2130477d3235a55b6728c2f.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers the registeredUsers property
func (m *DeviceItemRequestBuilder) RegisteredUsers()(*id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93.RegisteredUsersRequestBuilder) {
    return id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.registeredUsers.item collection
func (m *DeviceItemRequestBuilder) RegisteredUsersById(id string)(*id7b0a8873e3dc9a1070cd7d3a137f0ed2968cdf066a85a0810adb726fef41b0e.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return id7b0a8873e3dc9a1070cd7d3a137f0ed2968cdf066a85a0810adb726fef41b0e.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore the restore property
func (m *DeviceItemRequestBuilder) Restore()(*ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc.RestoreRequestBuilder) {
    return ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *DeviceItemRequestBuilder) TransitiveMemberOf()(*i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc.TransitiveMemberOfRequestBuilder) {
    return i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.transitiveMemberOf.item collection
func (m *DeviceItemRequestBuilder) TransitiveMemberOfById(id string)(*i17b2e41daecab12778ef8849a52a0050afea7ab215b7a063c8d145b656ff350c.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i17b2e41daecab12778ef8849a52a0050afea7ab215b7a063c8d145b656ff350c.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights the usageRights property
func (m *DeviceItemRequestBuilder) UsageRights()(*if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b.UsageRightsRequestBuilder) {
    return if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.usageRights.item collection
func (m *DeviceItemRequestBuilder) UsageRightsById(id string)(*i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
