package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder provides operations to manage the retentionEventTypes property of the microsoft.graph.security.triggerTypesRoot entity.
type TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderGetQueryParameters read the properties and relationships of a retentionEventType object.
type TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderGetQueryParameters
}
// TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderInternal instantiates a new TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder and sets the default values.
func NewTriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) {
    m := &TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/triggerTypes/retentionEventTypes/{retentionEventType%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder instantiates a new TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder and sets the default values.
func NewTriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a retentionEventType object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentioneventtype-delete?view=graph-rest-1.0
func (m *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a retentionEventType object.
// returns a RetentionEventTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentioneventtype-get?view=graph-rest-1.0
func (m *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateRetentionEventTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable), nil
}
// Patch update the navigation property retentionEventTypes in security
// returns a RetentionEventTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, requestConfiguration *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateRetentionEventTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable), nil
}
// ToDeleteRequestInformation delete a retentionEventType object.
// returns a *RequestInformation when successful
func (m *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/security/triggerTypes/retentionEventTypes/{retentionEventType%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a retentionEventType object.
// returns a *RequestInformation when successful
func (m *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property retentionEventTypes in security
// returns a *RequestInformation when successful
func (m *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, requestConfiguration *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/security/triggerTypes/retentionEventTypes/{retentionEventType%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder when successful
func (m *TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) WithUrl(rawUrl string)(*TriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder) {
    return NewTriggerTypesRetentionEventTypesRetentionEventTypeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
