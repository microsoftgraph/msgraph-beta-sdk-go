package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationSynchronizationLicenseAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The user role type to assign to license. Possible values are: student, teacher, faculty.
    appliesTo *EducationUserRole;
    // Represents the SKU identifiers of the licenses to assign.
    skuIds []string;
}
// Instantiates a new educationSynchronizationLicenseAssignment and sets the default values.
func NewEducationSynchronizationLicenseAssignment()(*EducationSynchronizationLicenseAssignment) {
    m := &EducationSynchronizationLicenseAssignment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationLicenseAssignment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appliesTo property value. The user role type to assign to license. Possible values are: student, teacher, faculty.
func (m *EducationSynchronizationLicenseAssignment) GetAppliesTo()(*EducationUserRole) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// Gets the skuIds property value. Represents the SKU identifiers of the licenses to assign.
func (m *EducationSynchronizationLicenseAssignment) GetSkuIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.skuIds
    }
}
// The deserialization information for the current model
func (m *EducationSynchronizationLicenseAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationUserRole)
        if err != nil {
            return err
        }
        cast := val.(EducationUserRole)
        m.SetAppliesTo(&cast)
        return nil
    }
    res["skuIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSkuIds(res)
        return nil
    }
    return res
}
func (m *EducationSynchronizationLicenseAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationSynchronizationLicenseAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAppliesTo() != nil {
        cast := m.GetAppliesTo().String()
        err := writer.WriteStringValue("appliesTo", &cast)
        if err != nil {
            return err
        }
    }
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EducationSynchronizationLicenseAssignment) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appliesTo property value. The user role type to assign to license. Possible values are: student, teacher, faculty.
// Parameters:
//  - value : Value to set for the appliesTo property.
func (m *EducationSynchronizationLicenseAssignment) SetAppliesTo(value *EducationUserRole)() {
    m.appliesTo = value
}
// Sets the skuIds property value. Represents the SKU identifiers of the licenses to assign.
// Parameters:
//  - value : Value to set for the skuIds property.
func (m *EducationSynchronizationLicenseAssignment) SetSkuIds(value []string)() {
    m.skuIds = value
}
