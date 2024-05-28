package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorgroupConnectorGroupRequestBuilder provides operations to manage the connectorGroup property of the microsoft.graph.application entity.
type ItemConnectorgroupConnectorGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorgroupConnectorGroupRequestBuilderGetQueryParameters the connectorGroup the application is using with Microsoft Entra application proxy. Nullable.
type ItemConnectorgroupConnectorGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConnectorgroupConnectorGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorgroupConnectorGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConnectorgroupConnectorGroupRequestBuilderGetQueryParameters
}
// NewItemConnectorgroupConnectorGroupRequestBuilderInternal instantiates a new ItemConnectorgroupConnectorGroupRequestBuilder and sets the default values.
func NewItemConnectorgroupConnectorGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupConnectorGroupRequestBuilder) {
    m := &ItemConnectorgroupConnectorGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/connectorGroup{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemConnectorgroupConnectorGroupRequestBuilder instantiates a new ItemConnectorgroupConnectorGroupRequestBuilder and sets the default values.
func NewItemConnectorgroupConnectorGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupConnectorGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorgroupConnectorGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the connectorGroup the application is using with Microsoft Entra application proxy. Nullable.
// returns a ConnectorGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorgroupConnectorGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConnectorgroupConnectorGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectorGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable), nil
}
// Ref provides operations to manage the collection of application entities.
// returns a *ItemConnectorgroupRefRequestBuilder when successful
func (m *ItemConnectorgroupConnectorGroupRequestBuilder) Ref()(*ItemConnectorgroupRefRequestBuilder) {
    return NewItemConnectorgroupRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the connectorGroup the application is using with Microsoft Entra application proxy. Nullable.
// returns a *RequestInformation when successful
func (m *ItemConnectorgroupConnectorGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorgroupConnectorGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemConnectorgroupConnectorGroupRequestBuilder when successful
func (m *ItemConnectorgroupConnectorGroupRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorgroupConnectorGroupRequestBuilder) {
    return NewItemConnectorgroupConnectorGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
