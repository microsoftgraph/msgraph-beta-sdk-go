package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder provides operations to manage the shiftsRoleDefinitions property of the microsoft.graph.schedule entity.
type ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetQueryParameters get shiftsRoleDefinitions from groups
type ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetQueryParameters
}
// ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderInternal instantiates a new ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder and sets the default values.
func NewItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) {
    m := &ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/schedule/shiftsRoleDefinitions/{shiftsRoleDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder instantiates a new ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder and sets the default values.
func NewItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property shiftsRoleDefinitions for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get shiftsRoleDefinitions from groups
// returns a ShiftsRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateShiftsRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable), nil
}
// Patch update the navigation property shiftsRoleDefinitions in groups
// returns a ShiftsRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateShiftsRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable), nil
}
// ToDeleteRequestInformation delete navigation property shiftsRoleDefinitions for groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/groups/{group%2Did}/team/schedule/shiftsRoleDefinitions/{shiftsRoleDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get shiftsRoleDefinitions from groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property shiftsRoleDefinitions in groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/groups/{group%2Did}/team/schedule/shiftsRoleDefinitions/{shiftsRoleDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) {
    return NewItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
