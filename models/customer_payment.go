package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerPayment provides operations to manage the financials singleton.
type CustomerPayment struct {
    Entity
    // The amount property
    amount *float64
    // The appliesToInvoiceId property
    appliesToInvoiceId *string
    // The appliesToInvoiceNumber property
    appliesToInvoiceNumber *string
    // The comment property
    comment *string
    // The contactId property
    contactId *string
    // The customer property
    customer Customerable
    // The customerId property
    customerId *string
    // The customerNumber property
    customerNumber *string
    // The description property
    description *string
    // The documentNumber property
    documentNumber *string
    // The externalDocumentNumber property
    externalDocumentNumber *string
    // The journalDisplayName property
    journalDisplayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lineNumber property
    lineNumber *int32
    // The postingDate property
    postingDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
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
    if m == nil {
        return nil
    } else {
        return m.amount
    }
}
// GetAppliesToInvoiceId gets the appliesToInvoiceId property value. The appliesToInvoiceId property
func (m *CustomerPayment) GetAppliesToInvoiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appliesToInvoiceId
    }
}
// GetAppliesToInvoiceNumber gets the appliesToInvoiceNumber property value. The appliesToInvoiceNumber property
func (m *CustomerPayment) GetAppliesToInvoiceNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appliesToInvoiceNumber
    }
}
// GetComment gets the comment property value. The comment property
func (m *CustomerPayment) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetContactId gets the contactId property value. The contactId property
func (m *CustomerPayment) GetContactId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactId
    }
}
// GetCustomer gets the customer property value. The customer property
func (m *CustomerPayment) GetCustomer()(Customerable) {
    if m == nil {
        return nil
    } else {
        return m.customer
    }
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *CustomerPayment) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *CustomerPayment) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
// GetDescription gets the description property value. The description property
func (m *CustomerPayment) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDocumentNumber gets the documentNumber property value. The documentNumber property
func (m *CustomerPayment) GetDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentNumber
    }
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *CustomerPayment) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
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
        val, err := n.GetStringValue()
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
        val, err := n.GetStringValue()
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
    if m == nil {
        return nil
    } else {
        return m.journalDisplayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPayment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLineNumber gets the lineNumber property value. The lineNumber property
func (m *CustomerPayment) GetLineNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lineNumber
    }
}
// GetPostingDate gets the postingDate property value. The postingDate property
func (m *CustomerPayment) GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.postingDate
    }
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
        err = writer.WriteStringValue("appliesToInvoiceId", m.GetAppliesToInvoiceId())
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
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
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
    if m != nil {
        m.amount = value
    }
}
// SetAppliesToInvoiceId sets the appliesToInvoiceId property value. The appliesToInvoiceId property
func (m *CustomerPayment) SetAppliesToInvoiceId(value *string)() {
    if m != nil {
        m.appliesToInvoiceId = value
    }
}
// SetAppliesToInvoiceNumber sets the appliesToInvoiceNumber property value. The appliesToInvoiceNumber property
func (m *CustomerPayment) SetAppliesToInvoiceNumber(value *string)() {
    if m != nil {
        m.appliesToInvoiceNumber = value
    }
}
// SetComment sets the comment property value. The comment property
func (m *CustomerPayment) SetComment(value *string)() {
    if m != nil {
        m.comment = value
    }
}
// SetContactId sets the contactId property value. The contactId property
func (m *CustomerPayment) SetContactId(value *string)() {
    if m != nil {
        m.contactId = value
    }
}
// SetCustomer sets the customer property value. The customer property
func (m *CustomerPayment) SetCustomer(value Customerable)() {
    if m != nil {
        m.customer = value
    }
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *CustomerPayment) SetCustomerId(value *string)() {
    if m != nil {
        m.customerId = value
    }
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *CustomerPayment) SetCustomerNumber(value *string)() {
    if m != nil {
        m.customerNumber = value
    }
}
// SetDescription sets the description property value. The description property
func (m *CustomerPayment) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDocumentNumber sets the documentNumber property value. The documentNumber property
func (m *CustomerPayment) SetDocumentNumber(value *string)() {
    if m != nil {
        m.documentNumber = value
    }
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *CustomerPayment) SetExternalDocumentNumber(value *string)() {
    if m != nil {
        m.externalDocumentNumber = value
    }
}
// SetJournalDisplayName sets the journalDisplayName property value. The journalDisplayName property
func (m *CustomerPayment) SetJournalDisplayName(value *string)() {
    if m != nil {
        m.journalDisplayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPayment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLineNumber sets the lineNumber property value. The lineNumber property
func (m *CustomerPayment) SetLineNumber(value *int32)() {
    if m != nil {
        m.lineNumber = value
    }
}
// SetPostingDate sets the postingDate property value. The postingDate property
func (m *CustomerPayment) SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.postingDate = value
    }
}
