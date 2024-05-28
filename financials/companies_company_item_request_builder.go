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
// returns a *CompaniesItemAccountsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Accounts()(*CompaniesItemAccountsRequestBuilder) {
    return NewCompaniesItemAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsPayable provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
// returns a *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsPayable()(*CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) {
    return NewCompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsReceivable provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
// returns a *CompaniesItemAgedaccountsreceivableAgedAccountsReceivableRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsReceivable()(*CompaniesItemAgedaccountsreceivableAgedAccountsReceivableRequestBuilder) {
    return NewCompaniesItemAgedaccountsreceivableAgedAccountsReceivableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompanyInformation provides operations to manage the companyInformation property of the microsoft.graph.company entity.
// returns a *CompaniesItemCompanyinformationCompanyInformationRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) CompanyInformation()(*CompaniesItemCompanyinformationCompanyInformationRequestBuilder) {
    return NewCompaniesItemCompanyinformationCompanyInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesCompanyItemRequestBuilderInternal instantiates a new CompaniesCompanyItemRequestBuilder and sets the default values.
func NewCompaniesCompanyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesCompanyItemRequestBuilder) {
    m := &CompaniesCompanyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesCompanyItemRequestBuilder instantiates a new CompaniesCompanyItemRequestBuilder and sets the default values.
func NewCompaniesCompanyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesCompanyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesCompanyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CountriesRegions provides operations to manage the countriesRegions property of the microsoft.graph.company entity.
// returns a *CompaniesItemCountriesregionsCountriesRegionsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) CountriesRegions()(*CompaniesItemCountriesregionsCountriesRegionsRequestBuilder) {
    return NewCompaniesItemCountriesregionsCountriesRegionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Currencies provides operations to manage the currencies property of the microsoft.graph.company entity.
// returns a *CompaniesItemCurrenciesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Currencies()(*CompaniesItemCurrenciesRequestBuilder) {
    return NewCompaniesItemCurrenciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPaymentJournals provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
// returns a *CompaniesItemCustomerpaymentjournalsCustomerPaymentJournalsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) CustomerPaymentJournals()(*CompaniesItemCustomerpaymentjournalsCustomerPaymentJournalsRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsCustomerPaymentJournalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPayments provides operations to manage the customerPayments property of the microsoft.graph.company entity.
// returns a *CompaniesItemCustomerpaymentsCustomerPaymentsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) CustomerPayments()(*CompaniesItemCustomerpaymentsCustomerPaymentsRequestBuilder) {
    return NewCompaniesItemCustomerpaymentsCustomerPaymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customers provides operations to manage the customers property of the microsoft.graph.company entity.
// returns a *CompaniesItemCustomersRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Customers()(*CompaniesItemCustomersRequestBuilder) {
    return NewCompaniesItemCustomersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dimensions provides operations to manage the dimensions property of the microsoft.graph.company entity.
// returns a *CompaniesItemDimensionsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Dimensions()(*CompaniesItemDimensionsRequestBuilder) {
    return NewCompaniesItemDimensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DimensionValues provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
// returns a *CompaniesItemDimensionvaluesDimensionValuesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) DimensionValues()(*CompaniesItemDimensionvaluesDimensionValuesRequestBuilder) {
    return NewCompaniesItemDimensionvaluesDimensionValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Employees provides operations to manage the employees property of the microsoft.graph.company entity.
// returns a *CompaniesItemEmployeesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Employees()(*CompaniesItemEmployeesRequestBuilder) {
    return NewCompaniesItemEmployeesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GeneralLedgerEntries provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
// returns a *CompaniesItemGeneralledgerentriesGeneralLedgerEntriesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) GeneralLedgerEntries()(*CompaniesItemGeneralledgerentriesGeneralLedgerEntriesRequestBuilder) {
    return NewCompaniesItemGeneralledgerentriesGeneralLedgerEntriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get companies from financials
// returns a Companyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesCompanyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesCompanyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Companyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *CompaniesItemItemcategoriesItemCategoriesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) ItemCategories()(*CompaniesItemItemcategoriesItemCategoriesRequestBuilder) {
    return NewCompaniesItemItemcategoriesItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Items provides operations to manage the items property of the microsoft.graph.company entity.
// returns a *CompaniesItemItemsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Items()(*CompaniesItemItemsRequestBuilder) {
    return NewCompaniesItemItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// JournalLines provides operations to manage the journalLines property of the microsoft.graph.company entity.
// returns a *CompaniesItemJournallinesJournalLinesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) JournalLines()(*CompaniesItemJournallinesJournalLinesRequestBuilder) {
    return NewCompaniesItemJournallinesJournalLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Journals provides operations to manage the journals property of the microsoft.graph.company entity.
// returns a *CompaniesItemJournalsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Journals()(*CompaniesItemJournalsRequestBuilder) {
    return NewCompaniesItemJournalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentMethods provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
// returns a *CompaniesItemPaymentmethodsPaymentMethodsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) PaymentMethods()(*CompaniesItemPaymentmethodsPaymentMethodsRequestBuilder) {
    return NewCompaniesItemPaymentmethodsPaymentMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentTerms provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
// returns a *CompaniesItemPaymenttermsPaymentTermsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) PaymentTerms()(*CompaniesItemPaymenttermsPaymentTermsRequestBuilder) {
    return NewCompaniesItemPaymenttermsPaymentTermsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Picture provides operations to manage the picture property of the microsoft.graph.company entity.
// returns a *CompaniesItemPictureRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Picture()(*CompaniesItemPictureRequestBuilder) {
    return NewCompaniesItemPictureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoiceLines provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
// returns a *CompaniesItemPurchaseinvoicelinesPurchaseInvoiceLinesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoiceLines()(*CompaniesItemPurchaseinvoicelinesPurchaseInvoiceLinesRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicelinesPurchaseInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoices provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
// returns a *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoices()(*CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalescreditmemolinesSalesCreditMemoLinesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemoLines()(*CompaniesItemSalescreditmemolinesSalesCreditMemoLinesRequestBuilder) {
    return NewCompaniesItemSalescreditmemolinesSalesCreditMemoLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemos provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemos()(*CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) {
    return NewCompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoiceLines()(*CompaniesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) {
    return NewCompaniesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoices provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalesinvoicesSalesInvoicesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoices()(*CompaniesItemSalesinvoicesSalesInvoicesRequestBuilder) {
    return NewCompaniesItemSalesinvoicesSalesInvoicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrderLines provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalesorderlinesSalesOrderLinesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesOrderLines()(*CompaniesItemSalesorderlinesSalesOrderLinesRequestBuilder) {
    return NewCompaniesItemSalesorderlinesSalesOrderLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrders provides operations to manage the salesOrders property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalesordersSalesOrdersRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesOrders()(*CompaniesItemSalesordersSalesOrdersRequestBuilder) {
    return NewCompaniesItemSalesordersSalesOrdersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesQuoteLines()(*CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) {
    return NewCompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuotes provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalesquotesSalesQuotesRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) SalesQuotes()(*CompaniesItemSalesquotesSalesQuotesRequestBuilder) {
    return NewCompaniesItemSalesquotesSalesQuotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethods provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
// returns a *CompaniesItemShipmentmethodsShipmentMethodsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) ShipmentMethods()(*CompaniesItemShipmentmethodsShipmentMethodsRequestBuilder) {
    return NewCompaniesItemShipmentmethodsShipmentMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaxAreas provides operations to manage the taxAreas property of the microsoft.graph.company entity.
// returns a *CompaniesItemTaxareasTaxAreasRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) TaxAreas()(*CompaniesItemTaxareasTaxAreasRequestBuilder) {
    return NewCompaniesItemTaxareasTaxAreasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaxGroups provides operations to manage the taxGroups property of the microsoft.graph.company entity.
// returns a *CompaniesItemTaxgroupsTaxGroupsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) TaxGroups()(*CompaniesItemTaxgroupsTaxGroupsRequestBuilder) {
    return NewCompaniesItemTaxgroupsTaxGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get companies from financials
// returns a *RequestInformation when successful
func (m *CompaniesCompanyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesCompanyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnitsOfMeasure provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
// returns a *CompaniesItemUnitsofmeasureUnitsOfMeasureRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) UnitsOfMeasure()(*CompaniesItemUnitsofmeasureUnitsOfMeasureRequestBuilder) {
    return NewCompaniesItemUnitsofmeasureUnitsOfMeasureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Vendors provides operations to manage the vendors property of the microsoft.graph.company entity.
// returns a *CompaniesItemVendorsRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) Vendors()(*CompaniesItemVendorsRequestBuilder) {
    return NewCompaniesItemVendorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesCompanyItemRequestBuilder when successful
func (m *CompaniesCompanyItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesCompanyItemRequestBuilder) {
    return NewCompaniesCompanyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
