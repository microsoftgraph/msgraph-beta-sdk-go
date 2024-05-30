package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
type ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetQueryParameters the list of Exchange On Premisis policies configured by the tenant.
type ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetQueryParameters
}
// ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagementExchangeOnPremisesPolicy entity.
// returns a *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder when successful
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ConditionalAccessSettings()(*ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) {
    return NewExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal instantiates a new ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder and sets the default values.
func NewExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    m := &ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies/{deviceManagementExchangeOnPremisesPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder instantiates a new ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder and sets the default values.
func NewExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exchangeOnPremisesPolicies for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of Exchange On Premisis policies configured by the tenant.
// returns a DeviceManagementExchangeOnPremisesPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property exchangeOnPremisesPolicies in deviceManagement
// returns a DeviceManagementExchangeOnPremisesPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property exchangeOnPremisesPolicies for deviceManagement
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of Exchange On Premisis policies configured by the tenant.
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property exchangeOnPremisesPolicies in deviceManagement
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder when successful
func (m *ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    return NewExchangeonpremisespoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
