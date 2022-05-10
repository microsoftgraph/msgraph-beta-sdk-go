package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03cd5e77ab904b1d9d550349ce2b3ebbf85f8d70d2e8b8b590a0a56fd2015ba7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads"
    i0b4f1c42df0b90df8324fcdc3d05dacfb7ec10ada52b568525afbe1170e2aded "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/acceptedsenders"
    i25eb47ae7697b9319fcd5d83ada2fa21152487eb147c4f64122edb9cefe1f108 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/rejectedsenders"
    i2b9fba376d9e50dec119a2b76f881e63464f58624784dcd38299878025a67804 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations"
    i2bd35fc7ee0e3b47892a6a9ac6681da77c04ed504bf6a27df9be171c04124c82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/transitivemembers"
    i3caabc7558a587078aa576bd406ac8e36462c5315aa787c227383276239ddaf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/owners"
    i4204ea34c102c3ce4c467802be0d53bb1383556f4d557aea2558fbd83b890729 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/planner"
    i42784f80ed0151a738857097c712ee740e2741f7125b03513cf613385e594abd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/photo"
    i439abe9b162df75dfbcaf099bc769c03e0b745c4a81613bf3693d6925425b90f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/memberswithlicenseerrors"
    i5329639eb88f9b6c2d094e3d613ed4be48df1dec22ac3d8578948b2ee7acb55c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/endpoints"
    i61ded72dd2c668133064a60901af9dfbc831e1234dad0d922f613cc63f0383b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/members"
    i6888ba6f35d8df3005b1564656e1139672d00d3b4e233ab6b6e42529d2668a1d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/permissiongrants"
    i948c98f0578a418011e66489ee14819387a29ca363e665bd9e6b6833eb81b2c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/settings"
    i99abdcc4b2a48834af6d3a8a5c763f2bf29900e887ca6322eed504e065e7fb1d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/transitivememberof"
    ia19a14f4592f72c5aaf511a5c4c9acbc58d1a25a9c72e465ea59220150de311d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drive"
    ia99118eee18982eb3f6db993dcb9410d4c9ebb209e60163d299459c0c4c48ff8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview"
    iaa1bb19b42e20c88755fea27d8d715ef075238eaccc0b99aae086b97ff7896f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites"
    ibf9f5bfd81bc51f27c6f7622d2716e416d3bc72bed000aac0c9bd2016cc7c5cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar"
    ic01b4bd542aed5e623d0174ba000503519bcf61ac6b3389f9ded32f59322130d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events"
    ic1345d1238f5c859082f75bf7acc0a381adacf246d31004ba0b9e4058d43cd70 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/memberof"
    ic334f592567eb72efeb140a8a5e6b295a7391e48d8cae7b40d06255b097bd065 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/approleassignments"
    ic43cdf9a1fe896889437915b395ebd4c733205d1cd85ce4e2bc887e2da810246 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/onenote"
    icc85e2d20c639309b9df91028ffd443e68b8282397644136f05d40f84f916c4a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives"
    id26a183b96757f0f6d11c405de6eca1dd7cf10d6690bd0e1cc12caa91ba6c834 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/extensions"
    id9863281ce7434786dbf8d3aed7dadfb704c7e696823b48a1953803552d01591 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team"
    ie9b9807dc102e9e8bc20c1a63bea677882cfe09998f8cf91272ab2f76f986d05 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/photos"
    ieb18de79e0803d7df13d9487a0e984337c598b7917f40133610d54b7d649f9af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/createdonbehalfof"
    if745a4f4d29c82b7e59ca0620c43b6bcb153533eaa27f42da6330c03e33268e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/grouplifecyclepolicies"
    i04b3202de3051930127ef81d4cea3a51f32385860038e4eed860e4f1d243b001 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item"
    i2eacf97c95ddea4544fe96c87b10bd038c801fcbf32d57ddf2d2001e45196cad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item"
    i39481ffa15e7f34e44dab349104ab2c9dadf17f1acc3f66399f23a308355c375 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/transitivemembers/item"
    i3e13d843b2b5dcf7c172975e7e6c70e07428e47378555f5dc9052931e7f51911 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/approleassignments/item"
    i3edb15cc31f308ec20a13f10e6d88af086aed805418c7524773fbbb198364d3c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item"
    i4c58781cd458d83370084da0854ff7857405bf9d17dea39e341b3e527b03c519 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/acceptedsenders/item"
    i54541f9d1495ddde68fef54603ff8af012bef3cc45cfc98824110533e2a46bdb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/rejectedsenders/item"
    i59114913391ee7f202b5ecc0406e29ff8bf10e34e624b2996cc895f437dea025 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/permissiongrants/item"
    i61def864ec04340cf14443e8234491b824fa99358b8244bbb81c25c9c553d3fb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/settings/item"
    i67ca9ce420caa4d4c776ca2f7a165cc531c005c92d712bf00ee32bec9a8df8c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/photos/item"
    i684f66013c693beb648e4cce7495c53967ee914b8f037d5a16d0771798ef25bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/members/item"
    i6cb607d8d53a4c9cca15c31ad1e3b9258f6cc64a89618f5943e81be20558f156 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/memberswithlicenseerrors/item"
    i7d4a66c9c3b30d366b18d4c8aadb1a07d435126a9085c582eaadad9d5c5d7df0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/owners/item"
    ib1a5e9786a7b38da1e6dbac7b45028addc8eb18f0b733663e4b01d6c279a1f5e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/grouplifecyclepolicies/item"
    ib1e9edbe0da07af2fb4bc4517d5cdf027145ab037defb4cb9f726dba5c9317a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item"
    ib30044a9f9f6c1afe0e40c150dcd73a9f6c115f315d27b04fb1cf011aa98a664 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item"
    ib68e2e9468b53cc5e31f64c27dbae6704bdbe981c7033ccdb6422b33a1532aa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/extensions/item"
    ic1022ee2743ae6ddda477185b6305a8e4918467c23583bafe76865f048126d93 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/endpoints/item"
    id33f3bdc65c0bd4a9e54e938eb9f7187bd8eb02b9da59924982d217cbb6ca260 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item"
    iea7a5a3d39b3001b6df2127b934fd0a286ea7f457bfd3812097940b4504ae870 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/transitivememberof/item"
    if60d926a6a6251e945069bf44b6e400fead0ad809f3ee60ea62831804db901fe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/memberof/item"
)

