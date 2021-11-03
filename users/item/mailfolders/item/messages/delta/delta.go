package delta

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type Delta struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookItem
    // The fileAttachment and itemAttachment attachments for the message.
    attachments []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment;
    // The Bcc: recipients for the message.
    bccRecipients []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    // The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
    body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody;
    // The first 255 characters of the message body. It is in text format.
    bodyPreview *string;
    // The Cc: recipients for the message.
    ccRecipients []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    // The ID of the conversation the email belongs to.
    conversationId *string;
    // Indicates the position of the message within the conversation.
    conversationIndex []byte;
    // The collection of open extensions defined for the message. Nullable.
    extensions []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension;
    // The flag value that indicates the status, start date, due date, or completion date for the message.
    flag *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag;
    // The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
    from *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    // Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
    hasAttachments *bool;
    // 
    importance *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance;
    // 
    inferenceClassification *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InferenceClassificationType;
    // 
    internetMessageHeaders []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InternetMessageHeader;
    // 
    internetMessageId *string;
    // 
    isDeliveryReceiptRequested *bool;
    // 
    isDraft *bool;
    // 
    isRead *bool;
    // 
    isReadReceiptRequested *bool;
    // A collection of mentions in the message, ordered by the createdDateTime from the newest to the oldest. By default, a GET /messages does not return this property unless you apply $expand on the property.
    mentions []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Mention;
    // 
    mentionsPreview *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MentionsPreview;
    // The collection of multi-value extended properties defined for the message. Nullable.
    multiValueExtendedProperties []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty;
    // 
    parentFolderId *string;
    // 
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    replyTo []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    // 
    sender *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    // 
    sentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of single-value extended properties defined for the message. Nullable.
    singleValueExtendedProperties []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty;
    // 
    subject *string;
    // 
    toRecipients []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    // 
    uniqueBody *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody;
    // 
    unsubscribeData []string;
    // 
    unsubscribeEnabled *bool;
    // 
    webLink *string;
}
// Instantiates a new delta and sets the default values.
func NewDelta()(*Delta) {
    m := &Delta{
        OutlookItem: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookItem(),
    }
    return m
}
// Gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
func (m *Delta) GetAttachments()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// Gets the bccRecipients property value. The Bcc: recipients for the message.
func (m *Delta) GetBccRecipients()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.bccRecipients
    }
}
// Gets the body property value. The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
func (m *Delta) GetBody()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// Gets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
func (m *Delta) GetBodyPreview()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bodyPreview
    }
}
// Gets the ccRecipients property value. The Cc: recipients for the message.
func (m *Delta) GetCcRecipients()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.ccRecipients
    }
}
// Gets the conversationId property value. The ID of the conversation the email belongs to.
func (m *Delta) GetConversationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conversationId
    }
}
// Gets the conversationIndex property value. Indicates the position of the message within the conversation.
func (m *Delta) GetConversationIndex()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.conversationIndex
    }
}
// Gets the extensions property value. The collection of open extensions defined for the message. Nullable.
func (m *Delta) GetExtensions()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the message.
func (m *Delta) GetFlag()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag) {
    if m == nil {
        return nil
    } else {
        return m.flag
    }
}
// Gets the from property value. The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
func (m *Delta) GetFrom()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// Gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
func (m *Delta) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// Gets the importance property value. 
func (m *Delta) GetImportance()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// Gets the inferenceClassification property value. 
func (m *Delta) GetInferenceClassification()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InferenceClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.inferenceClassification
    }
}
// Gets the internetMessageHeaders property value. 
func (m *Delta) GetInternetMessageHeaders()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InternetMessageHeader) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageHeaders
    }
}
// Gets the internetMessageId property value. 
func (m *Delta) GetInternetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageId
    }
}
// Gets the isDeliveryReceiptRequested property value. 
func (m *Delta) GetIsDeliveryReceiptRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeliveryReceiptRequested
    }
}
// Gets the isDraft property value. 
func (m *Delta) GetIsDraft()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDraft
    }
}
// Gets the isRead property value. 
func (m *Delta) GetIsRead()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRead
    }
}
// Gets the isReadReceiptRequested property value. 
func (m *Delta) GetIsReadReceiptRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadReceiptRequested
    }
}
// Gets the mentions property value. A collection of mentions in the message, ordered by the createdDateTime from the newest to the oldest. By default, a GET /messages does not return this property unless you apply $expand on the property.
func (m *Delta) GetMentions()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Mention) {
    if m == nil {
        return nil
    } else {
        return m.mentions
    }
}
// Gets the mentionsPreview property value. 
func (m *Delta) GetMentionsPreview()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MentionsPreview) {
    if m == nil {
        return nil
    } else {
        return m.mentionsPreview
    }
}
// Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Delta) GetMultiValueExtendedProperties()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// Gets the parentFolderId property value. 
func (m *Delta) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// Gets the receivedDateTime property value. 
func (m *Delta) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.receivedDateTime
    }
}
// Gets the replyTo property value. 
func (m *Delta) GetReplyTo()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.replyTo
    }
}
// Gets the sender property value. 
func (m *Delta) GetSender()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.sender
    }
}
// Gets the sentDateTime property value. 
func (m *Delta) GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sentDateTime
    }
}
// Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Delta) GetSingleValueExtendedProperties()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// Gets the subject property value. 
func (m *Delta) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the toRecipients property value. 
func (m *Delta) GetToRecipients()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.toRecipients
    }
}
// Gets the uniqueBody property value. 
func (m *Delta) GetUniqueBody()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.uniqueBody
    }
}
// Gets the unsubscribeData property value. 
func (m *Delta) GetUnsubscribeData()([]string) {
    if m == nil {
        return nil
    } else {
        return m.unsubscribeData
    }
}
// Gets the unsubscribeEnabled property value. 
func (m *Delta) GetUnsubscribeEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.unsubscribeEnabled
    }
}
// Gets the webLink property value. 
func (m *Delta) GetWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webLink
    }
}
// The deserialization information for the current model
func (m *Delta) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAttachment() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment))
        }
        m.SetAttachments(res)
        return nil
    }
    res["bccRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        }
        m.SetBccRecipients(res)
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewItemBody() })
        if err != nil {
            return err
        }
        m.SetBody(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody))
        return nil
    }
    res["bodyPreview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBodyPreview(val)
        return nil
    }
    res["ccRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        }
        m.SetCcRecipients(res)
        return nil
    }
    res["conversationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConversationId(val)
        return nil
    }
    res["conversationIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetConversationIndex(val)
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExtension() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["flag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewFollowupFlag() })
        if err != nil {
            return err
        }
        m.SetFlag(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag))
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        m.SetFrom(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasAttachments(val)
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseImportance)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance)
        m.SetImportance(&cast)
        return nil
    }
    res["inferenceClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseInferenceClassificationType)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InferenceClassificationType)
        m.SetInferenceClassification(&cast)
        return nil
    }
    res["internetMessageHeaders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewInternetMessageHeader() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InternetMessageHeader, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InternetMessageHeader))
        }
        m.SetInternetMessageHeaders(res)
        return nil
    }
    res["internetMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInternetMessageId(val)
        return nil
    }
    res["isDeliveryReceiptRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeliveryReceiptRequested(val)
        return nil
    }
    res["isDraft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDraft(val)
        return nil
    }
    res["isRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRead(val)
        return nil
    }
    res["isReadReceiptRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReadReceiptRequested(val)
        return nil
    }
    res["mentions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMention() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Mention, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Mention))
        }
        m.SetMentions(res)
        return nil
    }
    res["mentionsPreview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMentionsPreview() })
        if err != nil {
            return err
        }
        m.SetMentionsPreview(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MentionsPreview))
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty))
        }
        m.SetMultiValueExtendedProperties(res)
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentFolderId(val)
        return nil
    }
    res["receivedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReceivedDateTime(val)
        return nil
    }
    res["replyTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        }
        m.SetReplyTo(res)
        return nil
    }
    res["sender"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        m.SetSender(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        return nil
    }
    res["sentDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSentDateTime(val)
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty))
        }
        m.SetSingleValueExtendedProperties(res)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    res["toRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        }
        m.SetToRecipients(res)
        return nil
    }
    res["uniqueBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewItemBody() })
        if err != nil {
            return err
        }
        m.SetUniqueBody(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody))
        return nil
    }
    res["unsubscribeData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetUnsubscribeData(res)
        return nil
    }
    res["unsubscribeEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUnsubscribeEnabled(val)
        return nil
    }
    res["webLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebLink(val)
        return nil
    }
    return res
}
func (m *Delta) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Delta) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBccRecipients()))
        for i, v := range m.GetBccRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("bccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bodyPreview", m.GetBodyPreview())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCcRecipients()))
        for i, v := range m.GetCcRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conversationId", m.GetConversationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("conversationIndex", m.GetConversationIndex())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("flag", m.GetFlag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := m.GetImportance().String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInferenceClassification() != nil {
        cast := m.GetInferenceClassification().String()
        err = writer.WriteStringValue("inferenceClassification", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInternetMessageHeaders()))
        for i, v := range m.GetInternetMessageHeaders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("internetMessageHeaders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeliveryReceiptRequested", m.GetIsDeliveryReceiptRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDraft", m.GetIsDraft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRead", m.GetIsRead())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReadReceiptRequested", m.GetIsReadReceiptRequested())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMentions()))
        for i, v := range m.GetMentions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mentions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mentionsPreview", m.GetMentionsPreview())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReplyTo()))
        for i, v := range m.GetReplyTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("replyTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sender", m.GetSender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("sentDateTime", m.GetSentDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetToRecipients()))
        for i, v := range m.GetToRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("toRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("uniqueBody", m.GetUniqueBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("unsubscribeData", m.GetUnsubscribeData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("unsubscribeEnabled", m.GetUnsubscribeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webLink", m.GetWebLink())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
