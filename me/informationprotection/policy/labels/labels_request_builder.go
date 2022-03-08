package labels

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i56b3db10f72adbf23f88e78a9f44a5e425d78be8489a34c7f798cae7c1ed2b28 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/policy/labels/evaluateclassificationresults"
    iabff15acfeb1cae45775785ffb30c23869c92518718d8e058fafb9594de6abc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/policy/labels/extractlabel"
    ib470781a8a5e1b02a13e8da61551165f143f7bd30847d6dd93ebd2668e16371b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/policy/labels/count"
    if8658242e2270cf053c472c3bb976bc701b68a88e262fe82aaf480690250f5e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/policy/labels/evaluateapplication"
    ifbb3e1e672049193601dce7ec1938089746f914d97e177f929412bc53af934ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/policy/labels/evaluateremoval"
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
// LabelsRequestBuilderGetQueryParameters get labels from me
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
    m.urlTemplate = "{+baseurl}/me/informationProtection/policy/labels{?top,skip,search,filter,count,orderby,select,expand}";
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
func (m *LabelsRequestBuilder) Count()(*ib470781a8a5e1b02a13e8da61551165f143f7bd30847d6dd93ebd2668e16371b.CountRequestBuilder) {
    return ib470781a8a5e1b02a13e8da61551165f143f7bd30847d6dd93ebd2668e16371b.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get labels from me
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
// CreatePostRequestInformation create new navigation property to labels for me
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
func (m *LabelsRequestBuilder) EvaluateApplication()(*if8658242e2270cf053c472c3bb976bc701b68a88e262fe82aaf480690250f5e7.EvaluateApplicationRequestBuilder) {
    return if8658242e2270cf053c472c3bb976bc701b68a88e262fe82aaf480690250f5e7.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) EvaluateClassificationResults()(*i56b3db10f72adbf23f88e78a9f44a5e425d78be8489a34c7f798cae7c1ed2b28.EvaluateClassificationResultsRequestBuilder) {
    return i56b3db10f72adbf23f88e78a9f44a5e425d78be8489a34c7f798cae7c1ed2b28.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) EvaluateRemoval()(*ifbb3e1e672049193601dce7ec1938089746f914d97e177f929412bc53af934ad.EvaluateRemovalRequestBuilder) {
    return ifbb3e1e672049193601dce7ec1938089746f914d97e177f929412bc53af934ad.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) ExtractLabel()(*iabff15acfeb1cae45775785ffb30c23869c92518718d8e058fafb9594de6abc7.ExtractLabelRequestBuilder) {
    return iabff15acfeb1cae45775785ffb30c23869c92518718d8e058fafb9594de6abc7.NewExtractLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get labels from me
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
// Post create new navigation property to labels for me
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
