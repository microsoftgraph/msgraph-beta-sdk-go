package schedule

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
// ScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/schedule{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property schedule for teams
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
// CreatePutRequestInformation update the navigation property schedule in teams
func (m *ScheduleRequestBuilder) CreatePutRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property schedule for teams
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
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*if87404376e57eab6aad08d117694742e1fe4795d02d5b48ed6f8a4364fc7123f.OfferShiftRequestsRequestBuilder) {
    return if87404376e57eab6aad08d117694742e1fe4795d02d5b48ed6f8a4364fc7123f.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i6775a87f2ee913fe19344d96c72212bf22fe4b396b9628cb53d99cfcfa36fc9c.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return i6775a87f2ee913fe19344d96c72212bf22fe4b396b9628cb53d99cfcfa36fc9c.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*i199337207ee48c6878bdbfc09bf984d5d2e0a412f68bc573bf06fb43eb40ee05.OpenShiftChangeRequestsRequestBuilder) {
    return i199337207ee48c6878bdbfc09bf984d5d2e0a412f68bc573bf06fb43eb40ee05.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i3256ba159ae6d6aebb727827d05f7dc3a29809b6c6877a5e8d99e86294e67992.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return i3256ba159ae6d6aebb727827d05f7dc3a29809b6c6877a5e8d99e86294e67992.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShifts()(*i8be04cc4d216312c0521147201b4ccd98887b3e052a903d1c9082dbfc5024f79.OpenShiftsRequestBuilder) {
    return i8be04cc4d216312c0521147201b4ccd98887b3e052a903d1c9082dbfc5024f79.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*id2bc58c9f17bb58f2ec16eec5adb94eb886c6d7b8501cb63385534d930c6b669.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return id2bc58c9f17bb58f2ec16eec5adb94eb886c6d7b8501cb63385534d930c6b669.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Put update the navigation property schedule in teams
func (m *ScheduleRequestBuilder) Put(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.CreatePutRequestInformation(ctx, body, requestConfiguration);
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
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i039eec79a002c78fac8d396c29ad1d675afcebb55ccdd1292de00f1d563fa471.SchedulingGroupsRequestBuilder) {
    return i039eec79a002c78fac8d396c29ad1d675afcebb55ccdd1292de00f1d563fa471.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*ib7471221011c1f099711d83cdb09e6cf59391f0400222e6b44fe1265ce8c4493.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return ib7471221011c1f099711d83cdb09e6cf59391f0400222e6b44fe1265ce8c4493.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share provides operations to call the share method.
func (m *ScheduleRequestBuilder) Share()(*i6c650dcff368f3263028d9a1f3597fe942bb505ecb6a0bfe4d9a2f9c1cb050ec.ShareRequestBuilder) {
    return i6c650dcff368f3263028d9a1f3597fe942bb505ecb6a0bfe4d9a2f9c1cb050ec.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) Shifts()(*ia0fbf373cb1eb84246d8b686530235f40d552d7bc1e3ba491969c7f280f7b121.ShiftsRequestBuilder) {
    return ia0fbf373cb1eb84246d8b686530235f40d552d7bc1e3ba491969c7f280f7b121.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i5a2542a239328c1612eb039bf256b57c38c0b4fbcd9e6e31c0874a5686dfb70a.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return i5a2542a239328c1612eb039bf256b57c38c0b4fbcd9e6e31c0874a5686dfb70a.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i6f583c173999f95a06cd73ba38bd498d2605960a70ea7a68c0313911e664ca5a.SwapShiftsChangeRequestsRequestBuilder) {
    return i6f583c173999f95a06cd73ba38bd498d2605960a70ea7a68c0313911e664ca5a.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*iba89e12f6bd6b100e5e3bf3c77fc9a5410638666188bc55e58937d53123f89b0.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return iba89e12f6bd6b100e5e3bf3c77fc9a5410638666188bc55e58937d53123f89b0.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeCards provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeCards()(*ic15cdf757ce02167beb311786e0d20a101a2672e64fe1cf1c82661bcd88b4cf8.TimeCardsRequestBuilder) {
    return ic15cdf757ce02167beb311786e0d20a101a2672e64fe1cf1c82661bcd88b4cf8.NewTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeCardsById(id string)(*i8124577d611f0e7f1ce72b1f65c4fcfd0aa2c04a1fb342411bf6ed5d6d8774a2.TimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return i8124577d611f0e7f1ce72b1f65c4fcfd0aa2c04a1fb342411bf6ed5d6d8774a2.NewTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffReasons()(*ief0b2856782e484f49aa7eb312c8eff32f14d3d4711c16721080090829837222.TimeOffReasonsRequestBuilder) {
    return ief0b2856782e484f49aa7eb312c8eff32f14d3d4711c16721080090829837222.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*iff34320fa267865f50b2f1efcb26ca13e9686e30a7e06ca18ef1001b4174db56.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return iff34320fa267865f50b2f1efcb26ca13e9686e30a7e06ca18ef1001b4174db56.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffRequests()(*ib1d3752d34bff651eaf9242d8c3f304cbb01e570654e8846cf6feea9f086b36c.TimeOffRequestsRequestBuilder) {
    return ib1d3752d34bff651eaf9242d8c3f304cbb01e570654e8846cf6feea9f086b36c.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*i76807142d76611fc6c6d570af0c229489e3cc7f1d4b2e1faa6775c11c77bf17d.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return i76807142d76611fc6c6d570af0c229489e3cc7f1d4b2e1faa6775c11c77bf17d.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimesOff()(*i42830f0f1a1ed689e96387146a8377637687c5f28c5985adc329dac70ab51161.TimesOffRequestBuilder) {
    return i42830f0f1a1ed689e96387146a8377637687c5f28c5985adc329dac70ab51161.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i52b520af63df4ebfc696996a2c167386745f9c6c26b895d430d3a2a1cf25cc97.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return i52b520af63df4ebfc696996a2c167386745f9c6c26b895d430d3a2a1cf25cc97.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
