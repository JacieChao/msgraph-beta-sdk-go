package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder provides operations to call the renewUpload method.
type MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilderInternal instantiates a new RenewUploadRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder) {
    m := &MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedAndroidLobApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}/renewUpload", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder instantiates a new RenewUploadRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renews the SAS URI for an application file upload.
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder) Post(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation renews the SAS URI for an application file upload.
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder) {
    return NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemFilesItemRenewUploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
