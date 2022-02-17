package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i003dc15aa14498f60d98c0b7a47742a50027cefb258ba1c36d156092e60088ec "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/checkgrantedpermissionsforapp"
    i02fee0fcbb14d5af244cac0c2b5cc69f5aa82b85ec7d224e6dbae566d1646d5d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/acceptedsenders"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1d83a3478e3b6f201c24080c8d89208808c03b3d48283f162dd6366fb599966a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/planner"
    i1e180c98492486bc886a6593521a76055d900d97075423582dee1a58714edbbd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/approleassignments"
    i2005bae4c19e73b7a8ae6e53c6eeca1731a42ddc92f15f65cea5b27698152a6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/unsubscribebymail"
    i22f4b02027f167f9f99b20aa4eae4d0871053f60939cceceb5731344f9774944 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events"
    i3afd98d50e381574faac1f866015db15f1e2f6ddf7c88f46ed656c7cdfc5f039 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites"
    i44f2125d842af13a8f7093916e1621e93cdf4c5a20c2ead1b9f39afa6c04fd6d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/photos"
    i456bcb76576cf13cf032372ecb1047d00a01b9a75766cdae2c878c5b9b032274 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof"
    i4dd1e402eb350eebba650b4315bb6cb9b3541976d18f055116cc4ee87e741b11 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/evaluatedynamicmembership"
    i50da4bfa6fe7c38e6f565a09fc06474b9b56b07f730cda76b7630d562a519945 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i555eb1d39daf7dceb1ba9d79ad628592e7fc2485c12cd5568cefbf712c3fbab6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/restore"
    i55a8986bb385256f244980fc278061552d51921385d2ef7794e3d6dc2dd12d0c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drive"
    i5740a9548283ec6d1f78bd07c0e97c4b41045ec41cf82467585964a4625bb670 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/settings"
    i5b714432a34bff807e883f3a6a44ffaa08898cf27440b3259bf31ca5c9a5fb42 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members"
    i5c5a2c806ddf65cb48eca8a02a8bb45136bda5cf5cd7de5574ad78f10f604ba2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads"
    i5f5ccb8bcb6b8d2a93b9b86498908af53ad5ea888c300bf590b64d7866088597 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/getmemberobjects"
    i6391d4a344a566d46875fa0b0152d8ebbfcec62764c10d44a543893f06f3b135 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/photo"
    i79f77916b75b286d0c14bd1d97ba370f6acb69eac471d411f2d363facb3dcfcb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar"
    i8837566c7fd9bfa1e091bc591d6d52db9c0af6edf0fdc781f5e2f194afeffe0c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/extensions"
    i89e2b1f1f516d9a68cf412b1affa3e4875236e1ec9219fe013ea8cd481acff30 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/checkmemberobjects"
    i8b52f30152d353621c99bf7bb2121f0421df94ae1ba573498738cca015e14db2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/endpoints"
    i8daf117241c7766d7ff29c5cbcabcc082be0444b70bf3b2708dffc23076a87ad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives"
    i8e2cecc6f8d2e060ecf874e852120b3b67871f950b1351386257d6f90d17766e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/subscribebymail"
    i9395ea35b81079dd9cb84f9333ef55d0d2f4a345e948247f50098d8de82acda3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/assignlicense"
    i97e04fb71bfb479e053612252065865fa313254cd8a02906de5cd864c86e08c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/renew"
    ia430883e94cb7abce482535ac8c7253ae6b920d048575108445765c30b4d334f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview"
    ia89e1b878890aa299d6e9f09d3a24104bd1d3851534471fabeebd891ee4d3765 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof"
    ia8fbc097a625199f61d4952cd59e78f58c409e05e37ca9a81f3b8a0272d849bc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/rejectedsenders"
    ib51c2fbc52f6d50414f9159f3ead9786b621b227a51da00af9354cb6be6fe6b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/permissiongrants"
    ib5cb4c515d70f5792a169843ae6dafc1787c559ab4496f29705a0f0779ad2914 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/validateproperties"
    ib725062d5d2e7f70fa299e0ee2b220626e42ba87b03d5d4a6ba5a08f987c9bfe "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/createdonbehalfof"
    id5be5a063ce4126bc3169c53e0c8a80107c1e8167360b9e5cdd5c02543648a19 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/addfavorite"
    id76919b93194ca9fe5980601f9cb2f18adae36958ca7a1dbdb54cafc08718c7a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors"
    id7f04eb5e6dbe2b30622df95d1de790ebae357f05ba2293e3e9c6941f337805b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners"
    id8eff73a0a9f37276faebe6d96437b4ff5a95ad20baa4387c6008514e132cb37 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers"
    id94a57871bbab0456d687ee70bc7c1fd561519d5ccb62b8021883c0d5a0386fb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/grouplifecyclepolicies"
    idca52ceeaf754eed01a2ea52af72bb7be37a2691f73fab6b966a898b32c145e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations"
    if3ff5fc8989855f923b84a6d60a86bf95a0f7a494296a52a24c9222053e48e3c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/resetunseencount"
    if5b292810372bc43cee66c49bbbeedcfd82fc96959f410656ad93595b60b4235 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/checkmembergroups"
    if7b3a9a428f07b46ebfec8399e6db0dc5b41c70a7d9342046571dd8c49cfd9b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/getmembergroups"
    ifd55f11174a57d75508e4d147d81b76c4bf5b9bec52b3db9b112b7bd634b0949 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/removefavorite"
    iff7545973594d37acb39d1b4365ee7b6eb267ff59bd2b1417ba53e9446928c48 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote"
    i1b00a3ee52b67a17dd6db2a8f1625a38323a6135625081fb54b40b44232e89bf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item"
    i20ad8636c32bde1b97aa6dbe46a82bfcc7dec5240d558da351ccb4e7801c7cae "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/photos/item"
    i2a0b85f77a81e8f4f18cb792065ad27157bce3afd57f90f15a90fd83196308f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/permissiongrants/item"
    i40372e5830ab7b777e2a31cda4dd8ad81dac267969ff1e25f7608a8298cc1dad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item"
    i7d987f66a62915c3b7f44d0edad9deb15f79e59e58f2d9c5ad9b71bf48f15eec "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item"
    i963f882f3f2b0d209b4c91218281e65677ae124fdb171aff79661733b37fc788 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/extensions/item"
    i98092eb6315426df8755d93b1458a713452cb101bc01a6e4f43e716dc0143e8f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/endpoints/item"
    i9e89b77447821dbf477a68b13c79a69139c1946756898bdc20b285e001539ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item"
    i9f44ea42f796eb814c4952a992e9dc83e10e462621a158d473fd625c512ef4bc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/approleassignments/item"
    ib7c0afd37a4b7a5f14bbce80fa4fb24e2e0339815e98c21882b0619e5898b56e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/settings/item"
    id4a401ad94d1a1e42e6580fcd9d3d1f570370ff23b48a112b9e47f17279acf9f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/grouplifecyclepolicies/item"
    if3d2858206557e4f8b00ac2f74098e2775aa239b7ce50e535d3e73ff34555b4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item"
    ife496b0c6e03a5eae3a225783a9684a940a4c1aa359a015613fe0091a2704801 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item"
)

