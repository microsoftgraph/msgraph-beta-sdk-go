package unenrollassetsbyid

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// UnenrollAssetsByIdRequestBuilder builds and executes requests for operations under \admin\windows\updates\deployments\{deployment-id}\audience\exclusions\microsoft.graph.windowsUpdates.unenrollAssetsById
type UnenrollAssetsByIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnenrollAssetsByIdRequestBuilderPostOptions options for Post
type UnenrollAssetsByIdRequestBuilderPostOptions struct {
    // 
    Body *UnenrollAssetsByIdRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUnenrollAssetsByIdRequestBuilderInternal instantiates a new UnenrollAssetsByIdRequestBuilder and sets the default values.
func NewUnenrollAssetsByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnenrollAssetsByIdRequestBuilder) {
    m := &UnenrollAssetsByIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment_id}/audience/exclusions/microsoft.graph.windowsUpdates.unenrollAssetsById";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnenrollAssetsByIdRequestBuilder instantiates a new UnenrollAssetsByIdRequestBuilder and sets the default values.
func NewUnenrollAssetsByIdRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnenrollAssetsByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnenrollAssetsByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action unenrollAssetsById
func (m *UnenrollAssetsByIdRequestBuilder) CreatePostRequestInformation(options *UnenrollAssetsByIdRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action unenrollAssetsById
func (m *UnenrollAssetsByIdRequestBuilder) Post(options *UnenrollAssetsByIdRequestBuilderPostOptions)(error) {
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
