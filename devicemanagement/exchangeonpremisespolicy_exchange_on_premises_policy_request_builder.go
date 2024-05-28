package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder provides operations to manage the exchangeOnPremisesPolicy property of the microsoft.graph.deviceManagement entity.
type ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderGetQueryParameters the policy which controls mobile device access to Exchange On Premises
type ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderGetQueryParameters
}
// ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagementExchangeOnPremisesPolicy entity.
// returns a *ExchangeonpremisespolicyConditionalaccesssettingsConditionalAccessSettingsRequestBuilder when successful
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) ConditionalAccessSettings()(*ExchangeonpremisespolicyConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) {
    return NewExchangeonpremisespolicyConditionalaccesssettingsConditionalAccessSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderInternal instantiates a new ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder and sets the default values.
func NewExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) {
    m := &ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder instantiates a new ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder and sets the default values.
func NewExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exchangeOnPremisesPolicy for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the policy which controls mobile device access to Exchange On Premises
// returns a DeviceManagementExchangeOnPremisesPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, error) {
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
// Patch update the navigation property exchangeOnPremisesPolicy in deviceManagement
// returns a DeviceManagementExchangeOnPremisesPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, error) {
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
// ToDeleteRequestInformation delete navigation property exchangeOnPremisesPolicy for deviceManagement
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the policy which controls mobile device access to Exchange On Premises
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property exchangeOnPremisesPolicy in deviceManagement
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder when successful
func (m *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) WithUrl(rawUrl string)(*ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) {
    return NewExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
