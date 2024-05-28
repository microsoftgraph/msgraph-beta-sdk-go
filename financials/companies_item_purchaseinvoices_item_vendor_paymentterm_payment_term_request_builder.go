package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder provides operations to manage the paymentTerm property of the microsoft.graph.vendor entity.
type CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderGetQueryParameters get paymentTerm from financials
type CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderGetQueryParameters
}
// CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderInternal instantiates a new CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) {
    m := &CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices/{purchaseInvoice%2Did}/vendor/paymentTerm{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder instantiates a new CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property paymentTerm for financials
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get paymentTerm from financials
// returns a PaymentTermable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PaymentTermable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePaymentTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PaymentTermable), nil
}
// Patch update the navigation property paymentTerm in financials
// returns a PaymentTermable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PaymentTermable, requestConfiguration *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PaymentTermable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePaymentTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PaymentTermable), nil
}
// ToDeleteRequestInformation delete navigation property paymentTerm for financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get paymentTerm from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property paymentTerm in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PaymentTermable, requestConfiguration *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemVendorPaymenttermPaymentTermRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
