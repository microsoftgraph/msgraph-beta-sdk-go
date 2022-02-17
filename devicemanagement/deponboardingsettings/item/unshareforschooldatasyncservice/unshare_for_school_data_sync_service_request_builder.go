package unshareforschooldatasyncservice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// UnshareForSchoolDataSyncServiceRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\microsoft.graph.unshareForSchoolDataSyncService
type UnshareForSchoolDataSyncServiceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnshareForSchoolDataSyncServiceRequestBuilderPostOptions options for Post
type UnshareForSchoolDataSyncServiceRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUnshareForSchoolDataSyncServiceRequestBuilderInternal instantiates a new UnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewUnshareForSchoolDataSyncServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnshareForSchoolDataSyncServiceRequestBuilder) {
    m := &UnshareForSchoolDataSyncServiceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/microsoft.graph.unshareForSchoolDataSyncService";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnshareForSchoolDataSyncServiceRequestBuilder instantiates a new UnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewUnshareForSchoolDataSyncServiceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnshareForSchoolDataSyncServiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnshareForSchoolDataSyncServiceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action unshareForSchoolDataSyncService
func (m *UnshareForSchoolDataSyncServiceRequestBuilder) CreatePostRequestInformation(options *UnshareForSchoolDataSyncServiceRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action unshareForSchoolDataSyncService
func (m *UnshareForSchoolDataSyncServiceRequestBuilder) Post(options *UnshareForSchoolDataSyncServiceRequestBuilderPostOptions)(error) {
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
