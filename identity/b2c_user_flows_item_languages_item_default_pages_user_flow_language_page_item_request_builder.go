package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder provides operations to manage the defaultPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
type B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters collection of pages with the default content to display in a user flow for a specified language. This collection doesn't allow any kind of modification.
type B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters
}
// B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal instantiates a new B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    m := &B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder instantiates a new B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the identityContainer entity.
// returns a *B2cUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder when successful
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Content()(*B2cUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) {
    return NewB2cUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property defaultPages for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of pages with the default content to display in a user flow for a specified language. This collection doesn't allow any kind of modification.
// returns a UserFlowLanguagePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property defaultPages in identity
// returns a UserFlowLanguagePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property defaultPages for identity
// returns a *RequestInformation when successful
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of pages with the default content to display in a user flow for a specified language. This collection doesn't allow any kind of modification.
// returns a *RequestInformation when successful
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property defaultPages in identity
// returns a *RequestInformation when successful
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserFlowLanguagePageable, requestConfiguration *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder when successful
func (m *B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) WithUrl(rawUrl string)(*B2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    return NewB2cUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
