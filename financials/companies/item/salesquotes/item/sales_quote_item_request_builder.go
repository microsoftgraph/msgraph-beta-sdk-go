package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/currency"
    i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/paymentterm"
    i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/customer"
    i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/makeinvoice"
    if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/send"
    iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/shipmentmethod"
    iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines"
    id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines/item"
)

// SalesQuoteItemRequestBuilder provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
type SalesQuoteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SalesQuoteItemRequestBuilderGetQueryParameters get salesQuotes from financials
type SalesQuoteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SalesQuoteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesQuoteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SalesQuoteItemRequestBuilderGetQueryParameters
}
// SalesQuoteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesQuoteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSalesQuoteItemRequestBuilderInternal instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewSalesQuoteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesQuoteItemRequestBuilder) {
    m := &SalesQuoteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesQuotes/{salesQuote%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesQuoteItemRequestBuilder instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewSalesQuoteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesQuoteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesQuoteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesQuotes from financials
func (m *SalesQuoteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SalesQuoteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property salesQuotes in financials
func (m *SalesQuoteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *SalesQuoteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesQuote entity.
func (m *SalesQuoteItemRequestBuilder) Currency()(*i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca.CurrencyRequestBuilder) {
    return i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesQuote entity.
func (m *SalesQuoteItemRequestBuilder) Customer()(*i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5.CustomerRequestBuilder) {
    return i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesQuotes from financials
func (m *SalesQuoteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SalesQuoteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// MakeInvoice provides operations to call the makeInvoice method.
func (m *SalesQuoteItemRequestBuilder) MakeInvoice()(*i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b.MakeInvoiceRequestBuilder) {
    return i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b.NewMakeInvoiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesQuotes in financials
func (m *SalesQuoteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *SalesQuoteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesQuote entity.
func (m *SalesQuoteItemRequestBuilder) PaymentTerm()(*i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78.PaymentTermRequestBuilder) {
    return i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
func (m *SalesQuoteItemRequestBuilder) SalesQuoteLines()(*iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a.SalesQuoteLinesRequestBuilder) {
    return iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a.NewSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLinesById provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
func (m *SalesQuoteItemRequestBuilder) SalesQuoteLinesById(id string)(*id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce.SalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine%2Did"] = id
    }
    return id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce.NewSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Send provides operations to call the send method.
func (m *SalesQuoteItemRequestBuilder) Send()(*if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003.SendRequestBuilder) {
    return if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.salesQuote entity.
func (m *SalesQuoteItemRequestBuilder) ShipmentMethod()(*iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7.ShipmentMethodRequestBuilder) {
    return iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
