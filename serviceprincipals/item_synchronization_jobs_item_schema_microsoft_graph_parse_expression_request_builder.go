package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder provides operations to call the parseExpression method.
type ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilderInternal instantiates a new MicrosoftGraphParseExpressionRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder) {
    m := &ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/jobs/{synchronizationJob%2Did}/schema/microsoft.graph.parseExpression";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder instantiates a new MicrosoftGraphParseExpressionRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post parse a given string expression into an attributeMappingSource object. For more information about expressions, see Writing Expressions for Attribute Mappings in Azure Active Directory.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/synchronization-synchronizationschema-parseexpression?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseExpressionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateParseExpressionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseExpressionResponseable), nil
}
// ToPostRequestInformation parse a given string expression into an attributeMappingSource object. For more information about expressions, see Writing Expressions for Attribute Mappings in Azure Active Directory.
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
