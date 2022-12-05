package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.plannerPlanConfiguration entity.
type SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetQueryParameters get localizations from solutions
type SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetQueryParameters
}
// SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal instantiates a new PlannerPlanConfigurationLocalizationItemRequestBuilder and sets the default values.
func NewSolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) {
    m := &SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/planConfiguration/localizations/{plannerPlanConfigurationLocalization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder instantiates a new PlannerPlanConfigurationLocalizationItemRequestBuilder and sets the default values.
func NewSolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property localizations for solutions
func (m *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get localizations from solutions
func (m *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property localizations in solutions
func (m *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property localizations for solutions
func (m *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get localizations from solutions
func (m *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable), nil
}
// Patch update the navigation property localizations in solutions
func (m *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *SolutionsBusinessScenariosItemPlannerPlanConfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable), nil
}
