package schedule

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i039eec79a002c78fac8d396c29ad1d675afcebb55ccdd1292de00f1d563fa471 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/schedulinggroups"
    i199337207ee48c6878bdbfc09bf984d5d2e0a412f68bc573bf06fb43eb40ee05 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/openshiftchangerequests"
    i42830f0f1a1ed689e96387146a8377637687c5f28c5985adc329dac70ab51161 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timesoff"
    i6c650dcff368f3263028d9a1f3597fe942bb505ecb6a0bfe4d9a2f9c1cb050ec "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/share"
    i6f583c173999f95a06cd73ba38bd498d2605960a70ea7a68c0313911e664ca5a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/swapshiftschangerequests"
    i8be04cc4d216312c0521147201b4ccd98887b3e052a903d1c9082dbfc5024f79 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/openshifts"
    ia0fbf373cb1eb84246d8b686530235f40d552d7bc1e3ba491969c7f280f7b121 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/shifts"
    ib1d3752d34bff651eaf9242d8c3f304cbb01e570654e8846cf6feea9f086b36c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timeoffrequests"
    ic15cdf757ce02167beb311786e0d20a101a2672e64fe1cf1c82661bcd88b4cf8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards"
    ief0b2856782e484f49aa7eb312c8eff32f14d3d4711c16721080090829837222 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timeoffreasons"
    if87404376e57eab6aad08d117694742e1fe4795d02d5b48ed6f8a4364fc7123f "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/offershiftrequests"
    i3256ba159ae6d6aebb727827d05f7dc3a29809b6c6877a5e8d99e86294e67992 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/openshiftchangerequests/item"
    i52b520af63df4ebfc696996a2c167386745f9c6c26b895d430d3a2a1cf25cc97 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timesoff/item"
    i5a2542a239328c1612eb039bf256b57c38c0b4fbcd9e6e31c0874a5686dfb70a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/shifts/item"
    i6775a87f2ee913fe19344d96c72212bf22fe4b396b9628cb53d99cfcfa36fc9c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/offershiftrequests/item"
    i76807142d76611fc6c6d570af0c229489e3cc7f1d4b2e1faa6775c11c77bf17d "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timeoffrequests/item"
    i8124577d611f0e7f1ce72b1f65c4fcfd0aa2c04a1fb342411bf6ed5d6d8774a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item"
    ib7471221011c1f099711d83cdb09e6cf59391f0400222e6b44fe1265ce8c4493 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/schedulinggroups/item"
    iba89e12f6bd6b100e5e3bf3c77fc9a5410638666188bc55e58937d53123f89b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/swapshiftschangerequests/item"
    id2bc58c9f17bb58f2ec16eec5adb94eb886c6d7b8501cb63385534d930c6b669 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/openshifts/item"
    iff34320fa267865f50b2f1efcb26ca13e9686e30a7e06ca18ef1001b4174db56 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timeoffreasons/item"
)

