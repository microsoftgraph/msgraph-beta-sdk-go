package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesCompanyItemRequestBuilder provides operations to manage the companies property of the microsoft.graph.financials entity.
type FinancialsCompaniesCompanyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesCompanyItemRequestBuilderGetQueryParameters get companies from financials
type FinancialsCompaniesCompanyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesCompanyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesCompanyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesCompanyItemRequestBuilderGetQueryParameters
}
// Accounts provides operations to manage the accounts property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Accounts()(*FinancialsCompaniesItemAccountsRequestBuilder) {
    return NewFinancialsCompaniesItemAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountsById provides operations to manage the accounts property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) AccountsById(id string)(*FinancialsCompaniesItemAccountsAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["account%2Did"] = id
    }
    return NewFinancialsCompaniesItemAccountsAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgedAccountsPayable provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) AgedAccountsPayable()(*FinancialsCompaniesItemAgedAccountsPayableRequestBuilder) {
    return NewFinancialsCompaniesItemAgedAccountsPayableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgedAccountsPayableById provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) AgedAccountsPayableById(id string)(*FinancialsCompaniesItemAgedAccountsPayableAgedAccountsPayableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsPayable%2Did"] = id
    }
    return NewFinancialsCompaniesItemAgedAccountsPayableAgedAccountsPayableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgedAccountsReceivable provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) AgedAccountsReceivable()(*FinancialsCompaniesItemAgedAccountsReceivableRequestBuilder) {
    return NewFinancialsCompaniesItemAgedAccountsReceivableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgedAccountsReceivableById provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) AgedAccountsReceivableById(id string)(*FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsReceivable%2Did"] = id
    }
    return NewFinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompanyInformation provides operations to manage the companyInformation property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CompanyInformation()(*FinancialsCompaniesItemCompanyInformationRequestBuilder) {
    return NewFinancialsCompaniesItemCompanyInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompanyInformationById provides operations to manage the companyInformation property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CompanyInformationById(id string)(*FinancialsCompaniesItemCompanyInformationCompanyInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["companyInformation%2Did"] = id
    }
    return NewFinancialsCompaniesItemCompanyInformationCompanyInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewFinancialsCompaniesCompanyItemRequestBuilderInternal instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesCompanyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesCompanyItemRequestBuilder) {
    m := &FinancialsCompaniesCompanyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesCompanyItemRequestBuilder instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesCompanyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesCompanyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesCompanyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CountriesRegions provides operations to manage the countriesRegions property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CountriesRegions()(*FinancialsCompaniesItemCountriesRegionsRequestBuilder) {
    return NewFinancialsCompaniesItemCountriesRegionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CountriesRegionsById provides operations to manage the countriesRegions property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CountriesRegionsById(id string)(*FinancialsCompaniesItemCountriesRegionsCountryRegionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["countryRegion%2Did"] = id
    }
    return NewFinancialsCompaniesItemCountriesRegionsCountryRegionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation get companies from financials
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesCompanyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currencies provides operations to manage the currencies property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Currencies()(*FinancialsCompaniesItemCurrenciesRequestBuilder) {
    return NewFinancialsCompaniesItemCurrenciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CurrenciesById provides operations to manage the currencies property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CurrenciesById(id string)(*FinancialsCompaniesItemCurrenciesCurrencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["currency%2Did"] = id
    }
    return NewFinancialsCompaniesItemCurrenciesCurrencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomerPaymentJournals provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CustomerPaymentJournals()(*FinancialsCompaniesItemCustomerPaymentJournalsRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentJournalsById provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CustomerPaymentJournalsById(id string)(*FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPaymentJournal%2Did"] = id
    }
    return NewFinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomerPayments provides operations to manage the customerPayments property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CustomerPayments()(*FinancialsCompaniesItemCustomerPaymentsRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentsById provides operations to manage the customerPayments property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CustomerPaymentsById(id string)(*FinancialsCompaniesItemCustomerPaymentsCustomerPaymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPayment%2Did"] = id
    }
    return NewFinancialsCompaniesItemCustomerPaymentsCustomerPaymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Customers provides operations to manage the customers property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Customers()(*FinancialsCompaniesItemCustomersRequestBuilder) {
    return NewFinancialsCompaniesItemCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById provides operations to manage the customers property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) CustomersById(id string)(*FinancialsCompaniesItemCustomersCustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customer%2Did"] = id
    }
    return NewFinancialsCompaniesItemCustomersCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Dimensions provides operations to manage the dimensions property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Dimensions()(*FinancialsCompaniesItemDimensionsRequestBuilder) {
    return NewFinancialsCompaniesItemDimensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionsById provides operations to manage the dimensions property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) DimensionsById(id string)(*FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimension%2Did"] = id
    }
    return NewFinancialsCompaniesItemDimensionsDimensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DimensionValues provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) DimensionValues()(*FinancialsCompaniesItemDimensionValuesRequestBuilder) {
    return NewFinancialsCompaniesItemDimensionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionValuesById provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) DimensionValuesById(id string)(*FinancialsCompaniesItemDimensionValuesDimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimensionValue%2Did"] = id
    }
    return NewFinancialsCompaniesItemDimensionValuesDimensionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Employees provides operations to manage the employees property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Employees()(*FinancialsCompaniesItemEmployeesRequestBuilder) {
    return NewFinancialsCompaniesItemEmployeesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmployeesById provides operations to manage the employees property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) EmployeesById(id string)(*FinancialsCompaniesItemEmployeesEmployeeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["employee%2Did"] = id
    }
    return NewFinancialsCompaniesItemEmployeesEmployeeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GeneralLedgerEntries provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) GeneralLedgerEntries()(*FinancialsCompaniesItemGeneralLedgerEntriesRequestBuilder) {
    return NewFinancialsCompaniesItemGeneralLedgerEntriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GeneralLedgerEntriesById provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) GeneralLedgerEntriesById(id string)(*FinancialsCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["generalLedgerEntry%2Did"] = id
    }
    return NewFinancialsCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get companies from financials
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesCompanyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Companyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCompanyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Companyable), nil
}
// ItemCategories provides operations to manage the itemCategories property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) ItemCategories()(*FinancialsCompaniesItemItemCategoriesRequestBuilder) {
    return NewFinancialsCompaniesItemItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemCategoriesById provides operations to manage the itemCategories property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) ItemCategoriesById(id string)(*FinancialsCompaniesItemItemCategoriesItemCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemCategory%2Did"] = id
    }
    return NewFinancialsCompaniesItemItemCategoriesItemCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Items provides operations to manage the items property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Items()(*FinancialsCompaniesItemItemsRequestBuilder) {
    return NewFinancialsCompaniesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) ItemsById(id string)(*FinancialsCompaniesItemItemsItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["item%2Did"] = id
    }
    return NewFinancialsCompaniesItemItemsItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// JournalLines provides operations to manage the journalLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) JournalLines()(*FinancialsCompaniesItemJournalLinesRequestBuilder) {
    return NewFinancialsCompaniesItemJournalLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JournalLinesById provides operations to manage the journalLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) JournalLinesById(id string)(*FinancialsCompaniesItemJournalLinesJournalLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journalLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemJournalLinesJournalLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Journals provides operations to manage the journals property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Journals()(*FinancialsCompaniesItemJournalsRequestBuilder) {
    return NewFinancialsCompaniesItemJournalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JournalsById provides operations to manage the journals property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) JournalsById(id string)(*FinancialsCompaniesItemJournalsJournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journal%2Did"] = id
    }
    return NewFinancialsCompaniesItemJournalsJournalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PaymentMethods provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PaymentMethods()(*FinancialsCompaniesItemPaymentMethodsRequestBuilder) {
    return NewFinancialsCompaniesItemPaymentMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentMethodsById provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PaymentMethodsById(id string)(*FinancialsCompaniesItemPaymentMethodsPaymentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentMethod%2Did"] = id
    }
    return NewFinancialsCompaniesItemPaymentMethodsPaymentMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PaymentTerms provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PaymentTerms()(*FinancialsCompaniesItemPaymentTermsRequestBuilder) {
    return NewFinancialsCompaniesItemPaymentTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTermsById provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PaymentTermsById(id string)(*FinancialsCompaniesItemPaymentTermsPaymentTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentTerm%2Did"] = id
    }
    return NewFinancialsCompaniesItemPaymentTermsPaymentTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Picture provides operations to manage the picture property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Picture()(*FinancialsCompaniesItemPictureRequestBuilder) {
    return NewFinancialsCompaniesItemPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById provides operations to manage the picture property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PictureById(id string)(*FinancialsCompaniesItemPicturePictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return NewFinancialsCompaniesItemPicturePictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PurchaseInvoiceLines provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PurchaseInvoiceLines()(*FinancialsCompaniesItemPurchaseInvoiceLinesRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLinesById provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PurchaseInvoiceLinesById(id string)(*FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoiceLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PurchaseInvoices provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PurchaseInvoices()(*FinancialsCompaniesItemPurchaseInvoicesRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoicesById provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) PurchaseInvoicesById(id string)(*FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoice%2Did"] = id
    }
    return NewFinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesCreditMemoLines()(*FinancialsCompaniesItemSalesCreditMemoLinesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemoLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLinesById provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesCreditMemoLinesById(id string)(*FinancialsCompaniesItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemoLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesCreditMemos provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesCreditMemos()(*FinancialsCompaniesItemSalesCreditMemosRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemosById provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesCreditMemosById(id string)(*FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemo%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesInvoiceLines()(*FinancialsCompaniesItemSalesInvoiceLinesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLinesById provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesInvoiceLinesById(id string)(*FinancialsCompaniesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesInvoices provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesInvoices()(*FinancialsCompaniesItemSalesInvoicesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoicesById provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesInvoicesById(id string)(*FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoice%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesOrderLines provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesOrderLines()(*FinancialsCompaniesItemSalesOrderLinesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesOrderLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLinesById provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesOrderLinesById(id string)(*FinancialsCompaniesItemSalesOrderLinesSalesOrderLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrderLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesOrderLinesSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesOrders provides operations to manage the salesOrders property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesOrders()(*FinancialsCompaniesItemSalesOrdersRequestBuilder) {
    return NewFinancialsCompaniesItemSalesOrdersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrdersById provides operations to manage the salesOrders property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesOrdersById(id string)(*FinancialsCompaniesItemSalesOrdersSalesOrderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrder%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesOrdersSalesOrderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesQuoteLines()(*FinancialsCompaniesItemSalesQuoteLinesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLinesById provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesQuoteLinesById(id string)(*FinancialsCompaniesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesQuotes provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesQuotes()(*FinancialsCompaniesItemSalesQuotesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuotesById provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) SalesQuotesById(id string)(*FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuote%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ShipmentMethods provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) ShipmentMethods()(*FinancialsCompaniesItemShipmentMethodsRequestBuilder) {
    return NewFinancialsCompaniesItemShipmentMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethodsById provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) ShipmentMethodsById(id string)(*FinancialsCompaniesItemShipmentMethodsShipmentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shipmentMethod%2Did"] = id
    }
    return NewFinancialsCompaniesItemShipmentMethodsShipmentMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaxAreas provides operations to manage the taxAreas property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) TaxAreas()(*FinancialsCompaniesItemTaxAreasRequestBuilder) {
    return NewFinancialsCompaniesItemTaxAreasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaxAreasById provides operations to manage the taxAreas property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) TaxAreasById(id string)(*FinancialsCompaniesItemTaxAreasTaxAreaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxArea%2Did"] = id
    }
    return NewFinancialsCompaniesItemTaxAreasTaxAreaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaxGroups provides operations to manage the taxGroups property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) TaxGroups()(*FinancialsCompaniesItemTaxGroupsRequestBuilder) {
    return NewFinancialsCompaniesItemTaxGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaxGroupsById provides operations to manage the taxGroups property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) TaxGroupsById(id string)(*FinancialsCompaniesItemTaxGroupsTaxGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxGroup%2Did"] = id
    }
    return NewFinancialsCompaniesItemTaxGroupsTaxGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnitsOfMeasure provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) UnitsOfMeasure()(*FinancialsCompaniesItemUnitsOfMeasureRequestBuilder) {
    return NewFinancialsCompaniesItemUnitsOfMeasureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnitsOfMeasureById provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) UnitsOfMeasureById(id string)(*FinancialsCompaniesItemUnitsOfMeasureUnitOfMeasureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unitOfMeasure%2Did"] = id
    }
    return NewFinancialsCompaniesItemUnitsOfMeasureUnitOfMeasureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Vendors provides operations to manage the vendors property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) Vendors()(*FinancialsCompaniesItemVendorsRequestBuilder) {
    return NewFinancialsCompaniesItemVendorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VendorsById provides operations to manage the vendors property of the microsoft.graph.company entity.
func (m *FinancialsCompaniesCompanyItemRequestBuilder) VendorsById(id string)(*FinancialsCompaniesItemVendorsVendorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vendor%2Did"] = id
    }
    return NewFinancialsCompaniesItemVendorsVendorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
