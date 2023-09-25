package registryclient

import _ "unsafe"

type Embedme interface{}

type anonymuskc struct{}

type azurekeychain struct{}

type gcpkeychain struct{}

type autoRefreshSecrets struct {
	lister           interface{}
	imagePullSecrets interface{}
}

type awskeychain struct{}

func NewAutoRefreshSecretsKeychain(lister interface{}, imagePullSecrets interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (kc *autoRefreshSecrets) Resolve(resource interface{}) (interface{}, interface{}) {
	panic("stub")
}

func isACRRegistry(input interface{}) interface{} {
	panic("stub")
}

func isAWSRegistry(input interface{}) interface{} {
	panic("stub")
}

type Client struct {
	FetchImageDescriptor struct{}
}

type config struct {
	keychain  interface{}
	transport interface{}
	tracing   interface{}
}

type client struct {
	keychain  interface{}
	transport interface{}
}

type Option func(interface{}) interface{}

func New(options interface{}) (interface{}, interface{}) {
	panic("stub")
}

func NewOrDie(options interface{}) interface{} {
	panic("stub")
}

func WithKeychainPullSecrets(lister interface{}, imagePullSecrets interface{}) interface{} {
	panic("stub")
}

func WithCredentialProviders(credentialProviders interface{}) interface{} {
	panic("stub")
}

func WithAllowInsecureRegistry() interface{} {
	panic("stub")
}

func WithLocalKeychain() interface{} {
	panic("stub")
}

func WithTracing() interface{} {
	panic("stub")
}

func (c *client) BuildRemoteOption(ctx interface{}) interface{} {
	panic("stub")
}

func (c *client) FetchImageDescriptor(ctx interface{}, imageRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (c *client) Keychain() interface{} {
	panic("stub")
}

func (c *client) getTransport() interface{} {
	panic("stub")
}

func generateKeychainForPullSecrets(lister interface{}, imagePullSecrets interface{}) (interface{}, interface{}) {
	panic("stub")
}
