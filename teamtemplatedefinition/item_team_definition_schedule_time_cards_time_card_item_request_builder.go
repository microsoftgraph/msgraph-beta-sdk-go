package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
type ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderGetQueryParameters get the properties and relationships of a timeCard object by ID.
type ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClockOut provides operations to call the clockOut method.
// returns a *ItemTeamDefinitionScheduleTimeCardsItemClockOutRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) ClockOut()(*ItemTeamDefinitionScheduleTimeCardsItemClockOutRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeCardsItemClockOutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Confirm provides operations to call the confirm method.
// returns a *ItemTeamDefinitionScheduleTimeCardsItemConfirmRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) Confirm()(*ItemTeamDefinitionScheduleTimeCardsItemConfirmRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeCardsItemConfirmRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderInternal instantiates a new ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) {
    m := &ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/timeCards/{timeCard%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder instantiates a new ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a timeCard instance in a schedule.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/timecard-delete?view=graph-rest-1.0
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EndBreak provides operations to call the endBreak method.
// returns a *ItemTeamDefinitionScheduleTimeCardsItemEndBreakRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) EndBreak()(*ItemTeamDefinitionScheduleTimeCardsItemEndBreakRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeCardsItemEndBreakRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the properties and relationships of a timeCard object by ID.
// returns a TimeCardable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/timecard-get?view=graph-rest-1.0
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable), nil
}
// Patch replace an existing timeCard with updated values.
// returns a TimeCardable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/timecard-replace?view=graph-rest-1.0
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, requestConfiguration *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable), nil
}
// StartBreak provides operations to call the startBreak method.
// returns a *ItemTeamDefinitionScheduleTimeCardsItemStartBreakRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) StartBreak()(*ItemTeamDefinitionScheduleTimeCardsItemStartBreakRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeCardsItemStartBreakRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a timeCard instance in a schedule.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/timeCards/{timeCard%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of a timeCard object by ID.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation replace an existing timeCard with updated values.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, requestConfiguration *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/timeCards/{timeCard%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeCardsTimeCardItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
