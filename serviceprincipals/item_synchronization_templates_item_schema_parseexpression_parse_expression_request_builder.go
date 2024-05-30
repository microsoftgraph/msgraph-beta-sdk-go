package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder provides operations to call the parseExpression method.
type ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilderInternal instantiates a new ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder and sets the default values.
func NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder) {
    m := &ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/templates/{synchronizationTemplate%2Did}/schema/parseExpression", pathParameters),
    }
    return m
}
// NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder instantiates a new ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder and sets the default values.
func NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post parse a given string expression into an attributeMappingSource object. For more information about expressions, see Writing Expressions for Attribute Mappings in Microsoft Entra ID.
// returns a ParseExpressionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationschema-parseexpression?view=graph-rest-beta
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder) Post(ctx context.Context, body ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBodyable, requestConfiguration *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseExpressionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateParseExpressionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseExpressionResponseable), nil
}
// ToPostRequestInformation parse a given string expression into an attributeMappingSource object. For more information about expressions, see Writing Expressions for Attribute Mappings in Microsoft Entra ID.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBodyable, requestConfiguration *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder when successful
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder) {
    return NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
