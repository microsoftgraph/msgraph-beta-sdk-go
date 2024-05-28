package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
type ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderGetQueryParameters the list of Exchange On Premisis policies configured by the tenant.
type ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderGetQueryParameters struct {
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
// ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderGetQueryParameters
}
// ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementExchangeOnPremisesPolicyId provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder when successful
func (m *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) ByDeviceManagementExchangeOnPremisesPolicyId(deviceManagementExchangeOnPremisesPolicyId string)(*ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementExchangeOnPremisesPolicyId != "" {
        urlTplParams["deviceManagementExchangeOnPremisesPolicy%2Did"] = deviceManagementExchangeOnPremisesPolicyId
    }
    return NewExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderInternal instantiates a new ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder and sets the default values.
func NewExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) {
    m := &ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder instantiates a new ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder and sets the default values.
func NewExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ExchangeonpremisespoliciesCountRequestBuilder when successful
func (m *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) Count()(*ExchangeonpremisespoliciesCountRequestBuilder) {
    return NewExchangeonpremisespoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of Exchange On Premisis policies configured by the tenant.
// returns a DeviceManagementExchangeOnPremisesPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementExchangeOnPremisesPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyCollectionResponseable), nil
}
// Post create new navigation property to exchangeOnPremisesPolicies for deviceManagement
// returns a DeviceManagementExchangeOnPremisesPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable), nil
}
// ToGetRequestInformation the list of Exchange On Premisis policies configured by the tenant.
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to exchangeOnPremisesPolicies for deviceManagement
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder when successful
func (m *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) WithUrl(rawUrl string)(*ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) {
    return NewExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
