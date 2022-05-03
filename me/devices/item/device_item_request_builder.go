package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06de72e85c136548b632475aeec5362fa29354e14a94fe0f6befb3d7d0c7ee08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/extensions"
    i2d7c976ed5ceb8455050d02a0603e4daf6148156215fc4056825711cb8d13da3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof"
    i7767512c114fae7428f00ecc379bdb6c40770b44d591d609b62d288e8fb6e1a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/usagerights"
    ia09559070bc6c3c9adcb81ec9ee248e1edc481381bdc642dcbc165846ecfb443 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof"
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
// Commands the commands property
func (m *DeviceItemRequestBuilder) Commands()(*ib367c374a26076087ddba4f9b886fc919d9fa57297ddc426513b74650121bfbb.CommandsRequestBuilder) {
    return ib367c374a26076087ddba4f9b886fc919d9fa57297ddc426513b74650121bfbb.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item.commands.item collection
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
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property devices for me
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
// CreateGetRequestInformation get devices from me
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get devices from me
func (m *DeviceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property devices in me
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property devices for me
func (m *DeviceItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property devices for me
func (m *DeviceItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *DeviceItemRequestBuilder) Extensions()(*i06de72e85c136548b632475aeec5362fa29354e14a94fe0f6befb3d7d0c7ee08.ExtensionsRequestBuilder) {
    return i06de72e85c136548b632475aeec5362fa29354e14a94fe0f6befb3d7d0c7ee08.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item.extensions.item collection
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
func (m *DeviceItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get devices from me
func (m *DeviceItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable), nil
}
// MemberOf the memberOf property
func (m *DeviceItemRequestBuilder) MemberOf()(*ia09559070bc6c3c9adcb81ec9ee248e1edc481381bdc642dcbc165846ecfb443.MemberOfRequestBuilder) {
    return ia09559070bc6c3c9adcb81ec9ee248e1edc481381bdc642dcbc165846ecfb443.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item.memberOf.item collection
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
func (m *DeviceItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property devices in me
func (m *DeviceItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RegisteredOwners the registeredOwners property
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
// RegisteredUsers the registeredUsers property
func (m *DeviceItemRequestBuilder) RegisteredUsers()(*ib911ca91b9101c7c48053251c8d96b7673ac91153da66f0307d0d4e1cfffd8b2.RegisteredUsersRequestBuilder) {
    return ib911ca91b9101c7c48053251c8d96b7673ac91153da66f0307d0d4e1cfffd8b2.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item.registeredUsers.item collection
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
// TransitiveMemberOf the transitiveMemberOf property
func (m *DeviceItemRequestBuilder) TransitiveMemberOf()(*i2d7c976ed5ceb8455050d02a0603e4daf6148156215fc4056825711cb8d13da3.TransitiveMemberOfRequestBuilder) {
    return i2d7c976ed5ceb8455050d02a0603e4daf6148156215fc4056825711cb8d13da3.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item.transitiveMemberOf.item collection
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
// UsageRights the usageRights property
func (m *DeviceItemRequestBuilder) UsageRights()(*i7767512c114fae7428f00ecc379bdb6c40770b44d591d609b62d288e8fb6e1a7.UsageRightsRequestBuilder) {
    return i7767512c114fae7428f00ecc379bdb6c40770b44d591d609b62d288e8fb6e1a7.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item.usageRights.item collection
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
