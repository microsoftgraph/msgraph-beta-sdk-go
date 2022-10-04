package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i003dc15aa14498f60d98c0b7a47742a50027cefb258ba1c36d156092e60088ec "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/checkgrantedpermissionsforapp"
    i02fee0fcbb14d5af244cac0c2b5cc69f5aa82b85ec7d224e6dbae566d1646d5d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/acceptedsenders"
    i1d83a3478e3b6f201c24080c8d89208808c03b3d48283f162dd6366fb599966a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/planner"
    i1e180c98492486bc886a6593521a76055d900d97075423582dee1a58714edbbd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/approleassignments"
    i2005bae4c19e73b7a8ae6e53c6eeca1731a42ddc92f15f65cea5b27698152a6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/unsubscribebymail"
    i22f4b02027f167f9f99b20aa4eae4d0871053f60939cceceb5731344f9774944 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events"
    i3afd98d50e381574faac1f866015db15f1e2f6ddf7c88f46ed656c7cdfc5f039 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites"
    i44f2125d842af13a8f7093916e1621e93cdf4c5a20c2ead1b9f39afa6c04fd6d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/photos"
    i456bcb76576cf13cf032372ecb1047d00a01b9a75766cdae2c878c5b9b032274 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof"
    i4dd1e402eb350eebba650b4315bb6cb9b3541976d18f055116cc4ee87e741b11 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/evaluatedynamicmembership"
    i50da4bfa6fe7c38e6f565a09fc06474b9b56b07f730cda76b7630d562a519945 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team"
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
    i00e36e1582d4f23423ff6a9bad7a1f621b9d760c10ea34daa89cf8680ef0acdf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/acceptedsenders/item"
    i15f53920f12385cdf51477c6f7036fdcf921f6ebd032740528ef54c015df9193 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item"
    i1b00a3ee52b67a17dd6db2a8f1625a38323a6135625081fb54b40b44232e89bf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item"
    i20ad8636c32bde1b97aa6dbe46a82bfcc7dec5240d558da351ccb4e7801c7cae "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/photos/item"
    i2a0b85f77a81e8f4f18cb792065ad27157bce3afd57f90f15a90fd83196308f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/permissiongrants/item"
    i40372e5830ab7b777e2a31cda4dd8ad81dac267969ff1e25f7608a8298cc1dad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item"
    i608d9aacc6c4cf35e59e89c155e926bd9789e9ca3e6a2081ddd45db5c36cc760 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item"
    i6151bd928799b25e129a4dfec6b703211c3e9bf7515248a24b14f26bb40215af "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item"
    i7d987f66a62915c3b7f44d0edad9deb15f79e59e58f2d9c5ad9b71bf48f15eec "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item"
    i8d39a67aed0c83a0bbf2fd9b9d94eea9e359661e2a364b2a8708ac4eb19459be "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/rejectedsenders/item"
    i963f882f3f2b0d209b4c91218281e65677ae124fdb171aff79661733b37fc788 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/extensions/item"
    i98092eb6315426df8755d93b1458a713452cb101bc01a6e4f43e716dc0143e8f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/endpoints/item"
    i98f5640044bd75cd223ace6912483df5b3141e60b09af3052b22a942b5548c93 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item"
    i9e89b77447821dbf477a68b13c79a69139c1946756898bdc20b285e001539ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item"
    i9f44ea42f796eb814c4952a992e9dc83e10e462621a158d473fd625c512ef4bc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/approleassignments/item"
    ib7c0afd37a4b7a5f14bbce80fa4fb24e2e0339815e98c21882b0619e5898b56e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/settings/item"
    ibf461a69bb33550b4efa61060c58eb8a2bcb76939c703f24dd1f30ca47fea911 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item"
    id4a401ad94d1a1e42e6580fcd9d3d1f570370ff23b48a112b9e47f17279acf9f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/grouplifecyclepolicies/item"
    if3c1397857745b598974f240f08c291179dcf26394fc5f63c5ff79c8b72d6362 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item"
    if3d2858206557e4f8b00ac2f74098e2775aa239b7ce50e535d3e73ff34555b4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item"
    ife496b0c6e03a5eae3a225783a9684a940a4c1aa359a015613fe0091a2704801 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item"
)

