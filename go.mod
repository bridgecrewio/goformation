module github.com/bridgecrewio/goformation/v5

require (
	github.com/awslabs/goformation/v5 v5.2.7
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.12.0
	github.com/sanathkr/go-yaml v0.0.0-20170819195128-ed9d249f429b
	github.com/sanathkr/yaml v0.0.0-20170819201035-0056894fa522
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v0.0.0-20181112162635-ac52e6811b56
)

replace github.com/awslabs/goformation/v5 => github.com/bridgecrewio/goformation/v5 v5.0.0-20210823081757-99ed9bf3c0e5

go 1.13
