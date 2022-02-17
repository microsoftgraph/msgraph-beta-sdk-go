package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0ef421fcbedbb4f748d8ee83089e470ee90a86cffa6bcaf19c83b92c69251acf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/activatedeviceesim"
    i12945457f997038b76c1addce8a90c8ca62b883557ca17f4ed929e4d9f81333e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    i2b8eaaf12dca6b1caf6d92965f975d48582c938b998866539f203fc8303fb07e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/recoverpasscode"
    i367f8f7ad679313e0893a1f2ec58ebcaae8365d21e42bea2a7fd2fa07a5ed1e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    i368b6726b49736ebf310921e08269c416a09a0743daf3f162d12614d6ca3ab31 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/resizecloudpc"
    i36f67de05567e1f313aa022ec28754463c41e8274fe6fa52d4bebe9d0046c6de "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    i4e0a5676994c7333de5a0856accb5e58f8597e0bb575db0a84f7c4c78044e30d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/rebootnow"
    i50076d12f91727b222c982a4e996178d60e94eab4cff1c5274f8b7299bae4e7e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/getfilevaultkey"
    i58c1358147f29e5bd123e0035d47d8c5a3440dc7ec5cf3f35a5a5da00802c075 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/enablelostmode"
    i5fced4a6debf9e4ba4b285622456c818d5a430fafb44a0381dc38cf61455b9a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
    i60f02ba493f0397b4fc6fc0d1171fd9ca0ae0e65ae0cb82d91e3a037433573d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/syncdevice"
    i6b194c7bd85d2e1c766fde5cb1924ac0700e6f8aca3735127f45b594c35d2bbe "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/restorecloudpc"
    i6b6e6aabf460a64145f864799f8a336e61191a0d466e703b6864b8f831119a0a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    i8a35cf38797995d80cadd61c1064d88154bb1e3523cc76cef3a07277b75fff53 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/wipe"
    i91c28097084bee59d0e27fd1b49706a6affa75891a7d820f56da7ea648ccfb7b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    i9407f2cfbbcb08d37c3a18b3e2de49c3d90b9362555db628aba4ae9933df82c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/shutdown"
    i9d222256c612b74cd69ff50a80215f7443ea6ec60a356acb7fdd356d2f8e4e40 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/locatedevice"
    ia1246055c13ddb45ba20bf75fafa584a253388aa2b202fea2389585d5e11cd35 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/requestremoteassistance"
    iaac8ba21d10c7375dffc58c1cc79c402e0b0aef56eedf568a5c590dfd413b022 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/setdevicename"
    iab63c1e62a8342c3372a5043510c80bd2bd3e888e1a079928fab52bbb644e783 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    iac99263125505a31f860df3046e770400d14ecccf9c47583d6311f0e2601b031 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    iac9edd425a40bcb17792d31b3febf8abf88865540c98fa52a566d1f357c5d7cc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/remotelock"
    ib352a99be2884a24a144b519396cb05704120613f0fb0f4a7e26d3f7b0edb2e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    ib69e3ced126338b58bbea465961c0ca57b8985673512ebe164588ddfc541a1e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/getoemwarranty"
    ib96e7582fd56cd6244ca5177a0ea5d1918d37cc2310d308118145280a9e0abca "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    ibcd576a6d50756f1516742f671e2120eb57130ecf23494ac540fc8c3ecdc5d81 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/reenable"
    ic2049f2d0285028c14246a7cadf0d1963a3d3ea2b2fa9f5cb51cbd4c2d831470 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/ref"
    ic353e9a6d70cfb21952371c8920d2e1280732042943129a2bbe2169d6409e0fe "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
    ic57c2aa9dabd09bd487f6279023a83fe945844e253ff3e962774e58fca8d828f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/overridecompliancestate"
    ic873f4aab2c0eaaf2425aafcfd8258b3f92a0d37287acd1ecc6607ce2826cf67 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/disable"
    ic8ec55e1af553cc7692e0dba804c20b3ac62c6f32ff8e461b4cde590ac4f3913 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    ic92bfb5f57ca22fef17fceeacad00a609f92ac7c4568672799377236826d8e2b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/retire"
    icad7beff941791f99a6a5a2611873e0e8492c06550d509ff2446b4ac99be1496 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/deprovision"
    id30e6bacd67690d166b6a1b3778814a967d660032c1408ff701db48cc6a80496 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/bypassactivationlock"
    id5ab800455ffa3d9d0119e66cc2459e5a9f6ece2f4ba732d8a7dfda30a24ecae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/disablelostmode"
    ie14ba1c566da27f7a9c3942d2e3655382038d2b26276c612f344cc01df4fd3af "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    ief6073205d906b92412cb13264470179e14ca2852d07f5a4ee1e229dc092c2ba "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/resetpasscode"
    if280ca473c07280d7b754617cfc252392ca25c6ec275922c697a23307418ffda "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    if8f8944c5445fbde328af8be7e5d14b25b446b3ffeeef3954d2a2b203e464a40 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    if9da201826d7107fd319b4f9de1c348775636e2f10356a14ac6ea3db08183036 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item/manageddevice/playlostmodesound"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\deviceRunStates\{deviceHealthScriptDeviceState-id}\managedDevice
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
// ManagedDeviceRequestBuilderGetQueryParameters the managed device on which the device health script executed
type ManagedDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*i0ef421fcbedbb4f748d8ee83089e470ee90a86cffa6bcaf19c83b92c69251acf.ActivateDeviceEsimRequestBuilder) {
    return i0ef421fcbedbb4f748d8ee83089e470ee90a86cffa6bcaf19c83b92c69251acf.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*id30e6bacd67690d166b6a1b3778814a967d660032c1408ff701db48cc6a80496.BypassActivationLockRequestBuilder) {
    return id30e6bacd67690d166b6a1b3778814a967d660032c1408ff701db48cc6a80496.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i6b6e6aabf460a64145f864799f8a336e61191a0d466e703b6864b8f831119a0a.CleanWindowsDeviceRequestBuilder) {
    return i6b6e6aabf460a64145f864799f8a336e61191a0d466e703b6864b8f831119a0a.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript_id}/deviceRunStates/{deviceHealthScriptDeviceState_id}/managedDevice{?select,expand}";
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
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*i36f67de05567e1f313aa022ec28754463c41e8274fe6fa52d4bebe9d0046c6de.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i36f67de05567e1f313aa022ec28754463c41e8274fe6fa52d4bebe9d0046c6de.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed device on which the device health script executed
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*iab63c1e62a8342c3372a5043510c80bd2bd3e888e1a079928fab52bbb644e783.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return iab63c1e62a8342c3372a5043510c80bd2bd3e888e1a079928fab52bbb644e783.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*icad7beff941791f99a6a5a2611873e0e8492c06550d509ff2446b4ac99be1496.DeprovisionRequestBuilder) {
    return icad7beff941791f99a6a5a2611873e0e8492c06550d509ff2446b4ac99be1496.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*ic873f4aab2c0eaaf2425aafcfd8258b3f92a0d37287acd1ecc6607ce2826cf67.DisableRequestBuilder) {
    return ic873f4aab2c0eaaf2425aafcfd8258b3f92a0d37287acd1ecc6607ce2826cf67.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*id5ab800455ffa3d9d0119e66cc2459e5a9f6ece2f4ba732d8a7dfda30a24ecae.DisableLostModeRequestBuilder) {
    return id5ab800455ffa3d9d0119e66cc2459e5a9f6ece2f4ba732d8a7dfda30a24ecae.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*i58c1358147f29e5bd123e0035d47d8c5a3440dc7ec5cf3f35a5a5da00802c075.EnableLostModeRequestBuilder) {
    return i58c1358147f29e5bd123e0035d47d8c5a3440dc7ec5cf3f35a5a5da00802c075.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed device on which the device health script executed
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
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\deviceRunStates\{deviceHealthScriptDeviceState-id}\managedDevice\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*ib96e7582fd56cd6244ca5177a0ea5d1918d37cc2310d308118145280a9e0abca.GetCloudPcRemoteActionResultsRequestBuilder) {
    return ib96e7582fd56cd6244ca5177a0ea5d1918d37cc2310d308118145280a9e0abca.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\deviceRunStates\{deviceHealthScriptDeviceState-id}\managedDevice\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i50076d12f91727b222c982a4e996178d60e94eab4cff1c5274f8b7299bae4e7e.GetFileVaultKeyRequestBuilder) {
    return i50076d12f91727b222c982a4e996178d60e94eab4cff1c5274f8b7299bae4e7e.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\deviceRunStates\{deviceHealthScriptDeviceState-id}\managedDevice\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*i5fced4a6debf9e4ba4b285622456c818d5a430fafb44a0381dc38cf61455b9a8.GetNonCompliantSettingsRequestBuilder) {
    return i5fced4a6debf9e4ba4b285622456c818d5a430fafb44a0381dc38cf61455b9a8.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\deviceRunStates\{deviceHealthScriptDeviceState-id}\managedDevice\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*ib69e3ced126338b58bbea465961c0ca57b8985673512ebe164588ddfc541a1e7.GetOemWarrantyRequestBuilder) {
    return ib69e3ced126338b58bbea465961c0ca57b8985673512ebe164588ddfc541a1e7.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*i9d222256c612b74cd69ff50a80215f7443ea6ec60a356acb7fdd356d2f8e4e40.LocateDeviceRequestBuilder) {
    return i9d222256c612b74cd69ff50a80215f7443ea6ec60a356acb7fdd356d2f8e4e40.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i91c28097084bee59d0e27fd1b49706a6affa75891a7d820f56da7ea648ccfb7b.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i91c28097084bee59d0e27fd1b49706a6affa75891a7d820f56da7ea648ccfb7b.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*ic57c2aa9dabd09bd487f6279023a83fe945844e253ff3e962774e58fca8d828f.OverrideComplianceStateRequestBuilder) {
    return ic57c2aa9dabd09bd487f6279023a83fe945844e253ff3e962774e58fca8d828f.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*if9da201826d7107fd319b4f9de1c348775636e2f10356a14ac6ea3db08183036.PlayLostModeSoundRequestBuilder) {
    return if9da201826d7107fd319b4f9de1c348775636e2f10356a14ac6ea3db08183036.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*i4e0a5676994c7333de5a0856accb5e58f8597e0bb575db0a84f7c4c78044e30d.RebootNowRequestBuilder) {
    return i4e0a5676994c7333de5a0856accb5e58f8597e0bb575db0a84f7c4c78044e30d.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*i2b8eaaf12dca6b1caf6d92965f975d48582c938b998866539f203fc8303fb07e.RecoverPasscodeRequestBuilder) {
    return i2b8eaaf12dca6b1caf6d92965f975d48582c938b998866539f203fc8303fb07e.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*ibcd576a6d50756f1516742f671e2120eb57130ecf23494ac540fc8c3ecdc5d81.ReenableRequestBuilder) {
    return ibcd576a6d50756f1516742f671e2120eb57130ecf23494ac540fc8c3ecdc5d81.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*ic2049f2d0285028c14246a7cadf0d1963a3d3ea2b2fa9f5cb51cbd4c2d831470.RefRequestBuilder) {
    return ic2049f2d0285028c14246a7cadf0d1963a3d3ea2b2fa9f5cb51cbd4c2d831470.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*iac9edd425a40bcb17792d31b3febf8abf88865540c98fa52a566d1f357c5d7cc.RemoteLockRequestBuilder) {
    return iac9edd425a40bcb17792d31b3febf8abf88865540c98fa52a566d1f357c5d7cc.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*if280ca473c07280d7b754617cfc252392ca25c6ec275922c697a23307418ffda.ReprovisionCloudPcRequestBuilder) {
    return if280ca473c07280d7b754617cfc252392ca25c6ec275922c697a23307418ffda.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*ia1246055c13ddb45ba20bf75fafa584a253388aa2b202fea2389585d5e11cd35.RequestRemoteAssistanceRequestBuilder) {
    return ia1246055c13ddb45ba20bf75fafa584a253388aa2b202fea2389585d5e11cd35.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*ief6073205d906b92412cb13264470179e14ca2852d07f5a4ee1e229dc092c2ba.ResetPasscodeRequestBuilder) {
    return ief6073205d906b92412cb13264470179e14ca2852d07f5a4ee1e229dc092c2ba.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i368b6726b49736ebf310921e08269c416a09a0743daf3f162d12614d6ca3ab31.ResizeCloudPcRequestBuilder) {
    return i368b6726b49736ebf310921e08269c416a09a0743daf3f162d12614d6ca3ab31.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*i6b194c7bd85d2e1c766fde5cb1924ac0700e6f8aca3735127f45b594c35d2bbe.RestoreCloudPcRequestBuilder) {
    return i6b194c7bd85d2e1c766fde5cb1924ac0700e6f8aca3735127f45b594c35d2bbe.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*ic92bfb5f57ca22fef17fceeacad00a609f92ac7c4568672799377236826d8e2b.RetireRequestBuilder) {
    return ic92bfb5f57ca22fef17fceeacad00a609f92ac7c4568672799377236826d8e2b.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*ie14ba1c566da27f7a9c3942d2e3655382038d2b26276c612f344cc01df4fd3af.RevokeAppleVppLicensesRequestBuilder) {
    return ie14ba1c566da27f7a9c3942d2e3655382038d2b26276c612f344cc01df4fd3af.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*if8f8944c5445fbde328af8be7e5d14b25b446b3ffeeef3954d2a2b203e464a40.RotateBitLockerKeysRequestBuilder) {
    return if8f8944c5445fbde328af8be7e5d14b25b446b3ffeeef3954d2a2b203e464a40.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*iac99263125505a31f860df3046e770400d14ecccf9c47583d6311f0e2601b031.RotateFileVaultKeyRequestBuilder) {
    return iac99263125505a31f860df3046e770400d14ecccf9c47583d6311f0e2601b031.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*i12945457f997038b76c1addce8a90c8ca62b883557ca17f4ed929e4d9f81333e.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return i12945457f997038b76c1addce8a90c8ca62b883557ca17f4ed929e4d9f81333e.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*iaac8ba21d10c7375dffc58c1cc79c402e0b0aef56eedf568a5c590dfd413b022.SetDeviceNameRequestBuilder) {
    return iaac8ba21d10c7375dffc58c1cc79c402e0b0aef56eedf568a5c590dfd413b022.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*i9407f2cfbbcb08d37c3a18b3e2de49c3d90b9362555db628aba4ae9933df82c0.ShutDownRequestBuilder) {
    return i9407f2cfbbcb08d37c3a18b3e2de49c3d90b9362555db628aba4ae9933df82c0.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*i60f02ba493f0397b4fc6fc0d1171fd9ca0ae0e65ae0cb82d91e3a037433573d9.SyncDeviceRequestBuilder) {
    return i60f02ba493f0397b4fc6fc0d1171fd9ca0ae0e65ae0cb82d91e3a037433573d9.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*ic353e9a6d70cfb21952371c8920d2e1280732042943129a2bbe2169d6409e0fe.TriggerConfigurationManagerActionRequestBuilder) {
    return ic353e9a6d70cfb21952371c8920d2e1280732042943129a2bbe2169d6409e0fe.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i367f8f7ad679313e0893a1f2ec58ebcaae8365d21e42bea2a7fd2fa07a5ed1e1.UpdateWindowsDeviceAccountRequestBuilder) {
    return i367f8f7ad679313e0893a1f2ec58ebcaae8365d21e42bea2a7fd2fa07a5ed1e1.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*ib352a99be2884a24a144b519396cb05704120613f0fb0f4a7e26d3f7b0edb2e2.WindowsDefenderScanRequestBuilder) {
    return ib352a99be2884a24a144b519396cb05704120613f0fb0f4a7e26d3f7b0edb2e2.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*ic8ec55e1af553cc7692e0dba804c20b3ac62c6f32ff8e461b4cde590ac4f3913.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return ic8ec55e1af553cc7692e0dba804c20b3ac62c6f32ff8e461b4cde590ac4f3913.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*i8a35cf38797995d80cadd61c1064d88154bb1e3523cc76cef3a07277b75fff53.WipeRequestBuilder) {
    return i8a35cf38797995d80cadd61c1064d88154bb1e3523cc76cef3a07277b75fff53.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
