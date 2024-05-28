package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder provides operations to call the filterOperators method.
type ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetQueryParameters list all operators supported in the scoping filters.
type ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetQueryParameters struct {
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
// ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetQueryParameters
}
// NewItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderInternal instantiates a new ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder and sets the default values.
func NewItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder) {
    m := &ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/templates/{synchronizationTemplate%2Did}/schema/filterOperators(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder instantiates a new ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder and sets the default values.
func NewItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all operators supported in the scoping filters.
// Deprecated: This method is obsolete. Use GetAsFilterOperatorsGetResponse instead.
// returns a ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationschema-filteroperators?view=graph-rest-beta
func (m *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetRequestConfiguration)(ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsResponseable), nil
}
// GetAsFilterOperatorsGetResponse list all operators supported in the scoping filters.
// returns a ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationschema-filteroperators?view=graph-rest-beta
func (m *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder) GetAsFilterOperatorsGetResponse(ctx context.Context, requestConfiguration *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetRequestConfiguration)(ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsGetResponseable), nil
}
// ToGetRequestInformation list all operators supported in the scoping filters.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder when successful
func (m *ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder) {
    return NewItemSynchronizationTemplatesItemSchemaFilteroperatorsFilterOperatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
