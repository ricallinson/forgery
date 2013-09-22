package f

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestStack(t *testing.T) {

    Describe("defaultConfiguration()", func() {

        It("should return [1]", func() {
            s := &Server{
                settings: map[string]string{},
            }
            s.defaultConfiguration()
        })
    })
}