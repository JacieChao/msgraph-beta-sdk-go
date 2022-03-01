package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2bb7d0bba9f4dacf992341819d973086e188fe4bac69df578a46624982bc73af "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsqualityupdateprofiles/item/assignments"
    i36baea6ccabeba8dfb103008171c914e98a40320eea881bf06c1729f50cb81ce "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsqualityupdateprofiles/item/assign"
    i04d128338563166ddbb410fea932b830e3e2b8799a8d363b4dbb0797203fe43d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsqualityupdateprofiles/item/assignments/item"
)

// WindowsQualityUpdateProfileItemRequestBuilder builds and executes requests for operations under \deviceManagement\windowsQualityUpdateProfiles\{windowsQualityUpdateProfile-id}
type WindowsQualityUpdateProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsQualityUpdateProfileItemRequestBuilderDeleteOptions options for Delete
type WindowsQualityUpdateProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsQualityUpdateProfileItemRequestBuilderGetOptions options for Get
type WindowsQualityUpdateProfileItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsQualityUpdateProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsQualityUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows quality update profiles
type WindowsQualityUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsQualityUpdateProfileItemRequestBuilderPatchOptions options for Patch
type WindowsQualityUpdateProfileItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsQualityUpdateProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsQualityUpdateProfileItemRequestBuilder) Assign()(*i36baea6ccabeba8dfb103008171c914e98a40320eea881bf06c1729f50cb81ce.AssignRequestBuilder) {
    return i36baea6ccabeba8dfb103008171c914e98a40320eea881bf06c1729f50cb81ce.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsQualityUpdateProfileItemRequestBuilder) Assignments()(*i2bb7d0bba9f4dacf992341819d973086e188fe4bac69df578a46624982bc73af.AssignmentsRequestBuilder) {
    return i2bb7d0bba9f4dacf992341819d973086e188fe4bac69df578a46624982bc73af.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.windowsQualityUpdateProfiles.item.assignments.item collection
func (m *WindowsQualityUpdateProfileItemRequestBuilder) AssignmentsById(id string)(*i04d128338563166ddbb410fea932b830e3e2b8799a8d363b4dbb0797203fe43d.WindowsQualityUpdateProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsQualityUpdateProfileAssignment_id"] = id
    }
    return i04d128338563166ddbb410fea932b830e3e2b8799a8d363b4dbb0797203fe43d.NewWindowsQualityUpdateProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsQualityUpdateProfileItemRequestBuilderInternal instantiates a new WindowsQualityUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsQualityUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsQualityUpdateProfileItemRequestBuilder) {
    m := &WindowsQualityUpdateProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsQualityUpdateProfiles/{windowsQualityUpdateProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsQualityUpdateProfileItemRequestBuilder instantiates a new WindowsQualityUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsQualityUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsQualityUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsQualityUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of windows quality update profiles
func (m *WindowsQualityUpdateProfileItemRequestBuilder) CreateDeleteRequestInformation(options *WindowsQualityUpdateProfileItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of windows quality update profiles
func (m *WindowsQualityUpdateProfileItemRequestBuilder) CreateGetRequestInformation(options *WindowsQualityUpdateProfileItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of windows quality update profiles
func (m *WindowsQualityUpdateProfileItemRequestBuilder) CreatePatchRequestInformation(options *WindowsQualityUpdateProfileItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of windows quality update profiles
func (m *WindowsQualityUpdateProfileItemRequestBuilder) Delete(options *WindowsQualityUpdateProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of windows quality update profiles
func (m *WindowsQualityUpdateProfileItemRequestBuilder) Get(options *WindowsQualityUpdateProfileItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsQualityUpdateProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsQualityUpdateProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsQualityUpdateProfile), nil
}
// Patch a collection of windows quality update profiles
func (m *WindowsQualityUpdateProfileItemRequestBuilder) Patch(options *WindowsQualityUpdateProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}