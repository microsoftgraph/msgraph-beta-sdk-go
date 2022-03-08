package windowsinformationprotectionwipeactions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i36afe53cb0af59d3fc1458509022deef8945baf7304b3c51b9c1537d11a9dbce "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsinformationprotectionwipeactions/count"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// WindowsInformationProtectionWipeActionsRequestBuilder provides operations to manage the windowsInformationProtectionWipeActions property of the microsoft.graph.deviceAppManagement entity.
type WindowsInformationProtectionWipeActionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsInformationProtectionWipeActionsRequestBuilderGetOptions options for Get
type WindowsInformationProtectionWipeActionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsInformationProtectionWipeActionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsInformationProtectionWipeActionsRequestBuilderGetQueryParameters windows information protection wipe actions.
type WindowsInformationProtectionWipeActionsRequestBuilderGetQueryParameters struct {
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
// WindowsInformationProtectionWipeActionsRequestBuilderPostOptions options for Post
type WindowsInformationProtectionWipeActionsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsInformationProtectionWipeActionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWindowsInformationProtectionWipeActionsRequestBuilderInternal instantiates a new WindowsInformationProtectionWipeActionsRequestBuilder and sets the default values.
func NewWindowsInformationProtectionWipeActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionWipeActionsRequestBuilder) {
    m := &WindowsInformationProtectionWipeActionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/windowsInformationProtectionWipeActions{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsInformationProtectionWipeActionsRequestBuilder instantiates a new WindowsInformationProtectionWipeActionsRequestBuilder and sets the default values.
func NewWindowsInformationProtectionWipeActionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionWipeActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionWipeActionsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WindowsInformationProtectionWipeActionsRequestBuilder) Count()(*i36afe53cb0af59d3fc1458509022deef8945baf7304b3c51b9c1537d11a9dbce.CountRequestBuilder) {
    return i36afe53cb0af59d3fc1458509022deef8945baf7304b3c51b9c1537d11a9dbce.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation windows information protection wipe actions.
func (m *WindowsInformationProtectionWipeActionsRequestBuilder) CreateGetRequestInformation(options *WindowsInformationProtectionWipeActionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to windowsInformationProtectionWipeActions for deviceAppManagement
func (m *WindowsInformationProtectionWipeActionsRequestBuilder) CreatePostRequestInformation(options *WindowsInformationProtectionWipeActionsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get windows information protection wipe actions.
func (m *WindowsInformationProtectionWipeActionsRequestBuilder) Get(options *WindowsInformationProtectionWipeActionsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsInformationProtectionWipeActionCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateWindowsInformationProtectionWipeActionCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsInformationProtectionWipeActionCollectionResponseable), nil
}
// Post create new navigation property to windowsInformationProtectionWipeActions for deviceAppManagement
func (m *WindowsInformationProtectionWipeActionsRequestBuilder) Post(options *WindowsInformationProtectionWipeActionsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsInformationProtectionWipeActionable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsInformationProtectionWipeActionable), nil
}
