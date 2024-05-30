package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder provides operations to manage the regionalAndLanguageSettings property of the microsoft.graph.userSettings entity.
type ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderGetQueryParameters the user's preferences for languages, regional locale and date/time formatting.
type ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderGetQueryParameters
}
// ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderInternal instantiates a new ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder and sets the default values.
func NewItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) {
    m := &ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/settings/regionalAndLanguageSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder instantiates a new ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder and sets the default values.
func NewItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property regionalAndLanguageSettings for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the user's preferences for languages, regional locale and date/time formatting.
// returns a RegionalAndLanguageSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RegionalAndLanguageSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRegionalAndLanguageSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RegionalAndLanguageSettingsable), nil
}
// Patch update the navigation property regionalAndLanguageSettings in users
// returns a RegionalAndLanguageSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RegionalAndLanguageSettingsable, requestConfiguration *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RegionalAndLanguageSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRegionalAndLanguageSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RegionalAndLanguageSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property regionalAndLanguageSettings for users
// returns a *RequestInformation when successful
func (m *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the user's preferences for languages, regional locale and date/time formatting.
// returns a *RequestInformation when successful
func (m *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property regionalAndLanguageSettings in users
// returns a *RequestInformation when successful
func (m *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RegionalAndLanguageSettingsable, requestConfiguration *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder when successful
func (m *ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder) {
    return NewItemSettingsRegionalandlanguagesettingsRegionalAndLanguageSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
