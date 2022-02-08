package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSynchronizationLicenseAssignment 
type EducationSynchronizationLicenseAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The user role type to assign to license. Possible values are: student, teacher, faculty.
    appliesTo *EducationUserRole;
    // Represents the SKU identifiers of the licenses to assign.
    skuIds []string;
}
// NewEducationSynchronizationLicenseAssignment instantiates a new educationSynchronizationLicenseAssignment and sets the default values.
func NewEducationSynchronizationLicenseAssignment()(*EducationSynchronizationLicenseAssignment) {
    m := &EducationSynchronizationLicenseAssignment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationLicenseAssignment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppliesTo gets the appliesTo property value. The user role type to assign to license. Possible values are: student, teacher, faculty.
func (m *EducationSynchronizationLicenseAssignment) GetAppliesTo()(*EducationUserRole) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetSkuIds gets the skuIds property value. Represents the SKU identifiers of the licenses to assign.
func (m *EducationSynchronizationLicenseAssignment) GetSkuIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.skuIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationLicenseAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val.(*EducationUserRole))
        }
        return nil
    }
    res["skuIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSkuIds(res)
        }
        return nil
    }
    return res
}
func (m *EducationSynchronizationLicenseAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationSynchronizationLicenseAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err := writer.WriteStringValue("appliesTo", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSkuIds() != nil {
        err := writer.WriteCollectionOfStringValues("skuIds", m.GetSkuIds())
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
func (m *EducationSynchronizationLicenseAssignment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppliesTo sets the appliesTo property value. The user role type to assign to license. Possible values are: student, teacher, faculty.
func (m *EducationSynchronizationLicenseAssignment) SetAppliesTo(value *EducationUserRole)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetSkuIds sets the skuIds property value. Represents the SKU identifiers of the licenses to assign.
func (m *EducationSynchronizationLicenseAssignment) SetSkuIds(value []string)() {
    if m != nil {
        m.skuIds = value
    }
}
