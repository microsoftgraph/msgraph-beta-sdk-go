package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CompanyInformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCompanyInformation instantiates a new CompanyInformation and sets the default values.
func NewCompanyInformation()(*CompanyInformation) {
    m := &CompanyInformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCompanyInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompanyInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanyInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompanyInformation) GetAdditionalData()(map[string]any) {
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
// GetAddress gets the address property value. The address property
// returns a PostalAddressTypeable when successful
func (m *CompanyInformation) GetAddress()(PostalAddressTypeable) {
    val, err := m.GetBackingStore().Get("address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PostalAddressTypeable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CompanyInformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
// returns a *string when successful
func (m *CompanyInformation) GetCurrencyCode()(*string) {
    val, err := m.GetBackingStore().Get("currencyCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCurrentFiscalYearStartDate gets the currentFiscalYearStartDate property value. The currentFiscalYearStartDate property
// returns a *DateOnly when successful
func (m *CompanyInformation) GetCurrentFiscalYearStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("currentFiscalYearStartDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *CompanyInformation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *CompanyInformation) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFaxNumber gets the faxNumber property value. The faxNumber property
// returns a *string when successful
func (m *CompanyInformation) GetFaxNumber()(*string) {
    val, err := m.GetBackingStore().Get("faxNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompanyInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["currencyCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currentFiscalYearStartDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentFiscalYearStartDate(val)
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["faxNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaxNumber(val)
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
    res["industry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustry(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPicture(val)
        }
        return nil
    }
    res["taxRegistrationNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxRegistrationNumber(val)
        }
        return nil
    }
    res["website"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebsite(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *UUID when successful
func (m *CompanyInformation) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetIndustry gets the industry property value. The industry property
// returns a *string when successful
func (m *CompanyInformation) GetIndustry()(*string) {
    val, err := m.GetBackingStore().Get("industry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
// returns a *Time when successful
func (m *CompanyInformation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CompanyInformation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
// returns a *string when successful
func (m *CompanyInformation) GetPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("phoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPicture gets the picture property value. The picture property
// returns a []byte when successful
func (m *CompanyInformation) GetPicture()([]byte) {
    val, err := m.GetBackingStore().Get("picture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetTaxRegistrationNumber gets the taxRegistrationNumber property value. The taxRegistrationNumber property
// returns a *string when successful
func (m *CompanyInformation) GetTaxRegistrationNumber()(*string) {
    val, err := m.GetBackingStore().Get("taxRegistrationNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *CompanyInformation) GetWebsite()(*string) {
    val, err := m.GetBackingStore().Get("website")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CompanyInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("currentFiscalYearStartDate", m.GetCurrentFiscalYearStartDate())
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
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("faxNumber", m.GetFaxNumber())
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
    {
        err := writer.WriteStringValue("industry", m.GetIndustry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
    {
        err := writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("picture", m.GetPicture())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("taxRegistrationNumber", m.GetTaxRegistrationNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("website", m.GetWebsite())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyInformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddress sets the address property value. The address property
func (m *CompanyInformation) SetAddress(value PostalAddressTypeable)() {
    err := m.GetBackingStore().Set("address", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CompanyInformation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *CompanyInformation) SetCurrencyCode(value *string)() {
    err := m.GetBackingStore().Set("currencyCode", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrentFiscalYearStartDate sets the currentFiscalYearStartDate property value. The currentFiscalYearStartDate property
func (m *CompanyInformation) SetCurrentFiscalYearStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("currentFiscalYearStartDate", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CompanyInformation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmail sets the email property value. The email property
func (m *CompanyInformation) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetFaxNumber sets the faxNumber property value. The faxNumber property
func (m *CompanyInformation) SetFaxNumber(value *string)() {
    err := m.GetBackingStore().Set("faxNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. The id property
func (m *CompanyInformation) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetIndustry sets the industry property value. The industry property
func (m *CompanyInformation) SetIndustry(value *string)() {
    err := m.GetBackingStore().Set("industry", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CompanyInformation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CompanyInformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *CompanyInformation) SetPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("phoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetPicture sets the picture property value. The picture property
func (m *CompanyInformation) SetPicture(value []byte)() {
    err := m.GetBackingStore().Set("picture", value)
    if err != nil {
        panic(err)
    }
}
// SetTaxRegistrationNumber sets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *CompanyInformation) SetTaxRegistrationNumber(value *string)() {
    err := m.GetBackingStore().Set("taxRegistrationNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetWebsite sets the website property value. The website property
func (m *CompanyInformation) SetWebsite(value *string)() {
    err := m.GetBackingStore().Set("website", value)
    if err != nil {
        panic(err)
    }
}
type CompanyInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PostalAddressTypeable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCurrencyCode()(*string)
    GetCurrentFiscalYearStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetFaxNumber()(*string)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIndustry()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetPhoneNumber()(*string)
    GetPicture()([]byte)
    GetTaxRegistrationNumber()(*string)
    GetWebsite()(*string)
    SetAddress(value PostalAddressTypeable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCurrencyCode(value *string)()
    SetCurrentFiscalYearStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetFaxNumber(value *string)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIndustry(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetPhoneNumber(value *string)()
    SetPicture(value []byte)()
    SetTaxRegistrationNumber(value *string)()
    SetWebsite(value *string)()
}
