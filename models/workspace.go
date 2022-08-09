package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Workspace 
type Workspace struct {
    Place
    // The building property
    building *string
    // The capacity property
    capacity *int32
    // The emailAddress property
    emailAddress *string
    // The floorLabel property
    floorLabel *string
    // The floorNumber property
    floorNumber *int32
    // The isWheelChairAccessible property
    isWheelChairAccessible *bool
    // The label property
    label *string
    // The nickname property
    nickname *string
    // The tags property
    tags []string
}
// NewWorkspace instantiates a new Workspace and sets the default values.
func NewWorkspace()(*Workspace) {
    m := &Workspace{
        Place: *NewPlace(),
    }
    odataTypeValue := "#microsoft.graph.workspace";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkspaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkspaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkspace(), nil
}
// GetBuilding gets the building property value. The building property
func (m *Workspace) GetBuilding()(*string) {
    if m == nil {
        return nil
    } else {
        return m.building
    }
}
// GetCapacity gets the capacity property value. The capacity property
func (m *Workspace) GetCapacity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.capacity
    }
}
// GetEmailAddress gets the emailAddress property value. The emailAddress property
func (m *Workspace) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Workspace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Place.GetFieldDeserializers()
    res["building"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuilding(val)
        }
        return nil
    }
    res["capacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacity(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["floorLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorLabel(val)
        }
        return nil
    }
    res["floorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorNumber(val)
        }
        return nil
    }
    res["isWheelChairAccessible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsWheelChairAccessible(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["nickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickname(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetFloorLabel gets the floorLabel property value. The floorLabel property
func (m *Workspace) GetFloorLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.floorLabel
    }
}
// GetFloorNumber gets the floorNumber property value. The floorNumber property
func (m *Workspace) GetFloorNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.floorNumber
    }
}
// GetIsWheelChairAccessible gets the isWheelChairAccessible property value. The isWheelChairAccessible property
func (m *Workspace) GetIsWheelChairAccessible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isWheelChairAccessible
    }
}
// GetLabel gets the label property value. The label property
func (m *Workspace) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// GetNickname gets the nickname property value. The nickname property
func (m *Workspace) GetNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nickname
    }
}
// GetTags gets the tags property value. The tags property
func (m *Workspace) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Serialize serializes information the current object
func (m *Workspace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Place.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("building", m.GetBuilding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("capacity", m.GetCapacity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("floorLabel", m.GetFloorLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("floorNumber", m.GetFloorNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isWheelChairAccessible", m.GetIsWheelChairAccessible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickname", m.GetNickname())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuilding sets the building property value. The building property
func (m *Workspace) SetBuilding(value *string)() {
    if m != nil {
        m.building = value
    }
}
// SetCapacity sets the capacity property value. The capacity property
func (m *Workspace) SetCapacity(value *int32)() {
    if m != nil {
        m.capacity = value
    }
}
// SetEmailAddress sets the emailAddress property value. The emailAddress property
func (m *Workspace) SetEmailAddress(value *string)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetFloorLabel sets the floorLabel property value. The floorLabel property
func (m *Workspace) SetFloorLabel(value *string)() {
    if m != nil {
        m.floorLabel = value
    }
}
// SetFloorNumber sets the floorNumber property value. The floorNumber property
func (m *Workspace) SetFloorNumber(value *int32)() {
    if m != nil {
        m.floorNumber = value
    }
}
// SetIsWheelChairAccessible sets the isWheelChairAccessible property value. The isWheelChairAccessible property
func (m *Workspace) SetIsWheelChairAccessible(value *bool)() {
    if m != nil {
        m.isWheelChairAccessible = value
    }
}
// SetLabel sets the label property value. The label property
func (m *Workspace) SetLabel(value *string)() {
    if m != nil {
        m.label = value
    }
}
// SetNickname sets the nickname property value. The nickname property
func (m *Workspace) SetNickname(value *string)() {
    if m != nil {
        m.nickname = value
    }
}
// SetTags sets the tags property value. The tags property
func (m *Workspace) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
