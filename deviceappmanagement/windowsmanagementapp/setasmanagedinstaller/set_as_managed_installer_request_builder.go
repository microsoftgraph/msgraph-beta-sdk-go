package setasmanagedinstaller

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// SetAsManagedInstallerRequestBuilder builds and executes requests for operations under \deviceAppManagement\windowsManagementApp\microsoft.graph.setAsManagedInstaller
type SetAsManagedInstallerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetAsManagedInstallerRequestBuilderPostOptions options for Post
type SetAsManagedInstallerRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSetAsManagedInstallerRequestBuilderInternal instantiates a new SetAsManagedInstallerRequestBuilder and sets the default values.
func NewSetAsManagedInstallerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetAsManagedInstallerRequestBuilder) {
    m := &SetAsManagedInstallerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/windowsManagementApp/microsoft.graph.setAsManagedInstaller";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetAsManagedInstallerRequestBuilder instantiates a new SetAsManagedInstallerRequestBuilder and sets the default values.
func NewSetAsManagedInstallerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetAsManagedInstallerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetAsManagedInstallerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation set the Managed Installer status for the caller tenant
func (m *SetAsManagedInstallerRequestBuilder) CreatePostRequestInformation(options *SetAsManagedInstallerRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post set the Managed Installer status for the caller tenant
func (m *SetAsManagedInstallerRequestBuilder) Post(options *SetAsManagedInstallerRequestBuilderPostOptions)(error) {
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
