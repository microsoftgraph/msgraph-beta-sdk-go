// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ItemTeamScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamScheduleRequestBuilderGetQueryParameters the schedule of shifts for this team.
type ItemTeamScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamScheduleRequestBuilderGetQueryParameters
}
// ItemTeamScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamScheduleRequestBuilderInternal instantiates a new ItemTeamScheduleRequestBuilder and sets the default values.
func NewItemTeamScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleRequestBuilder) {
    m := &ItemTeamScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/schedule{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamScheduleRequestBuilder instantiates a new ItemTeamScheduleRequestBuilder and sets the default values.
func NewItemTeamScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// DayNotes provides operations to manage the dayNotes property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleDayNotesRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) DayNotes()(*ItemTeamScheduleDayNotesRequestBuilder) {
    return NewItemTeamScheduleDayNotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property schedule for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the schedule of shifts for this team.
// returns a Scheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
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
// returns a *ItemTeamScheduleOfferShiftRequestsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) OfferShiftRequests()(*ItemTeamScheduleOfferShiftRequestsRequestBuilder) {
    return NewItemTeamScheduleOfferShiftRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleOpenShiftChangeRequestsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) OpenShiftChangeRequests()(*ItemTeamScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewItemTeamScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleOpenShiftsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) OpenShifts()(*ItemTeamScheduleOpenShiftsRequestBuilder) {
    return NewItemTeamScheduleOpenShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put update the navigation property schedule in groups
// returns a Scheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleRequestBuilder) Put(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ItemTeamScheduleRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
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
// returns a *ItemTeamScheduleSchedulingGroupsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) SchedulingGroups()(*ItemTeamScheduleSchedulingGroupsRequestBuilder) {
    return NewItemTeamScheduleSchedulingGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Share provides operations to call the share method.
// returns a *ItemTeamScheduleShareRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) Share()(*ItemTeamScheduleShareRequestBuilder) {
    return NewItemTeamScheduleShareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleShiftsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) Shifts()(*ItemTeamScheduleShiftsRequestBuilder) {
    return NewItemTeamScheduleShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShiftsRoleDefinitions provides operations to manage the shiftsRoleDefinitions property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) ShiftsRoleDefinitions()(*ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) {
    return NewItemTeamScheduleShiftsRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleSwapShiftsChangeRequestsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) SwapShiftsChangeRequests()(*ItemTeamScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemTeamScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeCards provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleTimeCardsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) TimeCards()(*ItemTeamScheduleTimeCardsRequestBuilder) {
    return NewItemTeamScheduleTimeCardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleTimeOffReasonsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) TimeOffReasons()(*ItemTeamScheduleTimeOffReasonsRequestBuilder) {
    return NewItemTeamScheduleTimeOffReasonsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleTimeOffRequestsRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) TimeOffRequests()(*ItemTeamScheduleTimeOffRequestsRequestBuilder) {
    return NewItemTeamScheduleTimeOffRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleTimesOffRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) TimesOff()(*ItemTeamScheduleTimesOffRequestBuilder) {
    return NewItemTeamScheduleTimesOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property schedule for groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the schedule of shifts for this team.
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update the navigation property schedule in groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ItemTeamScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamScheduleRequestBuilder when successful
func (m *ItemTeamScheduleRequestBuilder) WithUrl(rawUrl string)(*ItemTeamScheduleRequestBuilder) {
    return NewItemTeamScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
