package macossoftwareupdateaccountsummaries

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i489cc74bef3004f249dde5c4a556d17b17597e69965458cdcc4754179c41709b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/macossoftwareupdateaccountsummaries/count"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// MacOSSoftwareUpdateAccountSummariesRequestBuilder provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
type MacOSSoftwareUpdateAccountSummariesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MacOSSoftwareUpdateAccountSummariesRequestBuilderGetOptions options for Get
type MacOSSoftwareUpdateAccountSummariesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MacOSSoftwareUpdateAccountSummariesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MacOSSoftwareUpdateAccountSummariesRequestBuilderGetQueryParameters the MacOS software update account summaries for this account.
type MacOSSoftwareUpdateAccountSummariesRequestBuilderGetQueryParameters struct {
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
// MacOSSoftwareUpdateAccountSummariesRequestBuilderPostOptions options for Post
type MacOSSoftwareUpdateAccountSummariesRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal instantiates a new MacOSSoftwareUpdateAccountSummariesRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MacOSSoftwareUpdateAccountSummariesRequestBuilder) {
    m := &MacOSSoftwareUpdateAccountSummariesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMacOSSoftwareUpdateAccountSummariesRequestBuilder instantiates a new MacOSSoftwareUpdateAccountSummariesRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummariesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MacOSSoftwareUpdateAccountSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MacOSSoftwareUpdateAccountSummariesRequestBuilder) Count()(*i489cc74bef3004f249dde5c4a556d17b17597e69965458cdcc4754179c41709b.CountRequestBuilder) {
    return i489cc74bef3004f249dde5c4a556d17b17597e69965458cdcc4754179c41709b.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the MacOS software update account summaries for this account.
func (m *MacOSSoftwareUpdateAccountSummariesRequestBuilder) CreateGetRequestInformation(options *MacOSSoftwareUpdateAccountSummariesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *MacOSSoftwareUpdateAccountSummariesRequestBuilder) CreatePostRequestInformation(options *MacOSSoftwareUpdateAccountSummariesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the MacOS software update account summaries for this account.
func (m *MacOSSoftwareUpdateAccountSummariesRequestBuilder) Get(options *MacOSSoftwareUpdateAccountSummariesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMacOSSoftwareUpdateAccountSummaryCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryCollectionResponseable), nil
}
// Post create new navigation property to macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *MacOSSoftwareUpdateAccountSummariesRequestBuilder) Post(options *MacOSSoftwareUpdateAccountSummariesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MacOSSoftwareUpdateAccountSummaryable), nil
}
