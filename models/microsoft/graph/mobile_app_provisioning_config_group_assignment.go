package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppProvisioningConfigGroupAssignment provides operations to manage the deviceAppManagement singleton.
type MobileAppProvisioningConfigGroupAssignment struct {
    Entity
    // The ID of the AAD group in which the app provisioning configuration is being targeted.
    targetGroupId *string;
}
// NewMobileAppProvisioningConfigGroupAssignment instantiates a new mobileAppProvisioningConfigGroupAssignment and sets the default values.
func NewMobileAppProvisioningConfigGroupAssignment()(*MobileAppProvisioningConfigGroupAssignment) {
    m := &MobileAppProvisioningConfigGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMobileAppProvisioningConfigGroupAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppProvisioningConfigGroupAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    return res
}
// GetTargetGroupId gets the targetGroupId property value. The ID of the AAD group in which the app provisioning configuration is being targeted.
func (m *MobileAppProvisioningConfigGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
func (m *MobileAppProvisioningConfigGroupAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MobileAppProvisioningConfigGroupAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetGroupId sets the targetGroupId property value. The ID of the AAD group in which the app provisioning configuration is being targeted.
func (m *MobileAppProvisioningConfigGroupAssignment) SetTargetGroupId(value *string)() {
    if m != nil {
        m.targetGroupId = value
    }
}
