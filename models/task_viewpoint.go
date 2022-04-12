package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TaskViewpoint 
type TaskViewpoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The categories associated with the task. Each category corresponds to the displayName property of an outlookCategory that the user has defined.
    categories []string
    // The date and time for a reminder alert of the task to occur.
    reminderDateTime DateTimeTimeZoneable
}
// NewTaskViewpoint instantiates a new taskViewpoint and sets the default values.
func NewTaskViewpoint()(*TaskViewpoint) {
    m := &TaskViewpoint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTaskViewpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskViewpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskViewpoint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TaskViewpoint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategories gets the categories property value. The categories associated with the task. Each category corresponds to the displayName property of an outlookCategory that the user has defined.
func (m *TaskViewpoint) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskViewpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["reminderDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetReminderDateTime gets the reminderDateTime property value. The date and time for a reminder alert of the task to occur.
func (m *TaskViewpoint) GetReminderDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.reminderDateTime
    }
}
// Serialize serializes information the current object
func (m *TaskViewpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCategories() != nil {
        err := writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reminderDateTime", m.GetReminderDateTime())
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
func (m *TaskViewpoint) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategories sets the categories property value. The categories associated with the task. Each category corresponds to the displayName property of an outlookCategory that the user has defined.
func (m *TaskViewpoint) SetCategories(value []string)() {
    if m != nil {
        m.categories = value
    }
}
// SetReminderDateTime sets the reminderDateTime property value. The date and time for a reminder alert of the task to occur.
func (m *TaskViewpoint) SetReminderDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.reminderDateTime = value
    }
}
