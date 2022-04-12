package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanyPortalBlockedAction blocked actions on the company portal as per platform and device ownership types
type CompanyPortalBlockedAction struct {
    // Device Action. Possible values are: unknown, remove, reset.
    action *CompanyPortalAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Device ownership type. Possible values are: unknown, company, personal.
    ownerType *OwnerType
    // Device OS/Platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
    platform *DevicePlatformType
}
// NewCompanyPortalBlockedAction instantiates a new companyPortalBlockedAction and sets the default values.
func NewCompanyPortalBlockedAction()(*CompanyPortalBlockedAction) {
    m := &CompanyPortalBlockedAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCompanyPortalBlockedActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanyPortalBlockedActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanyPortalBlockedAction(), nil
}
// GetAction gets the action property value. Device Action. Possible values are: unknown, remove, reset.
func (m *CompanyPortalBlockedAction) GetAction()(*CompanyPortalAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyPortalBlockedAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompanyPortalBlockedAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCompanyPortalAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*CompanyPortalAction))
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*OwnerType))
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    return res
}
// GetOwnerType gets the ownerType property value. Device ownership type. Possible values are: unknown, company, personal.
func (m *CompanyPortalBlockedAction) GetOwnerType()(*OwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// GetPlatform gets the platform property value. Device OS/Platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *CompanyPortalBlockedAction) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// Serialize serializes information the current object
func (m *CompanyPortalBlockedAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err := writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
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
// SetAction sets the action property value. Device Action. Possible values are: unknown, remove, reset.
func (m *CompanyPortalBlockedAction) SetAction(value *CompanyPortalAction)() {
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyPortalBlockedAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOwnerType sets the ownerType property value. Device ownership type. Possible values are: unknown, company, personal.
func (m *CompanyPortalBlockedAction) SetOwnerType(value *OwnerType)() {
    if m != nil {
        m.ownerType = value
    }
}
// SetPlatform sets the platform property value. Device OS/Platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *CompanyPortalBlockedAction) SetPlatform(value *DevicePlatformType)() {
    if m != nil {
        m.platform = value
    }
}
