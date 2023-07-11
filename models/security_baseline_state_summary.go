package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineStateSummary the security baseline compliance state summary for the security baseline of the account.
type SecurityBaselineStateSummary struct {
    Entity
}
// NewSecurityBaselineStateSummary instantiates a new securityBaselineStateSummary and sets the default values.
func NewSecurityBaselineStateSummary()(*SecurityBaselineStateSummary) {
    m := &SecurityBaselineStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityBaselineStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.securityBaselineCategoryStateSummary":
                        return NewSecurityBaselineCategoryStateSummary(), nil
                }
            }
        }
    }
    return NewSecurityBaselineStateSummary(), nil
}
// GetConflictCount gets the conflictCount property value. Number of conflict devices
func (m *SecurityBaselineStateSummary) GetConflictCount()(*int32) {
    val, err := m.GetBackingStore().Get("conflictCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetErrorCount gets the errorCount property value. Number of error devices
func (m *SecurityBaselineStateSummary) GetErrorCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictCount(val)
        }
        return nil
    }
    res["errorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["notSecureCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotSecureCount(val)
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
    res["secureCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureCount(val)
        }
        return nil
    }
    res["unknownCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownCount(val)
        }
        return nil
    }
    return res
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
func (m *SecurityBaselineStateSummary) GetNotApplicableCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicableCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotSecureCount gets the notSecureCount property value. Number of not secure devices
func (m *SecurityBaselineStateSummary) GetNotSecureCount()(*int32) {
    val, err := m.GetBackingStore().Get("notSecureCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SecurityBaselineStateSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecureCount gets the secureCount property value. Number of secure devices
func (m *SecurityBaselineStateSummary) GetSecureCount()(*int32) {
    val, err := m.GetBackingStore().Get("secureCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUnknownCount gets the unknownCount property value. Number of unknown devices
func (m *SecurityBaselineStateSummary) GetUnknownCount()(*int32) {
    val, err := m.GetBackingStore().Get("unknownCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityBaselineStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("conflictCount", m.GetConflictCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableCount", m.GetNotApplicableCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notSecureCount", m.GetNotSecureCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("secureCount", m.GetSecureCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownCount", m.GetUnknownCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConflictCount sets the conflictCount property value. Number of conflict devices
func (m *SecurityBaselineStateSummary) SetConflictCount(value *int32)() {
    err := m.GetBackingStore().Set("conflictCount", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *SecurityBaselineStateSummary) SetErrorCount(value *int32)() {
    err := m.GetBackingStore().Set("errorCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *SecurityBaselineStateSummary) SetNotApplicableCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicableCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotSecureCount sets the notSecureCount property value. Number of not secure devices
func (m *SecurityBaselineStateSummary) SetNotSecureCount(value *int32)() {
    err := m.GetBackingStore().Set("notSecureCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SecurityBaselineStateSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSecureCount sets the secureCount property value. Number of secure devices
func (m *SecurityBaselineStateSummary) SetSecureCount(value *int32)() {
    err := m.GetBackingStore().Set("secureCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnknownCount sets the unknownCount property value. Number of unknown devices
func (m *SecurityBaselineStateSummary) SetUnknownCount(value *int32)() {
    err := m.GetBackingStore().Set("unknownCount", value)
    if err != nil {
        panic(err)
    }
}
// SecurityBaselineStateSummaryable 
type SecurityBaselineStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflictCount()(*int32)
    GetErrorCount()(*int32)
    GetNotApplicableCount()(*int32)
    GetNotSecureCount()(*int32)
    GetOdataType()(*string)
    GetSecureCount()(*int32)
    GetUnknownCount()(*int32)
    SetConflictCount(value *int32)()
    SetErrorCount(value *int32)()
    SetNotApplicableCount(value *int32)()
    SetNotSecureCount(value *int32)()
    SetOdataType(value *string)()
    SetSecureCount(value *int32)()
    SetUnknownCount(value *int32)()
}
