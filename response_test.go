package f

import(
    // "fmt"
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestResponse(t *testing.T) {

    Describe("_()", func() {

        It("should return []", func() {
            AssertEqual(true, true)
        })
    })

    Report(t)
}