package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder provides operations to manage the templates property of the microsoft.graph.conditionalAccessRoot entity.
type ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderGetQueryParameters read the properties and relationships of a conditionalAccessTemplate object.
type ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderGetQueryParameters
}
// NewConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal instantiates a new ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder and sets the default values.
func NewConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    m := &ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/templates/{conditionalAccessTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder instantiates a new ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder and sets the default values.
func NewConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a conditionalAccessTemplate object.
// returns a ConditionalAccessTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conditionalaccesstemplate-get?view=graph-rest-1.0
func (m *ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConditionalAccessTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessTemplateable), nil
}
// ToGetRequestInformation read the properties and relationships of a conditionalAccessTemplate object.
// returns a *RequestInformation when successful
func (m *ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder when successful
func (m *ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) WithUrl(rawUrl string)(*ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    return NewConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
