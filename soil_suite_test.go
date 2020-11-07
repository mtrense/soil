package soil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSoil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Soil Suite")
}
