package wipeandblockmanagedapps

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// WipeAndBlockManagedAppsRequestBuilder builds and executes requests for operations under \users\{user-id}\microsoft.graph.wipeAndBlockManagedApps
type WipeAndBlockManagedAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WipeAndBlockManagedAppsRequestBuilderPostOptions options for Post
type WipeAndBlockManagedAppsRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWipeAndBlockManagedAppsRequestBuilderInternal instantiates a new WipeAndBlockManagedAppsRequestBuilder and sets the default values.
func NewWipeAndBlockManagedAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WipeAndBlockManagedAppsRequestBuilder) {
    m := &WipeAndBlockManagedAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.wipeAndBlockManagedApps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWipeAndBlockManagedAppsRequestBuilder instantiates a new WipeAndBlockManagedAppsRequestBuilder and sets the default values.
func NewWipeAndBlockManagedAppsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WipeAndBlockManagedAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWipeAndBlockManagedAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation blocks the managed app user from app check-in.
func (m *WipeAndBlockManagedAppsRequestBuilder) CreatePostRequestInformation(options *WipeAndBlockManagedAppsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post blocks the managed app user from app check-in.
func (m *WipeAndBlockManagedAppsRequestBuilder) Post(options *WipeAndBlockManagedAppsRequestBuilderPostOptions)(error) {
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
