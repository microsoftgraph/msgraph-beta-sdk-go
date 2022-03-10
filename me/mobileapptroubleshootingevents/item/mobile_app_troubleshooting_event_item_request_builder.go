package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i509469ad7d7d1bfb7c44fead3987e6c904f9e05130da83eef6fdcdcc595fadc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileapptroubleshootingevents/item/applogcollectionrequests"
    i59be9a91cd25f6b0023adba1357e480bf0af59b0340aa2e02c3ed40c236d28c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileapptroubleshootingevents/item/applogcollectionrequests/item"
)

// MobileAppTroubleshootingEventItemRequestBuilder provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.user entity.
type MobileAppTroubleshootingEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MobileAppTroubleshootingEventItemRequestBuilderDeleteOptions options for Delete
type MobileAppTroubleshootingEventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MobileAppTroubleshootingEventItemRequestBuilderGetOptions options for Get
type MobileAppTroubleshootingEventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters the list of mobile app troubleshooting events for this user.
type MobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MobileAppTroubleshootingEventItemRequestBuilderPatchOptions options for Patch
type MobileAppTroubleshootingEventItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppTroubleshootingEventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MobileAppTroubleshootingEventItemRequestBuilder) AppLogCollectionRequests()(*i509469ad7d7d1bfb7c44fead3987e6c904f9e05130da83eef6fdcdcc595fadc2.AppLogCollectionRequestsRequestBuilder) {
    return i509469ad7d7d1bfb7c44fead3987e6c904f9e05130da83eef6fdcdcc595fadc2.NewAppLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppLogCollectionRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mobileAppTroubleshootingEvents.item.appLogCollectionRequests.item collection
func (m *MobileAppTroubleshootingEventItemRequestBuilder) AppLogCollectionRequestsById(id string)(*i59be9a91cd25f6b0023adba1357e480bf0af59b0340aa2e02c3ed40c236d28c4.AppLogCollectionRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appLogCollectionRequest_id"] = id
    }
    return i59be9a91cd25f6b0023adba1357e480bf0af59b0340aa2e02c3ed40c236d28c4.NewAppLogCollectionRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMobileAppTroubleshootingEventItemRequestBuilderInternal instantiates a new MobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppTroubleshootingEventItemRequestBuilder) {
    m := &MobileAppTroubleshootingEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppTroubleshootingEventItemRequestBuilder instantiates a new MobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppTroubleshootingEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppTroubleshootingEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mobileAppTroubleshootingEvents for me
func (m *MobileAppTroubleshootingEventItemRequestBuilder) CreateDeleteRequestInformation(options *MobileAppTroubleshootingEventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MobileAppTroubleshootingEventItemRequestBuilder) CreateGetRequestInformation(options *MobileAppTroubleshootingEventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mobileAppTroubleshootingEvents in me
func (m *MobileAppTroubleshootingEventItemRequestBuilder) CreatePatchRequestInformation(options *MobileAppTroubleshootingEventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property mobileAppTroubleshootingEvents for me
func (m *MobileAppTroubleshootingEventItemRequestBuilder) Delete(options *MobileAppTroubleshootingEventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of mobile app troubleshooting events for this user.
func (m *MobileAppTroubleshootingEventItemRequestBuilder) Get(options *MobileAppTroubleshootingEventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppTroubleshootingEventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMobileAppTroubleshootingEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppTroubleshootingEventable), nil
}
// Patch update the navigation property mobileAppTroubleshootingEvents in me
func (m *MobileAppTroubleshootingEventItemRequestBuilder) Patch(options *MobileAppTroubleshootingEventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