// GroupItemRequestBuilder provides operations to manage the collection of group entities.
type GroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupItemRequestBuilderGetQueryParameters get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query. Because the **group** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **group** instance.
type GroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupItemRequestBuilderGetQueryParameters
}
// GroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptedSenders the acceptedSenders property
func (m *GroupItemRequestBuilder) AcceptedSenders()(*i02fee0fcbb14d5af244cac0c2b5cc69f5aa82b85ec7d224e6dbae566d1646d5d.AcceptedSendersRequestBuilder) {
    return i02fee0fcbb14d5af244cac0c2b5cc69f5aa82b85ec7d224e6dbae566d1646d5d.NewAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.acceptedSenders.item collection
func (m *GroupItemRequestBuilder) AcceptedSendersById(id string)(*i00e36e1582d4f23423ff6a9bad7a1f621b9d760c10ea34daa89cf8680ef0acdf.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i00e36e1582d4f23423ff6a9bad7a1f621b9d760c10ea34daa89cf8680ef0acdf.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AddFavorite the addFavorite property
func (m *GroupItemRequestBuilder) AddFavorite()(*id5be5a063ce4126bc3169c53e0c8a80107c1e8167360b9e5cdd5c02543648a19.AddFavoriteRequestBuilder) {
    return id5be5a063ce4126bc3169c53e0c8a80107c1e8167360b9e5cdd5c02543648a19.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignments the appRoleAssignments property
func (m *GroupItemRequestBuilder) AppRoleAssignments()(*i1e180c98492486bc886a6593521a76055d900d97075423582dee1a58714edbbd.AppRoleAssignmentsRequestBuilder) {
    return i1e180c98492486bc886a6593521a76055d900d97075423582dee1a58714edbbd.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.appRoleAssignments.item collection
func (m *GroupItemRequestBuilder) AppRoleAssignmentsById(id string)(*i9f44ea42f796eb814c4952a992e9dc83e10e462621a158d473fd625c512ef4bc.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return i9f44ea42f796eb814c4952a992e9dc83e10e462621a158d473fd625c512ef4bc.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupItemRequestBuilder) AssignLicense()(*i9395ea35b81079dd9cb84f9333ef55d0d2f4a345e948247f50098d8de82acda3.AssignLicenseRequestBuilder) {
    return i9395ea35b81079dd9cb84f9333ef55d0d2f4a345e948247f50098d8de82acda3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar the calendar property
func (m *GroupItemRequestBuilder) Calendar()(*i79f77916b75b286d0c14bd1d97ba370f6acb69eac471d411f2d363facb3dcfcb.CalendarRequestBuilder) {
    return i79f77916b75b286d0c14bd1d97ba370f6acb69eac471d411f2d363facb3dcfcb.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *GroupItemRequestBuilder) CalendarView()(*ia430883e94cb7abce482535ac8c7253ae6b920d048575108445765c30b4d334f.CalendarViewRequestBuilder) {
    return ia430883e94cb7abce482535ac8c7253ae6b920d048575108445765c30b4d334f.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item collection
func (m *GroupItemRequestBuilder) CalendarViewById(id string)(*i9e89b77447821dbf477a68b13c79a69139c1946756898bdc20b285e001539ea1.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i9e89b77447821dbf477a68b13c79a69139c1946756898bdc20b285e001539ea1.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupItemRequestBuilder) CheckGrantedPermissionsForApp()(*i003dc15aa14498f60d98c0b7a47742a50027cefb258ba1c36d156092e60088ec.CheckGrantedPermissionsForAppRequestBuilder) {
    return i003dc15aa14498f60d98c0b7a47742a50027cefb258ba1c36d156092e60088ec.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupItemRequestBuilder) CheckMemberGroups()(*if5b292810372bc43cee66c49bbbeedcfd82fc96959f410656ad93595b60b4235.CheckMemberGroupsRequestBuilder) {
    return if5b292810372bc43cee66c49bbbeedcfd82fc96959f410656ad93595b60b4235.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupItemRequestBuilder) CheckMemberObjects()(*i89e2b1f1f516d9a68cf412b1affa3e4875236e1ec9219fe013ea8cd481acff30.CheckMemberObjectsRequestBuilder) {
    return i89e2b1f1f516d9a68cf412b1affa3e4875236e1ec9219fe013ea8cd481acff30.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    m := &GroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Conversations the conversations property
func (m *GroupItemRequestBuilder) Conversations()(*idca52ceeaf754eed01a2ea52af72bb7be37a2691f73fab6b966a898b32c145e6.ConversationsRequestBuilder) {
    return idca52ceeaf754eed01a2ea52af72bb7be37a2691f73fab6b966a898b32c145e6.NewConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConversationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item collection
func (m *GroupItemRequestBuilder) ConversationsById(id string)(*i7d987f66a62915c3b7f44d0edad9deb15f79e59e58f2d9c5ad9b71bf48f15eec.ConversationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversation%2Did"] = id
    }
    return i7d987f66a62915c3b7f44d0edad9deb15f79e59e58f2d9c5ad9b71bf48f15eec.NewConversationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation deletes a group. When deleted, Microsoft 365 groups are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted. This isn't applicable to Security groups and Distribution groups which are permanently deleted immediately. To learn more, see deletedItems.
func (m *GroupItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatedOnBehalfOf the createdOnBehalfOf property
func (m *GroupItemRequestBuilder) CreatedOnBehalfOf()(*ib725062d5d2e7f70fa299e0ee2b220626e42ba87b03d5d4a6ba5a08f987c9bfe.CreatedOnBehalfOfRequestBuilder) {
    return ib725062d5d2e7f70fa299e0ee2b220626e42ba87b03d5d4a6ba5a08f987c9bfe.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query. Because the **group** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **group** instance.
func (m *GroupItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a group object.
func (m *GroupItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete deletes a group. When deleted, Microsoft 365 groups are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted. This isn't applicable to Security groups and Distribution groups which are permanently deleted immediately. To learn more, see deletedItems.
func (m *GroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Drive the drive property
func (m *GroupItemRequestBuilder) Drive()(*i55a8986bb385256f244980fc278061552d51921385d2ef7794e3d6dc2dd12d0c.DriveRequestBuilder) {
    return i55a8986bb385256f244980fc278061552d51921385d2ef7794e3d6dc2dd12d0c.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *GroupItemRequestBuilder) Drives()(*i8daf117241c7766d7ff29c5cbcabcc082be0444b70bf3b2708dffc23076a87ad.DrivesRequestBuilder) {
    return i8daf117241c7766d7ff29c5cbcabcc082be0444b70bf3b2708dffc23076a87ad.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item collection
func (m *GroupItemRequestBuilder) DrivesById(id string)(*i1b00a3ee52b67a17dd6db2a8f1625a38323a6135625081fb54b40b44232e89bf.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return i1b00a3ee52b67a17dd6db2a8f1625a38323a6135625081fb54b40b44232e89bf.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Endpoints the endpoints property
func (m *GroupItemRequestBuilder) Endpoints()(*i8b52f30152d353621c99bf7bb2121f0421df94ae1ba573498738cca015e14db2.EndpointsRequestBuilder) {
    return i8b52f30152d353621c99bf7bb2121f0421df94ae1ba573498738cca015e14db2.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.endpoints.item collection
func (m *GroupItemRequestBuilder) EndpointsById(id string)(*i98092eb6315426df8755d93b1458a713452cb101bc01a6e4f43e716dc0143e8f.EndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint%2Did"] = id
    }
    return i98092eb6315426df8755d93b1458a713452cb101bc01a6e4f43e716dc0143e8f.NewEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EvaluateDynamicMembership the evaluateDynamicMembership property
func (m *GroupItemRequestBuilder) EvaluateDynamicMembership()(*i4dd1e402eb350eebba650b4315bb6cb9b3541976d18f055116cc4ee87e741b11.EvaluateDynamicMembershipRequestBuilder) {
    return i4dd1e402eb350eebba650b4315bb6cb9b3541976d18f055116cc4ee87e741b11.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Events the events property
func (m *GroupItemRequestBuilder) Events()(*i22f4b02027f167f9f99b20aa4eae4d0871053f60939cceceb5731344f9774944.EventsRequestBuilder) {
    return i22f4b02027f167f9f99b20aa4eae4d0871053f60939cceceb5731344f9774944.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item collection
func (m *GroupItemRequestBuilder) EventsById(id string)(*i40372e5830ab7b777e2a31cda4dd8ad81dac267969ff1e25f7608a8298cc1dad.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i40372e5830ab7b777e2a31cda4dd8ad81dac267969ff1e25f7608a8298cc1dad.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *GroupItemRequestBuilder) Extensions()(*i8837566c7fd9bfa1e091bc591d6d52db9c0af6edf0fdc781f5e2f194afeffe0c.ExtensionsRequestBuilder) {
    return i8837566c7fd9bfa1e091bc591d6d52db9c0af6edf0fdc781f5e2f194afeffe0c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.extensions.item collection
func (m *GroupItemRequestBuilder) ExtensionsById(id string)(*i963f882f3f2b0d209b4c91218281e65677ae124fdb171aff79661733b37fc788.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i963f882f3f2b0d209b4c91218281e65677ae124fdb171aff79661733b37fc788.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query. Because the **group** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **group** instance.
func (m *GroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *GroupItemRequestBuilder) GetMemberGroups()(*if7b3a9a428f07b46ebfec8399e6db0dc5b41c70a7d9342046571dd8c49cfd9b6.GetMemberGroupsRequestBuilder) {
    return if7b3a9a428f07b46ebfec8399e6db0dc5b41c70a7d9342046571dd8c49cfd9b6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupItemRequestBuilder) GetMemberObjects()(*i5f5ccb8bcb6b8d2a93b9b86498908af53ad5ea888c300bf590b64d7866088597.GetMemberObjectsRequestBuilder) {
    return i5f5ccb8bcb6b8d2a93b9b86498908af53ad5ea888c300bf590b64d7866088597.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePolicies the groupLifecyclePolicies property
func (m *GroupItemRequestBuilder) GroupLifecyclePolicies()(*id94a57871bbab0456d687ee70bc7c1fd561519d5ccb62b8021883c0d5a0386fb.GroupLifecyclePoliciesRequestBuilder) {
    return id94a57871bbab0456d687ee70bc7c1fd561519d5ccb62b8021883c0d5a0386fb.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.groupLifecyclePolicies.item collection
func (m *GroupItemRequestBuilder) GroupLifecyclePoliciesById(id string)(*id4a401ad94d1a1e42e6580fcd9d3d1f570370ff23b48a112b9e47f17279acf9f.GroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return id4a401ad94d1a1e42e6580fcd9d3d1f570370ff23b48a112b9e47f17279acf9f.NewGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *GroupItemRequestBuilder) MemberOf()(*ia89e1b878890aa299d6e9f09d3a24104bd1d3851534471fabeebd891ee4d3765.MemberOfRequestBuilder) {
    return ia89e1b878890aa299d6e9f09d3a24104bd1d3851534471fabeebd891ee4d3765.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.memberOf.item collection
func (m *GroupItemRequestBuilder) MemberOfById(id string)(*i98f5640044bd75cd223ace6912483df5b3141e60b09af3052b22a942b5548c93.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i98f5640044bd75cd223ace6912483df5b3141e60b09af3052b22a942b5548c93.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *GroupItemRequestBuilder) Members()(*i5b714432a34bff807e883f3a6a44ffaa08898cf27440b3259bf31ca5c9a5fb42.MembersRequestBuilder) {
    return i5b714432a34bff807e883f3a6a44ffaa08898cf27440b3259bf31ca5c9a5fb42.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.members.item collection
func (m *GroupItemRequestBuilder) MembersById(id string)(*if3c1397857745b598974f240f08c291179dcf26394fc5f63c5ff79c8b72d6362.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return if3c1397857745b598974f240f08c291179dcf26394fc5f63c5ff79c8b72d6362.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MembersWithLicenseErrors the membersWithLicenseErrors property
func (m *GroupItemRequestBuilder) MembersWithLicenseErrors()(*id76919b93194ca9fe5980601f9cb2f18adae36958ca7a1dbdb54cafc08718c7a.MembersWithLicenseErrorsRequestBuilder) {
    return id76919b93194ca9fe5980601f9cb2f18adae36958ca7a1dbdb54cafc08718c7a.NewMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersWithLicenseErrorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.membersWithLicenseErrors.item collection
func (m *GroupItemRequestBuilder) MembersWithLicenseErrorsById(id string)(*i6151bd928799b25e129a4dfec6b703211c3e9bf7515248a24b14f26bb40215af.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i6151bd928799b25e129a4dfec6b703211c3e9bf7515248a24b14f26bb40215af.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *GroupItemRequestBuilder) Onenote()(*iff7545973594d37acb39d1b4365ee7b6eb267ff59bd2b1417ba53e9446928c48.OnenoteRequestBuilder) {
    return iff7545973594d37acb39d1b4365ee7b6eb267ff59bd2b1417ba53e9446928c48.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Owners the owners property
func (m *GroupItemRequestBuilder) Owners()(*id7f04eb5e6dbe2b30622df95d1de790ebae357f05ba2293e3e9c6941f337805b.OwnersRequestBuilder) {
    return id7f04eb5e6dbe2b30622df95d1de790ebae357f05ba2293e3e9c6941f337805b.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.owners.item collection
func (m *GroupItemRequestBuilder) OwnersById(id string)(*i608d9aacc6c4cf35e59e89c155e926bd9789e9ca3e6a2081ddd45db5c36cc760.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i608d9aacc6c4cf35e59e89c155e926bd9789e9ca3e6a2081ddd45db5c36cc760.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a group object.
func (m *GroupItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// PermissionGrants the permissionGrants property
func (m *GroupItemRequestBuilder) PermissionGrants()(*ib51c2fbc52f6d50414f9159f3ead9786b621b227a51da00af9354cb6be6fe6b5.PermissionGrantsRequestBuilder) {
    return ib51c2fbc52f6d50414f9159f3ead9786b621b227a51da00af9354cb6be6fe6b5.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.permissionGrants.item collection
func (m *GroupItemRequestBuilder) PermissionGrantsById(id string)(*i2a0b85f77a81e8f4f18cb792065ad27157bce3afd57f90f15a90fd83196308f8.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i2a0b85f77a81e8f4f18cb792065ad27157bce3afd57f90f15a90fd83196308f8.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *GroupItemRequestBuilder) Photo()(*i6391d4a344a566d46875fa0b0152d8ebbfcec62764c10d44a543893f06f3b135.PhotoRequestBuilder) {
    return i6391d4a344a566d46875fa0b0152d8ebbfcec62764c10d44a543893f06f3b135.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos the photos property
func (m *GroupItemRequestBuilder) Photos()(*i44f2125d842af13a8f7093916e1621e93cdf4c5a20c2ead1b9f39afa6c04fd6d.PhotosRequestBuilder) {
    return i44f2125d842af13a8f7093916e1621e93cdf4c5a20c2ead1b9f39afa6c04fd6d.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.photos.item collection
func (m *GroupItemRequestBuilder) PhotosById(id string)(*i20ad8636c32bde1b97aa6dbe46a82bfcc7dec5240d558da351ccb4e7801c7cae.ProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return i20ad8636c32bde1b97aa6dbe46a82bfcc7dec5240d558da351ccb4e7801c7cae.NewProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner the planner property
func (m *GroupItemRequestBuilder) Planner()(*i1d83a3478e3b6f201c24080c8d89208808c03b3d48283f162dd6366fb599966a.PlannerRequestBuilder) {
    return i1d83a3478e3b6f201c24080c8d89208808c03b3d48283f162dd6366fb599966a.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSenders the rejectedSenders property
func (m *GroupItemRequestBuilder) RejectedSenders()(*ia8fbc097a625199f61d4952cd59e78f58c409e05e37ca9a81f3b8a0272d849bc.RejectedSendersRequestBuilder) {
    return ia8fbc097a625199f61d4952cd59e78f58c409e05e37ca9a81f3b8a0272d849bc.NewRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.rejectedSenders.item collection
func (m *GroupItemRequestBuilder) RejectedSendersById(id string)(*i8d39a67aed0c83a0bbf2fd9b9d94eea9e359661e2a364b2a8708ac4eb19459be.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i8d39a67aed0c83a0bbf2fd9b9d94eea9e359661e2a364b2a8708ac4eb19459be.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RemoveFavorite the removeFavorite property
func (m *GroupItemRequestBuilder) RemoveFavorite()(*ifd55f11174a57d75508e4d147d81b76c4bf5b9bec52b3db9b112b7bd634b0949.RemoveFavoriteRequestBuilder) {
    return ifd55f11174a57d75508e4d147d81b76c4bf5b9bec52b3db9b112b7bd634b0949.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupItemRequestBuilder) Renew()(*i97e04fb71bfb479e053612252065865fa313254cd8a02906de5cd864c86e08c2.RenewRequestBuilder) {
    return i97e04fb71bfb479e053612252065865fa313254cd8a02906de5cd864c86e08c2.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupItemRequestBuilder) ResetUnseenCount()(*if3ff5fc8989855f923b84a6d60a86bf95a0f7a494296a52a24c9222053e48e3c.ResetUnseenCountRequestBuilder) {
    return if3ff5fc8989855f923b84a6d60a86bf95a0f7a494296a52a24c9222053e48e3c.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupItemRequestBuilder) Restore()(*i555eb1d39daf7dceb1ba9d79ad628592e7fc2485c12cd5568cefbf712c3fbab6.RestoreRequestBuilder) {
    return i555eb1d39daf7dceb1ba9d79ad628592e7fc2485c12cd5568cefbf712c3fbab6.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *GroupItemRequestBuilder) Settings()(*i5740a9548283ec6d1f78bd07c0e97c4b41045ec41cf82467585964a4625bb670.SettingsRequestBuilder) {
    return i5740a9548283ec6d1f78bd07c0e97c4b41045ec41cf82467585964a4625bb670.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.settings.item collection
func (m *GroupItemRequestBuilder) SettingsById(id string)(*ib7c0afd37a4b7a5f14bbce80fa4fb24e2e0339815e98c21882b0619e5898b56e.DirectorySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directorySetting%2Did"] = id
    }
    return ib7c0afd37a4b7a5f14bbce80fa4fb24e2e0339815e98c21882b0619e5898b56e.NewDirectorySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *GroupItemRequestBuilder) Sites()(*i3afd98d50e381574faac1f866015db15f1e2f6ddf7c88f46ed656c7cdfc5f039.SitesRequestBuilder) {
    return i3afd98d50e381574faac1f866015db15f1e2f6ddf7c88f46ed656c7cdfc5f039.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item collection
func (m *GroupItemRequestBuilder) SitesById(id string)(*ife496b0c6e03a5eae3a225783a9684a940a4c1aa359a015613fe0091a2704801.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return ife496b0c6e03a5eae3a225783a9684a940a4c1aa359a015613fe0091a2704801.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupItemRequestBuilder) SubscribeByMail()(*i8e2cecc6f8d2e060ecf874e852120b3b67871f950b1351386257d6f90d17766e.SubscribeByMailRequestBuilder) {
    return i8e2cecc6f8d2e060ecf874e852120b3b67871f950b1351386257d6f90d17766e.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Team the team property
func (m *GroupItemRequestBuilder) Team()(*i50da4bfa6fe7c38e6f565a09fc06474b9b56b07f730cda76b7630d562a519945.TeamRequestBuilder) {
    return i50da4bfa6fe7c38e6f565a09fc06474b9b56b07f730cda76b7630d562a519945.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Threads the threads property
func (m *GroupItemRequestBuilder) Threads()(*i5c5a2c806ddf65cb48eca8a02a8bb45136bda5cf5cd7de5574ad78f10f604ba2.ThreadsRequestBuilder) {
    return i5c5a2c806ddf65cb48eca8a02a8bb45136bda5cf5cd7de5574ad78f10f604ba2.NewThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreadsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item collection
func (m *GroupItemRequestBuilder) ThreadsById(id string)(*if3d2858206557e4f8b00ac2f74098e2775aa239b7ce50e535d3e73ff34555b4f.ConversationThreadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationThread%2Did"] = id
    }
    return if3d2858206557e4f8b00ac2f74098e2775aa239b7ce50e535d3e73ff34555b4f.NewConversationThreadItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *GroupItemRequestBuilder) TransitiveMemberOf()(*i456bcb76576cf13cf032372ecb1047d00a01b9a75766cdae2c878c5b9b032274.TransitiveMemberOfRequestBuilder) {
    return i456bcb76576cf13cf032372ecb1047d00a01b9a75766cdae2c878c5b9b032274.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.transitiveMemberOf.item collection
func (m *GroupItemRequestBuilder) TransitiveMemberOfById(id string)(*ibf461a69bb33550b4efa61060c58eb8a2bcb76939c703f24dd1f30ca47fea911.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ibf461a69bb33550b4efa61060c58eb8a2bcb76939c703f24dd1f30ca47fea911.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMembers the transitiveMembers property
func (m *GroupItemRequestBuilder) TransitiveMembers()(*id8eff73a0a9f37276faebe6d96437b4ff5a95ad20baa4387c6008514e132cb37.TransitiveMembersRequestBuilder) {
    return id8eff73a0a9f37276faebe6d96437b4ff5a95ad20baa4387c6008514e132cb37.NewTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.transitiveMembers.item collection
func (m *GroupItemRequestBuilder) TransitiveMembersById(id string)(*i15f53920f12385cdf51477c6f7036fdcf921f6ebd032740528ef54c015df9193.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i15f53920f12385cdf51477c6f7036fdcf921f6ebd032740528ef54c015df9193.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupItemRequestBuilder) UnsubscribeByMail()(*i2005bae4c19e73b7a8ae6e53c6eeca1731a42ddc92f15f65cea5b27698152a6a.UnsubscribeByMailRequestBuilder) {
    return i2005bae4c19e73b7a8ae6e53c6eeca1731a42ddc92f15f65cea5b27698152a6a.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupItemRequestBuilder) ValidateProperties()(*ib5cb4c515d70f5792a169843ae6dafc1787c559ab4496f29705a0f0779ad2914.ValidatePropertiesRequestBuilder) {
    return ib5cb4c515d70f5792a169843ae6dafc1787c559ab4496f29705a0f0779ad2914.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
