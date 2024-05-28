package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partners/billing"
)

// PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder provides operations to call the export method.
type PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderInternal instantiates a new PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) {
    m := &PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/partners/billing/reconciliation/billed/microsoft.graph.partners.billing.export", pathParameters),
    }
    return m
}
// NewPartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder instantiates a new PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post export the billed invoice reconciliation data.
// returns a Operationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partners-billing-billedreconciliation-export?view=graph-rest-beta
func (m *PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) Post(ctx context.Context, body PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportExportPostRequestBodyable, requestConfiguration *PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.Operationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.CreateOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.Operationable), nil
}
// ToPostRequestInformation export the billed invoice reconciliation data.
// returns a *RequestInformation when successful
func (m *PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportExportPostRequestBodyable, requestConfiguration *PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder when successful
func (m *PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) WithUrl(rawUrl string)(*PartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder) {
    return NewPartnersBillingReconciliationBilledMicrosoftgraphpartnersbillingexportMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
