package updaterelationships

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// UpdateRelationshipsRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\userStatuses\{userAppInstallStatus-id}\deviceStatuses\{mobileAppInstallStatus-id}\app\microsoft.graph.updateRelationships
type UpdateRelationshipsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdateRelationshipsRequestBuilderPostOptions options for Post
type UpdateRelationshipsRequestBuilderPostOptions struct {
    // 
    Body *UpdateRelationshipsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUpdateRelationshipsRequestBuilderInternal instantiates a new UpdateRelationshipsRequestBuilder and sets the default values.
func NewUpdateRelationshipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateRelationshipsRequestBuilder) {
    m := &UpdateRelationshipsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/userStatuses/{userAppInstallStatus_id}/deviceStatuses/{mobileAppInstallStatus_id}/app/microsoft.graph.updateRelationships";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateRelationshipsRequestBuilder instantiates a new UpdateRelationshipsRequestBuilder and sets the default values.
func NewUpdateRelationshipsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateRelationshipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateRelationshipsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateRelationships
func (m *UpdateRelationshipsRequestBuilder) CreatePostRequestInformation(options *UpdateRelationshipsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action updateRelationships
func (m *UpdateRelationshipsRequestBuilder) Post(options *UpdateRelationshipsRequestBuilderPostOptions)(error) {
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
