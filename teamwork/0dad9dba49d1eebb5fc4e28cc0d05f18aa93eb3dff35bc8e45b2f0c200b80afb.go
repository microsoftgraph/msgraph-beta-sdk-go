package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder provides operations to manage the plans property of the microsoft.graph.teamsChannelPlanner entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Archive provides operations to call the archive method.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemArchiveRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Archive()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemArchiveRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Buckets provides operations to manage the buckets property of the microsoft.graph.plannerPlan entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Buckets()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/planner/plans/{plannerPlan%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property plans for teamwork
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Details provides operations to manage the details property of the microsoft.graph.plannerPlan entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemDetailsRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Details()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemDetailsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
// returns a PlannerPlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable), nil
}
// MoveToContainer provides operations to call the moveToContainer method.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemMoveToContainerRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) MoveToContainer()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemMoveToContainerRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemMoveToContainerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property plans in teamwork
// returns a PlannerPlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable), nil
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Tasks()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property plans for teamwork
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property plans in teamwork
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unarchive provides operations to call the unarchive method.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemUnarchiveRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) Unarchive()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemUnarchiveRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansPlannerPlanItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
