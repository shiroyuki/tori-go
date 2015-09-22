package re

import "fmt"
import "testing"
import "../tameshigiri"

func localTestReSearchOne(
    t            *testing.T,
    pattern      string,
    sample       string,
    expectedList []string,
    expectedDict map[string]string,
) {
    var index    int
    var key      string
    var actual   *string
    var expected string

    var assertion = tameshigiri.NewAssertion(t)
    var compiled  = Compile(pattern)
    var result    = compiled.SearchOne(sample)

    var expectedListLength = len(expectedList)
    var expectedDictLength = len(expectedDict)
    var expectedTotalCount = expectedListLength + expectedDictLength

    if expectedTotalCount > 0 {
        assertion.IsTrue(result.HasAny(), "This should be found.")
    }

    assertion.Equals(
        expectedTotalCount,
        result.Count(),
        "Total count",
    )

    assertion.Equals(
        expectedListLength,
        result.CountIndices(),
        "Item-list count",
    )

    if expectedListLength > 0 {
        for index = range expectedList {
            actual   = result.Index(index)
            expected = expectedList[index]

            assertion.Equals(
                expected,
                *actual,
                fmt.Sprintf("List#%d", index),
            )
        }
    }

    assertion.Equals(
        expectedDictLength,
        result.CountKeys(),
        "Dictionary count",
    )

    if expectedDictLength > 0 {
        for key, expected = range expectedDict {
            actual = result.Key(key)

            assertion.Equals(
                expected,
                *actual,
                fmt.Sprintf("Dictionary[%s]", key),
            )
        }
    }
}

func localTestReSearchAll(
    t            *testing.T,
    pattern      string,
    sample       string,
    //expectedList []string,
    //expectedDict map[string][]string,
) {
    // var index    int
    // var key      string
    // var actual   *string
    // var expected string

    // var assertion = tameshigiri.NewAssertion(t)
    var compiled  = Compile(pattern)
    var result    = compiled.SearchAll(sample)

    fmt.Println(result.Dictionary)

    // var expectedListLength = len(expectedList)
    // var expectedDictLength = len(expectedDict)
    // var expectedTotalCount = expectedListLength + expectedDictLength
    //
    // if expectedTotalCount > 0 {
    //     assertion.IsTrue(result.HasAny(), "This should be found.")
    // }
    //
    // assertion.Equals(
    //     expectedTotalCount,
    //     result.Count(),
    //     "Total count",
    // )
    //
    // assertion.Equals(
    //     expectedListLength,
    //     result.CountIndices(),
    //     "Item-list count",
    // )
    //
    // if expectedListLength > 0 {
    //     for index = range expectedList {
    //         actual   = result.Index(index)
    //         expected = expectedList[index]
    //
    //         assertion.Equals(
    //             expected,
    //             *actual,
    //             fmt.Sprintf("List#%d", index),
    //         )
    //     }
    // }
    //
    // assertion.Equals(
    //     expectedDictLength,
    //     result.CountKeys(),
    //     "Dictionary count",
    // )
    //
    // if expectedDictLength > 0 {
    //     for key, expected = range expectedDict {
    //         actual = result.Key(key)
    //
    //         assertion.Equals(
    //             expected,
    //             *actual,
    //             fmt.Sprintf("Dictionary[%s]", key),
    //         )
    //     }
    // }
}

func TestReSearchOneBasic(t *testing.T) {
    localTestReSearchOne(
        t,
        "shiroyuki",
        "user-01-shiroyuki-as-admin",
        []string{ "shiroyuki" },
        make(map[string]string, 0),
    )
}

func TestReSearchOneAdvanced(t *testing.T) {
    var testPattern = "(/api)?/users/(?P<key>[^/]+)"

    localTestReSearchOne(
        t,
        testPattern,
        "/users",
        make([]string, 0),
        make(map[string]string, 0),
    )

    localTestReSearchOne(
        t,
        testPattern,
        "/users/shiroyuki",
        []string{
            "/users/shiroyuki",
            "",
        },
        map[string]string{
            "key": "shiroyuki",
        },
    )

    localTestReSearchOne(
        t,
        testPattern,
        "/api/users/shiroyuki",
        []string{
            "/api/users/shiroyuki",
            "/api",
        },
        map[string]string{
            "key": "shiroyuki",
        },
    )
}

func TestReSearchAll(t *testing.T) {
    var pattern = "<(?P<name>[^>]+)>"
    var sample  = "/api/v1/<abc>/<def>"

    localTestReSearchAll(
        t,
        pattern,
        sample,
    )
}
