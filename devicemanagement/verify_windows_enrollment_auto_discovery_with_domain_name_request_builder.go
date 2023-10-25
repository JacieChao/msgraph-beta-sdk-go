package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, domainName *string)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    m := &VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')", pathParameters),
    }
    if domainName != nil {
        m.BaseRequestBuilder.PathParameters["domainName"] = *domainName
    }
    return m
}
// NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function verifyWindowsEnrollmentAutoDiscovery
// Deprecated: This method is obsolete. Use GetAsVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponse instead.
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) Get(ctx context.Context, requestConfiguration *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable), nil
}
// GetAsVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponse invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) GetAsVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponse(ctx context.Context, requestConfiguration *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponseable), nil
}
// ToGetRequestInformation invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) WithUrl(rawUrl string)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
