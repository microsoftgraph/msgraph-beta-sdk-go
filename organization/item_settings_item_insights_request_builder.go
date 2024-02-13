package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSettingsItemInsightsRequestBuilder provides operations to manage the itemInsights property of the microsoft.graph.organizationSettings entity.
type ItemSettingsItemInsightsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsItemInsightsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsItemInsightsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSettingsItemInsightsRequestBuilderGetQueryParameters get the properties of an insightsSettings object for displaying or returning item insights in an organization. To learn how to customize the privacy of item insights in an organization, see Customize item insights privacy. 
type ItemSettingsItemInsightsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSettingsItemInsightsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsItemInsightsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSettingsItemInsightsRequestBuilderGetQueryParameters
}
// ItemSettingsItemInsightsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsItemInsightsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSettingsItemInsightsRequestBuilderInternal instantiates a new ItemSettingsItemInsightsRequestBuilder and sets the default values.
func NewItemSettingsItemInsightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsItemInsightsRequestBuilder) {
    m := &ItemSettingsItemInsightsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/settings/itemInsights{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSettingsItemInsightsRequestBuilder instantiates a new ItemSettingsItemInsightsRequestBuilder and sets the default values.
func NewItemSettingsItemInsightsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsItemInsightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsItemInsightsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property itemInsights for organization
// Deprecated: The Organization ItemInsights endpoint will stop returning data on January 1st, 2024. Please use the new Admin People ItemInsights endpoint. as of 2023-10/Beta:ItemInsightsOranizationSettings
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsItemInsightsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSettingsItemInsightsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties of an insightsSettings object for displaying or returning item insights in an organization. To learn how to customize the privacy of item insights in an organization, see Customize item insights privacy. 
// Deprecated: The Organization ItemInsights endpoint will stop returning data on January 1st, 2024. Please use the new Admin People ItemInsights endpoint. as of 2023-10/Beta:ItemInsightsOranizationSettings
// returns a InsightsSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/organizationsettings-list-iteminsights?view=graph-rest-1.0
func (m *ItemSettingsItemInsightsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSettingsItemInsightsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightsSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInsightsSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightsSettingsable), nil
}
// Patch update privacy settings to display or return the specified type of insights in an organization. The type of settings can be contact insights, item insights, or people insights. To learn more about customizing insights privacy for your organization, see:-  Customize item insights privacy -  Customize people insights privacy
// Deprecated: The Organization ItemInsights endpoint will stop returning data on January 1st, 2024. Please use the new Admin People ItemInsights endpoint. as of 2023-10/Beta:ItemInsightsOranizationSettings
// returns a InsightsSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/insightssettings-update?view=graph-rest-1.0
func (m *ItemSettingsItemInsightsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightsSettingsable, requestConfiguration *ItemSettingsItemInsightsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightsSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInsightsSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightsSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property itemInsights for organization
// Deprecated: The Organization ItemInsights endpoint will stop returning data on January 1st, 2024. Please use the new Admin People ItemInsights endpoint. as of 2023-10/Beta:ItemInsightsOranizationSettings
// returns a *RequestInformation when successful
func (m *ItemSettingsItemInsightsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsItemInsightsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/organization/{organization%2Did}/settings/itemInsights", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties of an insightsSettings object for displaying or returning item insights in an organization. To learn how to customize the privacy of item insights in an organization, see Customize item insights privacy. 
// Deprecated: The Organization ItemInsights endpoint will stop returning data on January 1st, 2024. Please use the new Admin People ItemInsights endpoint. as of 2023-10/Beta:ItemInsightsOranizationSettings
// returns a *RequestInformation when successful
func (m *ItemSettingsItemInsightsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsItemInsightsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update privacy settings to display or return the specified type of insights in an organization. The type of settings can be contact insights, item insights, or people insights. To learn more about customizing insights privacy for your organization, see:-  Customize item insights privacy -  Customize people insights privacy
// Deprecated: The Organization ItemInsights endpoint will stop returning data on January 1st, 2024. Please use the new Admin People ItemInsights endpoint. as of 2023-10/Beta:ItemInsightsOranizationSettings
// returns a *RequestInformation when successful
func (m *ItemSettingsItemInsightsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightsSettingsable, requestConfiguration *ItemSettingsItemInsightsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/organization/{organization%2Did}/settings/itemInsights", m.BaseRequestBuilder.PathParameters)
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
// Deprecated: The Organization ItemInsights endpoint will stop returning data on January 1st, 2024. Please use the new Admin People ItemInsights endpoint. as of 2023-10/Beta:ItemInsightsOranizationSettings
// returns a *ItemSettingsItemInsightsRequestBuilder when successful
func (m *ItemSettingsItemInsightsRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsItemInsightsRequestBuilder) {
    return NewItemSettingsItemInsightsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
