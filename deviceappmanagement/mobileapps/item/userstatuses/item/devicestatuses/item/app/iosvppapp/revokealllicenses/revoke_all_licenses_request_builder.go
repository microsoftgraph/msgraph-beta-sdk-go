package revokealllicenses

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// RevokeAllLicensesRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\userStatuses\{userAppInstallStatus-id}\deviceStatuses\{mobileAppInstallStatus-id}\app\microsoft.graph.iosVppApp\microsoft.graph.revokeAllLicenses
type RevokeAllLicensesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RevokeAllLicensesRequestBuilderPostOptions options for Post
type RevokeAllLicensesRequestBuilderPostOptions struct {
    // 
    Body *RevokeAllLicensesRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRevokeAllLicensesRequestBuilderInternal instantiates a new RevokeAllLicensesRequestBuilder and sets the default values.
func NewRevokeAllLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RevokeAllLicensesRequestBuilder) {
    m := &RevokeAllLicensesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/userStatuses/{userAppInstallStatus_id}/deviceStatuses/{mobileAppInstallStatus_id}/app/microsoft.graph.iosVppApp/microsoft.graph.revokeAllLicenses";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRevokeAllLicensesRequestBuilder instantiates a new RevokeAllLicensesRequestBuilder and sets the default values.
func NewRevokeAllLicensesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RevokeAllLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRevokeAllLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation revoke all assigned iOS VPP licenses for given app.
func (m *RevokeAllLicensesRequestBuilder) CreatePostRequestInformation(options *RevokeAllLicensesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post revoke all assigned iOS VPP licenses for given app.
func (m *RevokeAllLicensesRequestBuilder) Post(options *RevokeAllLicensesRequestBuilderPostOptions)(error) {
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
