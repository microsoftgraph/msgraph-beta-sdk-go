package messages

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/unfavorite"
    i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/favorite"
    i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/markunread"
    ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/markread"
    icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/archive"
    icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/unarchive"
)

// MessagesRequestBuilder builds and executes requests for operations under \admin\serviceAnnouncement\messages
type MessagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessagesRequestBuilderGetOptions options for Get
type MessagesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MessagesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessagesRequestBuilderGetQueryParameters a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
type MessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// MessagesRequestBuilderPostOptions options for Post
type MessagesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MessagesRequestBuilder) Archive()(*icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208.ArchiveRequestBuilder) {
    return icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessagesRequestBuilderInternal instantiates a new MessagesRequestBuilder and sets the default values.
func NewMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessagesRequestBuilder) {
    m := &MessagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessagesRequestBuilder instantiates a new MessagesRequestBuilder and sets the default values.
func NewMessagesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *MessagesRequestBuilder) CreateGetRequestInformation(options *MessagesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *MessagesRequestBuilder) CreatePostRequestInformation(options *MessagesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *MessagesRequestBuilder) Favorite()(*i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb.FavoriteRequestBuilder) {
    return i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb.NewFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *MessagesRequestBuilder) Get(options *MessagesRequestBuilderGetOptions)(*MessagesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessagesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*MessagesResponse), nil
}
func (m *MessagesRequestBuilder) MarkRead()(*ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e.MarkReadRequestBuilder) {
    return ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e.NewMarkReadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) MarkUnread()(*i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83.MarkUnreadRequestBuilder) {
    return i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83.NewMarkUnreadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post a collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *MessagesRequestBuilder) Post(options *MessagesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServiceUpdateMessage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage), nil
}
func (m *MessagesRequestBuilder) Unarchive()(*icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c.UnarchiveRequestBuilder) {
    return icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) Unfavorite()(*i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112.UnfavoriteRequestBuilder) {
    return i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112.NewUnfavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
