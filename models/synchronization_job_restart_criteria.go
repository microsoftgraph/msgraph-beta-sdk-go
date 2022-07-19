package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobRestartCriteria 
type SynchronizationJobRestartCriteria struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Comma-separated combination of the following values: Full, QuarantineState, Watermark, Escrows, ConnectorDataStore. Use Full if you want all of the options.
    resetScope *SynchronizationJobRestartScope
}
// NewSynchronizationJobRestartCriteria instantiates a new synchronizationJobRestartCriteria and sets the default values.
func NewSynchronizationJobRestartCriteria()(*SynchronizationJobRestartCriteria) {
    m := &SynchronizationJobRestartCriteria{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.synchronizationJobRestartCriteria";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSynchronizationJobRestartCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobRestartCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJobRestartCriteria(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationJobRestartCriteria) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobRestartCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["resetScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationJobRestartScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetScope(val.(*SynchronizationJobRestartScope))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationJobRestartCriteria) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetResetScope gets the resetScope property value. Comma-separated combination of the following values: Full, QuarantineState, Watermark, Escrows, ConnectorDataStore. Use Full if you want all of the options.
func (m *SynchronizationJobRestartCriteria) GetResetScope()(*SynchronizationJobRestartScope) {
    if m == nil {
        return nil
    } else {
        return m.resetScope
    }
}
// Serialize serializes information the current object
func (m *SynchronizationJobRestartCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResetScope() != nil {
        cast := (*m.GetResetScope()).String()
        err := writer.WriteStringValue("resetScope", &cast)
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
func (m *SynchronizationJobRestartCriteria) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationJobRestartCriteria) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetResetScope sets the resetScope property value. Comma-separated combination of the following values: Full, QuarantineState, Watermark, Escrows, ConnectorDataStore. Use Full if you want all of the options.
func (m *SynchronizationJobRestartCriteria) SetResetScope(value *SynchronizationJobRestartScope)() {
    if m != nil {
        m.resetScope = value
    }
}
