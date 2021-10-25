package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ServiceHealthIssuePost struct {
    additionalData map[string]interface{};
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    description *ItemBody;
    postType *PostType;
}
func NewServiceHealthIssuePost()(*ServiceHealthIssuePost) {
    m := &ServiceHealthIssuePost{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ServiceHealthIssuePost) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ServiceHealthIssuePost) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *ServiceHealthIssuePost) GetDescription()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *ServiceHealthIssuePost) GetPostType()(*PostType) {
    if m == nil {
        return nil
    } else {
        return m.postType
    }
}
func (m *ServiceHealthIssuePost) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetDescription(val.(*ItemBody))
        return nil
    }
    res["postType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePostType)
        if err != nil {
            return err
        }
        cast := val.(PostType)
        m.SetPostType(&cast)
        return nil
    }
    return res
}
func (m *ServiceHealthIssuePost) IsNil()(bool) {
    return m == nil
}
func (m *ServiceHealthIssuePost) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetPostType() != nil {
        cast := m.GetPostType().String()
        err := writer.WriteStringValue("postType", &cast)
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
func (m *ServiceHealthIssuePost) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ServiceHealthIssuePost) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *ServiceHealthIssuePost) SetDescription(value *ItemBody)() {
    m.description = value
}
func (m *ServiceHealthIssuePost) SetPostType(value *PostType)() {
    m.postType = value
}
