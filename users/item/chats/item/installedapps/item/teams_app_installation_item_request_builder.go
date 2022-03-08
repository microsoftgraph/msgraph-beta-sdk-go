package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia4904963dfe5c3efffef939d90da7785720ac0ba5f00225b9f28bebbcbbd6155 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/installedapps/item/upgrade"
    ida05b22008dc0d759522d849950617f5bf99f982dad6c7ea22082336178b44de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/installedapps/item/teamsapp"
    iee6ca560ef39b2c7bb0e76e9b8e00c261db6ee2e9f434a5ab6f7fa307f7dac1b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/installedapps/item/teamsappdefinition"
)

// TeamsAppInstallationItemRequestBuilder provides operations to manage the installedApps property of the microsoft.graph.chat entity.
type TeamsAppInstallationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamsAppInstallationItemRequestBuilderDeleteOptions options for Delete
type TeamsAppInstallationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamsAppInstallationItemRequestBuilderGetOptions options for Get
type TeamsAppInstallationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamsAppInstallationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamsAppInstallationItemRequestBuilderGetQueryParameters a collection of all the apps in the chat. Nullable.
type TeamsAppInstallationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamsAppInstallationItemRequestBuilderPatchOptions options for Patch
type TeamsAppInstallationItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamsAppInstallationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTeamsAppInstallationItemRequestBuilderInternal instantiates a new TeamsAppInstallationItemRequestBuilder and sets the default values.
func NewTeamsAppInstallationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationItemRequestBuilder) {
    m := &TeamsAppInstallationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/chats/{chat_id}/installedApps/{teamsAppInstallation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamsAppInstallationItemRequestBuilder instantiates a new TeamsAppInstallationItemRequestBuilder and sets the default values.
func NewTeamsAppInstallationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppInstallationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property installedApps for users
func (m *TeamsAppInstallationItemRequestBuilder) CreateDeleteRequestInformation(options *TeamsAppInstallationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationItemRequestBuilder) CreateGetRequestInformation(options *TeamsAppInstallationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property installedApps in users
func (m *TeamsAppInstallationItemRequestBuilder) CreatePatchRequestInformation(options *TeamsAppInstallationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property installedApps for users
func (m *TeamsAppInstallationItemRequestBuilder) Delete(options *TeamsAppInstallationItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationItemRequestBuilder) Get(options *TeamsAppInstallationItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamsAppInstallationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateTeamsAppInstallationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamsAppInstallationable), nil
}
// Patch update the navigation property installedApps in users
func (m *TeamsAppInstallationItemRequestBuilder) Patch(options *TeamsAppInstallationItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamsAppInstallationItemRequestBuilder) TeamsApp()(*ida05b22008dc0d759522d849950617f5bf99f982dad6c7ea22082336178b44de.TeamsAppRequestBuilder) {
    return ida05b22008dc0d759522d849950617f5bf99f982dad6c7ea22082336178b44de.NewTeamsAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationItemRequestBuilder) TeamsAppDefinition()(*iee6ca560ef39b2c7bb0e76e9b8e00c261db6ee2e9f434a5ab6f7fa307f7dac1b.TeamsAppDefinitionRequestBuilder) {
    return iee6ca560ef39b2c7bb0e76e9b8e00c261db6ee2e9f434a5ab6f7fa307f7dac1b.NewTeamsAppDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationItemRequestBuilder) Upgrade()(*ia4904963dfe5c3efffef939d90da7785720ac0ba5f00225b9f28bebbcbbd6155.UpgradeRequestBuilder) {
    return ia4904963dfe5c3efffef939d90da7785720ac0ba5f00225b9f28bebbcbbd6155.NewUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
