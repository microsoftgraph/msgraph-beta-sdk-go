package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ItemTeamDefinitionScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamDefinitionScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the provisionStatus property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the provisionStatusCode property. Clients can also inspect the configuration of the schedule.
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
// NewItemTeamDefinitionScheduleRequestBuilderInternal instantiates a new ItemTeamDefinitionScheduleRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleRequestBuilder) {
    m := &ItemTeamDefinitionScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionScheduleRequestBuilder instantiates a new ItemTeamDefinitionScheduleRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// DayNotes provides operations to manage the dayNotes property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleDayNotesRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) DayNotes()(*ItemTeamDefinitionScheduleDayNotesRequestBuilder) {
    return NewItemTeamDefinitionScheduleDayNotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property schedule for teamTemplateDefinition
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the provisionStatus property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the provisionStatusCode property. Clients can also inspect the configuration of the schedule.
// returns a Scheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/schedule-get?view=graph-rest-1.0
func (m *ItemTeamDefinitionScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleOfferShiftRequestsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) OfferShiftRequests()(*ItemTeamDefinitionScheduleOfferShiftRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleOfferShiftRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleOpenShiftChangeRequestsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) OpenShiftChangeRequests()(*ItemTeamDefinitionScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleOpenShiftsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) OpenShifts()(*ItemTeamDefinitionScheduleOpenShiftsRequestBuilder) {
    return NewItemTeamDefinitionScheduleOpenShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put update the navigation property schedule in teamTemplateDefinition
// returns a Scheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionScheduleRequestBuilder) Put(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleSchedulingGroupsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) SchedulingGroups()(*ItemTeamDefinitionScheduleSchedulingGroupsRequestBuilder) {
    return NewItemTeamDefinitionScheduleSchedulingGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Share provides operations to call the share method.
// returns a *ItemTeamDefinitionScheduleShareRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) Share()(*ItemTeamDefinitionScheduleShareRequestBuilder) {
    return NewItemTeamDefinitionScheduleShareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleShiftsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) Shifts()(*ItemTeamDefinitionScheduleShiftsRequestBuilder) {
    return NewItemTeamDefinitionScheduleShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShiftsRoleDefinitions provides operations to manage the shiftsRoleDefinitions property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleShiftsRoleDefinitionsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) ShiftsRoleDefinitions()(*ItemTeamDefinitionScheduleShiftsRoleDefinitionsRequestBuilder) {
    return NewItemTeamDefinitionScheduleShiftsRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) SwapShiftsChangeRequests()(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeCards provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleTimeCardsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeCards()(*ItemTeamDefinitionScheduleTimeCardsRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeCardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleTimeOffReasonsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeOffReasons()(*ItemTeamDefinitionScheduleTimeOffReasonsRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeOffReasonsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleTimeOffRequestsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimeOffRequests()(*ItemTeamDefinitionScheduleTimeOffRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeOffRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleTimesOffRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) TimesOff()(*ItemTeamDefinitionScheduleTimesOffRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimesOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property schedule for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the provisionStatus property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the provisionStatusCode property. Clients can also inspect the configuration of the schedule.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update the navigation property schedule in teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ItemTeamDefinitionScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamDefinitionScheduleRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionScheduleRequestBuilder) {
    return NewItemTeamDefinitionScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
