package chromeosonboardingsettings

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i03210ad32b09351152234b83ec305c89c9e57ede5b483a4b11073f1d5356f1ce "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/chromeosonboardingsettings/count"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    icd47eb9dea6f9de87c7b2bffac772b07ab4e53b5bd39f33b7ea3886e5e11f322 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/chromeosonboardingsettings/disconnect"
    ie99e2dac95e812895dc09a981511e5bcfcf3d589738bbaf03993925b5ede2ada "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/chromeosonboardingsettings/connect"
)

// ChromeOSOnboardingSettingsRequestBuilder provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type ChromeOSOnboardingSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChromeOSOnboardingSettingsRequestBuilderGetOptions options for Get
type ChromeOSOnboardingSettingsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChromeOSOnboardingSettingsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChromeOSOnboardingSettingsRequestBuilderGetQueryParameters collection of ChromeOSOnboardingSettings settings associated with account.
type ChromeOSOnboardingSettingsRequestBuilderGetQueryParameters struct {
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
// ChromeOSOnboardingSettingsRequestBuilderPostOptions options for Post
type ChromeOSOnboardingSettingsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChromeOSOnboardingSettingsable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ChromeOSOnboardingSettingsRequestBuilder) Connect()(*ie99e2dac95e812895dc09a981511e5bcfcf3d589738bbaf03993925b5ede2ada.ConnectRequestBuilder) {
    return ie99e2dac95e812895dc09a981511e5bcfcf3d589738bbaf03993925b5ede2ada.NewConnectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChromeOSOnboardingSettingsRequestBuilderInternal instantiates a new ChromeOSOnboardingSettingsRequestBuilder and sets the default values.
func NewChromeOSOnboardingSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChromeOSOnboardingSettingsRequestBuilder) {
    m := &ChromeOSOnboardingSettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/chromeOSOnboardingSettings{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChromeOSOnboardingSettingsRequestBuilder instantiates a new ChromeOSOnboardingSettingsRequestBuilder and sets the default values.
func NewChromeOSOnboardingSettingsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChromeOSOnboardingSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChromeOSOnboardingSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ChromeOSOnboardingSettingsRequestBuilder) Count()(*i03210ad32b09351152234b83ec305c89c9e57ede5b483a4b11073f1d5356f1ce.CountRequestBuilder) {
    return i03210ad32b09351152234b83ec305c89c9e57ede5b483a4b11073f1d5356f1ce.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of ChromeOSOnboardingSettings settings associated with account.
func (m *ChromeOSOnboardingSettingsRequestBuilder) CreateGetRequestInformation(options *ChromeOSOnboardingSettingsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to chromeOSOnboardingSettings for deviceManagement
func (m *ChromeOSOnboardingSettingsRequestBuilder) CreatePostRequestInformation(options *ChromeOSOnboardingSettingsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ChromeOSOnboardingSettingsRequestBuilder) Disconnect()(*icd47eb9dea6f9de87c7b2bffac772b07ab4e53b5bd39f33b7ea3886e5e11f322.DisconnectRequestBuilder) {
    return icd47eb9dea6f9de87c7b2bffac772b07ab4e53b5bd39f33b7ea3886e5e11f322.NewDisconnectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get collection of ChromeOSOnboardingSettings settings associated with account.
func (m *ChromeOSOnboardingSettingsRequestBuilder) Get(options *ChromeOSOnboardingSettingsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChromeOSOnboardingSettingsCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateChromeOSOnboardingSettingsCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChromeOSOnboardingSettingsCollectionResponseable), nil
}
// Post create new navigation property to chromeOSOnboardingSettings for deviceManagement
func (m *ChromeOSOnboardingSettingsRequestBuilder) Post(options *ChromeOSOnboardingSettingsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChromeOSOnboardingSettingsable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateChromeOSOnboardingSettingsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChromeOSOnboardingSettingsable), nil
}
