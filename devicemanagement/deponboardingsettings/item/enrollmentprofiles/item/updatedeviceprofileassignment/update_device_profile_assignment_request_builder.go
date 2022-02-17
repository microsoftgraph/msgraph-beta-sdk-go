package updatedeviceprofileassignment

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// UpdateDeviceProfileAssignmentRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\enrollmentProfiles\{enrollmentProfile-id}\microsoft.graph.updateDeviceProfileAssignment
type UpdateDeviceProfileAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdateDeviceProfileAssignmentRequestBuilderPostOptions options for Post
type UpdateDeviceProfileAssignmentRequestBuilderPostOptions struct {
    // 
    Body *UpdateDeviceProfileAssignmentRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUpdateDeviceProfileAssignmentRequestBuilderInternal instantiates a new UpdateDeviceProfileAssignmentRequestBuilder and sets the default values.
func NewUpdateDeviceProfileAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateDeviceProfileAssignmentRequestBuilder) {
    m := &UpdateDeviceProfileAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/enrollmentProfiles/{enrollmentProfile_id}/microsoft.graph.updateDeviceProfileAssignment";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateDeviceProfileAssignmentRequestBuilder instantiates a new UpdateDeviceProfileAssignmentRequestBuilder and sets the default values.
func NewUpdateDeviceProfileAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateDeviceProfileAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateDeviceProfileAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) CreatePostRequestInformation(options *UpdateDeviceProfileAssignmentRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action updateDeviceProfileAssignment
func (m *UpdateDeviceProfileAssignmentRequestBuilder) Post(options *UpdateDeviceProfileAssignmentRequestBuilderPostOptions)(error) {
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
