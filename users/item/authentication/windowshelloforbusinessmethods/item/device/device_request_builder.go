package device

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
// DeviceRequestBuilderGetQueryParameters the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
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
        urlTplParams["command_id"] = id
    }
    return i5680acfb979cbcbc9bf369845d5e7e99f9f26e626efd593414fc1a6d59bd14be.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod_id}/device{?select,expand}";
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
// CreateGetRequestInformation the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
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
        urlTplParams["extension_id"] = id
    }
    return i1605d0057c653186fcacf38730d68a93a064ce091805e5e1d4c6d378eefac6fa.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business registration information, this property is returned only on a single GET and when you specify ?$expand. For example, GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
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
        urlTplParams["directoryObject_id"] = id
    }
    return i6ab82c3fea9de45896b50e6171070905c9f7a8a1606a82a068c3254a4988d1b8.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
        urlTplParams["directoryObject_id"] = id
    }
    return iae7d2d0a647de4e274a820553b74075513948fb0331ef68df755d319366e359d.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return ibbf9f71f42ad003261e0ab32bc144aa7595cb52793e32a38b56cd7580f889e7b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i65f3107e1febe2b1e7e8b92d84172f33c47da6831c1484a1292d615876fa1ed1.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["usageRight_id"] = id
    }
    return ibdf383b4b0b4f963197b019e7e819b43b6bc49f66b4328eaa2f01436357e24f0.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
