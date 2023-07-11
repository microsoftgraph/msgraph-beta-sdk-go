package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntent entity that represents an intent to apply settings to a device
type DeviceManagementIntent struct {
    Entity
}
// NewDeviceManagementIntent instantiates a new deviceManagementIntent and sets the default values.
func NewDeviceManagementIntent()(*DeviceManagementIntent) {
    m := &DeviceManagementIntent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntent(), nil
}
// GetAssignments gets the assignments property value. Collection of assignments
func (m *DeviceManagementIntent) GetAssignments()([]DeviceManagementIntentAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementIntentAssignmentable)
    }
    return nil
}
// GetCategories gets the categories property value. Collection of setting categories within the intent
func (m *DeviceManagementIntent) GetCategories()([]DeviceManagementIntentSettingCategoryable) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementIntentSettingCategoryable)
    }
    return nil
}
// GetDescription gets the description property value. The user given description
func (m *DeviceManagementIntent) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *DeviceManagementIntent) GetDeviceSettingStateSummaries()([]DeviceManagementIntentDeviceSettingStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceSettingStateSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementIntentDeviceSettingStateSummaryable)
    }
    return nil
}
// GetDeviceStates gets the deviceStates property value. Collection of states of all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStates()([]DeviceManagementIntentDeviceStateable) {
    val, err := m.GetBackingStore().Get("deviceStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementIntentDeviceStateable)
    }
    return nil
}
// GetDeviceStateSummary gets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStateSummary()(DeviceManagementIntentDeviceStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceStateSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementIntentDeviceStateSummaryable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The user given display name
func (m *DeviceManagementIntent) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementIntentAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentSettingCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentSettingCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementIntentSettingCategoryable)
                }
            }
            m.SetCategories(res)
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
    res["deviceSettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentDeviceSettingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentDeviceSettingStateSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementIntentDeviceSettingStateSummaryable)
                }
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentDeviceStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementIntentDeviceStateable)
                }
            }
            m.SetDeviceStates(res)
        }
        return nil
    }
    res["deviceStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementIntentDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStateSummary(val.(DeviceManagementIntentDeviceStateSummaryable))
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
    res["isAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
        }
        return nil
    }
    res["isMigratingToConfigurationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMigratingToConfigurationPolicy(val)
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
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingInstanceable)
                }
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["userStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentUserStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentUserStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementIntentUserStateable)
                }
            }
            m.SetUserStates(res)
        }
        return nil
    }
    res["userStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementIntentUserStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStateSummary(val.(DeviceManagementIntentUserStateSummaryable))
        }
        return nil
    }
    return res
}
// GetIsAssigned gets the isAssigned property value. Signifies whether or not the intent is assigned to users
func (m *DeviceManagementIntent) GetIsAssigned()(*bool) {
    val, err := m.GetBackingStore().Get("isAssigned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsMigratingToConfigurationPolicy gets the isMigratingToConfigurationPolicy property value. Signifies whether or not the intent is being migrated to the configurationPolicies endpoint
func (m *DeviceManagementIntent) GetIsMigratingToConfigurationPolicy()(*bool) {
    val, err := m.GetBackingStore().Get("isMigratingToConfigurationPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the intent was last modified
func (m *DeviceManagementIntent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementIntent) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementIntent) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSettings gets the settings property value. Collection of all settings to be applied
func (m *DeviceManagementIntent) GetSettings()([]DeviceManagementSettingInstanceable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingInstanceable)
    }
    return nil
}
// GetTemplateId gets the templateId property value. The ID of the template this intent was created from (if any)
func (m *DeviceManagementIntent) GetTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("templateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserStates gets the userStates property value. Collection of states of all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStates()([]DeviceManagementIntentUserStateable) {
    val, err := m.GetBackingStore().Get("userStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementIntentUserStateable)
    }
    return nil
}
// GetUserStateSummary gets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStateSummary()(DeviceManagementIntentUserStateSummaryable) {
    val, err := m.GetBackingStore().Get("userStateSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementIntentUserStateSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
    if m.GetDeviceSettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStateSummary", m.GetDeviceStateSummary())
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
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMigratingToConfigurationPolicy", m.GetIsMigratingToConfigurationPolicy())
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetUserStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserStates()))
        for i, v := range m.GetUserStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStateSummary", m.GetUserStateSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. Collection of assignments
