package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0427fba48e94d62efa005dbfd923b9591f370d1812fc3f9f73ebbdd681128cc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/detectedapps"
    i0630fe5357f055c4a869c7fec5dac47d8109b185d3e70c56c5c5d602d0bf2b49 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/activatedeviceesim"
    i06cb710667f52f42d44fc7648cda823b35f4f00e9af1e891bde61f0cdab9ea57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/disablelostmode"
    i0a5aa6fcc00963be9687eacfad1f376ed54f197caf08fe68554c92823d951533 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/reenable"
    i0a7666943ca9e25dae81c945031c9324a6104c6661fd7c2ec7e6f0f7398a82ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/deviceconfigurationstates"
    i0e163ea29ba2a15db04b12fee3b41564a2a8159ed31ba0bfa3acb04215135a47 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/shutdown"
    i0e96e9d7fcd789b68a4e2dc1614a66232476f904c43a89f7562d22f94f9f34f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/resizecloudpc"
    i10e987396d980603fb74cb7a5c275ec758d4ee291025714cf5320202523b1649 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/recoverpasscode"
    i13639623e2dbbc8263977fb93102819219d2c34ab64843715747f7af29c46bab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/securitybaselinestates"
    i167c255245cd6e8951c54a9c2135998064837b89b7d4d0e0f1c42f2c557b5b8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/retire"
    i16de495ce65b087460ba21136c9b72e3b2ff9a664a0756760685c434e52936ab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/bypassactivationlock"
    i186b653434f13d6f68f59e5101c7347bca64c5d54c32438393e9e563d87e9a03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/removedevicefirmwareconfigurationinterfacemanagement"
    i1e7e9d2c76289c56aebdb9d202c39b53d102418f99f9090f015754be3f4bcc19 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/retrieveremotehelpsessionwithsessionkey"
    i1f766465498501463348486d49874ef974c6aa6e056820a5caf1c3a5e2690c21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/logcollectionrequests"
    i2895d06145d6f77d60b53e7a17c29b5b08ec73440e1edd1592e6b4e58dfd3bcc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/rebootnow"
    i2a07ab69055107ffdc042942f8b812f851c28433435fb264c6d61d6e456d4a76 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/restorecloudpc"
    i3457870a56ed4379638a45891e9b1f059c53eb0a34ddead12128729ae6824a94 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/requestremotehelpsessionaccess"
    i36fc8a6a1f1f21ea4c3b8e036968cd4b9f64fe509db56f6ee2a10b53e5e23daa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/remotelock"
    i3776f24495e66930f2a43cb3ea8f98452ba8e71357481f30e1f01270eaebd89f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/createdevicelogcollectionrequest"
    i3801d66be73ebd80eac8d4afbb42c4a4b730ab7f2f26c6e5cd4dba5aa9d5482d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/setdevicename"
    i3a451234d2b2a6a6387dc5a169c250dea302a38ca99f87e49d79cd06cafb5f4d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/setcloudpcreviewstatus"
    i3b897a8ccd665fdbe167f4e370c9723f97e54817cba754eb28388f5dd299a5c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/windowsdefenderupdatesignatures"
    i3c078ae2cc611b9ba288af4dd645655413e06a9d17701e150ea19af27569731f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/requestremoteassistance"
    i4ad2411bc4f698091cb4cddb8912e565be689d70f97726b35bf923595d409c66 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/logoutsharedappledeviceactiveuser"
    i54c70818c0cf574dd041352a9ef88e3f0a2615f9d8e7d2f709c6d3799d3f777b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/locatedevice"
    i56c71bb28041877d99370b5914a9d176e4843c54d0bad29cd8a79268ef82dcf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/disable"
    i582affabda78deb10f5f0232aca6f8fa5d5bdd13a972a8227a2fca5cd056c147 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/deprovision"
    i584451cba1485854ec39deb63934f106cfbee88649f118d333314d1bb41ed15f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/deleteuserfromsharedappledevice"
    i64cf356dab9c672515b0f94125c4d954263b127208bfbb9d3bf44d33c02c57d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/manageddevicemobileappconfigurationstates"
    i76321fab1e28309a28c2fd699420e10a6d9d27fecba21c5f5ff3c89eb5e9417e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/resetpasscode"
    i7f6406f0c90d2dcc9852f30f4ca6b1df88a331683c598d8de380eaafb75271d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/windowsdefenderscan"
    i89669e1da9ff42fe3d04becaf5b4a01bd3cf18efa328f7b9f7c83e42074b1417 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/users"
    i89c236eb4c587fbb0a24528cb0e513d10719fb9adcb5a8f1c90cbd0e5f6b6c4b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/syncdevice"
    i8f3dd0f37629f789adeb5e7d3157f12e9ff9e2054bf129c520ea0079258824e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/getnoncompliantsettings"
    i943f0e8246505776fb92732c1b8002c02a1a683900dc017b18fa61b9a8c55865 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/getoemwarranty"
    i9b6d8a4ba16e42bb61f37ca948b5f74c84d3f99fd30e04f0429b325d2da5c415 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/getcloudpcreviewstatus"
    ia4ee316c6c9eb7161a71f047d40762643af93b348a3b2896f494e510a93632b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/triggerconfigurationmanageraction"
    ia6866ca41722c5c1a3bb8d2a63452764f310f9bedb206627fbff9b413170d311 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/playlostmodesound"
    iaac25517cf1038417a09bf97f681ffe842dcbe5678bf131e589ee9fbac0c73b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/reprovisioncloudpc"
    ib181475d5b174fd0c21f46c5adfa4e4cd596823a4ea13c19d41a46320ddd9753 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/rotatefilevaultkey"
    ib1b899a4faefda1522c1f9a788dfd573c30299883bae2917a797c86c3aab83b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/windowsprotectionstate"
    ib9ddea00f470a770676c6891549ef1977895ac0cb688021e1e7ca6525a2bd30e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/sendcustomnotificationtocompanyportal"
    ibe1a8097ef4c4f880f68b7002449a07a962687a3a6b20d416b08587f9959968a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/devicecategory"
    ic202d8cd388d07aa6342f443a778ec94ff0b70a0d7616cb1a86e888f74c73e98 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/getfilevaultkey"
    ic31f6e50541d7a1a2cbdcc9661762fe7f00d68f3a4b2953edee9b926748ded33 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/revokeapplevpplicenses"
    ic3d46fe96eb7d3c0bbb5b76d1ae4693f9bdb9b5675396fff7df26f117611b089 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/wipe"
    ic7c5a2e2eaab60e96d43ceb51da74be9c7ac1bcacdb964857d2c4d2331784c70 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/overridecompliancestate"
    icc098ec94f41a3c760dbaa62bd4f987e48ad5fa85a68d19121319d92c2f15eaf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/createremotehelpsession"
    id2f34cb54ddaeb8cd225320949858738929421d17f72c94883305d6139b502e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/updatewindowsdeviceaccount"
    id8bc634ccb989ff6228ca3386217960759b949d6d7a4a6dcaef53ec645ce0bb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/getcloudpcremoteactionresults"
    ie9117b0fbaf882b6a5c75879c83237b22998bbf1a1fd4b0542ac3ad863d25751 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/cleanwindowsdevice"
    ie9b83a46344e78ccf7acd52e88cae665fee20e8bed3246165738e5094dc39421 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/assignmentfilterevaluationstatusdetails"
    iea52b437c1dc7016742d797ac45ae1410bba0259df3e4fa40788171f8006769f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/devicecompliancepolicystates"
    ief4557c513ab3bc198596f3b818d3711d27cece1311420a8546b241b3ba7df10 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/rotatebitlockerkeys"
    if450a7b981b5463dc312f04522c7033d72d19492fa65aa3ae61e295d24e6b43f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/enablelostmode"
    i3b81ded3502653dd3172717ded2cc7fa51f906ee50a26f3fe4b8e0dad9677e3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/devicecompliancepolicystates/item"
    i578db475d536dd6c248552b1ff9b6a95b1348d7941cff6a6948c720fdab0b754 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/detectedapps/item"
    i66e1a9a592cf352f3ff5264cb51be3f84dee6770b52c256db783a3766f23fd2c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/deviceconfigurationstates/item"
    i7aec23a974d5a99eded144b4d6e593437866486632b1be4e531e0cbabb55685e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/securitybaselinestates/item"
    i96054b50b70fc586b43b9fdad370e6044740fbb719443c11c6ce4fbb73f65207 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/manageddevicemobileappconfigurationstates/item"
    ia1c10427ec26c5213d3cb81c564d2ac164554b68fb635b97827129a4960420f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/logcollectionrequests/item"
    idad2f9d4b430a211b97b2722dfb0a947f9cbdee6e03d60f419b691db67fec1d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/assignmentfilterevaluationstatusdetails/item"
)

// ManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDeviceItemRequestBuilderGetQueryParameters
}
// ManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim the activateDeviceEsim property
func (m *ManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*i0630fe5357f055c4a869c7fec5dac47d8109b185d3e70c56c5c5d602d0bf2b49.ActivateDeviceEsimRequestBuilder) {
    return i0630fe5357f055c4a869c7fec5dac47d8109b185d3e70c56c5c5d602d0bf2b49.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails the assignmentFilterEvaluationStatusDetails property
func (m *ManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ie9b83a46344e78ccf7acd52e88cae665fee20e8bed3246165738e5094dc39421.AssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return ie9b83a46344e78ccf7acd52e88cae665fee20e8bed3246165738e5094dc39421.NewAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item.assignmentFilterEvaluationStatusDetails.item collection
func (m *ManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*idad2f9d4b430a211b97b2722dfb0a947f9cbdee6e03d60f419b691db67fec1d7.AssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return idad2f9d4b430a211b97b2722dfb0a947f9cbdee6e03d60f419b691db67fec1d7.NewAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock the bypassActivationLock property
func (m *ManagedDeviceItemRequestBuilder) BypassActivationLock()(*i16de495ce65b087460ba21136c9b72e3b2ff9a664a0756760685c434e52936ab.BypassActivationLockRequestBuilder) {
    return i16de495ce65b087460ba21136c9b72e3b2ff9a664a0756760685c434e52936ab.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice the cleanWindowsDevice property
func (m *ManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ie9117b0fbaf882b6a5c75879c83237b22998bbf1a1fd4b0542ac3ad863d25751.CleanWindowsDeviceRequestBuilder) {
    return ie9117b0fbaf882b6a5c75879c83237b22998bbf1a1fd4b0542ac3ad863d25751.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    m := &ManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedDevices for me
func (m *ManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedDevices for me
func (m *ManagedDeviceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateDeviceLogCollectionRequest the createDeviceLogCollectionRequest property
func (m *ManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*i3776f24495e66930f2a43cb3ea8f98452ba8e71357481f30e1f01270eaebd89f.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i3776f24495e66930f2a43cb3ea8f98452ba8e71357481f30e1f01270eaebd89f.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed devices associated with the user.
func (m *ManagedDeviceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the managed devices associated with the user.
func (m *ManagedDeviceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property managedDevices in me
func (m *ManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedDevices in me
func (m *ManagedDeviceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateRemoteHelpSession the createRemoteHelpSession property
func (m *ManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*icc098ec94f41a3c760dbaa62bd4f987e48ad5fa85a68d19121319d92c2f15eaf.CreateRemoteHelpSessionRequestBuilder) {
    return icc098ec94f41a3c760dbaa62bd4f987e48ad5fa85a68d19121319d92c2f15eaf.NewCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property managedDevices for me
func (m *ManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeleteUserFromSharedAppleDevice the deleteUserFromSharedAppleDevice property
func (m *ManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*i584451cba1485854ec39deb63934f106cfbee88649f118d333314d1bb41ed15f.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i584451cba1485854ec39deb63934f106cfbee88649f118d333314d1bb41ed15f.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision the deprovision property
func (m *ManagedDeviceItemRequestBuilder) Deprovision()(*i582affabda78deb10f5f0232aca6f8fa5d5bdd13a972a8227a2fca5cd056c147.DeprovisionRequestBuilder) {
    return i582affabda78deb10f5f0232aca6f8fa5d5bdd13a972a8227a2fca5cd056c147.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps the detectedApps property
func (m *ManagedDeviceItemRequestBuilder) DetectedApps()(*i0427fba48e94d62efa005dbfd923b9591f370d1812fc3f9f73ebbdd681128cc2.DetectedAppsRequestBuilder) {
    return i0427fba48e94d62efa005dbfd923b9591f370d1812fc3f9f73ebbdd681128cc2.NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item.detectedApps.item collection
func (m *ManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*i578db475d536dd6c248552b1ff9b6a95b1348d7941cff6a6948c720fdab0b754.DetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return i578db475d536dd6c248552b1ff9b6a95b1348d7941cff6a6948c720fdab0b754.NewDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory the deviceCategory property
func (m *ManagedDeviceItemRequestBuilder) DeviceCategory()(*ibe1a8097ef4c4f880f68b7002449a07a962687a3a6b20d416b08587f9959968a.DeviceCategoryRequestBuilder) {
    return ibe1a8097ef4c4f880f68b7002449a07a962687a3a6b20d416b08587f9959968a.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates the deviceCompliancePolicyStates property
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*iea52b437c1dc7016742d797ac45ae1410bba0259df3e4fa40788171f8006769f.DeviceCompliancePolicyStatesRequestBuilder) {
    return iea52b437c1dc7016742d797ac45ae1410bba0259df3e4fa40788171f8006769f.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item.deviceCompliancePolicyStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i3b81ded3502653dd3172717ded2cc7fa51f906ee50a26f3fe4b8e0dad9677e3a.DeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return i3b81ded3502653dd3172717ded2cc7fa51f906ee50a26f3fe4b8e0dad9677e3a.NewDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates the deviceConfigurationStates property
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*i0a7666943ca9e25dae81c945031c9324a6104c6661fd7c2ec7e6f0f7398a82ff.DeviceConfigurationStatesRequestBuilder) {
    return i0a7666943ca9e25dae81c945031c9324a6104c6661fd7c2ec7e6f0f7398a82ff.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item.deviceConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*i66e1a9a592cf352f3ff5264cb51be3f84dee6770b52c256db783a3766f23fd2c.DeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return i66e1a9a592cf352f3ff5264cb51be3f84dee6770b52c256db783a3766f23fd2c.NewDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable the disable property
func (m *ManagedDeviceItemRequestBuilder) Disable()(*i56c71bb28041877d99370b5914a9d176e4843c54d0bad29cd8a79268ef82dcf7.DisableRequestBuilder) {
    return i56c71bb28041877d99370b5914a9d176e4843c54d0bad29cd8a79268ef82dcf7.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode the disableLostMode property
func (m *ManagedDeviceItemRequestBuilder) DisableLostMode()(*i06cb710667f52f42d44fc7648cda823b35f4f00e9af1e891bde61f0cdab9ea57.DisableLostModeRequestBuilder) {
    return i06cb710667f52f42d44fc7648cda823b35f4f00e9af1e891bde61f0cdab9ea57.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode the enableLostMode property
func (m *ManagedDeviceItemRequestBuilder) EnableLostMode()(*if450a7b981b5463dc312f04522c7033d72d19492fa65aa3ae61e295d24e6b43f.EnableLostModeRequestBuilder) {
    return if450a7b981b5463dc312f04522c7033d72d19492fa65aa3ae61e295d24e6b43f.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
func (m *ManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// GetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
func (m *ManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*id8bc634ccb989ff6228ca3386217960759b949d6d7a4a6dcaef53ec645ce0bb1.GetCloudPcRemoteActionResultsRequestBuilder) {
    return id8bc634ccb989ff6228ca3386217960759b949d6d7a4a6dcaef53ec645ce0bb1.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*i9b6d8a4ba16e42bb61f37ca948b5f74c84d3f99fd30e04f0429b325d2da5c415.GetCloudPcReviewStatusRequestBuilder) {
    return i9b6d8a4ba16e42bb61f37ca948b5f74c84d3f99fd30e04f0429b325d2da5c415.NewGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ic202d8cd388d07aa6342f443a778ec94ff0b70a0d7616cb1a86e888f74c73e98.GetFileVaultKeyRequestBuilder) {
    return ic202d8cd388d07aa6342f443a778ec94ff0b70a0d7616cb1a86e888f74c73e98.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*i8f3dd0f37629f789adeb5e7d3157f12e9ff9e2054bf129c520ea0079258824e7.GetNonCompliantSettingsRequestBuilder) {
    return i8f3dd0f37629f789adeb5e7d3157f12e9ff9e2054bf129c520ea0079258824e7.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *ManagedDeviceItemRequestBuilder) GetOemWarranty()(*i943f0e8246505776fb92732c1b8002c02a1a683900dc017b18fa61b9a8c55865.GetOemWarrantyRequestBuilder) {
    return i943f0e8246505776fb92732c1b8002c02a1a683900dc017b18fa61b9a8c55865.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice the locateDevice property
func (m *ManagedDeviceItemRequestBuilder) LocateDevice()(*i54c70818c0cf574dd041352a9ef88e3f0a2615f9d8e7d2f709c6d3799d3f777b.LocateDeviceRequestBuilder) {
    return i54c70818c0cf574dd041352a9ef88e3f0a2615f9d8e7d2f709c6d3799d3f777b.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests the logCollectionRequests property
func (m *ManagedDeviceItemRequestBuilder) LogCollectionRequests()(*i1f766465498501463348486d49874ef974c6aa6e056820a5caf1c3a5e2690c21.LogCollectionRequestsRequestBuilder) {
    return i1f766465498501463348486d49874ef974c6aa6e056820a5caf1c3a5e2690c21.NewLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item.logCollectionRequests.item collection
func (m *ManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*ia1c10427ec26c5213d3cb81c564d2ac164554b68fb635b97827129a4960420f1.DeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return ia1c10427ec26c5213d3cb81c564d2ac164554b68fb635b97827129a4960420f1.NewDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser the logoutSharedAppleDeviceActiveUser property
func (m *ManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i4ad2411bc4f698091cb4cddb8912e565be689d70f97726b35bf923595d409c66.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i4ad2411bc4f698091cb4cddb8912e565be689d70f97726b35bf923595d409c66.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates the managedDeviceMobileAppConfigurationStates property
func (m *ManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*i64cf356dab9c672515b0f94125c4d954263b127208bfbb9d3bf44d33c02c57d4.ManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return i64cf356dab9c672515b0f94125c4d954263b127208bfbb9d3bf44d33c02c57d4.NewManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item.managedDeviceMobileAppConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*i96054b50b70fc586b43b9fdad370e6044740fbb719443c11c6ce4fbb73f65207.ManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return i96054b50b70fc586b43b9fdad370e6044740fbb719443c11c6ce4fbb73f65207.NewManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState the overrideComplianceState property
func (m *ManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ic7c5a2e2eaab60e96d43ceb51da74be9c7ac1bcacdb964857d2c4d2331784c70.OverrideComplianceStateRequestBuilder) {
    return ic7c5a2e2eaab60e96d43ceb51da74be9c7ac1bcacdb964857d2c4d2331784c70.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in me
func (m *ManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDeviceItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PlayLostModeSound the playLostModeSound property
func (m *ManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ia6866ca41722c5c1a3bb8d2a63452764f310f9bedb206627fbff9b413170d311.PlayLostModeSoundRequestBuilder) {
    return ia6866ca41722c5c1a3bb8d2a63452764f310f9bedb206627fbff9b413170d311.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow the rebootNow property
func (m *ManagedDeviceItemRequestBuilder) RebootNow()(*i2895d06145d6f77d60b53e7a17c29b5b08ec73440e1edd1592e6b4e58dfd3bcc.RebootNowRequestBuilder) {
    return i2895d06145d6f77d60b53e7a17c29b5b08ec73440e1edd1592e6b4e58dfd3bcc.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode the recoverPasscode property
func (m *ManagedDeviceItemRequestBuilder) RecoverPasscode()(*i10e987396d980603fb74cb7a5c275ec758d4ee291025714cf5320202523b1649.RecoverPasscodeRequestBuilder) {
    return i10e987396d980603fb74cb7a5c275ec758d4ee291025714cf5320202523b1649.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable the reenable property
func (m *ManagedDeviceItemRequestBuilder) Reenable()(*i0a5aa6fcc00963be9687eacfad1f376ed54f197caf08fe68554c92823d951533.ReenableRequestBuilder) {
    return i0a5aa6fcc00963be9687eacfad1f376ed54f197caf08fe68554c92823d951533.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock the remoteLock property
func (m *ManagedDeviceItemRequestBuilder) RemoteLock()(*i36fc8a6a1f1f21ea4c3b8e036968cd4b9f64fe509db56f6ee2a10b53e5e23daa.RemoteLockRequestBuilder) {
    return i36fc8a6a1f1f21ea4c3b8e036968cd4b9f64fe509db56f6ee2a10b53e5e23daa.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement the removeDeviceFirmwareConfigurationInterfaceManagement property
func (m *ManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*i186b653434f13d6f68f59e5101c7347bca64c5d54c32438393e9e563d87e9a03.RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return i186b653434f13d6f68f59e5101c7347bca64c5d54c32438393e9e563d87e9a03.NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc the reprovisionCloudPc property
func (m *ManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*iaac25517cf1038417a09bf97f681ffe842dcbe5678bf131e589ee9fbac0c73b9.ReprovisionCloudPcRequestBuilder) {
    return iaac25517cf1038417a09bf97f681ffe842dcbe5678bf131e589ee9fbac0c73b9.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance the requestRemoteAssistance property
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*i3c078ae2cc611b9ba288af4dd645655413e06a9d17701e150ea19af27569731f.RequestRemoteAssistanceRequestBuilder) {
    return i3c078ae2cc611b9ba288af4dd645655413e06a9d17701e150ea19af27569731f.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess the requestRemoteHelpSessionAccess property
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*i3457870a56ed4379638a45891e9b1f059c53eb0a34ddead12128729ae6824a94.RequestRemoteHelpSessionAccessRequestBuilder) {
    return i3457870a56ed4379638a45891e9b1f059c53eb0a34ddead12128729ae6824a94.NewRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode the resetPasscode property
func (m *ManagedDeviceItemRequestBuilder) ResetPasscode()(*i76321fab1e28309a28c2fd699420e10a6d9d27fecba21c5f5ff3c89eb5e9417e.ResetPasscodeRequestBuilder) {
    return i76321fab1e28309a28c2fd699420e10a6d9d27fecba21c5f5ff3c89eb5e9417e.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc the resizeCloudPc property
func (m *ManagedDeviceItemRequestBuilder) ResizeCloudPc()(*i0e96e9d7fcd789b68a4e2dc1614a66232476f904c43a89f7562d22f94f9f34f5.ResizeCloudPcRequestBuilder) {
    return i0e96e9d7fcd789b68a4e2dc1614a66232476f904c43a89f7562d22f94f9f34f5.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc the restoreCloudPc property
func (m *ManagedDeviceItemRequestBuilder) RestoreCloudPc()(*i2a07ab69055107ffdc042942f8b812f851c28433435fb264c6d61d6e456d4a76.RestoreCloudPcRequestBuilder) {
    return i2a07ab69055107ffdc042942f8b812f851c28433435fb264c6d61d6e456d4a76.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire the retire property
func (m *ManagedDeviceItemRequestBuilder) Retire()(*i167c255245cd6e8951c54a9c2135998064837b89b7d4d0e0f1c42f2c557b5b8e.RetireRequestBuilder) {
    return i167c255245cd6e8951c54a9c2135998064837b89b7d4d0e0f1c42f2c557b5b8e.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*i1e7e9d2c76289c56aebdb9d202c39b53d102418f99f9090f015754be3f4bcc19.RetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return i1e7e9d2c76289c56aebdb9d202c39b53d102418f99f9090f015754be3f4bcc19.NewRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses the revokeAppleVppLicenses property
func (m *ManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ic31f6e50541d7a1a2cbdcc9661762fe7f00d68f3a4b2953edee9b926748ded33.RevokeAppleVppLicensesRequestBuilder) {
    return ic31f6e50541d7a1a2cbdcc9661762fe7f00d68f3a4b2953edee9b926748ded33.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys the rotateBitLockerKeys property
func (m *ManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ief4557c513ab3bc198596f3b818d3711d27cece1311420a8546b241b3ba7df10.RotateBitLockerKeysRequestBuilder) {
    return ief4557c513ab3bc198596f3b818d3711d27cece1311420a8546b241b3ba7df10.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey the rotateFileVaultKey property
func (m *ManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ib181475d5b174fd0c21f46c5adfa4e4cd596823a4ea13c19d41a46320ddd9753.RotateFileVaultKeyRequestBuilder) {
    return ib181475d5b174fd0c21f46c5adfa4e4cd596823a4ea13c19d41a46320ddd9753.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates the securityBaselineStates property
func (m *ManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*i13639623e2dbbc8263977fb93102819219d2c34ab64843715747f7af29c46bab.SecurityBaselineStatesRequestBuilder) {
    return i13639623e2dbbc8263977fb93102819219d2c34ab64843715747f7af29c46bab.NewSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item.securityBaselineStates.item collection
func (m *ManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*i7aec23a974d5a99eded144b4d6e593437866486632b1be4e531e0cbabb55685e.SecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return i7aec23a974d5a99eded144b4d6e593437866486632b1be4e531e0cbabb55685e.NewSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal the sendCustomNotificationToCompanyPortal property
func (m *ManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ib9ddea00f470a770676c6891549ef1977895ac0cb688021e1e7ca6525a2bd30e.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return ib9ddea00f470a770676c6891549ef1977895ac0cb688021e1e7ca6525a2bd30e.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus the setCloudPcReviewStatus property
func (m *ManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*i3a451234d2b2a6a6387dc5a169c250dea302a38ca99f87e49d79cd06cafb5f4d.SetCloudPcReviewStatusRequestBuilder) {
    return i3a451234d2b2a6a6387dc5a169c250dea302a38ca99f87e49d79cd06cafb5f4d.NewSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName the setDeviceName property
func (m *ManagedDeviceItemRequestBuilder) SetDeviceName()(*i3801d66be73ebd80eac8d4afbb42c4a4b730ab7f2f26c6e5cd4dba5aa9d5482d.SetDeviceNameRequestBuilder) {
    return i3801d66be73ebd80eac8d4afbb42c4a4b730ab7f2f26c6e5cd4dba5aa9d5482d.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown the shutDown property
func (m *ManagedDeviceItemRequestBuilder) ShutDown()(*i0e163ea29ba2a15db04b12fee3b41564a2a8159ed31ba0bfa3acb04215135a47.ShutDownRequestBuilder) {
    return i0e163ea29ba2a15db04b12fee3b41564a2a8159ed31ba0bfa3acb04215135a47.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice the syncDevice property
func (m *ManagedDeviceItemRequestBuilder) SyncDevice()(*i89c236eb4c587fbb0a24528cb0e513d10719fb9adcb5a8f1c90cbd0e5f6b6c4b.SyncDeviceRequestBuilder) {
    return i89c236eb4c587fbb0a24528cb0e513d10719fb9adcb5a8f1c90cbd0e5f6b6c4b.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction the triggerConfigurationManagerAction property
func (m *ManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ia4ee316c6c9eb7161a71f047d40762643af93b348a3b2896f494e510a93632b1.TriggerConfigurationManagerActionRequestBuilder) {
    return ia4ee316c6c9eb7161a71f047d40762643af93b348a3b2896f494e510a93632b1.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount the updateWindowsDeviceAccount property
func (m *ManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*id2f34cb54ddaeb8cd225320949858738929421d17f72c94883305d6139b502e5.UpdateWindowsDeviceAccountRequestBuilder) {
    return id2f34cb54ddaeb8cd225320949858738929421d17f72c94883305d6139b502e5.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users the users property
func (m *ManagedDeviceItemRequestBuilder) Users()(*i89669e1da9ff42fe3d04becaf5b4a01bd3cf18efa328f7b9f7c83e42074b1417.UsersRequestBuilder) {
    return i89669e1da9ff42fe3d04becaf5b4a01bd3cf18efa328f7b9f7c83e42074b1417.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan the windowsDefenderScan property
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*i7f6406f0c90d2dcc9852f30f4ca6b1df88a331683c598d8de380eaafb75271d4.WindowsDefenderScanRequestBuilder) {
    return i7f6406f0c90d2dcc9852f30f4ca6b1df88a331683c598d8de380eaafb75271d4.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures the windowsDefenderUpdateSignatures property
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*i3b897a8ccd665fdbe167f4e370c9723f97e54817cba754eb28388f5dd299a5c3.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i3b897a8ccd665fdbe167f4e370c9723f97e54817cba754eb28388f5dd299a5c3.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState the windowsProtectionState property
func (m *ManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ib1b899a4faefda1522c1f9a788dfd573c30299883bae2917a797c86c3aab83b0.WindowsProtectionStateRequestBuilder) {
    return ib1b899a4faefda1522c1f9a788dfd573c30299883bae2917a797c86c3aab83b0.NewWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe the wipe property
func (m *ManagedDeviceItemRequestBuilder) Wipe()(*ic3d46fe96eb7d3c0bbb5b76d1ae4693f9bdb9b5675396fff7df26f117611b089.WipeRequestBuilder) {
    return ic3d46fe96eb7d3c0bbb5b76d1ae4693f9bdb9b5675396fff7df26f117611b089.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
