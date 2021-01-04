package terminus

type AzBlobTerminusConfig struct {
	ConnString string	`yaml:"connstring"`
	Container string	`yaml:"containerName"`
}