// GroupItemRequestBuilder provides operations to manage the joinedGroups property of the microsoft.graph.user entity.
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
// GroupItemRequestBuilderGetQueryParameters read-only. Nullable.
type GroupItemRequestBuilderGetQueryParameters struct {
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
func (m *GroupItemRequestBuilder) AcceptedSenders()(*i0b4f1c42df0b90df8324fcdc3d05dacfb7ec10ada52b568525afbe1170e2aded.AcceptedSendersRequestBuilder) {
    return i0b4f1c42df0b90df8324fcdc3d05dacfb7ec10ada52b568525afbe1170e2aded.NewAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.acceptedSenders.item collection
func (m *GroupItemRequestBuilder) AcceptedSendersById(id string)(*i4c58781cd458d83370084da0854ff7857405bf9d17dea39e341b3e527b03c519.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i4c58781cd458d83370084da0854ff7857405bf9d17dea39e341b3e527b03c519.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments the appRoleAssignments property
func (m *GroupItemRequestBuilder) AppRoleAssignments()(*ic334f592567eb72efeb140a8a5e6b295a7391e48d8cae7b40d06255b097bd065.AppRoleAssignmentsRequestBuilder) {
    return ic334f592567eb72efeb140a8a5e6b295a7391e48d8cae7b40d06255b097bd065.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.appRoleAssignments.item collection
func (m *GroupItemRequestBuilder) AppRoleAssignmentsById(id string)(*i3e13d843b2b5dcf7c172975e7e6c70e07428e47378555f5dc9052931e7f51911.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return i3e13d843b2b5dcf7c172975e7e6c70e07428e47378555f5dc9052931e7f51911.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *GroupItemRequestBuilder) Calendar()(*ibf9f5bfd81bc51f27c6f7622d2716e416d3bc72bed000aac0c9bd2016cc7c5cf.CalendarRequestBuilder) {
    return ibf9f5bfd81bc51f27c6f7622d2716e416d3bc72bed000aac0c9bd2016cc7c5cf.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *GroupItemRequestBuilder) CalendarView()(*ia99118eee18982eb3f6db993dcb9410d4c9ebb209e60163d299459c0c4c48ff8.CalendarViewRequestBuilder) {
    return ia99118eee18982eb3f6db993dcb9410d4c9ebb209e60163d299459c0c4c48ff8.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item collection
func (m *GroupItemRequestBuilder) CalendarViewById(id string)(*i04b3202de3051930127ef81d4cea3a51f32385860038e4eed860e4f1d243b001.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i04b3202de3051930127ef81d4cea3a51f32385860038e4eed860e4f1d243b001.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    m := &GroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}{?%24select}";
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
func (m *GroupItemRequestBuilder) Conversations()(*i2b9fba376d9e50dec119a2b76f881e63464f58624784dcd38299878025a67804.ConversationsRequestBuilder) {
    return i2b9fba376d9e50dec119a2b76f881e63464f58624784dcd38299878025a67804.NewConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConversationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item collection
func (m *GroupItemRequestBuilder) ConversationsById(id string)(*i2eacf97c95ddea4544fe96c87b10bd038c801fcbf32d57ddf2d2001e45196cad.ConversationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversation%2Did"] = id
    }
    return i2eacf97c95ddea4544fe96c87b10bd038c801fcbf32d57ddf2d2001e45196cad.NewConversationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property joinedGroups for me
func (m *GroupItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property joinedGroups for me
func (m *GroupItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupItemRequestBuilder) CreatedOnBehalfOf()(*ieb18de79e0803d7df13d9487a0e984337c598b7917f40133610d54b7d649f9af.CreatedOnBehalfOfRequestBuilder) {
    return ieb18de79e0803d7df13d9487a0e984337c598b7917f40133610d54b7d649f9af.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation read-only. Nullable.
func (m *GroupItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration read-only. Nullable.
func (m *GroupItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property joinedGroups in me
func (m *GroupItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property joinedGroups in me
func (m *GroupItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property joinedGroups for me
func (m *GroupItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property joinedGroups for me
func (m *GroupItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Drive the drive property
func (m *GroupItemRequestBuilder) Drive()(*ia19a14f4592f72c5aaf511a5c4c9acbc58d1a25a9c72e465ea59220150de311d.DriveRequestBuilder) {
    return ia19a14f4592f72c5aaf511a5c4c9acbc58d1a25a9c72e465ea59220150de311d.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *GroupItemRequestBuilder) Drives()(*icc85e2d20c639309b9df91028ffd443e68b8282397644136f05d40f84f916c4a.DrivesRequestBuilder) {
    return icc85e2d20c639309b9df91028ffd443e68b8282397644136f05d40f84f916c4a.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item collection
func (m *GroupItemRequestBuilder) DrivesById(id string)(*ib1e9edbe0da07af2fb4bc4517d5cdf027145ab037defb4cb9f726dba5c9317a2.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return ib1e9edbe0da07af2fb4bc4517d5cdf027145ab037defb4cb9f726dba5c9317a2.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Endpoints the endpoints property
func (m *GroupItemRequestBuilder) Endpoints()(*i5329639eb88f9b6c2d094e3d613ed4be48df1dec22ac3d8578948b2ee7acb55c.EndpointsRequestBuilder) {
    return i5329639eb88f9b6c2d094e3d613ed4be48df1dec22ac3d8578948b2ee7acb55c.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.endpoints.item collection
func (m *GroupItemRequestBuilder) EndpointsById(id string)(*ic1022ee2743ae6ddda477185b6305a8e4918467c23583bafe76865f048126d93.EndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint%2Did"] = id
    }
    return ic1022ee2743ae6ddda477185b6305a8e4918467c23583bafe76865f048126d93.NewEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events the events property
func (m *GroupItemRequestBuilder) Events()(*ic01b4bd542aed5e623d0174ba000503519bcf61ac6b3389f9ded32f59322130d.EventsRequestBuilder) {
    return ic01b4bd542aed5e623d0174ba000503519bcf61ac6b3389f9ded32f59322130d.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item collection
func (m *GroupItemRequestBuilder) EventsById(id string)(*ib30044a9f9f6c1afe0e40c150dcd73a9f6c115f315d27b04fb1cf011aa98a664.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return ib30044a9f9f6c1afe0e40c150dcd73a9f6c115f315d27b04fb1cf011aa98a664.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *GroupItemRequestBuilder) Extensions()(*id26a183b96757f0f6d11c405de6eca1dd7cf10d6690bd0e1cc12caa91ba6c834.ExtensionsRequestBuilder) {
    return id26a183b96757f0f6d11c405de6eca1dd7cf10d6690bd0e1cc12caa91ba6c834.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.extensions.item collection
func (m *GroupItemRequestBuilder) ExtensionsById(id string)(*ib68e2e9468b53cc5e31f64c27dbae6704bdbe981c7033ccdb6422b33a1532aa0.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ib68e2e9468b53cc5e31f64c27dbae6704bdbe981c7033ccdb6422b33a1532aa0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get read-only. Nullable.
func (m *GroupItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler read-only. Nullable.
func (m *GroupItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// GroupLifecyclePolicies the groupLifecyclePolicies property
func (m *GroupItemRequestBuilder) GroupLifecyclePolicies()(*if745a4f4d29c82b7e59ca0620c43b6bcb153533eaa27f42da6330c03e33268e8.GroupLifecyclePoliciesRequestBuilder) {
    return if745a4f4d29c82b7e59ca0620c43b6bcb153533eaa27f42da6330c03e33268e8.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.groupLifecyclePolicies.item collection
func (m *GroupItemRequestBuilder) GroupLifecyclePoliciesById(id string)(*ib1a5e9786a7b38da1e6dbac7b45028addc8eb18f0b733663e4b01d6c279a1f5e.GroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return ib1a5e9786a7b38da1e6dbac7b45028addc8eb18f0b733663e4b01d6c279a1f5e.NewGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *GroupItemRequestBuilder) MemberOf()(*ic1345d1238f5c859082f75bf7acc0a381adacf246d31004ba0b9e4058d43cd70.MemberOfRequestBuilder) {
    return ic1345d1238f5c859082f75bf7acc0a381adacf246d31004ba0b9e4058d43cd70.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.memberOf.item collection
func (m *GroupItemRequestBuilder) MemberOfById(id string)(*if60d926a6a6251e945069bf44b6e400fead0ad809f3ee60ea62831804db901fe.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return if60d926a6a6251e945069bf44b6e400fead0ad809f3ee60ea62831804db901fe.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *GroupItemRequestBuilder) Members()(*i61ded72dd2c668133064a60901af9dfbc831e1234dad0d922f613cc63f0383b5.MembersRequestBuilder) {
    return i61ded72dd2c668133064a60901af9dfbc831e1234dad0d922f613cc63f0383b5.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.members.item collection
func (m *GroupItemRequestBuilder) MembersById(id string)(*i684f66013c693beb648e4cce7495c53967ee914b8f037d5a16d0771798ef25bd.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i684f66013c693beb648e4cce7495c53967ee914b8f037d5a16d0771798ef25bd.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MembersWithLicenseErrors the membersWithLicenseErrors property
func (m *GroupItemRequestBuilder) MembersWithLicenseErrors()(*i439abe9b162df75dfbcaf099bc769c03e0b745c4a81613bf3693d6925425b90f.MembersWithLicenseErrorsRequestBuilder) {
    return i439abe9b162df75dfbcaf099bc769c03e0b745c4a81613bf3693d6925425b90f.NewMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersWithLicenseErrorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.membersWithLicenseErrors.item collection
func (m *GroupItemRequestBuilder) MembersWithLicenseErrorsById(id string)(*i6cb607d8d53a4c9cca15c31ad1e3b9258f6cc64a89618f5943e81be20558f156.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i6cb607d8d53a4c9cca15c31ad1e3b9258f6cc64a89618f5943e81be20558f156.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *GroupItemRequestBuilder) Onenote()(*ic43cdf9a1fe896889437915b395ebd4c733205d1cd85ce4e2bc887e2da810246.OnenoteRequestBuilder) {
    return ic43cdf9a1fe896889437915b395ebd4c733205d1cd85ce4e2bc887e2da810246.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Owners the owners property
func (m *GroupItemRequestBuilder) Owners()(*i3caabc7558a587078aa576bd406ac8e36462c5315aa787c227383276239ddaf5.OwnersRequestBuilder) {
    return i3caabc7558a587078aa576bd406ac8e36462c5315aa787c227383276239ddaf5.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.owners.item collection
func (m *GroupItemRequestBuilder) OwnersById(id string)(*i7d4a66c9c3b30d366b18d4c8aadb1a07d435126a9085c582eaadad9d5c5d7df0.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i7d4a66c9c3b30d366b18d4c8aadb1a07d435126a9085c582eaadad9d5c5d7df0.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property joinedGroups in me
func (m *GroupItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property joinedGroups in me
func (m *GroupItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PermissionGrants the permissionGrants property
func (m *GroupItemRequestBuilder) PermissionGrants()(*i6888ba6f35d8df3005b1564656e1139672d00d3b4e233ab6b6e42529d2668a1d.PermissionGrantsRequestBuilder) {
    return i6888ba6f35d8df3005b1564656e1139672d00d3b4e233ab6b6e42529d2668a1d.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.permissionGrants.item collection
func (m *GroupItemRequestBuilder) PermissionGrantsById(id string)(*i59114913391ee7f202b5ecc0406e29ff8bf10e34e624b2996cc895f437dea025.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i59114913391ee7f202b5ecc0406e29ff8bf10e34e624b2996cc895f437dea025.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *GroupItemRequestBuilder) Photo()(*i42784f80ed0151a738857097c712ee740e2741f7125b03513cf613385e594abd.PhotoRequestBuilder) {
    return i42784f80ed0151a738857097c712ee740e2741f7125b03513cf613385e594abd.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos the photos property
func (m *GroupItemRequestBuilder) Photos()(*ie9b9807dc102e9e8bc20c1a63bea677882cfe09998f8cf91272ab2f76f986d05.PhotosRequestBuilder) {
    return ie9b9807dc102e9e8bc20c1a63bea677882cfe09998f8cf91272ab2f76f986d05.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.photos.item collection
func (m *GroupItemRequestBuilder) PhotosById(id string)(*i67ca9ce420caa4d4c776ca2f7a165cc531c005c92d712bf00ee32bec9a8df8c2.ProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return i67ca9ce420caa4d4c776ca2f7a165cc531c005c92d712bf00ee32bec9a8df8c2.NewProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner the planner property
func (m *GroupItemRequestBuilder) Planner()(*i4204ea34c102c3ce4c467802be0d53bb1383556f4d557aea2558fbd83b890729.PlannerRequestBuilder) {
    return i4204ea34c102c3ce4c467802be0d53bb1383556f4d557aea2558fbd83b890729.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSenders the rejectedSenders property
func (m *GroupItemRequestBuilder) RejectedSenders()(*i25eb47ae7697b9319fcd5d83ada2fa21152487eb147c4f64122edb9cefe1f108.RejectedSendersRequestBuilder) {
    return i25eb47ae7697b9319fcd5d83ada2fa21152487eb147c4f64122edb9cefe1f108.NewRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.rejectedSenders.item collection
func (m *GroupItemRequestBuilder) RejectedSendersById(id string)(*i54541f9d1495ddde68fef54603ff8af012bef3cc45cfc98824110533e2a46bdb.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i54541f9d1495ddde68fef54603ff8af012bef3cc45cfc98824110533e2a46bdb.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Settings the settings property
func (m *GroupItemRequestBuilder) Settings()(*i948c98f0578a418011e66489ee14819387a29ca363e665bd9e6b6833eb81b2c6.SettingsRequestBuilder) {
    return i948c98f0578a418011e66489ee14819387a29ca363e665bd9e6b6833eb81b2c6.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.settings.item collection
func (m *GroupItemRequestBuilder) SettingsById(id string)(*i61def864ec04340cf14443e8234491b824fa99358b8244bbb81c25c9c553d3fb.DirectorySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directorySetting%2Did"] = id
    }
    return i61def864ec04340cf14443e8234491b824fa99358b8244bbb81c25c9c553d3fb.NewDirectorySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *GroupItemRequestBuilder) Sites()(*iaa1bb19b42e20c88755fea27d8d715ef075238eaccc0b99aae086b97ff7896f2.SitesRequestBuilder) {
    return iaa1bb19b42e20c88755fea27d8d715ef075238eaccc0b99aae086b97ff7896f2.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item collection
func (m *GroupItemRequestBuilder) SitesById(id string)(*i3edb15cc31f308ec20a13f10e6d88af086aed805418c7524773fbbb198364d3c.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return i3edb15cc31f308ec20a13f10e6d88af086aed805418c7524773fbbb198364d3c.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Team the team property
func (m *GroupItemRequestBuilder) Team()(*id9863281ce7434786dbf8d3aed7dadfb704c7e696823b48a1953803552d01591.TeamRequestBuilder) {
    return id9863281ce7434786dbf8d3aed7dadfb704c7e696823b48a1953803552d01591.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Threads the threads property
func (m *GroupItemRequestBuilder) Threads()(*i03cd5e77ab904b1d9d550349ce2b3ebbf85f8d70d2e8b8b590a0a56fd2015ba7.ThreadsRequestBuilder) {
    return i03cd5e77ab904b1d9d550349ce2b3ebbf85f8d70d2e8b8b590a0a56fd2015ba7.NewThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreadsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.threads.item collection
func (m *GroupItemRequestBuilder) ThreadsById(id string)(*id33f3bdc65c0bd4a9e54e938eb9f7187bd8eb02b9da59924982d217cbb6ca260.ConversationThreadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationThread%2Did"] = id
    }
    return id33f3bdc65c0bd4a9e54e938eb9f7187bd8eb02b9da59924982d217cbb6ca260.NewConversationThreadItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *GroupItemRequestBuilder) TransitiveMemberOf()(*i99abdcc4b2a48834af6d3a8a5c763f2bf29900e887ca6322eed504e065e7fb1d.TransitiveMemberOfRequestBuilder) {
    return i99abdcc4b2a48834af6d3a8a5c763f2bf29900e887ca6322eed504e065e7fb1d.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.transitiveMemberOf.item collection
func (m *GroupItemRequestBuilder) TransitiveMemberOfById(id string)(*iea7a5a3d39b3001b6df2127b934fd0a286ea7f457bfd3812097940b4504ae870.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return iea7a5a3d39b3001b6df2127b934fd0a286ea7f457bfd3812097940b4504ae870.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMembers the transitiveMembers property
func (m *GroupItemRequestBuilder) TransitiveMembers()(*i2bd35fc7ee0e3b47892a6a9ac6681da77c04ed504bf6a27df9be171c04124c82.TransitiveMembersRequestBuilder) {
    return i2bd35fc7ee0e3b47892a6a9ac6681da77c04ed504bf6a27df9be171c04124c82.NewTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.transitiveMembers.item collection
func (m *GroupItemRequestBuilder) TransitiveMembersById(id string)(*i39481ffa15e7f34e44dab349104ab2c9dadf17f1acc3f66399f23a308355c375.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i39481ffa15e7f34e44dab349104ab2c9dadf17f1acc3f66399f23a308355c375.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
