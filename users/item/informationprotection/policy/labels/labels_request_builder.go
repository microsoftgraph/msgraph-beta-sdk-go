package labels

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i815af1b5d2ba5bd7e5f98cda3daead6e66346dba70155ed163a163a6576409e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/evaluateclassificationresults"
    i899f24d803ea2e82f81f82cb89b445d890b85146bdccb8d94b936dd1f2312fd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/count"
    i8e9df7f67db474152e4a3949333d711503813269c2462b73b7c3fb524a46b0d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/evaluateremoval"
    id5310ddb400c09ed63cd8fa70c110b08804c18ed5839cdc5def5070ee8be4a53 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/extractlabel"
    ife1e0da99801d9046825648970ce58a1561136c8db63462e111d9f32f40f7682 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy/labels/evaluateapplication"
)

// LabelsRequestBuilder provides operations to manage the labels property of the microsoft.graph.informationProtectionPolicy entity.
type LabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// LabelsRequestBuilderGetOptions options for Get
type LabelsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *LabelsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// LabelsRequestBuilderGetQueryParameters get labels from users
type LabelsRequestBuilderGetQueryParameters struct {
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
// LabelsRequestBuilderPostOptions options for Post
type LabelsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabelable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LabelsRequestBuilder) {
    m := &LabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/informationProtection/policy/labels{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *LabelsRequestBuilder) Count()(*i899f24d803ea2e82f81f82cb89b445d890b85146bdccb8d94b936dd1f2312fd9.CountRequestBuilder) {
    return i899f24d803ea2e82f81f82cb89b445d890b85146bdccb8d94b936dd1f2312fd9.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get labels from users
func (m *LabelsRequestBuilder) CreateGetRequestInformation(options *LabelsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to labels for users
func (m *LabelsRequestBuilder) CreatePostRequestInformation(options *LabelsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *LabelsRequestBuilder) EvaluateApplication()(*ife1e0da99801d9046825648970ce58a1561136c8db63462e111d9f32f40f7682.EvaluateApplicationRequestBuilder) {
    return ife1e0da99801d9046825648970ce58a1561136c8db63462e111d9f32f40f7682.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) EvaluateClassificationResults()(*i815af1b5d2ba5bd7e5f98cda3daead6e66346dba70155ed163a163a6576409e5.EvaluateClassificationResultsRequestBuilder) {
    return i815af1b5d2ba5bd7e5f98cda3daead6e66346dba70155ed163a163a6576409e5.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) EvaluateRemoval()(*i8e9df7f67db474152e4a3949333d711503813269c2462b73b7c3fb524a46b0d5.EvaluateRemovalRequestBuilder) {
    return i8e9df7f67db474152e4a3949333d711503813269c2462b73b7c3fb524a46b0d5.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) ExtractLabel()(*id5310ddb400c09ed63cd8fa70c110b08804c18ed5839cdc5def5070ee8be4a53.ExtractLabelRequestBuilder) {
    return id5310ddb400c09ed63cd8fa70c110b08804c18ed5839cdc5def5070ee8be4a53.NewExtractLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get labels from users
func (m *LabelsRequestBuilder) Get(options *LabelsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabelCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateInformationProtectionLabelCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabelCollectionResponseable), nil
}
// Post create new navigation property to labels for users
func (m *LabelsRequestBuilder) Post(options *LabelsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabelable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateInformationProtectionLabelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabelable), nil
}
