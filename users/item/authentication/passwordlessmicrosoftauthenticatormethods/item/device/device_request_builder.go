package device

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2bb29e3be96b80c4d1ea202bf9795454e1ac40085fbd8c28dcd549c121f4a2c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/usagerights"
    i2e51ac4f4142590854e942ff543005c994156800001ab732ac27492f0464c529 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/extensions"
    i6c88bea172a45d33e766ba9c9c0c5a9cc98205bcfb3c0848fcee43c54a4d07cb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers"
    i8b40ca361810f5142e8fbad8621b4c7a2abdf1b7211f1b19a3a75dbb588f3a2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceRequestBuilderDeleteOptions options for Delete
type DeviceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceRequestBuilderGetOptions options for Get
type DeviceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceRequestBuilderGetQueryParameters get device from users
type DeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceRequestBuilderPatchOptions options for Patch
type DeviceRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["command_id"] = id
    }
    return i0b3d3eff6f0c692fa600a2cdc0e24410417bd326cff0c598f184fe73be909004.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod_id}/device{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property device for users
func (m *DeviceRequestBuilder) CreateDeleteRequestInformation(options *DeviceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get device from users
func (m *DeviceRequestBuilder) CreateGetRequestInformation(options *DeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property device in users
func (m *DeviceRequestBuilder) CreatePatchRequestInformation(options *DeviceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property device for users
func (m *DeviceRequestBuilder) Delete(options *DeviceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["extension_id"] = id
    }
    return i93aafe8bcd1a528ebb2b886da1bc4674d9896032247de7be28665296f2789ad0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get device from users
func (m *DeviceRequestBuilder) Get(options *DeviceRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable), nil
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i5b3935114692a197c635227b4624459fd562ad55fbd4eba43dfc168dcc1a08c8.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in users
func (m *DeviceRequestBuilder) Patch(options *DeviceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i69d90a742e916ade151e50cbab5a77ffd7bd6ab29ccf3035a60a517e0a0c3baf.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i409d9eed47e75160313c229e4e11041cc4f3fe1e57a6d39f4e6d32d6386cc472.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i146d79045ea17b9287dcf1e2b1127dd2a767666d763a9f93d1093c6122814b2c.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["usageRight_id"] = id
    }
    return i191b3f6554c0035f7a55eef4d87d782ea73ac3729d12d74a885232394ea732a2.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
