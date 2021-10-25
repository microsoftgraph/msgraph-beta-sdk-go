package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Company struct {
    Entity
    accounts []Account;
    agedAccountsPayable []AgedAccountsPayable;
    agedAccountsReceivable []AgedAccountsReceivable;
    businessProfileId *string;
    companyInformation []CompanyInformation;
    countriesRegions []CountryRegion;
    currencies []Currency;
    customerPaymentJournals []CustomerPaymentJournal;
    customerPayments []CustomerPayment;
    customers []Customer;
    dimensions []Dimension;
    dimensionValues []DimensionValue;
    displayName *string;
    employees []Employee;
    generalLedgerEntries []GeneralLedgerEntry;
    itemCategories []ItemCategory;
    items []Item;
    journalLines []JournalLine;
    journals []Journal;
    name *string;
    paymentMethods []PaymentMethod;
    paymentTerms []PaymentTerm;
    picture []Picture;
    purchaseInvoiceLines []PurchaseInvoiceLine;
    purchaseInvoices []PurchaseInvoice;
    salesCreditMemoLines []SalesCreditMemoLine;
    salesCreditMemos []SalesCreditMemo;
    salesInvoiceLines []SalesInvoiceLine;
    salesInvoices []SalesInvoice;
    salesOrderLines []SalesOrderLine;
    salesOrders []SalesOrder;
    salesQuoteLines []SalesQuoteLine;
    salesQuotes []SalesQuote;
    shipmentMethods []ShipmentMethod;
    systemVersion *string;
    taxAreas []TaxArea;
    taxGroups []TaxGroup;
    unitsOfMeasure []UnitOfMeasure;
    vendors []Vendor;
}
func NewCompany()(*Company) {
    m := &Company{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Company) GetAccounts()([]Account) {
    if m == nil {
        return nil
    } else {
        return m.accounts
    }
}
func (m *Company) GetAgedAccountsPayable()([]AgedAccountsPayable) {
    if m == nil {
        return nil
    } else {
        return m.agedAccountsPayable
    }
}
func (m *Company) GetAgedAccountsReceivable()([]AgedAccountsReceivable) {
    if m == nil {
        return nil
    } else {
        return m.agedAccountsReceivable
    }
}
func (m *Company) GetBusinessProfileId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessProfileId
    }
}
func (m *Company) GetCompanyInformation()([]CompanyInformation) {
    if m == nil {
        return nil
    } else {
        return m.companyInformation
    }
}
func (m *Company) GetCountriesRegions()([]CountryRegion) {
    if m == nil {
        return nil
    } else {
        return m.countriesRegions
    }
}
func (m *Company) GetCurrencies()([]Currency) {
    if m == nil {
        return nil
    } else {
        return m.currencies
    }
}
func (m *Company) GetCustomerPaymentJournals()([]CustomerPaymentJournal) {
    if m == nil {
        return nil
    } else {
        return m.customerPaymentJournals
    }
}
func (m *Company) GetCustomerPayments()([]CustomerPayment) {
    if m == nil {
        return nil
    } else {
        return m.customerPayments
    }
}
func (m *Company) GetCustomers()([]Customer) {
    if m == nil {
        return nil
    } else {
        return m.customers
    }
}
func (m *Company) GetDimensions()([]Dimension) {
    if m == nil {
        return nil
    } else {
        return m.dimensions
    }
}
func (m *Company) GetDimensionValues()([]DimensionValue) {
    if m == nil {
        return nil
    } else {
        return m.dimensionValues
    }
}
func (m *Company) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Company) GetEmployees()([]Employee) {
    if m == nil {
        return nil
    } else {
        return m.employees
    }
}
func (m *Company) GetGeneralLedgerEntries()([]GeneralLedgerEntry) {
    if m == nil {
        return nil
    } else {
        return m.generalLedgerEntries
    }
}
func (m *Company) GetItemCategories()([]ItemCategory) {
    if m == nil {
        return nil
    } else {
        return m.itemCategories
    }
}
func (m *Company) GetItems()([]Item) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
func (m *Company) GetJournalLines()([]JournalLine) {
    if m == nil {
        return nil
    } else {
        return m.journalLines
    }
}
func (m *Company) GetJournals()([]Journal) {
    if m == nil {
        return nil
    } else {
        return m.journals
    }
}
func (m *Company) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *Company) GetPaymentMethods()([]PaymentMethod) {
    if m == nil {
        return nil
    } else {
        return m.paymentMethods
    }
}
func (m *Company) GetPaymentTerms()([]PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerms
    }
}
func (m *Company) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
func (m *Company) GetPurchaseInvoiceLines()([]PurchaseInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoiceLines
    }
}
func (m *Company) GetPurchaseInvoices()([]PurchaseInvoice) {
    if m == nil {
        return nil
    } else {
        return m.purchaseInvoices
    }
}
func (m *Company) GetSalesCreditMemoLines()([]SalesCreditMemoLine) {
    if m == nil {
        return nil
    } else {
        return m.salesCreditMemoLines
    }
}
func (m *Company) GetSalesCreditMemos()([]SalesCreditMemo) {
    if m == nil {
        return nil
    } else {
        return m.salesCreditMemos
    }
}
func (m *Company) GetSalesInvoiceLines()([]SalesInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.salesInvoiceLines
    }
}
func (m *Company) GetSalesInvoices()([]SalesInvoice) {
    if m == nil {
        return nil
    } else {
        return m.salesInvoices
    }
}
func (m *Company) GetSalesOrderLines()([]SalesOrderLine) {
    if m == nil {
        return nil
    } else {
        return m.salesOrderLines
    }
}
func (m *Company) GetSalesOrders()([]SalesOrder) {
    if m == nil {
        return nil
    } else {
        return m.salesOrders
    }
}
func (m *Company) GetSalesQuoteLines()([]SalesQuoteLine) {
    if m == nil {
        return nil
    } else {
        return m.salesQuoteLines
    }
}
func (m *Company) GetSalesQuotes()([]SalesQuote) {
    if m == nil {
        return nil
    } else {
        return m.salesQuotes
    }
}
func (m *Company) GetShipmentMethods()([]ShipmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethods
    }
}
func (m *Company) GetSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.systemVersion
    }
}
func (m *Company) GetTaxAreas()([]TaxArea) {
    if m == nil {
        return nil
    } else {
        return m.taxAreas
    }
}
func (m *Company) GetTaxGroups()([]TaxGroup) {
    if m == nil {
        return nil
    } else {
        return m.taxGroups
    }
}
func (m *Company) GetUnitsOfMeasure()([]UnitOfMeasure) {
    if m == nil {
        return nil
    } else {
        return m.unitsOfMeasure
    }
}
func (m *Company) GetVendors()([]Vendor) {
    if m == nil {
        return nil
    } else {
        return m.vendors
    }
}
func (m *Company) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccount() })
        if err != nil {
            return err
        }
        res := make([]Account, len(val))
        for i, v := range val {
            res[i] = *(v.(*Account))
        }
        m.SetAccounts(res)
        return nil
    }
    res["agedAccountsPayable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgedAccountsPayable() })
        if err != nil {
            return err
        }
        res := make([]AgedAccountsPayable, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgedAccountsPayable))
        }
        m.SetAgedAccountsPayable(res)
        return nil
    }
    res["agedAccountsReceivable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgedAccountsReceivable() })
        if err != nil {
            return err
        }
        res := make([]AgedAccountsReceivable, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgedAccountsReceivable))
        }
        m.SetAgedAccountsReceivable(res)
        return nil
    }
    res["businessProfileId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBusinessProfileId(val)
        return nil
    }
    res["companyInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyInformation() })
        if err != nil {
            return err
        }
        res := make([]CompanyInformation, len(val))
        for i, v := range val {
            res[i] = *(v.(*CompanyInformation))
        }
        m.SetCompanyInformation(res)
        return nil
    }
    res["countriesRegions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCountryRegion() })
        if err != nil {
            return err
        }
        res := make([]CountryRegion, len(val))
        for i, v := range val {
            res[i] = *(v.(*CountryRegion))
        }
        m.SetCountriesRegions(res)
        return nil
    }
    res["currencies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrency() })
        if err != nil {
            return err
        }
        res := make([]Currency, len(val))
        for i, v := range val {
            res[i] = *(v.(*Currency))
        }
        m.SetCurrencies(res)
        return nil
    }
    res["customerPaymentJournals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomerPaymentJournal() })
        if err != nil {
            return err
        }
        res := make([]CustomerPaymentJournal, len(val))
        for i, v := range val {
            res[i] = *(v.(*CustomerPaymentJournal))
        }
        m.SetCustomerPaymentJournals(res)
        return nil
    }
    res["customerPayments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomerPayment() })
        if err != nil {
            return err
        }
        res := make([]CustomerPayment, len(val))
        for i, v := range val {
            res[i] = *(v.(*CustomerPayment))
        }
        m.SetCustomerPayments(res)
        return nil
    }
    res["customers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomer() })
        if err != nil {
            return err
        }
        res := make([]Customer, len(val))
        for i, v := range val {
            res[i] = *(v.(*Customer))
        }
        m.SetCustomers(res)
        return nil
    }
    res["dimensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDimension() })
        if err != nil {
            return err
        }
        res := make([]Dimension, len(val))
        for i, v := range val {
            res[i] = *(v.(*Dimension))
        }
        m.SetDimensions(res)
        return nil
    }
    res["dimensionValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDimensionValue() })
        if err != nil {
            return err
        }
        res := make([]DimensionValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*DimensionValue))
        }
        m.SetDimensionValues(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["employees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmployee() })
        if err != nil {
            return err
        }
        res := make([]Employee, len(val))
        for i, v := range val {
            res[i] = *(v.(*Employee))
        }
        m.SetEmployees(res)
        return nil
    }
    res["generalLedgerEntries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeneralLedgerEntry() })
        if err != nil {
            return err
        }
        res := make([]GeneralLedgerEntry, len(val))
        for i, v := range val {
            res[i] = *(v.(*GeneralLedgerEntry))
        }
        m.SetGeneralLedgerEntries(res)
        return nil
    }
    res["itemCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemCategory() })
        if err != nil {
            return err
        }
        res := make([]ItemCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemCategory))
        }
        m.SetItemCategories(res)
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItem() })
        if err != nil {
            return err
        }
        res := make([]Item, len(val))
        for i, v := range val {
            res[i] = *(v.(*Item))
        }
        m.SetItems(res)
        return nil
    }
    res["journalLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJournalLine() })
        if err != nil {
            return err
        }
        res := make([]JournalLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*JournalLine))
        }
        m.SetJournalLines(res)
        return nil
    }
    res["journals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJournal() })
        if err != nil {
            return err
        }
        res := make([]Journal, len(val))
        for i, v := range val {
            res[i] = *(v.(*Journal))
        }
        m.SetJournals(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["paymentMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentMethod() })
        if err != nil {
            return err
        }
        res := make([]PaymentMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*PaymentMethod))
        }
        m.SetPaymentMethods(res)
        return nil
    }
    res["paymentTerms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentTerm() })
        if err != nil {
            return err
        }
        res := make([]PaymentTerm, len(val))
        for i, v := range val {
            res[i] = *(v.(*PaymentTerm))
        }
        m.SetPaymentTerms(res)
        return nil
    }
    res["picture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPicture() })
        if err != nil {
            return err
        }
        res := make([]Picture, len(val))
        for i, v := range val {
            res[i] = *(v.(*Picture))
        }
        m.SetPicture(res)
        return nil
    }
    res["purchaseInvoiceLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPurchaseInvoiceLine() })
        if err != nil {
            return err
        }
        res := make([]PurchaseInvoiceLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*PurchaseInvoiceLine))
        }
        m.SetPurchaseInvoiceLines(res)
        return nil
    }
    res["purchaseInvoices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPurchaseInvoice() })
        if err != nil {
            return err
        }
        res := make([]PurchaseInvoice, len(val))
        for i, v := range val {
            res[i] = *(v.(*PurchaseInvoice))
        }
        m.SetPurchaseInvoices(res)
        return nil
    }
    res["salesCreditMemoLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesCreditMemoLine() })
        if err != nil {
            return err
        }
        res := make([]SalesCreditMemoLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesCreditMemoLine))
        }
        m.SetSalesCreditMemoLines(res)
        return nil
    }
    res["salesCreditMemos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesCreditMemo() })
        if err != nil {
            return err
        }
        res := make([]SalesCreditMemo, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesCreditMemo))
        }
        m.SetSalesCreditMemos(res)
        return nil
    }
    res["salesInvoiceLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesInvoiceLine() })
        if err != nil {
            return err
        }
        res := make([]SalesInvoiceLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesInvoiceLine))
        }
        m.SetSalesInvoiceLines(res)
        return nil
    }
    res["salesInvoices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesInvoice() })
        if err != nil {
            return err
        }
        res := make([]SalesInvoice, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesInvoice))
        }
        m.SetSalesInvoices(res)
        return nil
    }
    res["salesOrderLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesOrderLine() })
        if err != nil {
            return err
        }
        res := make([]SalesOrderLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesOrderLine))
        }
        m.SetSalesOrderLines(res)
        return nil
    }
    res["salesOrders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesOrder() })
        if err != nil {
            return err
        }
        res := make([]SalesOrder, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesOrder))
        }
        m.SetSalesOrders(res)
        return nil
    }
    res["salesQuoteLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesQuoteLine() })
        if err != nil {
            return err
        }
        res := make([]SalesQuoteLine, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesQuoteLine))
        }
        m.SetSalesQuoteLines(res)
        return nil
    }
    res["salesQuotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesQuote() })
        if err != nil {
            return err
        }
        res := make([]SalesQuote, len(val))
        for i, v := range val {
            res[i] = *(v.(*SalesQuote))
        }
        m.SetSalesQuotes(res)
        return nil
    }
    res["shipmentMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShipmentMethod() })
        if err != nil {
            return err
        }
        res := make([]ShipmentMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*ShipmentMethod))
        }
        m.SetShipmentMethods(res)
        return nil
    }
    res["systemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSystemVersion(val)
        return nil
    }
    res["taxAreas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTaxArea() })
        if err != nil {
            return err
        }
        res := make([]TaxArea, len(val))
        for i, v := range val {
            res[i] = *(v.(*TaxArea))
        }
        m.SetTaxAreas(res)
        return nil
    }
    res["taxGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTaxGroup() })
        if err != nil {
            return err
        }
        res := make([]TaxGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*TaxGroup))
        }
        m.SetTaxGroups(res)
        return nil
    }
    res["unitsOfMeasure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnitOfMeasure() })
        if err != nil {
            return err
        }
        res := make([]UnitOfMeasure, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnitOfMeasure))
        }
        m.SetUnitsOfMeasure(res)
        return nil
    }
    res["vendors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVendor() })
        if err != nil {
            return err
        }
        res := make([]Vendor, len(val))
        for i, v := range val {
            res[i] = *(v.(*Vendor))
        }
        m.SetVendors(res)
        return nil
    }
    return res
}
func (m *Company) IsNil()(bool) {
    return m == nil
}
func (m *Company) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
func (m *Company) SetAccounts(value []Account)() {
    m.accounts = value
}
func (m *Company) SetAgedAccountsPayable(value []AgedAccountsPayable)() {
    m.agedAccountsPayable = value
}
func (m *Company) SetAgedAccountsReceivable(value []AgedAccountsReceivable)() {
    m.agedAccountsReceivable = value
}
func (m *Company) SetBusinessProfileId(value *string)() {
    m.businessProfileId = value
}
func (m *Company) SetCompanyInformation(value []CompanyInformation)() {
    m.companyInformation = value
}
func (m *Company) SetCountriesRegions(value []CountryRegion)() {
    m.countriesRegions = value
}
func (m *Company) SetCurrencies(value []Currency)() {
    m.currencies = value
}
func (m *Company) SetCustomerPaymentJournals(value []CustomerPaymentJournal)() {
    m.customerPaymentJournals = value
}
func (m *Company) SetCustomerPayments(value []CustomerPayment)() {
    m.customerPayments = value
}
func (m *Company) SetCustomers(value []Customer)() {
    m.customers = value
}
func (m *Company) SetDimensions(value []Dimension)() {
    m.dimensions = value
}
func (m *Company) SetDimensionValues(value []DimensionValue)() {
    m.dimensionValues = value
}
func (m *Company) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Company) SetEmployees(value []Employee)() {
    m.employees = value
}
func (m *Company) SetGeneralLedgerEntries(value []GeneralLedgerEntry)() {
    m.generalLedgerEntries = value
}
func (m *Company) SetItemCategories(value []ItemCategory)() {
    m.itemCategories = value
}
func (m *Company) SetItems(value []Item)() {
    m.items = value
}
func (m *Company) SetJournalLines(value []JournalLine)() {
    m.journalLines = value
}
func (m *Company) SetJournals(value []Journal)() {
    m.journals = value
}
func (m *Company) SetName(value *string)() {
    m.name = value
}
func (m *Company) SetPaymentMethods(value []PaymentMethod)() {
    m.paymentMethods = value
}
func (m *Company) SetPaymentTerms(value []PaymentTerm)() {
    m.paymentTerms = value
}
func (m *Company) SetPicture(value []Picture)() {
    m.picture = value
}
func (m *Company) SetPurchaseInvoiceLines(value []PurchaseInvoiceLine)() {
    m.purchaseInvoiceLines = value
}
func (m *Company) SetPurchaseInvoices(value []PurchaseInvoice)() {
    m.purchaseInvoices = value
}
func (m *Company) SetSalesCreditMemoLines(value []SalesCreditMemoLine)() {
    m.salesCreditMemoLines = value
}
func (m *Company) SetSalesCreditMemos(value []SalesCreditMemo)() {
    m.salesCreditMemos = value
}
func (m *Company) SetSalesInvoiceLines(value []SalesInvoiceLine)() {
    m.salesInvoiceLines = value
}
func (m *Company) SetSalesInvoices(value []SalesInvoice)() {
    m.salesInvoices = value
}
func (m *Company) SetSalesOrderLines(value []SalesOrderLine)() {
    m.salesOrderLines = value
}
func (m *Company) SetSalesOrders(value []SalesOrder)() {
    m.salesOrders = value
}
func (m *Company) SetSalesQuoteLines(value []SalesQuoteLine)() {
    m.salesQuoteLines = value
}
func (m *Company) SetSalesQuotes(value []SalesQuote)() {
    m.salesQuotes = value
}
func (m *Company) SetShipmentMethods(value []ShipmentMethod)() {
    m.shipmentMethods = value
}
func (m *Company) SetSystemVersion(value *string)() {
    m.systemVersion = value
}
func (m *Company) SetTaxAreas(value []TaxArea)() {
    m.taxAreas = value
}
func (m *Company) SetTaxGroups(value []TaxGroup)() {
    m.taxGroups = value
}
func (m *Company) SetUnitsOfMeasure(value []UnitOfMeasure)() {
    m.unitsOfMeasure = value
}
func (m *Company) SetVendors(value []Vendor)() {
    m.vendors = value
}
