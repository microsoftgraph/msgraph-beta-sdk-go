package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i509469ad7d7d1bfb7c44fead3987e6c904f9e05130da83eef6fdcdcc595fadc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileapptroubleshootingevents/item/applogcollectionrequests"
    i59be9a91cd25f6b0023adba1357e480bf0af59b0340aa2e02c3ed40c236d28c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileapptroubleshootingevents/item/applogcollectionrequests/item"
)

// MobileAppTroubleshootingEventRequestBuilder builds and executes requests for operations under \me\mobileAppTroubleshootingEvents\{mobileAppTroubleshootingEvent-id}
type MobileAppTroubleshootingEventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MobileAppTroubleshootingEventRequestBuilderDeleteOptions options for Delete
type MobileAppTroubleshootingEventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MobileAppTroubleshootingEventRequestBuilderGetOptions options for Get
type MobileAppTroubleshootingEventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MobileAppTroubleshootingEventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MobileAppTroubleshootingEventRequestBuilderGetQueryParameters the list of mobile app troubleshooting events for this user.
type MobileAppTroubleshootingEventRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MobileAppTroubleshootingEventRequestBuilderPatchOptions options for Patch
type MobileAppTroubleshootingEventRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppTroubleshootingEvent;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MobileAppTroubleshootingEventRequestBuilder) AppLogCollectionRequests()(*i509469ad7d7d1bfb7c44fead3987e6c904f9e05130da83eef6fdcdcc595fadc2.AppLogCollectionRequestsRequestBuilder) {
    return i509469ad7d7d1bfb7c44fead3987e6c904f9e05130da83eef6fdcdcc595fadc2.NewAppLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppLogCollectionRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mobileAppTroubleshootingEvents.item.appLogCollectionRequests.item collection
func (m *MobileAppTroubleshootingEventRequestBuilder) AppLogCollectionRequestsById(id string)(*i59be9a91cd25f6b0023adba1357e480bf0af59b0340aa2e02c3ed40c236d28c4.AppLogCollectionRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appLogCollectionRequest_id"] = id
    }
    return i59be9a91cd25f6b0023adba1357e480bf0af59b0340aa2e02c3ed40c236d28c4.NewAppLogCollectionRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMobileAppTroubleshootingEventRequestBuilderInternal instantiates a new MobileAppTroubleshootingEventRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppTroubleshootingEventRequestBuilder) {
    m := &MobileAppTroubleshootingEventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppTroubleshootingEventRequestBuilder instantiates a new MobileAppTroubleshootingEventRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppTroubleshootingEventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppTroubleshootingEventRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of mobile app troubleshooting events for this user.
func (m *MobileAppTroubleshootingEventRequestBuilder) CreateDeleteRequestInformation(options *MobileAppTroubleshootingEventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of mobile app troubleshooting events for this user.
func (m *MobileAppTroubleshootingEventRequestBuilder) CreateGetRequestInformation(options *MobileAppTroubleshootingEventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of mobile app troubleshooting events for this user.
func (m *MobileAppTroubleshootingEventRequestBuilder) CreatePatchRequestInformation(options *MobileAppTroubleshootingEventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of mobile app troubleshooting events for this user.
func (m *MobileAppTroubleshootingEventRequestBuilder) Delete(options *MobileAppTroubleshootingEventRequestBuilderDeleteOptions)(error) {
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
// Get the list of mobile app troubleshooting events for this user.
func (m *MobileAppTroubleshootingEventRequestBuilder) Get(options *MobileAppTroubleshootingEventRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppTroubleshootingEvent, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMobileAppTroubleshootingEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppTroubleshootingEvent), nil
}
// Patch the list of mobile app troubleshooting events for this user.
func (m *MobileAppTroubleshootingEventRequestBuilder) Patch(options *MobileAppTroubleshootingEventRequestBuilderPatchOptions)(error) {
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
