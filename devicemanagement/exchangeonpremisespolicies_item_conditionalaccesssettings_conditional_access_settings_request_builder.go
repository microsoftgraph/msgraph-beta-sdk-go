package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagementExchangeOnPremisesPolicy entity.
type ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderGetQueryParameters the Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
type ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderGetQueryParameters
}
// ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderInternal instantiates a new ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder and sets the default values.
func NewExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) {
    m := &ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies/{deviceManagementExchangeOnPremisesPolicy%2Did}/conditionalAccessSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder instantiates a new ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder and sets the default values.
func NewExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property conditionalAccessSettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
// returns a OnPremisesConditionalAccessSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable), nil
}
// Patch update the navigation property conditionalAccessSettings in deviceManagement
// returns a OnPremisesConditionalAccessSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, requestConfiguration *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property conditionalAccessSettings for deviceManagement
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property conditionalAccessSettings in deviceManagement
// returns a *RequestInformation when successful
func (m *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, requestConfiguration *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder when successful
func (m *ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) WithUrl(rawUrl string)(*ExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) {
    return NewExchangeonpremisespoliciesItemConditionalaccesssettingsConditionalAccessSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
