package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Company 
type Company struct {
    Entity
    // The accounts property
    accounts []Accountable
    // The agedAccountsPayable property
    agedAccountsPayable []AgedAccountsPayableable
    // The agedAccountsReceivable property
    agedAccountsReceivable []AgedAccountsReceivableable
    // The businessProfileId property
    businessProfileId *string
    // The companyInformation property
    companyInformation []CompanyInformationable
    // The countriesRegions property
    countriesRegions []CountryRegionable
    // The currencies property
    currencies []Currencyable
    // The customerPaymentJournals property
    customerPaymentJournals []CustomerPaymentJournalable
    // The customerPayments property
    customerPayments []CustomerPaymentable
    // The customers property
    customers []Customerable
    // The dimensions property
    dimensions []Dimensionable
    // The dimensionValues property
    dimensionValues []DimensionValueable
    // The displayName property
    displayName *string
    // The employees property
    employees []Employeeable
    // The generalLedgerEntries property
    generalLedgerEntries []GeneralLedgerEntryable
    // The itemCategories property
    itemCategories []ItemCategoryable
    // The items property
    items []Itemable
    // The journalLines property
    journalLines []JournalLineable
    // The journals property
    journals []Journalable
    // The name property
    name *string
    // The paymentMethods property
    paymentMethods []PaymentMethodable
    // The paymentTerms property
    paymentTerms []PaymentTermable
    // The picture property
    picture []Pictureable
    // The purchaseInvoiceLines property
    purchaseInvoiceLines []PurchaseInvoiceLineable
    // The purchaseInvoices property
    purchaseInvoices []PurchaseInvoiceable
    // The salesCreditMemoLines property
    salesCreditMemoLines []SalesCreditMemoLineable
    // The salesCreditMemos property
    salesCreditMemos []SalesCreditMemoable
    // The salesInvoiceLines property
    salesInvoiceLines []SalesInvoiceLineable
    // The salesInvoices property
    salesInvoices []SalesInvoiceable
    // The salesOrderLines property
    salesOrderLines []SalesOrderLineable
    // The salesOrders property
    salesOrders []SalesOrderable
    // The salesQuoteLines property
    salesQuoteLines []SalesQuoteLineable
    // The salesQuotes property
    salesQuotes []SalesQuoteable
    // The shipmentMethods property
    shipmentMethods []ShipmentMethodable
    // The systemVersion property
    systemVersion *string
    // The taxAreas property
    taxAreas []TaxAreaable
    // The taxGroups property
    taxGroups []TaxGroupable
    // The unitsOfMeasure property
    unitsOfMeasure []UnitOfMeasureable
    // The vendors property
    vendors []Vendor_escapedable
}
// NewCompany instantiates a new Company and sets the default values.
func NewCompany()(*Company) {
    m := &Company{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCompanyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompany(), nil
}
// GetAccounts gets the accounts property value. The accounts property
func (m *Company) GetAccounts()([]Accountable) {
    return m.accounts
}
// GetAgedAccountsPayable gets the agedAccountsPayable property value. The agedAccountsPayable property
func (m *Company) GetAgedAccountsPayable()([]AgedAccountsPayableable) {
    return m.agedAccountsPayable
}
// GetAgedAccountsReceivable gets the agedAccountsReceivable property value. The agedAccountsReceivable property
func (m *Company) GetAgedAccountsReceivable()([]AgedAccountsReceivableable) {
    return m.agedAccountsReceivable
}
// GetBusinessProfileId gets the businessProfileId property value. The businessProfileId property
func (m *Company) GetBusinessProfileId()(*string) {
    return m.businessProfileId
}
// GetCompanyInformation gets the companyInformation property value. The companyInformation property
func (m *Company) GetCompanyInformation()([]CompanyInformationable) {
    return m.companyInformation
}
// GetCountriesRegions gets the countriesRegions property value. The countriesRegions property
func (m *Company) GetCountriesRegions()([]CountryRegionable) {
    return m.countriesRegions
}
// GetCurrencies gets the currencies property value. The currencies property
func (m *Company) GetCurrencies()([]Currencyable) {
    return m.currencies
}
// GetCustomerPaymentJournals gets the customerPaymentJournals property value. The customerPaymentJournals property
func (m *Company) GetCustomerPaymentJournals()([]CustomerPaymentJournalable) {
    return m.customerPaymentJournals
}
// GetCustomerPayments gets the customerPayments property value. The customerPayments property
func (m *Company) GetCustomerPayments()([]CustomerPaymentable) {
    return m.customerPayments
}
// GetCustomers gets the customers property value. The customers property
func (m *Company) GetCustomers()([]Customerable) {
    return m.customers
}
// GetDimensions gets the dimensions property value. The dimensions property
func (m *Company) GetDimensions()([]Dimensionable) {
    return m.dimensions
}
// GetDimensionValues gets the dimensionValues property value. The dimensionValues property
func (m *Company) GetDimensionValues()([]DimensionValueable) {
    return m.dimensionValues
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Company) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmployees gets the employees property value. The employees property
func (m *Company) GetEmployees()([]Employeeable) {
    return m.employees
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Company) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accounts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccountFromDiscriminatorValue , m.SetAccounts)
    res["agedAccountsPayable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAgedAccountsPayableFromDiscriminatorValue , m.SetAgedAccountsPayable)
    res["agedAccountsReceivable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAgedAccountsReceivableFromDiscriminatorValue , m.SetAgedAccountsReceivable)
    res["businessProfileId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBusinessProfileId)
    res["companyInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCompanyInformationFromDiscriminatorValue , m.SetCompanyInformation)
    res["countriesRegions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCountryRegionFromDiscriminatorValue , m.SetCountriesRegions)
    res["currencies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCurrencyFromDiscriminatorValue , m.SetCurrencies)
    res["customerPaymentJournals"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomerPaymentJournalFromDiscriminatorValue , m.SetCustomerPaymentJournals)
    res["customerPayments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomerPaymentFromDiscriminatorValue , m.SetCustomerPayments)
    res["customers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomerFromDiscriminatorValue , m.SetCustomers)
    res["dimensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDimensionFromDiscriminatorValue , m.SetDimensions)
    res["dimensionValues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDimensionValueFromDiscriminatorValue , m.SetDimensionValues)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["employees"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEmployeeFromDiscriminatorValue , m.SetEmployees)
    res["generalLedgerEntries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGeneralLedgerEntryFromDiscriminatorValue , m.SetGeneralLedgerEntries)
    res["itemCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemCategoryFromDiscriminatorValue , m.SetItemCategories)
    res["items"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemFromDiscriminatorValue , m.SetItems)
    res["journalLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJournalLineFromDiscriminatorValue , m.SetJournalLines)
    res["journals"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJournalFromDiscriminatorValue , m.SetJournals)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["paymentMethods"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePaymentMethodFromDiscriminatorValue , m.SetPaymentMethods)
    res["paymentTerms"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePaymentTermFromDiscriminatorValue , m.SetPaymentTerms)
    res["picture"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue , m.SetPicture)
    res["purchaseInvoiceLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePurchaseInvoiceLineFromDiscriminatorValue , m.SetPurchaseInvoiceLines)
    res["purchaseInvoices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePurchaseInvoiceFromDiscriminatorValue , m.SetPurchaseInvoices)
    res["salesCreditMemoLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesCreditMemoLineFromDiscriminatorValue , m.SetSalesCreditMemoLines)
    res["salesCreditMemos"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesCreditMemoFromDiscriminatorValue , m.SetSalesCreditMemos)
    res["salesInvoiceLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesInvoiceLineFromDiscriminatorValue , m.SetSalesInvoiceLines)
    res["salesInvoices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesInvoiceFromDiscriminatorValue , m.SetSalesInvoices)
    res["salesOrderLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesOrderLineFromDiscriminatorValue , m.SetSalesOrderLines)
    res["salesOrders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesOrderFromDiscriminatorValue , m.SetSalesOrders)
    res["salesQuoteLines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesQuoteLineFromDiscriminatorValue , m.SetSalesQuoteLines)
    res["salesQuotes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSalesQuoteFromDiscriminatorValue , m.SetSalesQuotes)
    res["shipmentMethods"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateShipmentMethodFromDiscriminatorValue , m.SetShipmentMethods)
    res["systemVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSystemVersion)
    res["taxAreas"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaxAreaFromDiscriminatorValue , m.SetTaxAreas)
    res["taxGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaxGroupFromDiscriminatorValue , m.SetTaxGroups)
    res["unitsOfMeasure"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnitOfMeasureFromDiscriminatorValue , m.SetUnitsOfMeasure)
    res["vendors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVendor_escapedFromDiscriminatorValue , m.SetVendors)
    return res
}
// GetGeneralLedgerEntries gets the generalLedgerEntries property value. The generalLedgerEntries property
func (m *Company) GetGeneralLedgerEntries()([]GeneralLedgerEntryable) {
    return m.generalLedgerEntries
}
// GetItemCategories gets the itemCategories property value. The itemCategories property
func (m *Company) GetItemCategories()([]ItemCategoryable) {
    return m.itemCategories
}
// GetItems gets the items property value. The items property
func (m *Company) GetItems()([]Itemable) {
    return m.items
}
// GetJournalLines gets the journalLines property value. The journalLines property
func (m *Company) GetJournalLines()([]JournalLineable) {
    return m.journalLines
}
// GetJournals gets the journals property value. The journals property
func (m *Company) GetJournals()([]Journalable) {
    return m.journals
}
// GetName gets the name property value. The name property
func (m *Company) GetName()(*string) {
    return m.name
}
// GetPaymentMethods gets the paymentMethods property value. The paymentMethods property
func (m *Company) GetPaymentMethods()([]PaymentMethodable) {
    return m.paymentMethods
}
// GetPaymentTerms gets the paymentTerms property value. The paymentTerms property
func (m *Company) GetPaymentTerms()([]PaymentTermable) {
    return m.paymentTerms
}
// GetPicture gets the picture property value. The picture property
func (m *Company) GetPicture()([]Pictureable) {
    return m.picture
}
// GetPurchaseInvoiceLines gets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *Company) GetPurchaseInvoiceLines()([]PurchaseInvoiceLineable) {
    return m.purchaseInvoiceLines
}
// GetPurchaseInvoices gets the purchaseInvoices property value. The purchaseInvoices property
func (m *Company) GetPurchaseInvoices()([]PurchaseInvoiceable) {
    return m.purchaseInvoices
}
// GetSalesCreditMemoLines gets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *Company) GetSalesCreditMemoLines()([]SalesCreditMemoLineable) {
    return m.salesCreditMemoLines
}
// GetSalesCreditMemos gets the salesCreditMemos property value. The salesCreditMemos property
func (m *Company) GetSalesCreditMemos()([]SalesCreditMemoable) {
    return m.salesCreditMemos
}
// GetSalesInvoiceLines gets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *Company) GetSalesInvoiceLines()([]SalesInvoiceLineable) {
    return m.salesInvoiceLines
}
// GetSalesInvoices gets the salesInvoices property value. The salesInvoices property
func (m *Company) GetSalesInvoices()([]SalesInvoiceable) {
    return m.salesInvoices
}
// GetSalesOrderLines gets the salesOrderLines property value. The salesOrderLines property
func (m *Company) GetSalesOrderLines()([]SalesOrderLineable) {
    return m.salesOrderLines
}
// GetSalesOrders gets the salesOrders property value. The salesOrders property
func (m *Company) GetSalesOrders()([]SalesOrderable) {
    return m.salesOrders
}
// GetSalesQuoteLines gets the salesQuoteLines property value. The salesQuoteLines property
func (m *Company) GetSalesQuoteLines()([]SalesQuoteLineable) {
    return m.salesQuoteLines
}
// GetSalesQuotes gets the salesQuotes property value. The salesQuotes property
func (m *Company) GetSalesQuotes()([]SalesQuoteable) {
    return m.salesQuotes
}
// GetShipmentMethods gets the shipmentMethods property value. The shipmentMethods property
func (m *Company) GetShipmentMethods()([]ShipmentMethodable) {
    return m.shipmentMethods
}
// GetSystemVersion gets the systemVersion property value. The systemVersion property
func (m *Company) GetSystemVersion()(*string) {
    return m.systemVersion
}
// GetTaxAreas gets the taxAreas property value. The taxAreas property
func (m *Company) GetTaxAreas()([]TaxAreaable) {
    return m.taxAreas
}
// GetTaxGroups gets the taxGroups property value. The taxGroups property
func (m *Company) GetTaxGroups()([]TaxGroupable) {
    return m.taxGroups
}
// GetUnitsOfMeasure gets the unitsOfMeasure property value. The unitsOfMeasure property
func (m *Company) GetUnitsOfMeasure()([]UnitOfMeasureable) {
    return m.unitsOfMeasure
}
// GetVendors gets the vendors property value. The vendors property
func (m *Company) GetVendors()([]Vendor_escapedable) {
    return m.vendors
}
// Serialize serializes information the current object
func (m *Company) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccounts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccounts())
        err = writer.WriteCollectionOfObjectValues("accounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsPayable() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAgedAccountsPayable())
        err = writer.WriteCollectionOfObjectValues("agedAccountsPayable", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsReceivable() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAgedAccountsReceivable())
        err = writer.WriteCollectionOfObjectValues("agedAccountsReceivable", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessProfileId", m.GetBusinessProfileId())
        if err != nil {
            return err
        }
    }
    if m.GetCompanyInformation() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCompanyInformation())
        err = writer.WriteCollectionOfObjectValues("companyInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCountriesRegions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCountriesRegions())
        err = writer.WriteCollectionOfObjectValues("countriesRegions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCurrencies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCurrencies())
        err = writer.WriteCollectionOfObjectValues("currencies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPaymentJournals() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomerPaymentJournals())
        err = writer.WriteCollectionOfObjectValues("customerPaymentJournals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPayments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomerPayments())
        err = writer.WriteCollectionOfObjectValues("customerPayments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomers())
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDimensions())
        err = writer.WriteCollectionOfObjectValues("dimensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensionValues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDimensionValues())
        err = writer.WriteCollectionOfObjectValues("dimensionValues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEmployees() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEmployees())
        err = writer.WriteCollectionOfObjectValues("employees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGeneralLedgerEntries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGeneralLedgerEntries())
        err = writer.WriteCollectionOfObjectValues("generalLedgerEntries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItemCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetItemCategories())
        err = writer.WriteCollectionOfObjectValues("itemCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetItems())
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournalLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetJournalLines())
        err = writer.WriteCollectionOfObjectValues("journalLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournals() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetJournals())
        err = writer.WriteCollectionOfObjectValues("journals", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetPaymentMethods() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPaymentMethods())
        err = writer.WriteCollectionOfObjectValues("paymentMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPaymentTerms() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPaymentTerms())
        err = writer.WriteCollectionOfObjectValues("paymentTerms", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPicture() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPicture())
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoiceLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPurchaseInvoiceLines())
        err = writer.WriteCollectionOfObjectValues("purchaseInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPurchaseInvoices())
        err = writer.WriteCollectionOfObjectValues("purchaseInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemoLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesCreditMemoLines())
        err = writer.WriteCollectionOfObjectValues("salesCreditMemoLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemos() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesCreditMemos())
        err = writer.WriteCollectionOfObjectValues("salesCreditMemos", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoiceLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesInvoiceLines())
        err = writer.WriteCollectionOfObjectValues("salesInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesInvoices())
        err = writer.WriteCollectionOfObjectValues("salesInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrderLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesOrderLines())
        err = writer.WriteCollectionOfObjectValues("salesOrderLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrders() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesOrders())
        err = writer.WriteCollectionOfObjectValues("salesOrders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuoteLines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesQuoteLines())
        err = writer.WriteCollectionOfObjectValues("salesQuoteLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuotes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSalesQuotes())
        err = writer.WriteCollectionOfObjectValues("salesQuotes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetShipmentMethods() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetShipmentMethods())
        err = writer.WriteCollectionOfObjectValues("shipmentMethods", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("systemVersion", m.GetSystemVersion())
        if err != nil {
            return err
        }
    }
    if m.GetTaxAreas() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaxAreas())
        err = writer.WriteCollectionOfObjectValues("taxAreas", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaxGroups() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaxGroups())
        err = writer.WriteCollectionOfObjectValues("taxGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUnitsOfMeasure() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUnitsOfMeasure())
        err = writer.WriteCollectionOfObjectValues("unitsOfMeasure", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVendors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetVendors())
        err = writer.WriteCollectionOfObjectValues("vendors", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. The accounts property
func (m *Company) SetAccounts(value []Accountable)() {
    m.accounts = value
}
// SetAgedAccountsPayable sets the agedAccountsPayable property value. The agedAccountsPayable property
func (m *Company) SetAgedAccountsPayable(value []AgedAccountsPayableable)() {
    m.agedAccountsPayable = value
}
// SetAgedAccountsReceivable sets the agedAccountsReceivable property value. The agedAccountsReceivable property
func (m *Company) SetAgedAccountsReceivable(value []AgedAccountsReceivableable)() {
    m.agedAccountsReceivable = value
}
// SetBusinessProfileId sets the businessProfileId property value. The businessProfileId property
func (m *Company) SetBusinessProfileId(value *string)() {
    m.businessProfileId = value
}
// SetCompanyInformation sets the companyInformation property value. The companyInformation property
func (m *Company) SetCompanyInformation(value []CompanyInformationable)() {
    m.companyInformation = value
}
// SetCountriesRegions sets the countriesRegions property value. The countriesRegions property
func (m *Company) SetCountriesRegions(value []CountryRegionable)() {
    m.countriesRegions = value
}
// SetCurrencies sets the currencies property value. The currencies property
func (m *Company) SetCurrencies(value []Currencyable)() {
    m.currencies = value
}
// SetCustomerPaymentJournals sets the customerPaymentJournals property value. The customerPaymentJournals property
func (m *Company) SetCustomerPaymentJournals(value []CustomerPaymentJournalable)() {
    m.customerPaymentJournals = value
}
// SetCustomerPayments sets the customerPayments property value. The customerPayments property
func (m *Company) SetCustomerPayments(value []CustomerPaymentable)() {
    m.customerPayments = value
}
// SetCustomers sets the customers property value. The customers property
func (m *Company) SetCustomers(value []Customerable)() {
    m.customers = value
}
// SetDimensions sets the dimensions property value. The dimensions property
func (m *Company) SetDimensions(value []Dimensionable)() {
    m.dimensions = value
}
// SetDimensionValues sets the dimensionValues property value. The dimensionValues property
func (m *Company) SetDimensionValues(value []DimensionValueable)() {
    m.dimensionValues = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Company) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmployees sets the employees property value. The employees property
func (m *Company) SetEmployees(value []Employeeable)() {
    m.employees = value
}
// SetGeneralLedgerEntries sets the generalLedgerEntries property value. The generalLedgerEntries property
func (m *Company) SetGeneralLedgerEntries(value []GeneralLedgerEntryable)() {
    m.generalLedgerEntries = value
}
// SetItemCategories sets the itemCategories property value. The itemCategories property
func (m *Company) SetItemCategories(value []ItemCategoryable)() {
    m.itemCategories = value
}
// SetItems sets the items property value. The items property
func (m *Company) SetItems(value []Itemable)() {
    m.items = value
}
// SetJournalLines sets the journalLines property value. The journalLines property
func (m *Company) SetJournalLines(value []JournalLineable)() {
    m.journalLines = value
}
// SetJournals sets the journals property value. The journals property
func (m *Company) SetJournals(value []Journalable)() {
    m.journals = value
}
// SetName sets the name property value. The name property
func (m *Company) SetName(value *string)() {
    m.name = value
}
// SetPaymentMethods sets the paymentMethods property value. The paymentMethods property
func (m *Company) SetPaymentMethods(value []PaymentMethodable)() {
    m.paymentMethods = value
}
// SetPaymentTerms sets the paymentTerms property value. The paymentTerms property
func (m *Company) SetPaymentTerms(value []PaymentTermable)() {
    m.paymentTerms = value
}
// SetPicture sets the picture property value. The picture property
func (m *Company) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetPurchaseInvoiceLines sets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *Company) SetPurchaseInvoiceLines(value []PurchaseInvoiceLineable)() {
    m.purchaseInvoiceLines = value
}
// SetPurchaseInvoices sets the purchaseInvoices property value. The purchaseInvoices property
func (m *Company) SetPurchaseInvoices(value []PurchaseInvoiceable)() {
    m.purchaseInvoices = value
}
// SetSalesCreditMemoLines sets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *Company) SetSalesCreditMemoLines(value []SalesCreditMemoLineable)() {
    m.salesCreditMemoLines = value
}
// SetSalesCreditMemos sets the salesCreditMemos property value. The salesCreditMemos property
func (m *Company) SetSalesCreditMemos(value []SalesCreditMemoable)() {
    m.salesCreditMemos = value
}
// SetSalesInvoiceLines sets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *Company) SetSalesInvoiceLines(value []SalesInvoiceLineable)() {
    m.salesInvoiceLines = value
}
// SetSalesInvoices sets the salesInvoices property value. The salesInvoices property
func (m *Company) SetSalesInvoices(value []SalesInvoiceable)() {
    m.salesInvoices = value
}
// SetSalesOrderLines sets the salesOrderLines property value. The salesOrderLines property
func (m *Company) SetSalesOrderLines(value []SalesOrderLineable)() {
    m.salesOrderLines = value
}
// SetSalesOrders sets the salesOrders property value. The salesOrders property
func (m *Company) SetSalesOrders(value []SalesOrderable)() {
    m.salesOrders = value
}
// SetSalesQuoteLines sets the salesQuoteLines property value. The salesQuoteLines property
func (m *Company) SetSalesQuoteLines(value []SalesQuoteLineable)() {
    m.salesQuoteLines = value
}
// SetSalesQuotes sets the salesQuotes property value. The salesQuotes property
func (m *Company) SetSalesQuotes(value []SalesQuoteable)() {
    m.salesQuotes = value
}
// SetShipmentMethods sets the shipmentMethods property value. The shipmentMethods property
func (m *Company) SetShipmentMethods(value []ShipmentMethodable)() {
    m.shipmentMethods = value
}
// SetSystemVersion sets the systemVersion property value. The systemVersion property
func (m *Company) SetSystemVersion(value *string)() {
    m.systemVersion = value
}
// SetTaxAreas sets the taxAreas property value. The taxAreas property
func (m *Company) SetTaxAreas(value []TaxAreaable)() {
    m.taxAreas = value
}
// SetTaxGroups sets the taxGroups property value. The taxGroups property
func (m *Company) SetTaxGroups(value []TaxGroupable)() {
    m.taxGroups = value
}
// SetUnitsOfMeasure sets the unitsOfMeasure property value. The unitsOfMeasure property
func (m *Company) SetUnitsOfMeasure(value []UnitOfMeasureable)() {
    m.unitsOfMeasure = value
}
// SetVendors sets the vendors property value. The vendors property
func (m *Company) SetVendors(value []Vendor_escapedable)() {
    m.vendors = value
}
