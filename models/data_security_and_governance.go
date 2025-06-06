// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DataSecurityAndGovernance struct {
    Entity
}
// NewDataSecurityAndGovernance instantiates a new DataSecurityAndGovernance and sets the default values.
func NewDataSecurityAndGovernance()(*DataSecurityAndGovernance) {
    m := &DataSecurityAndGovernance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDataSecurityAndGovernanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDataSecurityAndGovernanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.tenantDataSecurityAndGovernance":
                        return NewTenantDataSecurityAndGovernance(), nil
                    case "#microsoft.graph.userDataSecurityAndGovernance":
                        return NewUserDataSecurityAndGovernance(), nil
                }
            }
        }
    }
    return NewDataSecurityAndGovernance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DataSecurityAndGovernance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["sensitivityLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SensitivityLabelable)
                }
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    return res
}
// GetSensitivityLabels gets the sensitivityLabels property value. The sensitivityLabels property
// returns a []SensitivityLabelable when successful
func (m *DataSecurityAndGovernance) GetSensitivityLabels()([]SensitivityLabelable) {
    val, err := m.GetBackingStore().Get("sensitivityLabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SensitivityLabelable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DataSecurityAndGovernance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSensitivityLabels sets the sensitivityLabels property value. The sensitivityLabels property
func (m *DataSecurityAndGovernance) SetSensitivityLabels(value []SensitivityLabelable)() {
    err := m.GetBackingStore().Set("sensitivityLabels", value)
    if err != nil {
        panic(err)
    }
}
type DataSecurityAndGovernanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSensitivityLabels()([]SensitivityLabelable)
    SetSensitivityLabels(value []SensitivityLabelable)()
}
