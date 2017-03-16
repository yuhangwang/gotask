package convert


func MapLike(obj interface{}) interface{} {
    switch x := obj.(type) {
        case map[interface{}]interface{}:
            new_map := map[string]interface{}{}
            for k, v := range x {
                new_map[k.(string)] = MapLike(v)
            }
            return new_map
        case []interface{}:
            for i, v := range x {
                x[i] = MapLike(v)
            }
    }
    return obj
}