// Parameters:
//  - value : Value to set for the attachments property.
func (m *Delta) SetAttachments(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment)() {
    m.attachments = value
}
// Sets the bccRecipients property value. The Bcc: recipients for the message.
// Parameters:
//  - value : Value to set for the bccRecipients property.
func (m *Delta) SetBccRecipients(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    m.bccRecipients = value
}
// Sets the body property value. The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
// Parameters:
//  - value : Value to set for the body property.
func (m *Delta) SetBody(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody)() {
    m.body = value
}
// Sets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
// Parameters:
//  - value : Value to set for the bodyPreview property.
func (m *Delta) SetBodyPreview(value *string)() {
    m.bodyPreview = value
}
// Sets the ccRecipients property value. The Cc: recipients for the message.
// Parameters:
//  - value : Value to set for the ccRecipients property.
func (m *Delta) SetCcRecipients(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    m.ccRecipients = value
}
// Sets the conversationId property value. The ID of the conversation the email belongs to.
// Parameters:
//  - value : Value to set for the conversationId property.
func (m *Delta) SetConversationId(value *string)() {
    m.conversationId = value
}
// Sets the conversationIndex property value. Indicates the position of the message within the conversation.
// Parameters:
//  - value : Value to set for the conversationIndex property.
func (m *Delta) SetConversationIndex(value []byte)() {
    m.conversationIndex = value
}
// Sets the extensions property value. The collection of open extensions defined for the message. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *Delta) SetExtensions(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension)() {
    m.extensions = value
}
// Sets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the message.
// Parameters:
//  - value : Value to set for the flag property.
func (m *Delta) SetFlag(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag)() {
    m.flag = value
}
// Sets the from property value. The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
// Parameters:
//  - value : Value to set for the from property.
func (m *Delta) SetFrom(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    m.from = value
}
// Sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
// Parameters:
//  - value : Value to set for the hasAttachments property.
func (m *Delta) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// Sets the importance property value. 
// Parameters:
//  - value : Value to set for the importance property.
func (m *Delta) SetImportance(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance)() {
    m.importance = value
}
// Sets the inferenceClassification property value. 
// Parameters:
//  - value : Value to set for the inferenceClassification property.
func (m *Delta) SetInferenceClassification(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InferenceClassificationType)() {
    m.inferenceClassification = value
}
// Sets the internetMessageHeaders property value. 
// Parameters:
//  - value : Value to set for the internetMessageHeaders property.
func (m *Delta) SetInternetMessageHeaders(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InternetMessageHeader)() {
    m.internetMessageHeaders = value
}
// Sets the internetMessageId property value. 
// Parameters:
//  - value : Value to set for the internetMessageId property.
func (m *Delta) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// Sets the isDeliveryReceiptRequested property value. 
// Parameters:
//  - value : Value to set for the isDeliveryReceiptRequested property.
func (m *Delta) SetIsDeliveryReceiptRequested(value *bool)() {
    m.isDeliveryReceiptRequested = value
}
// Sets the isDraft property value. 
// Parameters:
//  - value : Value to set for the isDraft property.
func (m *Delta) SetIsDraft(value *bool)() {
    m.isDraft = value
}
// Sets the isRead property value. 
// Parameters:
//  - value : Value to set for the isRead property.
func (m *Delta) SetIsRead(value *bool)() {
    m.isRead = value
}
// Sets the isReadReceiptRequested property value. 
// Parameters:
//  - value : Value to set for the isReadReceiptRequested property.
func (m *Delta) SetIsReadReceiptRequested(value *bool)() {
    m.isReadReceiptRequested = value
}
// Sets the mentions property value. A collection of mentions in the message, ordered by the createdDateTime from the newest to the oldest. By default, a GET /messages does not return this property unless you apply $expand on the property.
// Parameters:
//  - value : Value to set for the mentions property.
func (m *Delta) SetMentions(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Mention)() {
    m.mentions = value
}
// Sets the mentionsPreview property value. 
// Parameters:
//  - value : Value to set for the mentionsPreview property.
func (m *Delta) SetMentionsPreview(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MentionsPreview)() {
    m.mentionsPreview = value
}
// Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
// Parameters:
//  - value : Value to set for the multiValueExtendedProperties property.
func (m *Delta) SetMultiValueExtendedProperties(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// Sets the parentFolderId property value. 
// Parameters:
//  - value : Value to set for the parentFolderId property.
func (m *Delta) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// Sets the receivedDateTime property value. 
// Parameters:
//  - value : Value to set for the receivedDateTime property.
func (m *Delta) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// Sets the replyTo property value. 
// Parameters:
//  - value : Value to set for the replyTo property.
func (m *Delta) SetReplyTo(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    m.replyTo = value
}
// Sets the sender property value. 
// Parameters:
//  - value : Value to set for the sender property.
func (m *Delta) SetSender(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    m.sender = value
}
// Sets the sentDateTime property value. 
// Parameters:
//  - value : Value to set for the sentDateTime property.
func (m *Delta) SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sentDateTime = value
}
// Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
// Parameters:
//  - value : Value to set for the singleValueExtendedProperties property.
func (m *Delta) SetSingleValueExtendedProperties(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
// Sets the subject property value. 
// Parameters:
//  - value : Value to set for the subject property.
func (m *Delta) SetSubject(value *string)() {
    m.subject = value
}
// Sets the toRecipients property value. 
// Parameters:
//  - value : Value to set for the toRecipients property.
func (m *Delta) SetToRecipients(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    m.toRecipients = value
}
// Sets the uniqueBody property value. 
// Parameters:
//  - value : Value to set for the uniqueBody property.
func (m *Delta) SetUniqueBody(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody)() {
    m.uniqueBody = value
}
// Sets the unsubscribeData property value. 
// Parameters:
//  - value : Value to set for the unsubscribeData property.
func (m *Delta) SetUnsubscribeData(value []string)() {
    m.unsubscribeData = value
}
// Sets the unsubscribeEnabled property value. 
// Parameters:
//  - value : Value to set for the unsubscribeEnabled property.
func (m *Delta) SetUnsubscribeEnabled(value *bool)() {
    m.unsubscribeEnabled = value
}
// Sets the webLink property value. 
// Parameters:
//  - value : Value to set for the webLink property.
func (m *Delta) SetWebLink(value *string)() {
    m.webLink = value
}
