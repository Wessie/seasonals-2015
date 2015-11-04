package conf

import (
	"encoding"
	"fmt"
	"sync"
)

type Configuration interface {
	// Default initializes the instance with default values
	Default()

	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

var ConfigDatabase = DatabaseFunc(ConfigurationFilename)

var (
	// needs to be hold when accessing ConfigMap
	ConfigMapLock sync.Mutex
	ConfigMap     = map[string]Configuration{}
	uniquePanic   = `non-unique configuration name: %s
      have argument: %v
  previous argument: %v`
)

// Register registers the Configuration to the name given, Configuration
// will have its contents loaded from persistent storage. After calling
// Register a Configuration should never be mutated by others.
func Register(name string, c Configuration) {
	// TODO: make sure c is a settable pointer-type
	// TODO: panic error messages should be clear and to-the-point
	ConfigMapLock.Lock()
	defer ConfigMapLock.Unlock()

	// check if the name is actually unique
	if c1, ok := ConfigMap[name]; ok {
		panic(fmt.Sprintf(uniqErr, name, c, c1))
	}
	ConfigMap[name] = c

	err := ConfigDatabase().Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			return err
		}

		bName := []byte(name)

		v := b.Get(bName)
		if v == nil {
			c.Default()
			nv, err := c.MarshalText()
			if err != nil {
				panic("invalid default value: " + err.Error())
			}
			return b.Put(bName, nv)
		}

		// TODO: handle broken input, keep track of c and v, ask for defaultification
		// if possible from the user.
		return c.UnmarshalText(v)
	})
	// TODO: differentiate between database and Unmarshal error, database is critical,
	// unmarshal requires user input.
	if err != nil {
		panic("fail" + err.Error())
	}

	return
}