// GroupRequestBuilder builds and executes requests for operations under \groups\{group-id}
type GroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupRequestBuilderDeleteOptions options for Delete
type GroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupRequestBuilderGetOptions options for Get
type GroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupRequestBuilderGetQueryParameters get entity from groups by key
type GroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupRequestBuilderPatchOptions options for Patch
type GroupRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Group;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GroupRequestBuilder) AcceptedSenders()(*i02fee0fcbb14d5af244cac0c2b5cc69f5aa82b85ec7d224e6dbae566d1646d5d.AcceptedSendersRequestBuilder) {
    return i02fee0fcbb14d5af244cac0c2b5cc69f5aa82b85ec7d224e6dbae566d1646d5d.NewAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) AddFavorite()(*id5be5a063ce4126bc3169c53e0c8a80107c1e8167360b9e5cdd5c02543648a19.AddFavoriteRequestBuilder) {
    return id5be5a063ce4126bc3169c53e0c8a80107c1e8167360b9e5cdd5c02543648a19.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) AppRoleAssignments()(*i1e180c98492486bc886a6593521a76055d900d97075423582dee1a58714edbbd.AppRoleAssignmentsRequestBuilder) {
    return i1e180c98492486bc886a6593521a76055d900d97075423582dee1a58714edbbd.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.appRoleAssignments.item collection
func (m *GroupRequestBuilder) AppRoleAssignmentsById(id string)(*i9f44ea42f796eb814c4952a992e9dc83e10e462621a158d473fd625c512ef4bc.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return i9f44ea42f796eb814c4952a992e9dc83e10e462621a158d473fd625c512ef4bc.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) AssignLicense()(*i9395ea35b81079dd9cb84f9333ef55d0d2f4a345e948247f50098d8de82acda3.AssignLicenseRequestBuilder) {
    return i9395ea35b81079dd9cb84f9333ef55d0d2f4a345e948247f50098d8de82acda3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Calendar()(*i79f77916b75b286d0c14bd1d97ba370f6acb69eac471d411f2d363facb3dcfcb.CalendarRequestBuilder) {
    return i79f77916b75b286d0c14bd1d97ba370f6acb69eac471d411f2d363facb3dcfcb.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CalendarView()(*ia430883e94cb7abce482535ac8c7253ae6b920d048575108445765c30b4d334f.CalendarViewRequestBuilder) {
    return ia430883e94cb7abce482535ac8c7253ae6b920d048575108445765c30b4d334f.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item collection
func (m *GroupRequestBuilder) CalendarViewById(id string)(*i9e89b77447821dbf477a68b13c79a69139c1946756898bdc20b285e001539ea1.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i9e89b77447821dbf477a68b13c79a69139c1946756898bdc20b285e001539ea1.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i003dc15aa14498f60d98c0b7a47742a50027cefb258ba1c36d156092e60088ec.CheckGrantedPermissionsForAppRequestBuilder) {
    return i003dc15aa14498f60d98c0b7a47742a50027cefb258ba1c36d156092e60088ec.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CheckMemberGroups()(*if5b292810372bc43cee66c49bbbeedcfd82fc96959f410656ad93595b60b4235.CheckMemberGroupsRequestBuilder) {
    return if5b292810372bc43cee66c49bbbeedcfd82fc96959f410656ad93595b60b4235.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) CheckMemberObjects()(*i89e2b1f1f516d9a68cf412b1affa3e4875236e1ec9219fe013ea8cd481acff30.CheckMemberObjectsRequestBuilder) {
    return i89e2b1f1f516d9a68cf412b1affa3e4875236e1ec9219fe013ea8cd481acff30.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupRequestBuilder) Conversations()(*idca52ceeaf754eed01a2ea52af72bb7be37a2691f73fab6b966a898b32c145e6.ConversationsRequestBuilder) {
    return idca52ceeaf754eed01a2ea52af72bb7be37a2691f73fab6b966a898b32c145e6.NewConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConversationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item collection
func (m *GroupRequestBuilder) ConversationsById(id string)(*i7d987f66a62915c3b7f44d0edad9deb15f79e59e58f2d9c5ad9b71bf48f15eec.ConversationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversation_id"] = id
    }
    return i7d987f66a62915c3b7f44d0edad9deb15f79e59e58f2d9c5ad9b71bf48f15eec.NewConversationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from groups
func (m *GroupRequestBuilder) CreateDeleteRequestInformation(options *GroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupRequestBuilder) CreatedOnBehalfOf()(*ib725062d5d2e7f70fa299e0ee2b220626e42ba87b03d5d4a6ba5a08f987c9bfe.CreatedOnBehalfOfRequestBuilder) {
    return ib725062d5d2e7f70fa299e0ee2b220626e42ba87b03d5d4a6ba5a08f987c9bfe.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entity from groups by key
func (m *GroupRequestBuilder) CreateGetRequestInformation(options *GroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in groups
func (m *GroupRequestBuilder) CreatePatchRequestInformation(options *GroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from groups
func (m *GroupRequestBuilder) Delete(options *GroupRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *GroupRequestBuilder) Drive()(*i55a8986bb385256f244980fc278061552d51921385d2ef7794e3d6dc2dd12d0c.DriveRequestBuilder) {
    return i55a8986bb385256f244980fc278061552d51921385d2ef7794e3d6dc2dd12d0c.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Drives()(*i8daf117241c7766d7ff29c5cbcabcc082be0444b70bf3b2708dffc23076a87ad.DrivesRequestBuilder) {
    return i8daf117241c7766d7ff29c5cbcabcc082be0444b70bf3b2708dffc23076a87ad.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item collection
func (m *GroupRequestBuilder) DrivesById(id string)(*i1b00a3ee52b67a17dd6db2a8f1625a38323a6135625081fb54b40b44232e89bf.DriveRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return i1b00a3ee52b67a17dd6db2a8f1625a38323a6135625081fb54b40b44232e89bf.NewDriveRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Endpoints()(*i8b52f30152d353621c99bf7bb2121f0421df94ae1ba573498738cca015e14db2.EndpointsRequestBuilder) {
    return i8b52f30152d353621c99bf7bb2121f0421df94ae1ba573498738cca015e14db2.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.endpoints.item collection
func (m *GroupRequestBuilder) EndpointsById(id string)(*i98092eb6315426df8755d93b1458a713452cb101bc01a6e4f43e716dc0143e8f.EndpointRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint_id"] = id
    }
    return i98092eb6315426df8755d93b1458a713452cb101bc01a6e4f43e716dc0143e8f.NewEndpointRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) EvaluateDynamicMembership()(*i4dd1e402eb350eebba650b4315bb6cb9b3541976d18f055116cc4ee87e741b11.EvaluateDynamicMembershipRequestBuilder) {
    return i4dd1e402eb350eebba650b4315bb6cb9b3541976d18f055116cc4ee87e741b11.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Events()(*i22f4b02027f167f9f99b20aa4eae4d0871053f60939cceceb5731344f9774944.EventsRequestBuilder) {
    return i22f4b02027f167f9f99b20aa4eae4d0871053f60939cceceb5731344f9774944.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item collection
func (m *GroupRequestBuilder) EventsById(id string)(*i40372e5830ab7b777e2a31cda4dd8ad81dac267969ff1e25f7608a8298cc1dad.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i40372e5830ab7b777e2a31cda4dd8ad81dac267969ff1e25f7608a8298cc1dad.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Extensions()(*i8837566c7fd9bfa1e091bc591d6d52db9c0af6edf0fdc781f5e2f194afeffe0c.ExtensionsRequestBuilder) {
    return i8837566c7fd9bfa1e091bc591d6d52db9c0af6edf0fdc781f5e2f194afeffe0c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.extensions.item collection
func (m *GroupRequestBuilder) ExtensionsById(id string)(*i963f882f3f2b0d209b4c91218281e65677ae124fdb171aff79661733b37fc788.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i963f882f3f2b0d209b4c91218281e65677ae124fdb171aff79661733b37fc788.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from groups by key
func (m *GroupRequestBuilder) Get(options *GroupRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Group, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroup() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Group), nil
}
func (m *GroupRequestBuilder) GetMemberGroups()(*if7b3a9a428f07b46ebfec8399e6db0dc5b41c70a7d9342046571dd8c49cfd9b6.GetMemberGroupsRequestBuilder) {
    return if7b3a9a428f07b46ebfec8399e6db0dc5b41c70a7d9342046571dd8c49cfd9b6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) GetMemberObjects()(*i5f5ccb8bcb6b8d2a93b9b86498908af53ad5ea888c300bf590b64d7866088597.GetMemberObjectsRequestBuilder) {
    return i5f5ccb8bcb6b8d2a93b9b86498908af53ad5ea888c300bf590b64d7866088597.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) GroupLifecyclePolicies()(*id94a57871bbab0456d687ee70bc7c1fd561519d5ccb62b8021883c0d5a0386fb.GroupLifecyclePoliciesRequestBuilder) {
    return id94a57871bbab0456d687ee70bc7c1fd561519d5ccb62b8021883c0d5a0386fb.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.groupLifecyclePolicies.item collection
func (m *GroupRequestBuilder) GroupLifecyclePoliciesById(id string)(*id4a401ad94d1a1e42e6580fcd9d3d1f570370ff23b48a112b9e47f17279acf9f.GroupLifecyclePolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy_id"] = id
    }
    return id4a401ad94d1a1e42e6580fcd9d3d1f570370ff23b48a112b9e47f17279acf9f.NewGroupLifecyclePolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) MemberOf()(*ia89e1b878890aa299d6e9f09d3a24104bd1d3851534471fabeebd891ee4d3765.MemberOfRequestBuilder) {
    return ia89e1b878890aa299d6e9f09d3a24104bd1d3851534471fabeebd891ee4d3765.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Members()(*i5b714432a34bff807e883f3a6a44ffaa08898cf27440b3259bf31ca5c9a5fb42.MembersRequestBuilder) {
    return i5b714432a34bff807e883f3a6a44ffaa08898cf27440b3259bf31ca5c9a5fb42.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) MembersWithLicenseErrors()(*id76919b93194ca9fe5980601f9cb2f18adae36958ca7a1dbdb54cafc08718c7a.MembersWithLicenseErrorsRequestBuilder) {
    return id76919b93194ca9fe5980601f9cb2f18adae36958ca7a1dbdb54cafc08718c7a.NewMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Onenote()(*iff7545973594d37acb39d1b4365ee7b6eb267ff59bd2b1417ba53e9446928c48.OnenoteRequestBuilder) {
    return iff7545973594d37acb39d1b4365ee7b6eb267ff59bd2b1417ba53e9446928c48.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Owners()(*id7f04eb5e6dbe2b30622df95d1de790ebae357f05ba2293e3e9c6941f337805b.OwnersRequestBuilder) {
    return id7f04eb5e6dbe2b30622df95d1de790ebae357f05ba2293e3e9c6941f337805b.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in groups
func (m *GroupRequestBuilder) Patch(options *GroupRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *GroupRequestBuilder) PermissionGrants()(*ib51c2fbc52f6d50414f9159f3ead9786b621b227a51da00af9354cb6be6fe6b5.PermissionGrantsRequestBuilder) {
    return ib51c2fbc52f6d50414f9159f3ead9786b621b227a51da00af9354cb6be6fe6b5.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.permissionGrants.item collection
func (m *GroupRequestBuilder) PermissionGrantsById(id string)(*i2a0b85f77a81e8f4f18cb792065ad27157bce3afd57f90f15a90fd83196308f8.ResourceSpecificPermissionGrantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return i2a0b85f77a81e8f4f18cb792065ad27157bce3afd57f90f15a90fd83196308f8.NewResourceSpecificPermissionGrantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Photo()(*i6391d4a344a566d46875fa0b0152d8ebbfcec62764c10d44a543893f06f3b135.PhotoRequestBuilder) {
    return i6391d4a344a566d46875fa0b0152d8ebbfcec62764c10d44a543893f06f3b135.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Photos()(*i44f2125d842af13a8f7093916e1621e93cdf4c5a20c2ead1b9f39afa6c04fd6d.PhotosRequestBuilder) {
    return i44f2125d842af13a8f7093916e1621e93cdf4c5a20c2ead1b9f39afa6c04fd6d.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.photos.item collection
func (m *GroupRequestBuilder) PhotosById(id string)(*i20ad8636c32bde1b97aa6dbe46a82bfcc7dec5240d558da351ccb4e7801c7cae.ProfilePhotoRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto_id"] = id
    }
    return i20ad8636c32bde1b97aa6dbe46a82bfcc7dec5240d558da351ccb4e7801c7cae.NewProfilePhotoRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Planner()(*i1d83a3478e3b6f201c24080c8d89208808c03b3d48283f162dd6366fb599966a.PlannerRequestBuilder) {
    return i1d83a3478e3b6f201c24080c8d89208808c03b3d48283f162dd6366fb599966a.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) RejectedSenders()(*ia8fbc097a625199f61d4952cd59e78f58c409e05e37ca9a81f3b8a0272d849bc.RejectedSendersRequestBuilder) {
    return ia8fbc097a625199f61d4952cd59e78f58c409e05e37ca9a81f3b8a0272d849bc.NewRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) RemoveFavorite()(*ifd55f11174a57d75508e4d147d81b76c4bf5b9bec52b3db9b112b7bd634b0949.RemoveFavoriteRequestBuilder) {
    return ifd55f11174a57d75508e4d147d81b76c4bf5b9bec52b3db9b112b7bd634b0949.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Renew()(*i97e04fb71bfb479e053612252065865fa313254cd8a02906de5cd864c86e08c2.RenewRequestBuilder) {
    return i97e04fb71bfb479e053612252065865fa313254cd8a02906de5cd864c86e08c2.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) ResetUnseenCount()(*if3ff5fc8989855f923b84a6d60a86bf95a0f7a494296a52a24c9222053e48e3c.ResetUnseenCountRequestBuilder) {
    return if3ff5fc8989855f923b84a6d60a86bf95a0f7a494296a52a24c9222053e48e3c.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Restore()(*i555eb1d39daf7dceb1ba9d79ad628592e7fc2485c12cd5568cefbf712c3fbab6.RestoreRequestBuilder) {
    return i555eb1d39daf7dceb1ba9d79ad628592e7fc2485c12cd5568cefbf712c3fbab6.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Settings()(*i5740a9548283ec6d1f78bd07c0e97c4b41045ec41cf82467585964a4625bb670.SettingsRequestBuilder) {
    return i5740a9548283ec6d1f78bd07c0e97c4b41045ec41cf82467585964a4625bb670.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.settings.item collection
func (m *GroupRequestBuilder) SettingsById(id string)(*ib7c0afd37a4b7a5f14bbce80fa4fb24e2e0339815e98c21882b0619e5898b56e.DirectorySettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directorySetting_id"] = id
    }
    return ib7c0afd37a4b7a5f14bbce80fa4fb24e2e0339815e98c21882b0619e5898b56e.NewDirectorySettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) Sites()(*i3afd98d50e381574faac1f866015db15f1e2f6ddf7c88f46ed656c7cdfc5f039.SitesRequestBuilder) {
    return i3afd98d50e381574faac1f866015db15f1e2f6ddf7c88f46ed656c7cdfc5f039.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item collection
func (m *GroupRequestBuilder) SitesById(id string)(*ife496b0c6e03a5eae3a225783a9684a940a4c1aa359a015613fe0091a2704801.SiteRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site_id"] = id
    }
    return ife496b0c6e03a5eae3a225783a9684a940a4c1aa359a015613fe0091a2704801.NewSiteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) SubscribeByMail()(*i8e2cecc6f8d2e060ecf874e852120b3b67871f950b1351386257d6f90d17766e.SubscribeByMailRequestBuilder) {
    return i8e2cecc6f8d2e060ecf874e852120b3b67871f950b1351386257d6f90d17766e.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Team()(*i50da4bfa6fe7c38e6f565a09fc06474b9b56b07f730cda76b7630d562a519945.TeamRequestBuilder) {
    return i50da4bfa6fe7c38e6f565a09fc06474b9b56b07f730cda76b7630d562a519945.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) Threads()(*i5c5a2c806ddf65cb48eca8a02a8bb45136bda5cf5cd7de5574ad78f10f604ba2.ThreadsRequestBuilder) {
    return i5c5a2c806ddf65cb48eca8a02a8bb45136bda5cf5cd7de5574ad78f10f604ba2.NewThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreadsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item collection
func (m *GroupRequestBuilder) ThreadsById(id string)(*if3d2858206557e4f8b00ac2f74098e2775aa239b7ce50e535d3e73ff34555b4f.ConversationThreadRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationThread_id"] = id
    }
    return if3d2858206557e4f8b00ac2f74098e2775aa239b7ce50e535d3e73ff34555b4f.NewConversationThreadRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupRequestBuilder) TransitiveMemberOf()(*i456bcb76576cf13cf032372ecb1047d00a01b9a75766cdae2c878c5b9b032274.TransitiveMemberOfRequestBuilder) {
    return i456bcb76576cf13cf032372ecb1047d00a01b9a75766cdae2c878c5b9b032274.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) TransitiveMembers()(*id8eff73a0a9f37276faebe6d96437b4ff5a95ad20baa4387c6008514e132cb37.TransitiveMembersRequestBuilder) {
    return id8eff73a0a9f37276faebe6d96437b4ff5a95ad20baa4387c6008514e132cb37.NewTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i2005bae4c19e73b7a8ae6e53c6eeca1731a42ddc92f15f65cea5b27698152a6a.UnsubscribeByMailRequestBuilder) {
    return i2005bae4c19e73b7a8ae6e53c6eeca1731a42ddc92f15f65cea5b27698152a6a.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupRequestBuilder) ValidateProperties()(*ib5cb4c515d70f5792a169843ae6dafc1787c559ab4496f29705a0f0779ad2914.ValidatePropertiesRequestBuilder) {
    return ib5cb4c515d70f5792a169843ae6dafc1787c559ab4496f29705a0f0779ad2914.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
