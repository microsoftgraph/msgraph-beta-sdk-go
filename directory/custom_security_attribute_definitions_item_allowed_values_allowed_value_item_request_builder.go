package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder provides operations to manage the allowedValues property of the microsoft.graph.customSecurityAttributeDefinition entity.
type CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderGetQueryParameters read the properties and relationships of an allowedValue object.
type CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderGetQueryParameters
}
// CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderInternal instantiates a new CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder and sets the default values.
func NewCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) {
    m := &CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/customSecurityAttributeDefinitions/{customSecurityAttributeDefinition%2Did}/allowedValues/{allowedValue%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder instantiates a new CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder and sets the default values.
func NewCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property allowedValues for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an allowedValue object.
// returns a AllowedValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/allowedvalue-get?view=graph-rest-1.0
func (m *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AllowedValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAllowedValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AllowedValueable), nil
}
// Patch update the properties of an allowedValue object.
// returns a AllowedValueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/allowedvalue-update?view=graph-rest-1.0
func (m *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AllowedValueable, requestConfiguration *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AllowedValueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAllowedValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AllowedValueable), nil
}
// ToDeleteRequestInformation delete navigation property allowedValues for directory
// returns a *RequestInformation when successful
func (m *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an allowedValue object.
// returns a *RequestInformation when successful
func (m *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an allowedValue object.
// returns a *RequestInformation when successful
func (m *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AllowedValueable, requestConfiguration *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder when successful
func (m *CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) WithUrl(rawUrl string)(*CustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder) {
    return NewCustomSecurityAttributeDefinitionsItemAllowedValuesAllowedValueItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
