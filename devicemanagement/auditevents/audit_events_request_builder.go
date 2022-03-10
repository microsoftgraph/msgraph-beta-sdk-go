package auditevents

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i030bbe38f2f5d2d8a0492b3f1370fddd253f4eb643fe50f74fdfdfcd967a1a8d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/auditevents/count"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i67727b0f27c2d57d41b9f444df8b33c187a7e9b7d4bf033012802d4cd0bc0365 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/auditevents/getauditactivitytypeswithcategory"
    ib4e5c6a7c7551cb9685aa79efe44df0077ee2b23b90f1c51ae212de08151e065 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/auditevents/getauditcategories"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// AuditEventsRequestBuilder provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
type AuditEventsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AuditEventsRequestBuilderGetOptions options for Get
type AuditEventsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuditEventsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuditEventsRequestBuilderGetQueryParameters the Audit Events
type AuditEventsRequestBuilderGetQueryParameters struct {
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
// AuditEventsRequestBuilderPostOptions options for Post
type AuditEventsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditEventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAuditEventsRequestBuilderInternal instantiates a new AuditEventsRequestBuilder and sets the default values.
func NewAuditEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditEventsRequestBuilder) {
    m := &AuditEventsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/auditEvents{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuditEventsRequestBuilder instantiates a new AuditEventsRequestBuilder and sets the default values.
func NewAuditEventsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditEventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditEventsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AuditEventsRequestBuilder) Count()(*i030bbe38f2f5d2d8a0492b3f1370fddd253f4eb643fe50f74fdfdfcd967a1a8d.CountRequestBuilder) {
    return i030bbe38f2f5d2d8a0492b3f1370fddd253f4eb643fe50f74fdfdfcd967a1a8d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the Audit Events
func (m *AuditEventsRequestBuilder) CreateGetRequestInformation(options *AuditEventsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to auditEvents for deviceManagement
func (m *AuditEventsRequestBuilder) CreatePostRequestInformation(options *AuditEventsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the Audit Events
func (m *AuditEventsRequestBuilder) Get(options *AuditEventsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditEventCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAuditEventCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditEventCollectionResponseable), nil
}
// GetAuditActivityTypesWithCategory provides operations to call the getAuditActivityTypes method.
func (m *AuditEventsRequestBuilder) GetAuditActivityTypesWithCategory(category *string)(*i67727b0f27c2d57d41b9f444df8b33c187a7e9b7d4bf033012802d4cd0bc0365.GetAuditActivityTypesWithCategoryRequestBuilder) {
    return i67727b0f27c2d57d41b9f444df8b33c187a7e9b7d4bf033012802d4cd0bc0365.NewGetAuditActivityTypesWithCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter, category);
}
// GetAuditCategories provides operations to call the getAuditCategories method.
func (m *AuditEventsRequestBuilder) GetAuditCategories()(*ib4e5c6a7c7551cb9685aa79efe44df0077ee2b23b90f1c51ae212de08151e065.GetAuditCategoriesRequestBuilder) {
    return ib4e5c6a7c7551cb9685aa79efe44df0077ee2b23b90f1c51ae212de08151e065.NewGetAuditCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to auditEvents for deviceManagement
func (m *AuditEventsRequestBuilder) Post(options *AuditEventsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditEventable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAuditEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditEventable), nil
}
