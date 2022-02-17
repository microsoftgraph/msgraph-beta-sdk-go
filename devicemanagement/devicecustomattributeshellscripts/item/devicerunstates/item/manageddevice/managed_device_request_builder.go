package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i10a67c6f91af12a6475c49deb7b841658e624d11a28b96f0005fe63565300656 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    i1434459f27e17211e4051252c38508a3ea081f365263a282daaf8c052be3b68f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    i1918978867688e2b31cc1e4e908d3224bbcf2868898a0d557cef15e59a29f4d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/recoverpasscode"
    i1b70c7a61d5012fea1a35c23d2d0fb19782b089f3fb59d96b12c6d93443877ad "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/playlostmodesound"
    i1d59ef095958e89f940b4f7af413e96f65e2a5af5b3df1c6a2446f027f637d31 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    i25c2363e78374b7f442235db7cf914443bad25cc68190a064f82bdb82012af8a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/locatedevice"
    i281fbc36f39caffd50bdff10c88e8bf19c0d787d733fadf01dbd328435bcd2cb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    i297c05b50cd34b1a33440cbb68b097a45621dcb984b4e36ea93d2f4858f0cdd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/resizecloudpc"
    i2c510e5e8ee7612f6c4c6e0e6c153a78a8acfa87826e181dca7f61f207799acb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/getfilevaultkey"
    i2e1d406bc76492978492dedbc28a6ed87a1ef2c9495574b26fefd5cfbd79f83a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/overridecompliancestate"
    i3a0f344d816a2f96437d84c47216f51802713254f071c20a23b7449b718cd483 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    i3c315c76c3477a7942cf456cfa852c1169b202ba95a2ed3fbe474026d2712fa1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/disable"
    i41182c6f733abbb0a32c415e28dd2d8207edf2d67faa17ecb6d7a560e0a4b6a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    i46a5cde9dd74c7c587ce527b943ce2693d0c1fd001cfa4c8a8d6375e2ccd3642 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    i4f5a29e568986148ba57f52ed2c678de6623818f31a69c88c8b44739d6635c00 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/setdevicename"
    i5347a5e7c45186583d5207d9524d7c46d69a6dee1a7406de4aec15ab487645ab "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/reenable"
    i5d2309d6b84a03223f84adcf25d5577b4d437db276ef88bacb664e795ec0be44 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/activatedeviceesim"
    i614642abab2d1d3c9f55ee5e807c31b618d0b1f8c7453c0cb9fa286c02008675 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    i65abdb83758be0b28889af0291b9e8c5951c9b7eadfcacf99c9f7ee13db5c7b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    i6681268629f2495fb8fd65275fb69d88bbf03037faafa27b52a95dae37847ab2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
    i749f50d8b6abf8ded5f0d365c787875535ed4b0f1bcf15028a8eac8a0e1dbbad "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/deprovision"
    i760913e7e587125757b649f07da59509805a2f3b3973b5399db206dd4e13345d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    i771ef0e97de44c662fc7e5449e520adf48e49ad49a56c6e42d7fe022fd776f32 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/disablelostmode"
    i7e695d6b8f7d6482543596c477c06b07a2b89348669b07533c21741d86be6db4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
    i84cbdff17ebd17c9903a72e865ead05d96b4785f202214f490fc889d91191cae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    i8d0fce4c6a7ccf72ab5bc182303bc4ea4f2ceff9137dfc86fe8c78f3776bf75c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/ref"
    i94a64351c1a35ebd2c4cb50b7d7c17c83c2ef54835e309408b6fd21cdd9d5788 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/restorecloudpc"
    i95e0ef1f8b9858554875db2a19f1dbf63f626cf2999f26730f3fa7a35fbaa6b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/remotelock"
    i9aa1f61b065201f6a46fabfd12faa480d93f32f530d2eac2b010aa111af09782 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/retire"
    i9d6d8322578470717e04b227bb1195c387945fcf39f5ced9e2942ebcd4c30b4f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/requestremoteassistance"
    ib0802318a9892d25b7b18f4e9e3c33231084a729ddd7eff58a11dd771467d42d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/getoemwarranty"
    ib29794a0828907b3384ff3ed5ab5512a0eb8441eab2dd318ffb1ffae77cf8b59 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    ib4acb56c39383f073fb5da91067424127fe9c9b6535abafb8d9d26ee64955dda "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/syncdevice"
    ibc5952e501b3af941ba7be230bb8b4b00bb04e65b4af40183c5c2758503f5520 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/resetpasscode"
    ibf9fe59e762e14d8b36f21b9f065ea16881188d1346389e300fd6f80038ff320 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/rebootnow"
    ic3f788f4730db1e51d899ebc71575c52b2694b370cb13f62d3bfe975712d686e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    icea6d962605f78d14601dab703fa480814bb2daa0673b094989696bd88e76d0d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/enablelostmode"
    ie6d765ec9b8489a6192194a539d4a38fcd55663da465ceb4475cf3f56b4a82e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/bypassactivationlock"
    iebc53f182b649a9cdba20349ba231d52c390a6fd0ae794ca286d74992342dd63 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/wipe"
    if4d895ba47ccb55f658bb2dcb060e008746c7a68683879c299197ed0643a6521 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item/manageddevice/shutdown"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice
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
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*i5d2309d6b84a03223f84adcf25d5577b4d437db276ef88bacb664e795ec0be44.ActivateDeviceEsimRequestBuilder) {
    return i5d2309d6b84a03223f84adcf25d5577b4d437db276ef88bacb664e795ec0be44.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*ie6d765ec9b8489a6192194a539d4a38fcd55663da465ceb4475cf3f56b4a82e7.BypassActivationLockRequestBuilder) {
    return ie6d765ec9b8489a6192194a539d4a38fcd55663da465ceb4475cf3f56b4a82e7.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i1434459f27e17211e4051252c38508a3ea081f365263a282daaf8c052be3b68f.CleanWindowsDeviceRequestBuilder) {
    return i1434459f27e17211e4051252c38508a3ea081f365263a282daaf8c052be3b68f.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice{?select,expand}";
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
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*i614642abab2d1d3c9f55ee5e807c31b618d0b1f8c7453c0cb9fa286c02008675.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i614642abab2d1d3c9f55ee5e807c31b618d0b1f8c7453c0cb9fa286c02008675.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*ib29794a0828907b3384ff3ed5ab5512a0eb8441eab2dd318ffb1ffae77cf8b59.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return ib29794a0828907b3384ff3ed5ab5512a0eb8441eab2dd318ffb1ffae77cf8b59.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*i749f50d8b6abf8ded5f0d365c787875535ed4b0f1bcf15028a8eac8a0e1dbbad.DeprovisionRequestBuilder) {
    return i749f50d8b6abf8ded5f0d365c787875535ed4b0f1bcf15028a8eac8a0e1dbbad.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*i3c315c76c3477a7942cf456cfa852c1169b202ba95a2ed3fbe474026d2712fa1.DisableRequestBuilder) {
    return i3c315c76c3477a7942cf456cfa852c1169b202ba95a2ed3fbe474026d2712fa1.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i771ef0e97de44c662fc7e5449e520adf48e49ad49a56c6e42d7fe022fd776f32.DisableLostModeRequestBuilder) {
    return i771ef0e97de44c662fc7e5449e520adf48e49ad49a56c6e42d7fe022fd776f32.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*icea6d962605f78d14601dab703fa480814bb2daa0673b094989696bd88e76d0d.EnableLostModeRequestBuilder) {
    return icea6d962605f78d14601dab703fa480814bb2daa0673b094989696bd88e76d0d.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*i65abdb83758be0b28889af0291b9e8c5951c9b7eadfcacf99c9f7ee13db5c7b7.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i65abdb83758be0b28889af0291b9e8c5951c9b7eadfcacf99c9f7ee13db5c7b7.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i2c510e5e8ee7612f6c4c6e0e6c153a78a8acfa87826e181dca7f61f207799acb.GetFileVaultKeyRequestBuilder) {
    return i2c510e5e8ee7612f6c4c6e0e6c153a78a8acfa87826e181dca7f61f207799acb.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*i7e695d6b8f7d6482543596c477c06b07a2b89348669b07533c21741d86be6db4.GetNonCompliantSettingsRequestBuilder) {
    return i7e695d6b8f7d6482543596c477c06b07a2b89348669b07533c21741d86be6db4.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*ib0802318a9892d25b7b18f4e9e3c33231084a729ddd7eff58a11dd771467d42d.GetOemWarrantyRequestBuilder) {
    return ib0802318a9892d25b7b18f4e9e3c33231084a729ddd7eff58a11dd771467d42d.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*i25c2363e78374b7f442235db7cf914443bad25cc68190a064f82bdb82012af8a.LocateDeviceRequestBuilder) {
    return i25c2363e78374b7f442235db7cf914443bad25cc68190a064f82bdb82012af8a.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i3a0f344d816a2f96437d84c47216f51802713254f071c20a23b7449b718cd483.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i3a0f344d816a2f96437d84c47216f51802713254f071c20a23b7449b718cd483.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*i2e1d406bc76492978492dedbc28a6ed87a1ef2c9495574b26fefd5cfbd79f83a.OverrideComplianceStateRequestBuilder) {
    return i2e1d406bc76492978492dedbc28a6ed87a1ef2c9495574b26fefd5cfbd79f83a.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*i1b70c7a61d5012fea1a35c23d2d0fb19782b089f3fb59d96b12c6d93443877ad.PlayLostModeSoundRequestBuilder) {
    return i1b70c7a61d5012fea1a35c23d2d0fb19782b089f3fb59d96b12c6d93443877ad.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*ibf9fe59e762e14d8b36f21b9f065ea16881188d1346389e300fd6f80038ff320.RebootNowRequestBuilder) {
    return ibf9fe59e762e14d8b36f21b9f065ea16881188d1346389e300fd6f80038ff320.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*i1918978867688e2b31cc1e4e908d3224bbcf2868898a0d557cef15e59a29f4d5.RecoverPasscodeRequestBuilder) {
    return i1918978867688e2b31cc1e4e908d3224bbcf2868898a0d557cef15e59a29f4d5.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*i5347a5e7c45186583d5207d9524d7c46d69a6dee1a7406de4aec15ab487645ab.ReenableRequestBuilder) {
    return i5347a5e7c45186583d5207d9524d7c46d69a6dee1a7406de4aec15ab487645ab.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*i8d0fce4c6a7ccf72ab5bc182303bc4ea4f2ceff9137dfc86fe8c78f3776bf75c.RefRequestBuilder) {
    return i8d0fce4c6a7ccf72ab5bc182303bc4ea4f2ceff9137dfc86fe8c78f3776bf75c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i95e0ef1f8b9858554875db2a19f1dbf63f626cf2999f26730f3fa7a35fbaa6b2.RemoteLockRequestBuilder) {
    return i95e0ef1f8b9858554875db2a19f1dbf63f626cf2999f26730f3fa7a35fbaa6b2.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*i1d59ef095958e89f940b4f7af413e96f65e2a5af5b3df1c6a2446f027f637d31.ReprovisionCloudPcRequestBuilder) {
    return i1d59ef095958e89f940b4f7af413e96f65e2a5af5b3df1c6a2446f027f637d31.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*i9d6d8322578470717e04b227bb1195c387945fcf39f5ced9e2942ebcd4c30b4f.RequestRemoteAssistanceRequestBuilder) {
    return i9d6d8322578470717e04b227bb1195c387945fcf39f5ced9e2942ebcd4c30b4f.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*ibc5952e501b3af941ba7be230bb8b4b00bb04e65b4af40183c5c2758503f5520.ResetPasscodeRequestBuilder) {
    return ibc5952e501b3af941ba7be230bb8b4b00bb04e65b4af40183c5c2758503f5520.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i297c05b50cd34b1a33440cbb68b097a45621dcb984b4e36ea93d2f4858f0cdd6.ResizeCloudPcRequestBuilder) {
    return i297c05b50cd34b1a33440cbb68b097a45621dcb984b4e36ea93d2f4858f0cdd6.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*i94a64351c1a35ebd2c4cb50b7d7c17c83c2ef54835e309408b6fd21cdd9d5788.RestoreCloudPcRequestBuilder) {
    return i94a64351c1a35ebd2c4cb50b7d7c17c83c2ef54835e309408b6fd21cdd9d5788.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*i9aa1f61b065201f6a46fabfd12faa480d93f32f530d2eac2b010aa111af09782.RetireRequestBuilder) {
    return i9aa1f61b065201f6a46fabfd12faa480d93f32f530d2eac2b010aa111af09782.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i84cbdff17ebd17c9903a72e865ead05d96b4785f202214f490fc889d91191cae.RevokeAppleVppLicensesRequestBuilder) {
    return i84cbdff17ebd17c9903a72e865ead05d96b4785f202214f490fc889d91191cae.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*i281fbc36f39caffd50bdff10c88e8bf19c0d787d733fadf01dbd328435bcd2cb.RotateBitLockerKeysRequestBuilder) {
    return i281fbc36f39caffd50bdff10c88e8bf19c0d787d733fadf01dbd328435bcd2cb.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*i760913e7e587125757b649f07da59509805a2f3b3973b5399db206dd4e13345d.RotateFileVaultKeyRequestBuilder) {
    return i760913e7e587125757b649f07da59509805a2f3b3973b5399db206dd4e13345d.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*ic3f788f4730db1e51d899ebc71575c52b2694b370cb13f62d3bfe975712d686e.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return ic3f788f4730db1e51d899ebc71575c52b2694b370cb13f62d3bfe975712d686e.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*i4f5a29e568986148ba57f52ed2c678de6623818f31a69c88c8b44739d6635c00.SetDeviceNameRequestBuilder) {
    return i4f5a29e568986148ba57f52ed2c678de6623818f31a69c88c8b44739d6635c00.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*if4d895ba47ccb55f658bb2dcb060e008746c7a68683879c299197ed0643a6521.ShutDownRequestBuilder) {
    return if4d895ba47ccb55f658bb2dcb060e008746c7a68683879c299197ed0643a6521.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*ib4acb56c39383f073fb5da91067424127fe9c9b6535abafb8d9d26ee64955dda.SyncDeviceRequestBuilder) {
    return ib4acb56c39383f073fb5da91067424127fe9c9b6535abafb8d9d26ee64955dda.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*i6681268629f2495fb8fd65275fb69d88bbf03037faafa27b52a95dae37847ab2.TriggerConfigurationManagerActionRequestBuilder) {
    return i6681268629f2495fb8fd65275fb69d88bbf03037faafa27b52a95dae37847ab2.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i41182c6f733abbb0a32c415e28dd2d8207edf2d67faa17ecb6d7a560e0a4b6a2.UpdateWindowsDeviceAccountRequestBuilder) {
    return i41182c6f733abbb0a32c415e28dd2d8207edf2d67faa17ecb6d7a560e0a4b6a2.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*i10a67c6f91af12a6475c49deb7b841658e624d11a28b96f0005fe63565300656.WindowsDefenderScanRequestBuilder) {
    return i10a67c6f91af12a6475c49deb7b841658e624d11a28b96f0005fe63565300656.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i46a5cde9dd74c7c587ce527b943ce2693d0c1fd001cfa4c8a8d6375e2ccd3642.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i46a5cde9dd74c7c587ce527b943ce2693d0c1fd001cfa4c8a8d6375e2ccd3642.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*iebc53f182b649a9cdba20349ba231d52c390a6fd0ae794ca286d74992342dd63.WipeRequestBuilder) {
    return iebc53f182b649a9cdba20349ba231d52c390a6fd0ae794ca286d74992342dd63.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
