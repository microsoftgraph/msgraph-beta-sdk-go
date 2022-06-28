package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionLabel 
type RetentionLabel struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents the user who created the retentionLabel.
    createdBy IdentitySetable
    // Represents the date and time in which the retentionLabel is created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Provides label information for the admin. Optional.
    descriptionForAdmins *string
    // Provides the label information for the user. Optional.
    descriptionForUsers *string
    // Unique string that defines a label name.
    displayName *string
    // Specifies whether the label is currently being used.
    isInUse *bool
    // Specifies the replacement label to be applied automatically after the retention period of the current label ends.
    labelToBeApplied *string
    // The user who last modified the retentionLabel.
    lastModifiedBy IdentitySetable
    // The latest date time when the retentionLabel was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewRetentionLabel instantiates a new RetentionLabel and sets the default values.
func NewRetentionLabel()(*RetentionLabel) {
    m := &RetentionLabel{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRetentionLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionLabel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetentionLabel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCreatedBy gets the createdBy property value. Represents the user who created the retentionLabel.
func (m *RetentionLabel) GetCreatedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Represents the date and time in which the retentionLabel is created.
func (m *RetentionLabel) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescriptionForAdmins gets the descriptionForAdmins property value. Provides label information for the admin. Optional.
func (m *RetentionLabel) GetDescriptionForAdmins()(*string) {
    if m == nil {
        return nil
    } else {
        return m.descriptionForAdmins
    }
}
// GetDescriptionForUsers gets the descriptionForUsers property value. Provides the label information for the user. Optional.
func (m *RetentionLabel) GetDescriptionForUsers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.descriptionForUsers
    }
}
// GetDisplayName gets the displayName property value. Unique string that defines a label name.
func (m *RetentionLabel) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["descriptionForAdmins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescriptionForAdmins(val)
        }
        return nil
    }
    res["descriptionForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescriptionForUsers(val)
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
    res["isInUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInUse(val)
        }
        return nil
    }
    res["labelToBeApplied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelToBeApplied(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
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
    return res
}
// GetIsInUse gets the isInUse property value. Specifies whether the label is currently being used.
func (m *RetentionLabel) GetIsInUse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInUse
    }
}
// GetLabelToBeApplied gets the labelToBeApplied property value. Specifies the replacement label to be applied automatically after the retention period of the current label ends.
func (m *RetentionLabel) GetLabelToBeApplied()(*string) {
    if m == nil {
        return nil
    } else {
        return m.labelToBeApplied
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the retentionLabel.
func (m *RetentionLabel) GetLastModifiedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date time when the retentionLabel was modified.
func (m *RetentionLabel) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Serialize serializes information the current object
func (m *RetentionLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForAdmins", m.GetDescriptionForAdmins())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForUsers", m.GetDescriptionForUsers())
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
        err = writer.WriteBoolValue("isInUse", m.GetIsInUse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("labelToBeApplied", m.GetLabelToBeApplied())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetentionLabel) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCreatedBy sets the createdBy property value. Represents the user who created the retentionLabel.
func (m *RetentionLabel) SetCreatedBy(value IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Represents the date and time in which the retentionLabel is created.
func (m *RetentionLabel) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescriptionForAdmins sets the descriptionForAdmins property value. Provides label information for the admin. Optional.
func (m *RetentionLabel) SetDescriptionForAdmins(value *string)() {
    if m != nil {
        m.descriptionForAdmins = value
    }
}
// SetDescriptionForUsers sets the descriptionForUsers property value. Provides the label information for the user. Optional.
func (m *RetentionLabel) SetDescriptionForUsers(value *string)() {
    if m != nil {
        m.descriptionForUsers = value
    }
}
// SetDisplayName sets the displayName property value. Unique string that defines a label name.
func (m *RetentionLabel) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsInUse sets the isInUse property value. Specifies whether the label is currently being used.
func (m *RetentionLabel) SetIsInUse(value *bool)() {
    if m != nil {
        m.isInUse = value
    }
}
// SetLabelToBeApplied sets the labelToBeApplied property value. Specifies the replacement label to be applied automatically after the retention period of the current label ends.
func (m *RetentionLabel) SetLabelToBeApplied(value *string)() {
    if m != nil {
        m.labelToBeApplied = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the retentionLabel.
func (m *RetentionLabel) SetLastModifiedBy(value IdentitySetable)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date time when the retentionLabel was modified.
func (m *RetentionLabel) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
