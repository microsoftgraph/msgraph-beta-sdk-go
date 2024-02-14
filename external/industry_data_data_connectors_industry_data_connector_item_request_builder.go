package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder provides operations to manage the dataConnectors property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderGetQueryParameters read the properties and relationships of an industryDataConnector object.
type IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderGetQueryParameters
}
// IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderInternal instantiates a new IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder and sets the default values.
func NewIndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) {
    m := &IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/dataConnectors/{industryDataConnector%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder instantiates a new IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder and sets the default values.
func NewIndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an industryDataConnector object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydataconnector-delete?view=graph-rest-1.0
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an industryDataConnector object.
// returns a IndustryDataConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydataconnector-get?view=graph-rest-1.0
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable), nil
}
// MicrosoftGraphIndustryDataValidate provides operations to call the validate method.
// returns a *IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder when successful
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) MicrosoftGraphIndustryDataValidate()(*IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder) {
    return NewIndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an industryDataConnector object.
// returns a IndustryDataConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydataconnector-update?view=graph-rest-1.0
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) Patch(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable, requestConfiguration *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderPatchRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable), nil
}
// SourceSystem provides operations to manage the sourceSystem property of the microsoft.graph.industryData.industryDataConnector entity.
// returns a *IndustryDataDataConnectorsItemSourceSystemRequestBuilder when successful
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) SourceSystem()(*IndustryDataDataConnectorsItemSourceSystemRequestBuilder) {
    return NewIndustryDataDataConnectorsItemSourceSystemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete an industryDataConnector object.
// returns a *RequestInformation when successful
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/external/industryData/dataConnectors/{industryDataConnector%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an industryDataConnector object.
// returns a *RequestInformation when successful
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an industryDataConnector object.
// returns a *RequestInformation when successful
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable, requestConfiguration *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/external/industryData/dataConnectors/{industryDataConnector%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder when successful
func (m *IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) WithUrl(rawUrl string)(*IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) {
    return NewIndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
