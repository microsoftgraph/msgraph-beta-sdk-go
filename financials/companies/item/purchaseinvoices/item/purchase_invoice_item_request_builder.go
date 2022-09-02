package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3d86c6bd7b31ece7f9e0d0acb1882ec4a41b7e93cf9d45845efb984eab9dfe90 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/post"
    i48c88f94a8a21b35ebe6d7459f6ba6cd96a3d5e75749e3dd372526fbcb96f6d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped"
    i66ab53c3b2ddae75121633bbd823599d4e4e48fe876a2f914d5a57f2e241b272 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/currency"
    ic84ce80e57a95f249407c29d4b5fec4a8c0714292046ec3ba7c175bab69f542a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/purchaseinvoicelines"
    i25bfa50bef1bcf1c425633cb974de803465fa55f263668c4ae6bb8fa47aaf7cb "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/purchaseinvoicelines/item"
)

// PurchaseInvoiceItemRequestBuilder provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
type PurchaseInvoiceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PurchaseInvoiceItemRequestBuilderGetQueryParameters get purchaseInvoices from financials
type PurchaseInvoiceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PurchaseInvoiceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PurchaseInvoiceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PurchaseInvoiceItemRequestBuilderGetQueryParameters
}
// PurchaseInvoiceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PurchaseInvoiceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPurchaseInvoiceItemRequestBuilderInternal instantiates a new PurchaseInvoiceItemRequestBuilder and sets the default values.
func NewPurchaseInvoiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PurchaseInvoiceItemRequestBuilder) {
    m := &PurchaseInvoiceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices/{purchaseInvoice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPurchaseInvoiceItemRequestBuilder instantiates a new PurchaseInvoiceItemRequestBuilder and sets the default values.
func NewPurchaseInvoiceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PurchaseInvoiceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPurchaseInvoiceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get purchaseInvoices from financials
func (m *PurchaseInvoiceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get purchaseInvoices from financials
func (m *PurchaseInvoiceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PurchaseInvoiceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property purchaseInvoices in financials
func (m *PurchaseInvoiceItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property purchaseInvoices in financials
func (m *PurchaseInvoiceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, requestConfiguration *PurchaseInvoiceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Currency the currency property
func (m *PurchaseInvoiceItemRequestBuilder) Currency()(*i66ab53c3b2ddae75121633bbd823599d4e4e48fe876a2f914d5a57f2e241b272.CurrencyRequestBuilder) {
    return i66ab53c3b2ddae75121633bbd823599d4e4e48fe876a2f914d5a57f2e241b272.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get purchaseInvoices from financials
func (m *PurchaseInvoiceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PurchaseInvoiceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable), nil
}
// Patch update the navigation property purchaseInvoices in financials
func (m *PurchaseInvoiceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, requestConfiguration *PurchaseInvoiceItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Post the post property
func (m *PurchaseInvoiceItemRequestBuilder) Post()(*i3d86c6bd7b31ece7f9e0d0acb1882ec4a41b7e93cf9d45845efb984eab9dfe90.PostRequestBuilder) {
    return i3d86c6bd7b31ece7f9e0d0acb1882ec4a41b7e93cf9d45845efb984eab9dfe90.NewPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLines the purchaseInvoiceLines property
func (m *PurchaseInvoiceItemRequestBuilder) PurchaseInvoiceLines()(*ic84ce80e57a95f249407c29d4b5fec4a8c0714292046ec3ba7c175bab69f542a.PurchaseInvoiceLinesRequestBuilder) {
    return ic84ce80e57a95f249407c29d4b5fec4a8c0714292046ec3ba7c175bab69f542a.NewPurchaseInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.purchaseInvoices.item.purchaseInvoiceLines.item collection
func (m *PurchaseInvoiceItemRequestBuilder) PurchaseInvoiceLinesById(id string)(*i25bfa50bef1bcf1c425633cb974de803465fa55f263668c4ae6bb8fa47aaf7cb.PurchaseInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoiceLine%2Did"] = id
    }
    return i25bfa50bef1bcf1c425633cb974de803465fa55f263668c4ae6bb8fa47aaf7cb.NewPurchaseInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Vendor_escaped the vendor property
func (m *PurchaseInvoiceItemRequestBuilder) Vendor_escaped()(*i48c88f94a8a21b35ebe6d7459f6ba6cd96a3d5e75749e3dd372526fbcb96f6d3.VendorRequestBuilder) {
    return i48c88f94a8a21b35ebe6d7459f6ba6cd96a3d5e75749e3dd372526fbcb96f6d3.NewVendorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
