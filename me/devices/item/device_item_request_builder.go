package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06de72e85c136548b632475aeec5362fa29354e14a94fe0f6befb3d7d0c7ee08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/extensions"
    i2d7c976ed5ceb8455050d02a0603e4daf6148156215fc4056825711cb8d13da3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceItemRequestBuilderDeleteOptions options for Delete
type DeviceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceItemRequestBuilderGetOptions options for Get
type DeviceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceItemRequestBuilderGetQueryParameters get devices from me
type DeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceItemRequestBuilderPatchOptions options for Patch
type DeviceItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["command_id"] = id
    }
    return i2af1416da365c0340a39d8ecc556e26557cb9f58ce6560792d2198c137dc55b5.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property devices for me
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get devices from me
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation(options *DeviceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property devices in me
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(options *DeviceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property devices for me
func (m *DeviceItemRequestBuilder) Delete(options *DeviceItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["extension_id"] = id
    }
    return i9fd4cc16d3894463aa8887400816cd691765b2f88389be9d15e959a82eb1a7ed.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get devices from me
func (m *DeviceItemRequestBuilder) Get(options *DeviceItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable), nil
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i2e2e998caa42ebe0146db25ae496c4ffcc4120a2786ba8d28386223ee7843679.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property devices in me
func (m *DeviceItemRequestBuilder) Patch(options *DeviceItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i4a299936ca517941a98d45670fb79ba0a2da5a3ca11d12e1c5bbd1579d31d553.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return iec54264a6c486f68435119cb745dbefdc6e215ba637f94299ef3561dd30389aa.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return ib92c409c5832dab5202e7c30ad3be7d04087f9f45c98fb4009798c7f34318347.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["usageRight_id"] = id
    }
    return i534ed1cb37dd623a3ea412e67b0dca570cb44b8681f511614639ffd95b6791c1.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
