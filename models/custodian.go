package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Custodian 
type Custodian struct {
    DataSourceContainer
    // Date and time the custodian acknowledged a hold notification.
    acknowledgedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identifies whether a custodian's sources were placed on hold during creation.
    applyHoldToSources *bool
    // Email address of the custodian.
    email *string
}
// NewCustodian instantiates a new Custodian and sets the default values.
func NewCustodian()(*Custodian) {
    m := &Custodian{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCustodianFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustodianFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustodian(), nil
}
// GetAcknowledgedDateTime gets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
func (m *Custodian) GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.acknowledgedDateTime
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Custodian) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplyHoldToSources gets the applyHoldToSources property value. Identifies whether a custodian's sources were placed on hold during creation.
func (m *Custodian) GetApplyHoldToSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyHoldToSources
    }
}
// GetEmail gets the email property value. Email address of the custodian.
func (m *Custodian) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Custodian) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["acknowledgedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcknowledgedDateTime(val)
        }
        return nil
    }
    res["applyHoldToSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyHoldToSources(val)
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
    return res
}
// Serialize serializes information the current object
func (m *Custodian) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acknowledgedDateTime", m.GetAcknowledgedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applyHoldToSources", m.GetApplyHoldToSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcknowledgedDateTime sets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
func (m *Custodian) SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.acknowledgedDateTime = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Custodian) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplyHoldToSources sets the applyHoldToSources property value. Identifies whether a custodian's sources were placed on hold during creation.
func (m *Custodian) SetApplyHoldToSources(value *bool)() {
    if m != nil {
        m.applyHoldToSources = value
    }
}
// SetEmail sets the email property value. Email address of the custodian.
func (m *Custodian) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
