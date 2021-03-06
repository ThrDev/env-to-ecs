package _unittests

import (
	"7factor.io/converter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Constants for testing.
const EmptyEnvironmentArray = `[]`
const SingleLineSingleInput = `A=B`
const MultiLineInput = `
A=B
D=E
`
const MultiLineWithSpacesInput = `

L=M

N=O
`
const SingleLineMultiInput = `W=X   Y=Z`
const InputWithSpaces = `X = 1 Y = 1`
const illegalJSONInput = `A=\"`
const MultiLineWithComments = `
# this is a comment
Q=R
S=T
`

// Base64 strings always end with an equals, handle this case.
const InputWithEquals = `A=abcdefg=`
const MultiLineInputWithEquals = `
A=abcdefg=
B=C
`
const InputWithHash = `WITHHASH=#FOO#`

const InputWithQuotes = `WITHQUOTES="this is a test"`
const InputWithQuotesAndSpaces = `A = "test string"`

const CrazyInput = `A=1 B = 2 C="test string" D = "another test string" E="1" F = "2"
G = "another test string"
H="test string"`

var _ = Describe("The ECS converter", func() {
	Context("When passed a blank file", func() {
		It("Returns an empty JSON blob", func() {
			converted, err := converter.ConvertInputToJson("")
			Expect(err).ToNot(BeNil())
			Expect(converted).To(Equal(EmptyEnvironmentArray))
		})
	})

	Context("When passed a single line file without comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(SingleLineSingleInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"}]`))
		})
	})

	Context("When passed a single line file with multiple values and spaces between the equals", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(InputWithSpaces)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"X","value":"1"},{"name":"Y","value":"1"}]`))
		})
	})

	Context("When passed a single line file with quotes and spaces", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(InputWithQuotesAndSpaces)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"test string"}]`))
		})
	})

	Context("When passed a multi-line file with newlines in between", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(MultiLineInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"B"},{"name":"D","value":"E"}]`))
		})
	})

	Context("When passed a multi-line file with spaces in between", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(MultiLineWithSpacesInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"L","value":"M"},{"name":"N","value":"O"}]`))
		})
	})

	Context("When passed a single line file with multiple items", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(SingleLineMultiInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"W","value":"X"},{"name":"Y","value":"Z"}]`))
		})
	})

	Context("When passed characters that cause the JSON translation to fail", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(illegalJSONInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"\\\""}]`))
		})
	})

	Context("When passed a multi-line file with comments", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(MultiLineWithComments)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"Q","value":"R"},{"name":"S","value":"T"}]`))
		})
	})

	Context("When passed a file with `#` in the `name` and `value` params and is otherwise valid", func() {
		It("Returns the expected JSON blob", func() {
			converted, err := converter.ConvertInputToJson(InputWithHash)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"WITHHASH","value":"#FOO#"}]`))
		})
	})

	Context("When called with a string variable with an equals sign in it", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson(InputWithEquals)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"abcdefg="}]`))
		})
	})

	Context("When called with a string variable with an equals sign in it and multiple variables", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson(MultiLineInputWithEquals)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"abcdefg="},{"name":"B","value":"C"}]`))
		})
	})

	Context("When called with a quoted string variable", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson(InputWithQuotes)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"WITHQUOTES","value":"this is a test"}]`))
		})
	})

	Context("When called with a string with multiple challenges", func() {
		It("Works as expected and doesn't throw an index out of range exception.", func() {
			converted, err := converter.ConvertInputToJson(CrazyInput)
			Expect(err).To(BeNil())
			Expect(converted).To(Equal(`[{"name":"A","value":"1"},{"name":"B","value":"2"},{"name":"C","value":"test string"},{"name":"D","value":"another test string"},{"name":"E","value":"1"},{"name":"F","value":"2"},{"name":"G","value":"another test string"},{"name":"H","value":"test string"}]`))
		})
	})
})
