package sphero

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGobotSphero(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sphero Suite")
}
