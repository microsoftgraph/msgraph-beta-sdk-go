package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagementExchangeOnPremisesPolicy entity.
type ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderGetQueryParameters the Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
type ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderGetQueryParameters
}
// ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderInternal instantiates a new ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder and sets the default values.
func NewExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) {
    m := &ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies/{deviceManagementExchangeOnPremisesPolicy%2Did}/conditionalAccessSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder instantiates a new ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder and sets the default values.
func NewExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property conditionalAccessSettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, error) {
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
func (m *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, requestConfiguration *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, error) {
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
func (m *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies/{deviceManagementExchangeOnPremisesPolicy%2Did}/conditionalAccessSettings", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
// returns a *RequestInformation when successful
func (m *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesConditionalAccessSettingsable, requestConfiguration *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies/{deviceManagementExchangeOnPremisesPolicy%2Did}/conditionalAccessSettings", m.BaseRequestBuilder.PathParameters)
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
// returns a *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder when successful
func (m *ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) WithUrl(rawUrl string)(*ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) {
    return NewExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
