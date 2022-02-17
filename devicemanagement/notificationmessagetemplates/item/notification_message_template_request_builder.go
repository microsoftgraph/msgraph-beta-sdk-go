package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i92bdddc6d39f7ceb527fbc3a0c281fdad40f0d03ea69fda531a27f1fb720fdc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/notificationmessagetemplates/item/sendtestmessage"
    ie131f760e231c41c85ab4a627830ad8e63680d71cfdfeb33b08cb086c57f4e8a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages"
    i8f7ffa70b6e5e6cd3a41114a81ef71a7ef4f81cad3974df901008b8a88e8cf8e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages/item"
)

// NotificationMessageTemplateRequestBuilder builds and executes requests for operations under \deviceManagement\notificationMessageTemplates\{notificationMessageTemplate-id}
type NotificationMessageTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NotificationMessageTemplateRequestBuilderDeleteOptions options for Delete
type NotificationMessageTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotificationMessageTemplateRequestBuilderGetOptions options for Get
type NotificationMessageTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NotificationMessageTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotificationMessageTemplateRequestBuilderGetQueryParameters the Notification Message Templates.
type NotificationMessageTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NotificationMessageTemplateRequestBuilderPatchOptions options for Patch
type NotificationMessageTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotificationMessageTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewNotificationMessageTemplateRequestBuilderInternal instantiates a new NotificationMessageTemplateRequestBuilder and sets the default values.
func NewNotificationMessageTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotificationMessageTemplateRequestBuilder) {
    m := &NotificationMessageTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotificationMessageTemplateRequestBuilder instantiates a new NotificationMessageTemplateRequestBuilder and sets the default values.
func NewNotificationMessageTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotificationMessageTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotificationMessageTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) CreateDeleteRequestInformation(options *NotificationMessageTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) CreateGetRequestInformation(options *NotificationMessageTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) CreatePatchRequestInformation(options *NotificationMessageTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) Delete(options *NotificationMessageTemplateRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) Get(options *NotificationMessageTemplateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotificationMessageTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewNotificationMessageTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotificationMessageTemplate), nil
}
func (m *NotificationMessageTemplateRequestBuilder) LocalizedNotificationMessages()(*ie131f760e231c41c85ab4a627830ad8e63680d71cfdfeb33b08cb086c57f4e8a.LocalizedNotificationMessagesRequestBuilder) {
    return ie131f760e231c41c85ab4a627830ad8e63680d71cfdfeb33b08cb086c57f4e8a.NewLocalizedNotificationMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizedNotificationMessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.notificationMessageTemplates.item.localizedNotificationMessages.item collection
func (m *NotificationMessageTemplateRequestBuilder) LocalizedNotificationMessagesById(id string)(*i8f7ffa70b6e5e6cd3a41114a81ef71a7ef4f81cad3974df901008b8a88e8cf8e.LocalizedNotificationMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["localizedNotificationMessage_id"] = id
    }
    return i8f7ffa70b6e5e6cd3a41114a81ef71a7ef4f81cad3974df901008b8a88e8cf8e.NewLocalizedNotificationMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the Notification Message Templates.
func (m *NotificationMessageTemplateRequestBuilder) Patch(options *NotificationMessageTemplateRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *NotificationMessageTemplateRequestBuilder) SendTestMessage()(*i92bdddc6d39f7ceb527fbc3a0c281fdad40f0d03ea69fda531a27f1fb720fdc0.SendTestMessageRequestBuilder) {
    return i92bdddc6d39f7ceb527fbc3a0c281fdad40f0d03ea69fda531a27f1fb720fdc0.NewSendTestMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
