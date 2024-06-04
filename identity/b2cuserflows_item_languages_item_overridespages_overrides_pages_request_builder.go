package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder provides operations to manage the overridesPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
type B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderGetQueryParameters get the userFlowLanguagePage resources from the overridesPages navigation property. These pages are used to customize the values shown to the user during a user journey in a user flow.
type B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderGetQueryParameters struct {
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
// B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderGetQueryParameters
}
// B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserFlowLanguagePageId provides operations to manage the overridesPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
// returns a *B2cuserflowsItemLanguagesItemOverridespagesUserFlowLanguagePageItemRequestBuilder when successful
func (m *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) ByUserFlowLanguagePageId(userFlowLanguagePageId string)(*B2cuserflowsItemLanguagesItemOverridespagesUserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userFlowLanguagePageId != "" {
        urlTplParams["userFlowLanguagePage%2Did"] = userFlowLanguagePageId
    }
    return NewB2cuserflowsItemLanguagesItemOverridespagesUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewB2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderInternal instantiates a new B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder and sets the default values.
func NewB2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) {
    m := &B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/overridesPages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewB2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder instantiates a new B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder and sets the default values.
func NewB2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *B2cuserflowsItemLanguagesItemOverridespagesCountRequestBuilder when successful
func (m *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) Count()(*B2cuserflowsItemLanguagesItemOverridespagesCountRequestBuilder) {
    return NewB2cuserflowsItemLanguagesItemOverridespagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the userFlowLanguagePage resources from the overridesPages navigation property. These pages are used to customize the values shown to the user during a user journey in a user flow.
// returns a UserFlowLanguagePageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/userflowlanguageconfiguration-list-overridespages?view=graph-rest-beta
func (m *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguagePageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageCollectionResponseable), nil
}
// Post create new navigation property to overridesPages for identity
// returns a UserFlowLanguagePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFlowLanguagePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable), nil
}
// ToGetRequestInformation get the userFlowLanguagePage resources from the overridesPages navigation property. These pages are used to customize the values shown to the user during a user journey in a user flow.
// returns a *RequestInformation when successful
func (m *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to overridesPages for identity
// returns a *RequestInformation when successful
func (m *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder when successful
func (m *B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) WithUrl(rawUrl string)(*B2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder) {
    return NewB2cuserflowsItemLanguagesItemOverridespagesOverridesPagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
