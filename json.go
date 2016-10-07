package json

import (
    "encoding/json"
    "io/ioutil"
    "strings"
    "strconv"
    "fmt"
)

// Hash : A helper struct which will have methods hanging off of it for ease of access
type Hash struct {
    data map[string]interface{}
}

// Get : Return the specified value based on the key(s) provided
func (h *Hash) Get(key string) (interface{}) {
    keys := strings.Split(key, ".")
    keyCount := len(keys)
    
    // if we only have one key and we're not going any deeper then just return this straight off the data map
    if keyCount == 1 {
        return h.data[keys[0]]
    }

    // otherwise if we're going in deeper than just set the first value to a map so we can recurse through it
    array := h.data[keys[0]].(map[string]interface{})
    for i, k := range keys {

        // skip the first one as we've already assigned it as the parent (array)
        if i == 0 {
            continue
        }
        
        // if going in further still than just set the new parent to array
        if keyCount - 1 > i {
            array = array[k].(map[string]interface{})
            continue
        }

        // handle floats and integers properly and return the correctly typed array
        item := array[k]
        if fmt.Sprintf("%T", item) == "float64" {
            stringedNumber := strconv.FormatFloat(item.(float64), 'g', -1, 64)
            stringPieces := strings.Split(stringedNumber, ".")

            // it's actually an integer so cast and return that
            if len(stringPieces) == 1 {
                return int64(item.(float64))
            }
        }

        return item
    }

    return nil
}

// FromFile : Read JSON object from a file and return a hash JSON object
func FromFile(file string) (*Hash, error) {
    bytes, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, err
    }

    hash := new(Hash)
    if err := json.Unmarshal(bytes, &hash.data); err != nil {
        return nil, err
    }
    return hash, nil
}

// FromBytes : Read JSON object from a byte array and return a hash JSON object
func FromBytes(bytes []byte) (*Hash, error) {
    hash := new(Hash)
    if err := json.Unmarshal(bytes, &hash.data); err != nil {
        return nil, err
    }
    return hash, nil
}