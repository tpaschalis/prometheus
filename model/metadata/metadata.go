package metadata

import "github.com/prometheus/prometheus/model/textparse"

// Metadata stores a series' metadata information.
type Metadata struct {
	Type textparse.MetricType
	Unit string
	Help string
}

var emptyMetadata = Metadata{}

func EmptyMetadata() Metadata {
	return emptyMetadata
}