func (m *DeviceManagementIntent) SetAssignments(value []DeviceManagementIntentAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCategories sets the categories property value. Collection of setting categories within the intent
func (m *DeviceManagementIntent) SetCategories(value []DeviceManagementIntentSettingCategoryable)() {
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The user given description
func (m *DeviceManagementIntent) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *DeviceManagementIntent) SetDeviceSettingStateSummaries(value []DeviceManagementIntentDeviceSettingStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceSettingStateSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceStates sets the deviceStates property value. Collection of states of all devices that the intent is applied to
func (m *DeviceManagementIntent) SetDeviceStates(value []DeviceManagementIntentDeviceStateable)() {
    err := m.GetBackingStore().Set("deviceStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceStateSummary sets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
func (m *DeviceManagementIntent) SetDeviceStateSummary(value DeviceManagementIntentDeviceStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceStateSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The user given display name
func (m *DeviceManagementIntent) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAssigned sets the isAssigned property value. Signifies whether or not the intent is assigned to users
func (m *DeviceManagementIntent) SetIsAssigned(value *bool)() {
    err := m.GetBackingStore().Set("isAssigned", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMigratingToConfigurationPolicy sets the isMigratingToConfigurationPolicy property value. Signifies whether or not the intent is being migrated to the configurationPolicies endpoint
func (m *DeviceManagementIntent) SetIsMigratingToConfigurationPolicy(value *bool)() {
    err := m.GetBackingStore().Set("isMigratingToConfigurationPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the intent was last modified
func (m *DeviceManagementIntent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementIntent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementIntent) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. Collection of all settings to be applied
func (m *DeviceManagementIntent) SetSettings(value []DeviceManagementSettingInstanceable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateId sets the templateId property value. The ID of the template this intent was created from (if any)
func (m *DeviceManagementIntent) SetTemplateId(value *string)() {
    err := m.GetBackingStore().Set("templateId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserStates sets the userStates property value. Collection of states of all users that the intent is applied to
func (m *DeviceManagementIntent) SetUserStates(value []DeviceManagementIntentUserStateable)() {
    err := m.GetBackingStore().Set("userStates", value)
    if err != nil {
        panic(err)
    }
}
// SetUserStateSummary sets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
func (m *DeviceManagementIntent) SetUserStateSummary(value DeviceManagementIntentUserStateSummaryable)() {
    err := m.GetBackingStore().Set("userStateSummary", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementIntentable 
type DeviceManagementIntentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceManagementIntentAssignmentable)
    GetCategories()([]DeviceManagementIntentSettingCategoryable)
    GetDescription()(*string)
    GetDeviceSettingStateSummaries()([]DeviceManagementIntentDeviceSettingStateSummaryable)
    GetDeviceStates()([]DeviceManagementIntentDeviceStateable)
    GetDeviceStateSummary()(DeviceManagementIntentDeviceStateSummaryable)
    GetDisplayName()(*string)
    GetIsAssigned()(*bool)
    GetIsMigratingToConfigurationPolicy()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetRoleScopeTagIds()([]string)
    GetSettings()([]DeviceManagementSettingInstanceable)
    GetTemplateId()(*string)
    GetUserStates()([]DeviceManagementIntentUserStateable)
    GetUserStateSummary()(DeviceManagementIntentUserStateSummaryable)
    SetAssignments(value []DeviceManagementIntentAssignmentable)()
    SetCategories(value []DeviceManagementIntentSettingCategoryable)()
    SetDescription(value *string)()
    SetDeviceSettingStateSummaries(value []DeviceManagementIntentDeviceSettingStateSummaryable)()
    SetDeviceStates(value []DeviceManagementIntentDeviceStateable)()
    SetDeviceStateSummary(value DeviceManagementIntentDeviceStateSummaryable)()
    SetDisplayName(value *string)()
    SetIsAssigned(value *bool)()
    SetIsMigratingToConfigurationPolicy(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetSettings(value []DeviceManagementSettingInstanceable)()
    SetTemplateId(value *string)()
    SetUserStates(value []DeviceManagementIntentUserStateable)()
    SetUserStateSummary(value DeviceManagementIntentUserStateSummaryable)()
}
