package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSummary 
type UserSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of failed tasks for users in a user summary.
    failedTasks *int32
    // The number of failed users in a user summary.
    failedUsers *int32
    // The OdataType property
    odataType *string
    // The number of successful users in a user summary.
    successfulUsers *int32
    // The total tasks of users in a user summary.
    totalTasks *int32
    // The total number of users in a user summary
    totalUsers *int32
}
// NewUserSummary instantiates a new userSummary and sets the default values.
func NewUserSummary()(*UserSummary) {
    m := &UserSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFailedTasks gets the failedTasks property value. The number of failed tasks for users in a user summary.
func (m *UserSummary) GetFailedTasks()(*int32) {
    return m.failedTasks
}
// GetFailedUsers gets the failedUsers property value. The number of failed users in a user summary.
func (m *UserSummary) GetFailedUsers()(*int32) {
    return m.failedUsers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedTasks(val)
        }
        return nil
    }
    res["failedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUsers(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["successfulUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulUsers(val)
        }
        return nil
    }
    res["totalTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTasks(val)
        }
        return nil
    }
    res["totalUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUsers(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetSuccessfulUsers gets the successfulUsers property value. The number of successful users in a user summary.
func (m *UserSummary) GetSuccessfulUsers()(*int32) {
    return m.successfulUsers
}
// GetTotalTasks gets the totalTasks property value. The total tasks of users in a user summary.
func (m *UserSummary) GetTotalTasks()(*int32) {
    return m.totalTasks
}
// GetTotalUsers gets the totalUsers property value. The total number of users in a user summary
func (m *UserSummary) GetTotalUsers()(*int32) {
    return m.totalUsers
}
// Serialize serializes information the current object
func (m *UserSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("failedTasks", m.GetFailedTasks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("failedUsers", m.GetFailedUsers())
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
    {
        err := writer.WriteInt32Value("successfulUsers", m.GetSuccessfulUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalTasks", m.GetTotalTasks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalUsers", m.GetTotalUsers())
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
func (m *UserSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFailedTasks sets the failedTasks property value. The number of failed tasks for users in a user summary.
func (m *UserSummary) SetFailedTasks(value *int32)() {
    m.failedTasks = value
}
// SetFailedUsers sets the failedUsers property value. The number of failed users in a user summary.
func (m *UserSummary) SetFailedUsers(value *int32)() {
    m.failedUsers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuccessfulUsers sets the successfulUsers property value. The number of successful users in a user summary.
func (m *UserSummary) SetSuccessfulUsers(value *int32)() {
    m.successfulUsers = value
}
// SetTotalTasks sets the totalTasks property value. The total tasks of users in a user summary.
func (m *UserSummary) SetTotalTasks(value *int32)() {
    m.totalTasks = value
}
// SetTotalUsers sets the totalUsers property value. The total number of users in a user summary
func (m *UserSummary) SetTotalUsers(value *int32)() {
    m.totalUsers = value
}
