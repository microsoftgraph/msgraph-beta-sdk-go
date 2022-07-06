package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineStateSummary the security baseline compliance state summary for the security baseline of the account.
type SecurityBaselineStateSummary struct {
    Entity
    // Number of conflict devices
    conflictCount *int32
    // Number of error devices
    errorCount *int32
    // Number of not applicable devices
    notApplicableCount *int32
    // Number of not secure devices
    notSecureCount *int32
    // Number of secure devices
    secureCount *int32
    // The type property
    type_escaped *string
    // Number of unknown devices
    unknownCount *int32
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
                mappingStr := *mappingValue
                switch mappingStr {
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
    if m == nil {
        return nil
    } else {
        return m.conflictCount
    }
}
// GetErrorCount gets the errorCount property value. Number of error devices
func (m *SecurityBaselineStateSummary) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
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
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
// GetNotSecureCount gets the notSecureCount property value. Number of not secure devices
func (m *SecurityBaselineStateSummary) GetNotSecureCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notSecureCount
    }
}
// GetSecureCount gets the secureCount property value. Number of secure devices
func (m *SecurityBaselineStateSummary) GetSecureCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.secureCount
    }
}
// GetType gets the type property value. The type property
func (m *SecurityBaselineStateSummary) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUnknownCount gets the unknownCount property value. Number of unknown devices
func (m *SecurityBaselineStateSummary) GetUnknownCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownCount
    }
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
        err = writer.WriteInt32Value("secureCount", m.GetSecureCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
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
    if m != nil {
        m.conflictCount = value
    }
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *SecurityBaselineStateSummary) SetErrorCount(value *int32)() {
    if m != nil {
        m.errorCount = value
    }
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *SecurityBaselineStateSummary) SetNotApplicableCount(value *int32)() {
    if m != nil {
        m.notApplicableCount = value
    }
}
// SetNotSecureCount sets the notSecureCount property value. Number of not secure devices
func (m *SecurityBaselineStateSummary) SetNotSecureCount(value *int32)() {
    if m != nil {
        m.notSecureCount = value
    }
}
// SetSecureCount sets the secureCount property value. Number of secure devices
func (m *SecurityBaselineStateSummary) SetSecureCount(value *int32)() {
    if m != nil {
        m.secureCount = value
    }
}
// SetType sets the type property value. The type property
func (m *SecurityBaselineStateSummary) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUnknownCount sets the unknownCount property value. Number of unknown devices
func (m *SecurityBaselineStateSummary) SetUnknownCount(value *int32)() {
    if m != nil {
        m.unknownCount = value
    }
}
