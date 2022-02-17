package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ib1036e35f7726bb63968292aa4eec005a101ab799cf62d3d86c2229a632e0fd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/item/attachments"
    iba6314f9ae414737a93628d9803d2f5d079b4663032af4e52dead50e53425e58 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/item/attachmentsarchive"
    ia518794b9649eb040a30d702e0d35056c1eef1acf65564115cc1be6834064a1f "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/item/attachments/item"
)

// ServiceUpdateMessageRequestBuilder builds and executes requests for operations under \admin\serviceAnnouncement\messages\{serviceUpdateMessage-id}
type ServiceUpdateMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServiceUpdateMessageRequestBuilderDeleteOptions options for Delete
type ServiceUpdateMessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceUpdateMessageRequestBuilderGetOptions options for Get
type ServiceUpdateMessageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServiceUpdateMessageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceUpdateMessageRequestBuilderGetQueryParameters a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
type ServiceUpdateMessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ServiceUpdateMessageRequestBuilderPatchOptions options for Patch
type ServiceUpdateMessageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ServiceUpdateMessageRequestBuilder) Attachments()(*ib1036e35f7726bb63968292aa4eec005a101ab799cf62d3d86c2229a632e0fd6.AttachmentsRequestBuilder) {
    return ib1036e35f7726bb63968292aa4eec005a101ab799cf62d3d86c2229a632e0fd6.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServiceUpdateMessageRequestBuilder) AttachmentsArchive()(*iba6314f9ae414737a93628d9803d2f5d079b4663032af4e52dead50e53425e58.AttachmentsArchiveRequestBuilder) {
    return iba6314f9ae414737a93628d9803d2f5d079b4663032af4e52dead50e53425e58.NewAttachmentsArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.serviceAnnouncement.messages.item.attachments.item collection
func (m *ServiceUpdateMessageRequestBuilder) AttachmentsById(id string)(*ia518794b9649eb040a30d702e0d35056c1eef1acf65564115cc1be6834064a1f.ServiceAnnouncementAttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceAnnouncementAttachment_id"] = id
    }
    return ia518794b9649eb040a30d702e0d35056c1eef1acf65564115cc1be6834064a1f.NewServiceAnnouncementAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewServiceUpdateMessageRequestBuilderInternal instantiates a new ServiceUpdateMessageRequestBuilder and sets the default values.
func NewServiceUpdateMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceUpdateMessageRequestBuilder) {
    m := &ServiceUpdateMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServiceUpdateMessageRequestBuilder instantiates a new ServiceUpdateMessageRequestBuilder and sets the default values.
func NewServiceUpdateMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceUpdateMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceUpdateMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) CreateDeleteRequestInformation(options *ServiceUpdateMessageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) CreateGetRequestInformation(options *ServiceUpdateMessageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) CreatePatchRequestInformation(options *ServiceUpdateMessageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) Delete(options *ServiceUpdateMessageRequestBuilderDeleteOptions)(error) {
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
// Get a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) Get(options *ServiceUpdateMessageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServiceUpdateMessage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage), nil
}
// Patch a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceUpdateMessageRequestBuilder) Patch(options *ServiceUpdateMessageRequestBuilderPatchOptions)(error) {
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
