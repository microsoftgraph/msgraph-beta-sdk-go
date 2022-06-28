package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusDetails 
type StatusDetails struct {
    StatusBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Additional details in case of error.
    additionalDetails *string
    // Categorizes the error code. Possible values are Failure, NonServiceFailure, Success.
    errorCategory *ProvisioningStatusErrorCategory
    // Unique error code if any occurred. Learn more
    errorCode *string
    // Summarizes the status and describes why the status happened.
    reason *string
    // Provides the resolution for the corresponding error.
    recommendedAction *string
}
// NewStatusDetails instantiates a new StatusDetails and sets the default values.
func NewStatusDetails()(*StatusDetails) {
    m := &StatusDetails{
        StatusBase: *NewStatusBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatusDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StatusDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalDetails gets the additionalDetails property value. Additional details in case of error.
func (m *StatusDetails) GetAdditionalDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
// GetErrorCategory gets the errorCategory property value. Categorizes the error code. Possible values are Failure, NonServiceFailure, Success.
func (m *StatusDetails) GetErrorCategory()(*ProvisioningStatusErrorCategory) {
    if m == nil {
        return nil
    } else {
        return m.errorCategory
    }
}
// GetErrorCode gets the errorCode property value. Unique error code if any occurred. Learn more
func (m *StatusDetails) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StatusBase.GetFieldDeserializers()
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetails(val)
        }
        return nil
    }
    res["errorCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningStatusErrorCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCategory(val.(*ProvisioningStatusErrorCategory))
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["recommendedAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedAction(val)
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. Summarizes the status and describes why the status happened.
func (m *StatusDetails) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetRecommendedAction gets the recommendedAction property value. Provides the resolution for the corresponding error.
func (m *StatusDetails) GetRecommendedAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedAction
    }
}
// Serialize serializes information the current object
func (m *StatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StatusBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    if m.GetErrorCategory() != nil {
        cast := (*m.GetErrorCategory()).String()
        err = writer.WriteStringValue("errorCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
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
func (m *StatusDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdditionalDetails sets the additionalDetails property value. Additional details in case of error.
func (m *StatusDetails) SetAdditionalDetails(value *string)() {
    if m != nil {
        m.additionalDetails = value
    }
}
// SetErrorCategory sets the errorCategory property value. Categorizes the error code. Possible values are Failure, NonServiceFailure, Success.
func (m *StatusDetails) SetErrorCategory(value *ProvisioningStatusErrorCategory)() {
    if m != nil {
        m.errorCategory = value
    }
}
// SetErrorCode sets the errorCode property value. Unique error code if any occurred. Learn more
func (m *StatusDetails) SetErrorCode(value *string)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetReason sets the reason property value. Summarizes the status and describes why the status happened.
func (m *StatusDetails) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetRecommendedAction sets the recommendedAction property value. Provides the resolution for the corresponding error.
func (m *StatusDetails) SetRecommendedAction(value *string)() {
    if m != nil {
        m.recommendedAction = value
    }
}
