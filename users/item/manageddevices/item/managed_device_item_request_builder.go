package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03cb1afc2f67d57d9400b63ecef9792e89e945fd724cde23e589a82704ab7b25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/reenable"
    i04d7afb27c87823a5b8a1150b777eb8f649cfb4b182cc3aaac7ea449dc9db35a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/bypassactivationlock"
    i09952ccb17be64d9f17a7199497738056ec9c19bcbfc405a6be070f42037e6ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/users"
    i0c6cbc7398c9959107ad2940b4f778cba01d65fed0d9113fa2c0334fe226b66a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/restorecloudpc"
    i0fbac5ff689886d9f56359ce526396257d8cedd55091afc91e4354241b0b3745 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/requestremoteassistance"
    i1110055695319f09ead72d0b84c9cc25aee4e201636cf621020af662be621c15 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/createremotehelpsession"
    i1262fc2601f26c6426f134578d76f0d6d82f780c631056a8228b3f0d99e6d860 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/manageddevicemobileappconfigurationstates"
    i1443aae9c03729a2846287e12ec51c60f2fcf8b160378feeeb93a607f140550a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/cleanwindowsdevice"
    i1868d182d5fc9971d12e89f687d24afd220b971669c2a3ca28b0153484943026 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/shutdown"
    i1a2c5f8ed2e88c9f4bbc42d1bfa4b2c5d073f119152a10b4b129dfbe7d2a285d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/getoemwarranty"
    i2b2643552578b57de1982c2b63e86ffdc342052fcd3810982705fe32f7ff4917 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/syncdevice"
    i33634be37bdeb06eea71e0090c6a79f402d72ef80000527b54246bdca02cebf1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/disable"
    i39cdf1e9f523efcaeac5364e8b420fd9fb42101199a12dd136dbd692cf97b084 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/reprovisioncloudpc"
    i3a22286aea8615d29e6227b44400fd41bca6f05137314e12ee291cfa32204bf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/getcloudpcremoteactionresults"
    i3bbd017aa7e7ee20890cf1da02d15ad98b393604067695a9261d601bc9f2dfbf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/overridecompliancestate"
    i3be03389bc06ed5635fc8bb332ea7b8376890c1b17e6e780bf622bd3c151ebe3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/recoverpasscode"
    i3c221fb6d1b78f6e8df33278a766126c47d277a9042f7c9ce09b1f3657d0732d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/triggerconfigurationmanageraction"
    i4101fa62c2c2e4d2e31497b0925eed3a6eb3e00833404994f24ad11e68f6bd63 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/getcloudpcreviewstatus"
    i48206b2710f4ac70b602ca49d43653c4a8ac350b1003c4ab8346019e970b1efa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/resizecloudpc"
    i4b3d2e939ce45ec79b71a7af66a4168df3c7e28dcfe2bda9d77ee9ec7aa446d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/windowsdefenderupdatesignatures"
    i549908e132e24a18449b95b78d5dc3e006594d67a6603ea62369e7e0e478279a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/devicecategory"
    i5a4f48b5d5324b3c44800ca37257ad3cad162f655a3888d7a91f498075e81edf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/rotatebitlockerkeys"
    i5ae4c7a466c77f89b7aa19644a7526bfa120aae290c0135467b1726a5dcc4b74 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/activatedeviceesim"
    i5eaa1df8338c7e51d6a5e0a3aedd92c65a628004057cd8870c4b80d2b6a42a1f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/devicecompliancepolicystates"
    i5f9dbc1f78df2d6c326b240d3a563359eacd62306ee9277ee070577b47abf63b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/windowsprotectionstate"
    i6ffc1f3fcca02682f232534eba1ff2f06f7fe9dd58da254195bcf0bcfc70003c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/deprovision"
    i717f74abbf6738dffeddfbcdc728ee94299b8c040b0d2f0d14c816ea9507fe77 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/updatewindowsdeviceaccount"
    i75586b46cfc283c36309a2002f1459f7dc8874c921839275c5fa8ac1d99c0343 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/retire"
    i7a83c814817fafd79999afddaf053ec951c91277e23c2bd41a99fba8b5d5ee34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/detectedapps"
    i7d2e9c178f5e0b6fb049b4235e7b1757e955cf9cedeaed593fdd5ac4f1020444 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/disablelostmode"
    i8c5e1230acb79d14f8042d28df07b27768b39cfb9125df2a16a3c6d1c398721e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/resetpasscode"
    i96030aea427e143875b290dc27d979fb70e39fce18323f6e08c6b18c63a5e752 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/logoutsharedappledeviceactiveuser"
    i99dec72ee11b01609d1196b9050a524115dc380fdce7b3b78d8e55baed22b886 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/assignmentfilterevaluationstatusdetails"
    i9e3c4371f99649398ab300d2008599e2276df3ba3799e553549f49a0d01216b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/getnoncompliantsettings"
    ia2b4412b80e65ed62eac30410676ee7e92fefcdb6749f1d1bee939ae19dc3342 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/requestremotehelpsessionaccess"
    ia799811f472a99f7f84a8da5854bc239c533c5d488c3d3542056f7e3155307de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/wipe"
    ib3f25970b02da96716057db9a359c312a0f7ba2a78e06f2600b31dadff527acf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/retrieveremotehelpsessionwithsessionkey"
    ib5439e6eec54aeb435296ae381016c23f9ae7c99bdca955a9e2cedfd3cdef69b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/deleteuserfromsharedappledevice"
    ib66cb64667854c2a0e7d0db75f68649e40b4c02fcac6575b6c9ccf277ab9effe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/createdevicelogcollectionrequest"
    ib95cc79f1b97000080dc32d7dd293c1f2ce2e44aecf96ac88921df3933e4134f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/logcollectionrequests"
    ib978621d4053b29aa4150a55a291f5b5a274a7eb361e04d0c30e09718fa3c7c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/setcloudpcreviewstatus"
    ic0864c6659f8ebf3d0509681f532da9ea6035431f79af35c788e47301ff47ec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/rebootnow"
    ic0d3d817ae531275a13e069c0b36d3a00e8760c3936ab3ad8e37e1fe6fcc985c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/locatedevice"
    ic2143d62d7de26fc53c10ab4c7bfa00666977ee02aa20b3753f7ebcd4c577ac0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/removedevicefirmwareconfigurationinterfacemanagement"
    ic81852c6f9b456991db784c2f5639da1dd538a719a88b0a12bff8c52b3ae0452 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/deviceconfigurationstates"
    icf604249bb52d7f630e1cbb61cbd31f8ce8cb16dcf5632e67657e20a8ecb9c4f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/setdevicename"
    id0142781f4ca93aa4ac79852cc6de070be2db4ac59477c0319a59d9b2517f24c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/windowsdefenderscan"
    id0295aec7d10d18caa37f46306bc3f11adff870f41a1bb3f9739c917285dfb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/enablelostmode"
    ide8af58ade8773c75746f4966ffd215d5bf2383d02a45542f74e833909741943 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/remotelock"
    idefab0fc1eb24323095fc70a6c26fc309eebd8620c5b0c8ff512c9d619fec852 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/playlostmodesound"
    if1c46a53b673c294b2a75b58f10dccd776f12d7ce1d09eb316a766dd3c2f3921 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/securitybaselinestates"
    if492591e12af2ad3d3e41f831337d2ca084939e12df6e1b346cf43b89d63fbd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/sendcustomnotificationtocompanyportal"
    if621a81999333fb9e0000f8374424b7789369d44aa439e4a9a613ba7dc6a68d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/revokeapplevpplicenses"
    ifc5d9d9efd260f9d829ab862d7bf5e0f4d2c77bbbc6efcd094b67c08d46039c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/getfilevaultkey"
    ifc94a62de815b7c98b22f78f93384ae9bdf40bc7c1493236253a4c72b66bca18 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/rotatefilevaultkey"
    i155341708cf86beaf312872f82d66abfe65af0c332ff01dd38d70ec1d3c966a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/manageddevicemobileappconfigurationstates/item"
    i22663a63c2777796a397a3f119677747180ea5380ec3cbc46e10bf2f7c1e10c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/assignmentfilterevaluationstatusdetails/item"
    i2a82226506d5f05d37438f9f94f66b38d770f96152002714c86a78a15dac14e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/logcollectionrequests/item"
    i65e95ece1102722d622003107d0aa2ccfe851d7e2fabb973183179f667d18a80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/devicecompliancepolicystates/item"
    i716662bfbf45796ad3e3202624b91ff30f232b287ece48fa85882cba00897d10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/detectedapps/item"
    ic00c9fedb4afbfc7936f03fe8ea7225c855e68f67c7490496a41c8f4dce8bded "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/deviceconfigurationstates/item"
    ie4fe9c1b0775a3ba213a1b61d1fec1140c6cb4da5fc41c5c5997e96597f42ba8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/securitybaselinestates/item"
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
func (m *ManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*i5ae4c7a466c77f89b7aa19644a7526bfa120aae290c0135467b1726a5dcc4b74.ActivateDeviceEsimRequestBuilder) {
    return i5ae4c7a466c77f89b7aa19644a7526bfa120aae290c0135467b1726a5dcc4b74.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails the assignmentFilterEvaluationStatusDetails property
func (m *ManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*i99dec72ee11b01609d1196b9050a524115dc380fdce7b3b78d8e55baed22b886.AssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return i99dec72ee11b01609d1196b9050a524115dc380fdce7b3b78d8e55baed22b886.NewAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.assignmentFilterEvaluationStatusDetails.item collection
func (m *ManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*i22663a63c2777796a397a3f119677747180ea5380ec3cbc46e10bf2f7c1e10c6.AssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return i22663a63c2777796a397a3f119677747180ea5380ec3cbc46e10bf2f7c1e10c6.NewAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock the bypassActivationLock property
func (m *ManagedDeviceItemRequestBuilder) BypassActivationLock()(*i04d7afb27c87823a5b8a1150b777eb8f649cfb4b182cc3aaac7ea449dc9db35a.BypassActivationLockRequestBuilder) {
    return i04d7afb27c87823a5b8a1150b777eb8f649cfb4b182cc3aaac7ea449dc9db35a.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice the cleanWindowsDevice property
func (m *ManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*i1443aae9c03729a2846287e12ec51c60f2fcf8b160378feeeb93a607f140550a.CleanWindowsDeviceRequestBuilder) {
    return i1443aae9c03729a2846287e12ec51c60f2fcf8b160378feeeb93a607f140550a.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    m := &ManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property managedDevices for users
func (m *ManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedDevices for users
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
func (m *ManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ib66cb64667854c2a0e7d0db75f68649e40b4c02fcac6575b6c9ccf277ab9effe.CreateDeviceLogCollectionRequestRequestBuilder) {
    return ib66cb64667854c2a0e7d0db75f68649e40b4c02fcac6575b6c9ccf277ab9effe.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePatchRequestInformation update the navigation property managedDevices in users
func (m *ManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedDevices in users
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
func (m *ManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*i1110055695319f09ead72d0b84c9cc25aee4e201636cf621020af662be621c15.CreateRemoteHelpSessionRequestBuilder) {
    return i1110055695319f09ead72d0b84c9cc25aee4e201636cf621020af662be621c15.NewCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property managedDevices for users
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
func (m *ManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ib5439e6eec54aeb435296ae381016c23f9ae7c99bdca955a9e2cedfd3cdef69b.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return ib5439e6eec54aeb435296ae381016c23f9ae7c99bdca955a9e2cedfd3cdef69b.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision the deprovision property
func (m *ManagedDeviceItemRequestBuilder) Deprovision()(*i6ffc1f3fcca02682f232534eba1ff2f06f7fe9dd58da254195bcf0bcfc70003c.DeprovisionRequestBuilder) {
    return i6ffc1f3fcca02682f232534eba1ff2f06f7fe9dd58da254195bcf0bcfc70003c.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps the detectedApps property
func (m *ManagedDeviceItemRequestBuilder) DetectedApps()(*i7a83c814817fafd79999afddaf053ec951c91277e23c2bd41a99fba8b5d5ee34.DetectedAppsRequestBuilder) {
    return i7a83c814817fafd79999afddaf053ec951c91277e23c2bd41a99fba8b5d5ee34.NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.detectedApps.item collection
func (m *ManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*i716662bfbf45796ad3e3202624b91ff30f232b287ece48fa85882cba00897d10.DetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return i716662bfbf45796ad3e3202624b91ff30f232b287ece48fa85882cba00897d10.NewDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory the deviceCategory property
func (m *ManagedDeviceItemRequestBuilder) DeviceCategory()(*i549908e132e24a18449b95b78d5dc3e006594d67a6603ea62369e7e0e478279a.DeviceCategoryRequestBuilder) {
    return i549908e132e24a18449b95b78d5dc3e006594d67a6603ea62369e7e0e478279a.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates the deviceCompliancePolicyStates property
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*i5eaa1df8338c7e51d6a5e0a3aedd92c65a628004057cd8870c4b80d2b6a42a1f.DeviceCompliancePolicyStatesRequestBuilder) {
    return i5eaa1df8338c7e51d6a5e0a3aedd92c65a628004057cd8870c4b80d2b6a42a1f.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.deviceCompliancePolicyStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i65e95ece1102722d622003107d0aa2ccfe851d7e2fabb973183179f667d18a80.DeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return i65e95ece1102722d622003107d0aa2ccfe851d7e2fabb973183179f667d18a80.NewDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates the deviceConfigurationStates property
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ic81852c6f9b456991db784c2f5639da1dd538a719a88b0a12bff8c52b3ae0452.DeviceConfigurationStatesRequestBuilder) {
    return ic81852c6f9b456991db784c2f5639da1dd538a719a88b0a12bff8c52b3ae0452.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.deviceConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*ic00c9fedb4afbfc7936f03fe8ea7225c855e68f67c7490496a41c8f4dce8bded.DeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return ic00c9fedb4afbfc7936f03fe8ea7225c855e68f67c7490496a41c8f4dce8bded.NewDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable the disable property
func (m *ManagedDeviceItemRequestBuilder) Disable()(*i33634be37bdeb06eea71e0090c6a79f402d72ef80000527b54246bdca02cebf1.DisableRequestBuilder) {
    return i33634be37bdeb06eea71e0090c6a79f402d72ef80000527b54246bdca02cebf1.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode the disableLostMode property
func (m *ManagedDeviceItemRequestBuilder) DisableLostMode()(*i7d2e9c178f5e0b6fb049b4235e7b1757e955cf9cedeaed593fdd5ac4f1020444.DisableLostModeRequestBuilder) {
    return i7d2e9c178f5e0b6fb049b4235e7b1757e955cf9cedeaed593fdd5ac4f1020444.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode the enableLostMode property
func (m *ManagedDeviceItemRequestBuilder) EnableLostMode()(*id0295aec7d10d18caa37f46306bc3f11adff870f41a1bb3f9739c917285dfb0e.EnableLostModeRequestBuilder) {
    return id0295aec7d10d18caa37f46306bc3f11adff870f41a1bb3f9739c917285dfb0e.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*i3a22286aea8615d29e6227b44400fd41bca6f05137314e12ee291cfa32204bf9.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i3a22286aea8615d29e6227b44400fd41bca6f05137314e12ee291cfa32204bf9.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*i4101fa62c2c2e4d2e31497b0925eed3a6eb3e00833404994f24ad11e68f6bd63.GetCloudPcReviewStatusRequestBuilder) {
    return i4101fa62c2c2e4d2e31497b0925eed3a6eb3e00833404994f24ad11e68f6bd63.NewGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ifc5d9d9efd260f9d829ab862d7bf5e0f4d2c77bbbc6efcd094b67c08d46039c0.GetFileVaultKeyRequestBuilder) {
    return ifc5d9d9efd260f9d829ab862d7bf5e0f4d2c77bbbc6efcd094b67c08d46039c0.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*i9e3c4371f99649398ab300d2008599e2276df3ba3799e553549f49a0d01216b1.GetNonCompliantSettingsRequestBuilder) {
    return i9e3c4371f99649398ab300d2008599e2276df3ba3799e553549f49a0d01216b1.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *ManagedDeviceItemRequestBuilder) GetOemWarranty()(*i1a2c5f8ed2e88c9f4bbc42d1bfa4b2c5d073f119152a10b4b129dfbe7d2a285d.GetOemWarrantyRequestBuilder) {
    return i1a2c5f8ed2e88c9f4bbc42d1bfa4b2c5d073f119152a10b4b129dfbe7d2a285d.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice the locateDevice property
func (m *ManagedDeviceItemRequestBuilder) LocateDevice()(*ic0d3d817ae531275a13e069c0b36d3a00e8760c3936ab3ad8e37e1fe6fcc985c.LocateDeviceRequestBuilder) {
    return ic0d3d817ae531275a13e069c0b36d3a00e8760c3936ab3ad8e37e1fe6fcc985c.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests the logCollectionRequests property
func (m *ManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ib95cc79f1b97000080dc32d7dd293c1f2ce2e44aecf96ac88921df3933e4134f.LogCollectionRequestsRequestBuilder) {
    return ib95cc79f1b97000080dc32d7dd293c1f2ce2e44aecf96ac88921df3933e4134f.NewLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.logCollectionRequests.item collection
func (m *ManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*i2a82226506d5f05d37438f9f94f66b38d770f96152002714c86a78a15dac14e9.DeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return i2a82226506d5f05d37438f9f94f66b38d770f96152002714c86a78a15dac14e9.NewDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser the logoutSharedAppleDeviceActiveUser property
func (m *ManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i96030aea427e143875b290dc27d979fb70e39fce18323f6e08c6b18c63a5e752.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i96030aea427e143875b290dc27d979fb70e39fce18323f6e08c6b18c63a5e752.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates the managedDeviceMobileAppConfigurationStates property
func (m *ManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*i1262fc2601f26c6426f134578d76f0d6d82f780c631056a8228b3f0d99e6d860.ManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return i1262fc2601f26c6426f134578d76f0d6d82f780c631056a8228b3f0d99e6d860.NewManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.managedDeviceMobileAppConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*i155341708cf86beaf312872f82d66abfe65af0c332ff01dd38d70ec1d3c966a2.ManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return i155341708cf86beaf312872f82d66abfe65af0c332ff01dd38d70ec1d3c966a2.NewManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState the overrideComplianceState property
func (m *ManagedDeviceItemRequestBuilder) OverrideComplianceState()(*i3bbd017aa7e7ee20890cf1da02d15ad98b393604067695a9261d601bc9f2dfbf.OverrideComplianceStateRequestBuilder) {
    return i3bbd017aa7e7ee20890cf1da02d15ad98b393604067695a9261d601bc9f2dfbf.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in users
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
func (m *ManagedDeviceItemRequestBuilder) PlayLostModeSound()(*idefab0fc1eb24323095fc70a6c26fc309eebd8620c5b0c8ff512c9d619fec852.PlayLostModeSoundRequestBuilder) {
    return idefab0fc1eb24323095fc70a6c26fc309eebd8620c5b0c8ff512c9d619fec852.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow the rebootNow property
func (m *ManagedDeviceItemRequestBuilder) RebootNow()(*ic0864c6659f8ebf3d0509681f532da9ea6035431f79af35c788e47301ff47ec1.RebootNowRequestBuilder) {
    return ic0864c6659f8ebf3d0509681f532da9ea6035431f79af35c788e47301ff47ec1.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode the recoverPasscode property
func (m *ManagedDeviceItemRequestBuilder) RecoverPasscode()(*i3be03389bc06ed5635fc8bb332ea7b8376890c1b17e6e780bf622bd3c151ebe3.RecoverPasscodeRequestBuilder) {
    return i3be03389bc06ed5635fc8bb332ea7b8376890c1b17e6e780bf622bd3c151ebe3.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable the reenable property
func (m *ManagedDeviceItemRequestBuilder) Reenable()(*i03cb1afc2f67d57d9400b63ecef9792e89e945fd724cde23e589a82704ab7b25.ReenableRequestBuilder) {
    return i03cb1afc2f67d57d9400b63ecef9792e89e945fd724cde23e589a82704ab7b25.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock the remoteLock property
func (m *ManagedDeviceItemRequestBuilder) RemoteLock()(*ide8af58ade8773c75746f4966ffd215d5bf2383d02a45542f74e833909741943.RemoteLockRequestBuilder) {
    return ide8af58ade8773c75746f4966ffd215d5bf2383d02a45542f74e833909741943.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement the removeDeviceFirmwareConfigurationInterfaceManagement property
func (m *ManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ic2143d62d7de26fc53c10ab4c7bfa00666977ee02aa20b3753f7ebcd4c577ac0.RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return ic2143d62d7de26fc53c10ab4c7bfa00666977ee02aa20b3753f7ebcd4c577ac0.NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc the reprovisionCloudPc property
func (m *ManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*i39cdf1e9f523efcaeac5364e8b420fd9fb42101199a12dd136dbd692cf97b084.ReprovisionCloudPcRequestBuilder) {
    return i39cdf1e9f523efcaeac5364e8b420fd9fb42101199a12dd136dbd692cf97b084.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance the requestRemoteAssistance property
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*i0fbac5ff689886d9f56359ce526396257d8cedd55091afc91e4354241b0b3745.RequestRemoteAssistanceRequestBuilder) {
    return i0fbac5ff689886d9f56359ce526396257d8cedd55091afc91e4354241b0b3745.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess the requestRemoteHelpSessionAccess property
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*ia2b4412b80e65ed62eac30410676ee7e92fefcdb6749f1d1bee939ae19dc3342.RequestRemoteHelpSessionAccessRequestBuilder) {
    return ia2b4412b80e65ed62eac30410676ee7e92fefcdb6749f1d1bee939ae19dc3342.NewRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode the resetPasscode property
func (m *ManagedDeviceItemRequestBuilder) ResetPasscode()(*i8c5e1230acb79d14f8042d28df07b27768b39cfb9125df2a16a3c6d1c398721e.ResetPasscodeRequestBuilder) {
    return i8c5e1230acb79d14f8042d28df07b27768b39cfb9125df2a16a3c6d1c398721e.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc the resizeCloudPc property
func (m *ManagedDeviceItemRequestBuilder) ResizeCloudPc()(*i48206b2710f4ac70b602ca49d43653c4a8ac350b1003c4ab8346019e970b1efa.ResizeCloudPcRequestBuilder) {
    return i48206b2710f4ac70b602ca49d43653c4a8ac350b1003c4ab8346019e970b1efa.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc the restoreCloudPc property
func (m *ManagedDeviceItemRequestBuilder) RestoreCloudPc()(*i0c6cbc7398c9959107ad2940b4f778cba01d65fed0d9113fa2c0334fe226b66a.RestoreCloudPcRequestBuilder) {
    return i0c6cbc7398c9959107ad2940b4f778cba01d65fed0d9113fa2c0334fe226b66a.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire the retire property
func (m *ManagedDeviceItemRequestBuilder) Retire()(*i75586b46cfc283c36309a2002f1459f7dc8874c921839275c5fa8ac1d99c0343.RetireRequestBuilder) {
    return i75586b46cfc283c36309a2002f1459f7dc8874c921839275c5fa8ac1d99c0343.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ib3f25970b02da96716057db9a359c312a0f7ba2a78e06f2600b31dadff527acf.RetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return ib3f25970b02da96716057db9a359c312a0f7ba2a78e06f2600b31dadff527acf.NewRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses the revokeAppleVppLicenses property
func (m *ManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*if621a81999333fb9e0000f8374424b7789369d44aa439e4a9a613ba7dc6a68d2.RevokeAppleVppLicensesRequestBuilder) {
    return if621a81999333fb9e0000f8374424b7789369d44aa439e4a9a613ba7dc6a68d2.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys the rotateBitLockerKeys property
func (m *ManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*i5a4f48b5d5324b3c44800ca37257ad3cad162f655a3888d7a91f498075e81edf.RotateBitLockerKeysRequestBuilder) {
    return i5a4f48b5d5324b3c44800ca37257ad3cad162f655a3888d7a91f498075e81edf.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey the rotateFileVaultKey property
func (m *ManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ifc94a62de815b7c98b22f78f93384ae9bdf40bc7c1493236253a4c72b66bca18.RotateFileVaultKeyRequestBuilder) {
    return ifc94a62de815b7c98b22f78f93384ae9bdf40bc7c1493236253a4c72b66bca18.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates the securityBaselineStates property
func (m *ManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*if1c46a53b673c294b2a75b58f10dccd776f12d7ce1d09eb316a766dd3c2f3921.SecurityBaselineStatesRequestBuilder) {
    return if1c46a53b673c294b2a75b58f10dccd776f12d7ce1d09eb316a766dd3c2f3921.NewSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.securityBaselineStates.item collection
func (m *ManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*ie4fe9c1b0775a3ba213a1b61d1fec1140c6cb4da5fc41c5c5997e96597f42ba8.SecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return ie4fe9c1b0775a3ba213a1b61d1fec1140c6cb4da5fc41c5c5997e96597f42ba8.NewSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal the sendCustomNotificationToCompanyPortal property
func (m *ManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*if492591e12af2ad3d3e41f831337d2ca084939e12df6e1b346cf43b89d63fbd6.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return if492591e12af2ad3d3e41f831337d2ca084939e12df6e1b346cf43b89d63fbd6.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus the setCloudPcReviewStatus property
func (m *ManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ib978621d4053b29aa4150a55a291f5b5a274a7eb361e04d0c30e09718fa3c7c5.SetCloudPcReviewStatusRequestBuilder) {
    return ib978621d4053b29aa4150a55a291f5b5a274a7eb361e04d0c30e09718fa3c7c5.NewSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName the setDeviceName property
func (m *ManagedDeviceItemRequestBuilder) SetDeviceName()(*icf604249bb52d7f630e1cbb61cbd31f8ce8cb16dcf5632e67657e20a8ecb9c4f.SetDeviceNameRequestBuilder) {
    return icf604249bb52d7f630e1cbb61cbd31f8ce8cb16dcf5632e67657e20a8ecb9c4f.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown the shutDown property
func (m *ManagedDeviceItemRequestBuilder) ShutDown()(*i1868d182d5fc9971d12e89f687d24afd220b971669c2a3ca28b0153484943026.ShutDownRequestBuilder) {
    return i1868d182d5fc9971d12e89f687d24afd220b971669c2a3ca28b0153484943026.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice the syncDevice property
func (m *ManagedDeviceItemRequestBuilder) SyncDevice()(*i2b2643552578b57de1982c2b63e86ffdc342052fcd3810982705fe32f7ff4917.SyncDeviceRequestBuilder) {
    return i2b2643552578b57de1982c2b63e86ffdc342052fcd3810982705fe32f7ff4917.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction the triggerConfigurationManagerAction property
func (m *ManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*i3c221fb6d1b78f6e8df33278a766126c47d277a9042f7c9ce09b1f3657d0732d.TriggerConfigurationManagerActionRequestBuilder) {
    return i3c221fb6d1b78f6e8df33278a766126c47d277a9042f7c9ce09b1f3657d0732d.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount the updateWindowsDeviceAccount property
func (m *ManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*i717f74abbf6738dffeddfbcdc728ee94299b8c040b0d2f0d14c816ea9507fe77.UpdateWindowsDeviceAccountRequestBuilder) {
    return i717f74abbf6738dffeddfbcdc728ee94299b8c040b0d2f0d14c816ea9507fe77.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users the users property
func (m *ManagedDeviceItemRequestBuilder) Users()(*i09952ccb17be64d9f17a7199497738056ec9c19bcbfc405a6be070f42037e6ac.UsersRequestBuilder) {
    return i09952ccb17be64d9f17a7199497738056ec9c19bcbfc405a6be070f42037e6ac.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan the windowsDefenderScan property
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*id0142781f4ca93aa4ac79852cc6de070be2db4ac59477c0319a59d9b2517f24c.WindowsDefenderScanRequestBuilder) {
    return id0142781f4ca93aa4ac79852cc6de070be2db4ac59477c0319a59d9b2517f24c.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures the windowsDefenderUpdateSignatures property
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*i4b3d2e939ce45ec79b71a7af66a4168df3c7e28dcfe2bda9d77ee9ec7aa446d1.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i4b3d2e939ce45ec79b71a7af66a4168df3c7e28dcfe2bda9d77ee9ec7aa446d1.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState the windowsProtectionState property
func (m *ManagedDeviceItemRequestBuilder) WindowsProtectionState()(*i5f9dbc1f78df2d6c326b240d3a563359eacd62306ee9277ee070577b47abf63b.WindowsProtectionStateRequestBuilder) {
    return i5f9dbc1f78df2d6c326b240d3a563359eacd62306ee9277ee070577b47abf63b.NewWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe the wipe property
func (m *ManagedDeviceItemRequestBuilder) Wipe()(*ia799811f472a99f7f84a8da5854bc239c533c5d488c3d3542056f7e3155307de.WipeRequestBuilder) {
    return ia799811f472a99f7f84a8da5854bc239c533c5d488c3d3542056f7e3155307de.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
