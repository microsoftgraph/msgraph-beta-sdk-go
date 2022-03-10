package schedule

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ScheduleRequestBuilderDeleteOptions options for Delete
type ScheduleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ScheduleRequestBuilderGetOptions options for Get
type ScheduleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ScheduleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ScheduleRequestBuilderGetQueryParameters the schedule of shifts for this team.
type ScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ScheduleRequestBuilderPatchOptions options for Patch
type ScheduleRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Scheduleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/team/schedule{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation(options *ScheduleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation(options *ScheduleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(options *ScheduleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) Delete(options *ScheduleRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Get(options *ScheduleRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateScheduleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Scheduleable), nil
}
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
        urlTplParams["offerShiftRequest_id"] = id
    }
    return i1c5847630cacaba4e461bf67ed06b314174b8c0237544034c4910b435c77c057.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["openShiftChangeRequest_id"] = id
    }
    return i0702837c365d60402a8e2afb2a9f9242c128c230b7a040f9b0877df5144b880c.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["openShift_id"] = id
    }
    return i88de33b4566f7164405d2be64dd887c8eca913cd81d2bdbe497fb8a5f2995b44.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) Patch(options *ScheduleRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["schedulingGroup_id"] = id
    }
    return i9fd19cada1022c07edd4f9efce5397d0335e72bbcd3c93925f0b087e1731e801.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Share()(*ia243166708454ebb79a7704a3d23d702bece80f278249c1514950a1811ada1af.ShareRequestBuilder) {
    return ia243166708454ebb79a7704a3d23d702bece80f278249c1514950a1811ada1af.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["shift_id"] = id
    }
    return ibb2b06668a1933ce9d6f97e412e8fb7607e69b118dff349a01a8b30286b128c1.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["swapShiftsChangeRequest_id"] = id
    }
    return ic5baf05ee4c9fb47c702bb81611bcdfbb0d8e387300f1492c4d692330fdc63d3.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["timeCard_id"] = id
    }
    return i558f442c9c105fecf6dc2808f3f0ccf3dd36eaf23e2ff8a5eda062563383b8ea.NewTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["timeOffReason_id"] = id
    }
    return i907a879837d4ad8193849a1e28175301bc33f56af47b3bec76c043fe8d433c5a.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["timeOffRequest_id"] = id
    }
    return id497e1d27dcac41a98180f806207ea790f3a08d79dee1aef6b1ac8a77e04618a.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["timeOff_id"] = id
    }
    return i91742ad5d329680920570d7cf6540a09d9b38abb7a1cdaf78bd4bb6f651bf999.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
