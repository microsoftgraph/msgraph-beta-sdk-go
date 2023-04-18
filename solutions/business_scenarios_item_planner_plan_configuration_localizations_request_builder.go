package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder provides operations to manage the localizations property of the microsoft.graph.plannerPlanConfiguration entity.
type BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderGetQueryParameters get a list of the plannerPlanConfigurationLocalization objects and their properties.
type BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderGetQueryParameters struct {
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
// BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderGetQueryParameters
}
// BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPlannerPlanConfigurationLocalizationId provides operations to manage the localizations property of the microsoft.graph.plannerPlanConfiguration entity.
func (m *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) ByPlannerPlanConfigurationLocalizationId(plannerPlanConfigurationLocalizationId string)(*BusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if plannerPlanConfigurationLocalizationId != "" {
        urlTplParams["plannerPlanConfigurationLocalization%2Did"] = plannerPlanConfigurationLocalizationId
    }
    return NewBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderInternal instantiates a new LocalizationsRequestBuilder and sets the default values.
func NewBusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) {
    m := &BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/planConfiguration/localizations{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewBusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder instantiates a new LocalizationsRequestBuilder and sets the default values.
func NewBusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) Count()(*BusinessScenariosItemPlannerPlanConfigurationLocalizationsCountRequestBuilder) {
    return NewBusinessScenariosItemPlannerPlanConfigurationLocalizationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the plannerPlanConfigurationLocalization objects and their properties.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/plannerplanconfiguration-list-localizations?view=graph-rest-1.0
func (m *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Post create a new plannerPlanConfigurationLocalization object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/plannerplanconfiguration-post-localizations?view=graph-rest-1.0
func (m *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get a list of the plannerPlanConfigurationLocalization objects and their properties.
func (m *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPostRequestInformation create a new plannerPlanConfigurationLocalization object.
func (m *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *BusinessScenariosItemPlannerPlanConfigurationLocalizationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
