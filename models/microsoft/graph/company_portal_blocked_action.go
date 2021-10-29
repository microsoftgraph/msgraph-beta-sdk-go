package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CompanyPortalBlockedAction struct {
    // Device Action. Possible values are: unknown, remove, reset.
    action *CompanyPortalAction;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Device ownership type. Possible values are: unknown, company, personal.
    ownerType *OwnerType;
    // Device OS/Platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
    platform *DevicePlatformType;
}
// Instantiates a new companyPortalBlockedAction and sets the default values.
func NewCompanyPortalBlockedAction()(*CompanyPortalBlockedAction) {
    m := &CompanyPortalBlockedAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the action property value. Device Action. Possible values are: unknown, remove, reset.
func (m *CompanyPortalBlockedAction) GetAction()(*CompanyPortalAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyPortalBlockedAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the ownerType property value. Device ownership type. Possible values are: unknown, company, personal.
func (m *CompanyPortalBlockedAction) GetOwnerType()(*OwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// Gets the platform property value. Device OS/Platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *CompanyPortalBlockedAction) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// The deserialization information for the current model
func (m *CompanyPortalBlockedAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCompanyPortalAction)
        if err != nil {
            return err
        }
        cast := val.(CompanyPortalAction)
        m.SetAction(&cast)
        return nil
    }
    res["ownerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        cast := val.(OwnerType)
        m.SetOwnerType(&cast)
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        cast := val.(DevicePlatformType)
        m.SetPlatform(&cast)
        return nil
    }
    return res
}
func (m *CompanyPortalBlockedAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CompanyPortalBlockedAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnerType() != nil {
        cast := m.GetOwnerType().String()
        err := writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := m.GetPlatform().String()
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
// Sets the action property value. Device Action. Possible values are: unknown, remove, reset.
// Parameters:
//  - value : Value to set for the action property.
func (m *CompanyPortalBlockedAction) SetAction(value *CompanyPortalAction)() {
    m.action = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CompanyPortalBlockedAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the ownerType property value. Device ownership type. Possible values are: unknown, company, personal.
// Parameters:
//  - value : Value to set for the ownerType property.
func (m *CompanyPortalBlockedAction) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
// Sets the platform property value. Device OS/Platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
// Parameters:
//  - value : Value to set for the platform property.
func (m *CompanyPortalBlockedAction) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
