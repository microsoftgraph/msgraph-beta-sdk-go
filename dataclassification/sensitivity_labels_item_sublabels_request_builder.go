package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SensitivityLabelsItemSublabelsRequestBuilder provides operations to manage the sublabels property of the microsoft.graph.sensitivityLabel entity.
type SensitivityLabelsItemSublabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SensitivityLabelsItemSublabelsRequestBuilderGetQueryParameters get sublabels from dataClassification
type SensitivityLabelsItemSublabelsRequestBuilderGetQueryParameters struct {
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
// SensitivityLabelsItemSublabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SensitivityLabelsItemSublabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SensitivityLabelsItemSublabelsRequestBuilderGetQueryParameters
}
// SensitivityLabelsItemSublabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SensitivityLabelsItemSublabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySensitivityLabelId1 provides operations to manage the sublabels property of the microsoft.graph.sensitivityLabel entity.
// returns a *SensitivityLabelsItemSublabelsSensitivityLabelItemRequestBuilder when successful
func (m *SensitivityLabelsItemSublabelsRequestBuilder) BySensitivityLabelId1(sensitivityLabelId1 string)(*SensitivityLabelsItemSublabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sensitivityLabelId1 != "" {
        urlTplParams["sensitivityLabel%2Did1"] = sensitivityLabelId1
    }
    return NewSensitivityLabelsItemSublabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSensitivityLabelsItemSublabelsRequestBuilderInternal instantiates a new SensitivityLabelsItemSublabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsItemSublabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensitivityLabelsItemSublabelsRequestBuilder) {
    m := &SensitivityLabelsItemSublabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/sensitivityLabels/{sensitivityLabel%2Did}/sublabels{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewSensitivityLabelsItemSublabelsRequestBuilder instantiates a new SensitivityLabelsItemSublabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsItemSublabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensitivityLabelsItemSublabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSensitivityLabelsItemSublabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *SensitivityLabelsItemSublabelsCountRequestBuilder when successful
func (m *SensitivityLabelsItemSublabelsRequestBuilder) Count()(*SensitivityLabelsItemSublabelsCountRequestBuilder) {
    return NewSensitivityLabelsItemSublabelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Evaluate provides operations to call the evaluate method.
// returns a *SensitivityLabelsItemSublabelsEvaluateRequestBuilder when successful
func (m *SensitivityLabelsItemSublabelsRequestBuilder) Evaluate()(*SensitivityLabelsItemSublabelsEvaluateRequestBuilder) {
    return NewSensitivityLabelsItemSublabelsEvaluateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get sublabels from dataClassification
// returns a SensitivityLabelCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SensitivityLabelsItemSublabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *SensitivityLabelsItemSublabelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityLabelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelCollectionResponseable), nil
}
// Post create new navigation property to sublabels for dataClassification
// returns a SensitivityLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SensitivityLabelsItemSublabelsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, requestConfiguration *SensitivityLabelsItemSublabelsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable), nil
}
// ToGetRequestInformation get sublabels from dataClassification
// returns a *RequestInformation when successful
func (m *SensitivityLabelsItemSublabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SensitivityLabelsItemSublabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to sublabels for dataClassification
// returns a *RequestInformation when successful
func (m *SensitivityLabelsItemSublabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, requestConfiguration *SensitivityLabelsItemSublabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/dataClassification/sensitivityLabels/{sensitivityLabel%2Did}/sublabels", m.BaseRequestBuilder.PathParameters)
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
// returns a *SensitivityLabelsItemSublabelsRequestBuilder when successful
func (m *SensitivityLabelsItemSublabelsRequestBuilder) WithUrl(rawUrl string)(*SensitivityLabelsItemSublabelsRequestBuilder) {
    return NewSensitivityLabelsItemSublabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
