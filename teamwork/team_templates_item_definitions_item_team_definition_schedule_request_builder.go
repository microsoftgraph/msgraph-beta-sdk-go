package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderGetQueryParameters
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/schedule{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property schedule for teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/schedule-get?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// MicrosoftGraphShare provides operations to call the share method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) MicrosoftGraphShare()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleMicrosoftGraphShareShareRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleMicrosoftGraphShareShareRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) OfferShiftRequests()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOfferShiftRequestsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) OfferShiftRequestsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) OpenShiftChangeRequests()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) OpenShifts()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) OpenShiftsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftsOpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleOpenShiftsOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Put update the navigation property schedule in teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) Put(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) SchedulingGroups()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSchedulingGroupsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) SchedulingGroupsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) Shifts()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleShiftsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) ShiftsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleShiftsShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleShiftsShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) SwapShiftsChangeRequests()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimeCards provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimeCards()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimeCardsById provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimeCardsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimeOffReasons()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffReasonsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimeOffReasonsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimeOffRequests()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffRequestsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimeOffRequestsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffRequestsTimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeOffRequestsTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimesOff()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimesOffRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) TimesOffById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimesOffTimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimesOffTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property schedule for teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update the navigation property schedule in teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
