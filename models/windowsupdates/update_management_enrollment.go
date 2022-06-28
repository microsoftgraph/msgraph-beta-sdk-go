package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateManagementEnrollment 
type UpdateManagementEnrollment struct {
    UpdatableAssetEnrollment
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The category of updates that the service manages. Supports a subset of the values for updateCategory. Possible values are: feature, unknownFutureValue.
    updateCategory *UpdateCategory
}
// NewUpdateManagementEnrollment instantiates a new UpdateManagementEnrollment and sets the default values.
func NewUpdateManagementEnrollment()(*UpdateManagementEnrollment) {
    m := &UpdateManagementEnrollment{
        UpdatableAssetEnrollment: *NewUpdatableAssetEnrollment(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateManagementEnrollmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateManagementEnrollmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateManagementEnrollment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateManagementEnrollment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateManagementEnrollment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UpdatableAssetEnrollment.GetFieldDeserializers()
    res["updateCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*UpdateCategory))
        }
        return nil
    }
    return res
}
// GetUpdateCategory gets the updateCategory property value. The category of updates that the service manages. Supports a subset of the values for updateCategory. Possible values are: feature, unknownFutureValue.
func (m *UpdateManagementEnrollment) GetUpdateCategory()(*UpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
// Serialize serializes information the current object
func (m *UpdateManagementEnrollment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UpdatableAssetEnrollment.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err = writer.WriteStringValue("updateCategory", &cast)
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
func (m *UpdateManagementEnrollment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUpdateCategory sets the updateCategory property value. The category of updates that the service manages. Supports a subset of the values for updateCategory. Possible values are: feature, unknownFutureValue.
func (m *UpdateManagementEnrollment) SetUpdateCategory(value *UpdateCategory)() {
    if m != nil {
        m.updateCategory = value
    }
}
