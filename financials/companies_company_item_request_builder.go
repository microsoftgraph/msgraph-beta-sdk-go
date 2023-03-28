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
// AccountsById provides operations to manage the accounts property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) AccountsById(id string)(*CompaniesItemAccountsAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["account%2Did"] = id
    }
    return NewCompaniesItemAccountsAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsPayable provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsPayable()(*CompaniesItemAgedAccountsPayableRequestBuilder) {
    return NewCompaniesItemAgedAccountsPayableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsPayableById provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsPayableById(id string)(*CompaniesItemAgedAccountsPayableAgedAccountsPayableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsPayable%2Did"] = id
    }
    return NewCompaniesItemAgedAccountsPayableAgedAccountsPayableItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsReceivable provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsReceivable()(*CompaniesItemAgedAccountsReceivableRequestBuilder) {
    return NewCompaniesItemAgedAccountsReceivableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgedAccountsReceivableById provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) AgedAccountsReceivableById(id string)(*CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsReceivable%2Did"] = id
    }
    return NewCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// CompanyInformation provides operations to manage the companyInformation property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CompanyInformation()(*CompaniesItemCompanyInformationRequestBuilder) {
    return NewCompaniesItemCompanyInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompanyInformationById provides operations to manage the companyInformation property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CompanyInformationById(id string)(*CompaniesItemCompanyInformationCompanyInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["companyInformation%2Did"] = id
    }
    return NewCompaniesItemCompanyInformationCompanyInformationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
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
// CountriesRegionsById provides operations to manage the countriesRegions property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CountriesRegionsById(id string)(*CompaniesItemCountriesRegionsCountryRegionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["countryRegion%2Did"] = id
    }
    return NewCompaniesItemCountriesRegionsCountryRegionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Currencies provides operations to manage the currencies property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Currencies()(*CompaniesItemCurrenciesRequestBuilder) {
    return NewCompaniesItemCurrenciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CurrenciesById provides operations to manage the currencies property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CurrenciesById(id string)(*CompaniesItemCurrenciesCurrencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["currency%2Did"] = id
    }
    return NewCompaniesItemCurrenciesCurrencyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPaymentJournals provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CustomerPaymentJournals()(*CompaniesItemCustomerPaymentJournalsRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPaymentJournalsById provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CustomerPaymentJournalsById(id string)(*CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPaymentJournal%2Did"] = id
    }
    return NewCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPayments provides operations to manage the customerPayments property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CustomerPayments()(*CompaniesItemCustomerPaymentsRequestBuilder) {
    return NewCompaniesItemCustomerPaymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomerPaymentsById provides operations to manage the customerPayments property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CustomerPaymentsById(id string)(*CompaniesItemCustomerPaymentsCustomerPaymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPayment%2Did"] = id
    }
    return NewCompaniesItemCustomerPaymentsCustomerPaymentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Customers provides operations to manage the customers property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Customers()(*CompaniesItemCustomersRequestBuilder) {
    return NewCompaniesItemCustomersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomersById provides operations to manage the customers property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) CustomersById(id string)(*CompaniesItemCustomersCustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customer%2Did"] = id
    }
    return NewCompaniesItemCustomersCustomerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Dimensions provides operations to manage the dimensions property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Dimensions()(*CompaniesItemDimensionsRequestBuilder) {
    return NewCompaniesItemDimensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DimensionsById provides operations to manage the dimensions property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) DimensionsById(id string)(*CompaniesItemDimensionsDimensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimension%2Did"] = id
    }
    return NewCompaniesItemDimensionsDimensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// DimensionValues provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) DimensionValues()(*CompaniesItemDimensionValuesRequestBuilder) {
    return NewCompaniesItemDimensionValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DimensionValuesById provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) DimensionValuesById(id string)(*CompaniesItemDimensionValuesDimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimensionValue%2Did"] = id
    }
    return NewCompaniesItemDimensionValuesDimensionValueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Employees provides operations to manage the employees property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Employees()(*CompaniesItemEmployeesRequestBuilder) {
    return NewCompaniesItemEmployeesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmployeesById provides operations to manage the employees property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) EmployeesById(id string)(*CompaniesItemEmployeesEmployeeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["employee%2Did"] = id
    }
    return NewCompaniesItemEmployeesEmployeeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// GeneralLedgerEntries provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) GeneralLedgerEntries()(*CompaniesItemGeneralLedgerEntriesRequestBuilder) {
    return NewCompaniesItemGeneralLedgerEntriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GeneralLedgerEntriesById provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) GeneralLedgerEntriesById(id string)(*CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["generalLedgerEntry%2Did"] = id
    }
    return NewCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
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
// ItemCategoriesById provides operations to manage the itemCategories property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) ItemCategoriesById(id string)(*CompaniesItemItemCategoriesItemCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemCategory%2Did"] = id
    }
    return NewCompaniesItemItemCategoriesItemCategoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Items provides operations to manage the items property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Items()(*CompaniesItemItemsRequestBuilder) {
    return NewCompaniesItemItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ItemsById provides operations to manage the items property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) ItemsById(id string)(*CompaniesItemItemsItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["item%2Did"] = id
    }
    return NewCompaniesItemItemsItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// JournalLines provides operations to manage the journalLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) JournalLines()(*CompaniesItemJournalLinesRequestBuilder) {
    return NewCompaniesItemJournalLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// JournalLinesById provides operations to manage the journalLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) JournalLinesById(id string)(*CompaniesItemJournalLinesJournalLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journalLine%2Did"] = id
    }
    return NewCompaniesItemJournalLinesJournalLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Journals provides operations to manage the journals property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Journals()(*CompaniesItemJournalsRequestBuilder) {
    return NewCompaniesItemJournalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// JournalsById provides operations to manage the journals property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) JournalsById(id string)(*CompaniesItemJournalsJournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journal%2Did"] = id
    }
    return NewCompaniesItemJournalsJournalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentMethods provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PaymentMethods()(*CompaniesItemPaymentMethodsRequestBuilder) {
    return NewCompaniesItemPaymentMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentMethodsById provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PaymentMethodsById(id string)(*CompaniesItemPaymentMethodsPaymentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentMethod%2Did"] = id
    }
    return NewCompaniesItemPaymentMethodsPaymentMethodItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentTerms provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PaymentTerms()(*CompaniesItemPaymentTermsRequestBuilder) {
    return NewCompaniesItemPaymentTermsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentTermsById provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PaymentTermsById(id string)(*CompaniesItemPaymentTermsPaymentTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentTerm%2Did"] = id
    }
    return NewCompaniesItemPaymentTermsPaymentTermItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Picture provides operations to manage the picture property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Picture()(*CompaniesItemPictureRequestBuilder) {
    return NewCompaniesItemPictureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PictureById provides operations to manage the picture property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PictureById(id string)(*CompaniesItemPicturePictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return NewCompaniesItemPicturePictureItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoiceLines provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoiceLines()(*CompaniesItemPurchaseInvoiceLinesRequestBuilder) {
    return NewCompaniesItemPurchaseInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoiceLinesById provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoiceLinesById(id string)(*CompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoiceLine%2Did"] = id
    }
    return NewCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoices provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoices()(*CompaniesItemPurchaseInvoicesRequestBuilder) {
    return NewCompaniesItemPurchaseInvoicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoicesById provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) PurchaseInvoicesById(id string)(*CompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoice%2Did"] = id
    }
    return NewCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemoLines()(*CompaniesItemSalesCreditMemoLinesRequestBuilder) {
    return NewCompaniesItemSalesCreditMemoLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemoLinesById provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemoLinesById(id string)(*CompaniesItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemoLine%2Did"] = id
    }
    return NewCompaniesItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemos provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemos()(*CompaniesItemSalesCreditMemosRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemosById provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesCreditMemosById(id string)(*CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemo%2Did"] = id
    }
    return NewCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoiceLines()(*CompaniesItemSalesInvoiceLinesRequestBuilder) {
    return NewCompaniesItemSalesInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoiceLinesById provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoiceLinesById(id string)(*CompaniesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine%2Did"] = id
    }
    return NewCompaniesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoices provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoices()(*CompaniesItemSalesInvoicesRequestBuilder) {
    return NewCompaniesItemSalesInvoicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoicesById provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesInvoicesById(id string)(*CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoice%2Did"] = id
    }
    return NewCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrderLines provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesOrderLines()(*CompaniesItemSalesOrderLinesRequestBuilder) {
    return NewCompaniesItemSalesOrderLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrderLinesById provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesOrderLinesById(id string)(*CompaniesItemSalesOrderLinesSalesOrderLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrderLine%2Did"] = id
    }
    return NewCompaniesItemSalesOrderLinesSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrders provides operations to manage the salesOrders property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesOrders()(*CompaniesItemSalesOrdersRequestBuilder) {
    return NewCompaniesItemSalesOrdersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesOrdersById provides operations to manage the salesOrders property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesOrdersById(id string)(*CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrder%2Did"] = id
    }
    return NewCompaniesItemSalesOrdersSalesOrderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesQuoteLines()(*CompaniesItemSalesQuoteLinesRequestBuilder) {
    return NewCompaniesItemSalesQuoteLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuoteLinesById provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesQuoteLinesById(id string)(*CompaniesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine%2Did"] = id
    }
    return NewCompaniesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuotes provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesQuotes()(*CompaniesItemSalesQuotesRequestBuilder) {
    return NewCompaniesItemSalesQuotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuotesById provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) SalesQuotesById(id string)(*CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuote%2Did"] = id
    }
    return NewCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethods provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) ShipmentMethods()(*CompaniesItemShipmentMethodsRequestBuilder) {
    return NewCompaniesItemShipmentMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethodsById provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) ShipmentMethodsById(id string)(*CompaniesItemShipmentMethodsShipmentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shipmentMethod%2Did"] = id
    }
    return NewCompaniesItemShipmentMethodsShipmentMethodItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// TaxAreas provides operations to manage the taxAreas property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) TaxAreas()(*CompaniesItemTaxAreasRequestBuilder) {
    return NewCompaniesItemTaxAreasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaxAreasById provides operations to manage the taxAreas property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) TaxAreasById(id string)(*CompaniesItemTaxAreasTaxAreaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxArea%2Did"] = id
    }
    return NewCompaniesItemTaxAreasTaxAreaItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// TaxGroups provides operations to manage the taxGroups property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) TaxGroups()(*CompaniesItemTaxGroupsRequestBuilder) {
    return NewCompaniesItemTaxGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaxGroupsById provides operations to manage the taxGroups property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) TaxGroupsById(id string)(*CompaniesItemTaxGroupsTaxGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxGroup%2Did"] = id
    }
    return NewCompaniesItemTaxGroupsTaxGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get companies from financials
func (m *CompaniesCompanyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesCompanyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UnitsOfMeasure provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) UnitsOfMeasure()(*CompaniesItemUnitsOfMeasureRequestBuilder) {
    return NewCompaniesItemUnitsOfMeasureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UnitsOfMeasureById provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) UnitsOfMeasureById(id string)(*CompaniesItemUnitsOfMeasureUnitOfMeasureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unitOfMeasure%2Did"] = id
    }
    return NewCompaniesItemUnitsOfMeasureUnitOfMeasureItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Vendors provides operations to manage the vendors property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) Vendors()(*CompaniesItemVendorsRequestBuilder) {
    return NewCompaniesItemVendorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VendorsById provides operations to manage the vendors property of the microsoft.graph.company entity.
func (m *CompaniesCompanyItemRequestBuilder) VendorsById(id string)(*CompaniesItemVendorsVendorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vendor%2Did"] = id
    }
    return NewCompaniesItemVendorsVendorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
