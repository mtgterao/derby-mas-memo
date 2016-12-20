package libraries

import (
    "fmt"
    "testing"
    "github.com/mtgterao/derby-mas-memo/libraries"
)

func TestMergeMap(t *testing.T) {
    map1 := map[string]interface{}{
        "key1": "value1",
        "key4": "value4",
    }

    map2 := map[string]interface{}{
        "key2": "value2",
    }

    map3 := map[string]interface{}{
        "key1": "value3",
    }

    // テスト結果予定のマップ
    mapAssume := map[string]string{
        "key1": "value3",
        "key2": "value2",
        "key4": "value4",
    }

    result := libraries.MergeMap(map1, map2, map3)

    // map長チェック
    if len(result) != 3 {
        t.Error(fmt.Sprintf("map length expected to differ. [len=%d]", len(result)))
    }

    checkAssume(t, mapAssume, result)
}

func TestMergeMapWithNil(t *testing.T) {
    map1 := map[string]interface{}{
        "key1": "value1",
        "key4": "value4",
    }

    var map2 map[string]interface{}

    map3 := map[string]interface{}{
        "key1": "value3",
    }

    // テスト結果予定のマップ
    mapAssume := map[string]string{
        "key1": "value3",
        "key4": "value4",
    }

    result := libraries.MergeMap(map1, map2, map3)

    if len(result) != 2 {
        t.Error(fmt.Sprintf("map length expected to differ. [len=%d]", len(result)))
    }

    checkAssume(t, mapAssume, result)
}

func checkAssume(t *testing.T, assume map[string]string, result map[string]interface{}) {
    for key, value := range result {

        valueAssume, ok := assume[key]

        if !ok {
            t.Error(fmt.Sprintf("containts invalid key. [key=%v]", key))
        } else {
            if value != valueAssume {
                t.Error(fmt.Sprintf("this value is different assume. [key=%v, value=%v]", key, value))
            }
        }
    }
}
