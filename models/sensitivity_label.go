package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityLabel provides operations to manage the collection of accessReview entities.
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
    odataTypeValue := "#microsoft.graph.sensitivityLabel";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSensitivityLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabel(), nil
}
// GetApplicableTo gets the applicableTo property value. The applicableTo property
func (m *SensitivityLabel) GetApplicableTo()(*SensitivityLabelTarget) {
    return m.applicableTo
}
// GetApplicationMode gets the applicationMode property value. The applicationMode property
func (m *SensitivityLabel) GetApplicationMode()(*ApplicationMode) {
    return m.applicationMode
}
// GetAssignedPolicies gets the assignedPolicies property value. The assignedPolicies property
func (m *SensitivityLabel) GetAssignedPolicies()([]LabelPolicyable) {
    return m.assignedPolicies
}
// GetAutoLabeling gets the autoLabeling property value. The autoLabeling property
func (m *SensitivityLabel) GetAutoLabeling()(AutoLabelingable) {
    return m.autoLabeling
}
// GetDescription gets the description property value. The description property
func (m *SensitivityLabel) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *SensitivityLabel) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSensitivityLabelTarget , m.SetApplicableTo)
    res["applicationMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseApplicationMode , m.SetApplicationMode)
    res["assignedPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLabelPolicyFromDiscriminatorValue , m.SetAssignedPolicies)
    res["autoLabeling"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAutoLabelingFromDiscriminatorValue , m.SetAutoLabeling)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isDefault"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDefault)
    res["isEndpointProtectionEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEndpointProtectionEnabled)
    res["labelActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLabelActionBaseFromDiscriminatorValue , m.SetLabelActions)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["sublabels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue , m.SetSublabels)
    res["toolTip"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetToolTip)
    return res
}
// GetIsDefault gets the isDefault property value. The isDefault property
func (m *SensitivityLabel) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetIsEndpointProtectionEnabled gets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *SensitivityLabel) GetIsEndpointProtectionEnabled()(*bool) {
    return m.isEndpointProtectionEnabled
}
// GetLabelActions gets the labelActions property value. The labelActions property
func (m *SensitivityLabel) GetLabelActions()([]LabelActionBaseable) {
    return m.labelActions
}
// GetName gets the name property value. The name property
func (m *SensitivityLabel) GetName()(*string) {
    return m.name
}
// GetPriority gets the priority property value. The priority property
func (m *SensitivityLabel) GetPriority()(*int32) {
    return m.priority
}
// GetSublabels gets the sublabels property value. The sublabels property
func (m *SensitivityLabel) GetSublabels()([]SensitivityLabelable) {
    return m.sublabels
}
// GetToolTip gets the toolTip property value. The toolTip property
func (m *SensitivityLabel) GetToolTip()(*string) {
    return m.toolTip
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignedPolicies())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLabelActions())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSublabels())
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
    m.applicableTo = value
}
// SetApplicationMode sets the applicationMode property value. The applicationMode property
func (m *SensitivityLabel) SetApplicationMode(value *ApplicationMode)() {
    m.applicationMode = value
}
// SetAssignedPolicies sets the assignedPolicies property value. The assignedPolicies property
func (m *SensitivityLabel) SetAssignedPolicies(value []LabelPolicyable)() {
    m.assignedPolicies = value
}
// SetAutoLabeling sets the autoLabeling property value. The autoLabeling property
func (m *SensitivityLabel) SetAutoLabeling(value AutoLabelingable)() {
    m.autoLabeling = value
}
// SetDescription sets the description property value. The description property
func (m *SensitivityLabel) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *SensitivityLabel) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *SensitivityLabel) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetIsEndpointProtectionEnabled sets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *SensitivityLabel) SetIsEndpointProtectionEnabled(value *bool)() {
    m.isEndpointProtectionEnabled = value
}
// SetLabelActions sets the labelActions property value. The labelActions property
func (m *SensitivityLabel) SetLabelActions(value []LabelActionBaseable)() {
    m.labelActions = value
}
// SetName sets the name property value. The name property
func (m *SensitivityLabel) SetName(value *string)() {
    m.name = value
}
// SetPriority sets the priority property value. The priority property
func (m *SensitivityLabel) SetPriority(value *int32)() {
    m.priority = value
}
// SetSublabels sets the sublabels property value. The sublabels property
func (m *SensitivityLabel) SetSublabels(value []SensitivityLabelable)() {
    m.sublabels = value
}
// SetToolTip sets the toolTip property value. The toolTip property
func (m *SensitivityLabel) SetToolTip(value *string)() {
    m.toolTip = value
}
