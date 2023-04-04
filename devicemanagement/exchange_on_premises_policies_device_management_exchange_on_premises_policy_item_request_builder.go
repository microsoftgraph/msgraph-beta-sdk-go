package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
type ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetQueryParameters the list of Exchange On Premisis policies configured by the tenant.
type ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetQueryParameters
}
// ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagementExchangeOnPremisesPolicy entity.
func (m *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ConditionalAccessSettings()(*ExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilder) {
    return NewExchangeOnPremisesPoliciesItemConditionalAccessSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal instantiates a new DeviceManagementExchangeOnPremisesPolicyItemRequestBuilder and sets the default values.
func NewExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    m := &ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies/{deviceManagementExchangeOnPremisesPolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder instantiates a new DeviceManagementExchangeOnPremisesPolicyItemRequestBuilder and sets the default values.
func NewExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exchangeOnPremisesPolicies for deviceManagement
func (m *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get the list of Exchange On Premisis policies configured by the tenant.
func (m *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the list of Exchange On Premisis policies configured by the tenant.
func (m *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property exchangeOnPremisesPolicies in deviceManagement
func (m *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeOnPremisesPolicyable, requestConfiguration *ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
