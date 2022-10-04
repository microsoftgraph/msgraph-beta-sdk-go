package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/shipmentmethod"
    i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/currency"
    i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/salesinvoicelines"
    i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/cancel"
    i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/post"
    i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/cancelandsend"
    i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/send"
    i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/customer"
    i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/paymentterm"
    iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/postandsend"
    if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/salesinvoicelines/item"
)

// SalesInvoiceItemRequestBuilder provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
type SalesInvoiceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SalesInvoiceItemRequestBuilderGetQueryParameters get salesInvoices from financials
type SalesInvoiceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SalesInvoiceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesInvoiceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SalesInvoiceItemRequestBuilderGetQueryParameters
}
// SalesInvoiceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesInvoiceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel the cancel property
func (m *SalesInvoiceItemRequestBuilder) Cancel()(*i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2.CancelRequestBuilder) {
    return i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CancelAndSend the cancelAndSend property
func (m *SalesInvoiceItemRequestBuilder) CancelAndSend()(*i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee.CancelAndSendRequestBuilder) {
    return i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee.NewCancelAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSalesInvoiceItemRequestBuilderInternal instantiates a new SalesInvoiceItemRequestBuilder and sets the default values.
func NewSalesInvoiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesInvoiceItemRequestBuilder) {
    m := &SalesInvoiceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesInvoiceItemRequestBuilder instantiates a new SalesInvoiceItemRequestBuilder and sets the default values.
func NewSalesInvoiceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesInvoiceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesInvoiceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesInvoices from financials
func (m *SalesInvoiceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SalesInvoiceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesInvoices in financials
func (m *SalesInvoiceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *SalesInvoiceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency the currency property
func (m *SalesInvoiceItemRequestBuilder) Currency()(*i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47.CurrencyRequestBuilder) {
    return i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer the customer property
func (m *SalesInvoiceItemRequestBuilder) Customer()(*i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889.CustomerRequestBuilder) {
    return i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesInvoices from financials
func (m *SalesInvoiceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SalesInvoiceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable), nil
}
// Patch update the navigation property salesInvoices in financials
func (m *SalesInvoiceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *SalesInvoiceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable), nil
}
// PaymentTerm the paymentTerm property
func (m *SalesInvoiceItemRequestBuilder) PaymentTerm()(*i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb.PaymentTermRequestBuilder) {
    return i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the post property
func (m *SalesInvoiceItemRequestBuilder) Post()(*i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b.PostRequestBuilder) {
    return i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b.NewPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PostAndSend the postAndSend property
func (m *SalesInvoiceItemRequestBuilder) PostAndSend()(*iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609.PostAndSendRequestBuilder) {
    return iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609.NewPostAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLines the salesInvoiceLines property
func (m *SalesInvoiceItemRequestBuilder) SalesInvoiceLines()(*i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7.SalesInvoiceLinesRequestBuilder) {
    return i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7.NewSalesInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesInvoices.item.salesInvoiceLines.item collection
func (m *SalesInvoiceItemRequestBuilder) SalesInvoiceLinesById(id string)(*if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae.SalesInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine%2Did"] = id
    }
    return if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae.NewSalesInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Send the send property
func (m *SalesInvoiceItemRequestBuilder) Send()(*i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3.SendRequestBuilder) {
    return i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethod the shipmentMethod property
func (m *SalesInvoiceItemRequestBuilder) ShipmentMethod()(*i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff.ShipmentMethodRequestBuilder) {
    return i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
