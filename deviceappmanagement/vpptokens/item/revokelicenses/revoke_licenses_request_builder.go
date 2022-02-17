package revokelicenses

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// RevokeLicensesRequestBuilder builds and executes requests for operations under \deviceAppManagement\vppTokens\{vppToken-id}\microsoft.graph.revokeLicenses
type RevokeLicensesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RevokeLicensesRequestBuilderPostOptions options for Post
type RevokeLicensesRequestBuilderPostOptions struct {
    // 
    Body *RevokeLicensesRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRevokeLicensesRequestBuilderInternal instantiates a new RevokeLicensesRequestBuilder and sets the default values.
func NewRevokeLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RevokeLicensesRequestBuilder) {
    m := &RevokeLicensesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/vppTokens/{vppToken_id}/microsoft.graph.revokeLicenses";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRevokeLicensesRequestBuilder instantiates a new RevokeLicensesRequestBuilder and sets the default values.
func NewRevokeLicensesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RevokeLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRevokeLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation revoke licenses associated with a specific appleVolumePurchaseProgramToken
func (m *RevokeLicensesRequestBuilder) CreatePostRequestInformation(options *RevokeLicensesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post revoke licenses associated with a specific appleVolumePurchaseProgramToken
func (m *RevokeLicensesRequestBuilder) Post(options *RevokeLicensesRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
