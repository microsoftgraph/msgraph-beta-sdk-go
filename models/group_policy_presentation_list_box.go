package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationListBox 
type GroupPolicyPresentationListBox struct {
    GroupPolicyPresentation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If this option is specified true the user must specify the registry subkey value and the registry subkey name. The list box shows two columns, one for the name and one for the data. The default value is false.
    explicitValue *bool
    // Not yet documented
    valuePrefix *string
}
// NewGroupPolicyPresentationListBox instantiates a new GroupPolicyPresentationListBox and sets the default values.
func NewGroupPolicyPresentationListBox()(*GroupPolicyPresentationListBox) {
    m := &GroupPolicyPresentationListBox{
        GroupPolicyPresentation: *NewGroupPolicyPresentation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupPolicyPresentationListBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationListBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationListBox(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationListBox) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExplicitValue gets the explicitValue property value. If this option is specified true the user must specify the registry subkey value and the registry subkey name. The list box shows two columns, one for the name and one for the data. The default value is false.
func (m *GroupPolicyPresentationListBox) GetExplicitValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.explicitValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationListBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentation.GetFieldDeserializers()
    res["explicitValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExplicitValue(val)
        }
        return nil
    }
    res["valuePrefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValuePrefix(val)
        }
        return nil
    }
    return res
}
// GetValuePrefix gets the valuePrefix property value. Not yet documented
func (m *GroupPolicyPresentationListBox) GetValuePrefix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valuePrefix
    }
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationListBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("explicitValue", m.GetExplicitValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valuePrefix", m.GetValuePrefix())
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
func (m *GroupPolicyPresentationListBox) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExplicitValue sets the explicitValue property value. If this option is specified true the user must specify the registry subkey value and the registry subkey name. The list box shows two columns, one for the name and one for the data. The default value is false.
func (m *GroupPolicyPresentationListBox) SetExplicitValue(value *bool)() {
    if m != nil {
        m.explicitValue = value
    }
}
// SetValuePrefix sets the valuePrefix property value. Not yet documented
func (m *GroupPolicyPresentationListBox) SetValuePrefix(value *string)() {
    if m != nil {
        m.valuePrefix = value
    }
}
