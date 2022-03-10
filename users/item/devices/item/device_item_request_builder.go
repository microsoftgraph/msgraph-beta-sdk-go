package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i2429f93f2ffa433727c0848b6c9658e4d3cb648472c4eef955bc41768ebe93a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners"
    i260a199d87e8fc8451c2cebd56f049767d3520abc7911b3280adebafc0b35c51 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof"
    i287d87b2c8c4a2a503f874a4e31a649feb57022d3ab8f9697721d4c249a16bee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers"
    i33ff2b6be6dfa57cc8a51bb78cd90fbe1626d3a8e3b5df364004aa3515144f2a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/usagerights"
    i5cae329fa48ff3f1b1990c9cc80fbbfb85184e94b824d0858c03687f67950543 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof"
    i7e864cabfa7df255cec98e1d19a04a50c2fa51c2eff7feff0b4ecebb21696f14 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/extensions"
    iaf34d7f762b2e74c284c42effd3462cff0e37f08a9f6986b076294804a2449ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/commands"
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
// DeviceItemRequestBuilderGetQueryParameters get devices from users
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
        urlTplParams["command_id"] = id
    }
    return i35eda896d73fbeb4732e5c4720890d6a590710f94bd582a4cff1752828a69af3.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/devices/{device_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property devices for users
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
// CreateGetRequestInformation get devices from users
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
// CreatePatchRequestInformation update the navigation property devices in users
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
// Delete delete navigation property devices for users
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
        urlTplParams["extension_id"] = id
    }
    return i15521158ed5380a593e0fb66e3ea064f39b1577a2ac9b6eda66652739a42c3a7.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get devices from users
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
        urlTplParams["directoryObject_id"] = id
    }
    return i7691315271fdbee9c1364b694cd1da63bf25ef5dbd68a8733beca60509601b78.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property devices in users
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
        urlTplParams["directoryObject_id"] = id
    }
    return ie640d03df2e875a4b8e918f08a8bdd3c0ef7575f5ac82da96969e2d5c6834087.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i14cc2fa3b60c5ce229c289cfacee26146c0b595b94a01d4ba5e8f2bc380db01b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i12d469d8755e67463c108cc666b7d20d9dd4f29509dd9f4401df2c1bf22efaf7.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["usageRight_id"] = id
    }
    return i00aa926b8f1486da2206b774db04d1cd018936bae2dcf0db49162361f332101f.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
