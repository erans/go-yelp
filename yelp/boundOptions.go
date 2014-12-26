package yelp

import (
	"fmt"
)

/**
 * Location is specified by a bounding box, defined by a southwest latitude/longitude and a northeast latitude/longitude geographic coordinate.
 * The bounding box format is defined as:
 * bounds=sw_latitude,sw_longitude|ne_latitude,ne_longitude
 */
type BoundOptions struct {
	Sw_latitude  float32 //	Southwest latitude of bounding box
	Sw_longitude float32 //	Southwest longitude of bounding box
	Ne_latitude  float32 //	Northeast latitude of bounding box
	Ne_longitude float32 //	Northeast longitude of bounding box
}

func (o BoundOptions) GetParameters() (params map[string]string, err error) {
	return map[string]string{
		"bounds": fmt.Sprintf("%v,%v|%v,%v",
			o.Sw_latitude,
			o.Sw_longitude,
			o.Ne_latitude,
			o.Ne_longitude,
		),
	}, nil
}
