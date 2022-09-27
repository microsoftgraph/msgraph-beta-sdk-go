package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01b9c3c8f91d82e0acf6987d9d8a9d45e8970e9cc61e6faf4e24f9e53ef7d2c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/remotelock"
    i08036a752a811a186fe4cfbc410dc173a0482af2aea88d8ef8789f64be1a8a50 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/windowsprotectionstate"
    i0a9788a3556a4ca6e533dc1166cb3e642160b8b23aadcf53f0f54da24e25ee50 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/overridecompliancestate"
    i0e5c94a184309b3002cff816ad6e7906fadaccefe925a5a71b0eaf7da1cee573 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/devicecategory"
    i124459d67c21db7422395135441513bd6a695df9fe6596ae0672a66ea9e85780 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/reenable"
    i19c77245974ed8a6ef9d21e448b4f80b08071e818653a52c1710261948d25db3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/requestremotehelpsessionaccess"
    i1a0e017de38d695b9d703bf70a4161eaea10b82614aa1d2ad5623c980ccc89b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/setdevicename"
    i1cfac124b8720ab5c719cf26bfc2d8d83f55a69d715b8813458367cee9373b60 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getnoncompliantsettings"
    i1e9b4ce64039f5b2e7c5ab1ff6519bc2b54393233bd0de932f805ffa42b56ed5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/syncdevice"
    i1ed96ff0c3371d2b8be19e68b3491f7f051e944528a42666ad60e9a3df769b5d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/logoutsharedappledeviceactiveuser"
    i1f1456a44a9b0409729cadca40fbbdf79a42d5a2574cf3d332d1deb5f5f62202 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/rotatebitlockerkeys"
    i2011cd53e380c8be4e418adb3871ddb5cf1719dfb3ca6ee1b408b5211399c258 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/windowsdefenderupdatesignatures"
    i37fecb3bffa962bb92bae6d06b3cd6c4cc85201a2aa5cf7d9d6e15db1e249758 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/revokeapplevpplicenses"
    i39ac59e2082190c94b4885c6ada792a6840b6ab9dcfcf6667e820bf042be62ff "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/setcloudpcreviewstatus"
    i3f7d5f2f782c35626e08180d2725ab649a5a2048dbc7a6dc1a4f3493ed5aed4e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/bypassactivationlock"
    i4014a9f77365c024ff8b81b517cc90c5db1deeaee89910692d948bd54130b4ab "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/users"
    i4566609add3e4bacdac10bdd501bfbc92cf44ca7aa6d3492d3c92cd84f86564a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/windowsdefenderscan"
    i4fa46045cdcf8f3536238ca61013da101771dfeaaa9fed96aaf28f151ec66162 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/devicecompliancepolicystates"
    i5103c90c1b0640a73cfc58b6d80783fcc03b5adb82eb2e098521278c8403ebbc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getoemwarranty"
    i511c6c2c2693c3294886a17581fa0e4993806d6e442616fc5613bb17d2f9a3e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/rebootnow"
    i54aaa976711b04a9d75b5a6b447501e988e9c492e4339da93ee652eda4f23f04 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/wipe"
    i5533b65c5e7ef6c9cf517e9d69b5af07f1cec79c9dc96399fcbe89d0ec21a145 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/restorecloudpc"
    i61cf99cf0bb474cf95f287f10d9cc99a9b9474bfc2c89d61e472f941642ebc12 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/enablelostmode"
    i63c852e0e601ed6586a514ddd8fcecf4281a23511e01d39fc62af91ac0292701 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/disablelostmode"
    i658174871e5ed8a97ec6404a2a51d6533ca0556946eb0312fbdea2092b5a0356 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/shutdown"
    i6e7ba191a5fdb7201d9a6092955773ebfa9ecb59bdbd1e57be20c7989ea54226 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/detectedapps"
    i6f21f93e6f07adafa56b75d1c7c78e9df57af9204e36460aad874e8c9617c701 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/playlostmodesound"
    i769d01770c020a962ac6bfb317cba092b2eec617a06d92b29070491a0e4d7002 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/createremotehelpsession"
    i77ea5e92da231c3462e6daf6429f47a312c08417255e97ac007a61b747062e52 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/manageddevicemobileappconfigurationstates"
    i7d468003528ad5d962d3cce255cb40c4f894227cc59c71443390e09ab8ca40db "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getcloudpcremoteactionresults"
    i7d4fd4db8c253e46353ca32adf5b0f49f79c9b16b0709c7adc865da540dc64f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/logcollectionrequests"
    i86e677b9939bcc1c1f4b1438dd0178ec6f297b7d3c5baa413730c16ab30ff0a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/cleanwindowsdevice"
    i9479492bd537ee0da9d7cbaf1a7bda614d8af9b2db9f340f4d162cfdab45d7d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/resetpasscode"
    i964b81837112ab4b6bfd1cf6220e1e8cb634c0c56d1238300c9c00340217450d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/resizecloudpc"
    i9b07e3d1ad25d8bac7ab04fef250c221091aa5e13eba25491edfadca2f67508b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/triggerconfigurationmanageraction"
    i9c36cb050999ff8fab3434eaa458ff65d63d50ae824d7965d7bd33cff08b058a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/rotatefilevaultkey"
    ia033ed63262ddffc302f0ba2894aff599a6138a3fb8355208388deb36543c735 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getcloudpcreviewstatus"
    ia36af35bae5da4cc0bc0059a8dd1595ea3098bbaf66a723ad498408f9329538d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/requestremoteassistance"
    ia8bf08ad44640f5a1416c953b553221a5b755b2af9acc8770a00a1522e4650a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/locatedevice"
    iabf61e946f47f041682806eff385ca77fa52f76f6756d068e2e853bebc2c381b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/sendcustomnotificationtocompanyportal"
    ib4d6bc26f8783e0021130bc5307c6d883c818e4a41f6cb4f64641101753c3a9c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deprovision"
    iba0e44812d3f93778c7d884706c3eb3fad191cee346e2a62205385f2c63dc366 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getfilevaultkey"
    ic294f6a4dcf11553c619c3d81c4bed2bbc144506e96d7e69f11523299905e00e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/removedevicefirmwareconfigurationinterfacemanagement"
    ic41b8c2fdd1ab9c9e0d3c3744785e36b1c26031b43bec2708dcfb79053ed891b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/activatedeviceesim"
    ic4a4e554ac185065d3036fce9e0be725936f6f53aefb12ab4ac38f451bb309a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/createdevicelogcollectionrequest"
    ic7b58bd1f7ec7e59ddafda5fb2b6789ce5abdfcba5b1e47d8b82e90e8daeb055 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/reprovisioncloudpc"
    icb6ec007efa06597d8fa269fd84ddc90f7dd27cc757bf2d9fc8f6249c2efc2da "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deleteuserfromsharedappledevice"
    id1fd9f59ff6efc1a25cf2a7067af3ac8949a4357a777c40891113b5c6f9a3a97 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/updatewindowsdeviceaccount"
    ida5dd88d263aad80639be0db6c9f33ac2ef69ddb1701e98adef543f2d2e39c09 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/disable"
    ie1efb01c1dca39b9cc510aab2614914c0ae675e30562a196406520ea8c0163e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/securitybaselinestates"
    ie44ea0c4c2478526c3387d0067c7beeb549f1ececdbea0c42e565e76c317dd35 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deviceconfigurationstates"
    ieeb100fe5ba14a8141a3d99450da9262e50bb100dbdddf0192a6178b9a775b45 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/retrieveremotehelpsessionwithsessionkey"
    if75bb98073f1131ee02501bf098984c52a24d708147db3112ea968e2fc7b2c03 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/retire"
    ifb1c8b230156b71d830111db728223a0fcfda4e39c490d27d18e1957c1c35ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/recoverpasscode"
    iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/assignmentfilterevaluationstatusdetails"
    i032a97666f32bb19742b05d13d034f3b01289570f99da941138e57983d0680b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/detectedapps/item"
    i0f7c3711eb6736d4154b8664f763bba6cd17670f2f39b4f1a6cf7fd15aa16312 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/devicecompliancepolicystates/item"
    i299912d62f14997dd65aac2ade28936d603ff421576c4948af1d264ab7e90bd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/manageddevicemobileappconfigurationstates/item"
    i8ab1046e6a31a43a95a886daf359031d08afecdbcaae268849df05ded7a31a06 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/assignmentfilterevaluationstatusdetails/item"
    i8bf3e9602de99585216828ab1a7ba1b28f64f44b09c1192ca701225a40bdb050 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/securitybaselinestates/item"
    ib34591ad39ce97862312af86b9741ca3b1a7d0746218c1f1269fecd90b2e0ee6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/logcollectionrequests/item"
    ibf3c9129b1fe1c1f3482c4943cfae633d290cb95cfe0dec78735ebb6afbf6ac2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deviceconfigurationstates/item"
)

// ManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
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
// ManagedDeviceItemRequestBuilderGetQueryParameters the list of managed devices.
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
func (m *ManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ic41b8c2fdd1ab9c9e0d3c3744785e36b1c26031b43bec2708dcfb79053ed891b.ActivateDeviceEsimRequestBuilder) {
    return ic41b8c2fdd1ab9c9e0d3c3744785e36b1c26031b43bec2708dcfb79053ed891b.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails the assignmentFilterEvaluationStatusDetails property
func (m *ManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308.AssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308.NewAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.assignmentFilterEvaluationStatusDetails.item collection
func (m *ManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*i8ab1046e6a31a43a95a886daf359031d08afecdbcaae268849df05ded7a31a06.AssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return i8ab1046e6a31a43a95a886daf359031d08afecdbcaae268849df05ded7a31a06.NewAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock the bypassActivationLock property
func (m *ManagedDeviceItemRequestBuilder) BypassActivationLock()(*i3f7d5f2f782c35626e08180d2725ab649a5a2048dbc7a6dc1a4f3493ed5aed4e.BypassActivationLockRequestBuilder) {
    return i3f7d5f2f782c35626e08180d2725ab649a5a2048dbc7a6dc1a4f3493ed5aed4e.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice the cleanWindowsDevice property
func (m *ManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*i86e677b9939bcc1c1f4b1438dd0178ec6f297b7d3c5baa413730c16ab30ff0a9.CleanWindowsDeviceRequestBuilder) {
    return i86e677b9939bcc1c1f4b1438dd0178ec6f297b7d3c5baa413730c16ab30ff0a9.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    m := &ManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property managedDevices for deviceManagement
func (m *ManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedDevices for deviceManagement
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
func (m *ManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ic4a4e554ac185065d3036fce9e0be725936f6f53aefb12ab4ac38f451bb309a7.CreateDeviceLogCollectionRequestRequestBuilder) {
    return ic4a4e554ac185065d3036fce9e0be725936f6f53aefb12ab4ac38f451bb309a7.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of managed devices.
func (m *ManagedDeviceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of managed devices.
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
// CreatePatchRequestInformation update the navigation property managedDevices in deviceManagement
func (m *ManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedDevices in deviceManagement
func (m *ManagedDeviceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateRemoteHelpSession the createRemoteHelpSession property
func (m *ManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*i769d01770c020a962ac6bfb317cba092b2eec617a06d92b29070491a0e4d7002.CreateRemoteHelpSessionRequestBuilder) {
    return i769d01770c020a962ac6bfb317cba092b2eec617a06d92b29070491a0e4d7002.NewCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property managedDevices for deviceManagement
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
func (m *ManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*icb6ec007efa06597d8fa269fd84ddc90f7dd27cc757bf2d9fc8f6249c2efc2da.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return icb6ec007efa06597d8fa269fd84ddc90f7dd27cc757bf2d9fc8f6249c2efc2da.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision the deprovision property
func (m *ManagedDeviceItemRequestBuilder) Deprovision()(*ib4d6bc26f8783e0021130bc5307c6d883c818e4a41f6cb4f64641101753c3a9c.DeprovisionRequestBuilder) {
    return ib4d6bc26f8783e0021130bc5307c6d883c818e4a41f6cb4f64641101753c3a9c.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps the detectedApps property
func (m *ManagedDeviceItemRequestBuilder) DetectedApps()(*i6e7ba191a5fdb7201d9a6092955773ebfa9ecb59bdbd1e57be20c7989ea54226.DetectedAppsRequestBuilder) {
    return i6e7ba191a5fdb7201d9a6092955773ebfa9ecb59bdbd1e57be20c7989ea54226.NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.detectedApps.item collection
func (m *ManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*i032a97666f32bb19742b05d13d034f3b01289570f99da941138e57983d0680b7.DetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return i032a97666f32bb19742b05d13d034f3b01289570f99da941138e57983d0680b7.NewDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory the deviceCategory property
func (m *ManagedDeviceItemRequestBuilder) DeviceCategory()(*i0e5c94a184309b3002cff816ad6e7906fadaccefe925a5a71b0eaf7da1cee573.DeviceCategoryRequestBuilder) {
    return i0e5c94a184309b3002cff816ad6e7906fadaccefe925a5a71b0eaf7da1cee573.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates the deviceCompliancePolicyStates property
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*i4fa46045cdcf8f3536238ca61013da101771dfeaaa9fed96aaf28f151ec66162.DeviceCompliancePolicyStatesRequestBuilder) {
    return i4fa46045cdcf8f3536238ca61013da101771dfeaaa9fed96aaf28f151ec66162.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.deviceCompliancePolicyStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i0f7c3711eb6736d4154b8664f763bba6cd17670f2f39b4f1a6cf7fd15aa16312.DeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return i0f7c3711eb6736d4154b8664f763bba6cd17670f2f39b4f1a6cf7fd15aa16312.NewDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates the deviceConfigurationStates property
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ie44ea0c4c2478526c3387d0067c7beeb549f1ececdbea0c42e565e76c317dd35.DeviceConfigurationStatesRequestBuilder) {
    return ie44ea0c4c2478526c3387d0067c7beeb549f1ececdbea0c42e565e76c317dd35.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.deviceConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*ibf3c9129b1fe1c1f3482c4943cfae633d290cb95cfe0dec78735ebb6afbf6ac2.DeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return ibf3c9129b1fe1c1f3482c4943cfae633d290cb95cfe0dec78735ebb6afbf6ac2.NewDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable the disable property
func (m *ManagedDeviceItemRequestBuilder) Disable()(*ida5dd88d263aad80639be0db6c9f33ac2ef69ddb1701e98adef543f2d2e39c09.DisableRequestBuilder) {
    return ida5dd88d263aad80639be0db6c9f33ac2ef69ddb1701e98adef543f2d2e39c09.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode the disableLostMode property
func (m *ManagedDeviceItemRequestBuilder) DisableLostMode()(*i63c852e0e601ed6586a514ddd8fcecf4281a23511e01d39fc62af91ac0292701.DisableLostModeRequestBuilder) {
    return i63c852e0e601ed6586a514ddd8fcecf4281a23511e01d39fc62af91ac0292701.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode the enableLostMode property
func (m *ManagedDeviceItemRequestBuilder) EnableLostMode()(*i61cf99cf0bb474cf95f287f10d9cc99a9b9474bfc2c89d61e472f941642ebc12.EnableLostModeRequestBuilder) {
    return i61cf99cf0bb474cf95f287f10d9cc99a9b9474bfc2c89d61e472f941642ebc12.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of managed devices.
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
func (m *ManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*i7d468003528ad5d962d3cce255cb40c4f894227cc59c71443390e09ab8ca40db.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i7d468003528ad5d962d3cce255cb40c4f894227cc59c71443390e09ab8ca40db.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ia033ed63262ddffc302f0ba2894aff599a6138a3fb8355208388deb36543c735.GetCloudPcReviewStatusRequestBuilder) {
    return ia033ed63262ddffc302f0ba2894aff599a6138a3fb8355208388deb36543c735.NewGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ManagedDeviceItemRequestBuilder) GetFileVaultKey()(*iba0e44812d3f93778c7d884706c3eb3fad191cee346e2a62205385f2c63dc366.GetFileVaultKeyRequestBuilder) {
    return iba0e44812d3f93778c7d884706c3eb3fad191cee346e2a62205385f2c63dc366.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*i1cfac124b8720ab5c719cf26bfc2d8d83f55a69d715b8813458367cee9373b60.GetNonCompliantSettingsRequestBuilder) {
    return i1cfac124b8720ab5c719cf26bfc2d8d83f55a69d715b8813458367cee9373b60.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *ManagedDeviceItemRequestBuilder) GetOemWarranty()(*i5103c90c1b0640a73cfc58b6d80783fcc03b5adb82eb2e098521278c8403ebbc.GetOemWarrantyRequestBuilder) {
    return i5103c90c1b0640a73cfc58b6d80783fcc03b5adb82eb2e098521278c8403ebbc.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice the locateDevice property
func (m *ManagedDeviceItemRequestBuilder) LocateDevice()(*ia8bf08ad44640f5a1416c953b553221a5b755b2af9acc8770a00a1522e4650a6.LocateDeviceRequestBuilder) {
    return ia8bf08ad44640f5a1416c953b553221a5b755b2af9acc8770a00a1522e4650a6.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests the logCollectionRequests property
func (m *ManagedDeviceItemRequestBuilder) LogCollectionRequests()(*i7d4fd4db8c253e46353ca32adf5b0f49f79c9b16b0709c7adc865da540dc64f4.LogCollectionRequestsRequestBuilder) {
    return i7d4fd4db8c253e46353ca32adf5b0f49f79c9b16b0709c7adc865da540dc64f4.NewLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.logCollectionRequests.item collection
func (m *ManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*ib34591ad39ce97862312af86b9741ca3b1a7d0746218c1f1269fecd90b2e0ee6.DeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return ib34591ad39ce97862312af86b9741ca3b1a7d0746218c1f1269fecd90b2e0ee6.NewDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser the logoutSharedAppleDeviceActiveUser property
func (m *ManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i1ed96ff0c3371d2b8be19e68b3491f7f051e944528a42666ad60e9a3df769b5d.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i1ed96ff0c3371d2b8be19e68b3491f7f051e944528a42666ad60e9a3df769b5d.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates the managedDeviceMobileAppConfigurationStates property
func (m *ManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*i77ea5e92da231c3462e6daf6429f47a312c08417255e97ac007a61b747062e52.ManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return i77ea5e92da231c3462e6daf6429f47a312c08417255e97ac007a61b747062e52.NewManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.managedDeviceMobileAppConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*i299912d62f14997dd65aac2ade28936d603ff421576c4948af1d264ab7e90bd0.ManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return i299912d62f14997dd65aac2ade28936d603ff421576c4948af1d264ab7e90bd0.NewManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState the overrideComplianceState property
func (m *ManagedDeviceItemRequestBuilder) OverrideComplianceState()(*i0a9788a3556a4ca6e533dc1166cb3e642160b8b23aadcf53f0f54da24e25ee50.OverrideComplianceStateRequestBuilder) {
    return i0a9788a3556a4ca6e533dc1166cb3e642160b8b23aadcf53f0f54da24e25ee50.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in deviceManagement
func (m *ManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// PlayLostModeSound the playLostModeSound property
func (m *ManagedDeviceItemRequestBuilder) PlayLostModeSound()(*i6f21f93e6f07adafa56b75d1c7c78e9df57af9204e36460aad874e8c9617c701.PlayLostModeSoundRequestBuilder) {
    return i6f21f93e6f07adafa56b75d1c7c78e9df57af9204e36460aad874e8c9617c701.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow the rebootNow property
func (m *ManagedDeviceItemRequestBuilder) RebootNow()(*i511c6c2c2693c3294886a17581fa0e4993806d6e442616fc5613bb17d2f9a3e1.RebootNowRequestBuilder) {
    return i511c6c2c2693c3294886a17581fa0e4993806d6e442616fc5613bb17d2f9a3e1.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode the recoverPasscode property
func (m *ManagedDeviceItemRequestBuilder) RecoverPasscode()(*ifb1c8b230156b71d830111db728223a0fcfda4e39c490d27d18e1957c1c35ab9.RecoverPasscodeRequestBuilder) {
    return ifb1c8b230156b71d830111db728223a0fcfda4e39c490d27d18e1957c1c35ab9.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable the reenable property
func (m *ManagedDeviceItemRequestBuilder) Reenable()(*i124459d67c21db7422395135441513bd6a695df9fe6596ae0672a66ea9e85780.ReenableRequestBuilder) {
    return i124459d67c21db7422395135441513bd6a695df9fe6596ae0672a66ea9e85780.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock the remoteLock property
func (m *ManagedDeviceItemRequestBuilder) RemoteLock()(*i01b9c3c8f91d82e0acf6987d9d8a9d45e8970e9cc61e6faf4e24f9e53ef7d2c6.RemoteLockRequestBuilder) {
    return i01b9c3c8f91d82e0acf6987d9d8a9d45e8970e9cc61e6faf4e24f9e53ef7d2c6.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement the removeDeviceFirmwareConfigurationInterfaceManagement property
func (m *ManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ic294f6a4dcf11553c619c3d81c4bed2bbc144506e96d7e69f11523299905e00e.RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return ic294f6a4dcf11553c619c3d81c4bed2bbc144506e96d7e69f11523299905e00e.NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc the reprovisionCloudPc property
func (m *ManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ic7b58bd1f7ec7e59ddafda5fb2b6789ce5abdfcba5b1e47d8b82e90e8daeb055.ReprovisionCloudPcRequestBuilder) {
    return ic7b58bd1f7ec7e59ddafda5fb2b6789ce5abdfcba5b1e47d8b82e90e8daeb055.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance the requestRemoteAssistance property
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ia36af35bae5da4cc0bc0059a8dd1595ea3098bbaf66a723ad498408f9329538d.RequestRemoteAssistanceRequestBuilder) {
    return ia36af35bae5da4cc0bc0059a8dd1595ea3098bbaf66a723ad498408f9329538d.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess the requestRemoteHelpSessionAccess property
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*i19c77245974ed8a6ef9d21e448b4f80b08071e818653a52c1710261948d25db3.RequestRemoteHelpSessionAccessRequestBuilder) {
    return i19c77245974ed8a6ef9d21e448b4f80b08071e818653a52c1710261948d25db3.NewRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode the resetPasscode property
func (m *ManagedDeviceItemRequestBuilder) ResetPasscode()(*i9479492bd537ee0da9d7cbaf1a7bda614d8af9b2db9f340f4d162cfdab45d7d4.ResetPasscodeRequestBuilder) {
    return i9479492bd537ee0da9d7cbaf1a7bda614d8af9b2db9f340f4d162cfdab45d7d4.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc the resizeCloudPc property
func (m *ManagedDeviceItemRequestBuilder) ResizeCloudPc()(*i964b81837112ab4b6bfd1cf6220e1e8cb634c0c56d1238300c9c00340217450d.ResizeCloudPcRequestBuilder) {
    return i964b81837112ab4b6bfd1cf6220e1e8cb634c0c56d1238300c9c00340217450d.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc the restoreCloudPc property
func (m *ManagedDeviceItemRequestBuilder) RestoreCloudPc()(*i5533b65c5e7ef6c9cf517e9d69b5af07f1cec79c9dc96399fcbe89d0ec21a145.RestoreCloudPcRequestBuilder) {
    return i5533b65c5e7ef6c9cf517e9d69b5af07f1cec79c9dc96399fcbe89d0ec21a145.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire the retire property
func (m *ManagedDeviceItemRequestBuilder) Retire()(*if75bb98073f1131ee02501bf098984c52a24d708147db3112ea968e2fc7b2c03.RetireRequestBuilder) {
    return if75bb98073f1131ee02501bf098984c52a24d708147db3112ea968e2fc7b2c03.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ieeb100fe5ba14a8141a3d99450da9262e50bb100dbdddf0192a6178b9a775b45.RetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return ieeb100fe5ba14a8141a3d99450da9262e50bb100dbdddf0192a6178b9a775b45.NewRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses the revokeAppleVppLicenses property
func (m *ManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*i37fecb3bffa962bb92bae6d06b3cd6c4cc85201a2aa5cf7d9d6e15db1e249758.RevokeAppleVppLicensesRequestBuilder) {
    return i37fecb3bffa962bb92bae6d06b3cd6c4cc85201a2aa5cf7d9d6e15db1e249758.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys the rotateBitLockerKeys property
func (m *ManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*i1f1456a44a9b0409729cadca40fbbdf79a42d5a2574cf3d332d1deb5f5f62202.RotateBitLockerKeysRequestBuilder) {
    return i1f1456a44a9b0409729cadca40fbbdf79a42d5a2574cf3d332d1deb5f5f62202.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey the rotateFileVaultKey property
func (m *ManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*i9c36cb050999ff8fab3434eaa458ff65d63d50ae824d7965d7bd33cff08b058a.RotateFileVaultKeyRequestBuilder) {
    return i9c36cb050999ff8fab3434eaa458ff65d63d50ae824d7965d7bd33cff08b058a.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates the securityBaselineStates property
func (m *ManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ie1efb01c1dca39b9cc510aab2614914c0ae675e30562a196406520ea8c0163e9.SecurityBaselineStatesRequestBuilder) {
    return ie1efb01c1dca39b9cc510aab2614914c0ae675e30562a196406520ea8c0163e9.NewSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.securityBaselineStates.item collection
func (m *ManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*i8bf3e9602de99585216828ab1a7ba1b28f64f44b09c1192ca701225a40bdb050.SecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return i8bf3e9602de99585216828ab1a7ba1b28f64f44b09c1192ca701225a40bdb050.NewSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal the sendCustomNotificationToCompanyPortal property
func (m *ManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*iabf61e946f47f041682806eff385ca77fa52f76f6756d068e2e853bebc2c381b.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return iabf61e946f47f041682806eff385ca77fa52f76f6756d068e2e853bebc2c381b.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus the setCloudPcReviewStatus property
func (m *ManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*i39ac59e2082190c94b4885c6ada792a6840b6ab9dcfcf6667e820bf042be62ff.SetCloudPcReviewStatusRequestBuilder) {
    return i39ac59e2082190c94b4885c6ada792a6840b6ab9dcfcf6667e820bf042be62ff.NewSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName the setDeviceName property
func (m *ManagedDeviceItemRequestBuilder) SetDeviceName()(*i1a0e017de38d695b9d703bf70a4161eaea10b82614aa1d2ad5623c980ccc89b4.SetDeviceNameRequestBuilder) {
    return i1a0e017de38d695b9d703bf70a4161eaea10b82614aa1d2ad5623c980ccc89b4.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown the shutDown property
func (m *ManagedDeviceItemRequestBuilder) ShutDown()(*i658174871e5ed8a97ec6404a2a51d6533ca0556946eb0312fbdea2092b5a0356.ShutDownRequestBuilder) {
    return i658174871e5ed8a97ec6404a2a51d6533ca0556946eb0312fbdea2092b5a0356.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice the syncDevice property
func (m *ManagedDeviceItemRequestBuilder) SyncDevice()(*i1e9b4ce64039f5b2e7c5ab1ff6519bc2b54393233bd0de932f805ffa42b56ed5.SyncDeviceRequestBuilder) {
    return i1e9b4ce64039f5b2e7c5ab1ff6519bc2b54393233bd0de932f805ffa42b56ed5.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction the triggerConfigurationManagerAction property
func (m *ManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*i9b07e3d1ad25d8bac7ab04fef250c221091aa5e13eba25491edfadca2f67508b.TriggerConfigurationManagerActionRequestBuilder) {
    return i9b07e3d1ad25d8bac7ab04fef250c221091aa5e13eba25491edfadca2f67508b.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount the updateWindowsDeviceAccount property
func (m *ManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*id1fd9f59ff6efc1a25cf2a7067af3ac8949a4357a777c40891113b5c6f9a3a97.UpdateWindowsDeviceAccountRequestBuilder) {
    return id1fd9f59ff6efc1a25cf2a7067af3ac8949a4357a777c40891113b5c6f9a3a97.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users the users property
func (m *ManagedDeviceItemRequestBuilder) Users()(*i4014a9f77365c024ff8b81b517cc90c5db1deeaee89910692d948bd54130b4ab.UsersRequestBuilder) {
    return i4014a9f77365c024ff8b81b517cc90c5db1deeaee89910692d948bd54130b4ab.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan the windowsDefenderScan property
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*i4566609add3e4bacdac10bdd501bfbc92cf44ca7aa6d3492d3c92cd84f86564a.WindowsDefenderScanRequestBuilder) {
    return i4566609add3e4bacdac10bdd501bfbc92cf44ca7aa6d3492d3c92cd84f86564a.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures the windowsDefenderUpdateSignatures property
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*i2011cd53e380c8be4e418adb3871ddb5cf1719dfb3ca6ee1b408b5211399c258.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i2011cd53e380c8be4e418adb3871ddb5cf1719dfb3ca6ee1b408b5211399c258.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState the windowsProtectionState property
func (m *ManagedDeviceItemRequestBuilder) WindowsProtectionState()(*i08036a752a811a186fe4cfbc410dc173a0482af2aea88d8ef8789f64be1a8a50.WindowsProtectionStateRequestBuilder) {
    return i08036a752a811a186fe4cfbc410dc173a0482af2aea88d8ef8789f64be1a8a50.NewWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe the wipe property
func (m *ManagedDeviceItemRequestBuilder) Wipe()(*i54aaa976711b04a9d75b5a6b447501e988e9c492e4339da93ee652eda4f23f04.WipeRequestBuilder) {
    return i54aaa976711b04a9d75b5a6b447501e988e9c492e4339da93ee652eda4f23f04.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
