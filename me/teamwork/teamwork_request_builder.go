package teamwork

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/sendactivitynotification"
    i200af14bc8a01d318341d7ffd4bd69adb3b9b97703c97f7fd8c69b8f1dc415df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/associatedteams"
    i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/installedapps"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/installedapps/item"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i884fd4ffb36d199e7dec37be072607e2769510674770c99847363a2f72615545 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/associatedteams/item"
)

// TeamworkRequestBuilder provides operations to manage the teamwork property of the microsoft.graph.user entity.
type TeamworkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamworkRequestBuilderDeleteOptions options for Delete
type TeamworkRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamworkRequestBuilderGetOptions options for Get
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
// TeamworkRequestBuilderGetQueryParameters a container for Microsoft Teams features available for the user. Read-only. Nullable.
type TeamworkRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamworkRequestBuilderPatchOptions options for Patch
type TeamworkRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserTeamworkable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TeamworkRequestBuilder) AssociatedTeams()(*i200af14bc8a01d318341d7ffd4bd69adb3b9b97703c97f7fd8c69b8f1dc415df.AssociatedTeamsRequestBuilder) {
    return i200af14bc8a01d318341d7ffd4bd69adb3b9b97703c97f7fd8c69b8f1dc415df.NewAssociatedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssociatedTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.teamwork.associatedTeams.item collection
func (m *TeamworkRequestBuilder) AssociatedTeamsById(id string)(*i884fd4ffb36d199e7dec37be072607e2769510674770c99847363a2f72615545.AssociatedTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["associatedTeamInfo_id"] = id
    }
    return i884fd4ffb36d199e7dec37be072607e2769510674770c99847363a2f72615545.NewAssociatedTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamworkRequestBuilderInternal instantiates a new TeamworkRequestBuilder and sets the default values.
func NewTeamworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamworkRequestBuilder) {
    m := &TeamworkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/teamwork{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkRequestBuilder instantiates a new TeamworkRequestBuilder and sets the default values.
func NewTeamworkRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamworkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property teamwork for me
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
// CreateGetRequestInformation a container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *TeamworkRequestBuilder) CreateGetRequestInformation(options *TeamworkRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
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
// CreatePatchRequestInformation update the navigation property teamwork in me
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
// Delete delete navigation property teamwork for me
func (m *TeamworkRequestBuilder) Delete(options *TeamworkRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *TeamworkRequestBuilder) Get(options *TeamworkRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserTeamworkable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserTeamworkFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserTeamworkable), nil
}
func (m *TeamworkRequestBuilder) InstalledApps()(*i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e.InstalledAppsRequestBuilder) {
    return i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.teamwork.installedApps.item collection
func (m *TeamworkRequestBuilder) InstalledAppsById(id string)(*i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b.UserScopeTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userScopeTeamsAppInstallation_id"] = id
    }
    return i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b.NewUserScopeTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamwork in me
func (m *TeamworkRequestBuilder) Patch(options *TeamworkRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamworkRequestBuilder) SendActivityNotification()(*i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1.SendActivityNotificationRequestBuilder) {
    return i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
