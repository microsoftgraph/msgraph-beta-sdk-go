package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partners/billing"
)

// PartnersBillingReconciliationRequestBuilder provides operations to manage the reconciliation property of the microsoft.graph.partners.billing.billing entity.
type PartnersBillingReconciliationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnersBillingReconciliationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingReconciliationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PartnersBillingReconciliationRequestBuilderGetQueryParameters represents details for billed invoice reconciliation data.
type PartnersBillingReconciliationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PartnersBillingReconciliationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingReconciliationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnersBillingReconciliationRequestBuilderGetQueryParameters
}
// PartnersBillingReconciliationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnersBillingReconciliationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Billed provides operations to manage the billed property of the microsoft.graph.partners.billing.billingReconciliation entity.
// returns a *PartnersBillingReconciliationBilledRequestBuilder when successful
func (m *PartnersBillingReconciliationRequestBuilder) Billed()(*PartnersBillingReconciliationBilledRequestBuilder) {
    return NewPartnersBillingReconciliationBilledRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPartnersBillingReconciliationRequestBuilderInternal instantiates a new PartnersBillingReconciliationRequestBuilder and sets the default values.
func NewPartnersBillingReconciliationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingReconciliationRequestBuilder) {
    m := &PartnersBillingReconciliationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/partners/billing/reconciliation{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPartnersBillingReconciliationRequestBuilder instantiates a new PartnersBillingReconciliationRequestBuilder and sets the default values.
func NewPartnersBillingReconciliationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnersBillingReconciliationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnersBillingReconciliationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property reconciliation for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingReconciliationRequestBuilder) Delete(ctx context.Context, requestConfiguration *PartnersBillingReconciliationRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents details for billed invoice reconciliation data.
// returns a BillingReconciliationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingReconciliationRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnersBillingReconciliationRequestBuilderGetRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BillingReconciliationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.CreateBillingReconciliationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BillingReconciliationable), nil
}
// Patch update the navigation property reconciliation in reports
// returns a BillingReconciliationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnersBillingReconciliationRequestBuilder) Patch(ctx context.Context, body ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BillingReconciliationable, requestConfiguration *PartnersBillingReconciliationRequestBuilderPatchRequestConfiguration)(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BillingReconciliationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.CreateBillingReconciliationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BillingReconciliationable), nil
}
// ToDeleteRequestInformation delete navigation property reconciliation for reports
// returns a *RequestInformation when successful
func (m *PartnersBillingReconciliationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PartnersBillingReconciliationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents details for billed invoice reconciliation data.
// returns a *RequestInformation when successful
func (m *PartnersBillingReconciliationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnersBillingReconciliationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property reconciliation in reports
// returns a *RequestInformation when successful
func (m *PartnersBillingReconciliationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ibc1d41e06c1155c6505b27a1d0c17b20692636238b0ddb2acdaa2b4548e4f67b.BillingReconciliationable, requestConfiguration *PartnersBillingReconciliationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnersBillingReconciliationRequestBuilder when successful
func (m *PartnersBillingReconciliationRequestBuilder) WithUrl(rawUrl string)(*PartnersBillingReconciliationRequestBuilder) {
    return NewPartnersBillingReconciliationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
