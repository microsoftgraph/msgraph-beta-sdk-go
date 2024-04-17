package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partners/billing"
)

// PartnersBillingUsageBilledRequestBuilder provides operations to manage the billed property of the microsoft.graph.partners.billing.azureUsage entity.
type PartnersBillingUsageBilledRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnersBillingUsageBilledRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageBilledRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PartnersBillingUsageBilledRequestBuilderGetQueryParameters represents details for billed Azure usage data.
type PartnersBillingUsageBilledRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PartnersBillingUsageBilledRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageBilledRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnersBillingUsageBilledRequestBuilderGetQueryParameters
}
// PartnersBillingUsageBilledRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageBilledRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnersBillingUsageBilledRequestBuilderInternal instantiates a new PartnersBillingUsageBilledRequestBuilder and sets the default values.
func NewPartnersBillingUsageBilledRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageBilledRequestBuilder) {
    m := &PartnersBillingUsageBilledRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/partners/billing/usage/billed{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPartnersBillingUsageBilledRequestBuilder instantiates a new PartnersBillingUsageBilledRequestBuilder and sets the default values.
func NewPartnersBillingUsageBilledRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageBilledRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnersBillingUsageBilledRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property billed for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingUsageBilledRequestBuilder) Delete(ctx context.Context, requestConfiguration *PartnersBillingUsageBilledRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents details for billed Azure usage data.
// returns a BilledUsageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingUsageBilledRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnersBillingUsageBilledRequestBuilderGetRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BilledUsageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.CreateBilledUsageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BilledUsageable), nil
}
// MicrosoftGraphPartnersBillingExport provides operations to call the export method.
// returns a *PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder when successful
func (m *PartnersBillingUsageBilledRequestBuilder) MicrosoftGraphPartnersBillingExport()(*PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    return NewPartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property billed in reports
// returns a BilledUsageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingUsageBilledRequestBuilder) Patch(ctx context.Context, body ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BilledUsageable, requestConfiguration *PartnersBillingUsageBilledRequestBuilderPatchRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BilledUsageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.CreateBilledUsageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BilledUsageable), nil
}
// ToDeleteRequestInformation delete navigation property billed for reports
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageBilledRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PartnersBillingUsageBilledRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents details for billed Azure usage data.
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageBilledRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnersBillingUsageBilledRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property billed in reports
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageBilledRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BilledUsageable, requestConfiguration *PartnersBillingUsageBilledRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnersBillingUsageBilledRequestBuilder when successful
func (m *PartnersBillingUsageBilledRequestBuilder) WithUrl(rawUrl string)(*PartnersBillingUsageBilledRequestBuilder) {
    return NewPartnersBillingUsageBilledRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
