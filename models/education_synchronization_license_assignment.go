package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationLicenseAssignment 
type EducationSynchronizationLicenseAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The user role type to assign to license. Possible values are: student, teacher, faculty.
    appliesTo *EducationUserRole
    // The OdataType property
    odataType *string
    // Represents the SKU identifiers of the licenses to assign.
    skuIds []string
}
// NewEducationSynchronizationLicenseAssignment instantiates a new educationSynchronizationLicenseAssignment and sets the default values.
func NewEducationSynchronizationLicenseAssignment()(*EducationSynchronizationLicenseAssignment) {
    m := &EducationSynchronizationLicenseAssignment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationSynchronizationLicenseAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationLicenseAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationLicenseAssignment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationLicenseAssignment) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppliesTo gets the appliesTo property value. The user role type to assign to license. Possible values are: student, teacher, faculty.
func (m *EducationSynchronizationLicenseAssignment) GetAppliesTo()(*EducationUserRole) {
    return m.appliesTo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationLicenseAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliesTo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEducationUserRole , m.SetAppliesTo)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["skuIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSkuIds)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationLicenseAssignment) GetOdataType()(*string) {
    return m.odataType
}
// GetSkuIds gets the skuIds property value. Represents the SKU identifiers of the licenses to assign.
func (m *EducationSynchronizationLicenseAssignment) GetSkuIds()([]string) {
    return m.skuIds
}
// Serialize serializes information the current object
func (m *EducationSynchronizationLicenseAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err := writer.WriteStringValue("appliesTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetAppliesTo sets the appliesTo property value. The user role type to assign to license. Possible values are: student, teacher, faculty.
func (m *EducationSynchronizationLicenseAssignment) SetAppliesTo(value *EducationUserRole)() {
    m.appliesTo = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationLicenseAssignment) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSkuIds sets the skuIds property value. Represents the SKU identifiers of the licenses to assign.
func (m *EducationSynchronizationLicenseAssignment) SetSkuIds(value []string)() {
    m.skuIds = value
}
