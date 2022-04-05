package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/stop"
    i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/provisionondemand"
    i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/restart"
    i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/pause"
    i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/schema"
    idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/start"
    ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/validatecredentials"
)

// SynchronizationJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
type SynchronizationJobItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SynchronizationJobItemRequestBuilderDeleteOptions options for Delete
type SynchronizationJobItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SynchronizationJobItemRequestBuilderGetOptions options for Get
type SynchronizationJobItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *SynchronizationJobItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SynchronizationJobItemRequestBuilderGetQueryParameters performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
type SynchronizationJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SynchronizationJobItemRequestBuilderPatchOptions options for Patch
type SynchronizationJobItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSynchronizationJobItemRequestBuilderInternal instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    m := &SynchronizationJobItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application_id}/synchronization/jobs/{synchronizationJob_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSynchronizationJobItemRequestBuilder instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property jobs for applications
func (m *SynchronizationJobItemRequestBuilder) CreateDeleteRequestInformation(options *SynchronizationJobItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *SynchronizationJobItemRequestBuilder) CreateGetRequestInformation(options *SynchronizationJobItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property jobs in applications
func (m *SynchronizationJobItemRequestBuilder) CreatePatchRequestInformation(options *SynchronizationJobItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property jobs for applications
func (m *SynchronizationJobItemRequestBuilder) Delete(options *SynchronizationJobItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *SynchronizationJobItemRequestBuilder) Get(options *SynchronizationJobItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationJobFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable), nil
}
// Patch update the navigation property jobs in applications
func (m *SynchronizationJobItemRequestBuilder) Patch(options *SynchronizationJobItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Pause the pause property
func (m *SynchronizationJobItemRequestBuilder) Pause()(*i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f.PauseRequestBuilder) {
    return i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisionOnDemand the provisionOnDemand property
func (m *SynchronizationJobItemRequestBuilder) ProvisionOnDemand()(*i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f.ProvisionOnDemandRequestBuilder) {
    return i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f.NewProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restart the restart property
func (m *SynchronizationJobItemRequestBuilder) Restart()(*i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef.RestartRequestBuilder) {
    return i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema the schema property
func (m *SynchronizationJobItemRequestBuilder) Schema()(*i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967.SchemaRequestBuilder) {
    return i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Start the start property
func (m *SynchronizationJobItemRequestBuilder) Start()(*idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237.StartRequestBuilder) {
    return idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop the stop property
func (m *SynchronizationJobItemRequestBuilder) Stop()(*i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c.StopRequestBuilder) {
    return i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateCredentials the validateCredentials property
func (m *SynchronizationJobItemRequestBuilder) ValidateCredentials()(*ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90.ValidateCredentialsRequestBuilder) {
    return ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90.NewValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
