package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// Company 
type Company struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCompany instantiates a new company and sets the default values.
func NewCompany()(*Company) {
    m := &Company{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCompanyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompany(), nil
}
// GetAccounts gets the accounts property value. The accounts property
func (m *Company) GetAccounts()([]Accountable) {
    val, err := m.GetBackingStore().Get("accounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Accountable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Company) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAgedAccountsPayable gets the agedAccountsPayable property value. The agedAccountsPayable property
func (m *Company) GetAgedAccountsPayable()([]AgedAccountsPayableable) {
    val, err := m.GetBackingStore().Get("agedAccountsPayable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AgedAccountsPayableable)
    }
    return nil
}
// GetAgedAccountsReceivable gets the agedAccountsReceivable property value. The agedAccountsReceivable property
func (m *Company) GetAgedAccountsReceivable()([]AgedAccountsReceivableable) {
    val, err := m.GetBackingStore().Get("agedAccountsReceivable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AgedAccountsReceivableable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *Company) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBusinessProfileId gets the businessProfileId property value. The businessProfileId property
func (m *Company) GetBusinessProfileId()(*string) {
    val, err := m.GetBackingStore().Get("businessProfileId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCompanyInformation gets the companyInformation property value. The companyInformation property
func (m *Company) GetCompanyInformation()([]CompanyInformationable) {
    val, err := m.GetBackingStore().Get("companyInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CompanyInformationable)
    }
    return nil
}
// GetCountriesRegions gets the countriesRegions property value. The countriesRegions property
func (m *Company) GetCountriesRegions()([]CountryRegionable) {
    val, err := m.GetBackingStore().Get("countriesRegions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CountryRegionable)
    }
    return nil
}
// GetCurrencies gets the currencies property value. The currencies property
func (m *Company) GetCurrencies()([]Currencyable) {
    val, err := m.GetBackingStore().Get("currencies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Currencyable)
    }
    return nil
}
// GetCustomerPaymentJournals gets the customerPaymentJournals property value. The customerPaymentJournals property
func (m *Company) GetCustomerPaymentJournals()([]CustomerPaymentJournalable) {
    val, err := m.GetBackingStore().Get("customerPaymentJournals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomerPaymentJournalable)
    }
    return nil
}
// GetCustomerPayments gets the customerPayments property value. The customerPayments property
func (m *Company) GetCustomerPayments()([]CustomerPaymentable) {
    val, err := m.GetBackingStore().Get("customerPayments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomerPaymentable)
    }
    return nil
}
// GetCustomers gets the customers property value. The customers property
func (m *Company) GetCustomers()([]Customerable) {
    val, err := m.GetBackingStore().Get("customers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Customerable)
    }
    return nil
}
// GetDimensions gets the dimensions property value. The dimensions property
func (m *Company) GetDimensions()([]Dimensionable) {
    val, err := m.GetBackingStore().Get("dimensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Dimensionable)
    }
    return nil
}
// GetDimensionValues gets the dimensionValues property value. The dimensionValues property
func (m *Company) GetDimensionValues()([]DimensionValueable) {
    val, err := m.GetBackingStore().Get("dimensionValues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DimensionValueable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Company) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmployees gets the employees property value. The employees property
func (m *Company) GetEmployees()([]Employeeable) {
    val, err := m.GetBackingStore().Get("employees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Employeeable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Company) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Accountable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Accountable)
                }
            }
            m.SetAccounts(res)
        }
        return nil
    }
    res["agedAccountsPayable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgedAccountsPayableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgedAccountsPayableable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AgedAccountsPayableable)
                }
            }
            m.SetAgedAccountsPayable(res)
        }
        return nil
    }
    res["agedAccountsReceivable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgedAccountsReceivableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgedAccountsReceivableable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AgedAccountsReceivableable)
                }
            }
            m.SetAgedAccountsReceivable(res)
        }
        return nil
    }
    res["businessProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessProfileId(val)
        }
        return nil
    }
    res["companyInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompanyInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanyInformationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CompanyInformationable)
                }
            }
            m.SetCompanyInformation(res)
        }
        return nil
    }
    res["countriesRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCountryRegionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CountryRegionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CountryRegionable)
                }
            }
            m.SetCountriesRegions(res)
        }
        return nil
    }
    res["currencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCurrencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Currencyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Currencyable)
                }
            }
            m.SetCurrencies(res)
        }
        return nil
    }
    res["customerPaymentJournals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerPaymentJournalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerPaymentJournalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomerPaymentJournalable)
                }
            }
            m.SetCustomerPaymentJournals(res)
        }
        return nil
    }
    res["customerPayments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerPaymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerPaymentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomerPaymentable)
                }
            }
            m.SetCustomerPayments(res)
        }
        return nil
    }
    res["customers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Customerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Customerable)
                }
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["dimensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDimensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Dimensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Dimensionable)
                }
            }
            m.SetDimensions(res)
        }
        return nil
    }
    res["dimensionValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDimensionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DimensionValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DimensionValueable)
                }
            }
            m.SetDimensionValues(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["employees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmployeeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Employeeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Employeeable)
                }
            }
            m.SetEmployees(res)
        }
        return nil
    }
    res["generalLedgerEntries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGeneralLedgerEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GeneralLedgerEntryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GeneralLedgerEntryable)
                }
            }
            m.SetGeneralLedgerEntries(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["itemCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemCategoryable)
                }
            }
            m.SetItemCategories(res)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Itemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Itemable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    res["journalLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJournalLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JournalLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JournalLineable)
                }
            }
            m.SetJournalLines(res)
        }
        return nil
    }
    res["journals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJournalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Journalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Journalable)
                }
            }
            m.SetJournals(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["paymentMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePaymentMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PaymentMethodable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PaymentMethodable)
                }
            }
            m.SetPaymentMethods(res)
        }
        return nil
    }
    res["paymentTerms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePaymentTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PaymentTermable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PaymentTermable)
                }
            }
            m.SetPaymentTerms(res)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pictureable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Pictureable)
                }
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["purchaseInvoiceLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePurchaseInvoiceLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PurchaseInvoiceLineable)
                }
            }
            m.SetPurchaseInvoiceLines(res)
        }
        return nil
    }
    res["purchaseInvoices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePurchaseInvoiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PurchaseInvoiceable)
                }
            }
            m.SetPurchaseInvoices(res)
        }
        return nil
    }
    res["salesCreditMemoLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesCreditMemoLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesCreditMemoLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesCreditMemoLineable)
                }
            }
            m.SetSalesCreditMemoLines(res)
        }
        return nil
    }
    res["salesCreditMemos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesCreditMemoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesCreditMemoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesCreditMemoable)
                }
            }
            m.SetSalesCreditMemos(res)
        }
        return nil
    }
    res["salesInvoiceLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesInvoiceLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesInvoiceLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesInvoiceLineable)
                }
            }
            m.SetSalesInvoiceLines(res)
        }
        return nil
    }
    res["salesInvoices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesInvoiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesInvoiceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesInvoiceable)
                }
            }
            m.SetSalesInvoices(res)
        }
        return nil
    }
    res["salesOrderLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesOrderLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesOrderLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesOrderLineable)
                }
            }
            m.SetSalesOrderLines(res)
        }
        return nil
    }
    res["salesOrders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesOrderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesOrderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesOrderable)
                }
            }
            m.SetSalesOrders(res)
        }
        return nil
    }
    res["salesQuoteLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesQuoteLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuoteLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesQuoteLineable)
                }
            }
            m.SetSalesQuoteLines(res)
        }
        return nil
    }
    res["salesQuotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesQuoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuoteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SalesQuoteable)
                }
            }
            m.SetSalesQuotes(res)
        }
        return nil
    }
    res["shipmentMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateShipmentMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ShipmentMethodable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ShipmentMethodable)
                }
            }
            m.SetShipmentMethods(res)
        }
        return nil
    }
    res["systemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemVersion(val)
        }
        return nil
    }
    res["taxAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaxAreaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaxAreaable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TaxAreaable)
                }
            }
            m.SetTaxAreas(res)
        }
        return nil
    }
    res["taxGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaxGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaxGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TaxGroupable)
                }
            }
            m.SetTaxGroups(res)
        }
        return nil
    }
    res["unitsOfMeasure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnitOfMeasureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnitOfMeasureable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnitOfMeasureable)
                }
            }
            m.SetUnitsOfMeasure(res)
        }
        return nil
    }
    res["vendors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVendorEscapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VendorEscapedable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VendorEscapedable)
                }
            }
            m.SetVendors(res)
        }
        return nil
    }
    return res
}
// GetGeneralLedgerEntries gets the generalLedgerEntries property value. The generalLedgerEntries property
func (m *Company) GetGeneralLedgerEntries()([]GeneralLedgerEntryable) {
    val, err := m.GetBackingStore().Get("generalLedgerEntries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GeneralLedgerEntryable)
    }
    return nil
}
// GetId gets the id property value. The id property
func (m *Company) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetItemCategories gets the itemCategories property value. The itemCategories property
func (m *Company) GetItemCategories()([]ItemCategoryable) {
    val, err := m.GetBackingStore().Get("itemCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ItemCategoryable)
    }
    return nil
}
// GetItems gets the items property value. The items property
func (m *Company) GetItems()([]Itemable) {
    val, err := m.GetBackingStore().Get("items")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Itemable)
    }
    return nil
}
// GetJournalLines gets the journalLines property value. The journalLines property
func (m *Company) GetJournalLines()([]JournalLineable) {
    val, err := m.GetBackingStore().Get("journalLines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]JournalLineable)
    }
    return nil
}
// GetJournals gets the journals property value. The journals property
func (m *Company) GetJournals()([]Journalable) {
    val, err := m.GetBackingStore().Get("journals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Journalable)
    }
    return nil
}
// GetName gets the name property value. The name property
func (m *Company) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Company) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPaymentMethods gets the paymentMethods property value. The paymentMethods property
func (m *Company) GetPaymentMethods()([]PaymentMethodable) {
    val, err := m.GetBackingStore().Get("paymentMethods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PaymentMethodable)
    }
    return nil
}
// GetPaymentTerms gets the paymentTerms property value. The paymentTerms property
func (m *Company) GetPaymentTerms()([]PaymentTermable) {
    val, err := m.GetBackingStore().Get("paymentTerms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PaymentTermable)
    }
    return nil
}
// GetPicture gets the picture property value. The picture property
func (m *Company) GetPicture()([]Pictureable) {
    val, err := m.GetBackingStore().Get("picture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Pictureable)
    }
    return nil
}
// GetPurchaseInvoiceLines gets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *Company) GetPurchaseInvoiceLines()([]PurchaseInvoiceLineable) {
    val, err := m.GetBackingStore().Get("purchaseInvoiceLines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PurchaseInvoiceLineable)
    }
    return nil
}
// GetPurchaseInvoices gets the purchaseInvoices property value. The purchaseInvoices property
func (m *Company) GetPurchaseInvoices()([]PurchaseInvoiceable) {
    val, err := m.GetBackingStore().Get("purchaseInvoices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PurchaseInvoiceable)
    }
    return nil
}
// GetSalesCreditMemoLines gets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *Company) GetSalesCreditMemoLines()([]SalesCreditMemoLineable) {
    val, err := m.GetBackingStore().Get("salesCreditMemoLines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesCreditMemoLineable)
    }
    return nil
}
// GetSalesCreditMemos gets the salesCreditMemos property value. The salesCreditMemos property
func (m *Company) GetSalesCreditMemos()([]SalesCreditMemoable) {
    val, err := m.GetBackingStore().Get("salesCreditMemos")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesCreditMemoable)
    }
    return nil
}
// GetSalesInvoiceLines gets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *Company) GetSalesInvoiceLines()([]SalesInvoiceLineable) {
    val, err := m.GetBackingStore().Get("salesInvoiceLines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesInvoiceLineable)
    }
    return nil
}
// GetSalesInvoices gets the salesInvoices property value. The salesInvoices property
func (m *Company) GetSalesInvoices()([]SalesInvoiceable) {
    val, err := m.GetBackingStore().Get("salesInvoices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesInvoiceable)
    }
    return nil
}
// GetSalesOrderLines gets the salesOrderLines property value. The salesOrderLines property
func (m *Company) GetSalesOrderLines()([]SalesOrderLineable) {
    val, err := m.GetBackingStore().Get("salesOrderLines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesOrderLineable)
    }
    return nil
}
// GetSalesOrders gets the salesOrders property value. The salesOrders property
func (m *Company) GetSalesOrders()([]SalesOrderable) {
    val, err := m.GetBackingStore().Get("salesOrders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesOrderable)
    }
    return nil
}
// GetSalesQuoteLines gets the salesQuoteLines property value. The salesQuoteLines property
func (m *Company) GetSalesQuoteLines()([]SalesQuoteLineable) {
    val, err := m.GetBackingStore().Get("salesQuoteLines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesQuoteLineable)
    }
    return nil
}
// GetSalesQuotes gets the salesQuotes property value. The salesQuotes property
func (m *Company) GetSalesQuotes()([]SalesQuoteable) {
    val, err := m.GetBackingStore().Get("salesQuotes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SalesQuoteable)
    }
    return nil
}
// GetShipmentMethods gets the shipmentMethods property value. The shipmentMethods property
func (m *Company) GetShipmentMethods()([]ShipmentMethodable) {
    val, err := m.GetBackingStore().Get("shipmentMethods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ShipmentMethodable)
    }
    return nil
}
// GetSystemVersion gets the systemVersion property value. The systemVersion property
func (m *Company) GetSystemVersion()(*string) {
    val, err := m.GetBackingStore().Get("systemVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTaxAreas gets the taxAreas property value. The taxAreas property
func (m *Company) GetTaxAreas()([]TaxAreaable) {
    val, err := m.GetBackingStore().Get("taxAreas")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TaxAreaable)
    }
    return nil
}
// GetTaxGroups gets the taxGroups property value. The taxGroups property
func (m *Company) GetTaxGroups()([]TaxGroupable) {
    val, err := m.GetBackingStore().Get("taxGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TaxGroupable)
    }
    return nil
}
// GetUnitsOfMeasure gets the unitsOfMeasure property value. The unitsOfMeasure property
func (m *Company) GetUnitsOfMeasure()([]UnitOfMeasureable) {
    val, err := m.GetBackingStore().Get("unitsOfMeasure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnitOfMeasureable)
    }
    return nil
}
// GetVendors gets the vendors property value. The vendors property
func (m *Company) GetVendors()([]VendorEscapedable) {
    val, err := m.GetBackingStore().Get("vendors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VendorEscapedable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Company) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("accounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsPayable() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgedAccountsPayable()))
        for i, v := range m.GetAgedAccountsPayable() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("agedAccountsPayable", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsReceivable() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgedAccountsReceivable()))
        for i, v := range m.GetAgedAccountsReceivable() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("agedAccountsReceivable", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("businessProfileId", m.GetBusinessProfileId())
        if err != nil {
            return err
        }
    }
    if m.GetCompanyInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyInformation()))
        for i, v := range m.GetCompanyInformation() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("companyInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCountriesRegions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCountriesRegions()))
        for i, v := range m.GetCountriesRegions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("countriesRegions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCurrencies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCurrencies()))
        for i, v := range m.GetCurrencies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("currencies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPaymentJournals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomerPaymentJournals()))
        for i, v := range m.GetCustomerPaymentJournals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("customerPaymentJournals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPayments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomerPayments()))
        for i, v := range m.GetCustomerPayments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("customerPayments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDimensions()))
        for i, v := range m.GetDimensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("dimensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensionValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDimensionValues()))
        for i, v := range m.GetDimensionValues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("dimensionValues", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEmployees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmployees()))
        for i, v := range m.GetEmployees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("employees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGeneralLedgerEntries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGeneralLedgerEntries()))
        for i, v := range m.GetGeneralLedgerEntries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("generalLedgerEntries", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetItemCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemCategories()))
        for i, v := range m.GetItemCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("itemCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournalLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJournalLines()))
        for i, v := range m.GetJournalLines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("journalLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJournals()))
        for i, v := range m.GetJournals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("journals", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPaymentMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPaymentMethods()))
        for i, v := range m.GetPaymentMethods() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("paymentMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPaymentTerms() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPaymentTerms()))
        for i, v := range m.GetPaymentTerms() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("paymentTerms", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPicture() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoiceLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPurchaseInvoiceLines()))
        for i, v := range m.GetPurchaseInvoiceLines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("purchaseInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPurchaseInvoices()))
        for i, v := range m.GetPurchaseInvoices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("purchaseInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemoLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesCreditMemoLines()))
        for i, v := range m.GetSalesCreditMemoLines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesCreditMemoLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemos() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesCreditMemos()))
        for i, v := range m.GetSalesCreditMemos() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesCreditMemos", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoiceLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesInvoiceLines()))
        for i, v := range m.GetSalesInvoiceLines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesInvoices()))
        for i, v := range m.GetSalesInvoices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrderLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesOrderLines()))
        for i, v := range m.GetSalesOrderLines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesOrderLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesOrders()))
        for i, v := range m.GetSalesOrders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesOrders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuoteLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesQuoteLines()))
        for i, v := range m.GetSalesQuoteLines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesQuoteLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesQuotes()))
        for i, v := range m.GetSalesQuotes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("salesQuotes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetShipmentMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShipmentMethods()))
        for i, v := range m.GetShipmentMethods() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("shipmentMethods", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("systemVersion", m.GetSystemVersion())
        if err != nil {
            return err
        }
    }
    if m.GetTaxAreas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaxAreas()))
        for i, v := range m.GetTaxAreas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("taxAreas", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaxGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaxGroups()))
        for i, v := range m.GetTaxGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("taxGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUnitsOfMeasure() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnitsOfMeasure()))
        for i, v := range m.GetUnitsOfMeasure() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("unitsOfMeasure", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVendors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVendors()))
        for i, v := range m.GetVendors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("vendors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. The accounts property
func (m *Company) SetAccounts(value []Accountable)() {
    err := m.GetBackingStore().Set("accounts", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Company) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAgedAccountsPayable sets the agedAccountsPayable property value. The agedAccountsPayable property
func (m *Company) SetAgedAccountsPayable(value []AgedAccountsPayableable)() {
    err := m.GetBackingStore().Set("agedAccountsPayable", value)
    if err != nil {
        panic(err)
    }
}
// SetAgedAccountsReceivable sets the agedAccountsReceivable property value. The agedAccountsReceivable property
func (m *Company) SetAgedAccountsReceivable(value []AgedAccountsReceivableable)() {
    err := m.GetBackingStore().Set("agedAccountsReceivable", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *Company) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBusinessProfileId sets the businessProfileId property value. The businessProfileId property
func (m *Company) SetBusinessProfileId(value *string)() {
    err := m.GetBackingStore().Set("businessProfileId", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyInformation sets the companyInformation property value. The companyInformation property
func (m *Company) SetCompanyInformation(value []CompanyInformationable)() {
    err := m.GetBackingStore().Set("companyInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetCountriesRegions sets the countriesRegions property value. The countriesRegions property
func (m *Company) SetCountriesRegions(value []CountryRegionable)() {
    err := m.GetBackingStore().Set("countriesRegions", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrencies sets the currencies property value. The currencies property
func (m *Company) SetCurrencies(value []Currencyable)() {
    err := m.GetBackingStore().Set("currencies", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerPaymentJournals sets the customerPaymentJournals property value. The customerPaymentJournals property
func (m *Company) SetCustomerPaymentJournals(value []CustomerPaymentJournalable)() {
    err := m.GetBackingStore().Set("customerPaymentJournals", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerPayments sets the customerPayments property value. The customerPayments property
func (m *Company) SetCustomerPayments(value []CustomerPaymentable)() {
    err := m.GetBackingStore().Set("customerPayments", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomers sets the customers property value. The customers property
func (m *Company) SetCustomers(value []Customerable)() {
    err := m.GetBackingStore().Set("customers", value)
    if err != nil {
        panic(err)
    }
}
// SetDimensions sets the dimensions property value. The dimensions property
func (m *Company) SetDimensions(value []Dimensionable)() {
    err := m.GetBackingStore().Set("dimensions", value)
    if err != nil {
        panic(err)
    }
}
// SetDimensionValues sets the dimensionValues property value. The dimensionValues property
func (m *Company) SetDimensionValues(value []DimensionValueable)() {
    err := m.GetBackingStore().Set("dimensionValues", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Company) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmployees sets the employees property value. The employees property
func (m *Company) SetEmployees(value []Employeeable)() {
    err := m.GetBackingStore().Set("employees", value)
    if err != nil {
        panic(err)
    }
}
// SetGeneralLedgerEntries sets the generalLedgerEntries property value. The generalLedgerEntries property
func (m *Company) SetGeneralLedgerEntries(value []GeneralLedgerEntryable)() {
    err := m.GetBackingStore().Set("generalLedgerEntries", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. The id property
func (m *Company) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetItemCategories sets the itemCategories property value. The itemCategories property
func (m *Company) SetItemCategories(value []ItemCategoryable)() {
    err := m.GetBackingStore().Set("itemCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetItems sets the items property value. The items property
func (m *Company) SetItems(value []Itemable)() {
    err := m.GetBackingStore().Set("items", value)
    if err != nil {
        panic(err)
    }
}
// SetJournalLines sets the journalLines property value. The journalLines property
func (m *Company) SetJournalLines(value []JournalLineable)() {
    err := m.GetBackingStore().Set("journalLines", value)
    if err != nil {
        panic(err)
    }
}
// SetJournals sets the journals property value. The journals property
func (m *Company) SetJournals(value []Journalable)() {
    err := m.GetBackingStore().Set("journals", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *Company) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Company) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPaymentMethods sets the paymentMethods property value. The paymentMethods property
func (m *Company) SetPaymentMethods(value []PaymentMethodable)() {
    err := m.GetBackingStore().Set("paymentMethods", value)
    if err != nil {
        panic(err)
    }
}
// SetPaymentTerms sets the paymentTerms property value. The paymentTerms property
func (m *Company) SetPaymentTerms(value []PaymentTermable)() {
    err := m.GetBackingStore().Set("paymentTerms", value)
    if err != nil {
        panic(err)
    }
}
// SetPicture sets the picture property value. The picture property
func (m *Company) SetPicture(value []Pictureable)() {
    err := m.GetBackingStore().Set("picture", value)
    if err != nil {
        panic(err)
    }
}
// SetPurchaseInvoiceLines sets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *Company) SetPurchaseInvoiceLines(value []PurchaseInvoiceLineable)() {
    err := m.GetBackingStore().Set("purchaseInvoiceLines", value)
    if err != nil {
        panic(err)
    }
}
// SetPurchaseInvoices sets the purchaseInvoices property value. The purchaseInvoices property
func (m *Company) SetPurchaseInvoices(value []PurchaseInvoiceable)() {
    err := m.GetBackingStore().Set("purchaseInvoices", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesCreditMemoLines sets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *Company) SetSalesCreditMemoLines(value []SalesCreditMemoLineable)() {
    err := m.GetBackingStore().Set("salesCreditMemoLines", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesCreditMemos sets the salesCreditMemos property value. The salesCreditMemos property
func (m *Company) SetSalesCreditMemos(value []SalesCreditMemoable)() {
    err := m.GetBackingStore().Set("salesCreditMemos", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesInvoiceLines sets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *Company) SetSalesInvoiceLines(value []SalesInvoiceLineable)() {
    err := m.GetBackingStore().Set("salesInvoiceLines", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesInvoices sets the salesInvoices property value. The salesInvoices property
func (m *Company) SetSalesInvoices(value []SalesInvoiceable)() {
    err := m.GetBackingStore().Set("salesInvoices", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesOrderLines sets the salesOrderLines property value. The salesOrderLines property
func (m *Company) SetSalesOrderLines(value []SalesOrderLineable)() {
    err := m.GetBackingStore().Set("salesOrderLines", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesOrders sets the salesOrders property value. The salesOrders property
func (m *Company) SetSalesOrders(value []SalesOrderable)() {
    err := m.GetBackingStore().Set("salesOrders", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesQuoteLines sets the salesQuoteLines property value. The salesQuoteLines property
func (m *Company) SetSalesQuoteLines(value []SalesQuoteLineable)() {
    err := m.GetBackingStore().Set("salesQuoteLines", value)
    if err != nil {
        panic(err)
    }
}
// SetSalesQuotes sets the salesQuotes property value. The salesQuotes property
func (m *Company) SetSalesQuotes(value []SalesQuoteable)() {
    err := m.GetBackingStore().Set("salesQuotes", value)
    if err != nil {
        panic(err)
    }
}
// SetShipmentMethods sets the shipmentMethods property value. The shipmentMethods property
func (m *Company) SetShipmentMethods(value []ShipmentMethodable)() {
    err := m.GetBackingStore().Set("shipmentMethods", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemVersion sets the systemVersion property value. The systemVersion property
func (m *Company) SetSystemVersion(value *string)() {
    err := m.GetBackingStore().Set("systemVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxAreas sets the taxAreas property value. The taxAreas property
func (m *Company) SetTaxAreas(value []TaxAreaable)() {
    err := m.GetBackingStore().Set("taxAreas", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxGroups sets the taxGroups property value. The taxGroups property
func (m *Company) SetTaxGroups(value []TaxGroupable)() {
    err := m.GetBackingStore().Set("taxGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetUnitsOfMeasure sets the unitsOfMeasure property value. The unitsOfMeasure property
func (m *Company) SetUnitsOfMeasure(value []UnitOfMeasureable)() {
    err := m.GetBackingStore().Set("unitsOfMeasure", value)
    if err != nil {
        panic(err)
    }
}
// SetVendors sets the vendors property value. The vendors property
func (m *Company) SetVendors(value []VendorEscapedable)() {
    err := m.GetBackingStore().Set("vendors", value)
    if err != nil {
        panic(err)
    }
}
// Companyable 
type Companyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()([]Accountable)
    GetAgedAccountsPayable()([]AgedAccountsPayableable)
    GetAgedAccountsReceivable()([]AgedAccountsReceivableable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBusinessProfileId()(*string)
    GetCompanyInformation()([]CompanyInformationable)
    GetCountriesRegions()([]CountryRegionable)
    GetCurrencies()([]Currencyable)
    GetCustomerPaymentJournals()([]CustomerPaymentJournalable)
    GetCustomerPayments()([]CustomerPaymentable)
    GetCustomers()([]Customerable)
    GetDimensions()([]Dimensionable)
    GetDimensionValues()([]DimensionValueable)
    GetDisplayName()(*string)
    GetEmployees()([]Employeeable)
    GetGeneralLedgerEntries()([]GeneralLedgerEntryable)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetItemCategories()([]ItemCategoryable)
    GetItems()([]Itemable)
    GetJournalLines()([]JournalLineable)
    GetJournals()([]Journalable)
    GetName()(*string)
    GetOdataType()(*string)
    GetPaymentMethods()([]PaymentMethodable)
    GetPaymentTerms()([]PaymentTermable)
    GetPicture()([]Pictureable)
    GetPurchaseInvoiceLines()([]PurchaseInvoiceLineable)
    GetPurchaseInvoices()([]PurchaseInvoiceable)
    GetSalesCreditMemoLines()([]SalesCreditMemoLineable)
    GetSalesCreditMemos()([]SalesCreditMemoable)
    GetSalesInvoiceLines()([]SalesInvoiceLineable)
    GetSalesInvoices()([]SalesInvoiceable)
    GetSalesOrderLines()([]SalesOrderLineable)
    GetSalesOrders()([]SalesOrderable)
    GetSalesQuoteLines()([]SalesQuoteLineable)
    GetSalesQuotes()([]SalesQuoteable)
    GetShipmentMethods()([]ShipmentMethodable)
    GetSystemVersion()(*string)
    GetTaxAreas()([]TaxAreaable)
    GetTaxGroups()([]TaxGroupable)
    GetUnitsOfMeasure()([]UnitOfMeasureable)
    GetVendors()([]VendorEscapedable)
    SetAccounts(value []Accountable)()
    SetAgedAccountsPayable(value []AgedAccountsPayableable)()
    SetAgedAccountsReceivable(value []AgedAccountsReceivableable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBusinessProfileId(value *string)()
    SetCompanyInformation(value []CompanyInformationable)()
    SetCountriesRegions(value []CountryRegionable)()
    SetCurrencies(value []Currencyable)()
    SetCustomerPaymentJournals(value []CustomerPaymentJournalable)()
    SetCustomerPayments(value []CustomerPaymentable)()
    SetCustomers(value []Customerable)()
    SetDimensions(value []Dimensionable)()
    SetDimensionValues(value []DimensionValueable)()
    SetDisplayName(value *string)()
    SetEmployees(value []Employeeable)()
    SetGeneralLedgerEntries(value []GeneralLedgerEntryable)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetItemCategories(value []ItemCategoryable)()
    SetItems(value []Itemable)()
    SetJournalLines(value []JournalLineable)()
    SetJournals(value []Journalable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPaymentMethods(value []PaymentMethodable)()
    SetPaymentTerms(value []PaymentTermable)()
    SetPicture(value []Pictureable)()
    SetPurchaseInvoiceLines(value []PurchaseInvoiceLineable)()
    SetPurchaseInvoices(value []PurchaseInvoiceable)()
    SetSalesCreditMemoLines(value []SalesCreditMemoLineable)()
    SetSalesCreditMemos(value []SalesCreditMemoable)()
    SetSalesInvoiceLines(value []SalesInvoiceLineable)()
    SetSalesInvoices(value []SalesInvoiceable)()
    SetSalesOrderLines(value []SalesOrderLineable)()
    SetSalesOrders(value []SalesOrderable)()
    SetSalesQuoteLines(value []SalesQuoteLineable)()
    SetSalesQuotes(value []SalesQuoteable)()
    SetShipmentMethods(value []ShipmentMethodable)()
    SetSystemVersion(value *string)()
    SetTaxAreas(value []TaxAreaable)()
    SetTaxGroups(value []TaxGroupable)()
    SetUnitsOfMeasure(value []UnitOfMeasureable)()
    SetVendors(value []VendorEscapedable)()
}
