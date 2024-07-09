package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partners/billing"
)

// PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder provides operations to call the export method.
type PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal instantiates a new PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    m := &PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/partners/billing/usage/billed/microsoft.graph.partners.billing.export", pathParameters),
    }
    return m
}
// NewPartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder instantiates a new PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder and sets the default values.
func NewPartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post export the billed Azure usage data.
// returns a Operationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partners-billing-billedusage-export?view=graph-rest-beta
func (m *PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder) Post(ctx context.Context, body PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportExportPostRequestBodyable, requestConfiguration *PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.Operationable, error) {
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
// ToPostRequestInformation export the billed Azure usage data.
// returns a *RequestInformation when successful
func (m *PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportExportPostRequestBodyable, requestConfiguration *PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder when successful
func (m *PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder) WithUrl(rawUrl string)(*PartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder) {
    return NewPartnersBillingUsageBilledMicrosoftGraphPartnersBillingExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
