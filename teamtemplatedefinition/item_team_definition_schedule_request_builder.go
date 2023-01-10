package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ItemTeamDefinitionScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamDefinitionScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type ItemTeamDefinitionScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionScheduleRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleRequestBuilder) {
    m := &ItemTeamDefinitionScheduleRequestBuilder{
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
// NewItemTeamDefinitionScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property schedule for teamTemplateDefinition
func (m *ItemTeamDefinitionScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/schedule-get?view=graph-rest-1.0
func (m *ItemTeamDefinitionScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
func (m *ItemTeamDefinitionScheduleRequestBuilder) OfferShiftRequests()(*ItemTeamDefinitionScheduleOfferShiftRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) OfferShiftRequestsById(id string)(*ItemTeamDefinitionScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) OpenShiftChangeRequests()(*ItemTeamDefinitionScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*ItemTeamDefinitionScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) OpenShifts()(*ItemTeamDefinitionScheduleOpenShiftsRequestBuilder) {
    return NewItemTeamDefinitionScheduleOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) OpenShiftsById(id string)(*ItemTeamDefinitionScheduleOpenShiftsOpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleOpenShiftsOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Put update the navigation property schedule in teamTemplateDefinition
func (m *ItemTeamDefinitionScheduleRequestBuilder) Put(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
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
func (m *ItemTeamDefinitionScheduleRequestBuilder) SchedulingGroups()(*ItemTeamDefinitionScheduleSchedulingGroupsRequestBuilder) {
    return NewItemTeamDefinitionScheduleSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) SchedulingGroupsById(id string)(*ItemTeamDefinitionScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share provides operations to call the share method.
func (m *ItemTeamDefinitionScheduleRequestBuilder) Share()(*ItemTeamDefinitionScheduleShareRequestBuilder) {
    return NewItemTeamDefinitionScheduleShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) Shifts()(*ItemTeamDefinitionScheduleShiftsRequestBuilder) {
    return NewItemTeamDefinitionScheduleShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) ShiftsById(id string)(*ItemTeamDefinitionScheduleShiftsShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleShiftsShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) SwapShiftsChangeRequests()(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeCards provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeCards()(*ItemTeamDefinitionScheduleTimeCardsRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeCardsById(id string)(*ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeOffReasons()(*ItemTeamDefinitionScheduleTimeOffReasonsRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeOffReasonsById(id string)(*ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeOffRequests()(*ItemTeamDefinitionScheduleTimeOffRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeOffRequestsById(id string)(*ItemTeamDefinitionScheduleTimeOffRequestsTimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleTimeOffRequestsTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimesOff()(*ItemTeamDefinitionScheduleTimesOffRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimesOffById(id string)(*ItemTeamDefinitionScheduleTimesOffTimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return NewItemTeamDefinitionScheduleTimesOffTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property schedule for teamTemplateDefinition
func (m *ItemTeamDefinitionScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *ItemTeamDefinitionScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation update the navigation property schedule in teamTemplateDefinition
func (m *ItemTeamDefinitionScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
