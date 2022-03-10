package mobileapps

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i149429e3eafc16419f958b84dc22e5575cc97c96b553de6bc0dbce71d8dadd3c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/count"
    i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/haspayloadlinks"
    i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/getmobileappcountwithstatus"
    i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/gettopmobileappswithstatuswithcount"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/validatexml"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// MobileAppsRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MobileAppsRequestBuilderGetOptions options for Get
type MobileAppsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MobileAppsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MobileAppsRequestBuilderGetQueryParameters the mobile apps.
type MobileAppsRequestBuilderGetQueryParameters struct {
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
// MobileAppsRequestBuilderPostOptions options for Post
type MobileAppsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMobileAppsRequestBuilderInternal instantiates a new MobileAppsRequestBuilder and sets the default values.
func NewMobileAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppsRequestBuilder) {
    m := &MobileAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppsRequestBuilder instantiates a new MobileAppsRequestBuilder and sets the default values.
func NewMobileAppsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppsRequestBuilder) Count()(*i149429e3eafc16419f958b84dc22e5575cc97c96b553de6bc0dbce71d8dadd3c.CountRequestBuilder) {
    return i149429e3eafc16419f958b84dc22e5575cc97c96b553de6bc0dbce71d8dadd3c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the mobile apps.
func (m *MobileAppsRequestBuilder) CreateGetRequestInformation(options *MobileAppsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to mobileApps for deviceAppManagement
func (m *MobileAppsRequestBuilder) CreatePostRequestInformation(options *MobileAppsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the mobile apps.
func (m *MobileAppsRequestBuilder) Get(options *MobileAppsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMobileAppCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppCollectionResponseable), nil
}
// GetMobileAppCountWithStatus provides operations to call the getMobileAppCount method.
func (m *MobileAppsRequestBuilder) GetMobileAppCountWithStatus(status *string)(*i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.GetMobileAppCountWithStatusRequestBuilder) {
    return i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.NewGetMobileAppCountWithStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter, status);
}
// GetTopMobileAppsWithStatusWithCount provides operations to call the getTopMobileApps method.
func (m *MobileAppsRequestBuilder) GetTopMobileAppsWithStatusWithCount(count *int64, status *string)(*i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.GetTopMobileAppsWithStatusWithCountRequestBuilder) {
    return i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.NewGetTopMobileAppsWithStatusWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count, status);
}
func (m *MobileAppsRequestBuilder) HasPayloadLinks()(*i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.HasPayloadLinksRequestBuilder) {
    return i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to mobileApps for deviceAppManagement
func (m *MobileAppsRequestBuilder) Post(options *MobileAppsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMobileAppFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileAppable), nil
}
func (m *MobileAppsRequestBuilder) ValidateXml()(*id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.ValidateXmlRequestBuilder) {
    return id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.NewValidateXmlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
