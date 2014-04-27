package sphero

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpheroDriver", func() {
	var (
		s *Sphero
		a *SpheroAdaptor
	)

	BeforeEach(func() {
		a = NewSpheroAdaptor()
		a.sp = sp{}
		s = NewSphero(a)
	})

	It("Must be able to Start", func() {
		Expect(s.Start()).To(Equal(true))
	})
	It("Must be able to Init", func() {
		Expect(s.Init()).To(Equal(true))
	})
	It("Must be able to Halt", func() {
		Expect(s.Halt()).To(Equal(true))
	})
})
