package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder provides operations to manage the shiftsRoleDefinitions property of the microsoft.graph.schedule entity.
type ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetQueryParameters get shiftsRoleDefinitions from teams
type ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetQueryParameters
}
// ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderInternal instantiates a new ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder and sets the default values.
func NewItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) {
    m := &ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/schedule/shiftsRoleDefinitions/{shiftsRoleDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder instantiates a new ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder and sets the default values.
func NewItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property shiftsRoleDefinitions for teams
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get shiftsRoleDefinitions from teams
// returns a ShiftsRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, error) {
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
// Patch update the navigation property shiftsRoleDefinitions in teams
// returns a ShiftsRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, requestConfiguration *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, error) {
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
// ToDeleteRequestInformation delete navigation property shiftsRoleDefinitions for teams
// returns a *RequestInformation when successful
func (m *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/teams/{team%2Did}/schedule/shiftsRoleDefinitions/{shiftsRoleDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get shiftsRoleDefinitions from teams
// returns a *RequestInformation when successful
func (m *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property shiftsRoleDefinitions in teams
// returns a *RequestInformation when successful
func (m *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, requestConfiguration *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/teams/{team%2Did}/schedule/shiftsRoleDefinitions/{shiftsRoleDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder when successful
func (m *ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*ItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) {
    return NewItemScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
