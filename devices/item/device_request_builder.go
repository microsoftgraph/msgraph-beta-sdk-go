package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/getmemberobjects"
    i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof"
    i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/checkmembergroups"
    i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/commands"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/getmembergroups"
    i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof"
    ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/checkmemberobjects"
    ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/restore"
    id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredusers"
    ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/extensions"
    iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners"
    if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/usagerights"
    i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/usagerights/item"
    ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/commands/item"
    if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/extensions/item"
)

type DeviceRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DeviceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *DeviceRequestBuilder) CheckMemberGroups()(*i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862.CheckMemberGroupsRequestBuilder) {
    return i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) CheckMemberObjects()(*ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06.CheckMemberObjectsRequestBuilder) {
    return ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) Commands()(*i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4.CommandsRequestBuilder) {
    return i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) CommandsById(id string)(*ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12.CommandRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command_id"] = id
    }
    return ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12.NewCommandRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/devices/{device_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DeviceRequestBuilder) CreateGetRequestInformation(q func (value *DeviceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DeviceRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DeviceRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Device, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DeviceRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceRequestBuilder) Extensions()(*ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc.ExtensionsRequestBuilder) {
    return ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceRequestBuilder) Get(q func (value *DeviceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Device, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDevice() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Device), nil
}
func (m *DeviceRequestBuilder) GetMemberGroups()(*i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4.GetMemberGroupsRequestBuilder) {
    return i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) GetMemberObjects()(*i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3.GetMemberObjectsRequestBuilder) {
    return i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) MemberOf()(*i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1.MemberOfRequestBuilder) {
    return i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Device, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceRequestBuilder) RegisteredOwners()(*iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043.RegisteredOwnersRequestBuilder) {
    return iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) RegisteredUsers()(*id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93.RegisteredUsersRequestBuilder) {
    return id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) Restore()(*ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc.RestoreRequestBuilder) {
    return ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc.TransitiveMemberOfRequestBuilder) {
    return i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) UsageRights()(*if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b.UsageRightsRequestBuilder) {
    return if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) UsageRightsById(id string)(*i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75.UsageRightRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight_id"] = id
    }
    return i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75.NewUsageRightRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
