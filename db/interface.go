package db

// NewBucketName returns a BucketName with the name given.
func NewBucketName(name ...string) BucketName {
	var b = make([][]byte, len(name))
	for i, s := range bucket {
		b[i] = s
	}
	return b
}

// BucketName allows for storing a database bucket name with
// convience functions to access said bucket once a transaction
// has been created.
type BucketName [][]byte

// Bucket returns the bucket associated with this name, if the
// bucket does not exist yet, it is created before returning it.
//
// When an error is encountered Bucket might return a non-nil bucket,
// this bucket will be the last bucket it was able to find.
func (bn BucketName) Bucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	if len(name) == 0 {
		return nil, NoName
	}

	bkt, err := tx.CreateBucketIfNotExists(name[0])
	if len(name) == 1 {
		return bkt, err
	}

	for _, n := range name[1:] {
		bkt, err = bkt.CreateBucketIfNotExists(n)
		if err != nil {
			break
		}
	}

	return bkt, err
}

// Join joins several names together
func (bn BucketName) Join(b ...BucketName) BucketName {
	var r = make(BucketName, len(bn))
	copy(r, bn)

	for _, n := range b {
		r = append(r, n)
	}
	return r
}