// ScheduleRequestBuilder builds and executes requests for operations under \teams\{team-id}\schedule
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
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Schedule;
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
    m.urlTemplate = "{+baseurl}/teams/{team_id}/schedule{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the schedule of shifts for this team.
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
// CreatePatchRequestInformation the schedule of shifts for this team.
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
// Delete the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Delete(options *ScheduleRequestBuilderDeleteOptions)(error) {
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
// Get the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Get(options *ScheduleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Schedule, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSchedule() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Schedule), nil
}
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*if87404376e57eab6aad08d117694742e1fe4795d02d5b48ed6f8a4364fc7123f.OfferShiftRequestsRequestBuilder) {
    return if87404376e57eab6aad08d117694742e1fe4795d02d5b48ed6f8a4364fc7123f.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i6775a87f2ee913fe19344d96c72212bf22fe4b396b9628cb53d99cfcfa36fc9c.OfferShiftRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest_id"] = id
    }
    return i6775a87f2ee913fe19344d96c72212bf22fe4b396b9628cb53d99cfcfa36fc9c.NewOfferShiftRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*i199337207ee48c6878bdbfc09bf984d5d2e0a412f68bc573bf06fb43eb40ee05.OpenShiftChangeRequestsRequestBuilder) {
    return i199337207ee48c6878bdbfc09bf984d5d2e0a412f68bc573bf06fb43eb40ee05.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i3256ba159ae6d6aebb727827d05f7dc3a29809b6c6877a5e8d99e86294e67992.OpenShiftChangeRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest_id"] = id
    }
    return i3256ba159ae6d6aebb727827d05f7dc3a29809b6c6877a5e8d99e86294e67992.NewOpenShiftChangeRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShifts()(*i8be04cc4d216312c0521147201b4ccd98887b3e052a903d1c9082dbfc5024f79.OpenShiftsRequestBuilder) {
    return i8be04cc4d216312c0521147201b4ccd98887b3e052a903d1c9082dbfc5024f79.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*id2bc58c9f17bb58f2ec16eec5adb94eb886c6d7b8501cb63385534d930c6b669.OpenShiftRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift_id"] = id
    }
    return id2bc58c9f17bb58f2ec16eec5adb94eb886c6d7b8501cb63385534d930c6b669.NewOpenShiftRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Patch(options *ScheduleRequestBuilderPatchOptions)(error) {
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
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i039eec79a002c78fac8d396c29ad1d675afcebb55ccdd1292de00f1d563fa471.SchedulingGroupsRequestBuilder) {
    return i039eec79a002c78fac8d396c29ad1d675afcebb55ccdd1292de00f1d563fa471.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*ib7471221011c1f099711d83cdb09e6cf59391f0400222e6b44fe1265ce8c4493.SchedulingGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup_id"] = id
    }
    return ib7471221011c1f099711d83cdb09e6cf59391f0400222e6b44fe1265ce8c4493.NewSchedulingGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Share()(*i6c650dcff368f3263028d9a1f3597fe942bb505ecb6a0bfe4d9a2f9c1cb050ec.ShareRequestBuilder) {
    return i6c650dcff368f3263028d9a1f3597fe942bb505ecb6a0bfe4d9a2f9c1cb050ec.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Shifts()(*ia0fbf373cb1eb84246d8b686530235f40d552d7bc1e3ba491969c7f280f7b121.ShiftsRequestBuilder) {
    return ia0fbf373cb1eb84246d8b686530235f40d552d7bc1e3ba491969c7f280f7b121.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i5a2542a239328c1612eb039bf256b57c38c0b4fbcd9e6e31c0874a5686dfb70a.ShiftRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift_id"] = id
    }
    return i5a2542a239328c1612eb039bf256b57c38c0b4fbcd9e6e31c0874a5686dfb70a.NewShiftRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i6f583c173999f95a06cd73ba38bd498d2605960a70ea7a68c0313911e664ca5a.SwapShiftsChangeRequestsRequestBuilder) {
    return i6f583c173999f95a06cd73ba38bd498d2605960a70ea7a68c0313911e664ca5a.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*iba89e12f6bd6b100e5e3bf3c77fc9a5410638666188bc55e58937d53123f89b0.SwapShiftsChangeRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest_id"] = id
    }
    return iba89e12f6bd6b100e5e3bf3c77fc9a5410638666188bc55e58937d53123f89b0.NewSwapShiftsChangeRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeCards()(*ic15cdf757ce02167beb311786e0d20a101a2672e64fe1cf1c82661bcd88b4cf8.TimeCardsRequestBuilder) {
    return ic15cdf757ce02167beb311786e0d20a101a2672e64fe1cf1c82661bcd88b4cf8.NewTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.timeCards.item collection
func (m *ScheduleRequestBuilder) TimeCardsById(id string)(*i8124577d611f0e7f1ce72b1f65c4fcfd0aa2c04a1fb342411bf6ed5d6d8774a2.TimeCardRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard_id"] = id
    }
    return i8124577d611f0e7f1ce72b1f65c4fcfd0aa2c04a1fb342411bf6ed5d6d8774a2.NewTimeCardRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffReasons()(*ief0b2856782e484f49aa7eb312c8eff32f14d3d4711c16721080090829837222.TimeOffReasonsRequestBuilder) {
    return ief0b2856782e484f49aa7eb312c8eff32f14d3d4711c16721080090829837222.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*iff34320fa267865f50b2f1efcb26ca13e9686e30a7e06ca18ef1001b4174db56.TimeOffReasonRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason_id"] = id
    }
    return iff34320fa267865f50b2f1efcb26ca13e9686e30a7e06ca18ef1001b4174db56.NewTimeOffReasonRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffRequests()(*ib1d3752d34bff651eaf9242d8c3f304cbb01e570654e8846cf6feea9f086b36c.TimeOffRequestsRequestBuilder) {
    return ib1d3752d34bff651eaf9242d8c3f304cbb01e570654e8846cf6feea9f086b36c.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*i76807142d76611fc6c6d570af0c229489e3cc7f1d4b2e1faa6775c11c77bf17d.TimeOffRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest_id"] = id
    }
    return i76807142d76611fc6c6d570af0c229489e3cc7f1d4b2e1faa6775c11c77bf17d.NewTimeOffRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimesOff()(*i42830f0f1a1ed689e96387146a8377637687c5f28c5985adc329dac70ab51161.TimesOffRequestBuilder) {
    return i42830f0f1a1ed689e96387146a8377637687c5f28c5985adc329dac70ab51161.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i52b520af63df4ebfc696996a2c167386745f9c6c26b895d430d3a2a1cf25cc97.TimeOffRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff_id"] = id
    }
    return i52b520af63df4ebfc696996a2c167386745f9c6c26b895d430d3a2a1cf25cc97.NewTimeOffRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
