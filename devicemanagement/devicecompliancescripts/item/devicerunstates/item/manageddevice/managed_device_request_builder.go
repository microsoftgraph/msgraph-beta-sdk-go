package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i01581a22ce7e5a64d906c396eb2487648e7fc5616dd526462c96fb1804776712 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/setdevicename"
    i079aa7122cbcb9f3bcabcaebc5456df593757b03e4a7f830e1f65760017469c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    i0c13cca989d36415dc739876dd40c2fe50666dc84e302ff4a8271964781cbb1c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/bypassactivationlock"
    i0c84353e36e032ce5d48899e9b05d7bc96dc6cf952f1955d9118461cd0e0d8dc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/resetpasscode"
    i0dc75172e3d904e3c17cac19feea3e4b7cd597cc635011f0ad7b9c6c287709ee "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/recoverpasscode"
    i1211cc294f9fd7988e03ca06f5748f9b3dfeea5382e8851c41bf5d4df3cacf21 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    i1597a8da0e02dca7679fe1618cf3bddf30236baa4af9a7246a462df67c4c6586 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    i15cfba020c7f51adfdde20b4c0376dca5692c3b3ee55027b933e16c0d4229c92 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/reenable"
    i1be050eee0cd3956c018b2b90c6895a4878635aa140951ff226127d1d653288f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/remotelock"
    i25f3c27c6c34dbf5803f756a9175392be1829732a5812ff63ee0122f12b370c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    i26df0f007e963cd3a705623d7bafd5650ac744a3b1791c990c144baad36e9c3f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/ref"
    i2b9e1fcd9b20d0aec90ef6d44f4807c8269454f4d1e9e3407dbda66646bff743 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/requestremoteassistance"
    i2f6544b9595167827b45d64abea05223fd58dc4adbbe92cae2ee0b8a42642a0d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/getfilevaultkey"
    i354053a0023bebcbeee0e8b762bd3d513b42b01cc6c47bfa8bcf007e6e9b8475 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/getoemwarranty"
    i4a12b28c2339dcb572fad82d46660a8804f60b823d44e64279c1c34e71d4d918 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/activatedeviceesim"
    i5652137b0bca56f767112ac80c752cb255c95ad94c6602dc48d6f727723195e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    i59444746a539bcd0f329d728391823fa56b9c4f4f8ae63d48e432b251ccc313c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/disablelostmode"
    i6995cf8f7df09fdb9bc3a5a89e90ab9fc450a603778061e035a9aa131fd5d23c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/deprovision"
    i6cb2c627859f5800a9be8dac6eac411445de0e016826259d1852a66d61818fe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    i6ecf5495e4fab82c3a367a1c96cd4a59f9d392eb4be8eac089e10cdd9a3ab1d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/playlostmodesound"
    i6ef42f5dbef5c78c5b9bdd5460abcd9acf92d0edc7b19db095ba82e6b61311c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    i719fc0b1001e8eddde05fddcf4847e153007eafbddc33bb2f3c2bfa27f57b82c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    i7500b45ea8e4c4cab57a2f4b3ed64aef5419e6288c72fbcf2aec568a74407900 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/enablelostmode"
    i791c50172280851f6cb7be1e46d1ac210a3f321e376e6373f41a2470d66e6d87 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    i7d535397f42f6dedf9272a7638a407a00110dde8d0085cd6ed96cf65969f5bd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    i84df8d194e15ca3d51c8a260a37304cb3f82c1bc343e113b2a030e8ea9056d4a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    i942384c3c0e1eaa9a24df6bbeed659769736f066612098bd8cbada29b9b45ca7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
    i9456791ee08080dcb0c3406c6ea5eb20b26b0153527631363f6da779f75f0c6c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/resizecloudpc"
    ia0857b0dde08665e1024073bfc89f9686daef4c3f3ae191a27a36f718e7d3a71 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    ia2fa58abc26b23475218e57813a303927c30e5d8ded402b78c47caa6c45f1186 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/locatedevice"
    iaa5b597ab54441b62cd593950545a3ea8dff1416237b75fa3de19f31c50e1941 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/retire"
    iae067e468b965b0060f03944ed5362315abc85304d5dcdb92ad489a5b11969f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/overridecompliancestate"
    icd91bf72fc830bf2bbd4a7783ed4e20bce8e10e4b0f0366e45480d0e09bd7989 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
    id1354c41e29427e9af76db11e4fff4ec1b0c32d3b59a19c8d4fb7df2e2e193c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/wipe"
    id6b285f5f2f97a9c55b3070ac89867a1fd35745e31fa9421209fd106705cc4f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    idcb846ccdb407c5a6cbdef2a9094c7176c66dd50d7e81c437f89b67cbde92c6a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/restorecloudpc"
    ie013cb25fe9a97afa50587c6d722b4d7b68d4b381d40d708e73847f73ff3331f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/syncdevice"
    ie090b0eac09482df7655ee2c8f36932e1593b3392bfe422286fa1ed326df0351 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/rebootnow"
    ie64c5288b12324e49e685657d510aeea4a23a4ad54a52a43390f378e2d842c5a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/shutdown"
    ifcb67fa9d1b7a57172e532cbdfc8d0e028e4807e0291c53f8093ef1f9b318425 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item/manageddevice/disable"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\deviceComplianceScripts\{deviceComplianceScript-id}\deviceRunStates\{deviceComplianceScriptDeviceState-id}\managedDevice
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
// ManagedDeviceRequestBuilderGetQueryParameters the managed device on which the device compliance script executed
type ManagedDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*i4a12b28c2339dcb572fad82d46660a8804f60b823d44e64279c1c34e71d4d918.ActivateDeviceEsimRequestBuilder) {
    return i4a12b28c2339dcb572fad82d46660a8804f60b823d44e64279c1c34e71d4d918.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*i0c13cca989d36415dc739876dd40c2fe50666dc84e302ff4a8271964781cbb1c.BypassActivationLockRequestBuilder) {
    return i0c13cca989d36415dc739876dd40c2fe50666dc84e302ff4a8271964781cbb1c.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i079aa7122cbcb9f3bcabcaebc5456df593757b03e4a7f830e1f65760017469c9.CleanWindowsDeviceRequestBuilder) {
    return i079aa7122cbcb9f3bcabcaebc5456df593757b03e4a7f830e1f65760017469c9.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceComplianceScripts/{deviceComplianceScript_id}/deviceRunStates/{deviceComplianceScriptDeviceState_id}/managedDevice{?select,expand}";
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
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*i791c50172280851f6cb7be1e46d1ac210a3f321e376e6373f41a2470d66e6d87.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i791c50172280851f6cb7be1e46d1ac210a3f321e376e6373f41a2470d66e6d87.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed device on which the device compliance script executed
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*id6b285f5f2f97a9c55b3070ac89867a1fd35745e31fa9421209fd106705cc4f5.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return id6b285f5f2f97a9c55b3070ac89867a1fd35745e31fa9421209fd106705cc4f5.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*i6995cf8f7df09fdb9bc3a5a89e90ab9fc450a603778061e035a9aa131fd5d23c.DeprovisionRequestBuilder) {
    return i6995cf8f7df09fdb9bc3a5a89e90ab9fc450a603778061e035a9aa131fd5d23c.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*ifcb67fa9d1b7a57172e532cbdfc8d0e028e4807e0291c53f8093ef1f9b318425.DisableRequestBuilder) {
    return ifcb67fa9d1b7a57172e532cbdfc8d0e028e4807e0291c53f8093ef1f9b318425.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i59444746a539bcd0f329d728391823fa56b9c4f4f8ae63d48e432b251ccc313c.DisableLostModeRequestBuilder) {
    return i59444746a539bcd0f329d728391823fa56b9c4f4f8ae63d48e432b251ccc313c.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*i7500b45ea8e4c4cab57a2f4b3ed64aef5419e6288c72fbcf2aec568a74407900.EnableLostModeRequestBuilder) {
    return i7500b45ea8e4c4cab57a2f4b3ed64aef5419e6288c72fbcf2aec568a74407900.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed device on which the device compliance script executed
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
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\deviceComplianceScripts\{deviceComplianceScript-id}\deviceRunStates\{deviceComplianceScriptDeviceState-id}\managedDevice\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*i84df8d194e15ca3d51c8a260a37304cb3f82c1bc343e113b2a030e8ea9056d4a.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i84df8d194e15ca3d51c8a260a37304cb3f82c1bc343e113b2a030e8ea9056d4a.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\deviceComplianceScripts\{deviceComplianceScript-id}\deviceRunStates\{deviceComplianceScriptDeviceState-id}\managedDevice\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i2f6544b9595167827b45d64abea05223fd58dc4adbbe92cae2ee0b8a42642a0d.GetFileVaultKeyRequestBuilder) {
    return i2f6544b9595167827b45d64abea05223fd58dc4adbbe92cae2ee0b8a42642a0d.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\deviceComplianceScripts\{deviceComplianceScript-id}\deviceRunStates\{deviceComplianceScriptDeviceState-id}\managedDevice\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*i942384c3c0e1eaa9a24df6bbeed659769736f066612098bd8cbada29b9b45ca7.GetNonCompliantSettingsRequestBuilder) {
    return i942384c3c0e1eaa9a24df6bbeed659769736f066612098bd8cbada29b9b45ca7.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\deviceComplianceScripts\{deviceComplianceScript-id}\deviceRunStates\{deviceComplianceScriptDeviceState-id}\managedDevice\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*i354053a0023bebcbeee0e8b762bd3d513b42b01cc6c47bfa8bcf007e6e9b8475.GetOemWarrantyRequestBuilder) {
    return i354053a0023bebcbeee0e8b762bd3d513b42b01cc6c47bfa8bcf007e6e9b8475.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*ia2fa58abc26b23475218e57813a303927c30e5d8ded402b78c47caa6c45f1186.LocateDeviceRequestBuilder) {
    return ia2fa58abc26b23475218e57813a303927c30e5d8ded402b78c47caa6c45f1186.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i6ef42f5dbef5c78c5b9bdd5460abcd9acf92d0edc7b19db095ba82e6b61311c6.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i6ef42f5dbef5c78c5b9bdd5460abcd9acf92d0edc7b19db095ba82e6b61311c6.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*iae067e468b965b0060f03944ed5362315abc85304d5dcdb92ad489a5b11969f6.OverrideComplianceStateRequestBuilder) {
    return iae067e468b965b0060f03944ed5362315abc85304d5dcdb92ad489a5b11969f6.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*i6ecf5495e4fab82c3a367a1c96cd4a59f9d392eb4be8eac089e10cdd9a3ab1d9.PlayLostModeSoundRequestBuilder) {
    return i6ecf5495e4fab82c3a367a1c96cd4a59f9d392eb4be8eac089e10cdd9a3ab1d9.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*ie090b0eac09482df7655ee2c8f36932e1593b3392bfe422286fa1ed326df0351.RebootNowRequestBuilder) {
    return ie090b0eac09482df7655ee2c8f36932e1593b3392bfe422286fa1ed326df0351.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*i0dc75172e3d904e3c17cac19feea3e4b7cd597cc635011f0ad7b9c6c287709ee.RecoverPasscodeRequestBuilder) {
    return i0dc75172e3d904e3c17cac19feea3e4b7cd597cc635011f0ad7b9c6c287709ee.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*i15cfba020c7f51adfdde20b4c0376dca5692c3b3ee55027b933e16c0d4229c92.ReenableRequestBuilder) {
    return i15cfba020c7f51adfdde20b4c0376dca5692c3b3ee55027b933e16c0d4229c92.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*i26df0f007e963cd3a705623d7bafd5650ac744a3b1791c990c144baad36e9c3f.RefRequestBuilder) {
    return i26df0f007e963cd3a705623d7bafd5650ac744a3b1791c990c144baad36e9c3f.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i1be050eee0cd3956c018b2b90c6895a4878635aa140951ff226127d1d653288f.RemoteLockRequestBuilder) {
    return i1be050eee0cd3956c018b2b90c6895a4878635aa140951ff226127d1d653288f.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*i5652137b0bca56f767112ac80c752cb255c95ad94c6602dc48d6f727723195e9.ReprovisionCloudPcRequestBuilder) {
    return i5652137b0bca56f767112ac80c752cb255c95ad94c6602dc48d6f727723195e9.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*i2b9e1fcd9b20d0aec90ef6d44f4807c8269454f4d1e9e3407dbda66646bff743.RequestRemoteAssistanceRequestBuilder) {
    return i2b9e1fcd9b20d0aec90ef6d44f4807c8269454f4d1e9e3407dbda66646bff743.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i0c84353e36e032ce5d48899e9b05d7bc96dc6cf952f1955d9118461cd0e0d8dc.ResetPasscodeRequestBuilder) {
    return i0c84353e36e032ce5d48899e9b05d7bc96dc6cf952f1955d9118461cd0e0d8dc.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i9456791ee08080dcb0c3406c6ea5eb20b26b0153527631363f6da779f75f0c6c.ResizeCloudPcRequestBuilder) {
    return i9456791ee08080dcb0c3406c6ea5eb20b26b0153527631363f6da779f75f0c6c.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*idcb846ccdb407c5a6cbdef2a9094c7176c66dd50d7e81c437f89b67cbde92c6a.RestoreCloudPcRequestBuilder) {
    return idcb846ccdb407c5a6cbdef2a9094c7176c66dd50d7e81c437f89b67cbde92c6a.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*iaa5b597ab54441b62cd593950545a3ea8dff1416237b75fa3de19f31c50e1941.RetireRequestBuilder) {
    return iaa5b597ab54441b62cd593950545a3ea8dff1416237b75fa3de19f31c50e1941.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i719fc0b1001e8eddde05fddcf4847e153007eafbddc33bb2f3c2bfa27f57b82c.RevokeAppleVppLicensesRequestBuilder) {
    return i719fc0b1001e8eddde05fddcf4847e153007eafbddc33bb2f3c2bfa27f57b82c.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*i25f3c27c6c34dbf5803f756a9175392be1829732a5812ff63ee0122f12b370c1.RotateBitLockerKeysRequestBuilder) {
    return i25f3c27c6c34dbf5803f756a9175392be1829732a5812ff63ee0122f12b370c1.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*i1211cc294f9fd7988e03ca06f5748f9b3dfeea5382e8851c41bf5d4df3cacf21.RotateFileVaultKeyRequestBuilder) {
    return i1211cc294f9fd7988e03ca06f5748f9b3dfeea5382e8851c41bf5d4df3cacf21.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*ia0857b0dde08665e1024073bfc89f9686daef4c3f3ae191a27a36f718e7d3a71.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return ia0857b0dde08665e1024073bfc89f9686daef4c3f3ae191a27a36f718e7d3a71.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*i01581a22ce7e5a64d906c396eb2487648e7fc5616dd526462c96fb1804776712.SetDeviceNameRequestBuilder) {
    return i01581a22ce7e5a64d906c396eb2487648e7fc5616dd526462c96fb1804776712.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*ie64c5288b12324e49e685657d510aeea4a23a4ad54a52a43390f378e2d842c5a.ShutDownRequestBuilder) {
    return ie64c5288b12324e49e685657d510aeea4a23a4ad54a52a43390f378e2d842c5a.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*ie013cb25fe9a97afa50587c6d722b4d7b68d4b381d40d708e73847f73ff3331f.SyncDeviceRequestBuilder) {
    return ie013cb25fe9a97afa50587c6d722b4d7b68d4b381d40d708e73847f73ff3331f.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*icd91bf72fc830bf2bbd4a7783ed4e20bce8e10e4b0f0366e45480d0e09bd7989.TriggerConfigurationManagerActionRequestBuilder) {
    return icd91bf72fc830bf2bbd4a7783ed4e20bce8e10e4b0f0366e45480d0e09bd7989.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i6cb2c627859f5800a9be8dac6eac411445de0e016826259d1852a66d61818fe7.UpdateWindowsDeviceAccountRequestBuilder) {
    return i6cb2c627859f5800a9be8dac6eac411445de0e016826259d1852a66d61818fe7.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*i1597a8da0e02dca7679fe1618cf3bddf30236baa4af9a7246a462df67c4c6586.WindowsDefenderScanRequestBuilder) {
    return i1597a8da0e02dca7679fe1618cf3bddf30236baa4af9a7246a462df67c4c6586.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i7d535397f42f6dedf9272a7638a407a00110dde8d0085cd6ed96cf65969f5bd7.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i7d535397f42f6dedf9272a7638a407a00110dde8d0085cd6ed96cf65969f5bd7.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*id1354c41e29427e9af76db11e4fff4ec1b0c32d3b59a19c8d4fb7df2e2e193c0.WipeRequestBuilder) {
    return id1354c41e29427e9af76db11e4fff4ec1b0c32d3b59a19c8d4fb7df2e2e193c0.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
