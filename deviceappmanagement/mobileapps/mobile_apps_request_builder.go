package mobileapps

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/haspayloadlinks"
    i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/getmobileappcountwithstatus"
    i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/gettopmobileappswithstatuswithcount"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/validatexml"
)

// MobileAppsRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps
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
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp;
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
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppsRequestBuilder instantiates a new MobileAppsRequestBuilder and sets the default values.
func NewMobileAppsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsRequestBuilderInternal(urlParams, requestAdapter)
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
// CreatePostRequestInformation the mobile apps.
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
func (m *MobileAppsRequestBuilder) Get(options *MobileAppsRequestBuilderGetOptions)(*MobileAppsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*MobileAppsResponse), nil
}
// GetMobileAppCountWithStatus builds and executes requests for operations under \deviceAppManagement\mobileApps\microsoft.graph.getMobileAppCount(status='{status}')
func (m *MobileAppsRequestBuilder) GetMobileAppCountWithStatus(status *string)(*i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.GetMobileAppCountWithStatusRequestBuilder) {
    return i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.NewGetMobileAppCountWithStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter, status);
}
// GetTopMobileAppsWithStatusWithCount builds and executes requests for operations under \deviceAppManagement\mobileApps\microsoft.graph.getTopMobileApps(status='{status}',count={count})
func (m *MobileAppsRequestBuilder) GetTopMobileAppsWithStatusWithCount(status *string, count *int64)(*i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.GetTopMobileAppsWithStatusWithCountRequestBuilder) {
    return i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.NewGetTopMobileAppsWithStatusWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, status, count);
}
func (m *MobileAppsRequestBuilder) HasPayloadLinks()(*i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.HasPayloadLinksRequestBuilder) {
    return i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the mobile apps.
func (m *MobileAppsRequestBuilder) Post(options *MobileAppsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMobileApp() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp), nil
}
func (m *MobileAppsRequestBuilder) ValidateXml()(*id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.ValidateXmlRequestBuilder) {
    return id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.NewValidateXmlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
