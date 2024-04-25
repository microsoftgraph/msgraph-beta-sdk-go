package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
type ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderGetQueryParameters the set of reasons for a time off in the schedule.
type ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal instantiates a new ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    m := &ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/timeOffReasons/{timeOffReason%2Did}{?%24select}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder instantiates a new ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property timeOffReasons for teamTemplateDefinition
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the set of reasons for a time off in the schedule.
// returns a TimeOffReasonable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeOffReasonFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable), nil
}
// Patch update the navigation property timeOffReasons in teamTemplateDefinition
// returns a TimeOffReasonable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable, requestConfiguration *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeOffReasonFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable), nil
}
// ToDeleteRequestInformation delete navigation property timeOffReasons for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the set of reasons for a time off in the schedule.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property timeOffReasons in teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable, requestConfiguration *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    return NewItemTeamDefinitionScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
