package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Company 
type Company struct {
    Entity
    // 
    accounts []Account;
    // 
    agedAccountsPayable []AgedAccountsPayable;
    // 
    agedAccountsReceivable []AgedAccountsReceivable;
    // 
    businessProfileId *string;
    // 
    companyInformation []CompanyInformation;
    // 
    countriesRegions []CountryRegion;
    // 
    currencies []Currency;
    // 
    customerPaymentJournals []CustomerPaymentJournal;
    // 
    customerPayments []CustomerPayment;
    // 
    customers []Customer;
    // 
    dimensions []Dimension;
    // 
    dimensionValues []DimensionValue;
    // 
    displayName *string;
    // 
    employees []Employee;
    // 
    generalLedgerEntries []GeneralLedgerEntry;
    // 
    itemCategories []ItemCategory;
    // 
    items []Item;
    // 
    journalLines []JournalLine;
    // 
    journals []Journal;
    // 
    name *string;
    // 
    paymentMethods []PaymentMethod;
    // 
    paymentTerms []PaymentTerm;
    // 
    picture []Picture;
    // 
    purchaseInvoiceLines []PurchaseInvoiceLine;
    // 
    purchaseInvoices []PurchaseInvoice;
    // 
    salesCreditMemoLines []SalesCreditMemoLine;
    // 
    salesCreditMemos []SalesCreditMemo;
    // 
    salesInvoiceLines []SalesInvoiceLine;
    // 
    salesInvoices []SalesInvoice;
    // 
    salesOrderLines []SalesOrderLine;
    // 
    salesOrders []SalesOrder;
    // 
    salesQuoteLines []SalesQuoteLine;
    // 
    salesQuotes []SalesQuote;
    // 
    shipmentMethods []ShipmentMethod;
    // 
    systemVersion *string;
    // 
    taxAreas []TaxArea;
    // 
    taxGroups []TaxGroup;
    // 
    unitsOfMeasure []UnitOfMeasure;
    // 
    vendors []Vendor_escaped;
}
// NewCompany instantiates a new company and sets the default values.
func NewCompany()(*Company) {
    m := &Company{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccounts gets the accounts property value. 
func (m *Company) GetAccounts()([]Account) {
    if m == nil {
        return nil
    } else {
        return m.accounts
    }
}
// GetAgedAccountsPayable gets the agedAccountsPayable property value. 
func (m *Company) GetAgedAccountsPayable()([]AgedAccountsPayable) {
    if m == nil {
        return nil
    } else {
        return m.agedAccountsPayable
    }
}
// GetAgedAccountsReceivable gets the agedAccountsReceivable property value. 
func (m *Company) GetAgedAccountsReceivable()([]AgedAccountsReceivable) {
    if m == nil {
        return nil
    } else {
        return m.agedAccountsReceivable
    }
}
// GetBusinessProfileId gets the businessProfileId property value. 
func (m *Company) GetBusinessProfileId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessProfileId
    }
}
// GetCompanyInformation gets the companyInformation property value. 
func (m *Company) GetCompanyInformation()([]CompanyInformation) {
    if m == nil {
        return nil
    } else {
        return m.companyInformation
    }
}
// GetCountriesRegions gets the countriesRegions property value. 
func (m *Company) GetCountriesRegions()([]CountryRegion) {
    if m == nil {
        return nil
    } else {
        return m.countriesRegions
    }
}
// GetCurrencies gets the currencies property value. 
func (m *Company) GetCurrencies()([]Currency) {
    if m == nil {
        return nil
    } else {
        return m.currencies
    }
}
// GetCustomerPaymentJournals gets the customerPaymentJournals property value. 
func (m *Company) GetCustomerPaymentJournals()([]CustomerPaymentJournal) {
    if m == nil {
        return nil
    } else {
        return m.customerPaymentJournals
    }
}
// GetCustomerPayments gets the customerPayments property value. 
func (m *Company) GetCustomerPayments()([]CustomerPayment) {
    if m == nil {
        return nil
    } else {
        return m.customerPayments
    }
}
// GetCustomers gets the customers property value. 
func (m *Company) GetCustomers()([]Customer) {
    if m == nil {
        return nil
    } else {
        return m.customers
    }
}
// GetDimensions gets the dimensions property value. 
func (m *Company) GetDimensions()([]Dimension) {
    if m == nil {
        return nil
    } else {
        return m.dimensions
    }
}
// GetDimensionValues gets the dimensionValues property value. 
func (m *Company) GetDimensionValues()([]DimensionValue) {
    if m == nil {
        return nil
    } else {
        return m.dimensionValues
    }
}
// GetDisplayName gets the displayName property value. 
func (m *Company) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmployees gets the employees property value. 
func (m *Company) GetEmployees()([]Employee) {
    if m == nil {
        return nil
    } else {
        return m.employees
    }
}
// GetGeneralLedgerEntries gets the generalLedgerEntries property value. 
func (m *Company) GetGeneralLedgerEntries()([]GeneralLedgerEntry) {
    if m == nil {
        return nil
    } else {
        return m.generalLedgerEntries
    }
}
// GetItemCategories gets the itemCategories property value. 
func (m *Company) GetItemCategories()([]ItemCategory) {
    if m == nil {
        return nil
    } else {
        return m.itemCategories
    }
}
// GetItems gets the items property value. 
func (m *Company) GetItems()([]Item) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetJournalLines gets the journalLines property value. 
func (m *Company) GetJournalLines()([]JournalLine) {
    if m == nil {
        return nil
    } else {
        return m.journalLines
    }
}
// GetJournals gets the journals property value. 
func (m *Company) GetJournals()([]Journal) {
    if m == nil {
        return nil
    } else {
        return m.journals
    }
}
// GetName gets the name property value. 
func (m *Company) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPaymentMethods gets the paymentMethods property value. 
func (m *Company) GetPaymentMethods()([]PaymentMethod) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethods
    }
}
// GetPaymentTerms gets the paymentTerms property value. 
func (m *Company) GetPaymentTerms()([]PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerms
    }
}
// GetPicture gets the picture property value. 
func (m *Company) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// GetPurchaseInvoiceLines gets the purchaseInvoiceLines property value. 
func (m *Company) GetPurchaseInvoiceLines()([]PurchaseInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoiceLines
    }
}
// GetPurchaseInvoices gets the purchaseInvoices property value. 
func (m *Company) GetPurchaseInvoices()([]PurchaseInvoice) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoices
    }
}
// GetSalesCreditMemoLines gets the salesCreditMemoLines property value. 
func (m *Company) GetSalesCreditMemoLines()([]SalesCreditMemoLine) {
    if m == nil {
        return nil
    } else {
        return m.salesCreditMemoLines
    }
}
// GetSalesCreditMemos gets the salesCreditMemos property value. 
func (m *Company) GetSalesCreditMemos()([]SalesCreditMemo) {
    if m == nil {
        return nil
    } else {
        return m.salesCreditMemos
    }
}
// GetSalesInvoiceLines gets the salesInvoiceLines property value. 
func (m *Company) GetSalesInvoiceLines()([]SalesInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.salesInvoiceLines
    }
}
// GetSalesInvoices gets the salesInvoices property value. 
func (m *Company) GetSalesInvoices()([]SalesInvoice) {
    if m == nil {
        return nil
    } else {
        return m.salesInvoices
    }
}
// GetSalesOrderLines gets the salesOrderLines property value. 
func (m *Company) GetSalesOrderLines()([]SalesOrderLine) {
    if m == nil {
        return nil
    } else {
        return m.salesOrderLines
    }
}
// GetSalesOrders gets the salesOrders property value. 
func (m *Company) GetSalesOrders()([]SalesOrder) {
    if m == nil {
        return nil
    } else {
        return m.salesOrders
    }
}
// GetSalesQuoteLines gets the salesQuoteLines property value. 
func (m *Company) GetSalesQuoteLines()([]SalesQuoteLine) {
    if m == nil {
        return nil
    } else {
        return m.salesQuoteLines
    }
}
// GetSalesQuotes gets the salesQuotes property value. 
func (m *Company) GetSalesQuotes()([]SalesQuote) {
    if m == nil {
        return nil
    } else {
        return m.salesQuotes
    }
}
// GetShipmentMethods gets the shipmentMethods property value. 
func (m *Company) GetShipmentMethods()([]ShipmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethods
    }
}
// GetSystemVersion gets the systemVersion property value. 
func (m *Company) GetSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.systemVersion
    }
}
// GetTaxAreas gets the taxAreas property value. 
func (m *Company) GetTaxAreas()([]TaxArea) {
    if m == nil {
        return nil
    } else {
        return m.taxAreas
    }
}
// GetTaxGroups gets the taxGroups property value. 
func (m *Company) GetTaxGroups()([]TaxGroup) {
    if m == nil {
        return nil
    } else {
        return m.taxGroups
    }
}
// GetUnitsOfMeasure gets the unitsOfMeasure property value. 
func (m *Company) GetUnitsOfMeasure()([]UnitOfMeasure) {
    if m == nil {
        return nil
    } else {
        return m.unitsOfMeasure
    }
}
// GetVendors gets the vendors property value. 
func (m *Company) GetVendors()([]Vendor_escaped) {
    if m == nil {
        return nil
    } else {
        return m.vendors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Company) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccount() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Account, len(val))
            for i, v := range val {
                res[i] = *(v.(*Account))
            }
            m.SetAccounts(res)
        }
        return nil
    }
    res["agedAccountsPayable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgedAccountsPayable() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgedAccountsPayable, len(val))
            for i, v := range val {
                res[i] = *(v.(*AgedAccountsPayable))
            }
            m.SetAgedAccountsPayable(res)
        }
        return nil
    }
    res["agedAccountsReceivable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgedAccountsReceivable() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgedAccountsReceivable, len(val))
            for i, v := range val {
                res[i] = *(v.(*AgedAccountsReceivable))
            }
            m.SetAgedAccountsReceivable(res)
        }
        return nil
    }
    res["businessProfileId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessProfileId(val)
        }
        return nil
    }
    res["companyInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanyInformation, len(val))
            for i, v := range val {
                res[i] = *(v.(*CompanyInformation))
            }
            m.SetCompanyInformation(res)
        }
        return nil
    }
    res["countriesRegions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCountryRegion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CountryRegion, len(val))
            for i, v := range val {
                res[i] = *(v.(*CountryRegion))
            }
            m.SetCountriesRegions(res)
        }
        return nil
    }
    res["currencies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrency() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Currency, len(val))
            for i, v := range val {
                res[i] = *(v.(*Currency))
            }
            m.SetCurrencies(res)
        }
        return nil
    }
    res["customerPaymentJournals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomerPaymentJournal() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerPaymentJournal, len(val))
            for i, v := range val {
                res[i] = *(v.(*CustomerPaymentJournal))
            }
            m.SetCustomerPaymentJournals(res)
        }
        return nil
    }
    res["customerPayments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomerPayment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerPayment, len(val))
            for i, v := range val {
                res[i] = *(v.(*CustomerPayment))
            }
            m.SetCustomerPayments(res)
        }
        return nil
    }
    res["customers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Customer, len(val))
            for i, v := range val {
                res[i] = *(v.(*Customer))
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["dimensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDimension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Dimension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Dimension))
            }
            m.SetDimensions(res)
        }
        return nil
    }
    res["dimensionValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDimensionValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DimensionValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*DimensionValue))
            }
            m.SetDimensionValues(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["employees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmployee() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Employee, len(val))
            for i, v := range val {
                res[i] = *(v.(*Employee))
            }
            m.SetEmployees(res)
        }
        return nil
    }
    res["generalLedgerEntries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeneralLedgerEntry() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GeneralLedgerEntry, len(val))
            for i, v := range val {
                res[i] = *(v.(*GeneralLedgerEntry))
            }
            m.SetGeneralLedgerEntries(res)
        }
        return nil
    }
    res["itemCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemCategory))
            }
            m.SetItemCategories(res)
        }
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Item, len(val))
            for i, v := range val {
                res[i] = *(v.(*Item))
            }
            m.SetItems(res)
        }
        return nil
    }
    res["journalLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJournalLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JournalLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*JournalLine))
            }
            m.SetJournalLines(res)
        }
        return nil
    }
    res["journals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJournal() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Journal, len(val))
            for i, v := range val {
                res[i] = *(v.(*Journal))
            }
            m.SetJournals(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["paymentMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PaymentMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*PaymentMethod))
            }
            m.SetPaymentMethods(res)
        }
        return nil
    }
    res["paymentTerms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentTerm() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PaymentTerm, len(val))
            for i, v := range val {
                res[i] = *(v.(*PaymentTerm))
            }
            m.SetPaymentTerms(res)
        }
        return nil
    }
    res["picture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPicture() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Picture, len(val))
            for i, v := range val {
                res[i] = *(v.(*Picture))
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["purchaseInvoiceLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPurchaseInvoiceLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*PurchaseInvoiceLine))
            }
            m.SetPurchaseInvoiceLines(res)
        }
        return nil
    }
    res["purchaseInvoices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPurchaseInvoice() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoice, len(val))
            for i, v := range val {
                res[i] = *(v.(*PurchaseInvoice))
            }
            m.SetPurchaseInvoices(res)
        }
        return nil
    }
    res["salesCreditMemoLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesCreditMemoLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesCreditMemoLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesCreditMemoLine))
            }
            m.SetSalesCreditMemoLines(res)
        }
        return nil
    }
    res["salesCreditMemos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesCreditMemo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesCreditMemo, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesCreditMemo))
            }
            m.SetSalesCreditMemos(res)
        }
        return nil
    }
    res["salesInvoiceLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesInvoiceLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesInvoiceLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesInvoiceLine))
            }
            m.SetSalesInvoiceLines(res)
        }
        return nil
    }
    res["salesInvoices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesInvoice() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesInvoice, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesInvoice))
            }
            m.SetSalesInvoices(res)
        }
        return nil
    }
    res["salesOrderLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesOrderLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesOrderLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesOrderLine))
            }
            m.SetSalesOrderLines(res)
        }
        return nil
    }
    res["salesOrders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesOrder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesOrder, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesOrder))
            }
            m.SetSalesOrders(res)
        }
        return nil
    }
    res["salesQuoteLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesQuoteLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuoteLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesQuoteLine))
            }
            m.SetSalesQuoteLines(res)
        }
        return nil
    }
    res["salesQuotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesQuote() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuote, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesQuote))
            }
            m.SetSalesQuotes(res)
        }
        return nil
    }
    res["shipmentMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShipmentMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ShipmentMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*ShipmentMethod))
            }
            m.SetShipmentMethods(res)
        }
        return nil
    }
    res["systemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemVersion(val)
        }
        return nil
    }
    res["taxAreas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTaxArea() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaxArea, len(val))
            for i, v := range val {
                res[i] = *(v.(*TaxArea))
            }
            m.SetTaxAreas(res)
        }
        return nil
    }
    res["taxGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTaxGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaxGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*TaxGroup))
            }
            m.SetTaxGroups(res)
        }
        return nil
    }
    res["unitsOfMeasure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnitOfMeasure() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnitOfMeasure, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnitOfMeasure))
            }
            m.SetUnitsOfMeasure(res)
        }
        return nil
    }
    res["vendors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVendor_escaped() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Vendor_escaped, len(val))
            for i, v := range val {
                res[i] = *(v.(*Vendor_escaped))
            }
            m.SetVendors(res)
        }
        return nil
    }
    return res
}
func (m *Company) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Company) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccounts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsPayable() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgedAccountsPayable()))
        for i, v := range m.GetAgedAccountsPayable() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agedAccountsPayable", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsReceivable() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgedAccountsReceivable()))
        for i, v := range m.GetAgedAccountsReceivable() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCompanyInformation()))
        for i, v := range m.GetCompanyInformation() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("companyInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCountriesRegions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCountriesRegions()))
        for i, v := range m.GetCountriesRegions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("countriesRegions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCurrencies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCurrencies()))
        for i, v := range m.GetCurrencies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("currencies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPaymentJournals() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomerPaymentJournals()))
        for i, v := range m.GetCustomerPaymentJournals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customerPaymentJournals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPayments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomerPayments()))
        for i, v := range m.GetCustomerPayments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customerPayments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDimensions()))
        for i, v := range m.GetDimensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dimensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensionValues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDimensionValues()))
        for i, v := range m.GetDimensionValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmployees()))
        for i, v := range m.GetEmployees() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("employees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGeneralLedgerEntries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGeneralLedgerEntries()))
        for i, v := range m.GetGeneralLedgerEntries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("generalLedgerEntries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItemCategories() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItemCategories()))
        for i, v := range m.GetItemCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("itemCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournalLines() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJournalLines()))
        for i, v := range m.GetJournalLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("journalLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournals() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJournals()))
        for i, v := range m.GetJournals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPaymentMethods()))
        for i, v := range m.GetPaymentMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("paymentMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPaymentTerms() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPaymentTerms()))
        for i, v := range m.GetPaymentTerms() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("paymentTerms", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPicture() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoiceLines() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPurchaseInvoiceLines()))
        for i, v := range m.GetPurchaseInvoiceLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("purchaseInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoices() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPurchaseInvoices()))
        for i, v := range m.GetPurchaseInvoices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("purchaseInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemoLines() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesCreditMemoLines()))
        for i, v := range m.GetSalesCreditMemoLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesCreditMemoLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemos() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesCreditMemos()))
        for i, v := range m.GetSalesCreditMemos() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesCreditMemos", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoiceLines() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesInvoiceLines()))
        for i, v := range m.GetSalesInvoiceLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoices() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesInvoices()))
        for i, v := range m.GetSalesInvoices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrderLines() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesOrderLines()))
        for i, v := range m.GetSalesOrderLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesOrderLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrders() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesOrders()))
        for i, v := range m.GetSalesOrders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesOrders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuoteLines() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesQuoteLines()))
        for i, v := range m.GetSalesQuoteLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesQuoteLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuotes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesQuotes()))
        for i, v := range m.GetSalesQuotes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesQuotes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetShipmentMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetShipmentMethods()))
        for i, v := range m.GetShipmentMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaxAreas()))
        for i, v := range m.GetTaxAreas() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taxAreas", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaxGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaxGroups()))
        for i, v := range m.GetTaxGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taxGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUnitsOfMeasure() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUnitsOfMeasure()))
        for i, v := range m.GetUnitsOfMeasure() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("unitsOfMeasure", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVendors() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVendors()))
        for i, v := range m.GetVendors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("vendors", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. 
func (m *Company) SetAccounts(value []Account)() {
    if m != nil {
        m.accounts = value
    }
}
// SetAgedAccountsPayable sets the agedAccountsPayable property value. 
func (m *Company) SetAgedAccountsPayable(value []AgedAccountsPayable)() {
    if m != nil {
        m.agedAccountsPayable = value
    }
}
// SetAgedAccountsReceivable sets the agedAccountsReceivable property value. 
func (m *Company) SetAgedAccountsReceivable(value []AgedAccountsReceivable)() {
    if m != nil {
        m.agedAccountsReceivable = value
    }
}
// SetBusinessProfileId sets the businessProfileId property value. 
func (m *Company) SetBusinessProfileId(value *string)() {
    if m != nil {
        m.businessProfileId = value
    }
}
// SetCompanyInformation sets the companyInformation property value. 
func (m *Company) SetCompanyInformation(value []CompanyInformation)() {
    if m != nil {
        m.companyInformation = value
    }
}
// SetCountriesRegions sets the countriesRegions property value. 
func (m *Company) SetCountriesRegions(value []CountryRegion)() {
    if m != nil {
        m.countriesRegions = value
    }
}
// SetCurrencies sets the currencies property value. 
func (m *Company) SetCurrencies(value []Currency)() {
    if m != nil {
        m.currencies = value
    }
}
// SetCustomerPaymentJournals sets the customerPaymentJournals property value. 
func (m *Company) SetCustomerPaymentJournals(value []CustomerPaymentJournal)() {
    if m != nil {
        m.customerPaymentJournals = value
    }
}
// SetCustomerPayments sets the customerPayments property value. 
func (m *Company) SetCustomerPayments(value []CustomerPayment)() {
    if m != nil {
        m.customerPayments = value
    }
}
// SetCustomers sets the customers property value. 
func (m *Company) SetCustomers(value []Customer)() {
    if m != nil {
        m.customers = value
    }
}
// SetDimensions sets the dimensions property value. 
func (m *Company) SetDimensions(value []Dimension)() {
    if m != nil {
        m.dimensions = value
    }
}
// SetDimensionValues sets the dimensionValues property value. 
func (m *Company) SetDimensionValues(value []DimensionValue)() {
    if m != nil {
        m.dimensionValues = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *Company) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmployees sets the employees property value. 
func (m *Company) SetEmployees(value []Employee)() {
    if m != nil {
        m.employees = value
    }
}
// SetGeneralLedgerEntries sets the generalLedgerEntries property value. 
func (m *Company) SetGeneralLedgerEntries(value []GeneralLedgerEntry)() {
    if m != nil {
        m.generalLedgerEntries = value
    }
}
// SetItemCategories sets the itemCategories property value. 
func (m *Company) SetItemCategories(value []ItemCategory)() {
    if m != nil {
        m.itemCategories = value
    }
}
// SetItems sets the items property value. 
func (m *Company) SetItems(value []Item)() {
    if m != nil {
        m.items = value
    }
}
// SetJournalLines sets the journalLines property value. 
func (m *Company) SetJournalLines(value []JournalLine)() {
    if m != nil {
        m.journalLines = value
    }
}
// SetJournals sets the journals property value. 
func (m *Company) SetJournals(value []Journal)() {
    if m != nil {
        m.journals = value
    }
}
// SetName sets the name property value. 
func (m *Company) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPaymentMethods sets the paymentMethods property value. 
func (m *Company) SetPaymentMethods(value []PaymentMethod)() {
    if m != nil {
        m.paymentMethods = value
    }
}
// SetPaymentTerms sets the paymentTerms property value. 
func (m *Company) SetPaymentTerms(value []PaymentTerm)() {
    if m != nil {
        m.paymentTerms = value
    }
}
// SetPicture sets the picture property value. 
func (m *Company) SetPicture(value []Picture)() {
    if m != nil {
        m.picture = value
    }
}
// SetPurchaseInvoiceLines sets the purchaseInvoiceLines property value. 
func (m *Company) SetPurchaseInvoiceLines(value []PurchaseInvoiceLine)() {
    if m != nil {
        m.purchaseInvoiceLines = value
    }
}
// SetPurchaseInvoices sets the purchaseInvoices property value. 
func (m *Company) SetPurchaseInvoices(value []PurchaseInvoice)() {
    if m != nil {
        m.purchaseInvoices = value
    }
}
// SetSalesCreditMemoLines sets the salesCreditMemoLines property value. 
func (m *Company) SetSalesCreditMemoLines(value []SalesCreditMemoLine)() {
    if m != nil {
        m.salesCreditMemoLines = value
    }
}
// SetSalesCreditMemos sets the salesCreditMemos property value. 
func (m *Company) SetSalesCreditMemos(value []SalesCreditMemo)() {
    if m != nil {
        m.salesCreditMemos = value
    }
}
// SetSalesInvoiceLines sets the salesInvoiceLines property value. 
func (m *Company) SetSalesInvoiceLines(value []SalesInvoiceLine)() {
    if m != nil {
        m.salesInvoiceLines = value
    }
}
// SetSalesInvoices sets the salesInvoices property value. 
func (m *Company) SetSalesInvoices(value []SalesInvoice)() {
    if m != nil {
        m.salesInvoices = value
    }
}
// SetSalesOrderLines sets the salesOrderLines property value. 
func (m *Company) SetSalesOrderLines(value []SalesOrderLine)() {
    if m != nil {
        m.salesOrderLines = value
    }
}
// SetSalesOrders sets the salesOrders property value. 
func (m *Company) SetSalesOrders(value []SalesOrder)() {
    if m != nil {
        m.salesOrders = value
    }
}
// SetSalesQuoteLines sets the salesQuoteLines property value. 
func (m *Company) SetSalesQuoteLines(value []SalesQuoteLine)() {
    if m != nil {
        m.salesQuoteLines = value
    }
}
// SetSalesQuotes sets the salesQuotes property value. 
func (m *Company) SetSalesQuotes(value []SalesQuote)() {
    if m != nil {
        m.salesQuotes = value
    }
}
// SetShipmentMethods sets the shipmentMethods property value. 
func (m *Company) SetShipmentMethods(value []ShipmentMethod)() {
    if m != nil {
        m.shipmentMethods = value
    }
}
// SetSystemVersion sets the systemVersion property value. 
func (m *Company) SetSystemVersion(value *string)() {
    if m != nil {
        m.systemVersion = value
    }
}
// SetTaxAreas sets the taxAreas property value. 
func (m *Company) SetTaxAreas(value []TaxArea)() {
    if m != nil {
        m.taxAreas = value
    }
}
// SetTaxGroups sets the taxGroups property value. 
func (m *Company) SetTaxGroups(value []TaxGroup)() {
    if m != nil {
        m.taxGroups = value
    }
}
// SetUnitsOfMeasure sets the unitsOfMeasure property value. 
func (m *Company) SetUnitsOfMeasure(value []UnitOfMeasure)() {
    if m != nil {
        m.unitsOfMeasure = value
    }
}
// SetVendors sets the vendors property value. 
func (m *Company) SetVendors(value []Vendor_escaped)() {
    if m != nil {
        m.vendors = value
    }
}
