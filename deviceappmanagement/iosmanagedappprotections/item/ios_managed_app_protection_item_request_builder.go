package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0217d7e2ce763c77314148f46efea694dae433696df6faeb30b3f32c12f8c6f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/iosmanagedappprotections/item/apps"
    i96bdcd6a8e15330fdcfeeaf87c61073aae0d06ba66a090f0507d66c527796caa "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/iosmanagedappprotections/item/deploymentsummary"
    ia68797db34f5d00a9aec277296562650717d30a23a6d00f131196cb025ff431d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/iosmanagedappprotections/item/apps/item"
)

// IosManagedAppProtectionItemRequestBuilder builds and executes requests for operations under \deviceAppManagement\iosManagedAppProtections\{iosManagedAppProtection-id}
type IosManagedAppProtectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IosManagedAppProtectionItemRequestBuilderDeleteOptions options for Delete
type IosManagedAppProtectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosManagedAppProtectionItemRequestBuilderGetOptions options for Get
type IosManagedAppProtectionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IosManagedAppProtectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosManagedAppProtectionItemRequestBuilderGetQueryParameters iOS managed app policies.
type IosManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IosManagedAppProtectionItemRequestBuilderPatchOptions options for Patch
type IosManagedAppProtectionItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosManagedAppProtection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IosManagedAppProtectionItemRequestBuilder) Apps()(*i0217d7e2ce763c77314148f46efea694dae433696df6faeb30b3f32c12f8c6f5.AppsRequestBuilder) {
    return i0217d7e2ce763c77314148f46efea694dae433696df6faeb30b3f32c12f8c6f5.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosManagedAppProtections.item.apps.item collection
func (m *IosManagedAppProtectionItemRequestBuilder) AppsById(id string)(*ia68797db34f5d00a9aec277296562650717d30a23a6d00f131196cb025ff431d.ManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp_id"] = id
    }
    return ia68797db34f5d00a9aec277296562650717d30a23a6d00f131196cb025ff431d.NewManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIosManagedAppProtectionItemRequestBuilderInternal instantiates a new IosManagedAppProtectionItemRequestBuilder and sets the default values.
func NewIosManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosManagedAppProtectionItemRequestBuilder) {
    m := &IosManagedAppProtectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosManagedAppProtectionItemRequestBuilder instantiates a new IosManagedAppProtectionItemRequestBuilder and sets the default values.
func NewIosManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation iOS managed app policies.
func (m *IosManagedAppProtectionItemRequestBuilder) CreateDeleteRequestInformation(options *IosManagedAppProtectionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation iOS managed app policies.
func (m *IosManagedAppProtectionItemRequestBuilder) CreateGetRequestInformation(options *IosManagedAppProtectionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation iOS managed app policies.
func (m *IosManagedAppProtectionItemRequestBuilder) CreatePatchRequestInformation(options *IosManagedAppProtectionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete iOS managed app policies.
func (m *IosManagedAppProtectionItemRequestBuilder) Delete(options *IosManagedAppProtectionItemRequestBuilderDeleteOptions)(error) {
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
func (m *IosManagedAppProtectionItemRequestBuilder) DeploymentSummary()(*i96bdcd6a8e15330fdcfeeaf87c61073aae0d06ba66a090f0507d66c527796caa.DeploymentSummaryRequestBuilder) {
    return i96bdcd6a8e15330fdcfeeaf87c61073aae0d06ba66a090f0507d66c527796caa.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get iOS managed app policies.
func (m *IosManagedAppProtectionItemRequestBuilder) Get(options *IosManagedAppProtectionItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosManagedAppProtection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIosManagedAppProtection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosManagedAppProtection), nil
}
// Patch iOS managed app policies.
func (m *IosManagedAppProtectionItemRequestBuilder) Patch(options *IosManagedAppProtectionItemRequestBuilderPatchOptions)(error) {
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