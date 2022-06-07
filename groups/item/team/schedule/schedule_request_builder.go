package schedule

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b02f34fcf99cf273f9a0db33db03577f0ecc0090b9a3196c49731ae7c647883 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/offershiftrequests"
    i1b4956c318a95907b289926f2f349a13a03bd53f204034869e6c18c123bd91b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/schedulinggroups"
    i2553c7a1f6bdb23b70b04eb4cdc272d54605b558a3fbb425988c5da4353eedd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timeoffreasons"
    i3ea948c00f86da5506b5fb19e35678df8a061723e613d2b43e1f0ca68a77fd00 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/swapshiftschangerequests"
    i61c8d6dd868142571ae9d414e92547ad946e104835d26bee54daa2289fe2f798 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timeoffrequests"
    i960a8c3fc76be9613b4a46ff94e0a6c9b3537c8149dc9b08f2d690e660624115 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/openshifts"
    ia243166708454ebb79a7704a3d23d702bece80f278249c1514950a1811ada1af "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/share"
    iafd022f789e9a015b0a92ab6c31ad7cbc606db6d49e19bc61ef1125f9cc96099 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timecards"
    idc27395c0530f5c8f5dd6a7a5c6b24c9785a6c3a5a809f6ac9b01f4d1acdd759 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/shifts"
    ie3b0c796101d4a55851a62a2e119eec054545d790525d0067008abb58f89b704 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/openshiftchangerequests"
    if8e94710df3baa717497fe994b443a127aa3760c2041e0b916d008049ac653dd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timesoff"
    i0702837c365d60402a8e2afb2a9f9242c128c230b7a040f9b0877df5144b880c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/openshiftchangerequests/item"
    i1c5847630cacaba4e461bf67ed06b314174b8c0237544034c4910b435c77c057 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/offershiftrequests/item"
    i558f442c9c105fecf6dc2808f3f0ccf3dd36eaf23e2ff8a5eda062563383b8ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timecards/item"
    i88de33b4566f7164405d2be64dd887c8eca913cd81d2bdbe497fb8a5f2995b44 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/openshifts/item"
    i907a879837d4ad8193849a1e28175301bc33f56af47b3bec76c043fe8d433c5a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timeoffreasons/item"
    i91742ad5d329680920570d7cf6540a09d9b38abb7a1cdaf78bd4bb6f651bf999 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timesoff/item"
    i9fd19cada1022c07edd4f9efce5397d0335e72bbcd3c93925f0b087e1731e801 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/schedulinggroups/item"
    ibb2b06668a1933ce9d6f97e412e8fb7607e69b118dff349a01a8b30286b128c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/shifts/item"
    ic5baf05ee4c9fb47c702bb81611bcdfbb0d8e387300f1492c4d692330fdc63d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/swapshiftschangerequests/item"
    id497e1d27dcac41a98180f806207ea790f3a08d79dee1aef6b1ac8a77e04618a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule/timeoffrequests/item"
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
// ScheduleRequestBuilderGetQueryParameters the schedule of shifts for this team.
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
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/schedule{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// OfferShiftRequests the offerShiftRequests property
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*i0b02f34fcf99cf273f9a0db33db03577f0ecc0090b9a3196c49731ae7c647883.OfferShiftRequestsRequestBuilder) {
    return i0b02f34fcf99cf273f9a0db33db03577f0ecc0090b9a3196c49731ae7c647883.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i1c5847630cacaba4e461bf67ed06b314174b8c0237544034c4910b435c77c057.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return i1c5847630cacaba4e461bf67ed06b314174b8c0237544034c4910b435c77c057.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests the openShiftChangeRequests property
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*ie3b0c796101d4a55851a62a2e119eec054545d790525d0067008abb58f89b704.OpenShiftChangeRequestsRequestBuilder) {
    return ie3b0c796101d4a55851a62a2e119eec054545d790525d0067008abb58f89b704.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i0702837c365d60402a8e2afb2a9f9242c128c230b7a040f9b0877df5144b880c.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return i0702837c365d60402a8e2afb2a9f9242c128c230b7a040f9b0877df5144b880c.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts the openShifts property
func (m *ScheduleRequestBuilder) OpenShifts()(*i960a8c3fc76be9613b4a46ff94e0a6c9b3537c8149dc9b08f2d690e660624115.OpenShiftsRequestBuilder) {
    return i960a8c3fc76be9613b4a46ff94e0a6c9b3537c8149dc9b08f2d690e660624115.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*i88de33b4566f7164405d2be64dd887c8eca913cd81d2bdbe497fb8a5f2995b44.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return i88de33b4566f7164405d2be64dd887c8eca913cd81d2bdbe497fb8a5f2995b44.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// SchedulingGroups the schedulingGroups property
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i1b4956c318a95907b289926f2f349a13a03bd53f204034869e6c18c123bd91b1.SchedulingGroupsRequestBuilder) {
    return i1b4956c318a95907b289926f2f349a13a03bd53f204034869e6c18c123bd91b1.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*i9fd19cada1022c07edd4f9efce5397d0335e72bbcd3c93925f0b087e1731e801.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return i9fd19cada1022c07edd4f9efce5397d0335e72bbcd3c93925f0b087e1731e801.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share the share property
func (m *ScheduleRequestBuilder) Share()(*ia243166708454ebb79a7704a3d23d702bece80f278249c1514950a1811ada1af.ShareRequestBuilder) {
    return ia243166708454ebb79a7704a3d23d702bece80f278249c1514950a1811ada1af.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts the shifts property
func (m *ScheduleRequestBuilder) Shifts()(*idc27395c0530f5c8f5dd6a7a5c6b24c9785a6c3a5a809f6ac9b01f4d1acdd759.ShiftsRequestBuilder) {
    return idc27395c0530f5c8f5dd6a7a5c6b24c9785a6c3a5a809f6ac9b01f4d1acdd759.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*ibb2b06668a1933ce9d6f97e412e8fb7607e69b118dff349a01a8b30286b128c1.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return ibb2b06668a1933ce9d6f97e412e8fb7607e69b118dff349a01a8b30286b128c1.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests the swapShiftsChangeRequests property
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i3ea948c00f86da5506b5fb19e35678df8a061723e613d2b43e1f0ca68a77fd00.SwapShiftsChangeRequestsRequestBuilder) {
    return i3ea948c00f86da5506b5fb19e35678df8a061723e613d2b43e1f0ca68a77fd00.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*ic5baf05ee4c9fb47c702bb81611bcdfbb0d8e387300f1492c4d692330fdc63d3.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return ic5baf05ee4c9fb47c702bb81611bcdfbb0d8e387300f1492c4d692330fdc63d3.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeCards the timeCards property
func (m *ScheduleRequestBuilder) TimeCards()(*iafd022f789e9a015b0a92ab6c31ad7cbc606db6d49e19bc61ef1125f9cc96099.TimeCardsRequestBuilder) {
    return iafd022f789e9a015b0a92ab6c31ad7cbc606db6d49e19bc61ef1125f9cc96099.NewTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.timeCards.item collection
func (m *ScheduleRequestBuilder) TimeCardsById(id string)(*i558f442c9c105fecf6dc2808f3f0ccf3dd36eaf23e2ff8a5eda062563383b8ea.TimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return i558f442c9c105fecf6dc2808f3f0ccf3dd36eaf23e2ff8a5eda062563383b8ea.NewTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons the timeOffReasons property
func (m *ScheduleRequestBuilder) TimeOffReasons()(*i2553c7a1f6bdb23b70b04eb4cdc272d54605b558a3fbb425988c5da4353eedd8.TimeOffReasonsRequestBuilder) {
    return i2553c7a1f6bdb23b70b04eb4cdc272d54605b558a3fbb425988c5da4353eedd8.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*i907a879837d4ad8193849a1e28175301bc33f56af47b3bec76c043fe8d433c5a.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return i907a879837d4ad8193849a1e28175301bc33f56af47b3bec76c043fe8d433c5a.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests the timeOffRequests property
func (m *ScheduleRequestBuilder) TimeOffRequests()(*i61c8d6dd868142571ae9d414e92547ad946e104835d26bee54daa2289fe2f798.TimeOffRequestsRequestBuilder) {
    return i61c8d6dd868142571ae9d414e92547ad946e104835d26bee54daa2289fe2f798.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*id497e1d27dcac41a98180f806207ea790f3a08d79dee1aef6b1ac8a77e04618a.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return id497e1d27dcac41a98180f806207ea790f3a08d79dee1aef6b1ac8a77e04618a.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff the timesOff property
func (m *ScheduleRequestBuilder) TimesOff()(*if8e94710df3baa717497fe994b443a127aa3760c2041e0b916d008049ac653dd.TimesOffRequestBuilder) {
    return if8e94710df3baa717497fe994b443a127aa3760c2041e0b916d008049ac653dd.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i91742ad5d329680920570d7cf6540a09d9b38abb7a1cdaf78bd4bb6f651bf999.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return i91742ad5d329680920570d7cf6540a09d9b38abb7a1cdaf78bd4bb6f651bf999.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
