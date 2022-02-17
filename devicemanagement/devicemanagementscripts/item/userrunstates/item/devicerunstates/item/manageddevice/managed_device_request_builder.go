package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i10f6744576cd776bd8de8f12ebe8cf8615566bc5f7fd8d948cf87e464fc18d72 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/ref"
    i2309f9a1593bcd546d2a5f3758b4196ff376a243bc8570f03a40650eab948658 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/resetpasscode"
    i29ef92215c3a0570d77bfa8d4a848729136fe5095e3b2ade06426f87b16a66d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/activatedeviceesim"
    i2b80899a960ad28128244d2535f97fb532457b869a5658df8ab8b416c8dc19b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/bypassactivationlock"
    i399d6e7104be0617d1cc28c6cbdfa6e39a10a68e036e7ec466bdee5b3ee9d4eb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/reenable"
    i3ec118cd2643856d79f1ef7abec50a6a052b3e040a245955b4c3a6a94d259456 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    i403dcba80f2cb4eda556b9fa4a77cddfcf327cdc1c57609d2062227c8b2e278b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/shutdown"
    i471340460219a9537a821de60517e9a1c787386de59c0c6f948ba06021e19dd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    i48615277135454168b173ab20d530a2f5a8a815e1b22cdb1ac56be6b3193e91c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/disable"
    i4beea8a0fcd6812c6b72927302ba2bae142a6b88d90d27270dc33cf655870fa4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/remotelock"
    i55d4b94e396bad9a76c0aa71af0ceab7719fbf24e6e97fa08040654aeb7839e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/playlostmodesound"
    i58c250dea6df053eb964e467255d843621c0227eebc40b0fd4d14d09da2bb507 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    i5f850d24a8ec95012aa5a4b161e3d0df8fca997d498138f9635bdbbda18d8228 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/retire"
    i64c3fef8688ebbcc66e46321f822049c485d3ee4c6139420125e9c69ede03466 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rebootnow"
    i6a27f8ff131a4abb5a48228bf7f0ac1085eccc06470057ffd2500e4b0a2484ee "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
    i7470f2b18ba33a6d81ee8adcba65f6bf7ae7a097d9f62a288459325699d268d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    i775ef06d38249f8802deaeaf95042896ed2421ca137aedd4d321b0433174b436 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/recoverpasscode"
    i7951ff84c03d488d41d5d718ae008f4bf55d7098ecc6cfda0cda277d7fdf1011 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/restorecloudpc"
    i7a68baac69887e8af4217c1d694a079001669a4df37190ec0483cbd34d6c1d91 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/setdevicename"
    i7e5e1a05c1297f2cd782c5c9ba54f7379fe2f714aaa62928cb7abe0124bb604c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getfilevaultkey"
    i8d8602658b2b07bda45b9be7c06ea033c1b2cfe38db5493e1fbc30d74d7c4620 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    i9eb1267c62444a3cc384a52459a534effcdd1288be8d163fa543fe99ddb6b79d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    i9f9c7c224fe319efce0b62161dd132d15c29815963df29526040e88fdee29c00 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/wipe"
    ia0939795c41941bda0750587af44dca064a923b4d54b34f571af6613831be401 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    ib8f2523d2206a7faaf8a92a41ed28323e6de488009f295cc76512549fd84057b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getoemwarranty"
    ic9d055cb3975a30afb90486305e9dfc69b8bb249d194703560d49f94224c553f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    icbbc184d6edecea06de2dbcc68e352f1879de8a9ede1f4ceb144f3c0dc570ca9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
    id19190b7e652bf91a3888697decd8fc5fcaf97c1df9c8d9bb7a130cc29879ac2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    id3340c49abb5b01cc6459703996e6066719a8089c162cf75db81e2da32a3b6a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    ida38c086cbf3309ea906f75f123d0b968f83bb97586c4577d7721535a9ea5319 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/overridecompliancestate"
    idaf153750a518c69544d2691ac5f436b715b42484cea2329e43ca5fbf09af25c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/syncdevice"
    idd0937487d4e7865ef09bf471120f20fed975cc48f74c411d675787d55cb7941 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/enablelostmode"
    iddb02630e8f24f7ed576d8a4d4626288d9610749ffd3bb58b6f1d5473fbc4448 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    iddc9cf0a62dad1d6282fde3ff766395f746dbda191d9ad7efb42d0f473a97ad4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    ie33b8d4c95cc420ae6ae59783b82519bc36e6aef4d9738e3a0a3fc7779d14d6e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/locatedevice"
    ie4652856a45d8649b3e905653bf3242fe78b6e9af75119aedcc10f4adbcc4c43 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/requestremoteassistance"
    ie98eef0691d9bc52473f9f7e7b54f6908a1335125f5a99d5b4ea41d31f5e293b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/resizecloudpc"
    ieba57c35679d54bc9735a0392f4997ffb76b5d59c981b5d3330e5bfcf8c2ebfc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    iec225655450ef7e6ba19746507d675be5cf1fe483cd29f92acd05e10aaa37928 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/disablelostmode"
    iffb2299353bba8ab51f78fd2f6a3378d03c06b782e233ffd249b57862b465ae1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item/devicerunstates/item/manageddevice/deprovision"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice
type ManagedDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceRequestBuilderGetOptions options for Get
type ManagedDeviceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceRequestBuilderGetQueryParameters the managed devices that executes the device management script.
type ManagedDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*i29ef92215c3a0570d77bfa8d4a848729136fe5095e3b2ade06426f87b16a66d3.ActivateDeviceEsimRequestBuilder) {
    return i29ef92215c3a0570d77bfa8d4a848729136fe5095e3b2ade06426f87b16a66d3.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*i2b80899a960ad28128244d2535f97fb532457b869a5658df8ab8b416c8dc19b4.BypassActivationLockRequestBuilder) {
    return i2b80899a960ad28128244d2535f97fb532457b869a5658df8ab8b416c8dc19b4.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i3ec118cd2643856d79f1ef7abec50a6a052b3e040a245955b4c3a6a94d259456.CleanWindowsDeviceRequestBuilder) {
    return i3ec118cd2643856d79f1ef7abec50a6a052b3e040a245955b4c3a6a94d259456.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript_id}/userRunStates/{deviceManagementScriptUserState_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceRequestBuilder instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*ic9d055cb3975a30afb90486305e9dfc69b8bb249d194703560d49f94224c553f.CreateDeviceLogCollectionRequestRequestBuilder) {
    return ic9d055cb3975a30afb90486305e9dfc69b8bb249d194703560d49f94224c553f.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed devices that executes the device management script.
func (m *ManagedDeviceRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*i8d8602658b2b07bda45b9be7c06ea033c1b2cfe38db5493e1fbc30d74d7c4620.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i8d8602658b2b07bda45b9be7c06ea033c1b2cfe38db5493e1fbc30d74d7c4620.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*iffb2299353bba8ab51f78fd2f6a3378d03c06b782e233ffd249b57862b465ae1.DeprovisionRequestBuilder) {
    return iffb2299353bba8ab51f78fd2f6a3378d03c06b782e233ffd249b57862b465ae1.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*i48615277135454168b173ab20d530a2f5a8a815e1b22cdb1ac56be6b3193e91c.DisableRequestBuilder) {
    return i48615277135454168b173ab20d530a2f5a8a815e1b22cdb1ac56be6b3193e91c.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*iec225655450ef7e6ba19746507d675be5cf1fe483cd29f92acd05e10aaa37928.DisableLostModeRequestBuilder) {
    return iec225655450ef7e6ba19746507d675be5cf1fe483cd29f92acd05e10aaa37928.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*idd0937487d4e7865ef09bf471120f20fed975cc48f74c411d675787d55cb7941.EnableLostModeRequestBuilder) {
    return idd0937487d4e7865ef09bf471120f20fed975cc48f74c411d675787d55cb7941.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices that executes the device management script.
func (m *ManagedDeviceRequestBuilder) Get(options *ManagedDeviceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedDevice() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice), nil
}
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*ia0939795c41941bda0750587af44dca064a923b4d54b34f571af6613831be401.GetCloudPcRemoteActionResultsRequestBuilder) {
    return ia0939795c41941bda0750587af44dca064a923b4d54b34f571af6613831be401.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i7e5e1a05c1297f2cd782c5c9ba54f7379fe2f714aaa62928cb7abe0124bb604c.GetFileVaultKeyRequestBuilder) {
    return i7e5e1a05c1297f2cd782c5c9ba54f7379fe2f714aaa62928cb7abe0124bb604c.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*i6a27f8ff131a4abb5a48228bf7f0ac1085eccc06470057ffd2500e4b0a2484ee.GetNonCompliantSettingsRequestBuilder) {
    return i6a27f8ff131a4abb5a48228bf7f0ac1085eccc06470057ffd2500e4b0a2484ee.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*ib8f2523d2206a7faaf8a92a41ed28323e6de488009f295cc76512549fd84057b.GetOemWarrantyRequestBuilder) {
    return ib8f2523d2206a7faaf8a92a41ed28323e6de488009f295cc76512549fd84057b.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*ie33b8d4c95cc420ae6ae59783b82519bc36e6aef4d9738e3a0a3fc7779d14d6e.LocateDeviceRequestBuilder) {
    return ie33b8d4c95cc420ae6ae59783b82519bc36e6aef4d9738e3a0a3fc7779d14d6e.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i9eb1267c62444a3cc384a52459a534effcdd1288be8d163fa543fe99ddb6b79d.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i9eb1267c62444a3cc384a52459a534effcdd1288be8d163fa543fe99ddb6b79d.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*ida38c086cbf3309ea906f75f123d0b968f83bb97586c4577d7721535a9ea5319.OverrideComplianceStateRequestBuilder) {
    return ida38c086cbf3309ea906f75f123d0b968f83bb97586c4577d7721535a9ea5319.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*i55d4b94e396bad9a76c0aa71af0ceab7719fbf24e6e97fa08040654aeb7839e6.PlayLostModeSoundRequestBuilder) {
    return i55d4b94e396bad9a76c0aa71af0ceab7719fbf24e6e97fa08040654aeb7839e6.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*i64c3fef8688ebbcc66e46321f822049c485d3ee4c6139420125e9c69ede03466.RebootNowRequestBuilder) {
    return i64c3fef8688ebbcc66e46321f822049c485d3ee4c6139420125e9c69ede03466.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*i775ef06d38249f8802deaeaf95042896ed2421ca137aedd4d321b0433174b436.RecoverPasscodeRequestBuilder) {
    return i775ef06d38249f8802deaeaf95042896ed2421ca137aedd4d321b0433174b436.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*i399d6e7104be0617d1cc28c6cbdfa6e39a10a68e036e7ec466bdee5b3ee9d4eb.ReenableRequestBuilder) {
    return i399d6e7104be0617d1cc28c6cbdfa6e39a10a68e036e7ec466bdee5b3ee9d4eb.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*i10f6744576cd776bd8de8f12ebe8cf8615566bc5f7fd8d948cf87e464fc18d72.RefRequestBuilder) {
    return i10f6744576cd776bd8de8f12ebe8cf8615566bc5f7fd8d948cf87e464fc18d72.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i4beea8a0fcd6812c6b72927302ba2bae142a6b88d90d27270dc33cf655870fa4.RemoteLockRequestBuilder) {
    return i4beea8a0fcd6812c6b72927302ba2bae142a6b88d90d27270dc33cf655870fa4.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*i58c250dea6df053eb964e467255d843621c0227eebc40b0fd4d14d09da2bb507.ReprovisionCloudPcRequestBuilder) {
    return i58c250dea6df053eb964e467255d843621c0227eebc40b0fd4d14d09da2bb507.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*ie4652856a45d8649b3e905653bf3242fe78b6e9af75119aedcc10f4adbcc4c43.RequestRemoteAssistanceRequestBuilder) {
    return ie4652856a45d8649b3e905653bf3242fe78b6e9af75119aedcc10f4adbcc4c43.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i2309f9a1593bcd546d2a5f3758b4196ff376a243bc8570f03a40650eab948658.ResetPasscodeRequestBuilder) {
    return i2309f9a1593bcd546d2a5f3758b4196ff376a243bc8570f03a40650eab948658.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*ie98eef0691d9bc52473f9f7e7b54f6908a1335125f5a99d5b4ea41d31f5e293b.ResizeCloudPcRequestBuilder) {
    return ie98eef0691d9bc52473f9f7e7b54f6908a1335125f5a99d5b4ea41d31f5e293b.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*i7951ff84c03d488d41d5d718ae008f4bf55d7098ecc6cfda0cda277d7fdf1011.RestoreCloudPcRequestBuilder) {
    return i7951ff84c03d488d41d5d718ae008f4bf55d7098ecc6cfda0cda277d7fdf1011.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*i5f850d24a8ec95012aa5a4b161e3d0df8fca997d498138f9635bdbbda18d8228.RetireRequestBuilder) {
    return i5f850d24a8ec95012aa5a4b161e3d0df8fca997d498138f9635bdbbda18d8228.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*id19190b7e652bf91a3888697decd8fc5fcaf97c1df9c8d9bb7a130cc29879ac2.RevokeAppleVppLicensesRequestBuilder) {
    return id19190b7e652bf91a3888697decd8fc5fcaf97c1df9c8d9bb7a130cc29879ac2.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*id3340c49abb5b01cc6459703996e6066719a8089c162cf75db81e2da32a3b6a2.RotateBitLockerKeysRequestBuilder) {
    return id3340c49abb5b01cc6459703996e6066719a8089c162cf75db81e2da32a3b6a2.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*i7470f2b18ba33a6d81ee8adcba65f6bf7ae7a097d9f62a288459325699d268d8.RotateFileVaultKeyRequestBuilder) {
    return i7470f2b18ba33a6d81ee8adcba65f6bf7ae7a097d9f62a288459325699d268d8.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*i471340460219a9537a821de60517e9a1c787386de59c0c6f948ba06021e19dd9.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return i471340460219a9537a821de60517e9a1c787386de59c0c6f948ba06021e19dd9.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*i7a68baac69887e8af4217c1d694a079001669a4df37190ec0483cbd34d6c1d91.SetDeviceNameRequestBuilder) {
    return i7a68baac69887e8af4217c1d694a079001669a4df37190ec0483cbd34d6c1d91.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*i403dcba80f2cb4eda556b9fa4a77cddfcf327cdc1c57609d2062227c8b2e278b.ShutDownRequestBuilder) {
    return i403dcba80f2cb4eda556b9fa4a77cddfcf327cdc1c57609d2062227c8b2e278b.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*idaf153750a518c69544d2691ac5f436b715b42484cea2329e43ca5fbf09af25c.SyncDeviceRequestBuilder) {
    return idaf153750a518c69544d2691ac5f436b715b42484cea2329e43ca5fbf09af25c.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*icbbc184d6edecea06de2dbcc68e352f1879de8a9ede1f4ceb144f3c0dc570ca9.TriggerConfigurationManagerActionRequestBuilder) {
    return icbbc184d6edecea06de2dbcc68e352f1879de8a9ede1f4ceb144f3c0dc570ca9.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*iddb02630e8f24f7ed576d8a4d4626288d9610749ffd3bb58b6f1d5473fbc4448.UpdateWindowsDeviceAccountRequestBuilder) {
    return iddb02630e8f24f7ed576d8a4d4626288d9610749ffd3bb58b6f1d5473fbc4448.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*iddc9cf0a62dad1d6282fde3ff766395f746dbda191d9ad7efb42d0f473a97ad4.WindowsDefenderScanRequestBuilder) {
    return iddc9cf0a62dad1d6282fde3ff766395f746dbda191d9ad7efb42d0f473a97ad4.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*ieba57c35679d54bc9735a0392f4997ffb76b5d59c981b5d3330e5bfcf8c2ebfc.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return ieba57c35679d54bc9735a0392f4997ffb76b5d59c981b5d3330e5bfcf8c2ebfc.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*i9f9c7c224fe319efce0b62161dd132d15c29815963df29526040e88fdee29c00.WipeRequestBuilder) {
    return i9f9c7c224fe319efce0b62161dd132d15c29815963df29526040e88fdee29c00.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
