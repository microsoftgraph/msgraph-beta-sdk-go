package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partners/billing"
)

// PartnersBillingUsageUnbilledRequestBuilder provides operations to manage the unbilled property of the microsoft.graph.partners.billing.azureUsage entity.
type PartnersBillingUsageUnbilledRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnersBillingUsageUnbilledRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageUnbilledRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PartnersBillingUsageUnbilledRequestBuilderGetQueryParameters represents details for unbilled Azure usage data.
type PartnersBillingUsageUnbilledRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PartnersBillingUsageUnbilledRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageUnbilledRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnersBillingUsageUnbilledRequestBuilderGetQueryParameters
}
// PartnersBillingUsageUnbilledRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageUnbilledRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnersBillingUsageUnbilledRequestBuilderInternal instantiates a new PartnersBillingUsageUnbilledRequestBuilder and sets the default values.
func NewPartnersBillingUsageUnbilledRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageUnbilledRequestBuilder) {
    m := &PartnersBillingUsageUnbilledRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/partners/billing/usage/unbilled{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPartnersBillingUsageUnbilledRequestBuilder instantiates a new PartnersBillingUsageUnbilledRequestBuilder and sets the default values.
func NewPartnersBillingUsageUnbilledRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageUnbilledRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnersBillingUsageUnbilledRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property unbilled for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingUsageUnbilledRequestBuilder) Delete(ctx context.Context, requestConfiguration *PartnersBillingUsageUnbilledRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents details for unbilled Azure usage data.
// returns a UnbilledUsageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingUsageUnbilledRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnersBillingUsageUnbilledRequestBuilderGetRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.UnbilledUsageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.CreateUnbilledUsageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.UnbilledUsageable), nil
}
// MicrosoftGraphPartnersBillingExport provides operations to call the export method.
// returns a *PartnersBillingUsageUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder when successful
func (m *PartnersBillingUsageUnbilledRequestBuilder) MicrosoftGraphPartnersBillingExport()(*PartnersBillingUsageUnbilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    return NewPartnersBillingUsageUnbilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property unbilled in reports
// returns a UnbilledUsageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingUsageUnbilledRequestBuilder) Patch(ctx context.Context, body ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.UnbilledUsageable, requestConfiguration *PartnersBillingUsageUnbilledRequestBuilderPatchRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.UnbilledUsageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.CreateUnbilledUsageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.UnbilledUsageable), nil
}
// ToDeleteRequestInformation delete navigation property unbilled for reports
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageUnbilledRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PartnersBillingUsageUnbilledRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/reports/partners/billing/usage/unbilled", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents details for unbilled Azure usage data.
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageUnbilledRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnersBillingUsageUnbilledRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property unbilled in reports
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageUnbilledRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.UnbilledUsageable, requestConfiguration *PartnersBillingUsageUnbilledRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/reports/partners/billing/usage/unbilled", m.BaseRequestBuilder.PathParameters)
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
// returns a *PartnersBillingUsageUnbilledRequestBuilder when successful
func (m *PartnersBillingUsageUnbilledRequestBuilder) WithUrl(rawUrl string)(*PartnersBillingUsageUnbilledRequestBuilder) {
    return NewPartnersBillingUsageUnbilledRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
