package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppleOwnerTypeEnrollmentType 
type AppleOwnerTypeEnrollmentType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The enrollment type. Possible values are: unknown, device, user.
    enrollmentType *AppleUserInitiatedEnrollmentType;
    // The owner type. Possible values are: unknown, company, personal.
    ownerType *ManagedDeviceOwnerType;
}
// NewAppleOwnerTypeEnrollmentType instantiates a new appleOwnerTypeEnrollmentType and sets the default values.
func NewAppleOwnerTypeEnrollmentType()(*AppleOwnerTypeEnrollmentType) {
    m := &AppleOwnerTypeEnrollmentType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppleOwnerTypeEnrollmentType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnrollmentType gets the enrollmentType property value. The enrollment type. Possible values are: unknown, device, user.
func (m *AppleOwnerTypeEnrollmentType) GetEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentType
    }
}
// GetOwnerType gets the ownerType property value. The owner type. Possible values are: unknown, company, personal.
func (m *AppleOwnerTypeEnrollmentType) GetOwnerType()(*ManagedDeviceOwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleOwnerTypeEnrollmentType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val.(*AppleUserInitiatedEnrollmentType))
        }
        return nil
    }
    res["ownerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    return res
}
func (m *AppleOwnerTypeEnrollmentType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppleOwnerTypeEnrollmentType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetEnrollmentType() != nil {
        cast := (*m.GetEnrollmentType()).String()
        err := writer.WriteStringValue("enrollmentType", &cast)
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppleOwnerTypeEnrollmentType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnrollmentType sets the enrollmentType property value. The enrollment type. Possible values are: unknown, device, user.
func (m *AppleOwnerTypeEnrollmentType) SetEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    if m != nil {
        m.enrollmentType = value
    }
}
// SetOwnerType sets the ownerType property value. The owner type. Possible values are: unknown, company, personal.
func (m *AppleOwnerTypeEnrollmentType) SetOwnerType(value *ManagedDeviceOwnerType)() {
    if m != nil {
        m.ownerType = value
    }
}
