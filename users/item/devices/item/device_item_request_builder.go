package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i17c4436168820ed9465c6ea216b9f889b64a2e73f84cc7d3493825322c62bca3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/getmembergroups"
    i2429f93f2ffa433727c0848b6c9658e4d3cb648472c4eef955bc41768ebe93a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners"
    i260a199d87e8fc8451c2cebd56f049767d3520abc7911b3280adebafc0b35c51 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof"
    i287d87b2c8c4a2a503f874a4e31a649feb57022d3ab8f9697721d4c249a16bee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers"
    i33ff2b6be6dfa57cc8a51bb78cd90fbe1626d3a8e3b5df364004aa3515144f2a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/usagerights"
    i5cae329fa48ff3f1b1990c9cc80fbbfb85184e94b824d0858c03687f67950543 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof"
    i6c404d58392755d5bc2022900737e6eb0513dd6c6ede8da370f111d117cc4ffa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/checkmembergroups"
    i7e273b78e0c1827e680be4fa377056e00c9023d9db253f52dc7635e67462f2f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/getmemberobjects"
    i7e864cabfa7df255cec98e1d19a04a50c2fa51c2eff7feff0b4ecebb21696f14 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/extensions"
    i9cf2733b2fc19d0e3fcfbca2e47e90f59f781c75f0740d99d06d83a617c2cdb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/checkmemberobjects"
    iaf34d7f762b2e74c284c42effd3462cff0e37f08a9f6986b076294804a2449ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/commands"
    ib9741e8fd6f934777f637ee867cb5ea268b06ade444d82ba5c33f83a0a508199 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/restore"
    i00aa926b8f1486da2206b774db04d1cd018936bae2dcf0db49162361f332101f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/usagerights/item"
    i12d469d8755e67463c108cc666b7d20d9dd4f29509dd9f4401df2c1bf22efaf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item"
    i14cc2fa3b60c5ce229c289cfacee26146c0b595b94a01d4ba5e8f2bc380db01b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item"
    i15521158ed5380a593e0fb66e3ea064f39b1577a2ac9b6eda66652739a42c3a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/extensions/item"
    i35eda896d73fbeb4732e5c4720890d6a590710f94bd582a4cff1752828a69af3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/commands/item"
    i7691315271fdbee9c1364b694cd1da63bf25ef5dbd68a8733beca60509601b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item"
    ie640d03df2e875a4b8e918f08a8bdd3c0ef7575f5ac82da96969e2d5c6834087 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item"
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
// DeviceItemRequestBuilderGetQueryParameters get devices from users
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
func (m *DeviceItemRequestBuilder) CheckMemberGroups()(*i6c404d58392755d5bc2022900737e6eb0513dd6c6ede8da370f111d117cc4ffa.CheckMemberGroupsRequestBuilder) {
    return i6c404d58392755d5bc2022900737e6eb0513dd6c6ede8da370f111d117cc4ffa.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DeviceItemRequestBuilder) CheckMemberObjects()(*i9cf2733b2fc19d0e3fcfbca2e47e90f59f781c75f0740d99d06d83a617c2cdb1.CheckMemberObjectsRequestBuilder) {
    return i9cf2733b2fc19d0e3fcfbca2e47e90f59f781c75f0740d99d06d83a617c2cdb1.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Commands the commands property
func (m *DeviceItemRequestBuilder) Commands()(*iaf34d7f762b2e74c284c42effd3462cff0e37f08a9f6986b076294804a2449ef.CommandsRequestBuilder) {
    return iaf34d7f762b2e74c284c42effd3462cff0e37f08a9f6986b076294804a2449ef.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.commands.item collection
func (m *DeviceItemRequestBuilder) CommandsById(id string)(*i35eda896d73fbeb4732e5c4720890d6a590710f94bd582a4cff1752828a69af3.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return i35eda896d73fbeb4732e5c4720890d6a590710f94bd582a4cff1752828a69af3.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
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
// NewDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property devices for users
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property devices for users
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
// CreateGetRequestInformation get devices from users
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get devices from users
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
// CreatePatchRequestInformation update the navigation property devices in users
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property devices in users
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
// Delete delete navigation property devices for users
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
func (m *DeviceItemRequestBuilder) Extensions()(*i7e864cabfa7df255cec98e1d19a04a50c2fa51c2eff7feff0b4ecebb21696f14.ExtensionsRequestBuilder) {
    return i7e864cabfa7df255cec98e1d19a04a50c2fa51c2eff7feff0b4ecebb21696f14.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.extensions.item collection
func (m *DeviceItemRequestBuilder) ExtensionsById(id string)(*i15521158ed5380a593e0fb66e3ea064f39b1577a2ac9b6eda66652739a42c3a7.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i15521158ed5380a593e0fb66e3ea064f39b1577a2ac9b6eda66652739a42c3a7.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get devices from users
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
func (m *DeviceItemRequestBuilder) GetMemberGroups()(*i17c4436168820ed9465c6ea216b9f889b64a2e73f84cc7d3493825322c62bca3.GetMemberGroupsRequestBuilder) {
    return i17c4436168820ed9465c6ea216b9f889b64a2e73f84cc7d3493825322c62bca3.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DeviceItemRequestBuilder) GetMemberObjects()(*i7e273b78e0c1827e680be4fa377056e00c9023d9db253f52dc7635e67462f2f3.GetMemberObjectsRequestBuilder) {
    return i7e273b78e0c1827e680be4fa377056e00c9023d9db253f52dc7635e67462f2f3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *DeviceItemRequestBuilder) MemberOf()(*i260a199d87e8fc8451c2cebd56f049767d3520abc7911b3280adebafc0b35c51.MemberOfRequestBuilder) {
    return i260a199d87e8fc8451c2cebd56f049767d3520abc7911b3280adebafc0b35c51.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.memberOf.item collection
func (m *DeviceItemRequestBuilder) MemberOfById(id string)(*i7691315271fdbee9c1364b694cd1da63bf25ef5dbd68a8733beca60509601b78.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i7691315271fdbee9c1364b694cd1da63bf25ef5dbd68a8733beca60509601b78.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property devices in users
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
func (m *DeviceItemRequestBuilder) RegisteredOwners()(*i2429f93f2ffa433727c0848b6c9658e4d3cb648472c4eef955bc41768ebe93a3.RegisteredOwnersRequestBuilder) {
    return i2429f93f2ffa433727c0848b6c9658e4d3cb648472c4eef955bc41768ebe93a3.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.registeredOwners.item collection
func (m *DeviceItemRequestBuilder) RegisteredOwnersById(id string)(*ie640d03df2e875a4b8e918f08a8bdd3c0ef7575f5ac82da96969e2d5c6834087.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ie640d03df2e875a4b8e918f08a8bdd3c0ef7575f5ac82da96969e2d5c6834087.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers the registeredUsers property
func (m *DeviceItemRequestBuilder) RegisteredUsers()(*i287d87b2c8c4a2a503f874a4e31a649feb57022d3ab8f9697721d4c249a16bee.RegisteredUsersRequestBuilder) {
    return i287d87b2c8c4a2a503f874a4e31a649feb57022d3ab8f9697721d4c249a16bee.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.registeredUsers.item collection
func (m *DeviceItemRequestBuilder) RegisteredUsersById(id string)(*i14cc2fa3b60c5ce229c289cfacee26146c0b595b94a01d4ba5e8f2bc380db01b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i14cc2fa3b60c5ce229c289cfacee26146c0b595b94a01d4ba5e8f2bc380db01b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore the restore property
func (m *DeviceItemRequestBuilder) Restore()(*ib9741e8fd6f934777f637ee867cb5ea268b06ade444d82ba5c33f83a0a508199.RestoreRequestBuilder) {
    return ib9741e8fd6f934777f637ee867cb5ea268b06ade444d82ba5c33f83a0a508199.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *DeviceItemRequestBuilder) TransitiveMemberOf()(*i5cae329fa48ff3f1b1990c9cc80fbbfb85184e94b824d0858c03687f67950543.TransitiveMemberOfRequestBuilder) {
    return i5cae329fa48ff3f1b1990c9cc80fbbfb85184e94b824d0858c03687f67950543.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.transitiveMemberOf.item collection
func (m *DeviceItemRequestBuilder) TransitiveMemberOfById(id string)(*i12d469d8755e67463c108cc666b7d20d9dd4f29509dd9f4401df2c1bf22efaf7.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i12d469d8755e67463c108cc666b7d20d9dd4f29509dd9f4401df2c1bf22efaf7.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights the usageRights property
func (m *DeviceItemRequestBuilder) UsageRights()(*i33ff2b6be6dfa57cc8a51bb78cd90fbe1626d3a8e3b5df364004aa3515144f2a.UsageRightsRequestBuilder) {
    return i33ff2b6be6dfa57cc8a51bb78cd90fbe1626d3a8e3b5df364004aa3515144f2a.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.usageRights.item collection
func (m *DeviceItemRequestBuilder) UsageRightsById(id string)(*i00aa926b8f1486da2206b774db04d1cd018936bae2dcf0db49162361f332101f.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return i00aa926b8f1486da2206b774db04d1cd018936bae2dcf0db49162361f332101f.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
