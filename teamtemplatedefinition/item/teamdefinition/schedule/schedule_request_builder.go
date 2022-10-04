package schedule

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i18babedba92c2d888766acef0dff2fd8edc22fa96789263b9bb42d928e5c6cdb "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timesoff"
    i1f68f20c2b3022f6604cdaad3791a0485267cf6768161698df2083d63172ee0e "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/openshiftchangerequests"
    i3245f0fcd9a00451852ec7427b5b98be90d7542ef20cc90ea254e0966eb22b0c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timeoffreasons"
    i4d7eb4eef104c75786c5698a0e578a47012f0fa05ce1cbcadb2e4fb8f43e748f "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/swapshiftschangerequests"
    i4e2c74a78a1cf3ce67c87b71a47247f13dc9c5bfcf2645dfb2c12998cb29062b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/schedulinggroups"
    i5d49f97dd15b04ecd522dc9fec8d3c9e488deb575458756087e9513f46dbb76f "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/shifts"
    i7b706e9f0549a43a09f04f019af092f3e6bc2fd89e47ad59ad124e709012ca42 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/offershiftrequests"
    i84258bff567ef49e93aa3af14de01340ebda287c640d2ec17e51e2e4ff29576f "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timecards"
    ib0d103bfd99a41bd4d577cc8e76a91ff5166d422c23d88ee9cb0f0849ff06a51 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/openshifts"
    icd9947a6d9cf19ee697c0cc91e57072942cf198ab7966450589f05363d1f7378 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timeoffrequests"
    ifea1a7cfea311f66f71eb8cb85646d65ec248f740c81e693443f21a81b745a34 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/share"
    i08c0f6758dba5877f755763a41cf112e2042b2d045b5cc9d917351bfef1e8631 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/openshiftchangerequests/item"
    i0b5c4415c0bd560460ccc7bbcb7ff705344c3f48144769d95758bd4aed62a91b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timeoffreasons/item"
    i3226f5d54241bf65626cab89693b992ad0a76a9d336631b2ccc002bd7aac93d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/swapshiftschangerequests/item"
    i52beaac159e84ebee0e794314fa964a7b82c7f4d26d654f38cfbd650eb9d8679 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/openshifts/item"
    i617741706d1adbab904cb27141cd001078a24bc129db7d5c85a7815f40753e51 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timeoffrequests/item"
    i7df724aaa2f465d8ab74b6c48f9e457d4c17774f72942ba36028ee3f3da329b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/shifts/item"
    i8b04b2b0b7ccf83aa848c52a0f3b33054a3cec2d100a6675e2c5c897ae3297bc "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timesoff/item"
    ib950ee686335fc62bcce4cffd2e3bfc18a87faa30a7b47e191566d94677a80ac "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/offershiftrequests/item"
    ic200bc563269458a0364637019f4511191258bea1444b75ffb3eed755987debc "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/timecards/item"
    ic5a1e6140ac803920f0bdbb1135a121da7a98938851437c25a5d02ad0faa45be "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule/schedulinggroups/item"
)

// ScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type ScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ScheduleRequestBuilderGetQueryParameters
}
// ScheduleRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schedule for teamTemplateDefinition
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schedule in teamTemplateDefinition
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property schedule for teamTemplateDefinition
func (m *ScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *ScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// OfferShiftRequests the offerShiftRequests property
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*i7b706e9f0549a43a09f04f019af092f3e6bc2fd89e47ad59ad124e709012ca42.OfferShiftRequestsRequestBuilder) {
    return i7b706e9f0549a43a09f04f019af092f3e6bc2fd89e47ad59ad124e709012ca42.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*ib950ee686335fc62bcce4cffd2e3bfc18a87faa30a7b47e191566d94677a80ac.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return ib950ee686335fc62bcce4cffd2e3bfc18a87faa30a7b47e191566d94677a80ac.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests the openShiftChangeRequests property
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*i1f68f20c2b3022f6604cdaad3791a0485267cf6768161698df2083d63172ee0e.OpenShiftChangeRequestsRequestBuilder) {
    return i1f68f20c2b3022f6604cdaad3791a0485267cf6768161698df2083d63172ee0e.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i08c0f6758dba5877f755763a41cf112e2042b2d045b5cc9d917351bfef1e8631.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return i08c0f6758dba5877f755763a41cf112e2042b2d045b5cc9d917351bfef1e8631.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts the openShifts property
func (m *ScheduleRequestBuilder) OpenShifts()(*ib0d103bfd99a41bd4d577cc8e76a91ff5166d422c23d88ee9cb0f0849ff06a51.OpenShiftsRequestBuilder) {
    return ib0d103bfd99a41bd4d577cc8e76a91ff5166d422c23d88ee9cb0f0849ff06a51.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*i52beaac159e84ebee0e794314fa964a7b82c7f4d26d654f38cfbd650eb9d8679.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return i52beaac159e84ebee0e794314fa964a7b82c7f4d26d654f38cfbd650eb9d8679.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in teamTemplateDefinition
func (m *ScheduleRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// SchedulingGroups the schedulingGroups property
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i4e2c74a78a1cf3ce67c87b71a47247f13dc9c5bfcf2645dfb2c12998cb29062b.SchedulingGroupsRequestBuilder) {
    return i4e2c74a78a1cf3ce67c87b71a47247f13dc9c5bfcf2645dfb2c12998cb29062b.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*ic5a1e6140ac803920f0bdbb1135a121da7a98938851437c25a5d02ad0faa45be.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return ic5a1e6140ac803920f0bdbb1135a121da7a98938851437c25a5d02ad0faa45be.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share the share property
func (m *ScheduleRequestBuilder) Share()(*ifea1a7cfea311f66f71eb8cb85646d65ec248f740c81e693443f21a81b745a34.ShareRequestBuilder) {
    return ifea1a7cfea311f66f71eb8cb85646d65ec248f740c81e693443f21a81b745a34.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts the shifts property
func (m *ScheduleRequestBuilder) Shifts()(*i5d49f97dd15b04ecd522dc9fec8d3c9e488deb575458756087e9513f46dbb76f.ShiftsRequestBuilder) {
    return i5d49f97dd15b04ecd522dc9fec8d3c9e488deb575458756087e9513f46dbb76f.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i7df724aaa2f465d8ab74b6c48f9e457d4c17774f72942ba36028ee3f3da329b6.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return i7df724aaa2f465d8ab74b6c48f9e457d4c17774f72942ba36028ee3f3da329b6.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests the swapShiftsChangeRequests property
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i4d7eb4eef104c75786c5698a0e578a47012f0fa05ce1cbcadb2e4fb8f43e748f.SwapShiftsChangeRequestsRequestBuilder) {
    return i4d7eb4eef104c75786c5698a0e578a47012f0fa05ce1cbcadb2e4fb8f43e748f.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*i3226f5d54241bf65626cab89693b992ad0a76a9d336631b2ccc002bd7aac93d6.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return i3226f5d54241bf65626cab89693b992ad0a76a9d336631b2ccc002bd7aac93d6.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeCards the timeCards property
func (m *ScheduleRequestBuilder) TimeCards()(*i84258bff567ef49e93aa3af14de01340ebda287c640d2ec17e51e2e4ff29576f.TimeCardsRequestBuilder) {
    return i84258bff567ef49e93aa3af14de01340ebda287c640d2ec17e51e2e4ff29576f.NewTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.timeCards.item collection
func (m *ScheduleRequestBuilder) TimeCardsById(id string)(*ic200bc563269458a0364637019f4511191258bea1444b75ffb3eed755987debc.TimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return ic200bc563269458a0364637019f4511191258bea1444b75ffb3eed755987debc.NewTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons the timeOffReasons property
func (m *ScheduleRequestBuilder) TimeOffReasons()(*i3245f0fcd9a00451852ec7427b5b98be90d7542ef20cc90ea254e0966eb22b0c.TimeOffReasonsRequestBuilder) {
    return i3245f0fcd9a00451852ec7427b5b98be90d7542ef20cc90ea254e0966eb22b0c.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*i0b5c4415c0bd560460ccc7bbcb7ff705344c3f48144769d95758bd4aed62a91b.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return i0b5c4415c0bd560460ccc7bbcb7ff705344c3f48144769d95758bd4aed62a91b.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests the timeOffRequests property
func (m *ScheduleRequestBuilder) TimeOffRequests()(*icd9947a6d9cf19ee697c0cc91e57072942cf198ab7966450589f05363d1f7378.TimeOffRequestsRequestBuilder) {
    return icd9947a6d9cf19ee697c0cc91e57072942cf198ab7966450589f05363d1f7378.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*i617741706d1adbab904cb27141cd001078a24bc129db7d5c85a7815f40753e51.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return i617741706d1adbab904cb27141cd001078a24bc129db7d5c85a7815f40753e51.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff the timesOff property
func (m *ScheduleRequestBuilder) TimesOff()(*i18babedba92c2d888766acef0dff2fd8edc22fa96789263b9bb42d928e5c6cdb.TimesOffRequestBuilder) {
    return i18babedba92c2d888766acef0dff2fd8edc22fa96789263b9bb42d928e5c6cdb.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i8b04b2b0b7ccf83aa848c52a0f3b33054a3cec2d100a6675e2c5c897ae3297bc.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return i8b04b2b0b7ccf83aa848c52a0f3b33054a3cec2d100a6675e2c5c897ae3297bc.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
