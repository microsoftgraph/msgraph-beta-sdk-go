package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppPolicySetItem 
type MobileAppPolicySetItem struct {
    PolicySetItem
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Install intent of the MobileAppPolicySetItem. Possible values are: available, required, uninstall, availableWithoutEnrollment.
    intent *InstallIntent
    // Settings of the MobileAppPolicySetItem.
    settings MobileAppAssignmentSettingsable
}
// NewMobileAppPolicySetItem instantiates a new MobileAppPolicySetItem and sets the default values.
func NewMobileAppPolicySetItem()(*MobileAppPolicySetItem) {
    m := &MobileAppPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppPolicySetItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppPolicySetItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    res["intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*InstallIntent))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppAssignmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(MobileAppAssignmentSettingsable))
        }
        return nil
    }
    return res
}
// GetIntent gets the intent property value. Install intent of the MobileAppPolicySetItem. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *MobileAppPolicySetItem) GetIntent()(*InstallIntent) {
    if m == nil {
        return nil
    } else {
        return m.intent
    }
}
// GetSettings gets the settings property value. Settings of the MobileAppPolicySetItem.
func (m *MobileAppPolicySetItem) GetSettings()(MobileAppAssignmentSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Serialize serializes information the current object
func (m *MobileAppPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
        err = writer.WriteStringValue("intent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
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
func (m *MobileAppPolicySetItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIntent sets the intent property value. Install intent of the MobileAppPolicySetItem. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *MobileAppPolicySetItem) SetIntent(value *InstallIntent)() {
    if m != nil {
        m.intent = value
    }
}
// SetSettings sets the settings property value. Settings of the MobileAppPolicySetItem.
func (m *MobileAppPolicySetItem) SetSettings(value MobileAppAssignmentSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
