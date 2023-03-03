package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerPayment 
type CustomerPayment struct {
    Entity
}
// NewCustomerPayment instantiates a new customerPayment and sets the default values.
func NewCustomerPayment()(*CustomerPayment) {
    m := &CustomerPayment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomerPaymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerPaymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomerPayment(), nil
}
// GetAmount gets the amount property value. The amount property
func (m *CustomerPayment) GetAmount()(*float64) {
    val, err := m.GetBackingStore().Get("amount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetAppliesToInvoiceId gets the appliesToInvoiceId property value. The appliesToInvoiceId property
func (m *CustomerPayment) GetAppliesToInvoiceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("appliesToInvoiceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetAppliesToInvoiceNumber gets the appliesToInvoiceNumber property value. The appliesToInvoiceNumber property
func (m *CustomerPayment) GetAppliesToInvoiceNumber()(*string) {
    val, err := m.GetBackingStore().Get("appliesToInvoiceNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetComment gets the comment property value. The comment property
func (m *CustomerPayment) GetComment()(*string) {
    val, err := m.GetBackingStore().Get("comment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContactId gets the contactId property value. The contactId property
func (m *CustomerPayment) GetContactId()(*string) {
    val, err := m.GetBackingStore().Get("contactId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomer gets the customer property value. The customer property
func (m *CustomerPayment) GetCustomer()(Customerable) {
    val, err := m.GetBackingStore().Get("customer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Customerable)
    }
    return nil
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *CustomerPayment) GetCustomerId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("customerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *CustomerPayment) GetCustomerNumber()(*string) {
    val, err := m.GetBackingStore().Get("customerNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *CustomerPayment) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDocumentNumber gets the documentNumber property value. The documentNumber property
func (m *CustomerPayment) GetDocumentNumber()(*string) {
    val, err := m.GetBackingStore().Get("documentNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *CustomerPayment) GetExternalDocumentNumber()(*string) {
    val, err := m.GetBackingStore().Get("externalDocumentNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomerPayment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmount(val)
        }
        return nil
    }
    res["appliesToInvoiceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesToInvoiceId(val)
        }
        return nil
    }
    res["appliesToInvoiceNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesToInvoiceNumber(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["contactId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactId(val)
        }
        return nil
    }
    res["customer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(Customerable))
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["customerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNumber(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["documentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentNumber(val)
        }
        return nil
    }
    res["externalDocumentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDocumentNumber(val)
        }
        return nil
    }
    res["journalDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJournalDisplayName(val)
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
    res["lineNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLineNumber(val)
        }
        return nil
    }
    res["postingDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostingDate(val)
        }
        return nil
    }
    return res
}
// GetJournalDisplayName gets the journalDisplayName property value. The journalDisplayName property
func (m *CustomerPayment) GetJournalDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("journalDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPayment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLineNumber gets the lineNumber property value. The lineNumber property
func (m *CustomerPayment) GetLineNumber()(*int32) {
    val, err := m.GetBackingStore().Get("lineNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPostingDate gets the postingDate property value. The postingDate property
func (m *CustomerPayment) GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("postingDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomerPayment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("amount", m.GetAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("appliesToInvoiceId", m.GetAppliesToInvoiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appliesToInvoiceNumber", m.GetAppliesToInvoiceNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactId", m.GetContactId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customer", m.GetCustomer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerNumber", m.GetCustomerNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("documentNumber", m.GetDocumentNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalDocumentNumber", m.GetExternalDocumentNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("journalDisplayName", m.GetJournalDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("lineNumber", m.GetLineNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("postingDate", m.GetPostingDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAmount sets the amount property value. The amount property
func (m *CustomerPayment) SetAmount(value *float64)() {
    err := m.GetBackingStore().Set("amount", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliesToInvoiceId sets the appliesToInvoiceId property value. The appliesToInvoiceId property
func (m *CustomerPayment) SetAppliesToInvoiceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("appliesToInvoiceId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliesToInvoiceNumber sets the appliesToInvoiceNumber property value. The appliesToInvoiceNumber property
func (m *CustomerPayment) SetAppliesToInvoiceNumber(value *string)() {
    err := m.GetBackingStore().Set("appliesToInvoiceNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetComment sets the comment property value. The comment property
func (m *CustomerPayment) SetComment(value *string)() {
    err := m.GetBackingStore().Set("comment", value)
    if err != nil {
        panic(err)
    }
}
// SetContactId sets the contactId property value. The contactId property
func (m *CustomerPayment) SetContactId(value *string)() {
    err := m.GetBackingStore().Set("contactId", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomer sets the customer property value. The customer property
func (m *CustomerPayment) SetCustomer(value Customerable)() {
    err := m.GetBackingStore().Set("customer", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *CustomerPayment) SetCustomerId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("customerId", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *CustomerPayment) SetCustomerNumber(value *string)() {
    err := m.GetBackingStore().Set("customerNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *CustomerPayment) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentNumber sets the documentNumber property value. The documentNumber property
func (m *CustomerPayment) SetDocumentNumber(value *string)() {
    err := m.GetBackingStore().Set("documentNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *CustomerPayment) SetExternalDocumentNumber(value *string)() {
    err := m.GetBackingStore().Set("externalDocumentNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetJournalDisplayName sets the journalDisplayName property value. The journalDisplayName property
func (m *CustomerPayment) SetJournalDisplayName(value *string)() {
    err := m.GetBackingStore().Set("journalDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPayment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLineNumber sets the lineNumber property value. The lineNumber property
func (m *CustomerPayment) SetLineNumber(value *int32)() {
    err := m.GetBackingStore().Set("lineNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetPostingDate sets the postingDate property value. The postingDate property
func (m *CustomerPayment) SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("postingDate", value)
    if err != nil {
        panic(err)
    }
}
// CustomerPaymentable 
type CustomerPaymentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmount()(*float64)
    GetAppliesToInvoiceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetAppliesToInvoiceNumber()(*string)
    GetComment()(*string)
    GetContactId()(*string)
    GetCustomer()(Customerable)
    GetCustomerId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetCustomerNumber()(*string)
    GetDescription()(*string)
    GetDocumentNumber()(*string)
    GetExternalDocumentNumber()(*string)
    GetJournalDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLineNumber()(*int32)
    GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAmount(value *float64)()
    SetAppliesToInvoiceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetAppliesToInvoiceNumber(value *string)()
    SetComment(value *string)()
    SetContactId(value *string)()
    SetCustomer(value Customerable)()
    SetCustomerId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetCustomerNumber(value *string)()
    SetDescription(value *string)()
    SetDocumentNumber(value *string)()
    SetExternalDocumentNumber(value *string)()
    SetJournalDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLineNumber(value *int32)()
    SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
