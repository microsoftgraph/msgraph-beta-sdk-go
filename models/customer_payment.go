package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerPayment provides operations to manage the collection of accessReviewDecision entities.
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
    odataTypeValue := "#microsoft.graph.customerPayment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomerPaymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerPaymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomerPayment(), nil
}
// GetAmount gets the amount property value. The amount property
func (m *CustomerPayment) GetAmount()(*float64) {
    return m.amount
}
// GetAppliesToInvoiceId gets the appliesToInvoiceId property value. The appliesToInvoiceId property
func (m *CustomerPayment) GetAppliesToInvoiceId()(*string) {
    return m.appliesToInvoiceId
}
// GetAppliesToInvoiceNumber gets the appliesToInvoiceNumber property value. The appliesToInvoiceNumber property
func (m *CustomerPayment) GetAppliesToInvoiceNumber()(*string) {
    return m.appliesToInvoiceNumber
}
// GetComment gets the comment property value. The comment property
func (m *CustomerPayment) GetComment()(*string) {
    return m.comment
}
// GetContactId gets the contactId property value. The contactId property
func (m *CustomerPayment) GetContactId()(*string) {
    return m.contactId
}
// GetCustomer gets the customer property value. The customer property
func (m *CustomerPayment) GetCustomer()(Customerable) {
    return m.customer
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *CustomerPayment) GetCustomerId()(*string) {
    return m.customerId
}
// GetCustomerNumber gets the customerNumber property value. The customerNumber property
func (m *CustomerPayment) GetCustomerNumber()(*string) {
    return m.customerNumber
}
// GetDescription gets the description property value. The description property
func (m *CustomerPayment) GetDescription()(*string) {
    return m.description
}
// GetDocumentNumber gets the documentNumber property value. The documentNumber property
func (m *CustomerPayment) GetDocumentNumber()(*string) {
    return m.documentNumber
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *CustomerPayment) GetExternalDocumentNumber()(*string) {
    return m.externalDocumentNumber
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomerPayment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["amount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetAmount)
    res["appliesToInvoiceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppliesToInvoiceId)
    res["appliesToInvoiceNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppliesToInvoiceNumber)
    res["comment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetComment)
    res["contactId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContactId)
    res["customer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCustomerFromDiscriminatorValue , m.SetCustomer)
    res["customerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerId)
    res["customerNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerNumber)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["documentNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDocumentNumber)
    res["externalDocumentNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalDocumentNumber)
    res["journalDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJournalDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["lineNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetLineNumber)
    res["postingDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetPostingDate)
    return res
}
// GetJournalDisplayName gets the journalDisplayName property value. The journalDisplayName property
func (m *CustomerPayment) GetJournalDisplayName()(*string) {
    return m.journalDisplayName
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPayment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLineNumber gets the lineNumber property value. The lineNumber property
func (m *CustomerPayment) GetLineNumber()(*int32) {
    return m.lineNumber
}
// GetPostingDate gets the postingDate property value. The postingDate property
func (m *CustomerPayment) GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.postingDate
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
    m.amount = value
}
// SetAppliesToInvoiceId sets the appliesToInvoiceId property value. The appliesToInvoiceId property
func (m *CustomerPayment) SetAppliesToInvoiceId(value *string)() {
    m.appliesToInvoiceId = value
}
// SetAppliesToInvoiceNumber sets the appliesToInvoiceNumber property value. The appliesToInvoiceNumber property
func (m *CustomerPayment) SetAppliesToInvoiceNumber(value *string)() {
    m.appliesToInvoiceNumber = value
}
// SetComment sets the comment property value. The comment property
func (m *CustomerPayment) SetComment(value *string)() {
    m.comment = value
}
// SetContactId sets the contactId property value. The contactId property
func (m *CustomerPayment) SetContactId(value *string)() {
    m.contactId = value
}
// SetCustomer sets the customer property value. The customer property
func (m *CustomerPayment) SetCustomer(value Customerable)() {
    m.customer = value
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *CustomerPayment) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetCustomerNumber sets the customerNumber property value. The customerNumber property
func (m *CustomerPayment) SetCustomerNumber(value *string)() {
    m.customerNumber = value
}
// SetDescription sets the description property value. The description property
func (m *CustomerPayment) SetDescription(value *string)() {
    m.description = value
}
// SetDocumentNumber sets the documentNumber property value. The documentNumber property
func (m *CustomerPayment) SetDocumentNumber(value *string)() {
    m.documentNumber = value
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. The externalDocumentNumber property
func (m *CustomerPayment) SetExternalDocumentNumber(value *string)() {
    m.externalDocumentNumber = value
}
// SetJournalDisplayName sets the journalDisplayName property value. The journalDisplayName property
func (m *CustomerPayment) SetJournalDisplayName(value *string)() {
    m.journalDisplayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CustomerPayment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLineNumber sets the lineNumber property value. The lineNumber property
func (m *CustomerPayment) SetLineNumber(value *int32)() {
    m.lineNumber = value
}
// SetPostingDate sets the postingDate property value. The postingDate property
func (m *CustomerPayment) SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.postingDate = value
}
