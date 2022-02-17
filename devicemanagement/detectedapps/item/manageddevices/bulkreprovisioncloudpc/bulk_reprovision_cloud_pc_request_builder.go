package bulkreprovisioncloudpc

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// BulkReprovisionCloudPcRequestBuilder builds and executes requests for operations under \deviceManagement\detectedApps\{detectedApp-id}\managedDevices\microsoft.graph.bulkReprovisionCloudPc
type BulkReprovisionCloudPcRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BulkReprovisionCloudPcRequestBuilderPostOptions options for Post
type BulkReprovisionCloudPcRequestBuilderPostOptions struct {
    // 
    Body *BulkReprovisionCloudPcRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewBulkReprovisionCloudPcRequestBuilderInternal instantiates a new BulkReprovisionCloudPcRequestBuilder and sets the default values.
func NewBulkReprovisionCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BulkReprovisionCloudPcRequestBuilder) {
    m := &BulkReprovisionCloudPcRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/detectedApps/{detectedApp_id}/managedDevices/microsoft.graph.bulkReprovisionCloudPc";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBulkReprovisionCloudPcRequestBuilder instantiates a new BulkReprovisionCloudPcRequestBuilder and sets the default values.
func NewBulkReprovisionCloudPcRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BulkReprovisionCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBulkReprovisionCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action bulkReprovisionCloudPc
func (m *BulkReprovisionCloudPcRequestBuilder) CreatePostRequestInformation(options *BulkReprovisionCloudPcRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action bulkReprovisionCloudPc
func (m *BulkReprovisionCloudPcRequestBuilder) Post(options *BulkReprovisionCloudPcRequestBuilderPostOptions)(error) {
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
