package device

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i19975c5ccfea112c52d6d181777f13c8465aba6d0b7abc183efab7063bc48c56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/extensions"
    i6a8737686a34b4bf2117f31e88734dffc4c31ccb312e8e8647a6abf876879cc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/usagerights"
    i796697e2d60bc942c0b41a4c781a3cec80275b0cec3a40ccef37c1d9538388b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof"
    i88f1e34d106b4b084f22e25197c92bcb77f352a1065577ff8926e20bc4fc0677 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/commands"
    i902bc40ab212f60615c3c0bed5af46e74f733cfaafb6cd3105585920eb8b5e7c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers"
    ib6fe11c4faa1d614e1cfd3d3e2f00ac4ae81e4b439fa73188fc62389f662cf1a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners"
    ifa5283eea5f8de4c362462e01201478ec482f97f73da616b2cd14e7daf50c99f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof"
    i076960593b764589d6a5b4f4b6a295b537db0d9bf5f46827c387e41e567609fb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item"
    i1ab4ad1285fd331d3179729a55e1293de9d19378fbb0532a4fc048df19b5f167 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item"
    ia20b2ae4d062077fb0f76f3e57dae4f07e0f9aa61d47898b9604e0b0755f4702 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/commands/item"
    ib37b449cd061e26a1707349f206eb2c2c3d83cbbf76e95fdded036ed54465ba4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item"
    ic7eb226dc87fa339549a18bc60ae70ca999ff13bf2865fade2ca1b337488bc98 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item"
    icad7d54ff44ff8aab9b025ac0bd30b1a3f78556995d4c769d96f4d72c7272ba3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/usagerights/item"
    ieb1bc5ca9975b15b7e9659b747ab5ef364b6ef4d57a2ced7d46cac26e305f5ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/extensions/item"
)

// DeviceRequestBuilder provides operations to manage the device property of the microsoft.graph.microsoftAuthenticatorAuthenticationMethod entity.
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
// DeviceRequestBuilderGetQueryParameters the registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
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
func (m *DeviceRequestBuilder) Commands()(*i88f1e34d106b4b084f22e25197c92bcb77f352a1065577ff8926e20bc4fc0677.CommandsRequestBuilder) {
    return i88f1e34d106b4b084f22e25197c92bcb77f352a1065577ff8926e20bc4fc0677.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item.device.commands.item collection
func (m *DeviceRequestBuilder) CommandsById(id string)(*ia20b2ae4d062077fb0f76f3e57dae4f07e0f9aa61d47898b9604e0b0755f4702.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command_id"] = id
    }
    return ia20b2ae4d062077fb0f76f3e57dae4f07e0f9aa61d47898b9604e0b0755f4702.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod_id}/device{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property device for me
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
// CreateGetRequestInformation the registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
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
// CreatePatchRequestInformation update the navigation property device in me
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
// Delete delete navigation property device for me
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
func (m *DeviceRequestBuilder) Extensions()(*i19975c5ccfea112c52d6d181777f13c8465aba6d0b7abc183efab7063bc48c56.ExtensionsRequestBuilder) {
    return i19975c5ccfea112c52d6d181777f13c8465aba6d0b7abc183efab7063bc48c56.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item.device.extensions.item collection
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*ieb1bc5ca9975b15b7e9659b747ab5ef364b6ef4d57a2ced7d46cac26e305f5ec.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ieb1bc5ca9975b15b7e9659b747ab5ef364b6ef4d57a2ced7d46cac26e305f5ec.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
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
func (m *DeviceRequestBuilder) MemberOf()(*ifa5283eea5f8de4c362462e01201478ec482f97f73da616b2cd14e7daf50c99f.MemberOfRequestBuilder) {
    return ifa5283eea5f8de4c362462e01201478ec482f97f73da616b2cd14e7daf50c99f.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item.device.memberOf.item collection
func (m *DeviceRequestBuilder) MemberOfById(id string)(*ib37b449cd061e26a1707349f206eb2c2c3d83cbbf76e95fdded036ed54465ba4.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ib37b449cd061e26a1707349f206eb2c2c3d83cbbf76e95fdded036ed54465ba4.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property device in me
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
func (m *DeviceRequestBuilder) RegisteredOwners()(*ib6fe11c4faa1d614e1cfd3d3e2f00ac4ae81e4b439fa73188fc62389f662cf1a.RegisteredOwnersRequestBuilder) {
    return ib6fe11c4faa1d614e1cfd3d3e2f00ac4ae81e4b439fa73188fc62389f662cf1a.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item.device.registeredOwners.item collection
func (m *DeviceRequestBuilder) RegisteredOwnersById(id string)(*i076960593b764589d6a5b4f4b6a295b537db0d9bf5f46827c387e41e567609fb.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i076960593b764589d6a5b4f4b6a295b537db0d9bf5f46827c387e41e567609fb.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceRequestBuilder) RegisteredUsers()(*i902bc40ab212f60615c3c0bed5af46e74f733cfaafb6cd3105585920eb8b5e7c.RegisteredUsersRequestBuilder) {
    return i902bc40ab212f60615c3c0bed5af46e74f733cfaafb6cd3105585920eb8b5e7c.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item.device.registeredUsers.item collection
func (m *DeviceRequestBuilder) RegisteredUsersById(id string)(*ic7eb226dc87fa339549a18bc60ae70ca999ff13bf2865fade2ca1b337488bc98.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ic7eb226dc87fa339549a18bc60ae70ca999ff13bf2865fade2ca1b337488bc98.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*i796697e2d60bc942c0b41a4c781a3cec80275b0cec3a40ccef37c1d9538388b9.TransitiveMemberOfRequestBuilder) {
    return i796697e2d60bc942c0b41a4c781a3cec80275b0cec3a40ccef37c1d9538388b9.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item.device.transitiveMemberOf.item collection
func (m *DeviceRequestBuilder) TransitiveMemberOfById(id string)(*i1ab4ad1285fd331d3179729a55e1293de9d19378fbb0532a4fc048df19b5f167.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i1ab4ad1285fd331d3179729a55e1293de9d19378fbb0532a4fc048df19b5f167.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceRequestBuilder) UsageRights()(*i6a8737686a34b4bf2117f31e88734dffc4c31ccb312e8e8647a6abf876879cc2.UsageRightsRequestBuilder) {
    return i6a8737686a34b4bf2117f31e88734dffc4c31ccb312e8e8647a6abf876879cc2.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item.device.usageRights.item collection
func (m *DeviceRequestBuilder) UsageRightsById(id string)(*icad7d54ff44ff8aab9b025ac0bd30b1a3f78556995d4c769d96f4d72c7272ba3.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight_id"] = id
    }
    return icad7d54ff44ff8aab9b025ac0bd30b1a3f78556995d4c769d96f4d72c7272ba3.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
