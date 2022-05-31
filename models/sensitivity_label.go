package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityLabel casts the previous resource to group.
type SensitivityLabel struct {
    Entity
    // The applicableTo property
    applicableTo *SensitivityLabelTarget
    // The applicationMode property
    applicationMode *ApplicationMode
    // The assignedPolicies property
    assignedPolicies []LabelPolicyable
    // The autoLabeling property
    autoLabeling AutoLabelingable
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The isDefault property
    isDefault *bool
    // The isEndpointProtectionEnabled property
    isEndpointProtectionEnabled *bool
    // The labelActions property
    labelActions []LabelActionBaseable
    // The name property
    name *string
    // The priority property
    priority *int32
    // The sublabels property
    sublabels []SensitivityLabelable
    // The toolTip property
    toolTip *string
}
// NewSensitivityLabel instantiates a new sensitivityLabel and sets the default values.
func NewSensitivityLabel()(*SensitivityLabel) {
    m := &SensitivityLabel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSensitivityLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabel(), nil
}
// GetApplicableTo gets the applicableTo property value. The applicableTo property
func (m *SensitivityLabel) GetApplicableTo()(*SensitivityLabelTarget) {
    if m == nil {
        return nil
    } else {
        return m.applicableTo
    }
}
// GetApplicationMode gets the applicationMode property value. The applicationMode property
func (m *SensitivityLabel) GetApplicationMode()(*ApplicationMode) {
    if m == nil {
        return nil
    } else {
        return m.applicationMode
    }
}
// GetAssignedPolicies gets the assignedPolicies property value. The assignedPolicies property
func (m *SensitivityLabel) GetAssignedPolicies()([]LabelPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.assignedPolicies
    }
}
// GetAutoLabeling gets the autoLabeling property value. The autoLabeling property
func (m *SensitivityLabel) GetAutoLabeling()(AutoLabelingable) {
    if m == nil {
        return nil
    } else {
        return m.autoLabeling
    }
}
// GetDescription gets the description property value. The description property
func (m *SensitivityLabel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *SensitivityLabel) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityLabelTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableTo(val.(*SensitivityLabelTarget))
        }
        return nil
    }
    res["applicationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationMode(val.(*ApplicationMode))
        }
        return nil
    }
    res["assignedPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(LabelPolicyable)
            }
            m.SetAssignedPolicies(res)
        }
        return nil
    }
    res["autoLabeling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutoLabelingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoLabeling(val.(AutoLabelingable))
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
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isEndpointProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEndpointProtectionEnabled(val)
        }
        return nil
    }
    res["labelActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelActionBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelActionBaseable, len(val))
            for i, v := range val {
                res[i] = v.(LabelActionBaseable)
            }
            m.SetLabelActions(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["sublabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                res[i] = v.(SensitivityLabelable)
            }
            m.SetSublabels(res)
        }
        return nil
    }
    res["toolTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToolTip(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. The isDefault property
func (m *SensitivityLabel) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsEndpointProtectionEnabled gets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *SensitivityLabel) GetIsEndpointProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEndpointProtectionEnabled
    }
}
// GetLabelActions gets the labelActions property value. The labelActions property
func (m *SensitivityLabel) GetLabelActions()([]LabelActionBaseable) {
    if m == nil {
        return nil
    } else {
        return m.labelActions
    }
}
// GetName gets the name property value. The name property
func (m *SensitivityLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPriority gets the priority property value. The priority property
func (m *SensitivityLabel) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetSublabels gets the sublabels property value. The sublabels property
func (m *SensitivityLabel) GetSublabels()([]SensitivityLabelable) {
    if m == nil {
        return nil
    } else {
        return m.sublabels
    }
}
// GetToolTip gets the toolTip property value. The toolTip property
func (m *SensitivityLabel) GetToolTip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.toolTip
    }
}
// Serialize serializes information the current object
func (m *SensitivityLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableTo() != nil {
        cast := (*m.GetApplicableTo()).String()
        err = writer.WriteStringValue("applicableTo", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicationMode() != nil {
        cast := (*m.GetApplicationMode()).String()
        err = writer.WriteStringValue("applicationMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignedPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedPolicies()))
        for i, v := range m.GetAssignedPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignedPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("autoLabeling", m.GetAutoLabeling())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEndpointProtectionEnabled", m.GetIsEndpointProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetLabelActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabelActions()))
        for i, v := range m.GetLabelActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("labelActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetSublabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSublabels()))
        for i, v := range m.GetSublabels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sublabels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("toolTip", m.GetToolTip())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableTo sets the applicableTo property value. The applicableTo property
func (m *SensitivityLabel) SetApplicableTo(value *SensitivityLabelTarget)() {
    if m != nil {
        m.applicableTo = value
    }
}
// SetApplicationMode sets the applicationMode property value. The applicationMode property
func (m *SensitivityLabel) SetApplicationMode(value *ApplicationMode)() {
    if m != nil {
        m.applicationMode = value
    }
}
// SetAssignedPolicies sets the assignedPolicies property value. The assignedPolicies property
func (m *SensitivityLabel) SetAssignedPolicies(value []LabelPolicyable)() {
    if m != nil {
        m.assignedPolicies = value
    }
}
// SetAutoLabeling sets the autoLabeling property value. The autoLabeling property
func (m *SensitivityLabel) SetAutoLabeling(value AutoLabelingable)() {
    if m != nil {
        m.autoLabeling = value
    }
}
// SetDescription sets the description property value. The description property
func (m *SensitivityLabel) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *SensitivityLabel) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *SensitivityLabel) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsEndpointProtectionEnabled sets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *SensitivityLabel) SetIsEndpointProtectionEnabled(value *bool)() {
    if m != nil {
        m.isEndpointProtectionEnabled = value
    }
}
// SetLabelActions sets the labelActions property value. The labelActions property
func (m *SensitivityLabel) SetLabelActions(value []LabelActionBaseable)() {
    if m != nil {
        m.labelActions = value
    }
}
// SetName sets the name property value. The name property
func (m *SensitivityLabel) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPriority sets the priority property value. The priority property
func (m *SensitivityLabel) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetSublabels sets the sublabels property value. The sublabels property
func (m *SensitivityLabel) SetSublabels(value []SensitivityLabelable)() {
    if m != nil {
        m.sublabels = value
    }
}
// SetToolTip sets the toolTip property value. The toolTip property
func (m *SensitivityLabel) SetToolTip(value *string)() {
    if m != nil {
        m.toolTip = value
    }
}
