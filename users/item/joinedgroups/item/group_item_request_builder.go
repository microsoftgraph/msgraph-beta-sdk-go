package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i151bd88b167789c7f379069cae0c5cf3ce7a33c6f1732d0fab52d7c413bd972e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events"
    i2d6466e475d8a636fc8f4f58f807c6b37484ddd1346e6c83507fedded6bfad02 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/approleassignments"
    i2e123d7e0a7dec41d5c6aca2603dd5abf115d35da35c8cb4bb86811ec04fd4c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/endpoints"
    i31f436f071c3daaebe4e21c67c01aeffbbeb0acf42a682d5bd0da1f8265ac096 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/planner"
    i37d1c6edd7c02656ff19812d0f443f3a240b0f736e09d1396fee36797f90c682 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/photo"
    i3e491681dc3c7fc87f42d306fdf4a6db96b4e176b78bf4ded4085d340518a194 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites"
    i445223cb5fb1554cd6f84a16220bdec0fafc528688ea42ea5f437b47439a70ee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/owners"
    i461649a33bb785166c01173e17a1397fef7d397f67457d9ca2744d261a4b9de2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/members"
    i46c907074dd13bc5887834efa148bd776d64cc248c900951c0e1e3478bf387b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team"
    i4caaab3ad40a9992ef480a614da8698cd73481940a876181504765a2d0deecca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/photos"
    i4d8e28a4bd1e09a42bf9015551639d380e947f7264b4d8a99fb14889a91e37c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives"
    i4dc03b0a66d1f721539bf9f15ca7ecb992d2e6bba823cc9dc1c01cdf6e7ced40 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/memberof"
    i525edd8673ece5c6e8f4053107a715d76ccf45601b2905693f14691e528de124 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drive"
    i5360ef4992db51254474fc3906e32cf456835d54312836251cda47eaa2f1661f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/acceptedsenders"
    i553997603462e15023ce8c734341ef27d17795823ae349d86df79c9ff9d8a94b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/extensions"
    i5e515dfd981af8a6e6c00f981dc4feeade359aa186dd9bb35be56a6681ca41a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/createdonbehalfof"
    i5ef5d6109700cf9b74ce0f4735be7684ee06a994dcf750744922b83617c58ccf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/permissiongrants"
    i5f8f7910d54070b0f7e884c8482d350b4b873dc167891fbd89ba73da2e1e0986 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads"
    ia753cb6a12c0b696be3b71817e34405b91e0be88a923be55114d53cf49082434 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/settings"
    ia8fcf9acc88115027a27af5e0383e948c22ee114948e81510e2845e54f30634e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations"
    iaf4233dd516b9cb4ad3f959649d34c1904eb731f852c1ed4cf818b448a6c6616 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview"
    ic294ab9d19fb30cf14f9b35e0a081fcc7e11428608e26db5a8018dc4c7d494ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/rejectedsenders"
    ic46a4dc410467cf4d6acd5973d8b7065d566186d73bab4001ee7357a13824c2a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/grouplifecyclepolicies"
    icf14a4a532b90c50e5f8b7f026c68df019ddb14511768f7b0df1b9ea05c8bd1a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/transitivemembers"
    iead0fb61869eb3cc76adc24a7157c079df7ea6a6160170afab077a54a31066f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/onenote"
    if4e6f600a3514349cc881b5b0b3cb5e742f222ebd09202bb92a87f30c638315e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar"
    ifa1a1fed53a335586bb0dddb85c852c234549b6f6a88aacc698e5b941fb2022e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/memberswithlicenseerrors"
    ifbed89bb7272e35484ac2dc99d2fdbb35cfa6c81abb043f47080800bfc4f0a48 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/transitivememberof"
    i0d900e574447595db6a7d9962c3ff45bb15157f2a9ceda3ca75462393a2a0103 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/rejectedsenders/item"
    i142d97539d5c7f0c4d8140f25af8e17432073cbe3732b414e6aa5a97f9ab50a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/photos/item"
    i1451cf64c5c0761a53bbc749fe6bd91cced323ef43352c4b4d6a3b6eb0beb80e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/approleassignments/item"
    i187529aafe259ee4819b83e6be2d7e6b77dbd743282200549459e07ca4dd3c79 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/acceptedsenders/item"
    i240fc320b71fb1dab200c27cdf37d9ebbf070eaca1406a7175520b78a58ae259 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item"
    i2d68757616ff94d59b32b253f2d16af53ce6394915b9b8fbe88a8a527349871d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/memberof/item"
    i33e227b32f13209673672ae7a38ed6696cc7d828159332bb4ae1f5327022b450 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/members/item"
    i5271a24d5f906cc3b6f65e88d1e01af948bd906fc6c480ab04e7fb6688f27534 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item"
    i55c56d558062f71f1446b14e68e829cc344b4970552daa8a61c6fb50debf2430 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item"
    i596929c19e65628c2d282a6486262d8d20bcb6dd817073a9ccc8fccfadb48cf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/endpoints/item"
    i61fc5189c26f38bdac1e13e87931337652e594efe0d92f2a10203cd6732f32df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/permissiongrants/item"
    i63c21f0d1e989b0923cde33368f86d94e7d49586d39585b403db033fbf6714b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/extensions/item"
    i65767a283d1bd8e37ad589483567950ffe4d54fec4319f8d2cc8957f634f5c9f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/owners/item"
    i7fbcdd8dbde4be9d9f484459dd838f2fb66ca0ba35a2c3334cd07fec9bcbf936 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/memberswithlicenseerrors/item"
    i81fc61d562fa2c0a2d5ac39c21b3b8f91931def1c191e3eeb01afbdde27a5ae3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/transitivemembers/item"
    i83c11cd83fc9d7781fbc5c92495b2b6a7d532c21b1630a06f79b248afbd23fef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item"
    i912942191b304ee201aeaa8ea87cca5a1101a458fbccbb508c42c75a57e5e13a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/transitivememberof/item"
    i941e71faec83df85c828a4d8333dcaf30d8e831292c063cbb43bed311b59cc93 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item"
    if0eea92e7b4aa17a8cd5570652467567e700ce471d79115f32dd029e959e4332 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/grouplifecyclepolicies/item"
    ifac9bb0df9a76a9a7e01bb81905add19d789c6bde1441edd6687fc8b575e2682 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item"
    ifc347c65375f4761c6131dd0524e19090cde8f5b20784aec198b784677b47203 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/settings/item"
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
func (m *GroupItemRequestBuilder) AcceptedSenders()(*i5360ef4992db51254474fc3906e32cf456835d54312836251cda47eaa2f1661f.AcceptedSendersRequestBuilder) {
    return i5360ef4992db51254474fc3906e32cf456835d54312836251cda47eaa2f1661f.NewAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.acceptedSenders.item collection
func (m *GroupItemRequestBuilder) AcceptedSendersById(id string)(*i187529aafe259ee4819b83e6be2d7e6b77dbd743282200549459e07ca4dd3c79.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i187529aafe259ee4819b83e6be2d7e6b77dbd743282200549459e07ca4dd3c79.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments the appRoleAssignments property
func (m *GroupItemRequestBuilder) AppRoleAssignments()(*i2d6466e475d8a636fc8f4f58f807c6b37484ddd1346e6c83507fedded6bfad02.AppRoleAssignmentsRequestBuilder) {
    return i2d6466e475d8a636fc8f4f58f807c6b37484ddd1346e6c83507fedded6bfad02.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.appRoleAssignments.item collection
func (m *GroupItemRequestBuilder) AppRoleAssignmentsById(id string)(*i1451cf64c5c0761a53bbc749fe6bd91cced323ef43352c4b4d6a3b6eb0beb80e.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return i1451cf64c5c0761a53bbc749fe6bd91cced323ef43352c4b4d6a3b6eb0beb80e.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *GroupItemRequestBuilder) Calendar()(*if4e6f600a3514349cc881b5b0b3cb5e742f222ebd09202bb92a87f30c638315e.CalendarRequestBuilder) {
    return if4e6f600a3514349cc881b5b0b3cb5e742f222ebd09202bb92a87f30c638315e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *GroupItemRequestBuilder) CalendarView()(*iaf4233dd516b9cb4ad3f959649d34c1904eb731f852c1ed4cf818b448a6c6616.CalendarViewRequestBuilder) {
    return iaf4233dd516b9cb4ad3f959649d34c1904eb731f852c1ed4cf818b448a6c6616.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item collection
func (m *GroupItemRequestBuilder) CalendarViewById(id string)(*i55c56d558062f71f1446b14e68e829cc344b4970552daa8a61c6fb50debf2430.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i55c56d558062f71f1446b14e68e829cc344b4970552daa8a61c6fb50debf2430.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    m := &GroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}{?%24select}";
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
func (m *GroupItemRequestBuilder) Conversations()(*ia8fcf9acc88115027a27af5e0383e948c22ee114948e81510e2845e54f30634e.ConversationsRequestBuilder) {
    return ia8fcf9acc88115027a27af5e0383e948c22ee114948e81510e2845e54f30634e.NewConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConversationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item collection
func (m *GroupItemRequestBuilder) ConversationsById(id string)(*i5271a24d5f906cc3b6f65e88d1e01af948bd906fc6c480ab04e7fb6688f27534.ConversationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversation%2Did"] = id
    }
    return i5271a24d5f906cc3b6f65e88d1e01af948bd906fc6c480ab04e7fb6688f27534.NewConversationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property joinedGroups for users
func (m *GroupItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property joinedGroups for users
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
func (m *GroupItemRequestBuilder) CreatedOnBehalfOf()(*i5e515dfd981af8a6e6c00f981dc4feeade359aa186dd9bb35be56a6681ca41a7.CreatedOnBehalfOfRequestBuilder) {
    return i5e515dfd981af8a6e6c00f981dc4feeade359aa186dd9bb35be56a6681ca41a7.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePatchRequestInformation update the navigation property joinedGroups in users
func (m *GroupItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property joinedGroups in users
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
// Delete delete navigation property joinedGroups for users
func (m *GroupItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property joinedGroups for users
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
func (m *GroupItemRequestBuilder) Drive()(*i525edd8673ece5c6e8f4053107a715d76ccf45601b2905693f14691e528de124.DriveRequestBuilder) {
    return i525edd8673ece5c6e8f4053107a715d76ccf45601b2905693f14691e528de124.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *GroupItemRequestBuilder) Drives()(*i4d8e28a4bd1e09a42bf9015551639d380e947f7264b4d8a99fb14889a91e37c5.DrivesRequestBuilder) {
    return i4d8e28a4bd1e09a42bf9015551639d380e947f7264b4d8a99fb14889a91e37c5.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item collection
func (m *GroupItemRequestBuilder) DrivesById(id string)(*i240fc320b71fb1dab200c27cdf37d9ebbf070eaca1406a7175520b78a58ae259.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return i240fc320b71fb1dab200c27cdf37d9ebbf070eaca1406a7175520b78a58ae259.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Endpoints the endpoints property
func (m *GroupItemRequestBuilder) Endpoints()(*i2e123d7e0a7dec41d5c6aca2603dd5abf115d35da35c8cb4bb86811ec04fd4c4.EndpointsRequestBuilder) {
    return i2e123d7e0a7dec41d5c6aca2603dd5abf115d35da35c8cb4bb86811ec04fd4c4.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.endpoints.item collection
func (m *GroupItemRequestBuilder) EndpointsById(id string)(*i596929c19e65628c2d282a6486262d8d20bcb6dd817073a9ccc8fccfadb48cf9.EndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint%2Did"] = id
    }
    return i596929c19e65628c2d282a6486262d8d20bcb6dd817073a9ccc8fccfadb48cf9.NewEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events the events property
func (m *GroupItemRequestBuilder) Events()(*i151bd88b167789c7f379069cae0c5cf3ce7a33c6f1732d0fab52d7c413bd972e.EventsRequestBuilder) {
    return i151bd88b167789c7f379069cae0c5cf3ce7a33c6f1732d0fab52d7c413bd972e.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item collection
func (m *GroupItemRequestBuilder) EventsById(id string)(*i83c11cd83fc9d7781fbc5c92495b2b6a7d532c21b1630a06f79b248afbd23fef.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i83c11cd83fc9d7781fbc5c92495b2b6a7d532c21b1630a06f79b248afbd23fef.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *GroupItemRequestBuilder) Extensions()(*i553997603462e15023ce8c734341ef27d17795823ae349d86df79c9ff9d8a94b.ExtensionsRequestBuilder) {
    return i553997603462e15023ce8c734341ef27d17795823ae349d86df79c9ff9d8a94b.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.extensions.item collection
func (m *GroupItemRequestBuilder) ExtensionsById(id string)(*i63c21f0d1e989b0923cde33368f86d94e7d49586d39585b403db033fbf6714b3.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i63c21f0d1e989b0923cde33368f86d94e7d49586d39585b403db033fbf6714b3.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *GroupItemRequestBuilder) GroupLifecyclePolicies()(*ic46a4dc410467cf4d6acd5973d8b7065d566186d73bab4001ee7357a13824c2a.GroupLifecyclePoliciesRequestBuilder) {
    return ic46a4dc410467cf4d6acd5973d8b7065d566186d73bab4001ee7357a13824c2a.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.groupLifecyclePolicies.item collection
func (m *GroupItemRequestBuilder) GroupLifecyclePoliciesById(id string)(*if0eea92e7b4aa17a8cd5570652467567e700ce471d79115f32dd029e959e4332.GroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return if0eea92e7b4aa17a8cd5570652467567e700ce471d79115f32dd029e959e4332.NewGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *GroupItemRequestBuilder) MemberOf()(*i4dc03b0a66d1f721539bf9f15ca7ecb992d2e6bba823cc9dc1c01cdf6e7ced40.MemberOfRequestBuilder) {
    return i4dc03b0a66d1f721539bf9f15ca7ecb992d2e6bba823cc9dc1c01cdf6e7ced40.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.memberOf.item collection
func (m *GroupItemRequestBuilder) MemberOfById(id string)(*i2d68757616ff94d59b32b253f2d16af53ce6394915b9b8fbe88a8a527349871d.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i2d68757616ff94d59b32b253f2d16af53ce6394915b9b8fbe88a8a527349871d.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *GroupItemRequestBuilder) Members()(*i461649a33bb785166c01173e17a1397fef7d397f67457d9ca2744d261a4b9de2.MembersRequestBuilder) {
    return i461649a33bb785166c01173e17a1397fef7d397f67457d9ca2744d261a4b9de2.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.members.item collection
func (m *GroupItemRequestBuilder) MembersById(id string)(*i33e227b32f13209673672ae7a38ed6696cc7d828159332bb4ae1f5327022b450.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i33e227b32f13209673672ae7a38ed6696cc7d828159332bb4ae1f5327022b450.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MembersWithLicenseErrors the membersWithLicenseErrors property
func (m *GroupItemRequestBuilder) MembersWithLicenseErrors()(*ifa1a1fed53a335586bb0dddb85c852c234549b6f6a88aacc698e5b941fb2022e.MembersWithLicenseErrorsRequestBuilder) {
    return ifa1a1fed53a335586bb0dddb85c852c234549b6f6a88aacc698e5b941fb2022e.NewMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersWithLicenseErrorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.membersWithLicenseErrors.item collection
func (m *GroupItemRequestBuilder) MembersWithLicenseErrorsById(id string)(*i7fbcdd8dbde4be9d9f484459dd838f2fb66ca0ba35a2c3334cd07fec9bcbf936.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i7fbcdd8dbde4be9d9f484459dd838f2fb66ca0ba35a2c3334cd07fec9bcbf936.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *GroupItemRequestBuilder) Onenote()(*iead0fb61869eb3cc76adc24a7157c079df7ea6a6160170afab077a54a31066f6.OnenoteRequestBuilder) {
    return iead0fb61869eb3cc76adc24a7157c079df7ea6a6160170afab077a54a31066f6.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Owners the owners property
func (m *GroupItemRequestBuilder) Owners()(*i445223cb5fb1554cd6f84a16220bdec0fafc528688ea42ea5f437b47439a70ee.OwnersRequestBuilder) {
    return i445223cb5fb1554cd6f84a16220bdec0fafc528688ea42ea5f437b47439a70ee.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.owners.item collection
func (m *GroupItemRequestBuilder) OwnersById(id string)(*i65767a283d1bd8e37ad589483567950ffe4d54fec4319f8d2cc8957f634f5c9f.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i65767a283d1bd8e37ad589483567950ffe4d54fec4319f8d2cc8957f634f5c9f.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property joinedGroups in users
func (m *GroupItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property joinedGroups in users
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
func (m *GroupItemRequestBuilder) PermissionGrants()(*i5ef5d6109700cf9b74ce0f4735be7684ee06a994dcf750744922b83617c58ccf.PermissionGrantsRequestBuilder) {
    return i5ef5d6109700cf9b74ce0f4735be7684ee06a994dcf750744922b83617c58ccf.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.permissionGrants.item collection
func (m *GroupItemRequestBuilder) PermissionGrantsById(id string)(*i61fc5189c26f38bdac1e13e87931337652e594efe0d92f2a10203cd6732f32df.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i61fc5189c26f38bdac1e13e87931337652e594efe0d92f2a10203cd6732f32df.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *GroupItemRequestBuilder) Photo()(*i37d1c6edd7c02656ff19812d0f443f3a240b0f736e09d1396fee36797f90c682.PhotoRequestBuilder) {
    return i37d1c6edd7c02656ff19812d0f443f3a240b0f736e09d1396fee36797f90c682.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos the photos property
func (m *GroupItemRequestBuilder) Photos()(*i4caaab3ad40a9992ef480a614da8698cd73481940a876181504765a2d0deecca.PhotosRequestBuilder) {
    return i4caaab3ad40a9992ef480a614da8698cd73481940a876181504765a2d0deecca.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.photos.item collection
func (m *GroupItemRequestBuilder) PhotosById(id string)(*i142d97539d5c7f0c4d8140f25af8e17432073cbe3732b414e6aa5a97f9ab50a1.ProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return i142d97539d5c7f0c4d8140f25af8e17432073cbe3732b414e6aa5a97f9ab50a1.NewProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner the planner property
func (m *GroupItemRequestBuilder) Planner()(*i31f436f071c3daaebe4e21c67c01aeffbbeb0acf42a682d5bd0da1f8265ac096.PlannerRequestBuilder) {
    return i31f436f071c3daaebe4e21c67c01aeffbbeb0acf42a682d5bd0da1f8265ac096.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSenders the rejectedSenders property
func (m *GroupItemRequestBuilder) RejectedSenders()(*ic294ab9d19fb30cf14f9b35e0a081fcc7e11428608e26db5a8018dc4c7d494ed.RejectedSendersRequestBuilder) {
    return ic294ab9d19fb30cf14f9b35e0a081fcc7e11428608e26db5a8018dc4c7d494ed.NewRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.rejectedSenders.item collection
func (m *GroupItemRequestBuilder) RejectedSendersById(id string)(*i0d900e574447595db6a7d9962c3ff45bb15157f2a9ceda3ca75462393a2a0103.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i0d900e574447595db6a7d9962c3ff45bb15157f2a9ceda3ca75462393a2a0103.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Settings the settings property
func (m *GroupItemRequestBuilder) Settings()(*ia753cb6a12c0b696be3b71817e34405b91e0be88a923be55114d53cf49082434.SettingsRequestBuilder) {
    return ia753cb6a12c0b696be3b71817e34405b91e0be88a923be55114d53cf49082434.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.settings.item collection
func (m *GroupItemRequestBuilder) SettingsById(id string)(*ifc347c65375f4761c6131dd0524e19090cde8f5b20784aec198b784677b47203.DirectorySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directorySetting%2Did"] = id
    }
    return ifc347c65375f4761c6131dd0524e19090cde8f5b20784aec198b784677b47203.NewDirectorySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *GroupItemRequestBuilder) Sites()(*i3e491681dc3c7fc87f42d306fdf4a6db96b4e176b78bf4ded4085d340518a194.SitesRequestBuilder) {
    return i3e491681dc3c7fc87f42d306fdf4a6db96b4e176b78bf4ded4085d340518a194.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item collection
func (m *GroupItemRequestBuilder) SitesById(id string)(*ifac9bb0df9a76a9a7e01bb81905add19d789c6bde1441edd6687fc8b575e2682.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return ifac9bb0df9a76a9a7e01bb81905add19d789c6bde1441edd6687fc8b575e2682.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Team the team property
func (m *GroupItemRequestBuilder) Team()(*i46c907074dd13bc5887834efa148bd776d64cc248c900951c0e1e3478bf387b7.TeamRequestBuilder) {
    return i46c907074dd13bc5887834efa148bd776d64cc248c900951c0e1e3478bf387b7.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Threads the threads property
func (m *GroupItemRequestBuilder) Threads()(*i5f8f7910d54070b0f7e884c8482d350b4b873dc167891fbd89ba73da2e1e0986.ThreadsRequestBuilder) {
    return i5f8f7910d54070b0f7e884c8482d350b4b873dc167891fbd89ba73da2e1e0986.NewThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreadsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item collection
func (m *GroupItemRequestBuilder) ThreadsById(id string)(*i941e71faec83df85c828a4d8333dcaf30d8e831292c063cbb43bed311b59cc93.ConversationThreadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationThread%2Did"] = id
    }
    return i941e71faec83df85c828a4d8333dcaf30d8e831292c063cbb43bed311b59cc93.NewConversationThreadItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *GroupItemRequestBuilder) TransitiveMemberOf()(*ifbed89bb7272e35484ac2dc99d2fdbb35cfa6c81abb043f47080800bfc4f0a48.TransitiveMemberOfRequestBuilder) {
    return ifbed89bb7272e35484ac2dc99d2fdbb35cfa6c81abb043f47080800bfc4f0a48.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.transitiveMemberOf.item collection
func (m *GroupItemRequestBuilder) TransitiveMemberOfById(id string)(*i912942191b304ee201aeaa8ea87cca5a1101a458fbccbb508c42c75a57e5e13a.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i912942191b304ee201aeaa8ea87cca5a1101a458fbccbb508c42c75a57e5e13a.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMembers the transitiveMembers property
func (m *GroupItemRequestBuilder) TransitiveMembers()(*icf14a4a532b90c50e5f8b7f026c68df019ddb14511768f7b0df1b9ea05c8bd1a.TransitiveMembersRequestBuilder) {
    return icf14a4a532b90c50e5f8b7f026c68df019ddb14511768f7b0df1b9ea05c8bd1a.NewTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.transitiveMembers.item collection
func (m *GroupItemRequestBuilder) TransitiveMembersById(id string)(*i81fc61d562fa2c0a2d5ac39c21b3b8f91931def1c191e3eeb01afbdde27a5ae3.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i81fc61d562fa2c0a2d5ac39c21b3b8f91931def1c191e3eeb01afbdde27a5ae3.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
