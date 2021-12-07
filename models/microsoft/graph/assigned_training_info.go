package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignedTrainingInfo 
type AssignedTrainingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of users who were assigned the training in an attack simulation and training campaign.
    assignedUserCount *int32;
    // Number of users who completed the training in an attack simulation and training campaign.
    completedUserCount *int32;
    // Display name of the training in an attack simulation and training campaign.
    displayName *string;
}
// NewAssignedTrainingInfo instantiates a new assignedTrainingInfo and sets the default values.
func NewAssignedTrainingInfo()(*AssignedTrainingInfo) {
    m := &AssignedTrainingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedTrainingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignedUserCount gets the assignedUserCount property value. Number of users who were assigned the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) GetAssignedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.assignedUserCount
    }
}
// GetCompletedUserCount gets the completedUserCount property value. Number of users who completed the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) GetCompletedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.completedUserCount
    }
}
// GetDisplayName gets the displayName property value. Display name of the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignedTrainingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedUserCount(val)
        }
        return nil
    }
    res["completedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedUserCount(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *AssignedTrainingInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignedTrainingInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("assignedUserCount", m.GetAssignedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("completedUserCount", m.GetCompletedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *AssignedTrainingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignedUserCount sets the assignedUserCount property value. Number of users who were assigned the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) SetAssignedUserCount(value *int32)() {
    if m != nil {
        m.assignedUserCount = value
    }
}
// SetCompletedUserCount sets the completedUserCount property value. Number of users who completed the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) SetCompletedUserCount(value *int32)() {
    if m != nil {
        m.completedUserCount = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
