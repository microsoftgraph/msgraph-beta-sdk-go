package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder provides operations to manage the localizations property of the microsoft.graph.plannerPlanConfiguration entity.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderGetQueryParameters localized names for the plan configuration.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderGetQueryParameters
}
// BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPlannerPlanConfigurationLocalizationId provides operations to manage the localizations property of the microsoft.graph.plannerPlanConfiguration entity.
// returns a *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) ByPlannerPlanConfigurationLocalizationId(plannerPlanConfigurationLocalizationId string)(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if plannerPlanConfigurationLocalizationId != "" {
        urlTplParams["plannerPlanConfigurationLocalization%2Did"] = plannerPlanConfigurationLocalizationId
    }
    return NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderInternal instantiates a new BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) {
    m := &BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/planConfiguration/localizations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder instantiates a new BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BusinessscenariosItemPlannerPlanconfigurationLocalizationsCountRequestBuilder when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) Count()(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsCountRequestBuilder) {
    return NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get localized names for the plan configuration.
// returns a PlannerPlanConfigurationLocalizationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationLocalizationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationCollectionResponseable), nil
}
// Post create new navigation property to localizations for solutions
// returns a PlannerPlanConfigurationLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable), nil
}
// ToGetRequestInformation localized names for the plan configuration.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to localizations for solutions
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) WithUrl(rawUrl string)(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) {
    return NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
