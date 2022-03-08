package assignedaccessmultimodeprofiles

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// AssignedAccessMultiModeProfilesRequestBuilder provides operations to call the assignedAccessMultiModeProfiles method.
type AssignedAccessMultiModeProfilesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AssignedAccessMultiModeProfilesRequestBuilderPostOptions options for Post
type AssignedAccessMultiModeProfilesRequestBuilderPostOptions struct {
    // 
    Body AssignedAccessMultiModeProfilesRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAssignedAccessMultiModeProfilesRequestBuilderInternal instantiates a new AssignedAccessMultiModeProfilesRequestBuilder and sets the default values.
func NewAssignedAccessMultiModeProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AssignedAccessMultiModeProfilesRequestBuilder) {
    m := &AssignedAccessMultiModeProfilesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/microsoft.graph.assignedAccessMultiModeProfiles";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAssignedAccessMultiModeProfilesRequestBuilder instantiates a new AssignedAccessMultiModeProfilesRequestBuilder and sets the default values.
func NewAssignedAccessMultiModeProfilesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AssignedAccessMultiModeProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignedAccessMultiModeProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action assignedAccessMultiModeProfiles
func (m *AssignedAccessMultiModeProfilesRequestBuilder) CreatePostRequestInformation(options *AssignedAccessMultiModeProfilesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action assignedAccessMultiModeProfiles
func (m *AssignedAccessMultiModeProfilesRequestBuilder) Post(options *AssignedAccessMultiModeProfilesRequestBuilderPostOptions)(error) {
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
