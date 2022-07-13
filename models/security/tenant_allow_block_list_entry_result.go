package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAllowBlockListEntryResult 
type TenantAllowBlockListEntryResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The entryType property
    entryType *TenantAllowBlockListEntryType
    // The expirationDateTime property
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identity property
    identity *string
    // The status property
    status *LongRunningOperationStatus
    // The value property
    value *string
}
// NewTenantAllowBlockListEntryResult instantiates a new tenantAllowBlockListEntryResult and sets the default values.
func NewTenantAllowBlockListEntryResult()(*TenantAllowBlockListEntryResult) {
    m := &TenantAllowBlockListEntryResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantAllowBlockListEntryResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAllowBlockListEntryResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAllowBlockListEntryResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAllowBlockListEntryResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEntryType gets the entryType property value. The entryType property
func (m *TenantAllowBlockListEntryResult) GetEntryType()(*TenantAllowBlockListEntryType) {
    if m == nil {
        return nil
    } else {
        return m.entryType
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *TenantAllowBlockListEntryResult) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAllowBlockListEntryResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["entryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantAllowBlockListEntryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntryType(val.(*TenantAllowBlockListEntryType))
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLongRunningOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*LongRunningOperationStatus))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity property
func (m *TenantAllowBlockListEntryResult) GetIdentity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
// GetStatus gets the status property value. The status property
func (m *TenantAllowBlockListEntryResult) GetStatus()(*LongRunningOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetValue gets the value property value. The value property
func (m *TenantAllowBlockListEntryResult) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *TenantAllowBlockListEntryResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEntryType() != nil {
        cast := (*m.GetEntryType()).String()
        err := writer.WriteStringValue("entryType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAllowBlockListEntryResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEntryType sets the entryType property value. The entryType property
func (m *TenantAllowBlockListEntryResult) SetEntryType(value *TenantAllowBlockListEntryType)() {
    if m != nil {
        m.entryType = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *TenantAllowBlockListEntryResult) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetIdentity sets the identity property value. The identity property
func (m *TenantAllowBlockListEntryResult) SetIdentity(value *string)() {
    if m != nil {
        m.identity = value
    }
}
// SetStatus sets the status property value. The status property
func (m *TenantAllowBlockListEntryResult) SetStatus(value *LongRunningOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetValue sets the value property value. The value property
func (m *TenantAllowBlockListEntryResult) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
