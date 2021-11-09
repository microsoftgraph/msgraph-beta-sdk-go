package teamwork

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/sendactivitynotification"
    i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/installedapps"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/installedapps/item"
)

// Builds and executes requests for operations under \me\teamwork
type TeamworkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type TeamworkRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type TeamworkRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamworkRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// A container for Microsoft Teams features available for the user. Read-only. Nullable.
type TeamworkRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type TeamworkRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserTeamwork;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new TeamworkRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTeamworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamworkRequestBuilder) {
    m := &TeamworkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/teamwork{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TeamworkRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTeamworkRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamworkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkRequestBuilderInternal(urlParams, requestAdapter)
}
// A container for Microsoft Teams features available for the user. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *TeamworkRequestBuilder) CreateDeleteRequestInformation(options *TeamworkRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// A container for Microsoft Teams features available for the user. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *TeamworkRequestBuilder) CreateGetRequestInformation(options *TeamworkRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
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
// A container for Microsoft Teams features available for the user. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *TeamworkRequestBuilder) CreatePatchRequestInformation(options *TeamworkRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// A container for Microsoft Teams features available for the user. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *TeamworkRequestBuilder) Delete(options *TeamworkRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// A container for Microsoft Teams features available for the user. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *TeamworkRequestBuilder) Get(options *TeamworkRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserTeamwork, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserTeamwork() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserTeamwork), nil
}
func (m *TeamworkRequestBuilder) InstalledApps()(*i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e.InstalledAppsRequestBuilder) {
    return i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.teamwork.installedApps.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *TeamworkRequestBuilder) InstalledAppsById(id string)(*i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b.UserScopeTeamsAppInstallationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userScopeTeamsAppInstallation_id"] = id
    }
    return i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b.NewUserScopeTeamsAppInstallationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// A container for Microsoft Teams features available for the user. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *TeamworkRequestBuilder) Patch(options *TeamworkRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamworkRequestBuilder) SendActivityNotification()(*i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1.SendActivityNotificationRequestBuilder) {
    return i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
