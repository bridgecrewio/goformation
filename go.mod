module github.com/bridgecrewio/goformation/v5

require (
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.12.0
	github.com/sanathkr/go-yaml v0.0.0-20170819195128-ed9d249f429b
	github.com/sanathkr/yaml v0.0.0-20170819201035-0056894fa522
	github.com/xeipuuv/gojsonschema v0.0.0-20181112162635-ac52e6811b56
)

replace github.com/awslabs/goformation/v5 => github.com/bridgecrewio/goformation/v5 master

go 1.13
