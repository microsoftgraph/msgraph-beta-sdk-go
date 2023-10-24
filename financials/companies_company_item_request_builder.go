package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesCompanyItemRequestBuilder provides operations to manage the companies property of the microsoft.graph.financials entity.
type CompaniesCompanyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesCompanyItemRequestBuilderGetQueryParameters get companies from financials
type CompaniesCompanyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesCompanyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesCompanyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesCompanyItemRequestBuilderGetQueryParameters
}
// Accounts provides operations to manage the accounts property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Accounts()(*CompaniesItemAccountsRequestBuilder) {
    return NewCompaniesItemAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsPayable provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsPayable()(*CompaniesItemAgedAccountsPayableRequestBuilder) {
    return NewCompaniesItemAgedAccountsPayableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsReceivable provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsReceivable()(*CompaniesItemAgedAccountsReceivableRequestBuilder) {
    return NewCompaniesItemAgedAccountsReceivableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompanyInformation provides operations to manage the companyInformation property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CompanyInformation()(*CompaniesItemCompanyInformationRequestBuilder) {
    return NewCompaniesItemCompanyInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesCompanyItemRequestBuilderInternal instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewCompaniesCompanyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesCompanyItemRequestBuilder) {
    m := &CompaniesCompanyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesCompanyItemRequestBuilder instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewCompaniesCompanyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesCompanyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesCompanyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CountriesRegions provides operations to manage the countriesRegions property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CountriesRegions()(*CompaniesItemCountriesRegionsRequestBuilder) {
    return NewCompaniesItemCountriesRegionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Currencies provides operations to manage the currencies property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Currencies()(*CompaniesItemCurrenciesRequestBuilder) {
    return NewCompaniesItemCurrenciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPaymentJournals provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CustomerPaymentJournals()(*CompaniesItemCustomerPaymentJournalsRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPayments provides operations to manage the customerPayments property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CustomerPayments()(*CompaniesItemCustomerPaymentsRequestBuilder) {
    return NewCompaniesItemCustomerPaymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customers provides operations to manage the customers property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Customers()(*CompaniesItemCustomersRequestBuilder) {
    return NewCompaniesItemCustomersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dimensions provides operations to manage the dimensions property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Dimensions()(*CompaniesItemDimensionsRequestBuilder) {
    return NewCompaniesItemDimensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DimensionValues provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) DimensionValues()(*CompaniesItemDimensionValuesRequestBuilder) {
    return NewCompaniesItemDimensionValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Employees provides operations to manage the employees property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Employees()(*CompaniesItemEmployeesRequestBuilder) {
    return NewCompaniesItemEmployeesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GeneralLedgerEntries provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) GeneralLedgerEntries()(*CompaniesItemGeneralLedgerEntriesRequestBuilder) {
    return NewCompaniesItemGeneralLedgerEntriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get companies from financials
func (m *CompaniesCompanyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesCompanyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Companyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCompanyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Companyable), nil
}
// ItemCategories provides operations to manage the itemCategories property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) ItemCategories()(*CompaniesItemItemCategoriesRequestBuilder) {
    return NewCompaniesItemItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Items provides operations to manage the items property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Items()(*CompaniesItemItemsRequestBuilder) {
    return NewCompaniesItemItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// JournalLines provides operations to manage the journalLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) JournalLines()(*CompaniesItemJournalLinesRequestBuilder) {
    return NewCompaniesItemJournalLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Journals provides operations to manage the journals property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Journals()(*CompaniesItemJournalsRequestBuilder) {
    return NewCompaniesItemJournalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentMethods provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PaymentMethods()(*CompaniesItemPaymentMethodsRequestBuilder) {
    return NewCompaniesItemPaymentMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentTerms provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PaymentTerms()(*CompaniesItemPaymentTermsRequestBuilder) {
    return NewCompaniesItemPaymentTermsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Picture provides operations to manage the picture property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Picture()(*CompaniesItemPictureRequestBuilder) {
    return NewCompaniesItemPictureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoiceLines provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoiceLines()(*CompaniesItemPurchaseInvoiceLinesRequestBuilder) {
    return NewCompaniesItemPurchaseInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoices provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoices()(*CompaniesItemPurchaseInvoicesRequestBuilder) {
    return NewCompaniesItemPurchaseInvoicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemoLines()(*CompaniesItemSalesCreditMemoLinesRequestBuilder) {
    return NewCompaniesItemSalesCreditMemoLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemos provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemos()(*CompaniesItemSalesCreditMemosRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoiceLines()(*CompaniesItemSalesInvoiceLinesRequestBuilder) {
    return NewCompaniesItemSalesInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoices provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoices()(*CompaniesItemSalesInvoicesRequestBuilder) {
    return NewCompaniesItemSalesInvoicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrderLines provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesOrderLines()(*CompaniesItemSalesOrderLinesRequestBuilder) {
    return NewCompaniesItemSalesOrderLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrders provides operations to manage the salesOrders property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesOrders()(*CompaniesItemSalesOrdersRequestBuilder) {
    return NewCompaniesItemSalesOrdersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesQuoteLines()(*CompaniesItemSalesQuoteLinesRequestBuilder) {
    return NewCompaniesItemSalesQuoteLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuotes provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesQuotes()(*CompaniesItemSalesQuotesRequestBuilder) {
    return NewCompaniesItemSalesQuotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethods provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) ShipmentMethods()(*CompaniesItemShipmentMethodsRequestBuilder) {
    return NewCompaniesItemShipmentMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaxAreas provides operations to manage the taxAreas property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) TaxAreas()(*CompaniesItemTaxAreasRequestBuilder) {
    return NewCompaniesItemTaxAreasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaxGroups provides operations to manage the taxGroups property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) TaxGroups()(*CompaniesItemTaxGroupsRequestBuilder) {
    return NewCompaniesItemTaxGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get companies from financials
func (m *CompaniesCompanyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesCompanyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// UnitsOfMeasure provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) UnitsOfMeasure()(*CompaniesItemUnitsOfMeasureRequestBuilder) {
    return NewCompaniesItemUnitsOfMeasureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Vendors provides operations to manage the vendors property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Vendors()(*CompaniesItemVendorsRequestBuilder) {
    return NewCompaniesItemVendorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CompaniesCompanyItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesCompanyItemRequestBuilder) {
    return NewCompaniesCompanyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
