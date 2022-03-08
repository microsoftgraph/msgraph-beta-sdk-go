package setandroiddeviceownerfullymanagedenrollmentstate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
type SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostOptions options for Post
type SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostOptions struct {
    // 
    Body SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal instantiates a new SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder and sets the default values.
func NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    m := &SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/microsoft.graph.setAndroidDeviceOwnerFullyManagedEnrollmentState";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder instantiates a new SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder and sets the default values.
func NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) CreatePostRequestInformation(options *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post sets the AndroidManagedStoreAccountEnterpriseSettings AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) Post(options *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
