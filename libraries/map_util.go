/**
 * map系のutility
 *
 */
package libraries

// 複数のマップをマージする（keyはstringのみ）
func MergeMap(target ...map[string]interface{}) map[string]interface{} {
    merged := map[string]interface{}{}

    for _, slice := range target {

        for key, value := range slice {
            merged[key] = value
        }
    }

    return merged
}
