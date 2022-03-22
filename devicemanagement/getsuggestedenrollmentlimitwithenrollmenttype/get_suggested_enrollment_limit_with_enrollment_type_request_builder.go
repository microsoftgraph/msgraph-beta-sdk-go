package getsuggestedenrollmentlimitwithenrollmenttype

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder provides operations to call the getSuggestedEnrollmentLimit method.
type GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions options for Get
type GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal instantiates a new GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder and sets the default values.
func NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, enrollmentType *string)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    m := &GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getSuggestedEnrollmentLimit(enrollmentType='{enrollmentType}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if enrollmentType != nil {
        urlTplParams["enrollmentType"] = *enrollmentType
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder instantiates a new GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder and sets the default values.
func NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) CreateGetRequestInformation(options *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) Get(options *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SuggestedEnrollmentLimitable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateSuggestedEnrollmentLimitFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SuggestedEnrollmentLimitable), nil
}
