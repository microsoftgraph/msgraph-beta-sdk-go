package device

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1015f9bb3ad4d71ca425daf115686bfdd695ab4c0381810b29025a8ec114cb0f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/usagerights"
    i2de41a35f0b2e2b743cfa637e4f8e1a0bf1a83a3844e919595dede2b2cda69f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof"
    i5a3f3491cada69786d8dc6347851ed9ae24c7af51641700ae2c11483dd059a93 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/commands"
    i96c91677b99acc9de182d98c2a17872d33edf94ec51d7250b0ce8e483055d658 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/extensions"
    ia53c025dd72ce120037bb81a9d91708b696d4f0b0db965fe5ee1c3fa3c4eb07e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners"
    ibdb584a249724bb612d8ee007d9f757c2e10cd0930784701089c80b4fa4b6ad2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers"
    ie794b487ccd4116fc57d6e29138be7ff04bff3efcca3b370a2b8993414357c09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof"
    i1605d0057c653186fcacf38730d68a93a064ce091805e5e1d4c6d378eefac6fa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/extensions/item"
    i5680acfb979cbcbc9bf369845d5e7e99f9f26e626efd593414fc1a6d59bd14be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/commands/item"
    i65f3107e1febe2b1e7e8b92d84172f33c47da6831c1484a1292d615876fa1ed1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item"
    i6ab82c3fea9de45896b50e6171070905c9f7a8a1606a82a068c3254a4988d1b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item"
    iae7d2d0a647de4e274a820553b74075513948fb0331ef68df755d319366e359d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item"
    ibbf9f71f42ad003261e0ab32bc144aa7595cb52793e32a38b56cd7580f889e7b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item"
    ibdf383b4b0b4f963197b019e7e819b43b6bc49f66b4328eaa2f01436357e24f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/usagerights/item"
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
// Commands the commands property
func (m *DeviceRequestBuilder) Commands()(*i5a3f3491cada69786d8dc6347851ed9ae24c7af51641700ae2c11483dd059a93.CommandsRequestBuilder) {
    return i5a3f3491cada69786d8dc6347851ed9ae24c7af51641700ae2c11483dd059a93.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.commands.item collection
func (m *DeviceRequestBuilder) CommandsById(id string)(*i5680acfb979cbcbc9bf369845d5e7e99f9f26e626efd593414fc1a6d59bd14be.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return i5680acfb979cbcbc9bf369845d5e7e99f9f26e626efd593414fc1a6d59bd14be.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property device for users
func (m *DeviceRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property device in users
func (m *DeviceRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property device for users
func (m *DeviceRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DeviceRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property device for users
func (m *DeviceRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DeviceRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *DeviceRequestBuilder) Extensions()(*i96c91677b99acc9de182d98c2a17872d33edf94ec51d7250b0ce8e483055d658.ExtensionsRequestBuilder) {
    return i96c91677b99acc9de182d98c2a17872d33edf94ec51d7250b0ce8e483055d658.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.extensions.item collection
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*i1605d0057c653186fcacf38730d68a93a064ce091805e5e1d4c6d378eefac6fa.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i1605d0057c653186fcacf38730d68a93a064ce091805e5e1d4c6d378eefac6fa.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetWithResponseHandler the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (m *DeviceRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
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
func (m *DeviceRequestBuilder) MemberOf()(*ie794b487ccd4116fc57d6e29138be7ff04bff3efcca3b370a2b8993414357c09.MemberOfRequestBuilder) {
    return ie794b487ccd4116fc57d6e29138be7ff04bff3efcca3b370a2b8993414357c09.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.memberOf.item collection
func (m *DeviceRequestBuilder) MemberOfById(id string)(*i6ab82c3fea9de45896b50e6171070905c9f7a8a1606a82a068c3254a4988d1b8.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i6ab82c3fea9de45896b50e6171070905c9f7a8a1606a82a068c3254a4988d1b8.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property device in users
func (m *DeviceRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property device in users
func (m *DeviceRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *DeviceRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *DeviceRequestBuilder) RegisteredOwners()(*ia53c025dd72ce120037bb81a9d91708b696d4f0b0db965fe5ee1c3fa3c4eb07e.RegisteredOwnersRequestBuilder) {
    return ia53c025dd72ce120037bb81a9d91708b696d4f0b0db965fe5ee1c3fa3c4eb07e.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.registeredOwners.item collection
func (m *DeviceRequestBuilder) RegisteredOwnersById(id string)(*iae7d2d0a647de4e274a820553b74075513948fb0331ef68df755d319366e359d.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return iae7d2d0a647de4e274a820553b74075513948fb0331ef68df755d319366e359d.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers the registeredUsers property
func (m *DeviceRequestBuilder) RegisteredUsers()(*ibdb584a249724bb612d8ee007d9f757c2e10cd0930784701089c80b4fa4b6ad2.RegisteredUsersRequestBuilder) {
    return ibdb584a249724bb612d8ee007d9f757c2e10cd0930784701089c80b4fa4b6ad2.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.registeredUsers.item collection
func (m *DeviceRequestBuilder) RegisteredUsersById(id string)(*ibbf9f71f42ad003261e0ab32bc144aa7595cb52793e32a38b56cd7580f889e7b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ibbf9f71f42ad003261e0ab32bc144aa7595cb52793e32a38b56cd7580f889e7b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*i2de41a35f0b2e2b743cfa637e4f8e1a0bf1a83a3844e919595dede2b2cda69f3.TransitiveMemberOfRequestBuilder) {
    return i2de41a35f0b2e2b743cfa637e4f8e1a0bf1a83a3844e919595dede2b2cda69f3.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.transitiveMemberOf.item collection
func (m *DeviceRequestBuilder) TransitiveMemberOfById(id string)(*i65f3107e1febe2b1e7e8b92d84172f33c47da6831c1484a1292d615876fa1ed1.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i65f3107e1febe2b1e7e8b92d84172f33c47da6831c1484a1292d615876fa1ed1.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsageRights the usageRights property
func (m *DeviceRequestBuilder) UsageRights()(*i1015f9bb3ad4d71ca425daf115686bfdd695ab4c0381810b29025a8ec114cb0f.UsageRightsRequestBuilder) {
    return i1015f9bb3ad4d71ca425daf115686bfdd695ab4c0381810b29025a8ec114cb0f.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item.device.usageRights.item collection
func (m *DeviceRequestBuilder) UsageRightsById(id string)(*ibdf383b4b0b4f963197b019e7e819b43b6bc49f66b4328eaa2f01436357e24f0.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return ibdf383b4b0b4f963197b019e7e819b43b6bc49f66b4328eaa2f01436357e24f0.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
